# AppDefinition

The definition of the Q App, specifying the cards and flow.

## Contents

**appDefinitionVersion**

The version of the app definition schema or specification.

Type: String

Required: Yes

**cards**

The cards that make up the Q App, such as text input, file upload, or query cards.

Type: Array of [Card](api-qapps-card.md) objects

Array Members: Minimum number of 0 items. Maximum number of 20 items.

Required: Yes

**canEdit**

A flag indicating whether the Q App's definition can be edited by the user.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/appdefinition.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/appdefinition.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/appdefinition.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QApps

AppDefinitionInput

All content copied from https://docs.aws.amazon.com/.
