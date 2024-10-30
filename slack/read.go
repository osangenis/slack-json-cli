package slack

import "encoding/json"

func NewMessageExport() *MessageExport {
	return &MessageExport{
		Messages: Messages{},
	}
}

// Add a list of messages to the export
// rawJson is expected to match the Messages type
// Typically, this is coming from a json file of a single day export
func (me *MessageExport) AddRawMessages(rawJson []byte) error {
	messages := Messages{}
	err := json.Unmarshal(rawJson, &messages)
	if err != nil {
		return err
	}
	me.Messages = append(me.Messages, messages...)
	messagesWithMissingReplies := me.Messages.WithMissingReplies()
	for _, m := range messagesWithMissingReplies {
		for _, r := range m.Replies {
			if r.Message != nil {
				continue
			}
			r.Message = me.Messages.FindByUserAndTs(r.User, r.Ts)
		}
	}
	return nil
}
