This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster

The `AWS::EMR::Cluster` resource specifies an Amazon EMR cluster. This cluster is a collection of Amazon EC2 instances that run open source big data frameworks and applications to process and analyze vast amounts of data. For more information, see the [Amazon EMR Management Guide](https://docs.aws.amazon.com/emr/latest/ManagementGuide).

Amazon EMR now supports launching task instance groups and task instance
fleets as part of the `AWS::EMR::Cluster` resource. This can be done by using
the `JobFlowInstancesConfig` property type's `TaskInstanceGroups` and
`TaskInstanceFleets` subproperties. Using these subproperties reduces delays
in provisioning task nodes compared to specifying task nodes with the
`AWS::EMR::InstanceGroupConfig` and
`AWS::EMR::InstanceFleetConfig` resources. Please refer to the examples at
the bottom of this page to learn how to use these subproperties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EMR::Cluster",
  "Properties" : {
      "AdditionalInfo" : Json,
      "Applications" : [ Application, ... ],
      "AutoScalingRole" : String,
      "AutoTerminationPolicy" : AutoTerminationPolicy,
      "BootstrapActions" : [ BootstrapActionConfig, ... ],
      "Configurations" : [ Configuration, ... ],
      "CustomAmiId" : String,
      "EbsRootVolumeIops" : Integer,
      "EbsRootVolumeSize" : Integer,
      "EbsRootVolumeThroughput" : Integer,
      "Instances" : JobFlowInstancesConfig,
      "JobFlowRole" : String,
      "KerberosAttributes" : KerberosAttributes,
      "LogEncryptionKmsKeyId" : String,
      "LogUri" : String,
      "ManagedScalingPolicy" : ManagedScalingPolicy,
      "Name" : String,
      "OSReleaseLabel" : String,
      "PlacementGroupConfigs" : [ PlacementGroupConfig, ... ],
      "ReleaseLabel" : String,
      "ScaleDownBehavior" : String,
      "SecurityConfiguration" : String,
      "ServiceRole" : String,
      "StepConcurrencyLevel" : Integer,
      "Steps" : [ StepConfig, ... ],
      "Tags" : [ Tag, ... ],
      "VisibleToAllUsers" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::EMR::Cluster
Properties:
  AdditionalInfo: Json
  Applications:
    - Application
  AutoScalingRole: String
  AutoTerminationPolicy:
    AutoTerminationPolicy
  BootstrapActions:
    - BootstrapActionConfig
  Configurations:
    - Configuration
  CustomAmiId: String
  EbsRootVolumeIops: Integer
  EbsRootVolumeSize: Integer
  EbsRootVolumeThroughput: Integer
  Instances:
    JobFlowInstancesConfig
  JobFlowRole: String
  KerberosAttributes:
    KerberosAttributes
  LogEncryptionKmsKeyId: String
  LogUri: String
  ManagedScalingPolicy:
    ManagedScalingPolicy
  Name: String
  OSReleaseLabel: String
  PlacementGroupConfigs:
    - PlacementGroupConfig
  ReleaseLabel: String
  ScaleDownBehavior: String
  SecurityConfiguration: String
  ServiceRole: String
  StepConcurrencyLevel: Integer
  Steps:
    - StepConfig
  Tags:
    - Tag
  VisibleToAllUsers: Boolean

```

## Properties

`AdditionalInfo`

A JSON string for selecting additional features.

_Required_: No

_Type_: Json

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `10280`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Applications`

The applications to install on this cluster, for example, Spark, Flink, Oozie, Zeppelin, and so on.

_Required_: No

_Type_: Array of [Application](aws-properties-emr-cluster-application.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AutoScalingRole`

An IAM role for automatic scaling policies. The default role is
`EMR_AutoScaling_DefaultRole`. The IAM role provides
permissions that the automatic scaling feature requires to launch and terminate Amazon EC2 instances in an instance group.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `10280`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AutoTerminationPolicy`

An auto-termination policy for an Amazon EMR cluster. An auto-termination policy
defines the amount of idle time in seconds after which a cluster automatically terminates.
For alternative cluster termination options, see [Control cluster\
termination](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-termination.html).

_Required_: No

_Type_: [AutoTerminationPolicy](aws-properties-emr-cluster-autoterminationpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BootstrapActions`

A list of bootstrap actions to run before Hadoop starts on the cluster nodes.

_Required_: No

_Type_: Array of [BootstrapActionConfig](aws-properties-emr-cluster-bootstrapactionconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configurations`

Applies only to Amazon EMR releases 4.x and later. The list of configurations
that are supplied to the Amazon EMR cluster.

_Required_: No

_Type_: Array of [Configuration](aws-properties-emr-cluster-configuration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomAmiId`

Available only in Amazon EMR releases 5.7.0 and later. The ID of a custom Amazon
EBS-backed Linux AMI if the cluster uses a custom AMI.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EbsRootVolumeIops`

The IOPS, of the Amazon EBS root device volume of the Linux AMI that is
used for each Amazon EC2 instance. Available in Amazon EMR releases 6.15.0 and
later.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EbsRootVolumeSize`

The size, in GiB, of the Amazon EBS root device volume of the Linux AMI that is
used for each Amazon EC2 instance. Available in Amazon EMR releases 4.x and
later.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EbsRootVolumeThroughput`

The throughput, in MiB/s, of the Amazon EBS root device volume of the Linux AMI that is
used for each Amazon EC2 instance. Available in Amazon EMR releases 6.15.0 and
later.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Instances`

A specification of the number and type of Amazon EC2 instances.

_Required_: Yes

_Type_: [JobFlowInstancesConfig](aws-properties-emr-cluster-jobflowinstancesconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobFlowRole`

Also called instance profile and Amazon EC2 role. An IAM role for
an Amazon EMR cluster. The Amazon EC2 instances of the cluster assume this
role. The default role is `EMR_EC2_DefaultRole`. In order to use the default
role, you must have already created it using the AWS CLI or console.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `10280`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KerberosAttributes`

Attributes for Kerberos configuration when Kerberos authentication is enabled using a
security configuration. For more information see [Use Kerberos Authentication](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-kerberos.html)
in the _Amazon EMR Management Guide_.

_Required_: No

_Type_: [KerberosAttributes](aws-properties-emr-cluster-kerberosattributes.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogEncryptionKmsKeyId`

The AWS KMS key used for encrypting log files. This attribute is only
available with Amazon EMR 5.30.0 and later, excluding Amazon EMR 6.0.0.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogUri`

The path to the Amazon S3 location where logs for this cluster are
stored.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManagedScalingPolicy`

Creates or updates a managed scaling policy for an Amazon EMR cluster. The
managed scaling policy defines the limits for resources, such as Amazon EC2
instances that can be added or terminated from a cluster. The policy only applies to the
core and task nodes. The master node cannot be scaled after initial configuration.

_Required_: No

_Type_: [ManagedScalingPolicy](aws-properties-emr-cluster-managedscalingpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the cluster. This parameter can't contain the characters <, >, $, \|, or \` (backtick).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OSReleaseLabel`

The Amazon Linux release specified in a cluster launch RunJobFlow request. If no Amazon
Linux release was specified, the default Amazon Linux release is shown in the
response.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PlacementGroupConfigs`

Property description not available.

_Required_: No

_Type_: Array of [PlacementGroupConfig](aws-properties-emr-cluster-placementgroupconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReleaseLabel`

The Amazon EMR release label, which determines the version of open-source
application packages installed on the cluster. Release labels are in the form
`emr-x.x.x`, where x.x.x is an Amazon EMR release version such as
`emr-5.14.0`. For more information about Amazon EMR release versions
and included application versions and features, see [https://docs.aws.amazon.com/emr/latest/ReleaseGuide/](https://docs.aws.amazon.com/emr/latest/ReleaseGuide). The release label applies only to Amazon EMR
releases version 4.0 and later. Earlier versions use `AmiVersion`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScaleDownBehavior`

The way that individual Amazon EC2 instances terminate when an automatic
scale-in activity occurs or an instance group is resized.
`TERMINATE_AT_INSTANCE_HOUR` indicates that Amazon EMR terminates
nodes at the instance-hour boundary, regardless of when the request to terminate the
instance was submitted. This option is only available with Amazon EMR 5.1.0 and
later and is the default for clusters created using that version.
`TERMINATE_AT_TASK_COMPLETION` indicates that Amazon EMR adds nodes
to a deny list and drains tasks from nodes before terminating the Amazon EC2
instances, regardless of the instance-hour boundary. With either behavior, Amazon EMR removes the least active nodes first and blocks instance termination if it could lead to
HDFS corruption. `TERMINATE_AT_TASK_COMPLETION` is available only in Amazon EMR releases 4.1.0 and later, and is the default for versions of Amazon EMR earlier than 5.1.0.

_Required_: No

_Type_: String

_Allowed values_: `TERMINATE_AT_INSTANCE_HOUR | TERMINATE_AT_TASK_COMPLETION`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityConfiguration`

The name of the security configuration applied to the cluster.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `10280`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceRole`

The IAM role that Amazon EMR assumes in order to access AWS resources on your behalf.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StepConcurrencyLevel`

Specifies the number of steps that can be executed concurrently. The default value is
`1`. The maximum value is `256`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Steps`

A list of steps to run.

_Required_: No

_Type_: Array of [StepConfig](aws-properties-emr-cluster-stepconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags associated with a cluster.

_Required_: No

_Type_: Array of [Tag](aws-properties-emr-cluster-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisibleToAllUsers`

Indicates whether the cluster is visible to all IAM users of the AWS account associated with the cluster. If this value is set to `true`, all IAM users of that AWS account can view and manage the cluster if they have the proper policy permissions set.
If this value is `false`, only the IAM user that created the cluster can view and manage it. This value can be changed using the SetVisibleToAllUsers action.

###### Note

When you create clusters directly through the EMR console or API, this value is set to `true` by default. However, for `AWS::EMR::Cluster` resources in CloudFormation, the default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns returns the cluster ID, such as j-1ABCD123AB1A.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The unique identifier for the cluster.

`MasterPublicDNS`

The public DNS name of the master node (instance), such as `ec2-12-123-123-123.us-west-2.compute.amazonaws.com`.

## Examples

Create cluster examples.

- [Create a cluster using a custom AMI for EC2 instances](#aws-resource-emr-cluster--examples--Create_a_cluster_using_a_custom_AMI_for_EC2_instances)

- [Create a cluster and specify the root volume size of EC2 instances](#aws-resource-emr-cluster--examples--Create_a_cluster_and_specify_the_root_volume_size_of_EC2_instances)

- [Create a cluster with Kerberos authentication](#aws-resource-emr-cluster--examples--Create_a_cluster_with_Kerberos_authentication)

- [Create a cluster with a managed scaling policy](#aws-resource-emr-cluster--examples--Create_a_cluster_with_a_managed_scaling_policy)

- [Create a cluster with task instance groups](#aws-resource-emr-cluster--examples--Create_a_cluster_with_task_instance_groups)

- [Create a cluster with a task instance fleet](#aws-resource-emr-cluster--examples--Create_a_cluster_with_a_task_instance_fleet)

### Create a cluster using a custom AMI for EC2 instances

The following example template specifies a cluster using a custom Amazon Linux AMI for the EC2 instances in the cluster.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Parameters" : {
    "CustomAmiId" : {
      "Type" : "String"
    },
    "InstanceType" : {
      "Type" : "String"
    },
    "ReleaseLabel" : {
      "Type" : "String"
    },
    "SubnetId" : {
      "Type" : "String"
    },
    "TerminationProtected" : {
      "Type" : "String",
      "Default" : "false"
    },
    "ElasticMapReducePrincipal" : {
      "Type" : "String"
    },
    "Ec2Principal" : {
      "Type" : "String"
    }
  },
  "Resources": {
    "cluster": {
      "Type": "AWS::EMR::Cluster",
      "Properties": {
        "CustomAmiId" : {"Ref" : "CustomAmiId"},
        "Instances": {
          "MasterInstanceGroup": {
            "InstanceCount": 1,
            "InstanceType": {"Ref" : "InstanceType"},
            "Market": "ON_DEMAND",
            "Name": "cfnMaster"
          },
          "CoreInstanceGroup": {
            "InstanceCount": 1,
            "InstanceType": {"Ref" : "InstanceType"},
            "Market": "ON_DEMAND",
            "Name": "cfnCore"
          },
          "TaskInstanceGroups": [
            {
              "InstanceCount": 1,
              "InstanceType": {"Ref" : "InstanceType"},
              "Market": "ON_DEMAND",
              "Name": "cfnTask-1"
            },
            {
              "InstanceCount": 1,
              "InstanceType": {"Ref" : "InstanceType"},
              "Market": "ON_DEMAND",
              "Name": "cfnTask-2"
            }
          ],
          "TerminationProtected" : {"Ref" : "TerminationProtected"},
          "Ec2SubnetId" : {"Ref" : "SubnetId"}
        },
        "Name": "CFNtest",
        "JobFlowRole" : {"Ref": "emrEc2InstanceProfile"},
        "ServiceRole" : {"Ref": "emrRole"},
        "ReleaseLabel" : {"Ref" : "ReleaseLabel"},
        "VisibleToAllUsers" : true,
        "Tags": [
          {
            "Key": "key1",
            "Value": "value1"
          }
        ]
      }
    },
    "emrRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": {"Ref" : "ElasticMapReducePrincipal"}
              },
              "Action": "sts:AssumeRole"
            }
          ]
        },
        "Path": "/",
        "ManagedPolicyArns": ["arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceRole"]
      }
    },
    "emrEc2Role": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": {"Ref" : "Ec2Principal"}
              },
              "Action": "sts:AssumeRole"
            }
          ]
        },
        "Path": "/",
        "ManagedPolicyArns": ["arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceforEC2Role"]
      }
    },
    "emrEc2InstanceProfile": {
      "Type": "AWS::IAM::InstanceProfile",
      "Properties": {
        "Path": "/",
        "Roles": [ {
          "Ref": "emrEc2Role"
        } ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  CustomAmiId:
    Type: String
  InstanceType:
    Type: String
  ReleaseLabel:
    Type: String
  SubnetId:
    Type: String
  TerminationProtected:
    Type: String
    Default: 'false'
  ElasticMapReducePrincipal:
    Type: String
  Ec2Principal:
    Type: String
Resources:
  cluster:
    Type: AWS::EMR::Cluster
    Properties:
      CustomAmiId: !Ref CustomAmiId
      Instances:
        MasterInstanceGroup:
          InstanceCount: 1
          InstanceType: !Ref InstanceType
          Market: ON_DEMAND
          Name: cfnMaster
        CoreInstanceGroup:
          InstanceCount: 1
          InstanceType: !Ref InstanceType
          Market: ON_DEMAND
          Name: cfnCore
        TaskInstanceGroups:
          - InstanceCount: 1
            InstanceType: !Ref InstanceType
            Market: ON_DEMAND
            Name: cfnTask-1
          - InstanceCount: 1
            InstanceType: !Ref InstanceType
            Market: ON_DEMAND
            Name: cfnTask-2
        TerminationProtected: !Ref TerminationProtected
        Ec2SubnetId: !Ref SubnetId
      Name: CFNtest
      JobFlowRole: !Ref emrEc2InstanceProfile
      ServiceRole: !Ref emrRole
      ReleaseLabel: !Ref ReleaseLabel
      VisibleToAllUsers: true
      Tags:
        - Key: key1
          Value: value1
  emrRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2008-10-17
        Statement:
          - Sid: ''
            Effect: Allow
            Principal:
              Service: !Ref ElasticMapReducePrincipal
            Action: 'sts:AssumeRole'
      Path: /
      ManagedPolicyArns:
        - 'arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceRole'
  emrEc2Role:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2008-10-17
        Statement:
          - Sid: ''
            Effect: Allow
            Principal:
              Service: !Ref Ec2Principal
            Action: 'sts:AssumeRole'
      Path: /
      ManagedPolicyArns:
        - 'arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceforEC2Role'
  emrEc2InstanceProfile:
    Type: AWS::IAM::InstanceProfile
    Properties:
      Path: /
      Roles:
        - !Ref emrEc2Role
```

### Create a cluster and specify the root volume size of EC2 instances

The following example template enables you to specify the size of the EBS root volume for cluster instances.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Parameters" : {
    "InstanceType" : {
      "Type" : "String"
    },
    "ReleaseLabel" : {
      "Type" : "String"
    },
    "SubnetId" : {
      "Type" : "String"
    },
    "TerminationProtected" : {
      "Type" : "String",
      "Default" : "false"
    },
    "EbsRootVolumeSize" : {
      "Type" : "String"
    }
  },
  "Resources": {
    "cluster": {
      "Type": "AWS::EMR::Cluster",
      "Properties": {
        "EbsRootVolumeSize" : {"Ref" : "EbsRootVolumeSize"},
        "Instances": {
          "MasterInstanceGroup": {
            "InstanceCount": 1,
            "InstanceType": {"Ref" : "InstanceType"},
            "Market": "ON_DEMAND",
            "Name": "cfnMaster"
          },
          "CoreInstanceGroup": {
            "InstanceCount": 1,
            "InstanceType": {"Ref" : "InstanceType"},
            "Market": "ON_DEMAND",
            "Name": "cfnCore"
          },
          "TaskInstanceGroups": [
            {
              "InstanceCount": 1,
              "InstanceType": {"Ref" : "InstanceType"},
              "Market": "ON_DEMAND",
              "Name": "cfnTask-1"
            },
            {
              "InstanceCount": 1,
              "InstanceType": {"Ref" : "InstanceType"},
              "Market": "ON_DEMAND",
              "Name": "cfnTask-2"
            }
          ],
          "TerminationProtected" : {"Ref" : "TerminationProtected"},
          "Ec2SubnetId" : {"Ref" : "SubnetId"}
        },
        "Name": "CFNtest",
        "JobFlowRole" : {"Ref": "emrEc2InstanceProfile"},
        "ServiceRole" : {"Ref": "emrRole"},
        "ReleaseLabel" : {"Ref" : "ReleaseLabel"},
        "VisibleToAllUsers" : true,
        "Tags": [
          {
            "Key": "key1",
            "Value": "value1"
          }
        ]
      }
    },
    "emrRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": "elasticmapreduce.amazonaws.com"
              },
              "Action": "sts:AssumeRole"
            }
          ]
        },
        "Path": "/",
        "ManagedPolicyArns": ["arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceRole"]
      }
    },
    "emrEc2Role": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": "ec2.amazonaws.com"
              },
              "Action": "sts:AssumeRole"
            }
          ]
        },
        "Path": "/",
        "ManagedPolicyArns": ["arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceforEC2Role"]
      }
    },
    "emrEc2InstanceProfile": {
      "Type": "AWS::IAM::InstanceProfile",
      "Properties": {
        "Path": "/",
        "Roles": [ {
          "Ref": "emrEc2Role"
        } ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  InstanceType:
    Type: String
  ReleaseLabel:
    Type: String
  SubnetId:
    Type: String
  TerminationProtected:
    Type: String
    Default: 'false'
  EbsRootVolumeSize:
    Type: String
Resources:
  cluster:
    Type: AWS::EMR::Cluster
    Properties:
      EbsRootVolumeSize: !Ref EbsRootVolumeSize
      Instances:
        MasterInstanceGroup:
          InstanceCount: 1
          InstanceType: !Ref InstanceType
          Market: ON_DEMAND
          Name: cfnMaster
        CoreInstanceGroup:
          InstanceCount: 1
          InstanceType: !Ref InstanceType
          Market: ON_DEMAND
          Name: cfnCore
        TaskInstanceGroups:
          - InstanceCount: 1
            InstanceType: !Ref InstanceType
            Market: ON_DEMAND
            Name: cfnTask-1
          - InstanceCount: 1
            InstanceType: !Ref InstanceType
            Market: ON_DEMAND
            Name: cfnTask-2
        TerminationProtected: !Ref TerminationProtected
        Ec2SubnetId: !Ref SubnetId
      Name: CFNtest
      JobFlowRole: !Ref emrEc2InstanceProfile
      ServiceRole: !Ref emrRole
      ReleaseLabel: !Ref ReleaseLabel
      VisibleToAllUsers: true
      Tags:
        - Key: key1
          Value: value1
  emrRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2008-10-17
        Statement:
          - Sid: ''
            Effect: Allow
            Principal:
              Service: elasticmapreduce.amazonaws.com
            Action: 'sts:AssumeRole'
      Path: /
      ManagedPolicyArns:
        - 'arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceRole'
  emrEc2Role:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2008-10-17
        Statement:
          - Sid: ''
            Effect: Allow
            Principal:
              Service: ec2.amazonaws.com
            Action: 'sts:AssumeRole'
      Path: /
      ManagedPolicyArns:
        - 'arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceforEC2Role'
  emrEc2InstanceProfile:
    Type: AWS::IAM::InstanceProfile
    Properties:
      Path: /
      Roles:
        - !Ref emrEc2Role

```

### Create a cluster with Kerberos authentication

The following example template enables you to specify the Kerberos authentication
configuration for an EMR cluster.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Parameters" : {
    "CrossRealmTrustPrincipalPassword" : {
      "Type" : "String"
    },
    "KdcAdminPassword" : {
      "Type" : "String"
    },
    "Realm" : {
      "Type" : "String"
    },
    "InstanceType" : {
      "Type" : "String"
    },
    "ReleaseLabel" : {
      "Type" : "String"
    },
    "SubnetId" : {
      "Type" : "String"
    }
  },
  "Resources": {
    "cluster": {
      "Type": "AWS::EMR::Cluster",
      "Properties": {
        "Instances": {
          "MasterInstanceGroup": {
            "InstanceCount": 1,
            "InstanceType": {"Ref" : "InstanceType"},
            "Market": "ON_DEMAND",
            "Name": "cfnMaster"
          },
          "CoreInstanceGroup": {
            "InstanceCount": 1,
            "InstanceType": {"Ref" : "InstanceType"},
            "Market": "ON_DEMAND",
            "Name": "cfnCore"
          },
          "TaskInstanceGroups": [
            {
              "InstanceCount": 1,
              "InstanceType": {"Ref" : "InstanceType"},
              "Market": "ON_DEMAND",
              "Name": "cfnTask-1"
            },
            {
              "InstanceCount": 1,
              "InstanceType": {"Ref" : "InstanceType"},
              "Market": "ON_DEMAND",
              "Name": "cfnTask-2"
            }
          ],
          "Ec2SubnetId" : {"Ref" : "SubnetId"}
        },
        "Name": "CFNtest2",
        "JobFlowRole" : {"Ref": "emrEc2InstanceProfile"},
        "KerberosAttributes" : {
          "CrossRealmTrustPrincipalPassword" : "CfnIntegrationTest-1",
          "KdcAdminPassword" : "CfnIntegrationTest-1",
          "Realm": "EC2.INTERNAL"
        },
        "ServiceRole" : {"Ref": "emrRole"},
        "ReleaseLabel" : {"Ref" : "ReleaseLabel"},
        "SecurityConfiguration" : {"Ref" : "securityConfiguration"},
        "VisibleToAllUsers" : true,
        "Tags": [
          {
            "Key": "key1",
            "Value": "value1"
          }
        ]
      }
    },
    "key" : {
      "Type" : "AWS::KMS::Key",
      "Properties" : {
        "KeyPolicy" : {
          "Version": "2012-10-17",
          "Id": "key-default-1",
          "Statement": [
            {
              "Sid": "Enable IAM User Permissions",
              "Effect": "Allow",
              "Principal": {
                "AWS": { "Fn::GetAtt" : ["emrEc2Role", "Arn"]}
              },
              "Action": "kms:*",
              "Resource": "*"
            },
            {
              "Sid": "Enable IAM User Permissions",
              "Effect": "Allow",
              "Principal": {
                "AWS": { "Fn::Join" : ["" , ["arn:aws:iam::", {"Ref" : "AWS::AccountId"} ,":root" ]] }
              },
              "Action": "kms:*",
              "Resource": "*"
            }
          ]
        }
      }
    },
    "securityConfiguration": {
      "Type" : "AWS::EMR::SecurityConfiguration",
      "Properties" : {
        "SecurityConfiguration" : {
          "AuthenticationConfiguration": {
            "KerberosConfiguration": {
              "Provider": "ClusterDedicatedKdc",
              "ClusterDedicatedKdcConfiguration": {
                "TicketLifetimeInHours": 24,
                "CrossRealmTrustConfiguration": {
                  "Realm": "AD.DOMAIN.COM",
                  "Domain": "ad.domain.com",
                  "AdminServer": "ad.domain.com",
                  "KdcServer": "ad.domain.com"
                }
              }
            }
          }
        }
      }
    },
    "emrRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": "elasticmapreduce.amazonaws.com"
              },
              "Action": "sts:AssumeRole"
            }
          ]
        },
        "Path": "/",
        "ManagedPolicyArns": ["arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceRole"]
      }
    },
    "emrEc2Role": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": "ec2.amazonaws.com"
              },
              "Action": "sts:AssumeRole"
            }
          ]
        },
        "Path": "/",
        "ManagedPolicyArns": ["arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceforEC2Role"]
      }
    },
    "emrEc2InstanceProfile": {
      "Type": "AWS::IAM::InstanceProfile",
      "Properties": {
        "Path": "/",
        "Roles": [ {
          "Ref": "emrEc2Role"
        } ]
      }
    }
  },
  "Outputs" : {
    "keyArn" : {
      "Value" : {"Fn::GetAtt" : ["key", "Arn"]}
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  CrossRealmTrustPrincipalPassword:
    Type: String
  KdcAdminPassword:
    Type: String
  Realm:
    Type: String
  InstanceType:
    Type: String
  ReleaseLabel:
    Type: String
  SubnetId:
    Type: String
Resources:
  cluster:
    Type: 'AWS::EMR::Cluster'
    Properties:
      Instances:
        MasterInstanceGroup:
          InstanceCount: 1
          InstanceType: !Ref InstanceType
          Market: ON_DEMAND
          Name: cfnMaster
        CoreInstanceGroup:
          InstanceCount: 1
          InstanceType: !Ref InstanceType
          Market: ON_DEMAND
          Name: cfnCore
        TaskInstanceGroups:
          - InstanceCount: 1
            InstanceType: !Ref InstanceType
            Market: ON_DEMAND
            Name: cfnTask-1
          - InstanceCount: 1
            InstanceType: !Ref InstanceType
            Market: ON_DEMAND
            Name: cfnTask-2
        Ec2SubnetId: !Ref SubnetId
      Name: CFNtest2
      JobFlowRole: !Ref emrEc2InstanceProfile
      KerberosAttributes:
        CrossRealmTrustPrincipalPassword: CfnIntegrationTest-1
        KdcAdminPassword: CfnIntegrationTest-1
        Realm: EC2.INTERNAL
      ServiceRole: !Ref emrRole
      ReleaseLabel: !Ref ReleaseLabel
      SecurityConfiguration: !Ref securityConfiguration
      VisibleToAllUsers: true
      Tags:
        - Key: key1
          Value: value1
  key:
    Type: 'AWS::KMS::Key'
    Properties:
      KeyPolicy:
        Version: 2012-10-17
        Id: key-default-1
        Statement:
          - Sid: Enable IAM User Permissions
            Effect: Allow
            Principal:
              AWS: !GetAtt
                - emrEc2Role
                - Arn
            Action: 'kms:*'
            Resource: '*'
          - Sid: Enable IAM User Permissions
            Effect: Allow
            Principal:
              AWS: !Join
                - ''
                - - 'arn:aws:iam::'
                  - !Ref 'AWS::AccountId'
                  - ':root'
            Action: 'kms:*'
            Resource: '*'
  securityConfiguration:
    Type: 'AWS::EMR::SecurityConfiguration'
    Properties:
      SecurityConfiguration:
        AuthenticationConfiguration:
          KerberosConfiguration:
            Provider: ClusterDedicatedKdc
            ClusterDedicatedKdcConfiguration:
              TicketLifetimeInHours: 24
              CrossRealmTrustConfiguration:
                Realm: AD.DOMAIN.COM
                Domain: ad.domain.com
                AdminServer: ad.domain.com
                KdcServer: ad.domain.com
  emrRole:
    Type: 'AWS::IAM::Role'
    Properties:
      AssumeRolePolicyDocument:
        Version: 2008-10-17
        Statement:
          - Sid: ''
            Effect: Allow
            Principal:
              Service: elasticmapreduce.amazonaws.com
            Action: 'sts:AssumeRole'
      Path: /
      ManagedPolicyArns:
        - 'arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceRole'
  emrEc2Role:
    Type: 'AWS::IAM::Role'
    Properties:
      AssumeRolePolicyDocument:
        Version: 2008-10-17
        Statement:
          - Sid: ''
            Effect: Allow
            Principal:
              Service: ec2.amazonaws.com
            Action: 'sts:AssumeRole'
      Path: /
      ManagedPolicyArns:
        - 'arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceforEC2Role'
  emrEc2InstanceProfile:
    Type: 'AWS::IAM::InstanceProfile'
    Properties:
      Path: /
      Roles:
        - !Ref emrEc2Role
Outputs:
  keyArn:
    Value: !GetAtt
      - key
      - Arn

```

### Create a cluster with a managed scaling policy

The following example template enables you to specify the managed scaling policy
for an EMR cluster.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Parameters" : {
    "InstanceType" : {
      "Type" : "String"
    },
    "ReleaseLabel" : {
      "Type" : "String"
    },
    "SubnetId" : {
      "Type" : "String"
    },
    "MinimumCapacityUnits" : {
      "Type" : "String"
    },
    "MaximumCapacityUnits" : {
      "Type" : "String"
    },
    "MaximumCoreCapacityUnits" : {
      "Type" : "String"
    },
    "MaximumOnDemandCapacityUnits" : {
      "Type" : "String"
    },
    "UnitType" : {
      "Type" : "String"
    }
  },
  "Resources": {
    "cluster": {
      "Type": "AWS::EMR::Cluster",
      "Properties": {
        "Instances": {
          "MasterInstanceGroup": {
            "InstanceCount": 1,
            "InstanceType": {"Ref" : "InstanceType"},
            "Market": "ON_DEMAND",
            "Name": "primary"
          },
          "CoreInstanceGroup": {
            "InstanceCount": 1,
            "InstanceType": {"Ref" : "InstanceType"},
            "Market": "ON_DEMAND",
            "Name": "core"
          },
          "TaskInstanceGroups": [
            {
              "InstanceCount": 1,
              "InstanceType": {"Ref" : "InstanceType"},
              "Market": "ON_DEMAND",
              "Name": "cfnTask-1"
            },
            {
              "InstanceCount": 1,
              "InstanceType": {"Ref" : "InstanceType"},
              "Market": "ON_DEMAND",
              "Name": "cfnTask-2"
            }
          ],
          "Ec2SubnetId" : {"Ref" : "SubnetId"}
        },
        "Name": "ManagedScalingExample",
        "JobFlowRole" : "EMR_EC2_DefaultRole",
        "ServiceRole" : "EMR_DefaultRole",
        "ReleaseLabel" : {"Ref" : "ReleaseLabel"},
        "ManagedScalingPolicy" : {
          "ComputeLimits" : {
            "MinimumCapacityUnits" : {"Ref": "MinimumCapacityUnits"},
            "MaximumCapacityUnits" : {"Ref": "MaximumCapacityUnits"},
            "MaximumCoreCapacityUnits" : {"Ref": "MaximumCoreCapacityUnits"},
            "MaximumOnDemandCapacityUnits" : {"Ref": "MaximumOnDemandCapacityUnits"},
            "UnitType" : {"Ref": "UnitType"}
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  InstanceType:
    Type: String
  ReleaseLabel:
    Type: String
  SubnetId:
    Type: String
  MinimumCapacityUnits:
    Type: String
  MaximumCapacityUnits:
    Type: String
  MaximumCoreCapacityUnits:
    Type: String
  MaximumOnDemandCapacityUnits:
    Type: String
  UnitType:
    Type: String
Resources:
  cluster:
    Type: 'AWS::EMR::Cluster'
    Properties:
      Instances:
        MasterInstanceGroup:
          InstanceCount: 1
          InstanceType: !Ref InstanceType
          Market: ON_DEMAND
          Name: primary
        CoreInstanceGroup:
          InstanceCount: 1
          InstanceType: !Ref InstanceType
          Market: ON_DEMAND
          Name: core
        TaskInstanceGroups:
          - InstanceCount: 1
            InstanceType: !Ref InstanceType
            Market: ON_DEMAND
            Name: cfnTask-1
          - InstanceCount: 1
            InstanceType: !Ref InstanceType
            Market: ON_DEMAND
            Name: cfnTask-2
        Ec2SubnetId: !Ref SubnetId
      Name: ManagedScalingExample
      JobFlowRole: EMR_EC2_DefaultRole
      ServiceRole: EMR_DefaultRole
      ReleaseLabel: !Ref ReleaseLabel
      ManagedScalingPolicy:
        ComputeLimits:
          MinimumCapacityUnits: !Ref MinimumCapacityUnits
          MaximumCapacityUnits: !Ref MaximumCapacityUnits
          MaximumCoreCapacityUnits: !Ref MaximumCoreCapacityUnits
          MaximumOnDemandCapacityUnits: !Ref MaximumOnDemandCapacityUnits
          UnitType: !Ref UnitType
```

### Create a cluster with task instance groups

The following example template enables you to create task instance groups
for an EMR cluster.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Parameters" : {
    "InstanceType" : {
      "Type" : "String"
    },
    "ReleaseLabel" : {
      "Type" : "String"
    },
    "SubnetId" : {
      "Type" : "String"
    },
    "TerminationProtected" : {
      "Type" : "String",
      "Default" : "false"
    },
    "ElasticMapReducePrincipal" : {
      "Type" : "String"
    },
    "Ec2Principal" : {
      "Type" : "String"
    }
  },
  "Resources": {
    "cluster": {
      "Type": "AWS::EMR::Cluster",
      "Properties": {
        "Instances": {
          "MasterInstanceGroup": {
            "InstanceCount": 1,
            "InstanceType": {"Ref" : "InstanceType"},
            "Market": "ON_DEMAND",
            "Name": "cfnMaster"
          },
          "CoreInstanceGroup": {
            "InstanceCount": 1,
            "InstanceType": {"Ref" : "InstanceType"},
            "Market": "ON_DEMAND",
            "Name": "cfnCore"
          },
          "TaskInstanceGroups": [
            {
              "InstanceCount": 1,
              "InstanceType": {"Ref" : "InstanceType"},
              "Market": "ON_DEMAND",
              "Name": "cfnTask-1"
            },
            {
              "InstanceCount": 1,
              "InstanceType": {"Ref" : "InstanceType"},
              "Market": "ON_DEMAND",
              "Name": "cfnTask-2"
            }
          ],
          "TerminationProtected" : {"Ref" : "TerminationProtected"},
          "Ec2SubnetId" : {"Ref" : "SubnetId"}
        },
        "Name": "CFNtest",
        "JobFlowRole" : {"Ref": "emrEc2InstanceProfile"},
        "ServiceRole" : {"Ref": "emrRole"},
        "ReleaseLabel" : {"Ref" : "ReleaseLabel"},

        "VisibleToAllUsers" : true,
        "Tags": [
          {
            "Key": "key1",
            "Value": "value1"
          }
        ]
      }
    },
    "emrRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": {"Ref" : "ElasticMapReducePrincipal"}
              },
              "Action": "sts:AssumeRole"
            }
          ]
        },
        "Path": "/",
        "ManagedPolicyArns": ["arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceRole"]
      }
    },
    "emrEc2Role": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": {"Ref" : "Ec2Principal"}
              },
              "Action": "sts:AssumeRole"
            }
          ]
        },
        "Path": "/",
        "ManagedPolicyArns": ["arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceforEC2Role"]
      }
    },
    "emrEc2InstanceProfile": {
      "Type": "AWS::IAM::InstanceProfile",
      "Properties": {
        "Path": "/",
        "Roles": [ {
          "Ref": "emrEc2Role"
        } ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  InstanceType:
    Type: String
  ReleaseLabel:
    Type: String
  SubnetId:
    Type: String
  TerminationProtected:
    Type: String
    Default: 'false'
  ElasticMapReducePrincipal:
    Type: String
  Ec2Principal:
    Type: String
Resources:
  cluster:
    Type: AWS::EMR::Cluster
    Properties:
      Instances:
        MasterInstanceGroup:
          InstanceCount: 1
          InstanceType: !Ref InstanceType
          Market: ON_DEMAND
          Name: cfnMaster
        CoreInstanceGroup:
          InstanceCount: 1
          InstanceType: !Ref InstanceType
          Market: ON_DEMAND
          Name: cfnCore
        TaskInstanceGroups:
          - InstanceCount: 1
            InstanceType: !Ref InstanceType
            Market: ON_DEMAND
            Name: cfnTask-1
          - InstanceCount: 1
            InstanceType: !Ref InstanceType
            Market: ON_DEMAND
            Name: cfnTask-2
        TerminationProtected: !Ref TerminationProtected
        Ec2SubnetId: !Ref SubnetId
      Name: CFNtest
      JobFlowRole: !Ref emrEc2InstanceProfile
      ServiceRole: !Ref emrRole
      ReleaseLabel: !Ref ReleaseLabel
      VisibleToAllUsers: true
      Tags:
        - Key: key1
          Value: value1
  emrRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2008-10-17
        Statement:
          - Sid: ''
            Effect: Allow
            Principal:
              Service: !Ref ElasticMapReducePrincipal
            Action: 'sts:AssumeRole'
      Path: /
      ManagedPolicyArns:
        - 'arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceRole'
  emrEc2Role:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2008-10-17
        Statement:
          - Sid: ''
            Effect: Allow
            Principal:
              Service: !Ref Ec2Principal
            Action: 'sts:AssumeRole'
      Path: /
      ManagedPolicyArns:
        - 'arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceforEC2Role'
  emrEc2InstanceProfile:
    Type: AWS::IAM::InstanceProfile
    Properties:
      Path: /
      Roles:
        - !Ref emrEc2Role
```

### Create a cluster with a task instance fleet

The following example template enables you to create a task instance fleet
for an EMR cluster.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Parameters" : {
    "InstanceType" : {
      "Type" : "String"
    },
    "ReleaseLabel" : {
      "Type" : "String"
    },
    "SubnetId" : {
      "Type" : "String"
    },
    "TerminationProtected" : {
      "Type" : "String",
      "Default" : "false"
    },
    "ElasticMapReducePrincipal" : {
      "Type" : "String"
    },
    "Ec2Principal" : {
      "Type" : "String"
    }
  },
  "Resources": {
    "cluster": {
      "Type": "AWS::EMR::Cluster",
      "Properties": {
        "Instances": {
          "MasterInstanceFleet": {
            "Name": "cfnMaster",
            "TargetOnDemandCapacity": 1,
            "TargetSpotCapacity": 0,
            "InstanceTypeConfigs": [
              {
                "InstanceType": {"Ref" : "InstanceType"},
                "WeightedCapacity": 1
              }
            ]
          },
          "CoreInstanceFleet": {
            "Name": "cfnCore",
            "TargetOnDemandCapacity": 1,
            "TargetSpotCapacity": 0,
            "InstanceTypeConfigs": [
              {
                "InstanceType": {"Ref" : "InstanceType"},
                "WeightedCapacity": 1
              }
            ]
          },
          "TaskInstanceFleets": [
            {
              "Name": "cfnTask",
              "TargetOnDemandCapacity": 1,
              "TargetSpotCapacity": 0,
              "InstanceTypeConfigs": [
                {
                  "InstanceType": {"Ref" : "InstanceType"},
                  "WeightedCapacity": 1
                }
              ]
            }
          ],
          "TerminationProtected" : {"Ref" : "TerminationProtected"},
          "Ec2SubnetIds" : [
            {"Ref" : "SubnetId"}
          ]
        },
        "Name": "CFNtest",
        "JobFlowRole" : {"Ref": "emrEc2InstanceProfile"},
        "ServiceRole" : {"Ref": "emrRole"},
        "ReleaseLabel" : {"Ref" : "ReleaseLabel"},
        "VisibleToAllUsers" : true,
        "Tags": [
          {
            "Key": "key1",
            "Value": "value1"
          }
        ]
      }
    },
    "emrRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": {"Ref" : "ElasticMapReducePrincipal"}
              },
              "Action": "sts:AssumeRole"
            }
          ]
        },
        "Path": "/",
        "ManagedPolicyArns": ["arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceRole"]
      }
    },
    "emrEc2Role": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",

              "Effect": "Allow",
              "Principal": {
                "Service": {"Ref" : "Ec2Principal"}
              },
              "Action": "sts:AssumeRole"
            }
          ]
        },
        "Path": "/",
        "ManagedPolicyArns": ["arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceforEC2Role"]
      }
    },
    "emrEc2InstanceProfile": {
      "Type": "AWS::IAM::InstanceProfile",
      "Properties": {
        "Path": "/",
        "Roles": [ {
          "Ref": "emrEc2Role"
        } ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  InstanceType:
    Type: String
  ReleaseLabel:
    Type: String
  SubnetId:
    Type: String
  TerminationProtected:
    Type: String
    Default: 'false'
  ElasticMapReducePrincipal:
    Type: String
  Ec2Principal:
    Type: String
Resources:
  cluster:
    Type: AWS::EMR::Cluster
    Properties:
      Instances:
        MasterInstanceFleet:
          Name: cfnMaster
          TargetOnDemandCapacity: 1
          TargetSpotCapacity: 0
          InstanceTypeConfigs:
            - InstanceType: !Ref InstanceType
              WeightedCapacity: 1
        CoreInstanceFleet:
          Name: cfnCore
          TargetOnDemandCapacity: 1
          TargetSpotCapacity: 0
          InstanceTypeConfigs:
            - InstanceType: !Ref InstanceType
              WeightedCapacity: 1
        TaskInstanceFleets:
          - Name: cfnTask
            TargetOnDemandCapacity: 1
            TargetSpotCapacity: 0
            InstanceTypeConfigs:
              - InstanceType: !Ref InstanceType
                WeightedCapacity: 1
        TerminationProtected: !Ref TerminationProtected
        Ec2SubnetIds:
          - !Ref SubnetId
      Name: CFNtest
      JobFlowRole: !Ref emrEc2InstanceProfile
      ServiceRole: !Ref emrRole
      ReleaseLabel: !Ref ReleaseLabel
      VisibleToAllUsers: true
      Tags:
        - Key: key1
          Value: value1
  emrRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2008-10-17
        Statement:
          - Sid: ''
            Effect: Allow
            Principal:
              Service: !Ref ElasticMapReducePrincipal
            Action: 'sts:AssumeRole'
      Path: /
      ManagedPolicyArns:
        - 'arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceRole'
  emrEc2Role:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2008-10-17
        Statement:
          - Sid: ''
            Effect: Allow
            Principal:
              Service: !Ref Ec2Principal
            Action: 'sts:AssumeRole'
      Path: /
      ManagedPolicyArns:
        - 'arn:aws:iam::aws:policy/service-role/AmazonElasticMapReduceforEC2Role'
  emrEc2InstanceProfile:
    Type: AWS::IAM::InstanceProfile
    Properties:
      Path: /
      Roles:
        - !Ref emrEc2Role
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon EMR

Application
