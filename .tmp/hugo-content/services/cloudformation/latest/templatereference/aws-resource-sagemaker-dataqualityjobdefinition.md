This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::DataQualityJobDefinition

Creates a definition for a job that monitors data quality and drift. For information
about model monitor, see [Amazon SageMaker AI Model\
Monitor](../../../sagemaker/latest/dg/model-monitor.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::DataQualityJobDefinition",
  "Properties" : {
      "DataQualityAppSpecification" : DataQualityAppSpecification,
      "DataQualityBaselineConfig" : DataQualityBaselineConfig,
      "DataQualityJobInput" : DataQualityJobInput,
      "DataQualityJobOutputConfig" : MonitoringOutputConfig,
      "EndpointName" : String,
      "JobDefinitionName" : String,
      "JobResources" : MonitoringResources,
      "NetworkConfig" : NetworkConfig,
      "RoleArn" : String,
      "StoppingCondition" : StoppingCondition,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::DataQualityJobDefinition
Properties:
  DataQualityAppSpecification:
    DataQualityAppSpecification
  DataQualityBaselineConfig:
    DataQualityBaselineConfig
  DataQualityJobInput:
    DataQualityJobInput
  DataQualityJobOutputConfig:
    MonitoringOutputConfig
  EndpointName: String
  JobDefinitionName: String
  JobResources:
    MonitoringResources
  NetworkConfig:
    NetworkConfig
  RoleArn: String
  StoppingCondition:
    StoppingCondition
  Tags:
    - Tag

```

## Properties

`DataQualityAppSpecification`

Specifies the container that runs the monitoring job.

_Required_: Yes

_Type_: [DataQualityAppSpecification](aws-properties-sagemaker-dataqualityjobdefinition-dataqualityappspecification.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataQualityBaselineConfig`

Configures the constraints and baselines for the monitoring job.

_Required_: No

_Type_: [DataQualityBaselineConfig](aws-properties-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataQualityJobInput`

A list of inputs for the monitoring job. Currently endpoints are supported as monitoring inputs.

_Required_: Yes

_Type_: [DataQualityJobInput](aws-properties-sagemaker-dataqualityjobdefinition-dataqualityjobinput.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataQualityJobOutputConfig`

The output configuration for monitoring jobs.

_Required_: Yes

_Type_: [MonitoringOutputConfig](aws-properties-sagemaker-dataqualityjobdefinition-monitoringoutputconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointName`

The name of the endpoint used to run the monitoring job.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JobDefinitionName`

The name for the monitoring job definition.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JobResources`

Identifies the resources to deploy for a monitoring job.

_Required_: Yes

_Type_: [MonitoringResources](aws-properties-sagemaker-dataqualityjobdefinition-monitoringresources.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkConfig`

Specifies networking configuration for the monitoring job.

_Required_: No

_Type_: [NetworkConfig](aws-properties-sagemaker-dataqualityjobdefinition-networkconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your
behalf.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StoppingCondition`

A time limit for how long the monitoring job is allowed to run before stopping.

_Required_: No

_Type_: [StoppingCondition](aws-properties-sagemaker-dataqualityjobdefinition-stoppingcondition.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-dataqualityjobdefinition-tag.md)

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the job definition.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The time when the job definition was created.

`JobDefinitionArn`

The Amazon Resource Name (ARN) of the job definition.

## Examples

### SageMaker DataQualityJobDefinition Example

The following example creates a Data Quality monitoring job defintion.

#### JSON

```json

{
   "AWSTemplateFormatVersion": "2010-09-09",
   "Description": "Basic SageMaker Hosting entities to create a data quality job definition",
   "Mappings": {
      "RegionMap": {
         "us-west-2": {
            "MyModelImage": "123456789012.dkr.ecr.us-west-2.amazonaws.com/mymodel:latest"
         },
         "us-east-2": {
            "MyModelImage": "123456789012.dkr.ecr.us-east-2.amazonaws.com/mymodel:latest"
         },
         "us-east-1": {
            "MyModelImage": "123456789012.dkr.ecr.us-east-1.amazonaws.com/mymodel:latest"
         },
         "eu-west-1": {
            "MyModelImage": "123456789012.dkr.ecr.eu-west-1.amazonaws.com/mymodel:latest"
         },
         "ap-northeast-1": {
            "MyModelImage": "123456789012.dkr.ecr.ap-northeast-1.amazonaws.com/mymodel:latest"
         },
         "ap-northeast-2": {
            "MyModelImage": "123456789012.dkr.ecr.ap-northeast-2.amazonaws.com/mymodel:latest"
         },
         "ap-southeast-2": {
            "MyModelImage": "123456789012.dkr.ecr.ap-southeast-2.amazonaws.com/mymodel:latest"
         },
         "eu-central-1": {
            "MyModelImage": "123456789012.dkr.ecr.eu-central-1.amazonaws.com/mymodel:latest"
         }
      }
   },
   "Resources": {
      "Endpoint": {
         "Type": "AWS::SageMaker::Endpoint",
         "Properties": {
            "EndpointConfigName": {
            "Fn::GetAtt": [
              "EndpointConfigWithDataCapture",
              "EndpointConfigName"
             ]

            }
         }
      },
      "EndpointConfigWithDataCapture": {
         "Type": "AWS::SageMaker::EndpointConfig",
         "Properties": {
            "ProductionVariants": [
               {
                  "InitialInstanceCount": 1,
                  "InitialVariantWeight": 1,
                  "InstanceType": "ml.t2.large",
                  "ModelName": {
                     "Fn::GetAtt": [
                       "Model",
                       "ModelName"
                      ]

                },
                  "VariantName": {
                     "Fn::GetAtt": [
                       "Model",
                       "ModelName"
                      ]

                }
               }
            ],
            "DataCaptureConfig": {
               "EnableCapture": true,
               "InitialSamplingPercentage": 100,
               "DestinationS3Uri": "s3://bucket/prefix",
               "KmsKeyId": "kmskeyid",
               "CaptureOptions": [
                  {
                     "CaptureMode": "Input"
                  },
                  {
                     "CaptureMode": "Output"
                  }
               ],
               "CaptureContentTypeHeader": {
                  "CsvContentTypes": [
                     "text/csv"
                  ],
                  "JsonContentTypes": [
                     "appplication/json"
                  ]
               }
            }
         }
      },
      "Model": {
         "Type": "AWS::SageMaker::Model",
         "Properties": {
            "PrimaryContainer": {
               "Image": { "Fn::FindInMap": [
                  "RegionMap",
                  {"Ref": "AWS::Region"},
                  "MyModelImage"
                 ]
               }
            },
            "ExecutionRoleArn": {
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
      "JobDefinitionExecutionRole": {
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
      "DataQualityJobDefinition": {
         "Type": "AWS::SageMaker::DataQualityJobDefinition",
         "Properties": {
            "DataQualityAppSpecification": {
               "ImageUri": {
                  "Fn::Sub": "123456789012.dkr.ecr.${AWS::Partition}.amazonaws.com/sagemaker-model-monitor-analyzer:latest"
               }
            },
            "DataQualityJobInput": {
               "EndpointInput": {
                  "EndpointName": {
                     "Fn::GetAtt": [
                       "Endpoint",
                       "EndpointName"
                      ]

                },
                  "LocalPath": "/opt/ml/processing/endpointdata"
               }
            },
            "DataQualityJobOutputConfig": {
               "MonitoringOutputs": [
                  {
                     "S3Output": {
                        "LocalPath": "/opt/ml/processing/localpath",
                        "S3Uri": {
                           "Fn::Sub": "s3://data-quality-job-definition-${AWS::AccountId}/output"
                        }
                     }
                  }
               ]
            },
            "JobResources": {
               "ClusterConfig": {
                  "InstanceCount": 1,
                  "InstanceType": "ml.m5.large",
                  "VolumeSizeInGB": 50
               }
            },
            "RoleArn": null,
            "StoppingCondition": {
               "MaxRuntimeInSeconds": 2000
            }
         }
      }
   }
}
```

#### YAML

```yaml

---

AWSTemplateFormatVersion: '2010-09-09'
Description: Basic SageMaker Hosting entities to create a data quality job definition

Mappings:
  RegionMap:
    "us-west-2":
      "MyModelImage": "123456789012.dkr.ecr.us-west-2.amazonaws.com/mymodel:latest"
    "us-east-2":
      "MyModelImage": "123456789012.dkr.ecr.us-east-2.amazonaws.com/mymodel:latest"
    "us-east-1":
      "MyModelImage": "123456789012.dkr.ecr.us-east-1.amazonaws.com/mymodel:latest"
    "eu-west-1":
      "MyModelImage": "123456789012.dkr.ecr.eu-west-1.amazonaws.com/mymodel:latest"
    "ap-northeast-1":
      "MyModelImage": "123456789012.dkr.ecr.ap-northeast-1.amazonaws.com/mymodel:latest"
    "ap-northeast-2":
      "MyModelImage": "123456789012.dkr.ecr.ap-northeast-2.amazonaws.com/mymodel:latest"
    "ap-southeast-2":
      "MyModelImage": "123456789012.dkr.ecr.ap-southeast-2.amazonaws.com/mymodel:latest"
    "eu-central-1":
      "MyModelImage": "123456789012.dkr.ecr.eu-central-1.amazonaws.com/mymodel:latest"

Resources:
  Endpoint:
    Type: "AWS::SageMaker::Endpoint"
    Properties:
      EndpointConfigName:
        !GetAtt EndpointConfigWithDataCapture.EndpointConfigName

  EndpointConfigWithDataCapture:
    Type: "AWS::SageMaker::EndpointConfig"
    Properties:
      ProductionVariants:
        - InitialInstanceCount: 1
          InitialVariantWeight: 1.0
          InstanceType: ml.t2.large
          ModelName: !GetAtt Model.ModelName
          VariantName: !GetAtt Model.ModelName
      DataCaptureConfig:
        EnableCapture: true
        InitialSamplingPercentage: 100
        DestinationS3Uri: s3://bucket/prefix
        KmsKeyId: kmskeyid
        CaptureOptions:
          - CaptureMode: Input
          - CaptureMode: Output
        CaptureContentTypeHeader:
          CsvContentTypes:
            - "text/csv"
          JsonContentTypes:
            - "appplication/json"

  Model:
    Type: "AWS::SageMaker::Model"
    Properties:
      PrimaryContainer:
        Image: !FindInMap [RegionMap, !Ref "AWS::Region", "MyModelImage"]
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

  JobDefinitionExecutionRole:
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

  DataQualityJobDefinition:
    Type: AWS::SageMaker::DataQualityJobDefinition
    Properties:
      DataQualityAppSpecification:
        ImageUri:
          Fn::Sub: "123456789012.dkr.ecr.${AWS::Partition}.amazonaws.com/sagemaker-model-monitor-analyzer:latest"
      DataQualityJobInput:
        EndpointInput:
          EndpointName: !GetAtt Endpoint.EndpointName
          LocalPath: "/opt/ml/processing/endpointdata"
      DataQualityJobOutputConfig:
        MonitoringOutputs:
          - S3Output:
              LocalPath: "/opt/ml/processing/localpath"
              S3Uri:
                Fn::Sub: "s3://data-quality-job-definition-${AWS::AccountId}/output"
      JobResources:
        ClusterConfig:
          InstanceCount: 1
          InstanceType: ml.m5.large
          VolumeSizeInGB: 50
      RoleArn: !GetAtt JobDefinitionExecutionRole.Arn
      StoppingCondition:
        MaxRuntimeInSeconds: 2000
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

BatchTransformInput

All content copied from https://docs.aws.amazon.com/.
