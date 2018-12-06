package queue

type Message struct {
	Msg     string
	Type    string
	To      string
	GroupId string
}
