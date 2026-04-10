This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::RealtimeLogConfig KinesisStreamConfig

Contains information about the Amazon Kinesis data stream where you are sending real-time
log data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RoleArn" : String,
  "StreamArn" : String
}

```

### YAML

```yaml

  RoleArn: String
  StreamArn: String

```

## Properties

`RoleArn`

The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM) role that CloudFront can use to
send real-time log data to your Kinesis data stream.

For more information the IAM role, see [Real-time log configuration IAM role](../../../amazoncloudfront/latest/developerguide/real-time-logs.md#understand-real-time-log-config-iam-role) in the
_Amazon CloudFront Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamArn`

The Amazon Resource Name (ARN) of the Kinesis data stream where you are sending
real-time log data.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EndPoint

AWS::CloudFront::ResponseHeadersPolicy

All content copied from https://docs.aws.amazon.com/.
