This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticBeanstalk::Environment

The AWS::ElasticBeanstalk::Environment resource is an AWS Elastic Beanstalk resource
type that specifies an Elastic Beanstalk environment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElasticBeanstalk::Environment",
  "Properties" : {
      "ApplicationName" : String,
      "CNAMEPrefix" : String,
      "Description" : String,
      "EnvironmentName" : String,
      "OperationsRole" : String,
      "OptionSettings" : [ OptionSetting, ... ],
      "PlatformArn" : String,
      "SolutionStackName" : String,
      "Tags" : [ Tag, ... ],
      "TemplateName" : String,
      "Tier" : Tier,
      "VersionLabel" : String
    }
}

```

### YAML

```yaml

Type: AWS::ElasticBeanstalk::Environment
Properties:
  ApplicationName: String
  CNAMEPrefix: String
  Description: String
  EnvironmentName: String
  OperationsRole: String
  OptionSettings:
    - OptionSetting
  PlatformArn: String
  SolutionStackName: String
  Tags:
    - Tag
  TemplateName: String
  Tier:
    Tier
  VersionLabel: String

```

## Properties

`ApplicationName`

The name of the application that is associated with this environment.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CNAMEPrefix`

If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is
generated automatically by appending a random alphanumeric string to the environment name.

_Required_: No

_Type_: String

_Minimum_: `4`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

Your description for this environment.

_Required_: No

_Type_: String

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentName`

A unique name for the environment.

Constraint: Must be from 4 to 40 characters in length. The name can contain only
letters, numbers, and hyphens. It can't start or end with a hyphen. This name must be unique
within a region in your account.

If you don't specify the `CNAMEPrefix` parameter, the environment name becomes part of
the CNAME, and therefore part of the visible URL for your application.

If you don't specify an environment name, AWS CloudFormation generates a unique physical
ID and uses that ID for the environment name. For more information, see [Name\
Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).

###### Important

If you specify a name, you cannot perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you must replace
the resource, specify a new name.

_Required_: No

_Type_: String

_Minimum_: `4`

_Maximum_: `40`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OperationsRole`

###### Important

The operations role feature of AWS Elastic Beanstalk is in beta release and is subject to change.

The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role. If specified, Elastic Beanstalk uses the operations role
for permissions to downstream services during this call and during subsequent calls acting on this environment. To specify an operations role, you must
have the `iam:PassRole` permission for the role.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptionSettings`

Key-value pairs defining configuration options for this environment, such as the instance
type. These options override the values that are defined in the solution stack or the [configuration template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html).
If you remove any options during a stack update, the removed options retain their current values.

_Required_: No

_Type_: Array of [OptionSetting](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticbeanstalk-environment-optionsetting.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlatformArn`

The Amazon Resource Name (ARN) of the custom platform to use with the environment. For more information, see [Custom Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html) in the _AWS Elastic Beanstalk Developer Guide_.

###### Note

If you specify `PlatformArn`, don't specify `SolutionStackName`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SolutionStackName`

The name of an Elastic Beanstalk solution stack (platform version) to use with the environment. If specified, Elastic Beanstalk sets the configuration values to the default
values associated with the specified solution stack. For a list of current solution stacks, see [Elastic Beanstalk Supported Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/platforms/platforms-supported.html) in the _AWS Elastic Beanstalk_
_Platforms_ guide.

###### Note

If you specify `SolutionStackName`, don't specify `PlatformArn` or `TemplateName`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Specifies the tags applied to resources in the environment.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticbeanstalk-environment-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateName`

The name of the Elastic Beanstalk configuration template to use with the environment.

###### Note

If you specify `TemplateName`, then don't specify `SolutionStackName`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tier`

Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support
a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.

_Required_: No

_Type_: [Tier](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticbeanstalk-environment-tier.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionLabel`

The name of the application version to deploy.

Default: If not specified, Elastic Beanstalk attempts to deploy the sample application.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`EndpointURL`

For load-balanced, autoscaling environments, the URL to the load balancer.
For single-instance environments, the IP address of the instance.

Example load balancer URL:

`awseb-myst-myen-132MQC4KRLAMD-1371280482.us-east-2.elb.amazonaws.com`

Example instance IP address:

`192.0.2.0`

## Examples

- [Simple Environment](#aws-resource-elasticbeanstalk-environment--examples--Simple_Environment)

- [Environment with Embedded Option Settings](#aws-resource-elasticbeanstalk-environment--examples--Environment_with_Embedded_Option_Settings)

- [Custom or Supported Platform](#aws-resource-elasticbeanstalk-environment--examples--Custom_or_Supported_Platform)

### Simple Environment

#### JSON

```json

{
   "Type" : "AWS::ElasticBeanstalk::Environment",
   "Properties" : {
      "ApplicationName" : { "Ref" : "sampleApplication" },
      "Description" :  "AWS Elastic Beanstalk Environment running PHP Sample Application",
      "EnvironmentName" :  "SamplePHPEnvironment",
      "TemplateName" : "DefaultConfiguration",
      "VersionLabel" : "Initial Version"
   }
}
```

#### YAML

```yaml

Type: AWS::ElasticBeanstalk::Environment
Properties:
  ApplicationName:
    Ref: sampleApplication
    Description: "AWS Elastic Beanstalk Environment running PHP Sample Application"
  EnvironmentName: SamplePHPEnvironment
  TemplateName: DefaultConfiguration
  VersionLabel: "Initial Version"
```

### Environment with Embedded Option Settings

#### JSON

```json

{
   "Type" : "AWS::ElasticBeanstalk::Environment",
   "Properties" : {
      "ApplicationName" : { "Ref" : "sampleApplication" },
      "Description" :  "AWS Elastic Beanstalk Environment running Python Sample Application",
      "EnvironmentName" :  "SamplePythonEnvironment",
      "SolutionStackName" : "64bit Amazon Linux 2017.03 v2.5.0 running Python 2.7",
      "OptionSettings" : [ {
         "Namespace" : "aws:autoscaling:launchconfiguration",
         "OptionName" : "EC2KeyName",
         "Value" : { "Ref" : "KeyName" }
      } ],
      "VersionLabel" : "Initial Version"
   }
}
```

#### YAML

```yaml

Type: AWS::ElasticBeanstalk::Environment
Properties:
  ApplicationName:
    Ref: sampleApplication
    Description: "AWS Elastic Beanstalk Environment running Python Sample Application"
  EnvironmentName: SamplePythonEnvironment
  SolutionStackName: "64bit Amazon Linux 2017.03 v2.5.0 running Python 2.7"
  OptionSettings:
    -
      Namespace: "aws:autoscaling:launchconfiguration"
      OptionName: EC2KeyName
      Value:
        Ref: KeyName
  VersionLabel: "Initial Version"
```

### Custom or Supported Platform

The following example contains parameters that enable specifying
`PlatformArn` for a custom platform or `SolutionStackName` for a
supported platform when creating the stack.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "Elasticbeanstalk test template",
  "Parameters": {
    "BeanstalkService": {
      "Type": "String"
    },
    "Ec2Service": {
      "Type": "String"
    },
    "Partition":{
      "Type": "String"
    },
    "SolutionStackName": {
      "Type": "String"
    },
    "PlatformArn": {
      "Type": "String"
    }
  },
  "Resources": {
    "Application": {
      "Type": "AWS::ElasticBeanstalk::Application",
      "Properties": {
        "ApplicationName": "SampleBeanstalkApp",
        "Description": "AWS Elastic Beanstalk Python Sample Application"
      }
    },
    "AppVersion": {
      "Type": "AWS::ElasticBeanstalk::ApplicationVersion",
      "Properties": {
        "ApplicationName": {"Ref" : "Application"},
        "Description": "Version 1.0",
        "SourceBundle": {
          "S3Bucket": {
            "Fn::Join": ["-", [ "elasticbeanstalk-samples", { "Ref" : "AWS::Region" } ] ] },
          "S3Key": "python-sample-20150402.zip"
        }
      }
    },
    "Environment": {
      "Type": "AWS::ElasticBeanstalk::Environment",
      "Properties": {
        "ApplicationName": {"Ref": "Application"},
        "Description": "AWS Elastic Beanstalk Environment running Python Sample Application",
        "PlatformArn": { "Ref" : "PlatformArn"},
        "SolutionStackName": {
          "Ref": "SolutionStackName"
        },
        "VersionLabel": {"Ref": "AppVersion"},
        "OptionSettings": [
          {
            "Namespace": "aws:autoscaling:launchconfiguration",
            "OptionName": "IamInstanceProfile",
            "Value": {
              "Ref": "InstanceProfile"
            }
          },
          {
            "Namespace": "aws:elasticbeanstalk:environment",
            "OptionName": "ServiceRole",
            "Value": {
              "Ref": "ServiceRole"
            }
          }
        ]
      }
    },
    "ServiceRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": {"Ref": "BeanstalkService"}
              },
              "Action": "sts:AssumeRole",
              "Condition": {
                "StringEquals": {
                  "sts:ExternalId": "elasticbeanstalk"
                }
              }
            }
          ]
        },
        "Policies": [
          {
            "PolicyName": "root",
            "PolicyDocument": {
              "Version": "2012-10-17",
              "Statement": [
                {
                  "Effect": "Allow",
                  "Action": [
                    "elasticloadbalancing:DescribeInstanceHealth",
                    "ec2:DescribeInstances",
                    "ec2:DescribeInstanceStatus",
                    "ec2:GetConsoleOutput",
                    "ec2:AssociateAddress",
                    "ec2:DescribeAddresses",
                    "ec2:DescribeSecurityGroups",
                    "sqs:GetQueueAttributes",
                    "sqs:GetQueueUrl",
                    "autoscaling:DescribeAutoScalingGroups",
                    "autoscaling:DescribeAutoScalingInstances",
                    "autoscaling:DescribeScalingActivities",
                    "autoscaling:DescribeNotificationConfigurations"
                  ],
                  "Resource": [
                    "*"
                  ]
                }
              ]
            }
          }
        ],
        "Path": "/"
      }
    },
    "InstanceProfile": {
      "Type": "AWS::IAM::InstanceProfile",
      "Properties": {
        "Path": "/",
        "Roles": [
          {
            "Ref": "InstanceProfileRole"
          }
        ]
      }
    },
    "InstanceProfileRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Principal": {
                "Service": [
                  {"Ref": "Ec2Service"}
                ]
              },
              "Action": [
                "sts:AssumeRole"
              ]
            }
          ]
        },
        "Policies": [
          {
            "PolicyName": "root",
            "PolicyDocument": {
              "Version": "2012-10-17",
              "Statement": [
                {
                  "Sid": "BucketAccess",
                  "Action": [
                    "s3:Get*",
                    "s3:List*",
                    "s3:PutObject"
                  ],
                  "Effect": "Allow",
                  "Resource": [
                    {
                      "Fn::Join": [
                        "",
                        [
                          "arn:",
                          {
                            "Ref": "Partition"
                          },
                          ":s3:::elasticbeanstalk-*-",
                          {
                            "Ref": "AWS::AccountId"
                          }
                        ]
                      ]
                    },
                    {
                      "Fn::Join": [
                        "",
                        [
                          "arn:",
                          {
                            "Ref": "Partition"
                          },
                          ":s3:::elasticbeanstalk-*-",
                          {
                            "Ref": "AWS::AccountId"
                          },
                          "/*"
                        ]
                      ]
                    },
                    {
                      "Fn::Join": [
                        "",
                        [
                          "arn:",
                          {
                            "Ref": "Partition"
                          },
                          ":s3:::elasticbeanstalk-*-",
                          {
                            "Ref": "AWS::AccountId"
                          },
                          "-*"
                        ]
                      ]
                    },
                    {
                      "Fn::Join": [
                        "",
                        [
                          "arn:",
                          {
                            "Ref": "Partition"
                          },
                          ":s3:::elasticbeanstalk-*-",
                          {
                            "Ref": "AWS::AccountId"
                          },
                          "-*/*"
                        ]
                      ]
                    }
                  ]
                },
                {
                  "Sid": "ECSAccess",
                  "Effect": "Allow",
                  "Action": [
                    "ecs:StartTask",
                    "ecs:StopTask",
                    "ecs:RegisterContainerInstance",
                    "ecs:DeregisterContainerInstance",
                    "ecs:DescribeContainerInstances",
                    "ecs:DiscoverPollEndpoint",
                    "ecs:Submit*",
                    "ecs:Poll"
                  ],
                  "Resource": "*"
                },
                {
                  "Sid": "QueueAccess",
                  "Action": [
                    "sqs:ChangeMessageVisibility",
                    "sqs:DeleteMessage",
                    "sqs:ReceiveMessage",
                    "sqs:SendMessage"
                  ],
                  "Effect": "Allow",
                  "Resource": "*"
                },
                {
                  "Sid": "DynamoPeriodicTasks",
                  "Action": [
                    "dynamodb:BatchGetItem",
                    "dynamodb:BatchWriteItem",
                    "dynamodb:DeleteItem",
                    "dynamodb:GetItem",
                    "dynamodb:PutItem",
                    "dynamodb:Query",
                    "dynamodb:Scan",
                    "dynamodb:UpdateItem"
                  ],
                  "Effect": "Allow",
                  "Resource": [
                    {
                      "Fn::Join": [
                        "",
                        [
                          "arn:",
                          {
                            "Ref": "Partition"
                          },
                          ":dynamodb:*:",
                          {
                            "Ref": "AWS::AccountId"
                          },
                          ":table/*-stack-AWSEBWorkerCronLeaderRegistry*"
                        ]
                      ]
                    }
                  ]
                },
                {
                  "Sid": "MetricsAccess",
                  "Action": [
                    "cloudwatch:PutMetricData"
                  ],
                  "Effect": "Allow",
                  "Resource": "*"
                }
              ]
            }
          }
        ],
        "Path": "/"
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Elasticbeanstalk test template
Parameters:
  BeanstalkService:
    Type: String
  Ec2Service:
    Type: String
  Partition:
    Type: String
  SolutionStackName:
    Type: String
  PlatformArn:
    Type: String
Resources:
  Application:
    Type: AWS::ElasticBeanstalk::Application
    Properties:
      ApplicationName: SampleBeanstalkApp
      Description: AWS Elastic Beanstalk Python Sample Application
  AppVersion:
    Type: AWS::ElasticBeanstalk::ApplicationVersion
    Properties:
      ApplicationName: !Ref Application
      Description: Version 1.0
      SourceBundle:
        S3Bucket: Fn::Join:
          - '-'
          -
          - 'elasticbeanstalk-samples'
          - !Ref 'AWS::Region'
        S3Key: python-sample-20150402.zip
  Environment:
    Type: AWS::ElasticBeanstalk::Environment
    Properties:
      ApplicationName: !Ref Application
      Description: AWS Elastic Beanstalk Environment running Python Sample Application
      PlatformArn: !Ref PlatformArn
      SolutionStackName: !Ref SolutionStackName
      VersionLabel: !Ref AppVersion
      OptionSettings:
        - Namespace: 'aws:autoscaling:launchconfiguration'
          OptionName: IamInstanceProfile
          Value: !Ref InstanceProfile
        - Namespace: 'aws:elasticbeanstalk:environment'
          OptionName: ServiceRole
          Value: !Ref ServiceRole
  ServiceRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2012-10-17
        Statement:
          - Sid: ''
            Effect: Allow
            Principal:
              Service: !Ref BeanstalkService
            Action: 'sts:AssumeRole'
            Condition:
              StringEquals:
                'sts:ExternalId': elasticbeanstalk
      Policies:
        - PolicyName: root
          PolicyDocument:
            Version: 2012-10-17
            Statement:
              - Effect: Allow
                Action:
                  - 'elasticloadbalancing:DescribeInstanceHealth'
                  - 'ec2:DescribeInstances'
                  - 'ec2:DescribeInstanceStatus'
                  - 'ec2:GetConsoleOutput'
                  - 'ec2:AssociateAddress'
                  - 'ec2:DescribeAddresses'
                  - 'ec2:DescribeSecurityGroups'
                  - 'sqs:GetQueueAttributes'
                  - 'sqs:GetQueueUrl'
                  - 'autoscaling:DescribeAutoScalingGroups'
                  - 'autoscaling:DescribeAutoScalingInstances'
                  - 'autoscaling:DescribeScalingActivities'
                  - 'autoscaling:DescribeNotificationConfigurations'
                Resource:
                  - '*'
      Path: /
  InstanceProfile:
    Type: AWS::IAM::InstanceProfile
    Properties:
      Path: /
      Roles:
        - !Ref InstanceProfileRole
  InstanceProfileRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2012-10-17
        Statement:
          - Effect: Allow
            Principal:
              Service:
                - !Ref Ec2Service
            Action:
              - 'sts:AssumeRole'
      Policies:
        - PolicyName: root
          PolicyDocument:
            Version: 2012-10-17
            Statement:
              - Sid: BucketAccess
                Action:
                  - 's3:Get*'
                  - 's3:List*'
                  - 's3:PutObject'
                Effect: Allow
                Resource:
                  - !Join
                    - ''
                    - - 'arn:'
                      - !Ref Partition
                      - ':s3:::elasticbeanstalk-*-'
                      - !Ref 'AWS::AccountId'
                  - !Join
                    - ''
                    - - 'arn:'
                      - !Ref Partition
                      - ':s3:::elasticbeanstalk-*-'
                      - !Ref 'AWS::AccountId'
                      - /*
                  - !Join
                    - ''
                    - - 'arn:'
                      - !Ref Partition
                      - ':s3:::elasticbeanstalk-*-'
                      - !Ref 'AWS::AccountId'
                      - '-*'
                  - !Join
                    - ''
                    - - 'arn:'
                      - !Ref Partition
                      - ':s3:::elasticbeanstalk-*-'
                      - !Ref 'AWS::AccountId'
                      - '-*/*'
              - Sid: ECSAccess
                Effect: Allow
                Action:
                  - 'ecs:StartTask'
                  - 'ecs:StopTask'
                  - 'ecs:RegisterContainerInstance'
                  - 'ecs:DeregisterContainerInstance'
                  - 'ecs:DescribeContainerInstances'
                  - 'ecs:DiscoverPollEndpoint'
                  - 'ecs:Submit*'
                  - 'ecs:Poll'
                Resource: '*'
              - Sid: QueueAccess
                Action:
                  - 'sqs:ChangeMessageVisibility'
                  - 'sqs:DeleteMessage'
                  - 'sqs:ReceiveMessage'
                  - 'sqs:SendMessage'
                Effect: Allow
                Resource: '*'
              - Sid: DynamoPeriodicTasks
                Action:
                  - 'dynamodb:BatchGetItem'
                  - 'dynamodb:BatchWriteItem'
                  - 'dynamodb:DeleteItem'
                  - 'dynamodb:GetItem'
                  - 'dynamodb:PutItem'
                  - 'dynamodb:Query'
                  - 'dynamodb:Scan'
                  - 'dynamodb:UpdateItem'
                Effect: Allow
                Resource:
                  - !Join
                    - ''
                    - - 'arn:'
                      - !Ref Partition
                      - ':dynamodb:*:'
                      - !Ref 'AWS::AccountId'
                      - ':table/*-stack-AWSEBWorkerCronLeaderRegistry*'
              - Sid: MetricsAccess
                Action:
                  - 'cloudwatch:PutMetricData'
                Effect: Allow
                Resource: '*'
      Path: /

```

## See also

- [Creating an AWS Elastic Beanstalk Environment](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/using-features.environments.html) in the _AWS Elastic Beanstalk Developer_
_Guide_

- [Managing Environments](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/using-features.managing.html) in the _AWS Elastic Beanstalk Developer_
_Guide_

- For a complete Elastic Beanstalk sample template, see [Elastic\
Beanstalk Template Snippets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-elasticbeanstalk.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SourceConfiguration

OptionSetting
