This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket ServerSideEncryptionRule

Specifies the default server-side encryption configuration.

###### Note

- **General purpose buckets** \- If you're specifying a customer
managed KMS key, we recommend using a fully qualified KMS key ARN. If you use a KMS key
alias instead, then AWS KMS resolves the key within the requester’s account. This behavior can
result in data that's encrypted with a KMS key that belongs to the requester, and not the bucket
owner.

- **Directory buckets** -
When you specify an [AWS KMS customer managed key](../../../kms/latest/developerguide/concepts.md#customer-cmk) for encryption in your directory bucket, only use the key ID or key ARN. The key alias format of the KMS key isn't supported.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlockedEncryptionTypes" : BlockedEncryptionTypes,
  "BucketKeyEnabled" : Boolean,
  "ServerSideEncryptionByDefault" : ServerSideEncryptionByDefault
}

```

### YAML

```yaml

  BlockedEncryptionTypes:
    BlockedEncryptionTypes
  BucketKeyEnabled: Boolean
  ServerSideEncryptionByDefault:
    ServerSideEncryptionByDefault

```

## Properties

`BlockedEncryptionTypes`

A bucket-level setting for Amazon S3 general purpose buckets used to prevent the upload of new objects encrypted with the specified server-side encryption type. For example, blocking an encryption type will block `PutObject`, `CopyObject`, `PostObject`, multipart upload, and replication requests to the bucket for objects with the specified encryption type. However, you can continue to read and list any pre-existing objects already encrypted with the specified encryption type. For more information, see [Blocking or unblocking SSE-C for a general purpose bucket](../../../s3/latest/userguide/blocking-unblocking-s3-c-encryption-gpb.md).

###### Note

Currently, this parameter only supports blocking or unblocking server-side encryption with customer-provided keys (SSE-C). For more information about SSE-C, see [Using server-side encryption with customer-provided keys (SSE-C)](../../../s3/latest/userguide/serversideencryptioncustomerkeys.md).

_Required_: No

_Type_: [BlockedEncryptionTypes](aws-properties-s3-bucket-blockedencryptiontypes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketKeyEnabled`

Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using
KMS (SSE-KMS) for new objects in the bucket. Existing objects are not affected. Setting the
`BucketKeyEnabled` element to `true` causes Amazon S3 to use an S3
Bucket Key. By default, S3 Bucket Key is not enabled.

For more information, see [Amazon S3 Bucket Keys](../../../s3/latest/dev/bucket-key.md) in the
_Amazon S3 User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerSideEncryptionByDefault`

Specifies the default server-side encryption to apply to new objects in the bucket. If a PUT Object
request doesn't specify any server-side encryption, this default encryption will be applied.

_Required_: No

_Type_: [ServerSideEncryptionByDefault](aws-properties-s3-bucket-serversideencryptionbydefault.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Create a bucket with default encryption](#aws-properties-s3-bucket-serversideencryptionrule--examples--Create_a_bucket_with_default_encryption)

- [Create a bucket using KMS server-side encryption with an S3 Bucket Key](#aws-properties-s3-bucket-serversideencryptionrule--examples--Create_a_bucket_using_KMS_server-side_encryption_with_an_S3_Bucket_Key)

### Create a bucket with default encryption

The following example creates a bucket with server-side bucket encryption configured.
This example uses encryption with KMS keys (SSE-KMS). You can use
dual-layer server-side encryption with AWS KMS keys (DSSE-KMS) by
specifying `aws:kms:dsse` for `SSEAlgorithm`. You can also use server-side encryption
with S3-managed keys (SSE-S3) by modifying the [Amazon S3 Bucket ServerSideEncryptionByDefault](../userguide/aws-properties-s3-bucket-serversideencryptionbydefault.md) property to specify
`AES256` for `SSEAlgorithm`. For more information, see [Using SSE-S3](../../../s3/latest/userguide/usingserversideencryption.md) in the _Amazon S3 User Guide_.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "S3 bucket with default encryption",
    "Resources": {
        "EncryptedS3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "BucketName": {
                    "Fn::Sub": "encryptedbucket-${AWS::Region}-${AWS::AccountId}"
                },
                "BucketEncryption": {
                    "ServerSideEncryptionConfiguration": [
                        {
                            "ServerSideEncryptionByDefault": {
                                "SSEAlgorithm": "aws:kms",
                                "KMSMasterKeyID": "KMS-KEY-ARN"
                            }
                        }
                    ]
                }
            },
            "DeletionPolicy": "Delete"
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: S3 bucket with default encryption
Resources:
  EncryptedS3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      BucketName: !Sub 'encryptedbucket-${AWS::Region}-${AWS::AccountId}'
      BucketEncryption:
        ServerSideEncryptionConfiguration:
          - ServerSideEncryptionByDefault:
              SSEAlgorithm: 'aws:kms'
              KMSMasterKeyID: KMS-KEY-ARN
    DeletionPolicy: Delete
```

### Create a bucket using KMS server-side encryption with an S3 Bucket Key

The following example creates a bucket that specifies default encryption using AWS KMS server-side encryption with an S3 Bucket Key. The example uses a
customer managed AWS KMS key.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "S3 bucket with default encryption using SSE-KMS with an S3 Bucket Key",
    "Resources": {
        "EncryptedS3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "BucketName": {
                    "Fn::Sub": "encryptedbucket-${AWS::Region}-${AWS::AccountId}"
                },
                "BucketEncryption": {
                    "ServerSideEncryptionConfiguration": [
                        {
                            "ServerSideEncryptionByDefault": {
                                "SSEAlgorithm": "aws:kms",
                                "KMSMasterKeyID": "KMS-KEY-ARN"
                            },
                            "BucketKeyEnabled": true
                        }
                    ]
                }
            },
            "DeletionPolicy": "Delete"
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: S3 bucket with default encryption using SSE-KMS with an S3 Bucket Key
Resources:
  EncryptedS3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      BucketName: !Sub 'encryptedbucket-${AWS::Region}-${AWS::AccountId}'
      BucketEncryption:
        ServerSideEncryptionConfiguration:
          - ServerSideEncryptionByDefault:
              SSEAlgorithm: 'aws:kms'
              KMSMasterKeyID: KMS-KEY-ARN
            BucketKeyEnabled: true
    DeletionPolicy: Delete
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServerSideEncryptionByDefault

SourceSelectionCriteria

All content copied from https://docs.aws.amazon.com/.
