package socketstorage

import "encoding/json"

// Message
type Message struct {
	Topic   string `json:"topic"`
	Payload any    `json:"payload"`
}

func NewMessage(topic string, payload any) Message {
	return Message{
		Topic:   topic,
		Payload: payload,
	}
}

func (message Message) Encode() []byte {
	jsonData, err := json.Marshal(message)
	if err != nil {
		panic("error encoding message: " + err.Error())
	}
	return jsonData

}

func DecodeMessage(raw []byte) (Message, error) {
	var m Message
	err := json.Unmarshal(raw, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}
