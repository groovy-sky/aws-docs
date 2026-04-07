This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket ObjectLockRule

Specifies the Object Lock rule for the specified object. Enable the this rule when you
apply `ObjectLockConfiguration` to a bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultRetention" : DefaultRetention
}

```

### YAML

```yaml

  DefaultRetention:
    DefaultRetention

```

## Properties

`DefaultRetention`

The default Object Lock retention mode and period that you want to apply to new objects
placed in the specified bucket. If Object Lock is turned on, bucket settings require both
`Mode` and a period of either `Days` or `Years`. You cannot
specify `Days` and `Years` at the same time. For more information about
allowable values for mode and period, see [DefaultRetention](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html).

_Required_: Conditional

_Type_: [DefaultRetention](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-defaultretention.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ObjectLockConfiguration

OwnershipControls
