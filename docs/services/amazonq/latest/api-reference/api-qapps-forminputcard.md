# FormInputCard

A card in an Amazon Q App that allows the user to submit a response.

## Contents

**dependencies**

Any dependencies or requirements for the form input card.

Type: Array of strings

Required: Yes

**id**

The unique identifier of the form input card.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**metadata**

The metadata that defines the form input card data.

Type: [FormInputCardMetadata](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_FormInputCardMetadata.html) object

Required: Yes

**title**

The title of the form input card.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 100.

Pattern: `[^{}\\"<>]+`

Required: Yes

**type**

The type of the card.

Type: String

Valid Values: `text-input | q-query | file-upload | q-plugin | form-input`

Required: Yes

**computeMode**

The compute mode of the form input card. This property determines whether individual
participants of a data collection session can submit multiple response or one response. A
compute mode of `append` shall allow participants to submit the same form multiple
times with different values. A compute mode of `replace` code> shall overwrite the
current value for each participant.

Type: String

Valid Values: `append | replace`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/FormInputCard)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/FormInputCard)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/FormInputCard)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FileUploadCardInput

FormInputCardInput
