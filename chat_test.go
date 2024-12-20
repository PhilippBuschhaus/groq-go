package groq_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/philippbuschhaus/groq-go"
	"github.com/philippbuschhaus/groq-go/pkg/models"
	"github.com/philippbuschhaus/groq-go/pkg/test"
	"github.com/philippbuschhaus/groq-go/pkg/tools"
	"github.com/stretchr/testify/assert"
)

func TestChat(t *testing.T) {
	ctx := context.Background()
	a := assert.New(t)
	ts := test.NewTestServer()
	returnObj := groq.ChatCompletionResponse{
		ID:      "chatcmpl-123",
		Object:  "chat.completion.chunk",
		Created: 1693721698,
		Model:   "llama3-groq-70b-8192-tool-use-preview",
		Choices: []groq.ChatCompletionChoice{
			{
				Index: 0,
				Message: groq.ChatCompletionMessage{
					Role:    groq.ChatMessageRoleAssistant,
					Content: "Hello!",
				},
			},
		},
	}
	ts.RegisterHandler("/v1/chat/completions", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsval, err := json.Marshal(returnObj)
		a.NoError(err)
		_, err = w.Write(jsval)
		if err != nil {
			t.Fatal(err)
		}
	})
	testS := ts.GroqTestServer()
	testS.Start()
	client, err := groq.NewClient(
		test.GetTestToken(),
		groq.WithBaseURL(testS.URL+"/v1"),
	)
	a.NoError(err)
	resp, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlama3Groq70B8192ToolUsePreview,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "Hello!",
			},
		},
		MaxTokens: 2000,
		Tools:     []tools.Tool{},
	})
	a.NoError(err)
	a.NotEmpty(resp.Choices[0].Message.Content)
}
