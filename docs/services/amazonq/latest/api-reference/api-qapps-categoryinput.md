# CategoryInput

A label that web experience users associate with a library item. Web experience users use
Categories to tag and filter library items.

## Contents

**id**

The unique identifier of the category.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**title**

The name of the category.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 30.

Pattern: `[a-zA-Z0-9_]+( [a-zA-Z0-9_]+)*`

Required: Yes

**color**

The color of the category, represented as a hexadecimal value of either 3 or 6
digits.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 7.

Pattern: `#([A-Fa-f0-9]{3}|[A-Fa-f0-9]{6})`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/categoryinput.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/categoryinput.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/categoryinput.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Category

ConversationMessage

All content copied from https://docs.aws.amazon.com/.
