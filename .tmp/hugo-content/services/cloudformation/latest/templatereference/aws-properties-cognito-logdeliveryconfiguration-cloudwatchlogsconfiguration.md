This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::LogDeliveryConfiguration CloudWatchLogsConfiguration

Configuration for the CloudWatch log group destination of user pool detailed activity
logging, or of user activity log export with advanced security features.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroupArn" : String
}

```

### YAML

```yaml

  LogGroupArn: String

```

## Properties

`LogGroupArn`

The Amazon Resource Name (arn) of a CloudWatch Logs log group where your user pool sends logs.
The log group must not be encrypted with AWS Key Management Service and must be in the same AWS account
as your user pool.

To send logs to log groups with a resource policy of a size greater than 5120
characters, configure a log group with a path that starts with
`/aws/vendedlogs`. For more information, see [Enabling\
logging from certain AWS services](../../../amazoncloudwatch/latest/logs/aws-logs-and-resource-policy.md).

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cognito::LogDeliveryConfiguration

FirehoseConfiguration

All content copied from https://docs.aws.amazon.com/.
