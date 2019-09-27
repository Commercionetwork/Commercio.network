package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentMetadata_Equals(t *testing.T) {
	metadata := DocumentMetadata{
		ContentUri: "http://example.com/metadata",
		Schema: &DocumentMetadataSchema{
			Uri:     "https://example.com/metadata/schema",
			Version: "1.0.0",
		},
	}
	assert.True(t, metadata.Equals(metadata))
}

func TestDocumentMetadata_Equals_DifferentContents(t *testing.T) {
	metadata := DocumentMetadata{ContentUri: "http://example.com/metadata"}

	other := DocumentMetadata{ContentUri: "https://example.com"}
	assert.False(t, metadata.Equals(other))
}

func TestDocumentMetadata_Equals_DifferentSchema(t *testing.T) {
	metadata := DocumentMetadata{ContentUri: "http://example.com/metadata"}

	other := DocumentMetadata{
		ContentUri: metadata.ContentUri,
		SchemaType: metadata.SchemaType,
		Schema: &DocumentMetadataSchema{
			Uri:     "https://example.com/metadata/schema",
			Version: "1.0.0",
		},
	}
	assert.False(t, metadata.Equals(other))
}

// ---------------
// --- Validate
// ---------------

func TestDocumentMetadata_Validate(t *testing.T) {
	validDocumentMetadata := DocumentMetadata{
		ContentUri: "http://www.contentUri.com",
		Schema: &DocumentMetadataSchema{
			Uri:     "http://www.contentUri.com",
			Version: "test",
		},
	}

	actual := validDocumentMetadata.Validate()
	assert.Nil(t, actual)
}

func TestDocumentMetadata_Validate_EmptyContentUri(t *testing.T) {
	invalidDocumentMetadata := DocumentMetadata{
		ContentUri: "   ",
		Schema: &DocumentMetadataSchema{
			Uri:     "http://www.contentUri.com",
			Version: "test",
		},
	}

	err := invalidDocumentMetadata.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "metadata.content_uri can't be empty", err.Error())
}

func TestDocumentMetadata_Validate_EmptyMetadataInfo(t *testing.T) {
	invalidDocumentMetadata := DocumentMetadata{
		ContentUri: "https://example.com/metadata",
		Schema:     nil,
		SchemaType: "",
	}

	err := invalidDocumentMetadata.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "either metadata.schema or metadata.schema_type must be defined", err.Error())
}

func TestDocumentMetadata_Validate_EmptySchemaUri(t *testing.T) {
	invalidDocumentMetadata := DocumentMetadata{
		ContentUri: "http://www.contentUri.com",
		Schema: &DocumentMetadataSchema{
			Uri:     "",
			Version: "test",
		},
	}

	err := invalidDocumentMetadata.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "metadata.schema.uri can't be empty", err.Error())
}

func TestDocumentMetadata_Validate_EmptySchemaVersion(t *testing.T) {
	invalidDocumentMetadata := DocumentMetadata{
		ContentUri: "http://www.contentUri.com",
		Schema: &DocumentMetadataSchema{
			Uri:     "http://www.contentUri.com",
			Version: "",
		},
	}

	err := invalidDocumentMetadata.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "metadata.schema.version can't be empty", err.Error())
}
