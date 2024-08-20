package telegram

import "fmt"

var to *Chat
var toMonitor *Chat
var enabled = false

type Chat struct {
	id string
}

func (c Chat) Recipient() string {
	return c.id
}
func (c Chat) SetID(id string) {
	c.id = id
}
func (c Chat) String() string {
	return fmt.Sprintf("%s", c.id)
}
