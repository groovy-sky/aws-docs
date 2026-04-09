# TextInputCard

A card in an Amazon Q App that allows the user to input text.

## Contents

**dependencies**

Any dependencies or requirements for the text input card.

Type: Array of strings

Required: Yes

**id**

The unique identifier of the text input card.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**title**

The title or label of the text input card.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 100.

Pattern: `[^{}\\"<>]+`

Required: Yes

**type**

The type of the card.

Type: String

Valid Values: `text-input | q-query | file-upload | q-plugin | form-input`

Required: Yes

**defaultValue**

The default value to pre-populate in the text input field.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 500.

Required: No

**placeholder**

The placeholder text to display in the text input field.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 500.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/textinputcard.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/textinputcard.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/textinputcard.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubmissionMutation

TextInputCardInput

All content copied from https://docs.aws.amazon.com/.
