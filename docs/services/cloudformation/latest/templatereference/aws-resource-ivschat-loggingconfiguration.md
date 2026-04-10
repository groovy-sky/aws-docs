This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVSChat::LoggingConfiguration

The `AWS::IVSChat::LoggingConfiguration` resource specifies an Amazon IVS logging configuration that allows clients to store and record
sent messages. For more information, see [CreateLoggingConfiguration](../../../../reference/ivs/latest/chatapireference/api-createloggingconfiguration.md) in the _Amazon Interactive Video Service Chat_
_API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IVSChat::LoggingConfiguration",
  "Properties" : {
      "DestinationConfiguration" : DestinationConfiguration,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IVSChat::LoggingConfiguration
Properties:
  DestinationConfiguration:
    DestinationConfiguration
  Name: String
  Tags:
    - Tag

```

## Properties

`DestinationConfiguration`

The DestinationConfiguration is a complex type that contains information about where chat content will be
logged.

_Required_: Yes

_Type_: [DestinationConfiguration](aws-properties-ivschat-loggingconfiguration-destinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Logging-configuration name. The value does not need to be unique.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]*$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-ivschat-loggingconfiguration-tag.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-ivschat-loggingconfiguration-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the logging-configuration ARN. For example:

`{ "Ref": "myLoggingConfiguration" }`

For the Amazon IVS logging configuration
`myLoggingConfiguration`, `Ref` returns the logging-configuration
ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The logging-configuration ARN. For example:
`arn:aws:ivschat:us-west-2:123456789012:logging-configuration/abcdABCDefgh`

`Id`

The logging-configuration ID. For example: `abcdABCDefgh`

`State`

Indicates the current state of the logging configuration. When the state is
`ACTIVE`, the configuration is ready to log a chat session. Valid values:
`CREATING` \| `CREATE_FAILED` \| `DELETING` \|
`DELETE_FAILED` \| `UPDATING` \| `UPDATE_FAILED` \|
`ACTIVE`.

## Examples

### Logging Configuration Template Examples

The following examples specify an Amazon IVS Chat Room that logs
interactions to an S3 bucket.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "Bucket": {
      "Type": "AWS::S3::Bucket"
    },
    "LogGroup": {
      "Type": "AWS::Logs::LogGroup"
    },
    "DeliveryStreamRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": {
            "Effect": "Allow",
            "Principal": {
              "Service": "firehose.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
          }
        }
      }
    },
    "DeliveryStream": {
      "Type": "AWS::KinesisFirehose::DeliveryStream",
      "Properties": {
        "S3DestinationConfiguration": {
          "BucketARN": {
            "Fn::GetAtt": [
              "Bucket",
              "Arn"
            ]
          },
          "RoleARN": {
            "Fn::GetAtt": [
              "DeliveryStreamRole",
              "Arn"
            ]
          }
        }
      }
    },
    "S3LoggingConfiguration": {
      "Type": "AWS::IVSChat::LoggingConfiguration",
      "Properties": {
        "Name": "S3",
        "DestinationConfiguration": {
          "S3": {
            "BucketName": {
              "Ref": "Bucket"
            }
          }
        }
      }
    },
    "CloudWatchLogsLoggingConfiguration": {
      "Type": "AWS::IVSChat::LoggingConfiguration",
      "Properties": {
        "Name": "CloudWatchLogs",
        "DestinationConfiguration": {
          "CloudWatchLogs": {
            "LogGroupName": {
              "Ref": "LogGroup"
            }
          }
        }
      }
    },
    "FirehoseLoggingConfiguration": {
      "Type": "AWS::IVSChat::LoggingConfiguration",
      "Properties": {
        "Name": "Firehose",
        "DestinationConfiguration": {
          "Firehose": {
            "DeliveryStreamName": {
              "Ref": "DeliveryStream"
            }
          }
        }
      }
    },
    "Room": {
      "Type": "AWS::IVSChat::Room",
      "Properties": {
        "Name": "LoggingRoom",
        "LoggingConfigurationIdentifiers": [
          { "Ref": "S3LoggingConfiguration" },
          { "Ref": "CloudWatchLogsLoggingConfiguration" },
          { "Ref": "FirehoseLoggingConfiguration" }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  Bucket:
    Type: AWS::S3::Bucket
  LogGroup:
    Type: AWS::Logs::LogGroup
  DeliveryStreamRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2012-10-17
        Statement:
          Effect: Allow
          Principal:
            Service: firehose.amazonaws.com
          Action: sts:AssumeRole
  DeliveryStream:
    Type: AWS::KinesisFirehose::DeliveryStream
    Properties:
      S3DestinationConfiguration:
        BucketARN: !GetAtt Bucket.Arn
        RoleARN: !GetAtt DeliveryStreamRole.Arn
  S3LoggingConfiguration:
    Type: AWS::IVSChat::LoggingConfiguration
    Properties:
      Name: S3
      DestinationConfiguration:
        S3:
          BucketName: !Ref Bucket
  CloudWatchLogsLoggingConfiguration:
    Type: AWS::IVSChat::LoggingConfiguration
    Properties:
      Name: CloudWatchLogs
      DestinationConfiguration:
        CloudWatchLogs:
          LogGroupName: !Ref LogGroup
  FirehoseLoggingConfiguration:
    Type: AWS::IVSChat::LoggingConfiguration
    Properties:
      Name: Firehose
      DestinationConfiguration:
        Firehose:
          DeliveryStreamName: !Ref DeliveryStream
  Room:
    Type: AWS::IVSChat::Room
    Properties:
      Name: LoggingRoom
      LoggingConfigurationIdentifiers:
        - !Ref S3LoggingConfiguration
        - !Ref CloudWatchLogsLoggingConfiguration
        - !Ref FirehoseLoggingConfiguration
```

## See also

- [Getting\
Started with Amazon Interactive Video Service](../../../ivs/latest/userguide/getting-started.md)

- [LoggingConfigurationSummary](../../../../reference/ivs/latest/chatapireference/api-loggingconfigurationsummary.md) API data type

- [DestinationConfiguration](../../../../reference/ivs/latest/chatapireference/api-destinationconfiguration.md) API data type

- [CloudWatchLogsDestinationConfiguration](../../../../reference/ivs/latest/chatapireference/api-cloudwatchlogsdestinationconfiguration.md) API data type

- [FirehoseDestinationConfiguration](../../../../reference/ivs/latest/chatapireference/api-firehosedestinationconfiguration.md) API data type

- [S3DestinationConfiguration](../../../../reference/ivs/latest/chatapireference/api-s3destinationconfiguration.md) API data type

- [CreateLoggingConfiguration](../../../../reference/ivs/latest/chatapireference/api-createloggingconfiguration.md) API endpoint

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon IVS Chat

CloudWatchLogsDestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
