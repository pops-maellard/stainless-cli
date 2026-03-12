// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/ee-cli/internal/mocktest"
	"github.com/stainless-sdks/ee-cli/internal/requestflag"
)

func TestPetsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "pets", "create",
			"--api-key", "string",
			"--name", "doggie",
			"--photo-url", "string",
			"--id", "10",
			"--category", "{id: 1, name: Dogs}",
			"--status", "available",
			"--tag", "{id: 0, name: name}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(petsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "pets", "create",
			"--api-key", "string",
			"--name", "doggie",
			"--photo-url", "string",
			"--id", "10",
			"--category.id", "1",
			"--category.name", "Dogs",
			"--status", "available",
			"--tag.id", "0",
			"--tag.name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: doggie\n" +
			"photoUrls:\n" +
			"  - string\n" +
			"id: 10\n" +
			"category:\n" +
			"  id: 1\n" +
			"  name: Dogs\n" +
			"status: available\n" +
			"tags:\n" +
			"  - id: 0\n" +
			"    name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "pets", "create",
			"--api-key", "string",
		)
	})
}

func TestPetsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "pets", "retrieve",
			"--api-key", "string",
			"--pet-id", "0",
		)
	})
}

func TestPetsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "pets", "update",
			"--api-key", "string",
			"--name", "doggie",
			"--photo-url", "string",
			"--id", "10",
			"--category", "{id: 1, name: Dogs}",
			"--status", "available",
			"--tag", "{id: 0, name: name}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(petsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "pets", "update",
			"--api-key", "string",
			"--name", "doggie",
			"--photo-url", "string",
			"--id", "10",
			"--category.id", "1",
			"--category.name", "Dogs",
			"--status", "available",
			"--tag.id", "0",
			"--tag.name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: doggie\n" +
			"photoUrls:\n" +
			"  - string\n" +
			"id: 10\n" +
			"category:\n" +
			"  id: 1\n" +
			"  name: Dogs\n" +
			"status: available\n" +
			"tags:\n" +
			"  - id: 0\n" +
			"    name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "pets", "update",
			"--api-key", "string",
		)
	})
}

func TestPetsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "pets", "delete",
			"--api-key", "string",
			"--pet-id", "0",
		)
	})
}

func TestPetsFindByStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "pets", "find-by-status",
			"--api-key", "string",
			"--status", "available",
		)
	})
}

func TestPetsFindByTags(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "pets", "find-by-tags",
			"--api-key", "string",
			"--tag", "string",
		)
	})
}

func TestPetsUpdateByID(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "pets", "update-by-id",
			"--api-key", "string",
			"--pet-id", "0",
			"--name", "name",
			"--status", "status",
		)
	})
}

func TestPetsUploadImage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "pets", "upload-image",
			"--api-key", "string",
			"--pet-id", "0",
			"--image", mocktest.TestFile(t, "Example data"),
			"--additional-metadata", "additionalMetadata",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("Example data")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "pets", "upload-image",
			"--api-key", "string",
			"--pet-id", "0",
			"--additional-metadata", "additionalMetadata",
		)
	})
}
