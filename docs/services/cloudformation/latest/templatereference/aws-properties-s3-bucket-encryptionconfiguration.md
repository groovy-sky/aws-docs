This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket EncryptionConfiguration

Specifies encryption-related information for an Amazon S3 bucket that is a destination for replicated
objects.

###### Note

If you're specifying a customer managed KMS key, we recommend using a fully qualified KMS key
ARN. If you use a KMS key alias instead, then AWS KMS resolves the key within the requester’s account.
This behavior can result in data that's encrypted with a KMS key that belongs to the requester, and
not the bucket owner.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReplicaKmsKeyID" : String
}

```

### YAML

```yaml

  ReplicaKmsKeyID: String

```

## Properties

`ReplicaKmsKeyID`

Specifies the ID (Key ARN or Alias ARN) of the customer managed AWS KMS key stored in AWS Key
Management Service (KMS) for the destination bucket. Amazon S3 uses this key to encrypt replica objects. Amazon S3
only supports symmetric encryption KMS keys. For more information, see [Asymmetric keys in AWS KMS](../../../kms/latest/developerguide/symmetric-asymmetric.md) in the
_AWS Key Management Service Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- AWS::S3::Bucket [Examples](../userguide/aws-properties-s3-bucket.md#aws-properties-s3-bucket--examples)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Destination

EventBridgeConfiguration

All content copied from https://docs.aws.amazon.com/.
