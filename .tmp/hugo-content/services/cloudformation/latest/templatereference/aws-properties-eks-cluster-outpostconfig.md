This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster OutpostConfig

The configuration of your local Amazon EKS cluster on an AWS Outpost. Before creating a
cluster on an Outpost, review [Creating a local\
cluster on an Outpost](../../../eks/latest/userguide/eks-outposts-local-cluster-create.md) in the _Amazon EKS User Guide_. This API isn't available for
Amazon EKS clusters on the AWS cloud.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ControlPlaneInstanceType" : String,
  "ControlPlanePlacement" : ControlPlanePlacement,
  "OutpostArns" : [ String, ... ]
}

```

### YAML

```yaml

  ControlPlaneInstanceType: String
  ControlPlanePlacement:
    ControlPlanePlacement
  OutpostArns:
    - String

```

## Properties

`ControlPlaneInstanceType`

The Amazon EC2 instance type that you want to use for your local Amazon EKS cluster on Outposts.
Choose an instance type based on the number of nodes that your cluster will have. For
more information, see [Capacity\
considerations](../../../eks/latest/userguide/eks-outposts-capacity-considerations.md) in the _Amazon EKS User Guide_.

The instance type that you specify is used for all Kubernetes control plane instances. The
instance type can't be changed after cluster creation. The control plane is not
automatically scaled by Amazon EKS.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ControlPlanePlacement`

An object representing the placement configuration for all the control plane instances
of your local Amazon EKS cluster on an AWS Outpost. For more information, see [Capacity considerations](../../../eks/latest/userguide/eks-outposts-capacity-considerations.md) in the _Amazon EKS User Guide_.

_Required_: No

_Type_: [ControlPlanePlacement](aws-properties-eks-cluster-controlplaneplacement.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutpostArns`

The ARN of the Outpost that you want to use for your local Amazon EKS cluster on
Outposts. Only a single Outpost ARN is supported.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggingTypeConfig

Provider

All content copied from https://docs.aws.amazon.com/.
