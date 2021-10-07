package common

import "regexp"

//IsYAMLFile :
func IsYAMLFile(filepath string) bool {
	isYaml := false
	r, err := regexp.MatchString(".yaml", filepath)
	if err == nil && r {
		isYaml = true
	}
	r, err = regexp.MatchString(".yml", filepath)
	if err == nil && r {
		isYaml = true
	}
	return isYaml
}
