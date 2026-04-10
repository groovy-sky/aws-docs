This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application S3ContentBaseLocation

The base location of the Amazon Data Analytics application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BasePath" : String,
  "BucketARN" : String
}

```

### YAML

```yaml

  BasePath: String
  BucketARN: String

```

## Properties

`BasePath`

The base path for the S3 bucket.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9/!-_.*'()]+$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketARN`

The Amazon Resource Name (ARN) of the S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RunConfiguration

S3ContentLocation

All content copied from https://docs.aws.amazon.com/.
