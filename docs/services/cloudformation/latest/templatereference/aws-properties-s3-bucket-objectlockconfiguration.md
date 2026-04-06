This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket ObjectLockConfiguration

Places an Object Lock configuration on the specified bucket. The rule specified in the
Object Lock configuration will be applied by default to every new object placed in the
specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ObjectLockEnabled" : String,
  "Rule" : ObjectLockRule
}

```

### YAML

```yaml

  ObjectLockEnabled: String
  Rule:
    ObjectLockRule

```

## Properties

`ObjectLockEnabled`

Indicates whether this bucket has an Object Lock configuration enabled. Enable
`ObjectLockEnabled` when you apply `ObjectLockConfiguration` to a bucket.

_Required_: No

_Type_: String

_Allowed values_: `Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rule`

Specifies the Object Lock rule for the specified object. Enable this rule when you apply
`ObjectLockConfiguration` to a bucket. If Object Lock is turned on, bucket
settings require both `Mode` and a period of either `Days` or
`Years`. You cannot specify `Days` and `Years` at the same
time. For more information, see [ObjectLockRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-objectlockrule.html) and [DefaultRetention](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html).

_Required_: Conditional

_Type_: [ObjectLockRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-objectlockrule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- AWS::S3::Bucket [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#aws-properties-s3-bucket--examples)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NotificationFilter

ObjectLockRule
