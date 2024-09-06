package groq //nolint:testpackage // testing private field

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/conneroisu/groq-go/internal/test"
	"github.com/stretchr/testify/assert"
)

// TestStreamReaderReturnsUnmarshalerErrors tests the stream reader returns an unmarshaler error.
func TestStreamReaderReturnsUnmarshalerErrors(t *testing.T) {
	stream := &streamReader[ChatCompletionStreamResponse]{
		errAccumulator: newErrorAccumulator(),
	}

	respErr := stream.unmarshalError()
	if respErr != nil {
		t.Fatalf("Did not return nil with empty buffer: %v", respErr)
	}

	err := stream.errAccumulator.Write([]byte("{"))
	if err != nil {
		t.Fatalf("%+v", err)
	}

	respErr = stream.unmarshalError()
	if respErr != nil {
		t.Fatalf("Did not return nil when unmarshaler failed: %v", respErr)
	}
}

// TestStreamReaderReturnsErrTooManyEmptyStreamMessages tests the stream reader returns an error when the stream has too many empty messages.
func TestStreamReaderReturnsErrTooManyEmptyStreamMessages(t *testing.T) {
	a := assert.New(t)
	stream := &streamReader[ChatCompletionStreamResponse]{
		emptyMessagesLimit: 3,
		reader: bufio.NewReader(
			bytes.NewReader([]byte("\n\n\n\n")),
		),
		errAccumulator: newErrorAccumulator(),
	}
	_, err := stream.Recv()
	a.ErrorIs(
		err,
		ErrTooManyEmptyStreamMessages{},
		"Did not return error when recv failed",
		err.Error(),
	)
}

// TestStreamReaderReturnsErrTestErrorAccumulatorWriteFailed tests the stream reader returns an error when the error accumulator fails to write.
func TestStreamReaderReturnsErrTestErrorAccumulatorWriteFailed(t *testing.T) {
	a := assert.New(t)
	stream := &streamReader[ChatCompletionStreamResponse]{
		reader: bufio.NewReader(bytes.NewReader([]byte("\n"))),
		errAccumulator: &DefaultErrorAccumulator{
			Buffer: &test.FailingErrorBuffer{},
		},
	}
	_, err := stream.Recv()
	a.ErrorIs(
		err,
		test.ErrTestErrorAccumulatorWriteFailed{},
		"Did not return error when write failed",
		err.Error(),
	)
}
