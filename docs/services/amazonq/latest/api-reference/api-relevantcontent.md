# RelevantContent

Represents a piece of content that is relevant to a search query.

## Contents

**content**

The actual content of the relevant item.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**documentAttributes**

Additional attributes of the document containing the relevant content.

Type: Array of [DocumentAttribute](api-documentattribute.md) objects

Array Members: Minimum number of 1 item. Maximum number of 500 items.

Required: No

**documentId**

The unique identifier of the document containing the relevant content.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1825.

Pattern: `\P{C}*`

Required: No

**documentTitle**

The title of the document containing the relevant content.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**documentUri**

The URI of the document containing the relevant content.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(https?|ftp|file)://([^\s]*)`

Required: No

**scoreAttributes**

Attributes related to the relevance score of the content.

Type: [ScoreAttributes](api-scoreattributes.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/relevantcontent.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/relevantcontent.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/relevantcontent.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QuickSightConfiguration

ResponseConfiguration

All content copied from https://docs.aws.amazon.com/.
