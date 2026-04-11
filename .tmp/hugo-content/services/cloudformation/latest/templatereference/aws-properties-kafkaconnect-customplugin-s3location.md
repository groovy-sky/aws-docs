This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::CustomPlugin S3Location

The location of an object in Amazon S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketArn" : String,
  "FileKey" : String,
  "ObjectVersion" : String
}

```

### YAML

```yaml

  BucketArn: String
  FileKey: String
  ObjectVersion: String

```

## Properties

`BucketArn`

The Amazon Resource Name (ARN) of an S3 bucket.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FileKey`

The file key for an object in an S3 bucket.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ObjectVersion`

The version of an object in an S3 bucket.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomPluginLocation

Tag

All content copied from https://docs.aws.amazon.com/.
