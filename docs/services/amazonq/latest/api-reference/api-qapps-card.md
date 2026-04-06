# Card

A card representing a component or step in an Amazon Q App's flow.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**fileUpload**

A container for the properties of the file upload card.

Type: [FileUploadCard](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_FileUploadCard.html) object

Required: No

**formInput**

A container for the properties of the form input card.

Type: [FormInputCard](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_FormInputCard.html) object

Required: No

**qPlugin**

A container for the properties of the plugin card.

Type: [QPluginCard](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_QPluginCard.html) object

Required: No

**qQuery**

A container for the properties of the query card.

Type: [QQueryCard](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_QQueryCard.html) object

Required: No

**textInput**

A container for the properties of the text input card.

Type: [TextInputCard](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_TextInputCard.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/Card)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/Card)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/Card)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BatchCreateCategoryInputCategory

CardInput
