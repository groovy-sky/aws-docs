# CardStatus

The current status and value of a card in an active Amazon Q App session.

## Contents

**currentState**

The current state of the card.

Type: String

Valid Values: `IN_PROGRESS | WAITING | COMPLETED | ERROR`

Required: Yes

**currentValue**

The current value or result associated with the card.

Type: String

Required: Yes

**submissions**

A list of previous submissions, if the card is a form card.

Type: Array of [Submission](api-qapps-submission.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/cardstatus.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/cardstatus.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/cardstatus.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CardInput

CardValue

All content copied from https://docs.aws.amazon.com/.
