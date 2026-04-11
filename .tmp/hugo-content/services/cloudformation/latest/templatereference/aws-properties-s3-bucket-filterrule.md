This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket FilterRule

Specifies the Amazon S3 object key name to filter on. An object key name is the name assigned to an
object in your Amazon S3 bucket. You specify whether to filter on the suffix or prefix of the object key
name. A prefix is a specific string of characters at the beginning of an object key name, which you can
use to organize objects. For example, you can start the key names of related objects with a prefix, such
as `2023-` or `engineering/`. Then, you can use `FilterRule` to find
objects in a bucket with key names that have the same prefix. A suffix is similar to a prefix, but it is
at the end of the object key name instead of at the beginning.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

The object key name prefix or suffix identifying one or more objects to which the filtering rule
applies. The maximum length is 1,024 characters. Overlapping prefixes and suffixes are not supported.
For more information, see [Configuring Event Notifications](../../../s3/latest/dev/notificationhowto.md) in the _Amazon S3 User Guide_.

_Required_: Yes

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value that the filter searches for in object key names.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventBridgeConfiguration

IntelligentTieringConfiguration

All content copied from https://docs.aws.amazon.com/.
