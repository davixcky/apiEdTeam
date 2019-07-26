package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const WEBHOOKURL = "https://hooks.slack.com/services/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"

type Slack struct{}

type requestBody struct {
	Text string `json:"text"`
}

func (*Slack) sendNotification(msg string) error {

	slackBody, _ := json.Marshal(requestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, WEBHOOKURL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack 🚫 ")
	}
	return nil
}

func (s *Slack) Created(msg string, id int) (err error) {
	msg = fmt.Sprintf("Se ha *subido* una nueva foto.\n"+
		"Su id es, *%d*. \n"+
		"Estado: 🆗\n ", id)

	return s.sendNotification(msg)
}

func (s *Slack) Deleted(msg string, id int) (err error) {
	msg = fmt.Sprintf("Se ha *eliminado* correctamente la foto.\n"+
		"Id => *%d*. \n"+
		"Estado: 🆗\n", id)

	return s.sendNotification(msg)
}
