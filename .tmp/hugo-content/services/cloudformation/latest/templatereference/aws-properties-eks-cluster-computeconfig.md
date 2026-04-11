This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster ComputeConfig

Indicates the current configuration of the compute capability on your EKS Auto Mode
cluster. For example, if the capability is enabled or disabled. If the compute
capability is enabled, EKS Auto Mode will create and delete EC2 Managed Instances in
your AWS account. For more information, see EKS Auto Mode compute capability in the
_Amazon EKS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "NodePools" : [ String, ... ],
  "NodeRoleArn" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  NodePools:
    - String
  NodeRoleArn: String

```

## Properties

`Enabled`

Request to enable or disable the compute capability on your EKS Auto Mode cluster. If
the compute capability is enabled, EKS Auto Mode will create and delete EC2 Managed
Instances in your AWS account.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodePools`

Configuration for node pools that defines the compute resources for your EKS Auto Mode
cluster. For more information, see EKS Auto Mode Node Pools in the
_Amazon EKS User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeRoleArn`

The ARN of the IAM Role EKS will assign to EC2 Managed Instances in your EKS Auto
Mode cluster. This value cannot be changed after the compute capability of EKS Auto Mode
is enabled. For more information, see the IAM Reference in the
_Amazon EKS User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterLogging

ControlPlanePlacement

All content copied from https://docs.aws.amazon.com/.
