package whatsmeow

import (
	"testing"

	"go.mau.fi/whatsmeow/proto/waE2E"
)

// A reply to a business's buttons or list message must not carry a
// <biz><buttons_response/></biz> or <biz><list_response/></biz> node: sends of
// such replies were rejected by the server with error 479 (smax-invalid).
func TestButtonAndListResponsesDoNotRequestBusinessMetadata(t *testing.T) {
	tests := map[string]*waE2E.Message{
		"buttons response": {ButtonsResponseMessage: &waE2E.ButtonsResponseMessage{}},
		"list response":    {ListResponseMessage: &waE2E.ListResponseMessage{}},
	}
	for name, msg := range tests {
		t.Run(name, func(t *testing.T) {
			if got := getButtonTypeFromMessage(msg); got != "" {
				t.Errorf("getButtonTypeFromMessage() = %q, want \"\"", got)
			}
		})
	}
}
