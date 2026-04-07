This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket NotificationFilter

Specifies object key name filtering rules. For information about key name filtering, see [Configuring\
event notifications using object key name filtering](../../../s3/latest/userguide/notification-how-to-filtering.md) in the
_Amazon S3 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Key" : S3KeyFilter
}

```

### YAML

```yaml

  S3Key:
    S3KeyFilter

```

## Properties

`S3Key`

A container for object key name prefix and suffix filtering rules.

_Required_: Yes

_Type_: [S3KeyFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-s3keyfilter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NotificationConfiguration

ObjectLockConfiguration
