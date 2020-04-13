package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/thought-machine/dracon/consumers"
	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Consumer struct {
	Client     HTTPClient
	Webhook    string
	LongFormat bool
}

func main() {
	var cons = &Consumer{}

	if err := cons.parseFlags(); err != nil {
		log.Fatal(err)
	}

	if consumers.Raw {
		responses, err := consumers.LoadToolResponse()
		if err != nil {
			log.Fatal(err)
		}
		cons.processRawMessages(responses)
	} else {
		responses, err := consumers.LoadEnrichedToolResponse()
		if err != nil {
			log.Fatal(err)
		}
		cons.processEnrichedMessages(responses)
	}
}

func (c *Consumer) init() {
	flag.StringVar(&c.Webhook, "Webhook", "", "the Webhook to push results to")
	flag.BoolVar(&c.LongFormat, "long", true, "post the full results to Webhook, not just metrics")
	c.Client = &http.Client{}
}

func (c *Consumer) parseFlags() error {
	if err := consumers.ParseFlags(); err != nil {
		return err
	}
	if len(c.Webhook) < 1 {
		return fmt.Errorf("Webhook is undefined")
	}
	return nil
}

func (c *Consumer) getRawIssue(scanStartTime time.Time, res *v1.LaunchToolResponse, iss *v1.Issue) ([]byte, error) {
	jBytes, err := json.Marshal(&fullDocument{
		ScanStartTime: scanStartTime,
		ScanID:        res.GetScanInfo().GetScanUuid(),
		ToolName:      res.GetToolName(),
		Source:        iss.GetSource(),
		Title:         iss.GetTitle(),
		Target:        iss.GetTarget(),
		Type:          iss.GetType(),
		Severity:      iss.GetSeverity(),
		CVSS:          iss.GetCvss(),
		Confidence:    iss.GetConfidence(),
		Description:   iss.GetDescription(),
		FirstFound:    scanStartTime,
		FalsePositive: false,
	})
	if err != nil {
		return []byte{}, err
	}
	return jBytes, nil
}

func (c *Consumer) getEnrichedIssue(scanStartTime time.Time, res *v1.EnrichedLaunchToolResponse, iss *v1.EnrichedIssue) ([]byte, error) {
	firstSeenTime, _ := ptypes.Timestamp(iss.GetFirstSeen())
	jBytes, err := json.Marshal(&fullDocument{
		ScanStartTime: scanStartTime,
		ScanID:        res.GetOriginalResults().GetScanInfo().GetScanUuid(),
		ToolName:      res.GetOriginalResults().GetToolName(),
		Source:        iss.GetRawIssue().GetSource(),
		Title:         iss.GetRawIssue().GetTitle(),
		Target:        iss.GetRawIssue().GetTarget(),
		Type:          iss.GetRawIssue().GetType(),
		Severity:      iss.GetRawIssue().GetSeverity(),
		CVSS:          iss.GetRawIssue().GetCvss(),
		Confidence:    iss.GetRawIssue().GetConfidence(),
		Description:   iss.GetRawIssue().GetDescription(),
		FirstFound:    firstSeenTime,
		FalsePositive: iss.GetFalsePositive(),
	})
	if err != nil {
		return []byte{}, err
	}
	return jBytes, nil
}

type fullDocument struct {
	ScanStartTime time.Time     `json:"scan_start_time"`
	ScanID        string        `json:"scan_id"`
	ToolName      string        `json:"tool_name"`
	Source        string        `json:"source"`
	Target        string        `json:"target"`
	Type          string        `json:"type"`
	Title         string        `json:"title"`
	Severity      v1.Severity   `json:"severity"`
	CVSS          float64       `json:"cvss"`
	Confidence    v1.Confidence `json:"confidence"`
	Description   string        `json:"description"`
	FirstFound    time.Time     `json:"first_found"`
	FalsePositive bool          `json:"false_positive"`
}

func (c *Consumer) processRawMessages(responses []*v1.LaunchToolResponse) {
	if c.LongFormat == false {
		scanStartTime, _ := ptypes.Timestamp(responses[0].GetScanInfo().GetScanStartTime())
		c.pushMetrics(responses[0].GetScanInfo().GetScanUuid(), c.countRawMessages(responses), scanStartTime)
		return
	}
	for _, res := range responses {
		scanStartTime, _ := ptypes.Timestamp(res.GetScanInfo().GetScanStartTime())
		for _, iss := range res.GetIssues() {
			b, err := c.getRawIssue(scanStartTime, res, iss)
			if err != nil {
				log.Fatal(err)
			}
			c.push(string(b))
		}
	}
}
func (c *Consumer) processEnrichedMessages(responses []*v1.EnrichedLaunchToolResponse) {
	if c.LongFormat == false {
		scanStartTime, _ := ptypes.Timestamp(responses[0].GetOriginalResults().GetScanInfo().GetScanStartTime())
		c.pushMetrics(responses[0].GetOriginalResults().GetScanInfo().GetScanUuid(), c.countEnrichedMessages(responses), scanStartTime)
		return
	}
	for _, res := range responses {
		scanStartTime, _ := ptypes.Timestamp(res.GetOriginalResults().GetScanInfo().GetScanStartTime())
		for _, iss := range res.GetIssues() {
			b, err := c.getEnrichedIssue(scanStartTime, res, iss)
			if err != nil {
				log.Fatal(err)
			}
			c.push(string(b))
		}
	}
}

func (c *Consumer) countRawMessages(responses []*v1.LaunchToolResponse) int {
	result := 0
	for _, res := range responses {
		result += len(res.GetIssues())
	}
	return result
}

func (c *Consumer) countEnrichedMessages(responses []*v1.EnrichedLaunchToolResponse) int {
	result := 0
	for _, res := range responses {
		result += len(res.GetIssues())
	}
	return result
}
func (c *Consumer) pushMetrics(scanUUID string, issuesNo int, scanStartTime time.Time) {
	message := fmt.Sprintf("Dracon scan %s started on %s has been completed with %d issues\n", scanUUID, scanStartTime, issuesNo)
	c.push(message)
}

func (c *Consumer) push(b string) error {
	type SlackRequestBody struct {
		Text string `json:"text"`
	}
	var err error
	body, _ := json.Marshal(SlackRequestBody{Text: b})
	req, err := http.NewRequest(http.MethodPost, c.Webhook, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	fmt.Println(c.Client)

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	fmt.Println(resp)
	fmt.Println(resp.Body)

	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}