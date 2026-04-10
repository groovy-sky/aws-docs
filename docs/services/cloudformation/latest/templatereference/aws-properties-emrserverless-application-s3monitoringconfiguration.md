This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application S3MonitoringConfiguration

The Amazon S3 configuration for monitoring log publishing. You can configure your jobs
to send log information to Amazon S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionKeyArn" : String,
  "LogUri" : String
}

```

### YAML

```yaml

  EncryptionKeyArn: String
  LogUri: String

```

## Properties

`EncryptionKeyArn`

The KMS key ARN to encrypt the logs published to the given Amazon S3 destination.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z0-9-]*):kms:[a-zA-Z0-9\-]*:(\d{12})?:key\/[a-zA-Z0-9-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LogUri`

The Amazon S3 destination URI for log publishing.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `10280`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PrometheusMonitoringConfiguration

SchedulerConfiguration

All content copied from https://docs.aws.amazon.com/.
