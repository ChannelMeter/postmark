package postmark

import (
	"io"
	"io/ioutil"
	"time"
)

type Header struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type Attachment struct {
	Name        string `json:"Name"`
	Content     []byte `json:"Content"`
	ContentType string `json:"ContentType"`
}

type Email struct {
	From        string       `json:"From"`
	To          string       `json:"To"`
	Cc          string       `json:"Cc,omitempty"`
	Bcc         string       `json:"Bcc,omitempty"`
	Subject     string       `json:"Subject,omitempty"`
	Tag         string       `json:"Tag,omitempty"`
	HtmlBody    string       `json:"HtmlBody"`
	TextBody    string       `json:"TextBody"`
	ReplyTo     string       `json:"ReplyTo,omitempty"`
	Headers     []Header     `json:"Headers,omitempty"`
	TrackOpens  bool         `json:"TrackOpens,omitempty"`
	Attachments []Attachment `json:"Attachments,omitempty"`
}

type EmailResponse struct {
	To          string    `json:"To"`
	SubmittedAt time.Time `json:"SubmittedAt"`
	MessageID   string    `json:"MessageID"`
	ErrorCode   int       `json:"ErrorCode"`
	Message     string    `json:"Message"`
}

func (e *Email) AddHeader(Name, Value string) {
	if e.Headers == nil {
		e.Headers = []Header{Header{Name: Name, Value: Value}}
	} else {
		e.Headers = append(e.Headers, Header{Name: Name, Value: Value})
	}
}

func (e *Email) DelHeader(Name string) bool {
	for i := range e.Headers {
		if e.Headers[i].Name == Name {
			e.Headers[i], e.Headers[len(e.Headers)-1], e.Headers = e.Headers[len(e.Headers)-1], Header{}, e.Headers[:len(e.Headers)-1]
			return true
		}
	}
	return false
}

func (e *Email) AddAttachment(Name string, Content []byte, ContentType string) {
	if e.Attachments == nil {
		e.Attachments = []Attachment{Attachment{Name: Name, Content: Content, ContentType: ContentType}}
	} else {
		e.Attachments = append(e.Attachments, Attachment{Name: Name, Content: Content, ContentType: ContentType})
	}
}

func (e *Email) AddAttachmentReader(Name string, Content io.Reader, ContentType string) error {
	if content, err := ioutil.ReadAll(Content); err == nil {
		if e.Attachments == nil {
			e.Attachments = []Attachment{Attachment{Name: Name, Content: content, ContentType: ContentType}}
		} else {
			e.Attachments = append(e.Attachments, Attachment{Name: Name, Content: content, ContentType: ContentType})
		}
	} else {
		return err
	}
	return nil
}

func (e *Email) DelAttachment(Name string) bool {
	for i := range e.Attachments {
		if e.Attachments[i].Name == Name {
			e.Attachments[i], e.Attachments[len(e.Attachments)-1], e.Attachments = e.Attachments[len(e.Attachments)-1], Attachment{}, e.Attachments[:len(e.Attachments)-1]
			return true
		}
	}
	return false
}
