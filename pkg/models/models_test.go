// Code generated by groq-modeler DO NOT EDIT.
//
// Created at: 2024-11-04 07:23:24
//
// groq-modeler Version 1.1.2
package models_test

import (
	"bytes"
	"context"
	"os"
	"testing"
	"time"

	"github.com/conneroisu/groq-go"
	"github.com/conneroisu/groq-go/pkg/models"
	"github.com/conneroisu/groq-go/pkg/moderation"
	"github.com/conneroisu/groq-go/pkg/test"
	"github.com/stretchr/testify/assert"

	_ "embed"
)

//go:embed testdata/whisper.mp3
var whisperBytes []byte

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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelGemma29BIt,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Gemma29BIt calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelGemma7BIt,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Gemma7BIt calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlama3170BVersatile,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Llama3170BVersatile calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlama318BInstant,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Llama318BInstant calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlama3211BTextPreview,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Llama3211BTextPreview calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlama3211BVisionPreview,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Llama3211BVisionPreview calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlama321BPreview,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Llama321BPreview calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlama323BPreview,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Llama323BPreview calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlama3290BTextPreview,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Llama3290BTextPreview calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlama3290BVisionPreview,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Llama3290BVisionPreview calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlama370B8192,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Llama370B8192 calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlama38B8192,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Llama38B8192 calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlama3Groq70B8192ToolUsePreview,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Llama3Groq70B8192ToolUsePreview calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlama3Groq8B8192ToolUsePreview,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Llama3Groq8B8192ToolUsePreview calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelLlavaV157B4096Preview,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model LlavaV157B4096Preview calling CreateChatCompletion")
	}
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
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.CreateChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: models.ModelMixtral8X7B32768,
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "What is a proface display?",
			},
		},
		MaxTokens:  10,
		RetryDelay: time.Second * 2,
	})
	a.NoError(err, "CreateChatCompletionJSON error")
	if len(response.Choices[0].Message.Content) == 0 {
		t.Errorf("response.Choices[0].Message.Content is empty for model Mixtral8X7B32768 calling CreateChatCompletion")
	}
}

// TestWhisperLargeV3 tests the WhisperLargeV3  transcription model.
//
// It ensures that the model is supported by the groq-go library, the groq API,
// and the operations are working as expected with the api call using this transcription
// model.
func TestWhisperLargeV3(t *testing.T) {
	if len(os.Getenv("UNIT")) < 1 {
		t.Skip("Skipping WhisperLargeV3 transcription test")
	}
	time.Sleep(time.Second * 5)
	a := assert.New(t)
	ctx := context.Background()
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	reader := bytes.NewReader(whisperBytes)
	response, err := client.CreateTranscription(ctx, groq.AudioRequest{
		Model:    models.ModelWhisperLargeV3,
		Reader:   reader,
		FilePath: "whisper.mp3",
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
		t.Skip("Skipping LlamaGuard38B moderation test")
	}
	time.Sleep(time.Second * 5)
	a := assert.New(t)
	ctx := context.Background()
	apiKey, err := test.GetAPIKey("GROQ_KEY")
	a.NoError(err, "GetAPIKey error")
	client, err := groq.NewClient(apiKey)
	a.NoError(err, "NewClient error")
	response, err := client.Moderate(ctx,
		[]groq.ChatCompletionMessage{
			{
				Role:    groq.ChatMessageRoleUser,
				Content: "I want to kill them.",
			},
		},
		models.ModelLlamaGuard38B,
	)
	a.NoError(err, "Moderation error")
	a.Equal(true, response.Flagged)
	a.Contains(
		response.Categories,
		moderation.CategoryViolentCrimes,
	)
}
