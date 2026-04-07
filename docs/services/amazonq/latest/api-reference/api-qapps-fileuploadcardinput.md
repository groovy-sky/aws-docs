# FileUploadCardInput

Represents a file upload card. It can optionally receive a `filename` and
`fileId` to set a default file. If not received, the user must provide the file
when the Q App runs.

## Contents

**id**

The unique identifier of the file upload card.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**title**

The title or label of the file upload card.

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

The identifier of a pre-uploaded file associated with the card.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: No

**filename**

The default filename to use for the file upload card.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 100.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/FileUploadCardInput)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/FileUploadCardInput)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/FileUploadCardInput)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FileUploadCard

FormInputCard
