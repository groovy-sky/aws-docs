This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Express::DirectoryBucket

The `AWS::S3Express::DirectoryBucket` resource defines an Amazon S3 directory
bucket in the same AWS Region where you create the AWS
CloudFormation stack.

To control how AWS CloudFormation handles the bucket when the stack is
deleted, you can set a deletion policy for your bucket. You can choose to
_retain_ the bucket or to _delete_ the bucket. For
more information, see [DeletionPolicy attribute](../userguide/aws-attribute-deletionpolicy.md).

###### Important

You can only delete empty buckets. Deletion fails for buckets that have
contents.

Permissions

The required permissions for CloudFormation to use are based on the operations
that are performed on the stack.

- Create

- s3express:CreateBucket

- s3express:ListAllMyDirectoryBuckets

- Read

- s3express:ListAllMyDirectoryBuckets

- ec2:DescribeAvailabilityZones

- Delete

- s3express:DeleteBucket

- s3express:ListAllMyDirectoryBuckets

- List

- s3express:ListAllMyDirectoryBuckets

- PutBucketEncryption

- s3express:PutEncryptionConfiguration

- To set a directory bucket default encryption with SSE-KMS, you must
also have the kms:GenerateDataKey and kms:Decrypt permissions in IAM
identity-based policies and AWS KMS key policies for
the target AWS KMS key.

- GetBucketEncryption

- s3express:GetBucketEncryption

- DeleteBucketEncryption

- s3express:PutEncryptionConfiguration

The following operations are related to
`AWS::S3Express::DirectoryBucket`:

- [CreateBucket](../../../s3/latest/api/api-createbucket.md)

- [ListDirectoryBuckets](../../../s3/latest/api/api-listdirectorybuckets.md)

- [DeleteBucket](../../../s3/latest/api/api-deletebucket.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Express::DirectoryBucket",
  "Properties" : {
      "BucketEncryption" : BucketEncryption,
      "BucketName" : String,
      "DataRedundancy" : String,
      "LifecycleConfiguration" : LifecycleConfiguration,
      "LocationName" : String,
      "MetricsConfigurations" : [ MetricsConfiguration, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::S3Express::DirectoryBucket
Properties:
  BucketEncryption:
    BucketEncryption
  BucketName: String
  DataRedundancy: String
  LifecycleConfiguration:
    LifecycleConfiguration
  LocationName: String
  MetricsConfigurations:
    - MetricsConfiguration
  Tags:
    - Tag

```

## Properties

`BucketEncryption`

Specifies default encryption for a bucket using server-side encryption with Amazon S3
managed keys (SSE-S3) or AWS KMS keys (SSE-KMS). For information about
default encryption for directory buckets, see [Setting and monitoring\
default encryption for directory buckets](../../../s3/latest/userguide/s3-express-bucket-encryption.md) in the _Amazon S3 User_
_Guide_.

_Required_: No

_Type_: [BucketEncryption](aws-properties-s3express-directorybucket-bucketencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketName`

A name for the bucket. The bucket name must contain only
lowercase letters, numbers, and hyphens (-).
A directory bucket name must be unique in the chosen Zone (Availability Zone or Local Zone). The bucket name must also follow the format `bucket_base_name--zone_id--x-s3`
(for example, `bucket_base_name--usw2-az1--x-s3`). If you don't specify a name, AWS CloudFormation
generates a unique ID and uses that ID for the bucket name.
For information about bucket naming restrictions,
see [Directory bucket naming rules](../../../s3/latest/userguide/directory-bucket-naming-rules.md)
in the _Amazon S3 User Guide_.

###### Important

If you specify a name, you can't perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you need to
replace the resource, specify a new name.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9][a-z0-9//.//-]*[a-z0-9]$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataRedundancy`

The number of Zone (Availability Zone or Local Zone) that's used for redundancy for the bucket.

_Required_: Yes

_Type_: String

_Allowed values_: `SingleAvailabilityZone | SingleLocalZone`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LifecycleConfiguration`

Container for lifecycle rules. You can add as many as 1000 rules.

For more information see, [Creating and managing a lifecycle configuration for directory buckets](../../../s3/latest/userguide/directory-buckets-objects-lifecycle.md) in the
_Amazon S3 User Guide_.

_Required_: No

_Type_: [LifecycleConfiguration](aws-properties-s3express-directorybucket-lifecycleconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocationName`

The name of the location where the bucket will be created.

For directory buckets, the name of the location is the Zone ID of the Availability Zone (AZ) or Local Zone (LZ) where
the bucket will be created. An example AZ ID value is `usw2-az1`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetricsConfigurations`

Property description not available.

_Required_: No

_Type_: Array of [MetricsConfiguration](aws-properties-s3express-directorybucket-metricsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of tags that you can apply to the S3 directory bucket. Tags are key-value pairs of metadata used to categorize and organize your buckets, track costs, and control access. For more information, see [Using tags with directory buckets](../../../s3/latest/userguide/directory-buckets-tagging.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-s3express-directorybucket-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the bucket name.

Example: `bucket_base_name--usw2-az1--x-s3`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of the specified bucket.

Example:
`arn:aws:s3express:us-west-2:account_id:bucket/bucket_base_name--usw2-az1--x-s3`

`AvailabilityZoneName`

Returns the code for the Availability Zone or the Local Zone where the directory bucket was created.

Example value for an Availability Zone code: _us-east-1f_

###### Note

An Availability Zone code might not represent the same physical location for
different AWS accounts. For more information, see [Availability Zones and Regions](../../../s3/latest/userguide/s3-express-endpoints.md) in the _Amazon S3 User_
_Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::S3Express::BucketPolicy

AbortIncompleteMultipartUpload

All content copied from https://docs.aws.amazon.com/.
