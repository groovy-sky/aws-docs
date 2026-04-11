This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::CapacityProvider TargetTrackingScalingPolicy

A scaling policy for the capacity provider that automatically adjusts capacity to maintain a target value for a specific metric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PredefinedMetricType" : String,
  "TargetValue" : Number
}

```

### YAML

```yaml

  PredefinedMetricType: String
  TargetValue: Number

```

## Properties

`PredefinedMetricType`

The predefined metric type to track for scaling decisions.

_Required_: Yes

_Type_: String

_Allowed values_: `LambdaCapacityProviderAverageCPUUtilization`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetValue`

The target value for the metric that the scaling policy attempts to maintain through scaling actions.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Lambda::CodeSigningConfig

All content copied from https://docs.aws.amazon.com/.
