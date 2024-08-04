package termii

import (
	"context"
	"net/http"
)

type Messaging struct {
	client *Client
}

func (m *Messaging) SendMessage(ctx context.Context, opts *SendMessageOptions) (*SentMessageResponse, *Response, error) {
	message := SentMessageResponse{}

	req, err := m.client.NewRequest(http.MethodPost, "sms/send", opts)
	if err != nil {
		return nil, nil, err
	}

	res, err := m.client.Do(ctx, req, message)
	if err != nil {
		return nil, nil, err
	}

	return &message, res, nil
}

func (m *Messaging) SendBulkMessage(ctx context.Context, opts *SendBulkMessageOptions) (*SentMessageResponse, *Response, error) {
	message := SentMessageResponse{}

	req, err := m.client.NewRequest(http.MethodPost, "sms/send/bulk", opts)
	if err != nil {
		return nil, nil, err
	}

	res, err := m.client.Do(ctx, req, message)
	if err != nil {
		return nil, nil, err
	}

	return &message, res, nil
}
