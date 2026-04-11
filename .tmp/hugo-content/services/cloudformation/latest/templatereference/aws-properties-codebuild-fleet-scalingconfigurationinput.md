This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Fleet ScalingConfigurationInput

The scaling configuration input of a compute fleet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxCapacity" : Integer,
  "ScalingType" : String,
  "TargetTrackingScalingConfigs" : [ TargetTrackingScalingConfiguration, ... ]
}

```

### YAML

```yaml

  MaxCapacity: Integer
  ScalingType: String
  TargetTrackingScalingConfigs:
    - TargetTrackingScalingConfiguration

```

## Properties

`MaxCapacity`

The maximum number of instances in the ﬂeet when auto-scaling.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingType`

The scaling type for a compute fleet.

_Required_: No

_Type_: String

_Allowed values_: `TARGET_TRACKING_SCALING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetTrackingScalingConfigs`

A list of `TargetTrackingScalingConfiguration` objects.

_Required_: No

_Type_: Array of [TargetTrackingScalingConfiguration](aws-properties-codebuild-fleet-targettrackingscalingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProxyConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
