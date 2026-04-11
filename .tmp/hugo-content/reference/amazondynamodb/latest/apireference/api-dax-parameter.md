# Parameter

Describes an individual setting that controls some aspect of DAX
behavior.

## Contents

###### Note

In the following list, the required parameters are described first.

**AllowedValues**

A range of values within which the parameter can be set.

Type: String

Required: No

**ChangeType**

The conditions under which changes to this parameter can be applied. For example,
`requires-reboot` indicates that a new value for this parameter will only
take effect if a node is rebooted.

Type: String

Valid Values: `IMMEDIATE | REQUIRES_REBOOT`

Required: No

**DataType**

The data type of the parameter. For example, `integer`:

Type: String

Required: No

**Description**

A description of the parameter

Type: String

Required: No

**IsModifiable**

Whether the customer is allowed to modify the parameter.

Type: String

Valid Values: `TRUE | FALSE | CONDITIONAL`

Required: No

**NodeTypeSpecificValues**

A list of node types, and specific parameter values for each node.

Type: Array of [NodeTypeSpecificValue](api-dax-nodetypespecificvalue.md) objects

Required: No

**ParameterName**

The name of the parameter.

Type: String

Required: No

**ParameterType**

Determines whether the parameter can be applied to any nodes, or only nodes of a
particular type.

Type: String

Valid Values: `DEFAULT | NODE_TYPE_SPECIFIC`

Required: No

**ParameterValue**

The value for the parameter.

Type: String

Required: No

**Source**

How the parameter is defined. For example, `system` denotes a
system-defined parameter.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/parameter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/parameter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/parameter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotificationConfiguration

ParameterGroup

All content copied from https://docs.aws.amazon.com/.
