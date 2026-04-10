This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard StaticFileS3SourceOptions

The structure that contains the Amazon S3 location to download the static file from.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "ObjectKey" : String,
  "Region" : String
}

```

### YAML

```yaml

  BucketName: String
  ObjectKey: String
  Region: String

```

## Properties

`BucketName`

The name of the Amazon S3 bucket.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectKey`

The identifier of the static file in the Amazon S3 bucket.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The Region of the Amazon S3 account that contains the bucket.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StaticFile

StaticFileSource

All content copied from https://docs.aws.amazon.com/.
