package slack

import (
	"fmt"
	"io"
)

type OutputFormat string

const OuputFormat_PlainText = "plain"
const TimeFormat = "2006-01-02 15:04:05"

func (me *MessageExport) Write(format OutputFormat, w io.Writer) {
	for _, message := range me.Messages {
		if message.IsReply != nil && *message.IsReply {
			continue
		}
		tsFormatted := message.TimeStamp.Format(TimeFormat)
		user := message.UserDisplayName()

		msgBody := fmt.Sprintf("[%v][%v]%v\n",
			tsFormatted, user, message.Text)
		w.Write([]byte(msgBody))
		for _, reply := range message.Replies {
			if reply.Message == nil {
				continue
			}
			tsFormatted := reply.Message.TimeStamp.Format(TimeFormat)
			user := reply.Message.UserDisplayName()
			msgReply := fmt.Sprintf("\t[%v][reply from %v]%v\n",
				tsFormatted, user, reply.Message.Text)
			w.Write([]byte(msgReply))
		}
	}
}
