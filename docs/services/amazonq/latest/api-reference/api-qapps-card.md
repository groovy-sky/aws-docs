# Card

A card representing a component or step in an Amazon Q App's flow.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**fileUpload**

A container for the properties of the file upload card.

Type: [FileUploadCard](api-qapps-fileuploadcard.md) object

Required: No

**formInput**

A container for the properties of the form input card.

Type: [FormInputCard](api-qapps-forminputcard.md) object

Required: No

**qPlugin**

A container for the properties of the plugin card.

Type: [QPluginCard](api-qapps-qplugincard.md) object

Required: No

**qQuery**

A container for the properties of the query card.

Type: [QQueryCard](api-qapps-qquerycard.md) object

Required: No

**textInput**

A container for the properties of the text input card.

Type: [TextInputCard](api-qapps-textinputcard.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/card.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/card.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/card.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchCreateCategoryInputCategory

CardInput

All content copied from https://docs.aws.amazon.com/.
