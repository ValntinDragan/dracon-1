package gojira

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/andygrunwald/go-jira"
	"github.com/trivago/tgo/tcontainer"
	yaml "gopkg.in/yaml.v2"
)

const (
	// EnvConfigPath the path towards the config.yaml file
	EnvConfigPath = "DRACON_JIRA_CONFIG_PATH"
)

type goJiraClient struct {
	JiraClient *jira.Client
	DryRunMode bool
	Configs    config
}

type customField struct {
	ID        string   `yaml:"id"`
	FieldType string   `yaml:"fieldType"`
	Values    []string `yaml:"values"`
}

type defaultValues struct {
	IssueFields  map[string][]string `yaml:"issueFields,omitempty"`
	CustomFields []customField       `yaml:"customFields,omitempty"`
}

type mappings struct {
	DraconField string `yaml:"draconField"`
	JiraField   string `yaml:"jiraField"`
	FieldType   string `yaml:"fieldType"`
}

type config struct {
	DefaultValues     defaultValues `yaml:"defaultValues"`
	Mappings          []mappings    `yaml:"mappings"`
	DescriptionExtras []string      `yaml:"addToDescription"`
}

// NewGoJiraClient returns a goJiraClient containing the authentication details and the configuration settings
func NewGoJiraClient(user, token, url string, dryRun bool) *goJiraClient {
	return &goJiraClient{
		JiraClient: authJiraClient(user, token, url),
		DryRunMode: dryRun,
		Configs:    getConfig(),
	}
}

// getConfig parses the configuration from config.yaml and returns a config struct
func getConfig() config {
	configFile, err := ioutil.ReadFile(os.Getenv(EnvConfigPath))
	if err != nil {
		log.Printf("Error while reading config file:   #%v ", err)
	}

	var newConfig config
	err = yaml.Unmarshal(configFile, &newConfig)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return newConfig
}

// authJiraClient authenticates the client with the given Username, API token, and URL domain
func authJiraClient(user, token, url string) *jira.Client {
	tp := jira.BasicAuthTransport{
		Username: user,
		Password: token,
	}
	jiraClient, err := jira.NewClient(tp.Client(), url)
	if err != nil {
		log.Fatalf("Unable to contact Jira: %s", err)
	}
	return jiraClient
}

// assembleIssue parses the Dracon message and serializes it into a Jira Issue object
func (client goJiraClient) assembleIssue(message string) *jira.Issue {
	// Assign the default CustomField values specified in the configuration
	customFields := tcontainer.NewMarshalMap()
	for _, cf := range client.Configs.DefaultValues.CustomFields {
		customFields[cf.ID] = makeCustomField(cf.FieldType, cf.Values)
	}

	// Parse Dracon message into a hashmap
	draconResult, err := parseDraconMessage(message)
	if err != nil {
		log.Fatalf("Could not parse message: %v", err)
	}

	// Mappings the Dracon Result fields to their corresponding Jira fields specified in the configuration
	for _, m := range client.Configs.Mappings {
		customFields[m.JiraField] = makeCustomField(m.FieldType, []string{draconResult[m.DraconField]})
	}

	return &jira.Issue{
		Fields: &jira.IssueFields{
			Project:         makeProjectField(client.Configs.DefaultValues.IssueFields["project"][0]),
			Type:            makeIssueTypeField(client.Configs.DefaultValues.IssueFields["issueType"][0]),
			Description:     makeDescription(draconResult, client.Configs.DescriptionExtras),
			Summary:         makeSummary(draconResult),
			Components:      makeComponentsField(client.Configs.DefaultValues.IssueFields["components"]),
			AffectsVersions: makeAffectsVersionsField(client.Configs.DefaultValues.IssueFields["affectsVersions"]),
			Labels:          client.Configs.DefaultValues.IssueFields["labels"],
			Unknowns:        customFields,
		},
	}
}

// CreateIssue creates a new issue in Jira
func (client goJiraClient) CreateIssue(message string) error {
	issue := client.assembleIssue(message)

	if !client.DryRunMode {
		ri, resp, err := client.JiraClient.Issue.Create(issue)
		if err != nil {
			body, _ := ioutil.ReadAll(resp.Body)
			log.Printf("Error occurred posting to Jira. Response body:\n%s", body)
			return err
		}
		log.Printf("Created Jira Issue ID %s. jira_key=%s", ri.ID, string(ri.Key))
	} else {
		log.Printf("Dry run mode. The following issue would have been created: '%s'", issue.Fields.Summary)
	}
	return nil
}