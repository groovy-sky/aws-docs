This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RAM::ResourceShare Tag

A structure containing a tag. A tag is metadata that you can attach to your resources
to help organize and categorize them. You can also use them to help you secure your
resources. For more information, see [Controlling access to AWS resources\
using tags](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html).

For more information about tags, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the
_AWS General Reference Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The key, or name, attached to the tag. Every tag must have a key. Key names are case
sensitive.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The string value attached to the tag. The value can be an empty string. Key values are
case sensitive.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::RAM::ResourceShare

Next
