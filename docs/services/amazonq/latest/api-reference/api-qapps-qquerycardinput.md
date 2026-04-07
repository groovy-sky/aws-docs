# QQueryCardInput

The input shape for defining a query card in an Amazon Q App.

## Contents

**id**

The unique identifier of the query card.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**prompt**

The prompt or instructions displayed for the query card.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 50000.

Required: Yes

**title**

The title or label of the query card.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 100.

Pattern: `[^{}\\"<>]+`

Required: Yes

**type**

The type of the card.

Type: String

Valid Values: `text-input | q-query | file-upload | q-plugin | form-input`

Required: Yes

**attributeFilter**

Turns on filtering of responses based on document attributes or metadata fields.

Type: [AttributeFilter](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_AttributeFilter.html) object

Required: No

**outputSource**

The source or type of output to generate for the query card.

Type: String

Valid Values: `approved-sources | llm`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/QQueryCardInput)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/QQueryCardInput)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/QQueryCardInput)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QQueryCard

SessionSharingConfiguration
