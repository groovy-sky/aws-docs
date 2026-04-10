This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster JobFlowInstancesConfig

`JobFlowInstancesConfig` is a property of the `AWS::EMR::Cluster` resource. `JobFlowInstancesConfig` defines the instance groups or instance fleets that comprise the cluster. `JobFlowInstancesConfig` must contain either `InstanceFleetConfig` or `InstanceGroupConfig`. They cannot be used together.

You can now define task instance groups or task instance fleets using the
`TaskInstanceGroups` and `TaskInstanceFleets` subproperties. Using
these subproperties reduces delays in provisioning task nodes compared to specifying task
nodes with the `InstanceFleetConfig` and `InstanceGroupConfig`
resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalMasterSecurityGroups" : [ String, ... ],
  "AdditionalSlaveSecurityGroups" : [ String, ... ],
  "CoreInstanceFleet" : InstanceFleetConfig,
  "CoreInstanceGroup" : InstanceGroupConfig,
  "Ec2KeyName" : String,
  "Ec2SubnetId" : String,
  "Ec2SubnetIds" : [ String, ... ],
  "EmrManagedMasterSecurityGroup" : String,
  "EmrManagedSlaveSecurityGroup" : String,
  "HadoopVersion" : String,
  "KeepJobFlowAliveWhenNoSteps" : Boolean,
  "MasterInstanceFleet" : InstanceFleetConfig,
  "MasterInstanceGroup" : InstanceGroupConfig,
  "Placement" : PlacementType,
  "ServiceAccessSecurityGroup" : String,
  "TaskInstanceFleets" : [ InstanceFleetConfig, ... ],
  "TaskInstanceGroups" : [ InstanceGroupConfig, ... ],
  "TerminationProtected" : Boolean,
  "UnhealthyNodeReplacement" : Boolean
}

```

### YAML

```yaml

  AdditionalMasterSecurityGroups:
    - String
  AdditionalSlaveSecurityGroups:
    - String
  CoreInstanceFleet:
    InstanceFleetConfig
  CoreInstanceGroup:
    InstanceGroupConfig
  Ec2KeyName: String
  Ec2SubnetId: String
  Ec2SubnetIds:
    - String
  EmrManagedMasterSecurityGroup: String
  EmrManagedSlaveSecurityGroup: String
  HadoopVersion: String
  KeepJobFlowAliveWhenNoSteps: Boolean
  MasterInstanceFleet:
    InstanceFleetConfig
  MasterInstanceGroup:
    InstanceGroupConfig
  Placement:
    PlacementType
  ServiceAccessSecurityGroup: String
  TaskInstanceFleets:
    - InstanceFleetConfig
  TaskInstanceGroups:
    - InstanceGroupConfig
  TerminationProtected: Boolean
  UnhealthyNodeReplacement: Boolean

```

## Properties

`AdditionalMasterSecurityGroups`

A list of additional Amazon EC2 security group IDs for the master node.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdditionalSlaveSecurityGroups`

A list of additional Amazon EC2 security group IDs for the core and task
nodes.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CoreInstanceFleet`

Describes the EC2 instances and instance configurations for the core instance fleet when using clusters with the instance fleet configuration.

_Required_: No

_Type_: [InstanceFleetConfig](aws-properties-emr-cluster-instancefleetconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CoreInstanceGroup`

Describes the EC2 instances and instance configurations for core instance groups when using clusters with the uniform instance group configuration.

_Required_: No

_Type_: [InstanceGroupConfig](aws-properties-emr-cluster-instancegroupconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ec2KeyName`

The name of the Amazon EC2 key pair that can be used to connect to the master
node using SSH as the user called "hadoop."

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ec2SubnetId`

Applies to clusters that use the uniform instance group configuration. To launch the
cluster in Amazon Virtual Private Cloud (Amazon VPC), set this parameter to the
identifier of the Amazon VPC subnet where you want the cluster to launch. If you do
not specify this value and your account supports EC2-Classic, the cluster launches in
EC2-Classic.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ec2SubnetIds`

Applies to clusters that use the instance fleet configuration. When multiple Amazon EC2 subnet IDs are specified, Amazon EMR evaluates them and launches
instances in the optimal subnet.

###### Note

The instance fleet configuration is available only in Amazon EMR releases
4.8.0 and later, excluding 5.0.x versions.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmrManagedMasterSecurityGroup`

The identifier of the Amazon EC2 security group for the master node. If you
specify `EmrManagedMasterSecurityGroup`, you must also specify
`EmrManagedSlaveSecurityGroup`.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmrManagedSlaveSecurityGroup`

The identifier of the Amazon EC2 security group for the core and task nodes. If
you specify `EmrManagedSlaveSecurityGroup`, you must also specify
`EmrManagedMasterSecurityGroup`.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HadoopVersion`

Applies only to Amazon EMR release versions earlier than 4.0. The Hadoop version
for the cluster. Valid inputs are "0.18" (no longer maintained), "0.20" (no longer
maintained), "0.20.205" (no longer maintained), "1.0.3", "2.2.0", or "2.4.0". If you do not
set this value, the default of 0.18 is used, unless the `AmiVersion` parameter
is set in the RunJobFlow call, in which case the default version of Hadoop for that AMI
version is used.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeepJobFlowAliveWhenNoSteps`

Specifies whether the cluster should remain available after completing all steps.
Defaults to `false`. For more information about configuring cluster termination,
see [Control Cluster Termination](../../../emr/latest/managementguide/emr-plan-termination.md) in the _EMR Management_
_Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterInstanceFleet`

Describes the EC2 instances and instance configurations for the master instance fleet when using clusters with the instance fleet configuration.

_Required_: No

_Type_: [InstanceFleetConfig](aws-properties-emr-cluster-instancefleetconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterInstanceGroup`

Describes the EC2 instances and instance configurations for the master instance group when using clusters with the uniform instance group configuration.

_Required_: No

_Type_: [InstanceGroupConfig](aws-properties-emr-cluster-instancegroupconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Placement`

The Availability Zone in which the cluster runs.

_Required_: No

_Type_: [PlacementType](aws-properties-emr-cluster-placementtype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceAccessSecurityGroup`

The identifier of the Amazon EC2 security group for the Amazon EMR
service to access clusters in VPC private subnets.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskInstanceFleets`

Describes the EC2 instances and instance configurations for the task instance fleets when using clusters with the instance fleet configuration. These task instance fleets are added to the cluster as part of the cluster launch. Each task instance fleet must have a unique name specified so that CloudFormation can differentiate between the task instance fleets.

###### Note

You can currently specify only one task instance fleet for a cluster. After creating the cluster, you can only modify the mutable properties of `InstanceFleetConfig`, which are `TargetOnDemandCapacity` and `TargetSpotCapacity`. Modifying any other property results in cluster replacement.

###### Important

To allow a maximum of 30 Amazon EC2 instance types per fleet, include `TaskInstanceFleets` when you create your cluster.
If you create your cluster without `TaskInstanceFleets`, Amazon EMR uses its default allocation strategy, which allows for a maximum of five Amazon EC2 instance types.

_Required_: No

_Type_: Array of [InstanceFleetConfig](aws-properties-emr-cluster-instancefleetconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskInstanceGroups`

Describes the EC2 instances and instance configurations for task instance groups when using clusters with the uniform instance group configuration. These task instance groups are added to the cluster as part of the cluster launch. Each task instance group must have a unique name specified so that CloudFormation can differentiate between the task instance groups.

###### Note

After creating the cluster, you can only modify the mutable properties of `InstanceGroupConfig`, which are `AutoScalingPolicy` and `InstanceCount`. Modifying any other property results in cluster replacement.

_Required_: No

_Type_: Array of [InstanceGroupConfig](aws-properties-emr-cluster-instancegroupconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TerminationProtected`

Specifies whether to lock the cluster to prevent the Amazon EC2 instances from
being terminated by API call, user intervention, or in the event of a job-flow
error.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnhealthyNodeReplacement`

Indicates whether Amazon EMR should gracefully replace core nodes
that have degraded within the cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceTypeConfig

KerberosAttributes

All content copied from https://docs.aws.amazon.com/.
