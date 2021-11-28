package v1

type Data struct {
	Broker   string    `json:"broker"`
	Topic    string    `json:"topic,omitempty"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Topic   string   `json:"topic,omitempty"`
	Key     string   `json:"key,omitempty"`
	Value   string   `json:"value"`
	Headers []Header `json:"headers,omitempty"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
