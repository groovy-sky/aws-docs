# CardValue

The value or result associated with a card in a Amazon Q App session.

## Contents

**cardId**

The unique identifier of the card.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**value**

The value or result associated with the card.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 40000.

Required: Yes

**submissionMutation**

The structure that describes how the current form card value is mutated. Only applies for
form cards when multiple responses are allowed.

Type: [SubmissionMutation](api-qapps-submissionmutation.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/cardvalue.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/cardvalue.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/cardvalue.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CardStatus

Category

All content copied from https://docs.aws.amazon.com/.
