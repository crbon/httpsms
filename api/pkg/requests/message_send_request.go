package requests

import (
	"time"

	"github.com/nyaruka/phonenumbers"

	"github.com/NdoleStudio/http-sms-manager/pkg/services"
)

// MessageSend is the payload for sending and SMS message
type MessageSend struct {
	request
	From    string `json:"from" example:"+18005550199"`
	To      string `json:"to" example:"+18005550100"`
	Content string `json:"content" example:"This is a sample text message"`
}

// Sanitize sets defaults to MessageReceive
func (input *MessageSend) Sanitize() MessageSend {
	input.To = input.sanitizeAddress(input.To)
	input.From = input.sanitizeAddress(input.From)
	return *input
}

// ToMessageSendParams converts MessageSend to services.MessageSendParams
func (input MessageSend) ToMessageSendParams(source string) services.MessageSendParams {
	from, _ := phonenumbers.Parse(input.From, phonenumbers.UNKNOWN_REGION)
	to, _ := phonenumbers.Parse(input.To, phonenumbers.UNKNOWN_REGION)

	return services.MessageSendParams{
		Source:            source,
		Owner:             *from,
		RequestReceivedAt: time.Now().UTC(),
		Contact:           *to,
		Content:           input.Content,
	}
}
