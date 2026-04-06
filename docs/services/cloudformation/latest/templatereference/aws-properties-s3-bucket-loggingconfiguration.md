This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket LoggingConfiguration

Describes where logs are stored and the prefix that Amazon S3 assigns to all log object
keys for a bucket. For examples and more information, see [PUT Bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html) in the
_Amazon S3 API Reference_.

###### Note

To successfully complete the `AWS::S3::Bucket LoggingConfiguration` request,
you must have `s3:PutObject` and `s3:PutObjectAcl` in your IAM
permissions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationBucketName" : String,
  "LogFilePrefix" : String,
  "TargetObjectKeyFormat" : TargetObjectKeyFormat
}

```

### YAML

```yaml

  DestinationBucketName: String
  LogFilePrefix: String
  TargetObjectKeyFormat:
    TargetObjectKeyFormat

```

## Properties

`DestinationBucketName`

The name of the bucket where Amazon S3 should store server access log files. You can store
log files in any bucket that you own. By default, logs are stored in the bucket where the
`LoggingConfiguration` property is defined.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogFilePrefix`

A prefix for all log object keys. If you store log files from multiple Amazon S3 buckets in a single
bucket, you can use a prefix to distinguish which log files came from which bucket.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetObjectKeyFormat`

Amazon S3 key format for log objects. Only one format, either PartitionedPrefix or SimplePrefix, is allowed.

_Required_: No

_Type_: [TargetObjectKeyFormat](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-targetobjectkeyformat.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Log access requests for a specific S3 bucket](#aws-properties-s3-bucket-loggingconfiguration--examples--Log_access_requests_for_a_specific_S3_bucket)

- [Setting up logging configurations with log prefixes based on event time](#aws-properties-s3-bucket-loggingconfiguration--examples--Setting_up_logging_configurations_with_log_prefixes_based_on_event_time)

- [Setting up logging configurations with log prefixes based on delivery time](#aws-properties-s3-bucket-loggingconfiguration--examples--Setting_up_logging_configurations_with_log_prefixes_based_on_delivery_time)

- [Setting up logging configurations with a simple prefix](#aws-properties-s3-bucket-loggingconfiguration--examples--Setting_up_logging_configurations_with_a_simple_prefix)

### Log access requests for a specific S3 bucket

The following example template creates two S3 buckets. The `LoggingBucket`
bucket store the logs from the `S3Bucket` bucket. To receive logs from the
`S3Bucket` bucket, the logging bucket requires log delivery write
permissions.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "LoggingConfiguration": {
                    "DestinationBucketName": {
                        "Ref": "LoggingBucket"
                    },
                    "LogFilePrefix": "testing-logs"
                }
            }
        },
        "LoggingBucket": {
            "Type": "AWS::S3::Bucket"
        },
        "S3BucketPolicy": {
            "Type": "AWS::S3::BucketPolicy",
            "Properties": {
                "Bucket": {
                    "Ref": "LoggingBucket"
                },
                "PolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Action": [
                                "s3:PutObject"
                            ],
                            "Effect": "Allow",
                            "Principal": {
                                "Service": "logging.s3.amazonaws.com"
                            },
                            "Resource": {
                                "Fn::Join": [
                                    "",
                                    [
                                        "arn:aws:s3:::",
                                        {
                                            "Ref": "LoggingBucket"
                                        },
                                        "/*"
                                    ]
                                ]
                            },
                            "Condition": {
                                "ArnLike": {
                                    "aws:SourceArn": {
                                        "Fn::GetAtt": [
                                            "S3Bucket",
                                            "Arn"
                                        ]
                                    }
                                },
                                "StringEquals": {
                                    "aws:SourceAccount": {
                                        "Fn::Sub": "${AWS::AccountId}"
                                    }
                                }
                            }
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "BucketName": {
            "Value": {
                "Ref": "S3Bucket"
            },
            "Description": "Name of the sample Amazon S3 bucket with a logging configuration."
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  S3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      LoggingConfiguration:
        DestinationBucketName: !Ref LoggingBucket
        LogFilePrefix: testing-logs
  LoggingBucket:
    Type: 'AWS::S3::Bucket'
  S3BucketPolicy:
    Type: 'AWS::S3::BucketPolicy'
    Properties:
      Bucket: !Ref LoggingBucket
      PolicyDocument:
        Version: 2012-10-17
        Statement:
          - Action:
              - 's3:PutObject'
            Effect: Allow
            Principal:
              Service: logging.s3.amazonaws.com
            Resource: !Join
              - ''
              - - 'arn:aws:s3:::'
                - !Ref LoggingBucket
                - /*
            Condition:
              ArnLike:
                'aws:SourceArn': !GetAtt
                  - S3Bucket
                  - Arn
              StringEquals:
                'aws:SourceAccount': !Sub '${AWS::AccountId}'
Outputs:
  BucketName:
    Value: !Ref S3Bucket
    Description: Name of the sample Amazon S3 bucket with a logging configuration.

```

### Setting up logging configurations with log prefixes based on event time

The following example template configures the `DOC-EXAMPLE-BUCKET` destination bucket with a `logs/` prefix and event time log delivery.

#### JSON

```json

        "LoggingConfiguration": {
            "DestinationBucketName": "DOC-EXAMPLE-BUCKET",
            "LogFilePrefix": "logs/",
            "TargetObjectKeyFormat": {
                "PartitionedPrefix": {
                    "PartitionDateSource": "EventTime"
                }
            }
        }

```

#### YAML

```yaml

        LoggingConfiguration:
          DestinationBucketName: "DOC-EXAMPLE-BUCKET"
          LogFilePrefix: logs/
          TargetObjectKeyFormat:
            PartitionedPrefix:
              PartitionDateSource: EventTime

```

### Setting up logging configurations with log prefixes based on delivery time

The following example template configures the `DOC-EXAMPLE-BUCKET` destination bucket with a `logs/` prefix and delivery time log delivery.

#### JSON

```json

        "LoggingConfiguration": {
            "DestinationBucketName": "DOC-EXAMPLE-BUCKET",
            "LogFilePrefix": "logs/",
            "TargetObjectKeyFormat": {
                "PartitionedPrefix": {
                    "PartitionDateSource": "DeliveryTime"
                }
            }
        }

```

#### YAML

```yaml

        LoggingConfiguration:
          DestinationBucketName: "DOC-EXAMPLE-BUCKET"
          LogFilePrefix: logs/
          TargetObjectKeyFormat:
            PartitionedPrefix:
              PartitionDateSource: DeliveryTime

```

### Setting up logging configurations with a simple prefix

The following example template configures the `DOC-EXAMPLE-BUCKET` destination bucket with a `logs/` prefix and simple prefix delivery.

#### JSON

```json

        "LoggingConfiguration": {
            "DestinationBucketName": "DOC-EXAMPLE-BUCKET",
            "LogFilePrefix": "logs/",
            "TargetObjectKeyFormat": {
                "SimplePrefix": {}
            }
        }

```

#### YAML

```yaml

        LoggingConfiguration:
          DestinationBucketName: "DOC-EXAMPLE-BUCKET"
          LogFilePrefix: logs/
          TargetObjectKeyFormat:
            SimplePrefix: {}

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LifecycleConfiguration

MetadataConfiguration
