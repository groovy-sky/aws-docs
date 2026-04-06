# MessageUsefulnessFeedback

End user feedback on an AI-generated web experience chat message usefulness.

## Contents

**submittedAt**

The timestamp for when the feedback was submitted.

Type: Timestamp

Required: Yes

**usefulness**

The usefulness value assigned by an end user to a message.

Type: String

Valid Values: `USEFUL | NOT_USEFUL`

Required: Yes

**comment**

A comment given by an end user on the usefulness of an AI-generated chat
message.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `\P{C}*`

Required: No

**reason**

The reason for a usefulness rating.

Type: String

Valid Values: `NOT_FACTUALLY_CORRECT | HARMFUL_OR_UNSAFE | INCORRECT_OR_MISSING_SOURCES | NOT_HELPFUL | FACTUALLY_CORRECT | COMPLETE | RELEVANT_SOURCES | HELPFUL | NOT_BASED_ON_DOCUMENTS | NOT_COMPLETE | NOT_CONCISE | OTHER`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/MessageUsefulnessFeedback)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/MessageUsefulnessFeedback)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/MessageUsefulnessFeedback)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Message

MetadataEvent
