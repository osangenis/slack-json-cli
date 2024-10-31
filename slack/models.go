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
	IsReply         *bool
}

type Replies []*Reply
type Reply struct {
	User    string `json:"user"`
	Ts      string `json:"ts"`
	Message *Message
}

type UserProfile struct {
	FirstName   string `json:"first_name"`
	RealName    string `json:"real_name"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
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

func (m Message) UserDisplayName() string {
	if m.UserProfile != nil {
		if m.UserProfile.DisplayName != "" {
			return m.UserProfile.DisplayName
		}
		if m.UserProfile.RealName != "" {
			return m.UserProfile.RealName
		}
		if m.UserProfile.Name != "" {
			return m.UserProfile.Name
		}
	}
	return m.User
}
