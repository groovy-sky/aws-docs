This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket S3TablesDestination

The destination information for a V1 S3 Metadata configuration. The destination table bucket must
be in the same Region and AWS account as the general purpose bucket. The specified metadata table name
must be unique within the `aws_s3_metadata` namespace in the destination table bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TableArn" : String,
  "TableBucketArn" : String,
  "TableName" : String,
  "TableNamespace" : String
}

```

### YAML

```yaml

  TableArn: String
  TableBucketArn: String
  TableName: String
  TableNamespace: String

```

## Properties

`TableArn`

The Amazon Resource Name (ARN) for the metadata table in the metadata table configuration. The specified metadata table name must be unique within the `aws_s3_metadata` namespace in the destination table bucket.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableBucketArn`

The Amazon Resource Name (ARN) for the table bucket that's specified as the destination in the
metadata table configuration. The destination table bucket must be in the same Region and AWS account
as the general purpose bucket.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name for the metadata table in your metadata table configuration. The specified metadata table
name must be unique within the `aws_s3_metadata` namespace in the destination table bucket.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableNamespace`

The table bucket namespace for the metadata table in your metadata table configuration. This value is always `aws_s3_metadata`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3KeyFilter

ServerSideEncryptionByDefault

All content copied from https://docs.aws.amazon.com/.
