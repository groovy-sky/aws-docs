This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Macie::CustomDataIdentifier Tag

Specifies a tag (key-value pair) to apply to a custom data identifier. A _tag_ is a label that you can define and associate with AWS
resources, including certain types of Amazon Macie resources. Each tag consists of a _tag key_ and an associated
_tag value_. A _tag key_ is a general label that acts as a category for a more specific
tag value. Each tag key must be unique and it can have only one tag value. A _tag value_ acts as a descriptor for a tag key. Tag keys and
values are case sensitive. They can contain letters, numbers, spaces, or the following symbols: \_ . : / = + - @

For more information, see [Tagging \
Macie resources](../../../macie/latest/user/tagging-resources.md) in the _Amazon Macie User Guide_.

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

The name of the tag key. A tag key can contain up to 128 UTF-8 characters.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The tag value to associate with the specified tag key ( `Key`). A tag value can contain up to 256 UTF-8 characters. A tag value
cannot be null, but it can be an empty string.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Macie::CustomDataIdentifier

AWS::Macie::FindingsFilter

All content copied from https://docs.aws.amazon.com/.
