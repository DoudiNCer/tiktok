package model

import "time"

type Message struct {
	Id          int64     `json:"id" column:"id"`
	Receiver_Id int64     `column:"reciver_id"`
	Listener_Id int64     `column:"listener_id"`
	Text        string    `column:"text"`
	CreatedAt   time.Time `coumn:"create_time"`
}

func (m *Message) Tablename() string {
	return "message"
}
