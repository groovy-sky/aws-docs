# ScalingConfiguration

Contains the scaling configuration of an Aurora Serverless v1 DB cluster.

For more information, see [Using Amazon Aurora Serverless v1](../../../../services/amazonrds/latest/aurorauserguide/aurora-serverless.md) in the
_Amazon Aurora User Guide_.

## Contents

###### Note

In the following list, the required parameters are described first.

**AutoPause**

Indicates whether to allow or disallow automatic pause for an Aurora DB cluster in `serverless` DB engine mode.
A DB cluster can be paused only when it's idle (it has no connections).

###### Note

If a DB cluster is paused for more than seven days, the DB cluster might be backed up with a snapshot.
In this case, the DB cluster is restored when there is a request to connect to it.

Type: Boolean

Required: No

**MaxCapacity**

The maximum capacity for an Aurora DB cluster in `serverless` DB engine mode.

For Aurora MySQL, valid capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, and `256`.

For Aurora PostgreSQL, valid capacity values are `2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`.

The maximum capacity must be greater than or equal to the minimum capacity.

Type: Integer

Required: No

**MinCapacity**

The minimum capacity for an Aurora DB cluster in `serverless` DB engine mode.

For Aurora MySQL, valid capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, and `256`.

For Aurora PostgreSQL, valid capacity values are `2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`.

The minimum capacity must be less than or equal to the maximum capacity.

Type: Integer

Required: No

**SecondsBeforeTimeout**

The amount of time, in seconds, that Aurora Serverless v1 tries to find a scaling point
to perform seamless scaling before enforcing the timeout action. The default is 300.

Specify a value between 60 and 600 seconds.

Type: Integer

Required: No

**SecondsUntilAutoPause**

The time, in seconds, before an Aurora DB cluster in `serverless` mode is paused.

Specify a value between 300 and 86,400 seconds.

Type: Integer

Required: No

**TimeoutAction**

The action to take when the timeout is reached, either `ForceApplyCapacityChange` or `RollbackCapacityChange`.

`ForceApplyCapacityChange` sets the capacity to the specified value as soon as possible.

`RollbackCapacityChange`, the default, ignores the capacity change if a scaling point isn't found in the timeout period.

###### Important

If you specify `ForceApplyCapacityChange`, connections that
prevent Aurora Serverless v1 from finding a scaling point might be dropped.

For more information, see [Autoscaling for Aurora Serverless v1](../../../../services/amazonrds/latest/aurorauserguide/aurora-serverless-how-it-works-aurora-serverless-how-it-works-auto-scaling.md) in the _Amazon Aurora User Guide_.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/scalingconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/scalingconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/scalingconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ScalarReferenceDetails

ScalingConfigurationInfo
