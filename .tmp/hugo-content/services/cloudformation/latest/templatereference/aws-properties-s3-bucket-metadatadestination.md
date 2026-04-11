This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket MetadataDestination

The destination information for the S3 Metadata configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TableBucketArn" : String,
  "TableBucketType" : String,
  "TableNamespace" : String
}

```

### YAML

```yaml

  TableBucketArn: String
  TableBucketType: String
  TableNamespace: String

```

## Properties

`TableBucketArn`

The Amazon Resource Name (ARN) of the table bucket where the metadata configuration is stored.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableBucketType`

The type of the table bucket where the metadata configuration is stored. The `aws`
value indicates an AWS managed table bucket, and the `customer` value indicates a
customer-managed table bucket. V2 metadata configurations are stored in AWS managed table
buckets, and V1 metadata configurations are stored in customer-managed table buckets.

_Required_: Yes

_Type_: String

_Allowed values_: `aws | customer`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableNamespace`

The namespace in the table bucket where the metadata tables for a metadata configuration are
stored.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetadataConfiguration

MetadataTableConfiguration

All content copied from https://docs.aws.amazon.com/.
