package model

type Event struct {
	ID       uint32 `json:"id,omitempty" db:"id"`
	UniqueID string `json:"unique_id,omitempty" db:"unique_id"`
	Sender   string `json:"sender,omitempty" db:"sender"`
	Receiver string `json:"receiver,omitempty" db:"receiver"`
	Event    string `json:"event,omitempty" db:"event"`
	Subtitle string `json:"subtitle,omitempty" db:"subtitle"`
	Body     string `json:"body,omitempty" db:"body"`
	Template uint16 `json:"template,omitempty" db:"template"`
	Created  string `json:"created,omitempty" db:"created"`
	Sent     string `json:"sent,omitempty" db:"sent"`
}

type MailAddress struct {
	ID          uint16
	MailAddress string
	FirstName   string
	Name        string
	Status      uint8
	Created     string
}

type MsgTemplate struct {
	ID          uint16
	MsgSubtitle string
	MsgBody     string
	Created     string
}