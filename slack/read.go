package slack

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

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
	for _, m := range messages {
		m.TimeStamp = fromUnixTs(m.Ts)
	}
	lastYearMessages := Messages{}
	lastYear := time.Now().Add(-1 * 365 * 24 * time.Hour)
	for _, m := range messages {
		if m.TimeStamp.After(lastYear) {
			lastYearMessages = append(lastYearMessages, m)
		}
	}
	me.Messages = append(me.Messages, lastYearMessages...)
	messagesWithMissingReplies := me.Messages.WithMissingReplies()
	for _, m := range messagesWithMissingReplies {
		for _, r := range m.Replies {
			if r.Message != nil {
				continue
			}
			r.Message = me.Messages.FindByUserAndTs(r.User, r.Ts)
			if r.Message != nil {
				r.Message.IsReply = ptrTrue()
			}
		}
	}
	return nil
}

func ptrTrue() *bool {
	True := true
	return &True
}

func fromUnixTs(ts string) time.Time {
	timeComps := strings.Split(ts, ".")
	millis, err := strconv.ParseInt(timeComps[0], 10, 64)
	if err != nil {
		return time.Unix(0, 0)
	}
	nanos, err := strconv.ParseInt(timeComps[1], 10, 64)
	if err != nil {
		return time.Unix(0, 0)
	}

	return time.Unix(millis, nanos)
}
