package gojira

import (
	"testing"

	"github.com/andygrunwald/go-jira"
	"github.com/stretchr/testify/assert"
)

var sampleResult = map[string]string{
	"scan_start_time": "0001-01-01T00:00:00Z",
	"scan_id":         "babbb83-4627-41c6-8ba0-70ee866290e9",
	"tool_name":       "spotbugs",
	"source":          "//foo/bar:baz",
	"target":          "//foo1/bar1:baz2",
	"type":            "test type",
	"title":           "Unit Test Title",
	"severity_text":   "Info",
	"cvss":            "0.000",
	"confidence_text": "Info",
	"description":     "this is a test description",
	"first_found":     "0001-01-01T00:00:00Z",
	"false_positive":  "true",
}

func TestGetConfig(t *testing.T) {
	_ = getConfig()
}

func TestParseDraconMessage(t *testing.T) {
	res, err := parseDraconMessage(sampleMessage)

	assert.NoError(t, err)
	assert.EqualValues(t, res, sampleResult)
}

func TestMakeCustomField(t *testing.T) {
	res1 := makeCustomField("single-value", []string{"test-value"})
	exp1 := map[string]string{"value": "test-value"}

	res2 := makeCustomField("multi-value", []string{"value1", "value2", "value3"})
	exp2 := []map[string]string{
		{"value": "value1"},
		{"value": "value2"},
		{"value": "value3"},
	}

	res3 := makeCustomField("float", []string{"4.22"})
	exp3 := 4.22

	assert.EqualValues(t, res1, exp1)
	assert.EqualValues(t, res2, exp2)
	assert.Equal(t, res3, exp3)
}

func TestMakeProjectField(t *testing.T) {
	res1 := makeProjectField("CORP")
	exp1 := jira.Project{
		Key: "CORP",
	}

	assert.EqualValues(t, res1, exp1)
}

func TestMakeIssueTypeField(t *testing.T) {
	res1 := makeIssueTypeField("Vulnerability")
	exp1 := jira.IssueType{
		Name: "Vulnerability",
	}

	assert.EqualValues(t, res1, exp1)
}

func TestMakeComponentsField(t *testing.T) {
	res1 := makeComponentsField([]string{"Engineering", "Quality", "Services"})
	exp1 := []*jira.Component{
		&jira.Component{
			Name: "Engineering",
		},
		&jira.Component{
			Name: "Quality",
		},
		&jira.Component{
			Name: "Services",
		},
	}

	assert.EqualValues(t, res1, exp1)
}

func TestMakeAffectsVersionsField(t *testing.T) {
	res1 := makeAffectsVersionsField([]string{"Release A", "Release B"})
	exp1 := []*jira.AffectsVersion{
		&jira.AffectsVersion{
			Name: "Release A",
		},
		&jira.AffectsVersion{
			Name: "Release B",
		},
	}

	assert.EqualValues(t, res1, exp1)
}

func TestMakeDescription(t *testing.T) {
	extras := []string{"tool_name", "target", "confidence_text"}
	res := makeDescription(sampleResult, extras)
	exp := "This issue was automatically generated by the Dracon security pipeline.\n\n" +
		"*this is a test description*\n\n" +
		"{code:}\n" +
		"tool_name:                 spotbugs\n" +
		"target:                    //foo1/bar1:baz2\n" +
		"confidence_text:           Info\n" +
		"{code}\n"
	assert.Equal(t, res, exp)
}

func TestMakeSummary(t *testing.T) {
	res := makeSummary(sampleResult)
	exp := "bar1:baz2 Unit Test Title"

	assert.Equal(t, res, exp)
}
