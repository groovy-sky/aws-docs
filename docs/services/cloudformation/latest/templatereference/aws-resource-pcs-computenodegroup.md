This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::ComputeNodeGroup

Creates an AWS PCS compute node group resource. For more information, see
[Creating a compute node group\
in AWS PCS](../../../pcs/latest/userguide/working-with-cng-create.md) in the _AWS PCS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PCS::ComputeNodeGroup",
  "Properties" : {
      "AmiId" : String,
      "ClusterId" : String,
      "CustomLaunchTemplate" : CustomLaunchTemplate,
      "IamInstanceProfileArn" : String,
      "InstanceConfigs" : [ InstanceConfig, ... ],
      "Name" : String,
      "PurchaseOption" : String,
      "ScalingConfiguration" : ScalingConfiguration,
      "SlurmConfiguration" : SlurmConfiguration,
      "SpotOptions" : SpotOptions,
      "SubnetIds" : [ String, ... ],
      "Tags" : String
    }
}

```

### YAML

```yaml

Type: AWS::PCS::ComputeNodeGroup
Properties:
  AmiId: String
  ClusterId: String
  CustomLaunchTemplate:
    CustomLaunchTemplate
  IamInstanceProfileArn: String
  InstanceConfigs:
    - InstanceConfig
  Name: String
  PurchaseOption: String
  ScalingConfiguration:
    ScalingConfiguration
  SlurmConfiguration:
    SlurmConfiguration
  SpotOptions:
    SpotOptions
  SubnetIds:
    - String
  Tags: String

```

## Properties

`AmiId`

The ID of the Amazon Machine Image (AMI) that AWS PCS uses to launch
instances. If not provided, AWS PCS uses the AMI ID specified in the custom
launch template.

_Required_: No

_Type_: String

_Pattern_: `^ami-[a-z0-9]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterId`

The ID of the cluster of the compute node group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomLaunchTemplate`

An Amazon EC2 launch template AWS PCS uses to launch compute
nodes.

_Required_: Yes

_Type_: [CustomLaunchTemplate](aws-properties-pcs-computenodegroup-customlaunchtemplate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamInstanceProfileArn`

The Amazon Resource Name (ARN) of the IAM instance profile used to pass an IAM role when
launching EC2 instances. The role contained in your instance profile must have the
`pcs:RegisterComputeNodeGroupInstance` permission and the role name must
start with `AWSPCS` or must have the path `/aws-pcs/`. For more
information, see [IAM instance profiles for\
AWS PCS](../../../pcs/latest/userguide/security-instance-profiles.md) in the _AWS PCS User_
_Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws([a-zA-Z-]{0,10})?:iam::[0-9]{12}:instance-profile/.{1,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceConfigs`

A list of EC2 instance configurations that AWS PCS can provision in the compute node group.

_Required_: Yes

_Type_: Array of [InstanceConfig](aws-properties-pcs-computenodegroup-instanceconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name that identifies the compute node group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PurchaseOption`

Specifies how EC2 instances are purchased on your behalf. AWS PCS supports On-Demand Instances, Spot Instances,
and Amazon EC2 Capacity Blocks for ML. For more information, see [Amazon EC2 billing and purchasing options](../../../ec2/latest/userguide/instance-purchasing-options.md) in the _Amazon Elastic Compute Cloud User_
_Guide_. For more information about AWS PCS support for
Capacity Blocks, see [Using Amazon EC2 Capacity\
Blocks for ML with AWS PCS](../../../pcs/latest/userguide/capacity-blocks.md)
in the _AWS PCS User Guide_.
If you don't provide this option, it defaults to On-Demand.

_Required_: No

_Type_: String

_Allowed values_: `ONDEMAND | SPOT | CAPACITY_BLOCK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingConfiguration`

Specifies the boundaries of the compute node group auto scaling.

_Required_: Yes

_Type_: [ScalingConfiguration](aws-properties-pcs-computenodegroup-scalingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlurmConfiguration`

Additional options related to the Slurm scheduler.

_Required_: No

_Type_: [SlurmConfiguration](aws-properties-pcs-computenodegroup-slurmconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpotOptions`

Additional configuration when you specify `SPOT` as the
`purchaseOption` for the `CreateComputeNodeGroup` API
action.

_Required_: No

_Type_: [SpotOptions](aws-properties-pcs-computenodegroup-spotoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

The list of subnet IDs where instances are provisioned by the compute node group. The
subnets must be in the same VPC as the cluster.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

1 or more tags added to the resource. Each tag consists of a tag key and tag value. The
tag value is optional and can be an empty string.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The unique Amazon Resource Name (ARN) of the compute node group.

`ErrorInfo`

The list of errors that occurred during compute node group provisioning.

`Id`

The generated unique ID of the compute node group.

`Status`

The provisioning status of the compute node group.

###### Note

The provisioning status doesn't indicate the overall health of the compute node
group.

###### Important

The resource enters the `SUSPENDING` and `SUSPENDED` states
when the scheduler is beyond end of life and we have suspended the cluster. When in
these states, you can't use the cluster. The cluster controller is down and all compute
instances are terminated. The resources still count toward your service quotas. You can
delete a resource if its status is `SUSPENDED`. For more information, see
[Frequently asked questions about Slurm versions in AWS PCS](../../../pcs/latest/userguide/slurm-versions-faq.md) in the
_AWS PCS User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlurmRest

CustomLaunchTemplate

All content copied from https://docs.aws.amazon.com/.
