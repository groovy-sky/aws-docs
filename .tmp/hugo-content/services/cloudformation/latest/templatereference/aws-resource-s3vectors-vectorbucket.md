This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Vectors::VectorBucket

Defines an Amazon S3 vector bucket in the same AWS Region where you create the AWS CloudFormation stack.

Vector buckets are specialized storage containers designed for storing and managing vector data used in machine learning and AI applications. They provide optimized storage and retrieval capabilities for high-dimensional vector data.

To control how AWS CloudFormation handles the bucket when the stack is deleted, you can set a deletion policy for your bucket. You can choose to _retain_ the bucket or to _delete_ the bucket. For more information, see [DeletionPolicy attribute](../userguide/aws-attribute-deletionpolicy.md).

###### Important

You can only delete empty vector buckets. Deletion fails for buckets that have contents.

Permissions

The required permissions for CloudFormation to use are based on the operations that are performed on the stack.

- Create

- s3vectors:CreateVectorBucket

- s3vectors:GetVectorBucket

- kms:GenerateDataKey (if using KMS encryption)

- Read

- s3vectors:GetVectorBucket

- kms:GenerateDataKey (if using KMS encryption)

- Delete

- s3vectors:DeleteVectorBucket

- s3vectors:GetVectorBucket

- kms:GenerateDataKey (if using KMS encryption)

- List

- s3vectors:ListVectorBuckets

- kms:GenerateDataKey (if using KMS encryption)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Vectors::VectorBucket",
  "Properties" : {
      "EncryptionConfiguration" : EncryptionConfiguration,
      "Tags" : [ Tag, ... ],
      "VectorBucketName" : String
    }
}

```

### YAML

```yaml

Type: AWS::S3Vectors::VectorBucket
Properties:
  EncryptionConfiguration:
    EncryptionConfiguration
  Tags:
    - Tag
  VectorBucketName: String

```

## Properties

`EncryptionConfiguration`

The encryption configuration for the vector bucket.

_Required_: No

_Type_: [EncryptionConfiguration](aws-properties-s3vectors-vectorbucket-encryptionconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-s3vectors-vectorbucket-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VectorBucketName`

A name for the vector bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). The bucket name must be unique in the same AWS account for each AWS Region. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name.

The bucket name must be between 3 and 63 characters long and must not contain uppercase characters or underscores.

###### Important

If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the vector bucket ARN.

Example: `arn:aws:s3vectors:us-east-1:123456789012:bucket/amzn-s3-demo-vector-bucket`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

Returns the date and time when the vector bucket was created.

Example: `2024-12-21T10:30:00Z`

`VectorBucketArn`

Returns the Amazon Resource Name (ARN) of the specified vector bucket.

Example: `arn:aws:s3vectors:us-east-1:123456789012:bucket/amzn-s3-demo-vector-bucket`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
