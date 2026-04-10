This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Fleet TargetTrackingScalingConfiguration

Defines when a new instance is auto-scaled into the compute fleet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetricType" : String,
  "TargetValue" : Number
}

```

### YAML

```yaml

  MetricType: String
  TargetValue: Number

```

## Properties

`MetricType`

The metric type to determine auto-scaling.

_Required_: No

_Type_: String

_Allowed values_: `FLEET_UTILIZATION_RATE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetValue`

The value of `metricType` when to start scaling.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

VpcConfig

All content copied from https://docs.aws.amazon.com/.
