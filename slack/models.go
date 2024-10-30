package slack

import "time"

type MessageExport struct {
	Messages Messages
}
type Messages []*Message
type Message struct {
	User            string `json:"user"`
	Type            string `json:"type"`
	TimeStamp       time.Time
	Ts              string       `json:"ts"`
	UserProfile     *UserProfile `json:"user_profile"`
	ThreadTimeStamp time.Time
	ThreadTs        string  `json:"thread_ts"`
	Text            string  `json:"text"`
	ClientMsgID     string  `json:"client_msg_id"`
	Replies         Replies `json:"replies"`
}

type Replies []*Reply
type Reply struct {
	User    string `json:"user"`
	Ts      string `json:"ts"`
	Message *Message
}

type UserProfile struct {
	RealName string `json:"real_name"`
	UserName string `json:"name"`
}

func (l Messages) WithMissingReplies() Messages {
	res := Messages{}
	for _, m := range l {
		for _, r := range m.Replies {
			if r.Message == nil {
				res = append(res, m)
			}
		}
	}
	return res
}

func (l Messages) FindByUserAndTs(user string, ts string) *Message {
	for _, m := range l {
		if m.User == user && m.Ts == ts {
			return m
		}
	}
	return nil
}
