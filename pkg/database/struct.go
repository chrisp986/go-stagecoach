package database

type EventBuffer struct {
	UUID     uint32
	Sender   uint16
	Receiver uint16
	Event    uint16
	Subtitle string
	Body     string
	Template uint16
	Created  string
	Sent     string
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
