This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Vectors::Index

The `AWS::S3Vectors::Index` resource defines a vector index within an Amazon S3 vector bucket. For more information, see [Creating a vector index in a vector bucket](../../../s3/latest/userguide/s3-vectors-create-index.md) in the _Amazon Simple Storage Service User Guide_.

You must specify either `VectorBucketName` or `VectorBucketArn` to identify the bucket that contains the index.

To control how AWS CloudFormation handles the vector index when the stack is deleted, you can set a deletion policy for your index. You can choose to _retain_ the index or to _delete_ the index. For more information, see [DeletionPolicy attribute](../userguide/aws-attribute-deletionpolicy.md).

Permissions

The required permissions for CloudFormation to use are based on the operations that are performed on the stack.

- Create

- s3vectors:CreateIndex

- s3vectors:GetIndex

- Read

- s3vectors:GetIndex

- Delete

- s3vectors:DeleteIndex

- s3vectors:GetIndex

- List

- s3vectors:ListIndexes

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Vectors::Index",
  "Properties" : {
      "DataType" : String,
      "Dimension" : Integer,
      "DistanceMetric" : String,
      "EncryptionConfiguration" : EncryptionConfiguration,
      "IndexName" : String,
      "MetadataConfiguration" : MetadataConfiguration,
      "Tags" : [ Tag, ... ],
      "VectorBucketArn" : String,
      "VectorBucketName" : String
    }
}

```

### YAML

```yaml

Type: AWS::S3Vectors::Index
Properties:
  DataType: String
  Dimension: Integer
  DistanceMetric: String
  EncryptionConfiguration:
    EncryptionConfiguration
  IndexName: String
  MetadataConfiguration:
    MetadataConfiguration
  Tags:
    - Tag
  VectorBucketArn: String
  VectorBucketName: String

```

## Properties

`DataType`

The data type of the vectors to be inserted into the vector index. Currently, only `float32` is supported, which represents 32-bit floating-point numbers.

_Required_: Yes

_Type_: String

_Allowed values_: `float32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Dimension`

The dimensions of the vectors to be inserted into the vector index. This value must be between 1 and 4096, inclusive. All vectors stored in the index must have the same number of dimensions.

The dimension value affects the storage requirements and search performance. Higher dimensions require more storage space and may impact search latency.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DistanceMetric`

The distance metric to be used for similarity search. Valid values are:

- `cosine` \- Measures the cosine of the angle between two vectors.

- `euclidean` \- Measures the straight-line distance between two points in multi-dimensional space. Lower values indicate greater similarity.

_Required_: Yes

_Type_: String

_Allowed values_: `cosine | euclidean`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EncryptionConfiguration`

The encryption configuration for a vector index. By default, if you don't specify, all new vectors in the vector index will use the encryption configuration of the vector bucket.

_Required_: No

_Type_: [EncryptionConfiguration](aws-properties-s3vectors-index-encryptionconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IndexName`

The name of the vector index to create. The index name must be between 3 and 63 characters long and can contain only lowercase letters, numbers, hyphens (-), and dots (.). The index name must be unique within the vector bucket.

If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the index name.

###### Important

If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetadataConfiguration`

The metadata configuration for the vector index.

_Required_: No

_Type_: [MetadataConfiguration](aws-properties-s3vectors-index-metadataconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-s3vectors-index-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VectorBucketArn`

The Amazon Resource Name (ARN) of the vector bucket that contains the vector index.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VectorBucketName`

The name of the vector bucket that contains the vector index.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the index ARN.

Example: `arn:aws:s3vectors:us-east-1:123456789012:bucket/amzn-s3-demo-vector-bucket/index/my-index`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

Returns the date and time when the vector index was created.

Example: `2024-12-21T10:30:00Z`

`IndexArn`

Returns the Amazon Resource Name (ARN) of the specified index.

Example: `arn:aws:s3vectors:us-east-1:123456789012:bucket/amzn-s3-demo-vector-bucket/index/my-index`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon S3 Vectors

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
