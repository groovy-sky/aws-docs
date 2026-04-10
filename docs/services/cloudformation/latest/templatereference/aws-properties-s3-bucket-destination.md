This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket Destination

Specifies information about where to publish analysis or configuration results for an
Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketAccountId" : String,
  "BucketArn" : String,
  "Format" : String,
  "Prefix" : String
}

```

### YAML

```yaml

  BucketAccountId: String
  BucketArn: String
  Format: String
  Prefix: String

```

## Properties

`BucketAccountId`

The account ID that owns the destination S3 bucket. If no account ID is provided, the owner is not
validated before exporting data.

###### Note

Although this value is optional, we strongly recommend that you set it to help prevent problems
if the destination bucket ownership changes.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketArn`

The Amazon Resource Name (ARN) of the bucket to which data is exported.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

Specifies the file format used when exporting data to Amazon S3.

_Allowed values_: `CSV` \| `ORC` \|
`Parquet`

_Required_: Yes

_Type_: String

_Allowed values_: `CSV | ORC | Parquet`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The prefix to use when exporting data. The prefix is prepended to all results.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- AWS::S3::Bucket [Examples](../userguide/aws-properties-s3-bucket.md#aws-properties-s3-bucket--examples)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteMarkerReplication

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
