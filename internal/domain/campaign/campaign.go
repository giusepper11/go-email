package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedAt time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(emails) == 0 {
		return nil, errors.New("emails is required")
	}

	contacts := make([]Contact, len(emails))
	for ix, email := range emails {
		contacts[ix].Email = email
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
	}, nil
}
