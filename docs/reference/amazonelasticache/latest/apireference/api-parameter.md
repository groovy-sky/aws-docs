# Parameter

Describes an individual setting that controls some aspect of ElastiCache
behavior.

## Contents

###### Note

In the following list, the required parameters are described first.

**AllowedValues**

The valid range of values for the parameter.

Type: String

Required: No

**ChangeType**

Indicates whether a change to the parameter is applied immediately or requires a
reboot for the change to be applied. You can force a reboot or wait until the next
maintenance window's reboot. For more information, see [Rebooting a\
Cluster](../../../../services/amazonelasticache/latest/dg/clusters-rebooting.md).

Type: String

Valid Values: `immediate | requires-reboot`

Required: No

**DataType**

The valid data type for the parameter.

Type: String

Required: No

**Description**

A description of the parameter.

Type: String

Required: No

**IsModifiable**

Indicates whether ( `true`) or not ( `false`) the parameter can be
modified. Some parameters have security or operational implications that prevent them
from being changed.

Type: Boolean

Required: No

**MinimumEngineVersion**

The earliest cache engine version to which the parameter can apply.

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

The source of the parameter.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/parameter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/parameter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/parameter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotificationConfiguration

ParameterNameValue

All content copied from https://docs.aws.amazon.com/.
