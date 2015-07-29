package inflect

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetInstanceID reads the instance ID from the metadata URL.
func GetInstanceID() (string, error) {
	res, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// GetRegionFromInstance attempts to read the instance identity document to
// parse out what region an EC2 instance is in.
func GetRegionFromInstance() (string, error) {
	res, err := http.Get("http://169.254.169.254/latest/dynamic/instance-identity/document")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var document map[string]string
	err = json.Unmarshal(body, &document)
	if err != nil {
		return "", err
	}

	return document["region"], nil
}

// GetRegionFromARN parses the given ARN string, returning the AWS region.
func GetRegionFromARN(arn string) (string, error) {
	parts := strings.Split(arn, ":")
	if len(parts) != 6 {
		return "", errors.New("Could not inflect AWS region: ARN does not look valid")
	}
	return string(parts[3]), nil
}
