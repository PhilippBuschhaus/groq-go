package toolhouse_test

import (
	"context"
	"os"
	"testing"

	"github.com/conneroisu/groq-go"
	"github.com/conneroisu/groq-go/extensions/toolhouse"
	"github.com/conneroisu/groq-go/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestNewExtension(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	if !test.IsUnitTest() {
		t.Skip("Skipping Toolhouse extension test")
	}

	ext, err := toolhouse.NewExtension(os.Getenv("TOOLHOUSE_API_KEY"),
		toolhouse.WithMetadata(map[string]any{
			"id":       "conner",
			"timezone": 5,
		}),
		toolhouse.WithLogger(test.DefaultLogger),
	)
	a.NoError(err)
	client, err := groq.NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err)
	history := []groq.ChatCompletionMessage{
		{
			Role:    groq.ChatMessageRoleUser,
			Content: "Write a python function to print the first 10 prime numbers containing the number 3 then respond with the answer. DO NOT GUESS WHAT THE OUTPUT SHOULD BE. MAKE SURE TO CALL THE TOOL GIVEN.",
		},
	}
	tooling, err := ext.GetTools(ctx)
	a.NoError(err)
	re, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model:      groq.ModelLlama3Groq70B8192ToolUsePreview,
		Messages:   history,
		Tools:      tooling,
		ToolChoice: "required",
	})
	a.NoError(err)
	history = append(history, re.Choices[0].Message)
	print(history[len(history)-1].ToolCalls[len(history[len(history)-1].ToolCalls)-1].Function.Arguments)
	r, err := ext.Run(ctx, re)
	a.NoError(err)
	history = append(history, r...)
	finalr, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model:     groq.ModelLlama3Groq70B8192ToolUsePreview,
		Messages:  history,
		MaxTokens: 2000,
	})
	a.NoError(err)
	history = append(history, finalr.Choices[0].Message)
	print(history[len(history)-1].Content)
}