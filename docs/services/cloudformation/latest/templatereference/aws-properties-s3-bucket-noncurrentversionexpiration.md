This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket NoncurrentVersionExpiration

Specifies when noncurrent object versions expire. Upon expiration, Amazon S3
permanently deletes the noncurrent object versions. You set this lifecycle configuration
action on a bucket that has versioning enabled (or suspended) to request that Amazon S3 delete noncurrent object versions at a specific period in the object's
lifetime. For more information about setting a lifecycle rule configuration, see [AWS::S3::Bucket Rule](../userguide/aws-properties-s3-bucket-lifecycleconfig-rule.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NewerNoncurrentVersions" : Integer,
  "NoncurrentDays" : Integer
}

```

### YAML

```yaml

  NewerNoncurrentVersions: Integer
  NoncurrentDays: Integer

```

## Properties

`NewerNoncurrentVersions`

Specifies how many noncurrent versions Amazon S3 will retain. If there are this
many more recent noncurrent versions, Amazon S3 will take the associated action. For
more information about noncurrent versions, see [Lifecycle configuration\
elements](../../../s3/latest/userguide/intro-lifecycle-rules.md) in the _Amazon S3 User Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoncurrentDays`

Specifies the number of days an object is noncurrent before Amazon S3 can
perform the associated action. For information about the noncurrent days calculations, see
[How\
Amazon S3 Calculates When an Object Became Noncurrent](../../../s3/latest/dev/intro-lifecycle-rules.md#non-current-days-calculations) in the _Amazon S3_
_User Guide_.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricsConfiguration

NoncurrentVersionTransition

All content copied from https://docs.aws.amazon.com/.
