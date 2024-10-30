package slack

import (
	"fmt"
	"io"
)

type OutputFormat string

const OuputFormat_PlainText = "plain"

func (me *MessageExport) Write(format OutputFormat, w io.Writer) {
	for _, message := range me.Messages {
		msgBody := fmt.Sprintf("[%v][%v]%v\n",
			message.TimeStamp, message.User, message.Text)
		w.Write([]byte(msgBody))
		for _, reply := range message.Replies {
			if reply.Message == nil {
				continue
			}
			msgReply := fmt.Sprintf("\t[%v][reply from %v]%v\n",
				reply.Message.TimeStamp, reply.Message.User, reply.Message.Text)
			w.Write([]byte(msgReply))
		}
	}
}
