# FailedAttachmentEvent

A failed file upload during web experience chat.

## Contents

**attachment**

The details of a file uploaded during chat.

Type: [AttachmentOutput](api-attachmentoutput.md) object

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

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/failedattachmentevent.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/failedattachmentevent.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/failedattachmentevent.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ErrorDetail

FailedDocument

All content copied from https://docs.aws.amazon.com/.
