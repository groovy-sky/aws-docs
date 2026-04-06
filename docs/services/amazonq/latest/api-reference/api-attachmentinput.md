# AttachmentInput

This is either a file directly uploaded into a web experience chat or a reference to an existing attachment that is part of a web experience chat.

## Contents

**copyFrom**

A reference to an existing attachment.

Type: [CopyFromSource](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_CopyFromSource.html) object

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/AttachmentInput)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/AttachmentInput)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/AttachmentInput)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Attachment

AttachmentInputEvent
