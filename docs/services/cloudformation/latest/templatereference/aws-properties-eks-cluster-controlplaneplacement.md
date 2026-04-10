This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster ControlPlanePlacement

The placement configuration for all the control plane instances of your local Amazon EKS
cluster on an AWS Outpost. For more information, see [Capacity\
considerations](../../../eks/latest/userguide/eks-outposts-capacity-considerations.md) in the _Amazon EKS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupName" : String
}

```

### YAML

```yaml

  GroupName: String

```

## Properties

`GroupName`

The name of the placement group for the Kubernetes control plane instances. This
property is only used for a local cluster on an AWS Outpost.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComputeConfig

ControlPlaneScalingConfig

All content copied from https://docs.aws.amazon.com/.
