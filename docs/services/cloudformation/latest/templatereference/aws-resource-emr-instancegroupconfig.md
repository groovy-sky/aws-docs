This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceGroupConfig

Use `InstanceGroupConfig` to define instance groups for an EMR cluster. A cluster can not use both instance groups and instance fleets. For more information, see [Create a Cluster with Instance Fleets or Uniform Instance Groups](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html) in the _Amazon EMR Management Guide_.

###### Important

You can currently only add task instance groups to a cluster with this resource. If you use
this resource, CloudFormation waits for the cluster launch to complete before adding the
task instance group to the cluster. In order to add task instance groups to the cluster
as part of the cluster launch and minimize delays in provisioning task nodes, use the
`TaskInstanceGroups` subproperty for the [AWS::EMR::Cluster JobFlowInstancesConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html) property instead. To use this
subproperty, see [AWS::EMR::Cluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html) for examples.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EMR::InstanceGroupConfig",
  "Properties" : {
      "AutoScalingPolicy" : AutoScalingPolicy,
      "BidPrice" : String,
      "Configurations" : [ Configuration, ... ],
      "CustomAmiId" : String,
      "EbsConfiguration" : EbsConfiguration,
      "InstanceCount" : Integer,
      "InstanceRole" : String,
      "InstanceType" : String,
      "JobFlowId" : String,
      "Market" : String,
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::EMR::InstanceGroupConfig
Properties:
  AutoScalingPolicy:
    AutoScalingPolicy
  BidPrice: String
  Configurations:
    - Configuration
  CustomAmiId: String
  EbsConfiguration:
    EbsConfiguration
  InstanceCount: Integer
  InstanceRole: String
  InstanceType: String
  JobFlowId: String
  Market: String
  Name: String

```

## Properties

`AutoScalingPolicy`

`AutoScalingPolicy` is a subproperty of `InstanceGroupConfig`. `AutoScalingPolicy` defines how an instance group dynamically adds and terminates EC2 instances in response to the value of a CloudWatch metric. For more information, see [Using Automatic Scaling in Amazon EMR](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html) in the _Amazon EMR Management Guide_.

_Required_: No

_Type_: [AutoScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emr-instancegroupconfig-autoscalingpolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BidPrice`

If specified, indicates that the instance group uses Spot Instances. This is the maximum price you are willing to pay for Spot Instances. Specify `OnDemandPrice` to set the amount equal to the On-Demand price, or specify an amount in USD.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configurations`

###### Note

Amazon EMR releases 4.x or later.

The list of configurations supplied for an Amazon EMR cluster instance group.
You can specify a separate configuration for each instance group (master, core, and
task).

_Required_: No

_Type_: Array of [Configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emr-instancegroupconfig-configuration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomAmiId`

The custom AMI ID to use for the provisioned instance group.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EbsConfiguration`

`EbsConfiguration` determines the EBS volumes to attach to EMR cluster instances.

_Required_: No

_Type_: [EbsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emr-instancegroupconfig-ebsconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceCount`

Target number of instances for the instance group.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceRole`

The role of the instance group in the cluster.

_Allowed Values_: TASK

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceType`

The Amazon EC2 instance type for all instances in the instance group.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JobFlowId`

The ID of an Amazon EMR cluster that you want to associate this instance group with.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Market`

Market type of the Amazon EC2 instances used to create a cluster node.

_Required_: No

_Type_: String

_Allowed values_: `ON_DEMAND | SPOT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Friendly name given to the instance group.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns returns the ID of the instance group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

Example instance group configuration specifications.

- [Add a task instance group](#aws-resource-emr-instancegroupconfig--examples--Add_a_task_instance_group)

- [Specify an Automatic Scaling Policy](#aws-resource-emr-instancegroupconfig--examples--Specify_an_Automatic_Scaling_Policy)

### Add a task instance group

The following example adds a task instance group to a cluster named `TestCluster`.

#### JSON

```json

"TestInstanceGroupConfig": {
    "Type": "AWS: : EMR: : InstanceGroupConfig",
    "Properties": {
        "InstanceCount": 2,
        "InstanceType": "m3.xlarge",
        "InstanceRole": "TASK",
        "Market": "ON_DEMAND",
        "Name": "cfnTask2",
        "JobFlowId": {
            "Ref": "cluster"
        }
    }
}
```

#### YAML

```yaml

TestInstanceGroupConfig:
  Properties:
    InstanceCount: 2
    InstanceRole: TASK
    InstanceType: m3.xlarge
    JobFlowId:
      Ref: cluster
    Market: ON_DEMAND
    Name: cfnTask2
  Type: "AWS::EMR::InstanceGroupConfig"
```

### Specify an Automatic Scaling Policy

The following example defines an `AutoScalingPolicy` for an `InstanceGroupConfig` resource.

#### JSON

```json

"MyInstanceGroupConfig": {
  "Type": "AWS::EMR::InstanceGroupConfig",
  "Properties": {
    "InstanceCount": 1,
    "InstanceType": {
      "Ref": "InstanceType"
    },
    "InstanceRole": "TASK",
    "Market": "ON_DEMAND",
    "Name": "cfnTask",
    "JobFlowId": {
      "Ref": "MyCluster"
    },
    "AutoScalingPolicy": {
      "Constraints": {
        "MinCapacity": {
          "Ref": "MinCapacity"
        },
        "MaxCapacity": {
          "Ref": "MaxCapacity"
        }
      },
      "Rules": [
        {
          "Name": "Scale-out",
          "Description": "Scale-out policy",
          "Action": {
            "SimpleScalingPolicyConfiguration": {
              "AdjustmentType": "CHANGE_IN_CAPACITY",
              "ScalingAdjustment": 1,
              "CoolDown": 300
            }
          },
          "Trigger": {
            "CloudWatchAlarmDefinition": {
              "Dimensions": [
                {
                  "Key": "JobFlowId",
                  "Value": "${emr.clusterId}"
                }
              ],
              "EvaluationPeriods": 1,
              "Namespace": "AWS/ElasticMapReduce",
              "Period": 300,
              "ComparisonOperator": "LESS_THAN",
              "Statistic": "AVERAGE",
              "Threshold": 15,
              "Unit": "PERCENT",
              "MetricName": "YARNMemoryAvailablePercentage"
            }
          }
        },
        {
          "Name": "Scale-in",
          "Description": "Scale-in policy",
          "Action": {
            "SimpleScalingPolicyConfiguration": {
              "AdjustmentType": "CHANGE_IN_CAPACITY",
              "ScalingAdjustment": -1,
              "CoolDown": 300
            }
          },
          "Trigger": {
            "CloudWatchAlarmDefinition": {
              "Dimensions": [
                {
                  "Key": "JobFlowId",
                  "Value": "${emr.clusterId}"
                }
              ],
              "EvaluationPeriods": 1,
              "Namespace": "AWS/ElasticMapReduce",
              "Period": 300,
              "ComparisonOperator": "GREATER_THAN",
              "Statistic": "AVERAGE",
              "Threshold": 75,
              "Unit": "PERCENT",
              "MetricName": "YARNMemoryAvailablePercentage"
            }
          }
        }
      ]
    }
  }
}
```

#### YAML

```yaml

MyInstanceGroupConfig:
  Type: 'AWS::EMR::InstanceGroupConfig'
  Properties:
    InstanceCount: 1
    InstanceType: !Ref InstanceType
    InstanceRole: TASK
    Market: ON_DEMAND
    Name: cfnTask
    JobFlowId: !Ref MyCluster
    AutoScalingPolicy:
      Constraints:
        MinCapacity: !Ref MinCapacity
        MaxCapacity: !Ref MaxCapacity
      Rules:
        - Name: Scale-out
          Description: Scale-out policy
          Action:
            SimpleScalingPolicyConfiguration:
              AdjustmentType: CHANGE_IN_CAPACITY
              ScalingAdjustment: 1
              CoolDown: 300
          Trigger:
            CloudWatchAlarmDefinition:
              Dimensions:
                - Key: JobFlowId
                  Value: '${emr.clusterId}'
              EvaluationPeriods: 1
              Namespace: AWS/ElasticMapReduce
              Period: 300
              ComparisonOperator: LESS_THAN
              Statistic: AVERAGE
              Threshold: 15
              Unit: PERCENT
              MetricName: YARNMemoryAvailablePercentage
        - Name: Scale-in
          Description: Scale-in policy
          Action:
            SimpleScalingPolicyConfiguration:
              AdjustmentType: CHANGE_IN_CAPACITY
              ScalingAdjustment: -1
              CoolDown: 300
          Trigger:
            CloudWatchAlarmDefinition:
              Dimensions:
                - Key: JobFlowId
                  Value: '${emr.clusterId}'
              EvaluationPeriods: 1
              Namespace: AWS/ElasticMapReduce
              Period: 300
              ComparisonOperator: GREATER_THAN
              Statistic: AVERAGE
              Threshold: 75
              Unit: PERCENT
              MetricName: YARNMemoryAvailablePercentage
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VolumeSpecification

AutoScalingPolicy
