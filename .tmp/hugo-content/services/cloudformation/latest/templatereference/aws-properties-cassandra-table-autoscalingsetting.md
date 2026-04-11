This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Table AutoScalingSetting

The optional auto scaling settings for a table with provisioned throughput capacity.

To turn on auto scaling for a table in `throughputMode:PROVISIONED`,
you must specify the following parameters.

Configure the minimum and maximum capacity units. The auto scaling policy ensures that capacity never goes below the
minimum or above the maximum range.

- `minimumUnits`: The minimum level of throughput the table should always be ready to support. The value must be between 1
and the max throughput per second quota for your account (40,000 by default).

- `maximumUnits`: The maximum level of throughput the table should always be ready to
support. The value must be between 1 and the max throughput per second quota for your
account (40,000 by default).

- `scalingPolicy`: Amazon Keyspaces supports the `target tracking` scaling policy.
The auto scaling target is a percentage of the provisioned capacity of the table.

For more information, see [Managing throughput capacity automatically with Amazon Keyspaces auto scaling](../../../keyspaces/latest/devguide/autoscaling.md) in the _Amazon Keyspaces Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoScalingDisabled" : Boolean,
  "MaximumUnits" : Integer,
  "MinimumUnits" : Integer,
  "ScalingPolicy" : ScalingPolicy
}

```

### YAML

```yaml

  AutoScalingDisabled: Boolean
  MaximumUnits: Integer
  MinimumUnits: Integer
  ScalingPolicy:
    ScalingPolicy

```

## Properties

`AutoScalingDisabled`

This optional parameter enables auto scaling for the table if set to `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumUnits`

Manage costs by specifying the maximum amount of throughput to provision. The value must be between 1
and the max throughput per second quota for your account (40,000 by default).

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumUnits`

The minimum level of throughput the table should always be ready to support. The value must be between 1
and the max throughput per second quota for your account (40,000 by default).

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingPolicy`

Amazon Keyspaces supports the `target tracking` auto scaling policy. With this policy, Amazon Keyspaces auto scaling
ensures that the table's ratio of consumed to provisioned capacity stays at or near the target value that you specify. You
define the target value as a percentage between 20 and 90.

_Required_: No

_Type_: [ScalingPolicy](aws-properties-cassandra-table-scalingpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cassandra::Table

AutoScalingSpecification

All content copied from https://docs.aws.amazon.com/.
