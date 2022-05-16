package tests

// +test integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/mkdirJava/golangAwsPipelinePlayground/integrationTests/tests/generated_test_stub"
	"github.com/stretchr/testify/require"
)

func TestTodo(t *testing.T) {
	client := generated_test_stub.NewClient(&http.Client{}, "http://test_service:8080/query")
	todoResponse, toDoTransmissionErr := client.GetMyTodos(context.TODO())
	require.NoError(t, toDoTransmissionErr)
	require.Empty(t, todoResponse.Todos)
}
