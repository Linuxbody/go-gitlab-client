package cmd

import (
	"fmt"
	"strings"
)

var resources = map[string][]string{
	"group":            {"group_id"},
	"project":          {"project_id"},
	"user":             {"user_id"},
	"runner":           {"runner_id"},
	"namespace":        {"namespace_id"},
	"project-branch":   {"project_id", "branch_name"},
	"project-badge":    {"project_id", "badge_id"},
	"project-job":      {"project_id", "job_id"},
	"project-pipeline": {"project_id", "pipeline_id"},
	"group-var":        {"group_id", "var_key"},
	"project-var":      {"project_id", "var_key"},
	"project-hook":     {"project_id", "hook_id"},
}

var resourceTypes []string

func init() {
	for key := range resources {
		resourceTypes = append(resourceTypes, key)
	}
}

func resourceCmd(cmdName, resourceType string) string {
	resourceIdKeys := resources[resourceType]

	var upperCased []string
	for _, key := range resourceIdKeys {
		upperCased = append(upperCased, strings.ToUpper(key))
	}

	return fmt.Sprintf(
		"%s %s",
		cmdName,
		strings.Join(upperCased, " "),
	)
}

func isValidResourceType(resourceType string) bool {
	for _, t := range resourceTypes {
		if t == resourceType {
			return true
		}
	}

	return false
}
