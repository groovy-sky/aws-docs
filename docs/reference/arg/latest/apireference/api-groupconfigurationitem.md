# GroupConfigurationItem

An item in a group configuration. A group service configuration can have one or more
items. For details about group service configuration syntax, see [Service configurations for\
resource groups](about-slg.md).

## Contents

###### Note

In the following list, the required parameters are described first.

**Type**

Specifies the type of group configuration item. Each item must have a unique value for
`type`. For the list of types that you can specify for a configuration
item, see [Supported resource types and\
parameters](about-slg.md#about-slg-types).

Type: String

Length Constraints: Maximum length of 40.

Pattern: `AWS::[a-zA-Z0-9]+::[a-zA-Z0-9]+`

Required: Yes

**Parameters**

A collection of parameters for this group configuration item. For the list of
parameters that you can use with each configuration item type, see [Supported\
resource types and parameters](about-slg.md#about-slg-types).

Type: Array of [GroupConfigurationParameter](api-groupconfigurationparameter.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/resource-groups-2017-11-27/groupconfigurationitem.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/resource-groups-2017-11-27/groupconfigurationitem.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/resource-groups-2017-11-27/groupconfigurationitem.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GroupConfiguration

GroupConfigurationParameter

All content copied from https://docs.aws.amazon.com/.
