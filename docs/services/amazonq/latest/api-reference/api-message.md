# Message

A message in an Amazon Q Business web experience.

## Contents

**actionExecution**

Performs an Amazon Q Business plugin action during a non-streaming chat
conversation.

Type: [ActionExecution](api-actionexecution.md) object

Required: No

**actionReview**

An output event that Amazon Q Business returns to an user who wants to perform a plugin
action during a non-streaming chat conversation. It contains information about the
selected action with a list of possible user input fields, some pre-populated by
Amazon Q Business.

Type: [ActionReview](api-actionreview.md) object

Required: No

**attachments**

A file directly uploaded into an Amazon Q Business web experience chat.

Type: Array of [AttachmentOutput](api-attachmentoutput.md) objects

Required: No

**body**

The content of the Amazon Q Business web experience message.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `\P{C}*$}`

Required: No

**messageId**

The identifier of the Amazon Q Business web experience message.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**sourceAttribution**

The source documents used to generate Amazon Q Business web experience message.

Type: Array of [SourceAttribution](api-sourceattribution.md) objects

Required: No

**time**

The timestamp of the first Amazon Q Business web experience message.

Type: Timestamp

Required: No

**type**

The type of Amazon Q Business message, whether `HUMAN` or `AI`
generated.

Type: String

Valid Values: `USER | SYSTEM`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/Message)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/Message)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/Message)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MemberUser

MessageUsefulnessFeedback
