package gojira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/andygrunwald/go-jira"
	yaml "gopkg.in/yaml.v2"
)

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

// parseDraconMessage returns a hashmap of the parsed dracon message
func parseDraconMessage(message string) (map[string]string, error) {
	jBytes := []byte(message)
	var draconResult map[string]string
	err := json.Unmarshal(jBytes, &draconResult)
	return draconResult, err
}

// makeCustomField returns the appropriate interface for a jira CustomField given it's type and values
// :param fieldType: the type of the field in Jira (single-value, multi-value, float)
// :param values: list of values to be filled in
// :return the appropriate interface for a CustomField, given the corresponding fieldType and value(s)
func makeCustomField(fieldType string, values []string) interface{} {
	switch fieldType {
	case "single-value":
		return map[string]string{"value": values[0]}
	case "multi-value":
		cf := []map[string]string{}
		for _, v := range values {
			cf = append(cf, map[string]string{"value": v})
		}
		return cf
	case "float":
		if f, err := strconv.ParseFloat(values[0], 64); err == nil {
			return f
		} else {
			log.Fatalf("Error parsing float field-type: %v", err)
		}
	default:
		return nil
	}
	return nil
}

// makeProjectField returns a jira.Project object with the given key value
func makeProjectField(value string) jira.Project {
	return jira.Project{
		Key: value,
	}
}

// makeIssueTypeField returns a jira.IssueType object with the given name value
func makeIssueTypeField(value string) jira.IssueType {
	return jira.IssueType{
		Name: value,
	}
}

// makeComponentsField returns a []jira.Components object with the given name value(s)
func makeComponentsField(values []string) []*jira.Component {
	components := []*jira.Component{}
	for _, v := range values {
		components = append(components, &jira.Component{Name: v})
	}
	return components
}

// makeAffectsVersionsField returns a []jira.AffectsVersion object with the given name value(s)
func makeAffectsVersionsField(values []string) []*jira.AffectsVersion {
	affectsVersions := []*jira.AffectsVersion{}
	for _, v := range values {
		affectsVersions = append(affectsVersions, &jira.AffectsVersion{Name: v})
	}
	return affectsVersions
}

// makeDescription creates the description of an issue's enhanced with extra information from the Dracon Result
// depending on the tool used, diffrent fields are more suitable as description
func makeDescription(draconResult map[string]string, extras []string) string {
	desc := "This issue was automatically generated by the Dracon security pipeline.\n\n"
	switch draconResult["tool_name"] {
	case "gosec":
		desc = desc + "*" + draconResult["title"] + "*" + "\n\n"
	case "spotbugs":
		desc = desc + "*" + draconResult["description"] + "*" + "\n\n"
	case "bandit":
		desc = desc + "*" + draconResult["description"] + "*" + "\n\n"
	}

	// Append the extra fields to the description
	if len(extras) > 0 {
		desc = desc + "{code:}" + "\n"
		for _, s := range extras {
			desc = desc + fmt.Sprintf("%s: %*s\n", s, 25-len(s)+len(draconResult[s]), draconResult[s])
		}
		desc = desc + "{code}" + "\n"
	}
	return desc
}

// makeSummary creates the Summary/Title of an issue
// depending on the tool used, different fields are more suitable as summary
func makeSummary(draconResult map[string]string) string {
	switch draconResult["tool_name"] {
	case "gosec":
		return filepath.Base(draconResult["target"]) + " " + draconResult["description"]
	case "spotbugs":
		return filepath.Base(draconResult["target"]) + " " + draconResult["title"]
	case "bandit":
		// Note: setting character limit to 100 until the bandit producer's 'title' gets fixed
		return fmt.Sprintf("%.120s", (filepath.Base(draconResult["target"]) + " " + draconResult["description"]))
	}
	return ""
}
