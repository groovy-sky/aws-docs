This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceGroupConfig ScalingAction

`ScalingAction` is a subproperty of the `ScalingRule` property type. `ScalingAction` determines the type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Market" : String,
  "SimpleScalingPolicyConfiguration" : SimpleScalingPolicyConfiguration
}

```

### YAML

```yaml

  Market: String
  SimpleScalingPolicyConfiguration:
    SimpleScalingPolicyConfiguration

```

## Properties

`Market`

Not available for instance groups. Instance groups use the market type specified for the
group.

_Required_: No

_Type_: String

_Allowed values_: `ON_DEMAND | SPOT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SimpleScalingPolicyConfiguration`

The type of adjustment the automatic scaling activity makes when triggered, and the
periodicity of the adjustment.

_Required_: Yes

_Type_: [SimpleScalingPolicyConfiguration](aws-properties-emr-instancegroupconfig-simplescalingpolicyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricDimension

ScalingConstraints

All content copied from https://docs.aws.amazon.com/.
