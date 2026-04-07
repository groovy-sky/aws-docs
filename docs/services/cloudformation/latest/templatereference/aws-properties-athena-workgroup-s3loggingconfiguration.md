This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::WorkGroup S3LoggingConfiguration

Configuration settings for delivering logs to Amazon S3 buckets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "KmsKey" : String,
  "LogLocation" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  KmsKey: String
  LogLocation: String

```

## Properties

`Enabled`

Enables S3 log delivery.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKey`

The KMS key ARN to encrypt the logs published to the given Amazon S3 destination.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:kms:([a-z0-9\-]+):\d{12}:key/?[a-zA-Z_0-9+=,.@\-_/]+$|^arn:aws[a-z\-]*:kms:([a-z0-9\-]+):\d{12}:alias/?[a-zA-Z_0-9+=,.@\-_/]+$|^alias/[a-zA-Z0-9/_-]+$|[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogLocation`

The Amazon S3 destination URI for log publishing.

_Required_: No

_Type_: String

_Pattern_: `^s3://[a-z0-9][a-z0-9\-]*[a-z0-9](/.*)?$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResultConfigurationUpdates

Tag
