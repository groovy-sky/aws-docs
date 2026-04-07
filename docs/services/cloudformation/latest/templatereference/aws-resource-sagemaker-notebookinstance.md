This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::NotebookInstance

The `AWS::SageMaker::NotebookInstance` resource creates an Amazon SageMaker notebook instance. A
notebook instance is a machine learning (ML) compute instance running on a Jupyter notebook. For more information,
see [Use Notebook Instances](https://docs.aws.amazon.com/sagemaker/latest/dg/nbi.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::NotebookInstance",
  "Properties" : {
      "AcceleratorTypes" : [ String, ... ],
      "AdditionalCodeRepositories" : [ String, ... ],
      "DefaultCodeRepository" : String,
      "DirectInternetAccess" : String,
      "InstanceMetadataServiceConfiguration" : InstanceMetadataServiceConfiguration,
      "InstanceType" : String,
      "KmsKeyId" : String,
      "LifecycleConfigName" : String,
      "NotebookInstanceName" : String,
      "PlatformIdentifier" : String,
      "RoleArn" : String,
      "RootAccess" : String,
      "SecurityGroupIds" : [ String, ... ],
      "SubnetId" : String,
      "Tags" : [ Tag, ... ],
      "VolumeSizeInGB" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::NotebookInstance
Properties:
  AcceleratorTypes:
    - String
  AdditionalCodeRepositories:
    - String
  DefaultCodeRepository: String
  DirectInternetAccess: String
  InstanceMetadataServiceConfiguration:
    InstanceMetadataServiceConfiguration
  InstanceType: String
  KmsKeyId: String
  LifecycleConfigName: String
  NotebookInstanceName: String
  PlatformIdentifier: String
  RoleArn: String
  RootAccess: String
  SecurityGroupIds:
    - String
  SubnetId: String
  Tags:
    - Tag
  VolumeSizeInGB: Integer

```

## Properties

`AcceleratorTypes`

A list of Amazon Elastic Inference (EI) instance types to associate with the notebook instance. Currently, only
one instance type can be associated with a notebook instance. For more information, see [Using Elastic Inference in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html).

_Valid Values:_ `ml.eia1.medium | ml.eia1.large | ml.eia1.xlarge | ml.eia2.medium |
    ml.eia2.large | ml.eia2.xlarge`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdditionalCodeRepositories`

An array of up to three Git repositories associated with the notebook instance. These
can be either the names of Git repositories stored as resources in your account, or the
URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html)
or in any other Git repository. These repositories are cloned at the same level as the
default repository of your notebook instance. For more information, see [Associating Git\
Repositories with SageMaker AI Notebook Instances](https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultCodeRepository`

The Git repository associated with the notebook instance as its default code
repository. This can be either the name of a Git repository stored as a resource in your
account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html)
or in any other Git repository. When you open a notebook instance, it opens in the
directory that contains this repository. For more information, see [Associating Git\
Repositories with SageMaker AI Notebook Instances](https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).

_Required_: No

_Type_: String

_Pattern_: `https://([^/]+)/?(.*)$|^[a-zA-Z0-9](-*[a-zA-Z0-9])*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DirectInternetAccess`

Sets whether SageMaker AI provides internet access to the notebook instance. If
you set this to `Disabled` this notebook instance is able to access resources
only in your VPC, and is not be able to connect to SageMaker AI training and
endpoint services unless you configure a NAT Gateway in your VPC.

For more information, see [Notebook Instances Are Internet-Enabled by Default](https://docs.aws.amazon.com/sagemaker/latest/dg/appendix-additional-considerations.html#appendix-notebook-and-internet-access). You can set the value
of this parameter to `Disabled` only if you set a value for the
`SubnetId` parameter.

_Required_: No

_Type_: String

_Allowed values_: `Enabled | Disabled`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceMetadataServiceConfiguration`

Information on the IMDS configuration of the notebook instance

_Required_: No

_Type_: [InstanceMetadataServiceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-notebookinstance-instancemetadataserviceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

The type of ML compute instance to launch for the notebook instance.

###### Note

Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance
and starts it up again to update it.

_Required_: Yes

_Type_: String

_Allowed values_: `ml.t2.medium | ml.t2.large | ml.t2.xlarge | ml.t2.2xlarge | ml.t3.medium | ml.t3.large | ml.t3.xlarge | ml.t3.2xlarge | ml.m4.xlarge | ml.m4.2xlarge | ml.m4.4xlarge | ml.m4.10xlarge | ml.m4.16xlarge | ml.m5.xlarge | ml.m5.2xlarge | ml.m5.4xlarge | ml.m5.12xlarge | ml.m5.24xlarge | ml.m5d.large | ml.m5d.xlarge | ml.m5d.2xlarge | ml.m5d.4xlarge | ml.m5d.8xlarge | ml.m5d.12xlarge | ml.m5d.16xlarge | ml.m5d.24xlarge | ml.c4.xlarge | ml.c4.2xlarge | ml.c4.4xlarge | ml.c4.8xlarge | ml.c5.xlarge | ml.c5.2xlarge | ml.c5.4xlarge | ml.c5.9xlarge | ml.c5.18xlarge | ml.c5d.xlarge | ml.c5d.2xlarge | ml.c5d.4xlarge | ml.c5d.9xlarge | ml.c5d.18xlarge | ml.p2.xlarge | ml.p2.8xlarge | ml.p2.16xlarge | ml.p3.2xlarge | ml.p3.8xlarge | ml.p3.16xlarge | ml.p3dn.24xlarge | ml.g4dn.xlarge | ml.g4dn.2xlarge | ml.g4dn.4xlarge | ml.g4dn.8xlarge | ml.g4dn.12xlarge | ml.g4dn.16xlarge | ml.r5.large | ml.r5.xlarge | ml.r5.2xlarge | ml.r5.4xlarge | ml.r5.8xlarge | ml.r5.12xlarge | ml.r5.16xlarge | ml.r5.24xlarge | ml.g5.xlarge | ml.g5.2xlarge | ml.g5.4xlarge | ml.g5.8xlarge | ml.g5.16xlarge | ml.g5.12xlarge | ml.g5.24xlarge | ml.g5.48xlarge | ml.inf1.xlarge | ml.inf1.2xlarge | ml.inf1.6xlarge | ml.inf1.24xlarge | ml.trn1.2xlarge | ml.trn1.32xlarge | ml.trn1n.32xlarge | ml.inf2.xlarge | ml.inf2.8xlarge | ml.inf2.24xlarge | ml.inf2.48xlarge | ml.p4d.24xlarge | ml.p4de.24xlarge | ml.p5.48xlarge | ml.p6-b200.48xlarge | ml.m6i.large | ml.m6i.xlarge | ml.m6i.2xlarge | ml.m6i.4xlarge | ml.m6i.8xlarge | ml.m6i.12xlarge | ml.m6i.16xlarge | ml.m6i.24xlarge | ml.m6i.32xlarge | ml.m7i.large | ml.m7i.xlarge | ml.m7i.2xlarge | ml.m7i.4xlarge | ml.m7i.8xlarge | ml.m7i.12xlarge | ml.m7i.16xlarge | ml.m7i.24xlarge | ml.m7i.48xlarge | ml.c6i.large | ml.c6i.xlarge | ml.c6i.2xlarge | ml.c6i.4xlarge | ml.c6i.8xlarge | ml.c6i.12xlarge | ml.c6i.16xlarge | ml.c6i.24xlarge | ml.c6i.32xlarge | ml.c7i.large | ml.c7i.xlarge | ml.c7i.2xlarge | ml.c7i.4xlarge | ml.c7i.8xlarge | ml.c7i.12xlarge | ml.c7i.16xlarge | ml.c7i.24xlarge | ml.c7i.48xlarge | ml.r6i.large | ml.r6i.xlarge | ml.r6i.2xlarge | ml.r6i.4xlarge | ml.r6i.8xlarge | ml.r6i.12xlarge | ml.r6i.16xlarge | ml.r6i.24xlarge | ml.r6i.32xlarge | ml.r7i.large | ml.r7i.xlarge | ml.r7i.2xlarge | ml.r7i.4xlarge | ml.r7i.8xlarge | ml.r7i.12xlarge | ml.r7i.16xlarge | ml.r7i.24xlarge | ml.r7i.48xlarge | ml.m6id.large | ml.m6id.xlarge | ml.m6id.2xlarge | ml.m6id.4xlarge | ml.m6id.8xlarge | ml.m6id.12xlarge | ml.m6id.16xlarge | ml.m6id.24xlarge | ml.m6id.32xlarge | ml.c6id.large | ml.c6id.xlarge | ml.c6id.2xlarge | ml.c6id.4xlarge | ml.c6id.8xlarge | ml.c6id.12xlarge | ml.c6id.16xlarge | ml.c6id.24xlarge | ml.c6id.32xlarge | ml.r6id.large | ml.r6id.xlarge | ml.r6id.2xlarge | ml.r6id.4xlarge | ml.r6id.8xlarge | ml.r6id.12xlarge | ml.r6id.16xlarge | ml.r6id.24xlarge | ml.r6id.32xlarge | ml.g6.xlarge | ml.g6.2xlarge | ml.g6.4xlarge | ml.g6.8xlarge | ml.g6.12xlarge | ml.g6.16xlarge | ml.g6.24xlarge | ml.g6.48xlarge`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The Amazon Resource Name (ARN) of a AWS Key Management Service key that
SageMaker AI uses to encrypt data on the storage volume attached to your
notebook instance. The KMS key you provide must be enabled. For information, see [Enabling and\
Disabling Keys](https://docs.aws.amazon.com/kms/latest/developerguide/enabling-keys.html) in the _AWS Key Management Service_
_Developer Guide_.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:/_-]*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LifecycleConfigName`

The name of a lifecycle configuration to associate with the notebook instance. For information about lifecycle
configurations, see [Customize a\
Notebook Instance](https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html) in the _Amazon SageMaker Developer Guide_.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9])*`

_Minimum_: `0`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotebookInstanceName`

The name of the new notebook instance.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9])*`

_Minimum_: `0`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PlatformIdentifier`

The platform identifier of the notebook instance runtime environment. The default value is `notebook-al2-v2`.

_Required_: No

_Type_: String

_Pattern_: `(notebook-al1-v1|notebook-al2-v1|notebook-al2-v2|notebook-al2-v3|notebook-al2023-v1)`

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

When you send any requests to AWS resources from the notebook
instance, SageMaker AI assumes this role to perform tasks on your behalf. You must
grant this role necessary permissions so SageMaker AI can perform these tasks. The
policy must allow the SageMaker AI service principal (sagemaker.amazonaws.com)
permissions to assume this role. For more information, see [SageMaker AI Roles](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).

###### Note

To be able to pass this role to SageMaker AI, the caller of this API must
have the `iam:PassRole` permission.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RootAccess`

Whether root access is enabled or disabled for users of the notebook instance. The
default value is `Enabled`.

###### Note

Lifecycle configurations need root access to be able to set up a notebook
instance. Because of this, lifecycle configurations associated with a notebook
instance always run with root access even if you disable root access for
users.

_Required_: No

_Type_: String

_Allowed values_: `Enabled | Disabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

The VPC security group IDs, in the form sg-xxxxxxxx. The security groups must be
for the same VPC as specified in the subnet.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetId`

The ID of the subnet in a VPC to which you would like to have a connectivity from
your ML compute instance.

_Required_: No

_Type_: String

_Pattern_: `[-0-9a-zA-Z]+`

_Minimum_: `0`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of key-value pairs to apply to this resource.

For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost\
Allocation Tags](../../../awsaccountbilling/latest/aboutv2/cost-alloc-tags.md#allocation-what).

You can add tags later by using the `CreateTags` API.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-notebookinstance-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeSizeInGB`

The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.

###### Note

Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance
and starts it up again to update it.

_Required_: No

_Type_: Integer

_Minimum_: `5`

_Maximum_: `16384`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the notebook instance, such as
`arn:aws:sagemaker:us-west-2:012345678901:notebook-instance/mynotebookinstance`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The name of the notebook instance.

`NotebookInstanceName`

The name of the notebook instance, such as `MyNotebookInstance`.

## Examples

### SageMaker Notebook Instance Example

The following example creates a notebook instance.

#### JSON

```json

{
    "Description": "Create Basic Notebook",
    "Resources": {
        "BasicNotebookInstance": {
            "Type": "AWS::SageMaker::NotebookInstance",
            "Properties": {
                "InstanceType": "ml.t2.large",
                "RoleArn": {
                    "Fn::GetAtt": [
                        "ExecutionRole",
                        "Arn"
                    ]
                }
            }
        },
        "ExecutionRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "AssumeRolePolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": [
                                    "sagemaker.amazonaws.com"
                                ]
                            },
                            "Action": [
                                "sts:AssumeRole"
                            ]
                        }
                    ]
                },
                "Path": "/",
                "ManagedPolicyArns": [
                    {
                        "Fn::Sub": "arn:${AWS::Partition}:iam::aws:policy/AmazonSageMakerFullAccess"
                    }
                ]
            }
        }
    },
    "Outputs": {
        "BasicNotebookInstanceId": {
            "Value": {
                "Ref": "BasicNotebookInstance"
            }
        }
    }
}
```

#### YAML

```yaml

Description: "Create basic notebook instance"
Resources:
  BasicNotebookInstance:
    Type: "AWS::SageMaker::NotebookInstance"
    Properties:
      InstanceType: "ml.t2.large"
      RoleArn: !GetAtt ExecutionRole.Arn
  ExecutionRole:
    Type: "AWS::IAM::Role"
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          -
            Effect: "Allow"
            Principal:
              Service:
                - "sagemaker.amazonaws.com"
            Action:
              - "sts:AssumeRole"
      Path: "/"
      ManagedPolicyArns:
        - !Sub "arn:${AWS::Partition}:iam::aws:policy/AmazonSageMakerFullAccess"
Outputs:
  BasicNotebookInstanceId:
    Value: !Ref BasicNotebookInstance
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcConfig

InstanceMetadataServiceConfiguration
