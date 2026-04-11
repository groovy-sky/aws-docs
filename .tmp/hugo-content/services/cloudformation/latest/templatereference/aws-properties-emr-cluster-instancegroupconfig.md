This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster InstanceGroupConfig

Use `InstanceGroupConfig` to define instance groups for an EMR cluster. A cluster can not use both instance groups and instance fleets. For more information, see [Create a Cluster with Instance Fleets or Uniform Instance Groups](../../../emr/latest/managementguide/emr-instance-group-configuration.md) in the _Amazon EMR Management Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoScalingPolicy" : AutoScalingPolicy,
  "BidPrice" : String,
  "Configurations" : [ Configuration, ... ],
  "CustomAmiId" : String,
  "EbsConfiguration" : EbsConfiguration,
  "InstanceCount" : Integer,
  "InstanceType" : String,
  "Market" : String,
  "Name" : String
}

```

### YAML

```yaml

  AutoScalingPolicy:
    AutoScalingPolicy
  BidPrice: String
  Configurations:
    - Configuration
  CustomAmiId: String
  EbsConfiguration:
    EbsConfiguration
  InstanceCount: Integer
  InstanceType: String
  Market: String
  Name: String

```

## Properties

`AutoScalingPolicy`

`AutoScalingPolicy` is a subproperty of the [InstanceGroupConfig](../userguide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.md) property type that specifies the constraints and rules of an automatic scaling policy in Amazon EMR. The automatic scaling policy defines how an instance group dynamically adds and terminates EC2 instances in response to the value of a CloudWatch metric. Only core and task instance groups can use automatic scaling policies. For more information, see [Using Automatic Scaling in Amazon EMR](../../../emr/latest/managementguide/emr-automatic-scaling.md).

_Required_: No

_Type_: [AutoScalingPolicy](aws-properties-emr-cluster-autoscalingpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BidPrice`

If specified, indicates that the instance group uses Spot Instances. This is the maximum price you are willing to pay for Spot Instances. Specify `OnDemandPrice` to set the amount equal to the On-Demand price, or specify an amount in USD.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Configurations`

###### Note

Amazon EMR releases 4.x or later.

The list of configurations supplied for an Amazon EMR cluster instance group.
You can specify a separate configuration for each instance group (master, core, and
task).

_Required_: No

_Type_: Array of [Configuration](aws-properties-emr-cluster-configuration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomAmiId`

The custom AMI ID to use for the provisioned instance group.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EbsConfiguration`

EBS configurations that will be attached to each Amazon EC2 instance in the
instance group.

_Required_: No

_Type_: [EbsConfiguration](aws-properties-emr-cluster-ebsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceCount`

Target number of instances for the instance group.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

The Amazon EC2 instance type for all instances in the instance group.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Market`

Market type of the Amazon EC2 instances used to create a cluster node.

_Required_: No

_Type_: String

_Allowed values_: `ON_DEMAND | SPOT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Friendly name given to the instance group.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceFleetResizingSpecifications

InstanceTypeConfig

All content copied from https://docs.aws.amazon.com/.
