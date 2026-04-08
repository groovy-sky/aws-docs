# Parameter

This data type is used as a request parameter in the
`ModifyDBParameterGroup` and `ResetDBParameterGroup` actions.

This data type is used as a response element in the
`DescribeEngineDefaultParameters` and `DescribeDBParameters` actions.

## Contents

###### Note

In the following list, the required parameters are described first.

**AllowedValues**

Specifies the valid range of values for the parameter.

Type: String

Required: No

**ApplyMethod**

Indicates when to apply parameter updates.

Type: String

Valid Values: `immediate | pending-reboot`

Required: No

**ApplyType**

Specifies the engine specific parameters type.

Type: String

Required: No

**DataType**

Specifies the valid data type for the parameter.

Type: String

Required: No

**Description**

Provides a description of the parameter.

Type: String

Required: No

**IsModifiable**

Indicates whether ( `true`) or not ( `false`) the parameter can be modified.
Some parameters have security or operational implications
that prevent them from being changed.

Type: Boolean

Required: No

**MinimumEngineVersion**

The earliest engine version to which the parameter can apply.

Type: String

Required: No

**ParameterName**

The name of the parameter.

Type: String

Required: No

**ParameterValue**

The value of the parameter.

Type: String

Required: No

**Source**

The source of the parameter value.

Type: String

Required: No

**SupportedEngineModes.member.N**

The valid DB engine modes.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/parameter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/parameter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/parameter.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Outpost

PendingCloudwatchLogsExports
