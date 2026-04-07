This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application

Creates an Amazon Kinesis Data Analytics application. For information about creating a
Kinesis Data Analytics application, see [Creating an\
Application](https://docs.aws.amazon.com/managed-flink/latest/java/getting-started.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KinesisAnalyticsV2::Application",
  "Properties" : {
      "ApplicationConfiguration" : ApplicationConfiguration,
      "ApplicationDescription" : String,
      "ApplicationMaintenanceConfiguration" : ApplicationMaintenanceConfiguration,
      "ApplicationMode" : String,
      "ApplicationName" : String,
      "RunConfiguration" : RunConfiguration,
      "RuntimeEnvironment" : String,
      "ServiceExecutionRole" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::KinesisAnalyticsV2::Application
Properties:
  ApplicationConfiguration:
    ApplicationConfiguration
  ApplicationDescription: String
  ApplicationMaintenanceConfiguration:
    ApplicationMaintenanceConfiguration
  ApplicationMode: String
  ApplicationName: String
  RunConfiguration:
    RunConfiguration
  RuntimeEnvironment: String
  ServiceExecutionRole: String
  Tags:
    - Tag

```

## Properties

`ApplicationConfiguration`

Use this parameter to configure the application.

_Required_: No

_Type_: [ApplicationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationDescription`

The description of the application.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationMaintenanceConfiguration`

Specifies the maintenance window parameters for a Kinesis Data Analytics
application.

_Required_: No

_Type_: [ApplicationMaintenanceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-applicationmaintenanceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationMode`

To create a Kinesis Data Analytics Studio notebook, you must set the mode to
`INTERACTIVE`. However, for a Kinesis Data Analytics for Apache Flink
application, the mode is optional.

_Required_: No

_Type_: String

_Allowed values_: `INTERACTIVE | STREAMING`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ApplicationName`

The name of the application.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RunConfiguration`

Describes the starting parameters for an Managed Service for Apache Flink application.

_Required_: No

_Type_: [RunConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-runconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuntimeEnvironment`

The runtime environment for the application.

_Required_: Yes

_Type_: String

_Allowed values_: `SQL-1_0 | FLINK-1_6 | FLINK-1_8 | ZEPPELIN-FLINK-1_0 | FLINK-1_11 | FLINK-1_13 | ZEPPELIN-FLINK-2_0 | FLINK-1_15 | ZEPPELIN-FLINK-3_0 | FLINK-1_18 | FLINK-1_19 | FLINK-1_20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceExecutionRole`

Specifies the IAM role that the application uses to access external resources.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of one or more tags to assign to the application. A tag is a key-value pair
that identifies an application. Note that the maximum number of application tags
includes system tags. The maximum number of user-defined application tags is 50.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-tag.html)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Creating an Amazon Kinesis Data Analytics Application using Apache Flink](#aws-resource-kinesisanalyticsv2-application--examples--Creating_an_Amazon_Kinesis_Data_Analytics_Application_using_Apache_Flink)

- [Creating an Amazon Kinesis Data Analytics Studio Application](#aws-resource-kinesisanalyticsv2-application--examples--Creating_an_Amazon_Kinesis_Data_Analytics_Studio_Application)

### Creating an Amazon Kinesis Data Analytics Application using Apache Flink

The following example shows how to create a simple application by using a
deployment package from Amazon S3. You must add permissions to the IAM role to
access any streams that your code requires.

#### JSON

```json

{
    "Description": "Simple KDA Flink application",
    "Parameters": {
        "CodeBucketArn": {
            "Type": "String"
        },
        "CodeKey": {
            "Type": "String"
        }
    },
    "Resources": {
        "MyApplication": {
            "Type": "AWS::KinesisAnalyticsV2::Application",
            "Properties": {
                "RuntimeEnvironment": "FLINK-1_15",
                "ServiceExecutionRole": {
                    "Fn::GetAtt": [
                        "ServiceExecutionRole",
                        "Arn"
                    ]
                },
                "ApplicationConfiguration": {
                    "ApplicationCodeConfiguration": {
                        "CodeContent": {
                            "S3ContentLocation": {
                                "BucketARN": {
                                    "Ref": "CodeBucketArn"
                                },
                                "FileKey": {
                                    "Ref": "CodeKey"
                                }
                            }
                        },
                        "CodeContentType": "ZIPFILE"
                    }
                }
            }
        },
        "ServiceExecutionRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "AssumeRolePolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": "kinesisanalytics.amazonaws.com"
                            },
                            "Action": "sts:AssumeRole"
                        }
                    ]
                },
                "Path": "/",
                "Policies": [
                    {
                        "PolicyName": "s3-code-access",
                        "PolicyDocument": {
                            "Version": "2012-10-17",
                            "Statement": [
                                {
                                    "Effect": "Allow",
                                    "Action": [
                                        "s3:GetObject"
                                    ],
                                    "Resource": [
                                        {
                                            "Fn::Sub": "${CodeBucketArn}/${CodeKey}"
                                        }
                                    ]
                                }
                            ]
                        }
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Description: Simple KDA Flink application
Parameters:
  CodeBucketArn:
    Type: String
  CodeKey:
    Type: String
​
Resources:
  MyApplication:
    Type: AWS::KinesisAnalyticsV2::Application
    Properties:
      RuntimeEnvironment: FLINK-1_15
      ServiceExecutionRole: !GetAtt ServiceExecutionRole.Arn
      ApplicationConfiguration:
        ApplicationCodeConfiguration:
          CodeContent:
            S3ContentLocation:
              BucketARN: !Ref CodeBucketArn
              FileKey: !Ref CodeKey
          CodeContentType: 'ZIPFILE'
​
  ServiceExecutionRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2012-10-17
        Statement:
          - Effect: Allow
            Principal:
              Service: kinesisanalytics.amazonaws.com
            Action: 'sts:AssumeRole'
      Path: /
      Policies:
        - PolicyName: s3-code-access
          PolicyDocument:
            Version: 2012-10-17
            Statement:
              - Effect: Allow
                Action:
                  - s3:GetObject
                Resource:
                  - !Sub "${CodeBucketArn}/${CodeKey}"
```

### Creating an Amazon Kinesis Data Analytics Studio Application

The following example shows how to create a simple Studio application with an
Amazon Glue database. You must add permissions to the IAM role to create or
access any streams you require, and any that already exist must be added to the
Glue database.

#### JSON

```json

{
    "Description": "KDA Studio application",
    "Parameters": {
        "GlueDatabaseName": {
            "Type": "String"
        }
    },
    "Resources": {
        "MyApplication": {
            "Type": "AWS::KinesisAnalyticsV2::Application",
            "Properties": {
                "ApplicationMode": "INTERACTIVE",
                "RuntimeEnvironment": "ZEPPELIN-FLINK-3_0",
                "ServiceExecutionRole": {
                    "Fn::GetAtt": [
                        "ServiceExecutionRole",
                        "Arn"
                    ]
                },
                "ApplicationConfiguration": {
                    "FlinkApplicationConfiguration": {
                        "ParallelismConfiguration": {
                            "Parallelism": 4,
                            "ConfigurationType": "CUSTOM"
                        }
                    },
                    "ZeppelinApplicationConfiguration": {
                        "CatalogConfiguration": {
                            "GlueDataCatalogConfiguration": {
                                "DatabaseARN": {
                                    "Fn::Sub": "arn:aws:glue:${AWS::Region}:${AWS::AccountId}:database/${GlueDatabase}"
                                }
                            }
                        },
                        "CustomArtifactsConfiguration": [
                            {
                                "ArtifactType": "DEPENDENCY_JAR",
                                "MavenReference": {
                                    "GroupId": "org.apache.flink",
                                    "ArtifactId": "flink-sql-connector-kinesis",
                                    "Version": "1.15.4"
                                }
                            },
                            {
                                "ArtifactType": "DEPENDENCY_JAR",
                                "MavenReference": {
                                    "GroupId": "org.apache.flink",
                                    "ArtifactId": "flink-connector-kafka",
                                    "Version": "1.15.4"
                                }
                            },
                            {
                                "ArtifactType": "DEPENDENCY_JAR",
                                "MavenReference": {
                                    "GroupId": "software.amazon.msk",
                                    "ArtifactId": "aws-msk-iam-auth",
                                    "Version": "1.1.6"
                                }
                            }
                        ]
                    }
                }
            }
        },
        "GlueDatabase": {
            "Type": "AWS::Glue::Database",
            "Properties": {
                "CatalogId": {
                    "Ref": "AWS::AccountId"
                },
                "DatabaseInput": {
                    "Name": {
                        "Ref": "GlueDatabaseName"
                    },
                    "Description": "My glue database"
                }
            }
        },
        "ServiceExecutionRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "AssumeRolePolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": "kinesisanalytics.amazonaws.com"
                            },
                            "Action": "sts:AssumeRole"
                        }
                    ]
                },
                "Path": "/",
                "Policies": [
                    {
                        "PolicyName": "glue-access",
                        "PolicyDocument": {
                            "Version": "2012-10-17",
                            "Statement": [
                                {
                                    "Effect": "Allow",
                                    "Action": [
                                        "glue:GetConnection",
                                        "glue:GetTable",
                                        "glue:GetTables",
                                        "glue:CreateTable",
                                        "glue:UpdateTable",
                                        "glue:GetDatabases",
                                        "glue:GetUserDefinedFunction"
                                    ],
                                    "Resource": [
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:glue:${AWS::Region}:${AWS::AccountId}:connection/*"
                                        },
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:glue:${AWS::Region}:${AWS::AccountId}:table/*"
                                        },
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:glue:${AWS::Region}:${AWS::AccountId}:database/${GlueDatabase}/*"
                                        },
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:glue:${AWS::Region}:${AWS::AccountId}:catalog"
                                        },
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:glue:${AWS::Region}:${AWS::AccountId}:userDefinedFunction/*"
                                        }
                                    ]
                                },
                                {
                                    "Effect": "Allow",
                                    "Action": [
                                        "glue:GetDatabase"
                                    ],
                                    "Resource": [
                                        "*"
                                    ]
                                }
                            ]
                        }
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Description: KDA Studio application
Parameters:
  GlueDatabaseName:
    Type: String

Resources:
  MyApplication:
    Type: AWS::KinesisAnalyticsV2::Application
    Properties:
      ApplicationMode: INTERACTIVE
      RuntimeEnvironment: ZEPPELIN-FLINK-3_0
      ServiceExecutionRole: !GetAtt ServiceExecutionRole.Arn
      ApplicationConfiguration:
        FlinkApplicationConfiguration:
          ParallelismConfiguration:
            Parallelism: 4
            ConfigurationType: CUSTOM
        ZeppelinApplicationConfiguration:
          CatalogConfiguration:
            GlueDataCatalogConfiguration:
              DatabaseARN: !Sub "arn:aws:glue:${AWS::Region}:${AWS::AccountId}:database/${GlueDatabase}"
          CustomArtifactsConfiguration:
            - ArtifactType: DEPENDENCY_JAR
              MavenReference:
                GroupId: org.apache.flink
                ArtifactId: flink-sql-connector-kinesis
                Version: 1.15.4
            - ArtifactType: DEPENDENCY_JAR
              MavenReference:
                GroupId: org.apache.flink
                ArtifactId: flink-connector-kafka
                Version: 1.15.4
            - ArtifactType: DEPENDENCY_JAR
              MavenReference:
                GroupId: software.amazon.msk
                ArtifactId: aws-msk-iam-auth
                Version: 1.1.6

  GlueDatabase:
    Type: AWS::Glue::Database
    Properties:
      CatalogId: !Ref AWS::AccountId
      DatabaseInput:
        Name: !Ref GlueDatabaseName
        Description: My glue database

  ServiceExecutionRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2012-10-17
        Statement:
          - Effect: Allow
            Principal:
              Service: kinesisanalytics.amazonaws.com
            Action: 'sts:AssumeRole'
      Path: /
      Policies:
        - PolicyName: glue-access
          PolicyDocument:
            Version: 2012-10-17
            Statement:
              - Effect: Allow
                Action:
                  - glue:GetConnection
                  - glue:GetTable
                  - glue:GetTables
                  - glue:CreateTable
                  - glue:UpdateTable
                  - glue:GetDatabases
                  - glue:GetUserDefinedFunction
                Resource:
                  - !Sub "arn:${AWS::Partition}:glue:${AWS::Region}:${AWS::AccountId}:connection/*"
                  - !Sub "arn:${AWS::Partition}:glue:${AWS::Region}:${AWS::AccountId}:table/*"
                  - !Sub "arn:${AWS::Partition}:glue:${AWS::Region}:${AWS::AccountId}:database/${GlueDatabase}/*"
                  - !Sub "arn:${AWS::Partition}:glue:${AWS::Region}:${AWS::AccountId}:catalog"
                  - !Sub "arn:${AWS::Partition}:glue:${AWS::Region}:${AWS::AccountId}:userDefinedFunction/*"
              - Effect: Allow
                Action:
                  - glue:GetDatabase
                Resource:
                  - "*"
```

## See also

- [CreateApplication](https://docs.aws.amazon.com/managed-flink/latest/apiv2/API_CreateApplication.html) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Managed Service for Apache Flink V2

ApplicationCodeConfiguration
