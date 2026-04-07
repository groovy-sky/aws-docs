This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket NoncurrentVersionTransition

Container for the transition rule that describes when noncurrent objects transition to the
`STANDARD_IA`, `ONEZONE_IA`, `INTELLIGENT_TIERING`,
`GLACIER_IR`, `GLACIER`, or `DEEP_ARCHIVE` storage class.
If your bucket is versioning-enabled (or versioning is suspended), you can set this action to
request that Amazon S3 transition noncurrent object versions to the `STANDARD_IA`,
`ONEZONE_IA`, `INTELLIGENT_TIERING`, `GLACIER_IR`,
`GLACIER`, or `DEEP_ARCHIVE` storage class at a specific period in the
object's lifetime. If you specify this property, don't specify the
`NoncurrentVersionTransitions` property.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NewerNoncurrentVersions" : Integer,
  "StorageClass" : String,
  "TransitionInDays" : Integer
}

```

### YAML

```yaml

  NewerNoncurrentVersions: Integer
  StorageClass: String
  TransitionInDays: Integer

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

`StorageClass`

The class of storage used to store the object.

_Required_: Yes

_Type_: String

_Allowed values_: `DEEP_ARCHIVE | GLACIER | Glacier | GLACIER_IR | INTELLIGENT_TIERING | ONEZONE_IA | STANDARD_IA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitionInDays`

Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
For information about the noncurrent days calculations, see [How Amazon S3 Calculates\
How Long an Object Has Been Noncurrent](https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#non-current-days-calculations) in the _Amazon S3 User Guide_.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NoncurrentVersionExpiration

NotificationConfiguration
