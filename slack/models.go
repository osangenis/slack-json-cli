package slack

import "time"

type MessageExport struct {
	Messages Messages
}
type Messages []Message
type Message struct {
	User            string       `json:"user"`
	Type            string       `json:"type"`
	TimeStamp       time.Time    `json:"ts"`
	UserProfile     *UserProfile `json:"user_profile"`
	ThreadTimeStamp time.Time    `json:"thread_ts"`
	Text            string       `json:"text"`
}
type UserProfile struct {
	RealName string `json:"real_name"`
	UserName string `json:"name"`
}
