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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/PredictAppDefinition)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/PredictAppDefinition)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/PredictAppDefinition)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PermissionOutput

PredictQAppInputOptions
