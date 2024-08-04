package termii

type Channel string

const (
	ChannelGeneric  Channel = "generic"
	ChannelDND      Channel = "dnd"
	ChannelWhatsapp Channel = "whatsapp"
)

type SendMessageOptions struct {
	To      string  `json:"to"`
	From    string  `json:"from"`
	SMS     string  `json:"sms"`
	Type    string  `json:"type"`
	Channel Channel `json:"channel"`
	Media   Media   `json:"media"`
	APIKey  string  `json:"api_key"`
}

type SendBulkMessageOptions struct {
	To      []string `json:"to"`
	From    string   `json:"from"`
	SMS     string   `json:"sms"`
	Type    string   `json:"type"`
	Channel Channel  `json:"channel"`
	Media   Media    `json:"media"`
	APIKey  string   `json:"api_key"`
}

type Media struct {
	URL     string `json:"url"`
	Caption string `json:"caption"`
}

type SentMessageResponse struct {
	MessageID string `json:"message_id"`
	Message   string `json:"message"`
}
