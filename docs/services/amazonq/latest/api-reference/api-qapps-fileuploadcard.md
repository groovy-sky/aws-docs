# FileUploadCard

A card in an Amazon Q App that allows the user to upload a file.

## Contents

**dependencies**

Any dependencies or requirements for the file upload card.

Type: Array of strings

Required: Yes

**id**

The unique identifier of the file upload card.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**title**

The title of the file upload card.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 100.

Pattern: `[^{}\\"<>]+`

Required: Yes

**type**

The type of the card.

Type: String

Valid Values: `text-input | q-query | file-upload | q-plugin | form-input`

Required: Yes

**allowOverride**

A flag indicating if the user can override the default file for the upload card.

Type: Boolean

Required: No

**fileId**

The unique identifier of the file associated with the card.

Type: String

Required: No

**filename**

The name of the file being uploaded.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/FileUploadCard)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/FileUploadCard)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/FileUploadCard)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DocumentAttributeValue

FileUploadCardInput
