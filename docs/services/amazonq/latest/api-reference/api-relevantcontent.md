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

Type: [ScoreAttributes](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_ScoreAttributes.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/RelevantContent)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/RelevantContent)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/RelevantContent)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QuickSightConfiguration

ResponseConfiguration
