This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBCluster ServerlessV2ScalingConfiguration

The `ServerlessV2ScalingConfiguration` property type specifies the scaling
configuration of an Aurora Serverless V2 DB cluster. For more information, see [Using Amazon Aurora Serverless v2](../../../amazonrds/latest/aurorauserguide/aurora-serverless-v2.md) in the _Amazon Aurora User_
_Guide_.

If you have an Aurora cluster, you must set this attribute before you add a DB instance
that uses the `db.serverless` DB instance class. For more information, see
[Clusters that use Aurora Serverless v2 must have a capacity range specified](../../../amazonrds/latest/aurorauserguide/aurora-serverless-v2-requirements.md#aurora-serverless-v2.requirements.capacity-range) in
the _Amazon Aurora User Guide_.

This property is only supported for Aurora Serverless v2. For Aurora Serverless v1, use the `ScalingConfiguration` property.

Valid for: Aurora Serverless v2 DB clusters

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxCapacity" : Number,
  "MinCapacity" : Number,
  "SecondsUntilAutoPause" : Integer
}

```

### YAML

```yaml

  MaxCapacity: Number
  MinCapacity: Number
  SecondsUntilAutoPause: Integer

```

## Properties

`MaxCapacity`

The maximum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster.
You can specify ACU values in half-step increments, such as 40, 40.5, 41, and so on. The largest value
that you can use is 128.

The maximum capacity must be higher than 0.5 ACUs. For more information, see [Choosing the maximum Aurora Serverless v2 capacity setting for a cluster](../../../amazonrds/latest/aurorauserguide/aurora-serverless-v2-setting-capacity.md#aurora-serverless-v2.max_capacity_considerations) in the
_Amazon Aurora User Guide_.

Aurora automatically sets certain parameters for Aurora Serverless V2 DB instances to
values that depend on the maximum ACU value in the capacity range. When you update the
maximum capacity value, the `ParameterApplyStatus` value for the DB instance
changes to `pending-reboot`. You can update the parameter values by rebooting
the DB instance after changing the capacity range.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinCapacity`

The minimum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster.
You can specify ACU values in half-step increments, such as 8, 8.5, 9, and so on.
For Aurora versions that support the Aurora Serverless v2 auto-pause feature, the smallest value that you can use is 0.
For versions that don't support Aurora Serverless v2 auto-pause, the smallest value that you can use is 0.5.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondsUntilAutoPause`

Specifies the number of seconds an Aurora Serverless v2 DB instance must be idle before
Aurora attempts to automatically pause it.

Specify a value between 300 seconds (five minutes) and 86,400 seconds (one day). The default is 300 seconds.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example specifies a scaling configuration for an Aurora Serverless v2 DB cluster.

### Specifying a scaling configuration for a Serverless v2 DB cluster

#### JSON

```json

"ServerlessV2ScalingConfiguration":{
   "MinCapacity": 9,
   "MaxCapacity": 42
}
```

#### YAML

```yaml

ServerlessV2ScalingConfiguration:
  MinCapacity: 9
  MaxCapacity: 42
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScalingConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
