// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/ee-cli/internal/mocktest"
)

func TestStoreOrdersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "store:orders", "create",
			"--api-key", "string",
			"--id", "10",
			"--complete=true",
			"--pet-id", "198772",
			"--quantity", "7",
			"--ship-date", "'2019-12-27T18:11:19.117Z'",
			"--status", "approved",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: 10\n" +
			"complete: true\n" +
			"petId: 198772\n" +
			"quantity: 7\n" +
			"shipDate: '2019-12-27T18:11:19.117Z'\n" +
			"status: approved\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "store:orders", "create",
			"--api-key", "string",
		)
	})
}

func TestStoreOrdersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "store:orders", "retrieve",
			"--api-key", "string",
			"--order-id", "0",
		)
	})
}

func TestStoreOrdersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "store:orders", "delete",
			"--api-key", "string",
			"--order-id", "0",
		)
	})
}
