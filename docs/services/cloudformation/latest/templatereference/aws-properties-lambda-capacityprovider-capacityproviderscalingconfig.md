This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::CapacityProvider CapacityProviderScalingConfig

Configuration that defines how the capacity provider scales compute instances based on demand and policies.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxVCpuCount" : Integer,
  "ScalingMode" : String,
  "ScalingPolicies" : [ TargetTrackingScalingPolicy, ... ]
}

```

### YAML

```yaml

  MaxVCpuCount: Integer
  ScalingMode: String
  ScalingPolicies:
    - TargetTrackingScalingPolicy

```

## Properties

`MaxVCpuCount`

The maximum number of vCPUs that the capacity provider can provision across all compute instances.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `15000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingMode`

The scaling mode that determines how the capacity provider responds to changes in demand.

_Required_: No

_Type_: String

_Allowed values_: `Auto | Manual`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingPolicies`

A list of target tracking scaling policies for the capacity provider.

_Required_: No

_Type_: Array of [TargetTrackingScalingPolicy](aws-properties-lambda-capacityprovider-targettrackingscalingpolicy.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityProviderPermissionsConfig

CapacityProviderVpcConfig

All content copied from https://docs.aws.amazon.com/.
