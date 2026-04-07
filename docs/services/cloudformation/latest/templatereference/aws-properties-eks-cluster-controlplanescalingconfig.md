This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster ControlPlaneScalingConfig

The control plane scaling tier configuration. For more information, see EKS Provisioned Control Plane in the Amazon EKS User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Tier" : String
}

```

### YAML

```yaml

  Tier: String

```

## Properties

`Tier`

The control plane scaling tier configuration. Available options are `standard`, `tier-xl`, `tier-2xl`, `tier-4xl, or tier-8xl`. For more information, see EKS Provisioned Control Plane in the Amazon EKS User Guide.

_Required_: No

_Type_: String

_Allowed values_: `standard | tier-xl | tier-2xl | tier-4xl | tier-8xl | tier-ultra`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ControlPlanePlacement

ElasticLoadBalancing
