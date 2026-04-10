This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application S3ContentLocation

The location of an application or a custom artifact.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketARN" : String,
  "FileKey" : String,
  "ObjectVersion" : String
}

```

### YAML

```yaml

  BucketARN: String
  FileKey: String
  ObjectVersion: String

```

## Properties

`BucketARN`

The Amazon Resource Name (ARN) for the S3 bucket containing the application code.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileKey`

The file key for the object containing the application code.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectVersion`

The version of the object containing the application code.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [S3ContentLocation](../../../managed-flink/latest/apiv2/api-s3contentlocation.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3ContentBaseLocation

SqlApplicationConfiguration

All content copied from https://docs.aws.amazon.com/.
