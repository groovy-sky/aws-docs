This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ClusterAutoScalingConfig

Specifies the autoscaling configuration for a HyperPod cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoScalerType" : String,
  "Mode" : String
}

```

### YAML

```yaml

  AutoScalerType: String
  Mode: String

```

## Properties

`AutoScalerType`

The type of autoscaler to use. Currently supported value is
`Karpenter`.

_Required_: No

_Type_: String

_Allowed values_: `Karpenter`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

Describes whether autoscaling is enabled or disabled for the cluster. Valid values are
`Enable` and `Disable`.

_Required_: Yes

_Type_: String

_Allowed values_: `Enable | Disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CapacitySizeConfig

ClusterCapacityRequirements
