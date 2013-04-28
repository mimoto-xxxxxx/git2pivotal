package pivotal

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"net/http"
)

func clientGET(urlStr, pivotalToken string) (*http.Request, error) {
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-TrackerToken", pivotalToken)
	return req, nil
}

func clientPOST(urlStr, pivotalToken, contentType string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", urlStr, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-TrackerToken", pivotalToken)
	req.Header.Add("Content-Type", contentType)
	return req, nil
}

func AddNote(token string, project string, story string, note string) error {
	req, err := clientPOST(
		fmt.Sprintf("http://www.pivotaltracker.com/services/v3/projects/%s/stories/%s/notes", project, story),
		token,
		"application/xml",
		bytes.NewBufferString(fmt.Sprintf("<note><text>%s</text></note>", html.EscapeString(note))))
	if err != nil {
		return err
	}
	resp, err := (&http.Client{}).Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	return nil
}
