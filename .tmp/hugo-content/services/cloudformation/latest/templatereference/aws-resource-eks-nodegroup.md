This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Nodegroup

Creates a managed node group for an Amazon EKS cluster.

You can only create a node group for your cluster that is equal to the current Kubernetes
version for the cluster. All node groups are created with the latest AMI release version
for the respective minor Kubernetes version of the cluster, unless you deploy a custom AMI
using a launch template.

For later updates, you will only be able to update a node group using a launch
template only if it was originally deployed with a launch template. Additionally, the
launch template ID or name must match what was used when the node group was created. You
can update the launch template version with necessary changes. For more information
about using launch templates, see [Customizing managed nodes with\
launch templates](../../../eks/latest/userguide/launch-templates.md).

An Amazon EKS managed node group is an Amazon EC2 Auto Scaling group and associated Amazon EC2 instances that
are managed by AWS for an Amazon EKS cluster. For more information, see [Managed\
node groups](../../../eks/latest/userguide/managed-node-groups.md) in the _Amazon EKS User Guide_.

###### Note

Windows AMI types are only supported for commercial AWS Regions that support
Windows on Amazon EKS.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EKS::Nodegroup",
  "Properties" : {
      "AmiType" : String,
      "CapacityType" : String,
      "ClusterName" : String,
      "DiskSize" : Integer,
      "ForceUpdateEnabled" : Boolean,
      "InstanceTypes" : [ String, ... ],
      "Labels" : {Key: Value, ...},
      "LaunchTemplate" : LaunchTemplateSpecification,
      "NodegroupName" : String,
      "NodeRepairConfig" : NodeRepairConfig,
      "NodeRole" : String,
      "ReleaseVersion" : String,
      "RemoteAccess" : RemoteAccess,
      "ScalingConfig" : ScalingConfig,
      "Subnets" : [ String, ... ],
      "Tags" : {Key: Value, ...},
      "Taints" : [ Taint, ... ],
      "UpdateConfig" : UpdateConfig,
      "Version" : String,
      "WarmPoolConfig" : WarmPoolConfig
    }
}

```

### YAML

```yaml

Type: AWS::EKS::Nodegroup
Properties:
  AmiType: String
  CapacityType: String
  ClusterName: String
  DiskSize: Integer
  ForceUpdateEnabled: Boolean
  InstanceTypes:
    - String
  Labels:
    Key: Value
  LaunchTemplate:
    LaunchTemplateSpecification
  NodegroupName: String
  NodeRepairConfig:
    NodeRepairConfig
  NodeRole: String
  ReleaseVersion: String
  RemoteAccess:
    RemoteAccess
  ScalingConfig:
    ScalingConfig
  Subnets:
    - String
  Tags:
    Key: Value
  Taints:
    - Taint
  UpdateConfig:
    UpdateConfig
  Version: String
  WarmPoolConfig:
    WarmPoolConfig

```

## Properties

`AmiType`

The AMI type for your node group. If you specify `launchTemplate`, and your launch template uses a custom AMI,
then don't specify `amiType`, or the node group deployment
will fail. If your launch template uses a Windows custom AMI, then add
`eks:kube-proxy-windows` to your Windows nodes `rolearn` in
the `aws-auth` `ConfigMap`. For more information about using launch templates with Amazon EKS, see [Customizing managed nodes with launch templates](../../../eks/latest/userguide/launch-templates.md) in the _Amazon EKS User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `AL2_x86_64 | AL2_x86_64_GPU | AL2_ARM_64 | CUSTOM | BOTTLEROCKET_ARM_64 | BOTTLEROCKET_x86_64 | BOTTLEROCKET_ARM_64_FIPS | BOTTLEROCKET_x86_64_FIPS | BOTTLEROCKET_ARM_64_NVIDIA | BOTTLEROCKET_x86_64_NVIDIA | BOTTLEROCKET_ARM_64_NVIDIA_FIPS | BOTTLEROCKET_x86_64_NVIDIA_FIPS | WINDOWS_CORE_2019_x86_64 | WINDOWS_FULL_2019_x86_64 | WINDOWS_CORE_2022_x86_64 | WINDOWS_FULL_2022_x86_64 | WINDOWS_CORE_2025_x86_64 | WINDOWS_FULL_2025_x86_64 | AL2023_x86_64_STANDARD | AL2023_ARM_64_STANDARD | AL2023_x86_64_NEURON | AL2023_x86_64_NVIDIA | AL2023_ARM_64_NVIDIA`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CapacityType`

The capacity type of your managed node group.

_Required_: No

_Type_: String

_Allowed values_: `ON_DEMAND | SPOT | CAPACITY_BLOCK`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterName`

The name of your cluster.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DiskSize`

The root device disk size (in GiB) for your node group instances. The default disk
size is 20 GiB for Linux and Bottlerocket. The default disk size is 50 GiB for Windows.
If you specify `launchTemplate`, then don't specify `diskSize`, or the node group
deployment will fail. For more information about using launch templates with Amazon EKS, see [Customizing managed nodes with launch templates](../../../eks/latest/userguide/launch-templates.md) in the _Amazon EKS User Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ForceUpdateEnabled`

Force the update if any `Pod` on the existing node group can't be drained
due to a `Pod` disruption budget issue. If an update fails because all Pods
can't be drained, you can force the update after it fails to terminate the old node
whether or not any `Pod` is running on the node.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceTypes`

Specify the instance types for a node group. If you specify a GPU instance type, make
sure to also specify an applicable GPU AMI type with the `amiType` parameter.
If you specify `launchTemplate`, then you can specify zero or one instance
type in your launch template _or_ you can specify 0-20 instance types
for `instanceTypes`. If however, you specify an instance type in your launch
template _and_ specify any `instanceTypes`, the node group
deployment will fail. If you don't specify an instance type in a launch template or for
`instanceTypes`, then `t3.medium` is used, by default. If you
specify `Spot` for `capacityType`, then we recommend specifying
multiple values for `instanceTypes`. For more information, see [Managed node group capacity types](../../../eks/latest/userguide/managed-node-groups.md#managed-node-group-capacity-types) and [Customizing managed nodes with\
launch templates](../../../eks/latest/userguide/launch-templates.md) in the _Amazon EKS User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Labels`

The Kubernetes `labels` applied to the nodes in the node group.

###### Note

Only `labels` that are applied with the Amazon EKS API are shown here. There
may be other Kubernetes `labels` applied to the nodes in this group.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchTemplate`

An object representing a node group's launch template specification. When using this
object, don't directly specify `instanceTypes`, `diskSize`, or
`remoteAccess`. You cannot later specify a different launch template ID
or name than what was used to create the node group.

Make sure that the launch template meets the requirements in
`launchTemplateSpecification`. Also refer to [Customizing managed nodes with\
launch templates](../../../eks/latest/userguide/launch-templates.md) in the _Amazon EKS User Guide_.

_Required_: No

_Type_: [LaunchTemplateSpecification](aws-properties-eks-nodegroup-launchtemplatespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodegroupName`

The unique name to give your node group.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NodeRepairConfig`

The node auto repair configuration for the node group.

_Required_: No

_Type_: [NodeRepairConfig](aws-properties-eks-nodegroup-noderepairconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeRole`

The Amazon Resource Name (ARN) of the IAM role to associate with your node group. The Amazon EKS worker
node `kubelet` daemon makes calls to AWS APIs on your behalf. Nodes receive
permissions for these API calls through an IAM instance profile and associated
policies. Before you can launch nodes and register them into a cluster, you must create
an IAM role for those nodes to use when they are launched. For more information, see
[Amazon EKS\
node IAM role](../../../eks/latest/userguide/create-node-role.md) in the _Amazon EKS User Guide_.
If you specify `launchTemplate`, then don't specify `
                            IamInstanceProfile
                        ` in your launch template, or the node group
deployment will fail. For more information about using launch templates with Amazon EKS, see [Customizing managed nodes with launch templates](../../../eks/latest/userguide/launch-templates.md) in the _Amazon EKS User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReleaseVersion`

The AMI version of the Amazon EKS optimized AMI to use with your node group
(for example, `1.14.7-YYYYMMDD`). By default, the latest
available AMI version for the node group's current Kubernetes version is used. For more
information, see [Amazon EKS optimized\
Linux AMI Versions](../../../eks/latest/userguide/eks-linux-ami-versions.md) in the _Amazon EKS User_
_Guide_.

###### Note

Changing this value triggers an update of the node group if one is available. You
can't update other properties at the same time as updating `Release
                    Version`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoteAccess`

The remote access configuration to use with your node group. For Linux, the protocol
is SSH. For Windows, the protocol is RDP. If you specify `launchTemplate`, then don't specify
`remoteAccess`, or the node group deployment will fail.
For more information about using launch templates with Amazon EKS, see [Customizing managed nodes with launch templates](../../../eks/latest/userguide/launch-templates.md) in the _Amazon EKS User Guide_.

_Required_: No

_Type_: [RemoteAccess](aws-properties-eks-nodegroup-remoteaccess.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScalingConfig`

The scaling configuration details for the Auto Scaling group that is created for your
node group.

_Required_: No

_Type_: [ScalingConfig](aws-properties-eks-nodegroup-scalingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subnets`

The subnets to use for the Auto Scaling group that is created for your node group.
If you specify `launchTemplate`, then don't specify `
                            SubnetId
                        ` in your launch template, or the node group deployment
will fail. For more information about using launch templates with Amazon EKS, see [Customizing managed nodes with launch templates](../../../eks/latest/userguide/launch-templates.md) in the _Amazon EKS User Guide_.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata that assists with categorization and organization.
Each tag consists of a key and an optional value. You define both. Tags don't
propagate to any other cluster or AWS resources.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Taints`

The Kubernetes taints to be applied to the nodes in the node group when they are created.
Effect is one of `No_Schedule`, `Prefer_No_Schedule`, or
`No_Execute`. Kubernetes taints can be used together with tolerations to
control how workloads are scheduled to your nodes. For more information, see [Node taints on managed node groups](../../../eks/latest/userguide/node-taints-managed-node-groups.md).

_Required_: No

_Type_: Array of [Taint](aws-properties-eks-nodegroup-taint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdateConfig`

The node group update configuration.

_Required_: No

_Type_: [UpdateConfig](aws-properties-eks-nodegroup-updateconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The Kubernetes version to use for your managed nodes. By default, the Kubernetes
version of the cluster is used, and this is the only accepted specified value. If you
specify `launchTemplate`, and your launch template uses a custom AMI, then
don't specify `version`, or the node group deployment will fail. For more
information about using launch templates with Amazon EKS, see [Launch\
template support](../../../eks/latest/userguide/launch-templates.md) in the _Amazon EKS User Guide_.

###### Note

You can't update other properties at the same time as updating
`Version`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WarmPoolConfig`

Property description not available.

_Required_: No

_Type_: [WarmPoolConfig](aws-properties-eks-nodegroup-warmpoolconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myNodegroup" }`

For the Amazon EKS node group `myNodegroup`, Ref returns the
physical resource ID of the node group. For example,
`cluster-name/nodegroup_name`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) associated with the managed node group.

`ClusterName`

The name of your cluster.

`NodegroupName`

The name associated with an Amazon EKS managed node group.

## Examples

### Create a managed node group

The following example creates an Amazon EKS managed node group named
`standard` in the `prod` cluster.

#### JSON

```json

{
    "Resources": {
        "EKSNodegroup": {
            "Type": "AWS::EKS::Nodegroup",
            "Properties": {
                "ClusterName": "prod",
                "NodeRole": "arn:aws:iam::012345678910:role/eksInstanceRole",
                "ScalingConfig": {
                    "MinSize": 3,
                    "DesiredSize": 5,
                    "MaxSize": 7
                },
                "Labels": {
                    "Key1": "Value1",
                    "Key2": "Value2"
                },
                "Subnets": [
                    "subnet-6782e71e",
                    "subnet-e7e761ac"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  EKSNodegroup:
    Type: 'AWS::EKS::Nodegroup'
    Properties:
      ClusterName: prod
      NodeRole: 'arn:aws:iam::012345678910:role/eksInstanceRole'
      ScalingConfig:
        MinSize: 3
        DesiredSize: 5
        MaxSize: 7
      Labels:
        Key1: Value1
        Key2: Value2
      Subnets:
        - subnet-6782e71e
        - subnet-e7e761ac
```

## See also

- [Managed node\
groups](../../../eks/latest/userguide/managed-node-groups.md) in the _Amazon EKS User Guide_.

- [CreateNodegroup](../../../../reference/eks/latest/apireference/api-createnodegroup.md) in the _Amazon EKS API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

LaunchTemplateSpecification

All content copied from https://docs.aws.amazon.com/.
