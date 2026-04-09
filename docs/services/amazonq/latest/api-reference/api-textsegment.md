# TextSegment

Provides information about a text extract in a chat response that can be attributed to
a source document.

## Contents

**beginOffset**

The zero-based location in the response string where the source attribution
starts.

Type: Integer

Required: No

**endOffset**

The zero-based location in the response string where the source attribution
ends.

Type: Integer

Required: No

**mediaId**

_This member has been deprecated._

The identifier of the media object associated with the text segment in the source attribution.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**mediaMimeType**

_This member has been deprecated._

The MIME type (image/png) of the media object associated with the text segment in the source attribution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**snippetExcerpt**

The relevant text excerpt from a source that was used to generate a citation text
segment in an Amazon Q Business chat response.

Type: [SnippetExcerpt](api-snippetexcerpt.md) object

Required: No

**sourceDetails**

Source information for a segment of extracted text, including its media type.

Type: [SourceDetails](api-sourcedetails.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/textsegment.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/textsegment.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/textsegment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TextOutputEvent

TopicConfiguration

All content copied from https://docs.aws.amazon.com/.
