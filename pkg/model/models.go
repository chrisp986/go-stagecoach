package model

type Event struct {
	ID       uint32 `json:"id,omitempty" db:"id"`
	UniqueID string `json:"unique_id,omitempty" db:"unique_id"`
	Sender   string `json:"sender,omitempty" db:"sender"`
	Receiver string `json:"receiver,omitempty" db:"receiver"`
	Template string `json:"template,omitempty" db:"template"`
	Subtitle string `json:"subtitle,omitempty" db:"subtitle"`
	Body     string `json:"body,omitempty" db:"body"`
	Created  string `json:"created" db:"created"`
	SentDate string `json:"sentdate,omitempty" db:"sent"`
	Sent     int    `json:"sent,omitempty" db:"sent"`
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
