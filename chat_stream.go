package groq

import (
	"context"
	"net/http"
)

// ChatCompletionStreamChoiceDelta represents a response structure for chat completion API.
type ChatCompletionStreamChoiceDelta struct {
	Content      string        `json:"content,omitempty"`
	Role         string        `json:"role,omitempty"`
	FunctionCall *FunctionCall `json:"function_call,omitempty"`
	ToolCalls    []ToolCall    `json:"tool_calls,omitempty"`
}

// ChatCompletionStreamChoice represents a response structure for chat completion API.
type ChatCompletionStreamChoice struct {
	Index                int                             `json:"index"`
	Delta                ChatCompletionStreamChoiceDelta `json:"delta"`
	FinishReason         FinishReason                    `json:"finish_reason"`
	ContentFilterResults ContentFilterResults            `json:"content_filter_results,omitempty"`
}

// PromptFilterResult represents a response structure for chat completion API.
type PromptFilterResult struct {
	Index                int                  `json:"index"`
	ContentFilterResults ContentFilterResults `json:"content_filter_results,omitempty"`
}

// ChatCompletionStreamResponse represents a response structure for chat completion API.
type ChatCompletionStreamResponse struct {
	ID                  string                       `json:"id"`                              // ID is the identifier for the chat completion stream response.
	Object              string                       `json:"object"`                          // Object is the object type of the chat completion stream response.
	Created             int64                        `json:"created"`                         // Created is the creation time of the chat completion stream response.
	Model               string                       `json:"model"`                           // Model is the model used for the chat completion stream response.
	Choices             []ChatCompletionStreamChoice `json:"choices"`                         // Choices is the choices for the chat completion stream response.
	SystemFingerprint   string                       `json:"system_fingerprint"`              // SystemFingerprint is the system fingerprint for the chat completion stream response.
	PromptAnnotations   []PromptAnnotation           `json:"prompt_annotations,omitempty"`    // PromptAnnotations is the prompt annotations for the chat completion stream response.
	PromptFilterResults []PromptFilterResult         `json:"prompt_filter_results,omitempty"` // PromptFilterResults is the prompt filter results for the chat completion stream response.
	// Usage is an optional field that will only be present when you set stream_options: {"include_usage": true} in your request.
	// When present, it contains a null value except for the last chunk which contains the token usage statistics
	// for the entire request.
	Usage *Usage `json:"usage,omitempty"`
}

// ChatCompletionStream is a stream of ChatCompletionStreamResponse.
//
// Note: Perhaps it is more elegant to abstract Stream using generics.
type ChatCompletionStream struct {
	*streamReader[ChatCompletionStreamResponse]
}

// CreateChatCompletionStream is an API call to create a chat completion w/ streaming
// support.
//
// If set, tokens will be sent as data-only server-sent events as they become
// available, with the stream terminated by a data: [DONE] message.
func (c *Client) CreateChatCompletionStream(
	ctx context.Context,
	request ChatCompletionRequest,
) (stream *ChatCompletionStream, err error) {
	urlSuffix := chatCompletionsSuffix
	if !checkEndpointSupportsModel(urlSuffix, request.Model) {
		return stream, ErrChatCompletionInvalidModel{model: request.Model}
	}
	request.Stream = true
	req, err := c.newRequest(
		ctx,
		http.MethodPost,
		c.fullURL(urlSuffix, withModel(request.Model)),
		withBody(request),
	)
	if err != nil {
		return nil, err
	}
	resp, err := sendRequestStream[ChatCompletionStreamResponse](c, req)
	if err != nil {
		return
	}
	stream = &ChatCompletionStream{
		streamReader: resp,
	}
	return
}