# AttachmentOutput

The details of a file uploaded during chat.

## Contents

**attachmentId**

The unique identifier of the Amazon Q Business attachment.

Type: String

Pattern: `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`

Required: No

**conversationId**

The unique identifier of the Amazon Q Business conversation.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**error**

An error associated with a file uploaded during chat.

Type: [ErrorDetail](api-errordetail.md) object

Required: No

**name**

The name of a file uploaded during chat.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `\P{C}*`

Required: No

**status**

The status of a file uploaded during chat.

Type: String

Valid Values: `FAILED | SUCCESS`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/attachmentoutput.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/attachmentoutput.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/attachmentoutput.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AttachmentInputEvent

AttachmentsConfiguration

All content copied from https://docs.aws.amazon.com/.
