This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Files::FileSystem

The `AWS::S3Files::FileSystem` resource specifies an Amazon S3 Files
file system scoped to a bucket or prefix within a bucket, enabling file system access
to S3 data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Files::FileSystem",
  "Properties" : {
      "AcceptBucketWarning" : Boolean,
      "Bucket" : String,
      "ClientToken" : String,
      "KmsKeyId" : String,
      "Prefix" : String,
      "RoleArn" : String,
      "SynchronizationConfiguration" : SynchronizationConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::S3Files::FileSystem
Properties:
  AcceptBucketWarning: Boolean
  Bucket: String
  ClientToken: String
  KmsKeyId: String
  Prefix: String
  RoleArn: String
  SynchronizationConfiguration:
    SynchronizationConfiguration
  Tags:
    - Tag

```

## Properties

`AcceptBucketWarning`

A boolean that indicates you have read and accept the warning about the S3 bucket
being used for the file system. Set this to `true` to acknowledge the
warning.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Bucket`

The Amazon Resource Name (ARN) of the S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws[a-zA-Z0-9-]*:s3:::.+)$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClientToken`

A string of up to 64 ASCII characters that Amazon S3 Files uses to ensure
idempotent creation.

_Required_: No

_Type_: String

_Pattern_: `^(.+)$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The ID of the AWS KMS key used to encrypt the file system.

_Required_: No

_Type_: String

_Pattern_: `^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|mrk-[0-9a-f]{32}|alias/[a-zA-Z0-9/_-]+|(arn:aws[-a-z]*:kms:[a-z0-9-]+:\d{12}:((key/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})|(key/mrk-[0-9a-f]{32})|(alias/[a-zA-Z0-9/_-]+))))$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Prefix`

The S3 key prefix that scopes the file system. When specified, the file system
provides access only to objects under this prefix in the bucket.

_Required_: No

_Type_: String

_Pattern_: `^(|.*/)$`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role used for S3 access.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*)?:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SynchronizationConfiguration`

The synchronization configuration for the file system, including import data rules
and expiration data rules that control how data is synchronized between S3 and the
file system.

_Required_: No

_Type_: [SynchronizationConfiguration](aws-properties-s3files-filesystem-synchronizationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-s3files-filesystem-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the file system.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The time when the file system was created.

`FileSystemArn`

The Amazon Resource Name (ARN) of the file system.

`FileSystemId`

The ID of the file system.

`OwnerId`

The AWS account ID of the file system owner.

`Status`

The current status of the file system.

`StatusMessage`

Additional information about the file system status.

`SynchronizationConfiguration.LatestVersionNumber`

The latest version number of the synchronization configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RootDirectory

ExpirationDataRule

All content copied from https://docs.aws.amazon.com/.
