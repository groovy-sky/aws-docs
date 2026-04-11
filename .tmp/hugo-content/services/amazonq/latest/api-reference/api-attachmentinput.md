# AttachmentInput

This is either a file directly uploaded into a web experience chat or a reference to an existing attachment that is part of a web experience chat.

## Contents

**copyFrom**

A reference to an existing attachment.

Type: [CopyFromSource](api-copyfromsource.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**data**

The contents of the attachment.

Type: Base64-encoded binary data object

Required: No

**name**

The filename of the attachment.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `\P{C}*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/attachmentinput.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/attachmentinput.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/attachmentinput.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Attachment

AttachmentInputEvent

All content copied from https://docs.aws.amazon.com/.
