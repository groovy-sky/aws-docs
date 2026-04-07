This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBCluster ScalingConfiguration

The `ScalingConfiguration` property type specifies the scaling
configuration of an Aurora Serverless v1 DB cluster.

For more information, see [Using Amazon Aurora\
Serverless](../../../amazonrds/latest/aurorauserguide/aurora-serverless.md) in the _Amazon Aurora User Guide_.

This property is only supported for Aurora Serverless v1. For Aurora Serverless v2, Use the `ServerlessV2ScalingConfiguration` property.

Valid for: Aurora Serverless v1 DB clusters only

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoPause" : Boolean,
  "MaxCapacity" : Integer,
  "MinCapacity" : Integer,
  "SecondsBeforeTimeout" : Integer,
  "SecondsUntilAutoPause" : Integer,
  "TimeoutAction" : String
}

```

### YAML

```yaml

  AutoPause: Boolean
  MaxCapacity: Integer
  MinCapacity: Integer
  SecondsBeforeTimeout: Integer
  SecondsUntilAutoPause: Integer
  TimeoutAction: String

```

## Properties

`AutoPause`

Indicates whether to allow or disallow automatic pause for an Aurora DB cluster in `serverless` DB engine mode.
A DB cluster can be paused only when it's idle (it has no connections).

###### Note

If a DB cluster is paused for more than seven days, the DB cluster might be backed up with a snapshot.
In this case, the DB cluster is restored when there is a request to connect to it.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxCapacity`

The maximum capacity for an Aurora DB cluster in `serverless` DB engine mode.

For Aurora MySQL, valid capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, and `256`.

For Aurora PostgreSQL, valid capacity values are `2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`.

The maximum capacity must be greater than or equal to the minimum capacity.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinCapacity`

The minimum capacity for an Aurora DB cluster in `serverless` DB engine mode.

For Aurora MySQL, valid capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, and `256`.

For Aurora PostgreSQL, valid capacity values are `2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`.

The minimum capacity must be less than or equal to the maximum capacity.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondsBeforeTimeout`

The amount of time, in seconds, that Aurora Serverless v1 tries to find a scaling point
to perform seamless scaling before enforcing the timeout action. The default is 300.

Specify a value between 60 and 600 seconds.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondsUntilAutoPause`

The time, in seconds, before an Aurora DB cluster in `serverless` mode is paused.

Specify a value between 300 and 86,400 seconds.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutAction`

The action to take when the timeout is reached, either `ForceApplyCapacityChange` or `RollbackCapacityChange`.

`ForceApplyCapacityChange` sets the capacity to the specified value as soon as possible.

`RollbackCapacityChange`, the default, ignores the capacity change if a scaling point isn't found in the timeout period.

###### Important

If you specify `ForceApplyCapacityChange`, connections that
prevent Aurora Serverless v1 from finding a scaling point might be dropped.

For more information, see [Autoscaling for Aurora Serverless v1](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.auto-scaling) in the _Amazon Aurora User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example specifies a scaling configuration for a DB cluster.

### Specifying a scaling configuration

For a sample template that configures an Aurora Serverless DB cluster, see
[Creating an Amazon Aurora Serverless DB Cluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#aws-resource-rds-dbcluster--examples--Creating_an_Amazon_Aurora_Serverless_DB_Cluster).

#### JSON

```json

"ScalingConfiguration":{
   "AutoPause":true,
   "MinCapacity": 4,
   "MaxCapacity": 32,
   "SecondsUntilAutoPause":1000
}
```

#### YAML

```yaml

ScalingConfiguration:
  AutoPause: true
  MinCapacity: 4
  MaxCapacity: 32
  SecondsUntilAutoPause: 1000
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ReadEndpoint

ServerlessV2ScalingConfiguration
