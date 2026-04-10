This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::MonitoringSchedule

The `AWS::SageMaker::MonitoringSchedule` resource is an Amazon SageMaker resource type that regularly
starts SageMaker processing Jobs to monitor the data captured for a SageMaker endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::MonitoringSchedule",
  "Properties" : {
      "EndpointName" : String,
      "FailureReason" : String,
      "LastMonitoringExecutionSummary" : MonitoringExecutionSummary,
      "MonitoringScheduleConfig" : MonitoringScheduleConfig,
      "MonitoringScheduleName" : String,
      "MonitoringScheduleStatus" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::MonitoringSchedule
Properties:
  EndpointName: String
  FailureReason: String
  LastMonitoringExecutionSummary:
    MonitoringExecutionSummary
  MonitoringScheduleConfig:
    MonitoringScheduleConfig
  MonitoringScheduleName: String
  MonitoringScheduleStatus: String
  Tags:
    - Tag

```

## Properties

`EndpointName`

The name of the endpoint using the monitoring schedule.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureReason`

Contains the reason a monitoring job failed, if it failed.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastMonitoringExecutionSummary`

Describes metadata on the last execution to run, if there was one.

_Required_: No

_Type_: [MonitoringExecutionSummary](aws-properties-sagemaker-monitoringschedule-monitoringexecutionsummary.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringScheduleConfig`

The configuration object that specifies the monitoring schedule and defines the monitoring job.

_Required_: Yes

_Type_: [MonitoringScheduleConfig](aws-properties-sagemaker-monitoringschedule-monitoringscheduleconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringScheduleName`

The name of the monitoring schedule.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MonitoringScheduleStatus`

The status of the monitoring schedule.

_Required_: No

_Type_: String

_Allowed values_: `Pending | Failed | Scheduled | Stopped`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-monitoringschedule-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the monitoring schedule.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The time when the monitoring schedule was created.

`LastModifiedTime`

The last time that the monitoring schedule was modified.

`MonitoringScheduleArn`

The Amazon Resource Name (ARN) of the monitoring schedule.

## Examples

### SageMaker MonitoringSchedule Example

The following example creates a monitoring schedule for a SageMaker endpoint.

#### JSON

```json

{
   "AWSTemplateFormatVersion": "2010-09-09",
   "Description": "Basic SageMaker Hosting entities to create a monitoring schedule",
   "Mappings": {
      "RegionMap": {
         "us-west-2": {
            "NullTransformer": "123456789012.dkr.ecr.us-west-2.amazonaws.com/mymodel:latest"
         },
         "us-east-2": {
            "NullTransformer": "123456789012.dkr.ecr.us-east-2.amazonaws.com/mymodel:latest"
         },
         "us-east-1": {
            "NullTransformer": "123456789012.dkr.ecr.us-east-1.amazonaws.com/mymodel:latest"
         },
         "eu-west-1": {
            "NullTransformer": "123456789012.dkr.ecr.eu-west-1.amazonaws.com/mymodel:latest"
         },
         "ap-northeast-1": {
            "NullTransformer": "123456789012.dkr.ecr.ap-northeast-1.amazonaws.com/mymodel:latest"
         },
         "ap-northeast-2": {
            "NullTransformer": "123456789012.dkr.ecr.ap-northeast-2.amazonaws.com/mymodel:latest"
         },
         "ap-southeast-2": {
            "NullTransformer": "123456789012.dkr.ecr.ap-southeast-2.amazonaws.com/mymodel:latest"
         },
         "eu-central-1": {
            "NullTransformer": "123456789012.dkr.ecr.eu-central-1.amazonaws.com/mymodel:latest"
         }
      }
   },
   "Resources": {
      "Endpoint": {
         "Type": "AWS::SageMaker::Endpoint",
         "Properties": {
            "EndpointConfigName": null
         }
      },
      "EndpointConfig": {
         "Type": "AWS::SageMaker::EndpointConfig",
         "Properties": {
            "ProductionVariants": [
               {
                  "InitialInstanceCount": 1,
                  "InitialVariantWeight": 1,
                  "InstanceType": "ml.t2.large",
                  "ModelName": null,
                  "VariantName": null
               }
            ]
         }
      },
      "Model": {
         "Type": "AWS::SageMaker::Model",
         "Properties": {
            "PrimaryContainer": {
               "Image": null
            },
            "ExecutionRoleArn": null
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
            "Policies": [
               {
                  "PolicyName": "root",
                  "PolicyDocument": {
                     "Version": "2012-10-17",
                     "Statement": [
                        {
                           "Effect": "Allow",
                           "Action": "*",
                           "Resource": "*"
                        }
                     ]
                  }
               }
            ]
         }
      },
      "MonitoringScheduleExecutionRole": {
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
            "ManagedPolicyArns": [
               {
                  "Fn::Sub": "arn:${AWS::Partition}:iam::aws:policy/AmazonSageMakerFullAccess"
               },
               {
                  "Fn::Sub": "arn:${AWS::Partition}:iam::aws:policy/AmazonS3FullAccess"
               },
               {
                  "Fn::Sub": "arn:${AWS::Partition}:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
               }
            ]
         }
      },
      "MonitoringSchedule": {
         "Type": "AWS::SageMaker::MonitoringSchedule",
         "Properties": {
            "MonitoringScheduleConfig": {
               "MonitoringJobDefinition": {
                  "MonitoringAppSpecification": {
                     "ImageUri": {
                        "Fn::Sub": "123456789012.dkr.ecr.${AWS::Partition}.amazonaws.com/sagemaker-model-monitor-analyzer:latest"
                     }
                  },
                  "MonitoringInputs": [
                     {
                        "EndpointInput": {
                           "EndpointName": {
                              "Fn::ImportValue": "CanaryEndpointName"
                           },
                           "LocalPath": "/opt/ml/processing/endpointdata"
                        }
                     }
                  ],
                  "MonitoringOutputConfig": {
                     "MonitoringOutputs": [
                        {
                           "S3Output": {
                              "LocalPath": "/opt/ml/processing/localpath",
                              "S3Uri": "s3://endpoint-data-capture/myEndpoint"
                           }
                        }
                     ]
                  },
                  "MonitoringResources": {
                     "ClusterConfig": {
                        "InstanceCount": 1,
                        "InstanceType": "ml.m5.large",
                        "VolumeSizeInGB": 50
                     }
                  },
                  "RoleArn": null
               },
               "ScheduleConfig": {
                  "ScheduleExpression": "cron(0 * ? * * *)"
               }
            },
            "MonitoringScheduleName": "BasicMonitoringSchedule"
         }
      }
   }
}
```

#### YAML

```yaml

---
AWSTemplateFormatVersion: '2010-09-09'
Description: Basic SageMaker Hosting entities to create a monitoring schedule

Description: "Basic Hosting entities test.  We need models to create endpoint configs."
Mappings:
  RegionMap:
    "us-west-2":
      "NullTransformer": "123456789012.dkr.ecr.us-west-2.amazonaws.com/mymodel:latest"
    "us-east-2":
      "NullTransformer": "123456789012.dkr.ecr.us-east-2.amazonaws.com/mymodel:latest"
    "us-east-1":
      "NullTransformer": "123456789012.dkr.ecr.us-east-1.amazonaws.com/mymodel:latest"
    "eu-west-1":
      "NullTransformer": "123456789012.dkr.ecr.eu-west-1.amazonaws.com/mymodel:latest"
    "ap-northeast-1":
      "NullTransformer": "123456789012.dkr.ecr.ap-northeast-1.amazonaws.com/mymodel:latest"
    "ap-northeast-2":
      "NullTransformer": "123456789012.dkr.ecr.ap-northeast-2.amazonaws.com/mymodel:latest"
    "ap-southeast-2":
      "NullTransformer": "123456789012.dkr.ecr.ap-southeast-2.amazonaws.com/mymodel:latest"
    "eu-central-1":
      "NullTransformer": "123456789012.dkr.ecr.eu-central-1.amazonaws.com/mymodel:latest"
Resources:
  Endpoint:
    Type: "AWS::SageMaker::Endpoint"
    Properties:
      EndpointConfigName:
        !GetAtt EndpointConfig.EndpointConfigName
  EndpointConfig:
    Type: "AWS::SageMaker::EndpointConfig"
    Properties:
      ProductionVariants:
        - InitialInstanceCount: 1
          InitialVariantWeight: 1.0
          InstanceType: ml.t2.large
          ModelName: !GetAtt Model.ModelName
          VariantName: !GetAtt Model.ModelName
  Model:
    Type: "AWS::SageMaker::Model"
    Properties:
      PrimaryContainer:
        Image: !FindInMap [RegionMap, !Ref "AWS::Region", "NullTransformer"]
      ExecutionRoleArn: !GetAtt ExecutionRole.Arn

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
      Policies:
        -
          PolicyName: "root"
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
              -
                Effect: "Allow"
                Action: "*"
                Resource: "*"

  MonitoringScheduleExecutionRole:
    Type: "AWS::IAM::Role"
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          - Effect: "Allow"
            Principal:
              Service:
                - "sagemaker.amazonaws.com"
            Action:
              - "sts:AssumeRole"
      ManagedPolicyArns:
        - Fn::Sub: "arn:${AWS::Partition}:iam::aws:policy/AmazonSageMakerFullAccess"
        - Fn::Sub: "arn:${AWS::Partition}:iam::aws:policy/AmazonS3FullAccess"
        - Fn::Sub: "arn:${AWS::Partition}:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"

  MonitoringSchedule:
    Type: AWS::SageMaker::MonitoringSchedule
    Properties:
      MonitoringScheduleConfig:
        MonitoringJobDefinition:
          MonitoringAppSpecification:
            ImageUri:
              Fn::Sub: "123456789012.dkr.ecr.${AWS::Partition}.amazonaws.com/sagemaker-model-monitor-analyzer:latest"
          MonitoringInputs:
            - EndpointInput:
                EndpointName:
                  Fn::ImportValue: CanaryEndpointName
                LocalPath: "/opt/ml/processing/endpointdata"
          MonitoringOutputConfig:
            MonitoringOutputs:
              - S3Output:
                  LocalPath: "/opt/ml/processing/localpath"
                  S3Uri: s3://endpoint-data-capture/myEndpoint
          MonitoringResources:
            ClusterConfig:
              InstanceCount: 1
              InstanceType: ml.m5.large
              VolumeSizeInGB: 50
          RoleArn: !GetAtt MonitoringScheduleExecutionRole.Arn
        ScheduleConfig:
          ScheduleExpression: cron(0 * ? * * *)
      MonitoringScheduleName: BasicMonitoringSchedule

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfig

BaselineConfig

All content copied from https://docs.aws.amazon.com/.
