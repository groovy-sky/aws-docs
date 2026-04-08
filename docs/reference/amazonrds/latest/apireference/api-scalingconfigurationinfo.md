# ScalingConfigurationInfo

The scaling configuration for an Aurora DB cluster in `serverless` DB engine mode.

For more information, see [Using Amazon Aurora Serverless v1](../../../../services/amazonrds/latest/aurorauserguide/aurora-serverless.md) in the
_Amazon Aurora User Guide_.

## Contents

###### Note

In the following list, the required parameters are described first.

**AutoPause**

Indicates whether automatic pause is allowed for the Aurora DB cluster
in `serverless` DB engine mode.

When the value is set to false for an Aurora Serverless v1 DB cluster, the DB cluster automatically resumes.

Type: Boolean

Required: No

**MaxCapacity**

The maximum capacity for an Aurora DB cluster in `serverless` DB engine mode.

Type: Integer

Required: No

**MinCapacity**

The minimum capacity for an Aurora DB cluster in `serverless` DB engine mode.

Type: Integer

Required: No

**SecondsBeforeTimeout**

The number of seconds before scaling times out. What happens when an attempted scaling action times out
is determined by the `TimeoutAction` setting.

Type: Integer

Required: No

**SecondsUntilAutoPause**

The remaining amount of time, in seconds, before the Aurora DB cluster in
`serverless` mode is paused. A DB cluster can be paused only when
it's idle (it has no connections).

Type: Integer

Required: No

**TimeoutAction**

The action that occurs when Aurora times out while attempting to change the capacity of an
Aurora Serverless v1 cluster. The value is either `ForceApplyCapacityChange` or
`RollbackCapacityChange`.

`ForceApplyCapacityChange`, the default, sets the capacity to the specified value as soon as possible.

`RollbackCapacityChange` ignores the capacity change if a scaling point isn't found in the timeout period.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/scalingconfigurationinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/scalingconfigurationinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/scalingconfigurationinfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScalingConfiguration

ServerlessV2FeaturesSupport

All content copied from https://docs.aws.amazon.com/.
