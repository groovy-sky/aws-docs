This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::TableBucket

Creates a table bucket. For more information, see [Creating a table bucket](../../../s3/latest/userguide/s3-tables-buckets-create.md) in the _Amazon Simple Storage Service User Guide_.

Permissions

- You must have the `s3tables:CreateTableBucket` permission to use this operation.

- If you use this operation with the optional `encryptionConfiguration` parameter you must have the `s3tables:PutTableBucketEncryption` permission.

AWS Cloud Development Kit (AWS CDK)

To use S3 Tables AWS CDK constructs, add the `@aws-cdk/aws-s3tables-alpha` dependency with one of the following options:

- NPM: `npm i @aws-cdk/aws-s3tables-alpha`

- Yarn: `yarn add @aws-cdk/aws-s3tables-alpha`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Tables::TableBucket",
  "Properties" : {
      "EncryptionConfiguration" : EncryptionConfiguration,
      "MetricsConfiguration" : MetricsConfiguration,
      "StorageClassConfiguration" : StorageClassConfiguration,
      "TableBucketName" : String,
      "Tags" : [ Tag, ... ],
      "UnreferencedFileRemoval" : UnreferencedFileRemoval
    }
}

```

### YAML

```yaml

Type: AWS::S3Tables::TableBucket
Properties:
  EncryptionConfiguration:
    EncryptionConfiguration
  MetricsConfiguration:
    MetricsConfiguration
  StorageClassConfiguration:
    StorageClassConfiguration
  TableBucketName: String
  Tags:
    - Tag
  UnreferencedFileRemoval:
    UnreferencedFileRemoval

```

## Properties

`EncryptionConfiguration`

Configuration specifying how data should be encrypted. This structure defines the encryption algorithm and optional KMS key to be used for server-side encryption.

_Required_: No

_Type_: [EncryptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3tables-tablebucket-encryptionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricsConfiguration`

Property description not available.

_Required_: No

_Type_: [MetricsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3tables-tablebucket-metricsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageClassConfiguration`

The configuration details for the storage class of tables or table buckets. This allows you to optimize storage costs by selecting the appropriate storage class based on your access patterns and performance requirements.

_Required_: No

_Type_: [StorageClassConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3tables-tablebucket-storageclassconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableBucketName`

The name for the table bucket.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3tables-tablebucket-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnreferencedFileRemoval`

The unreferenced file removal settings for your table bucket. Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots. For more information, see the [_Amazon S3 User Guide_](../../../s3/latest/userguide/s3-table-buckets-maintenance.md).

_Required_: No

_Type_: [UnreferencedFileRemoval](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3tables-tablebucket-unreferencedfileremoval.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the table bucket name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`TableBucketARN`

The Amazon Resource Name (ARN) of the table bucket.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

EncryptionConfiguration
