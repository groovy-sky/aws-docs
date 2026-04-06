# SourceAttribution

The documents used to generate an Amazon Q Business web experience response.

## Contents

**citationNumber**

The number attached to a citation in an Amazon Q Business generated response.

Type: Integer

Required: No

**datasourceId**

The identifier of the data source from which the document was ingested. This field is not present if the document is ingested by directly calling the BatchPutDocument API (similar to checkDocumentAccess). If the document is from a file-upload data source, the datasource will be "uploaded-docs-file-stat-datasourceid".

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**documentId**

The unique identifier of the source document used in the citation, obtained from the Amazon Q Business index during chat response generation. This ID is used as input to the `GetDocumentContent` API to retrieve the actual document content for user verification.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**indexId**

The identifier of the index containing the source document's metadata and access control information. This links the citation back to the specific Amazon Q Business index where the document's searchable content and permissions are stored.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**snippet**

The content extract from the document on which the generated response is based.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**textMessageSegments**

A text extract from a source document that is used for source attribution.

Type: Array of [TextSegment](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_TextSegment.html) objects

Required: No

**title**

The title of the document which is the source for the Amazon Q Business generated
response.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**updatedAt**

The Unix timestamp when the Amazon Q Business application was last updated.

Type: Timestamp

Required: No

**url**

The URL of the document which is the source for the Amazon Q Business generated response.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/SourceAttribution)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/SourceAttribution)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/SourceAttribution)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SnippetExcerpt

SourceDetails
