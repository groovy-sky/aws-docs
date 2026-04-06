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

Type: Array of [Submission](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_Submission.html) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/CardStatus)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/CardStatus)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/CardStatus)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CardInput

CardValue
