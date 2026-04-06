# FailedAttachmentEvent

A failed file upload during web experience chat.

## Contents

**attachment**

The details of a file uploaded during chat.

Type: [AttachmentOutput](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_AttachmentOutput.html) object

Required: No

**conversationId**

The identifier of the conversation associated with the failed file upload.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**systemMessageId**

The identifier of the AI-generated message associated with the file upload.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**userMessageId**

The identifier of the end user chat message associated with the file upload.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/FailedAttachmentEvent)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/FailedAttachmentEvent)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/FailedAttachmentEvent)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ErrorDetail

FailedDocument
