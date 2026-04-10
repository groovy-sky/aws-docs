This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::Table

Creates a new table associated with the given namespace in a table bucket. For more information, see [Creating an Amazon S3 table](../../../s3/latest/userguide/s3-tables-create.md) in the _Amazon Simple Storage Service User Guide_.

Permissions

- You must have the `s3tables:CreateTable` permission to use this operation.

- If you use this operation with the optional `metadata` request parameter you must have the `s3tables:PutTableData` permission.

- If you use this operation with the optional `encryptionConfiguration` request parameter you must have the `s3tables:PutTableEncryption` permission.

###### Note

Additionally, If you choose SSE-KMS encryption you must grant the S3 Tables maintenance principal access to your KMS key. For more information, see [Permissions requirements for S3 Tables SSE-KMS encryption](../../../s3/latest/userguide/s3-tables-kms-permissions.md).

AWS Cloud Development Kit (AWS CDK)

To use S3 Tables AWS CDK constructs, add the `@aws-cdk/aws-s3tables-alpha` dependency with one of the following options:

- NPM: `npm i @aws-cdk/aws-s3tables-alpha`

- Yarn: `yarn add @aws-cdk/aws-s3tables-alpha`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Tables::Table",
  "Properties" : {
      "Compaction" : Compaction,
      "IcebergMetadata" : IcebergMetadata,
      "Namespace" : String,
      "OpenTableFormat" : String,
      "SnapshotManagement" : SnapshotManagement,
      "StorageClassConfiguration" : StorageClassConfiguration,
      "TableBucketARN" : String,
      "TableName" : String,
      "Tags" : [ Tag, ... ],
      "WithoutMetadata" : String
    }
}

```

### YAML

```yaml

Type: AWS::S3Tables::Table
Properties:
  Compaction:
    Compaction
  IcebergMetadata:
    IcebergMetadata
  Namespace: String
  OpenTableFormat: String
  SnapshotManagement:
    SnapshotManagement
  StorageClassConfiguration:
    StorageClassConfiguration
  TableBucketARN: String
  TableName: String
  Tags:
    - Tag
  WithoutMetadata: String

```

## Properties

`Compaction`

Contains details about the compaction settings for an Iceberg table.

_Required_: No

_Type_: [Compaction](aws-properties-s3tables-table-compaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IcebergMetadata`

Contains details about the metadata for an Iceberg table.

_Required_: No

_Type_: [IcebergMetadata](aws-properties-s3tables-table-icebergmetadata.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Namespace`

The name of the namespace.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenTableFormat`

The format of the table.

_Required_: Yes

_Type_: String

_Allowed values_: `ICEBERG`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotManagement`

Contains details about the Iceberg snapshot management settings for the table.

_Required_: No

_Type_: [SnapshotManagement](aws-properties-s3tables-table-snapshotmanagement.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageClassConfiguration`

The configuration details for the storage class of tables or table buckets. This allows you to optimize storage costs by selecting the appropriate storage class based on your access patterns and performance requirements.

_Required_: No

_Type_: [StorageClassConfiguration](aws-properties-s3tables-table-storageclassconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableBucketARN`

The Amazon Resource Name (ARN) of the table bucket to create the table in.

_Required_: Yes

_Type_: String

_Pattern_: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableName`

The name for the table.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9a-z_]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-s3tables-table-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WithoutMetadata`

Indicates that you don't want to specify a schema for the table. This property is mutually exclusive to `IcebergMetadata`, and its only possible value is `Yes`.

_Required_: No

_Type_: String

_Allowed values_: `Yes`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the table name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`TableARN`

The Amazon Resource Name (ARN) of the table.

`VersionToken`

The version token of the table.

`WarehouseLocation`

The warehouse location of the table.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::S3Tables::Namespace

Compaction

All content copied from https://docs.aws.amazon.com/.
