# PredictAppDefinition

The definition of an Amazon Q App generated based on input such as a conversation or problem
statement.

## Contents

**appDefinition**

The definition specifying the cards and flow of the generated Q App.

Type: [AppDefinitionInput](api-qapps-appdefinitioninput.md) object

Required: Yes

**title**

The title of the generated Q App definition.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 100.

Pattern: `[^{}\\"<>]+`

Required: Yes

**description**

The description of the generated Q App definition.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 500.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/predictappdefinition.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/predictappdefinition.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/predictappdefinition.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PermissionOutput

PredictQAppInputOptions

All content copied from https://docs.aws.amazon.com/.
