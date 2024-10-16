// Code generated by groq-modeler DO NOT EDIT.
//
// Created at: 2024-10-16 13:30:12
//
// groq-modeler Version 1.1.2
package groq

type (
	model string

	// Endpoint is the endpoint for the groq api.
	// string
	Endpoint string

	// ChatModel is the type for chat models present on the groq api.
	ChatModel model

	// ModerationModel is the type for moderation models present on the groq api.
	ModerationModel model

	// AudioModel is the type for audio models present on the groq api.
	AudioModel model
)

const (
	chatCompletionsSuffix Endpoint = "/chat/completions"
	transcriptionsSuffix  Endpoint = "/audio/transcriptions"
	translationsSuffix    Endpoint = "/audio/translations"
	embeddingsSuffix      Endpoint = "/embeddings"
	moderationsSuffix     Endpoint = "/moderations"
)

var (
	// ModelGemma29BIt is an AI text chat model.
	//
	// It is created/provided by Google.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelGemma29BIt ChatModel = "gemma2-9b-it"
	// ModelGemma7BIt is an AI text chat model.
	//
	// It is created/provided by Google.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelGemma7BIt ChatModel = "gemma-7b-it"
	// ModelLlama3170BVersatile is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 131072 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelLlama3170BVersatile ChatModel = "llama-3.1-70b-versatile"
	// ModelLlama318BInstant is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 131072 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelLlama318BInstant ChatModel = "llama-3.1-8b-instant"
	// ModelLlama3211BTextPreview is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelLlama3211BTextPreview ChatModel = "llama-3.2-11b-text-preview"
	// ModelLlama3211BVisionPreview is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelLlama3211BVisionPreview ChatModel = "llama-3.2-11b-vision-preview"
	// ModelLlama321BPreview is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelLlama321BPreview ChatModel = "llama-3.2-1b-preview"
	// ModelLlama323BPreview is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelLlama323BPreview ChatModel = "llama-3.2-3b-preview"
	// ModelLlama3290BTextPreview is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelLlama3290BTextPreview ChatModel = "llama-3.2-90b-text-preview"
	// ModelLlama3290BVisionPreview is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelLlama3290BVisionPreview ChatModel = "llama-3.2-90b-vision-preview"
	// ModelLlama370B8192 is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelLlama370B8192 ChatModel = "llama3-70b-8192"
	// ModelLlama38B8192 is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelLlama38B8192 ChatModel = "llama3-8b-8192"
	// ModelLlama3Groq70B8192ToolUsePreview is an AI text chat model.
	//
	// It is created/provided by Groq.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelLlama3Groq70B8192ToolUsePreview ChatModel = "llama3-groq-70b-8192-tool-use-preview"
	// ModelLlama3Groq8B8192ToolUsePreview is an AI text chat model.
	//
	// It is created/provided by Groq.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelLlama3Groq8B8192ToolUsePreview ChatModel = "llama3-groq-8b-8192-tool-use-preview"
	// ModelLlavaV157B4096Preview is an AI text chat model.
	//
	// It is created/provided by Other.
	//
	// It has 4096 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelLlavaV157B4096Preview ChatModel = "llava-v1.5-7b-4096-preview"
	// ModelMixtral8X7B32768 is an AI text chat model.
	//
	// It is created/provided by Mistral AI.
	//
	// It has 32768 context window.
	//
	// It can be used with the following client methods:
	//	- CreateChatCompletion
	// 	- CreateChatCompletionStream
	// 	- CreateChatCompletionJSON
	ModelMixtral8X7B32768 ChatModel = "mixtral-8x7b-32768"
	// ModelWhisperLargeV3 is an AI audio transcription model.
	//
	// It is created/provided by OpenAI.
	//
	// It has 448 context window.
	//
	// It can be used with the following client methods:
	//	- CreateTranscription
	// 	- CreateTranslation
	ModelWhisperLargeV3 AudioModel = "whisper-large-v3"
	// ModelLlamaGuard38B is an AI moderation model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- Moderate
	ModelLlamaGuard38B ModerationModel = "llama-guard-3-8b"
)
