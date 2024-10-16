// Code generated by groq-modeler DO NOT EDIT.
//
// Created at: 2024-10-16 14:22:50
//
// groq-modeler Version 1.1.2
package groq

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestChatModelsGemma29BIt tests the Gemma29BIt model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsGemma29BIt(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Gemma29BIt test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelGemma29BIt,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Gemma29BIt calling CreateChatCompletionJSON")
}

// TestChatModelsGemma7BIt tests the Gemma7BIt model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsGemma7BIt(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Gemma7BIt test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelGemma7BIt,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Gemma7BIt calling CreateChatCompletionJSON")
}

// TestChatModelsLlama3170BVersatile tests the Llama3170BVersatile model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsLlama3170BVersatile(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Llama3170BVersatile test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelLlama3170BVersatile,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Llama3170BVersatile calling CreateChatCompletionJSON")
}

// TestChatModelsLlama318BInstant tests the Llama318BInstant model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsLlama318BInstant(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Llama318BInstant test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelLlama318BInstant,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Llama318BInstant calling CreateChatCompletionJSON")
}

// TestChatModelsLlama3211BTextPreview tests the Llama3211BTextPreview model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsLlama3211BTextPreview(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Llama3211BTextPreview test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelLlama3211BTextPreview,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Llama3211BTextPreview calling CreateChatCompletionJSON")
}

// TestChatModelsLlama3211BVisionPreview tests the Llama3211BVisionPreview model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsLlama3211BVisionPreview(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Llama3211BVisionPreview test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelLlama3211BVisionPreview,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Llama3211BVisionPreview calling CreateChatCompletionJSON")
}

// TestChatModelsLlama321BPreview tests the Llama321BPreview model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsLlama321BPreview(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Llama321BPreview test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelLlama321BPreview,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Llama321BPreview calling CreateChatCompletionJSON")
}

// TestChatModelsLlama323BPreview tests the Llama323BPreview model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsLlama323BPreview(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Llama323BPreview test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelLlama323BPreview,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Llama323BPreview calling CreateChatCompletionJSON")
}

// TestChatModelsLlama3290BTextPreview tests the Llama3290BTextPreview model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsLlama3290BTextPreview(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Llama3290BTextPreview test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelLlama3290BTextPreview,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Llama3290BTextPreview calling CreateChatCompletionJSON")
}

// TestChatModelsLlama3290BVisionPreview tests the Llama3290BVisionPreview model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsLlama3290BVisionPreview(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Llama3290BVisionPreview test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelLlama3290BVisionPreview,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Llama3290BVisionPreview calling CreateChatCompletionJSON")
}

// TestChatModelsLlama370B8192 tests the Llama370B8192 model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsLlama370B8192(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Llama370B8192 test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelLlama370B8192,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Llama370B8192 calling CreateChatCompletionJSON")
}

// TestChatModelsLlama38B8192 tests the Llama38B8192 model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsLlama38B8192(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Llama38B8192 test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelLlama38B8192,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Llama38B8192 calling CreateChatCompletionJSON")
}

// TestChatModelsLlama3Groq70B8192ToolUsePreview tests the Llama3Groq70B8192ToolUsePreview model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsLlama3Groq70B8192ToolUsePreview(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Llama3Groq70B8192ToolUsePreview test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelLlama3Groq70B8192ToolUsePreview,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Llama3Groq70B8192ToolUsePreview calling CreateChatCompletionJSON")
}

// TestChatModelsLlama3Groq8B8192ToolUsePreview tests the Llama3Groq8B8192ToolUsePreview model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsLlama3Groq8B8192ToolUsePreview(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Llama3Groq8B8192ToolUsePreview test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelLlama3Groq8B8192ToolUsePreview,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Llama3Groq8B8192ToolUsePreview calling CreateChatCompletionJSON")
}

// TestChatModelsLlavaV157B4096Preview tests the LlavaV157B4096Preview model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsLlavaV157B4096Preview(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping LlavaV157B4096Preview test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelLlavaV157B4096Preview,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model LlavaV157B4096Preview calling CreateChatCompletionJSON")
}

// TestChatModelsMixtral8X7B32768 tests the Mixtral8X7B32768 model.
//
// It ensures that the model is supported by the groq-go library and the groq
// API. // and the operations are working as expected for the specific model type.
func TestChatModelsMixtral8X7B32768(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping Mixtral8X7B32768 test")
	}
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, ChatCompletionRequest{
		Model: ModelMixtral8X7B32768,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  5,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	a.NotEmpty(response.Choices[0].Message.Content, "response.Choices[0].Message.Content is empty for model Mixtral8X7B32768 calling CreateChatCompletionJSON")
}

// TestWhisperLargeV3 tests the WhisperLargeV3  transcription model.
//
// It ensures that the model is supported by the groq-go library, the groq API,
// and the operations are working as expected with the api call using this transcription
// model.
func TestWhisperLargeV3(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping WhisperLargeV3 test")
	}
	time.Sleep(time.Second * 5)
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.CreateTranscription(ctx, AudioRequest{
		Model:    ModelWhisperLargeV3,
		FilePath: "./examples/audio-lex-fridman/The Roman Emperors who went insane Gregory Aldrete and Lex Fridman.mp3",
	})
	a.NoError(err, "CreateTranscription error")
	a.NotEmpty(response.Text, "response.Text is empty for model WhisperLargeV3 calling CreateTranscription")
}

// TestLlamaGuard38B tests the LlamaGuard38B model.
//
// It ensures that the model is supported by the groq-go library, the groq API,
// and the operations are working as expected for the specific model type.
func TestLlamaGuard38B(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping LlamaGuard38B test")
	}
	time.Sleep(time.Second * 5)
	a := assert.New(t)
	ctx := context.Background()
	client, err := NewClient(os.Getenv("GROQ_KEY"))
	a.NoError(err, "NewClient error")
	response, err := client.Moderate(ctx, ModerationRequest{
		Model: ModelLlamaGuard38B,
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleUser,
				Content: "I want to kill them.",
			},
		},
	})
	a.NoError(err, "Moderation error")
	a.Equal(true, response.Flagged)
	a.Contains(
		response.Categories,
		CategoryViolentCrimes,
	)
}