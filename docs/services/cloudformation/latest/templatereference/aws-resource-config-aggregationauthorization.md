This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::AggregationAuthorization

An object that represents the authorizations granted to
aggregator accounts and regions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Config::AggregationAuthorization",
  "Properties" : {
      "AuthorizedAccountId" : String,
      "AuthorizedAwsRegion" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Config::AggregationAuthorization
Properties:
  AuthorizedAccountId: String
  AuthorizedAwsRegion: String
  Tags:
    - Tag

```

## Properties

`AuthorizedAccountId`

The 12-digit account ID of the account authorized to aggregate
data.

_Required_: Yes

_Type_: String

_Pattern_: `^\d{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuthorizedAwsRegion`

The region authorized to collect aggregated data.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of tag object.

_Required_: No

_Type_: Array of [Tag](aws-properties-config-aggregationauthorization-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the AggregationAuthorization, such as `arn:aws:config:us-east-1:123456789012:aggregation-authorization/987654321012/us-west-2`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`AggregationAuthorizationArn`

The Amazon Resource Name (ARN) of the aggregation object.

## Examples

- [Authorize Another Account](#aws-resource-config-aggregationauthorization--examples--Authorize_Another_Account)

- [Aggregation Authorization](#aws-resource-config-aggregationauthorization--examples--Aggregation_Authorization)

### Authorize Another Account

The following example creates an AggregationAuthorization that authorizes another account to aggregate your AWS Config data into a specific region.

#### JSON

```json

"AggregationAuthorization": {
    "Type": "AWS::Config::AggregationAuthorization",
    "Properties": {
        "AuthorizedAccountId": 123456789012,
        "AuthorizedAwsRegion": "us-west-2"
    }
}
```

#### YAML

```yaml

AggregationAuthorization:
  Type: "AWS::Config::AggregationAuthorization"
  Properties:
    AuthorizedAccountId: 123456789012
    AuthorizedAwsRegion: us-west-2
```

### Aggregation Authorization

The following example enables AWS Config and creates an AWS Config rule, an
aggregator, and an authorization.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Enable Config",
    "Metadata": {
        "AWS::CloudFormation::Interface": {
            "ParameterGroups": [
                {
                    "Label": {
                        "default": "Configuration Recorder Configuration"
                    },
                    "Parameters": [
                        "GlobalResourceTypesRegion"
                    ]
                },
                {
                    "Label": {
                        "default": "Configuration Aggregator Configuration"
                    },
                    "Parameters": [
                        "AggregatorAccount",
                        "AggregatorRegion",
                        "SourceAccounts",
                        "SourceRegions"
                    ]
                }
            ],
            "ParameterLabels": {
                "GlobalResourceTypesRegion": {
                    "default": "Global resource types region"
                },
                "AggregatorAccount": {
                    "default": "Aggregator account"
                },
                "AggregatorRegion": {
                    "default": "Aggregator region"
                },
                "SourceAccounts": {
                    "default": "Source accounts"
                },
                "SourceRegions": {
                    "default": "Source regions"
                }
            }
        }
    },
    "Parameters": {
        "GlobalResourceTypesRegion": {
            "Type": "String",
            "Default": "us-east-1",
            "Description": "AWS region used to record global resources types"
        },
        "AggregatorAccount": {
            "Type": "String",
            "Description": "Account ID of the aggregator"
        },
        "AggregatorRegion": {
            "Type": "String",
            "Default": "us-east-1",
            "Description": "AWS region of the aggregator"
        },
        "SourceAccounts": {
            "Type": "CommaDelimitedList",
            "Description": "List of source accounts to aggregate"
        },
        "SourceRegions": {
            "Type": "CommaDelimitedList",
            "Description": "List of regions to aggregate"
        }
    },
    "Conditions": {
        "IncludeGlobalResourceTypes": {
            "Fn::Equals": [
                {
                    "Ref": "GlobalResourceTypesRegion"
                },
                {
                    "Ref": "AWS::Region"
                }
            ]
        },
        "CreateAggregator": {
            "Fn::And": [
                {
                    "Fn::Equals": [
                        {
                            "Ref": "AggregatorAccount"
                        },
                        {
                            "Ref": "AWS::AccountId"
                        }
                    ]
                },
                {
                    "Fn::Equals": [
                        {
                            "Ref": "AggregatorRegion"
                        },
                        {
                            "Ref": "AWS::Region"
                        }
                    ]
                }
            ]
        },
        "CreateAuthorization": {
            "Fn::Not": [
                {
                    "Fn::Equals": [
                        {
                            "Ref": "AggregatorAccount"
                        },
                        {
                            "Ref": "AWS::AccountId"
                        }
                    ]
                }
            ]
        }
    },
    "Resources": {
        "ConfigBucket": {
            "DeletionPolicy": "Retain",
            "Type": "AWS::S3::Bucket"
        },
        "ConfigBucketPolicy": {
            "Type": "AWS::S3::BucketPolicy",
            "Properties": {
                "Bucket": {
                    "Ref": "ConfigBucket"
                },
                "PolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Sid": "AWSConfigBucketPermissionsCheck",
                            "Effect": "Allow",
                            "Principal": {
                                "Service": [
                                    "config.amazonaws.com"
                                ]
                            },
                            "Action": "s3:GetBucketAcl",
                            "Resource": [
                                {
                                    "Fn::Sub": "arn:aws:s3:::${ConfigBucket}"
                                }
                            ]
                        },
                        {
                            "Sid": "AWSConfigBucketDelivery",
                            "Effect": "Allow",
                            "Principal": {
                                "Service": [
                                    "config.amazonaws.com"
                                ]
                            },
                            "Action": "s3:PutObject",
                            "Resource": [
                                {
                                    "Fn::Sub": "arn:aws:s3:::${ConfigBucket}/AWSLogs/${AWS::AccountId}/*"
                                }
                            ]
                        }
                    ]
                }
            }
        },
        "ConfigRecorderRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "AssumeRolePolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": [
                                    "config.amazonaws.com"
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
                    "arn:aws:iam::aws:policy/service-role/AWSConfigRole"
                ]
            }
        },
        "ConfigRecorder": {
            "Type": "AWS::Config::ConfigurationRecorder",
            "DependsOn": [
                "ConfigRecorderRole",
                "ConfigBucketPolicy"
            ],
            "Properties": {
                "RoleARN": {
                    "Fn::GetAtt": [
                        "ConfigRecorderRole",
                        "Arn"
                    ]
                },
                "RecordingGroup": {
                    "AllSupported": true,
                    "IncludeGlobalResourceTypes": {
                        "Fn::If": [
                            "IncludeGlobalResourceTypes",
                            true,
                            false
                        ]
                    }
                }
            }
        },
        "DeliveryChannel": {
            "Type": "AWS::Config::DeliveryChannel",
            "DependsOn": [
                "ConfigBucketPolicy"
            ],
            "Properties": {
                "Name": "default",
                "S3BucketName": {
                    "Ref": "ConfigBucket"
                }
            }
        },
        "S3BucketPublicReadRule": {
            "Type": "AWS::Config::ConfigRule",
            "DependsOn": [
                "ConfigRecorder"
            ],
            "Properties": {
                "ConfigRuleName": "stackset-s3-bucket-public-read-prohibited",
                "Description": "s3-bucket-public-read-prohibited from stackset",
                "Scope": {
                    "ComplianceResourceTypes": [
                        "AWS::S3::Bucket"
                    ]
                },
                "Source": {
                    "Owner": "AWS",
                    "SourceIdentifier": "S3_BUCKET_PUBLIC_READ_PROHIBITED"
                }
            }
        },
        "ConfigAggregator": {
            "Type": "AWS::Config::ConfigurationAggregator",
            "Condition": "CreateAggregator",
            "Properties": {
                "Name": "name",
                "AccountAggregationSources": [
                    {
                        "AccountIds": {
                            "Ref": "SourceAccounts"
                        },
                        "AwsRegions": {
                            "Ref": "SourceRegions"
                        }
                    }
                ]
            }
        },
        "AggregationAuthorization": {
            "Type": "AWS::Config::AggregationAuthorization",
            "Condition": "CreateAuthorization",
            "Properties": {
                "AuthorizedAccountId": {
                    "Ref": "AggregatorAccount"
                },
                "AuthorizedAwsRegion": {
                    "Ref": "AggregatorRegion"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Enable Config

Metadata:
  AWS::CloudFormation::Interface:
    ParameterGroups:
      - Label:
          default: Configuration Recorder Configuration
        Parameters:
          - GlobalResourceTypesRegion
      - Label:
          default: Configuration Aggregator Configuration
        Parameters:
          - AggregatorAccount
          - AggregatorRegion
          - SourceAccounts
          - SourceRegions
    ParameterLabels:
      GlobalResourceTypesRegion:
        default: Global resource types region
      AggregatorAccount:
        default: Aggregator account
      AggregatorRegion:
        default: Aggregator region
      SourceAccounts:
        default: Source accounts
      SourceRegions:
        default: Source regions

Parameters:
  GlobalResourceTypesRegion:
    Type: String
    Default: us-east-1
    Description: AWS region used to record global resources types
  AggregatorAccount:
    Type: String
    Description: Account ID of the aggregator
  AggregatorRegion:
    Type: String
    Default: us-east-1
    Description: AWS region of the aggregator
  SourceAccounts:
    Type: CommaDelimitedList
    Description: List of source accounts to aggregate
  SourceRegions:
    Type: CommaDelimitedList
    Description: List of regions to aggregate

Conditions:
  IncludeGlobalResourceTypes: !Equals
    - !Ref GlobalResourceTypesRegion
    - !Ref AWS::Region
  CreateAggregator: !And
    - !Equals
      - !Ref AggregatorAccount
      - !Ref AWS::AccountId
    - !Equals
      - !Ref AggregatorRegion
      - !Ref AWS::Region
  CreateAuthorization: !Not
    - !Equals
      - !Ref AggregatorAccount
      - !Ref AWS::AccountId

Resources:

  ConfigBucket:
    DeletionPolicy: Retain
    Type: AWS::S3::Bucket

  ConfigBucketPolicy:
    Type: AWS::S3::BucketPolicy
    Properties:
      Bucket: !Ref ConfigBucket
      PolicyDocument:
        Version: 2012-10-17
        Statement:
          - Sid: AWSConfigBucketPermissionsCheck
            Effect: Allow
            Principal:
              Service:
                - config.amazonaws.com
            Action: s3:GetBucketAcl
            Resource:
              - !Sub "arn:aws:s3:::${ConfigBucket}"
          - Sid: AWSConfigBucketDelivery
            Effect: Allow
            Principal:
              Service:
                - config.amazonaws.com
            Action: s3:PutObject
            Resource:
              - !Sub "arn:aws:s3:::${ConfigBucket}/AWSLogs/${AWS::AccountId}/*"

  ConfigRecorderRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2012-10-17
        Statement:
          - Effect: Allow
            Principal:
              Service:
                - config.amazonaws.com
            Action:
              - sts:AssumeRole
      Path: /
      ManagedPolicyArns:
        - arn:aws:iam::aws:policy/service-role/AWSConfigRole

  ConfigRecorder:
    Type: AWS::Config::ConfigurationRecorder
    DependsOn:
      - ConfigRecorderRole
      - ConfigBucketPolicy
    Properties:
      RoleARN: !GetAtt ConfigRecorderRole.Arn
      RecordingGroup:
        AllSupported: True
        IncludeGlobalResourceTypes: !If
          - IncludeGlobalResourceTypes
          - True
          - False

  DeliveryChannel:
    Type: AWS::Config::DeliveryChannel
    DependsOn:
      - ConfigBucketPolicy
    Properties:
      Name: default
      S3BucketName: !Ref ConfigBucket

  S3BucketPublicReadRule:
    Type: AWS::Config::ConfigRule
    DependsOn:
      - ConfigRecorder
    Properties:
      ConfigRuleName: stackset-s3-bucket-public-read-prohibited
      Description: s3-bucket-public-read-prohibited from stackset
      Scope:
        ComplianceResourceTypes:
          - AWS::S3::Bucket
      Source:
        Owner: AWS
        SourceIdentifier: S3_BUCKET_PUBLIC_READ_PROHIBITED

  ConfigAggregator:
    Type: AWS::Config::ConfigurationAggregator
    Condition: CreateAggregator
    Properties:
    ConfigurationAggregatorName: name
      AccountAggregationSources:
        - AccountIds: !Ref SourceAccounts
          AwsRegions: !Ref SourceRegions

  AggregationAuthorization:
    Type: AWS::Config::AggregationAuthorization
    Condition: CreateAuthorization
    Properties:
      AuthorizedAccountId: !Ref AggregatorAccount
      AuthorizedAwsRegion: !Ref AggregatorRegion
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Config

Tag

All content copied from https://docs.aws.amazon.com/.
