package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Client defines the HTTP client interface
type Client interface {
	FetchURL(url string, headers map[string]string) ([]byte, error)
	ToJSON(data []byte, v interface{}) error
	FromJSON(v interface{}) ([]byte, error)
	GetLatestReleaseVersion(repo string) (string, error)
	GetPackageVersion(repo, filePath string) (string, error)
}

// client implements the Client interface
type client struct{}

// NewClient returns a new instance of Client
func NewClient() Client {
	return &client{}
}

// FetchURL makes an HTTP GET request with headers
func (c *client) FetchURL(url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// ToJSON unmarshals a byte slice into a struct
func (c *client) ToJSON(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

// FromJSON marshals a struct into a byte slice
func (c *client) FromJSON(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetLatestReleaseVersion fetches the latest release version from GitHub
func (c *client) GetLatestReleaseVersion(repo string) (string, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	headers := map[string]string{
		"Accept": "application/vnd.github.v3+json",
	}

	data, err := c.FetchURL(apiURL, headers)
	if err != nil {
		return "", err
	}

	var release struct {
		TagName string `json:"tag_name"`
	}

	err = c.ToJSON(data, &release)
	if err != nil {
		return "", err
	}

	if release.TagName == "" {
		return "", errors.New("no release found")
	}

	return release.TagName, nil
}

// GetPackageVersion retrieves the version from a specific file in the repo (e.g., package.json, go.mod)
func (c *client) GetPackageVersion(repo, filePath string) (string, error) {
	apiURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/main/%s", repo, filePath)
	headers := map[string]string{}

	data, err := c.FetchURL(apiURL, headers)
	if err != nil {
		return "", err
	}

	// Detect the file type and extract the version
	if strings.HasSuffix(filePath, "package.json") {
		return extractVersionFromPackageJSON(data)
	} else if strings.HasSuffix(filePath, "go.mod") {
		return extractVersionFromGoMod(data)
	}

	// Add support for more file types as needed
	return "", errors.New("unsupported file type for version extraction")
}

// extractVersionFromPackageJSON extracts version from package.json data
func extractVersionFromPackageJSON(data []byte) (string, error) {
	var packageJSON struct {
		Version string `json:"version"`
	}

	err := json.Unmarshal(data, &packageJSON)
	if err != nil {
		return "", err
	}

	if packageJSON.Version == "" {
		return "", errors.New("version not found in package.json")
	}

	return packageJSON.Version, nil
}

// extractVersionFromGoMod extracts version from go.mod data
func extractVersionFromGoMod(data []byte) (string, error) {
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "go ") {
			// Example of a line: go 1.16
			version := strings.TrimSpace(strings.TrimPrefix(line, "go"))
			return version, nil
		}
	}

	return "", errors.New("version not found in go.mod")
}
