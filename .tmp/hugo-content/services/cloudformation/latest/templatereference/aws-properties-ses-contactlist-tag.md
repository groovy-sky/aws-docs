This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ContactList Tag

A tag is a label that you optionally define and associate with a resource, such as a
contact list. Tags can help you categorize and manage resources in different ways, such
as by purpose, owner, environment, or other criteria. A resource can have as many as 50
tags.

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

One part of a key-value pair that defines a tag. The maximum length of a tag key is
128 characters. The minimum length is 1 character.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The optional part of a key-value pair that defines a tag. The maximum length of a tag
value is 256 characters. The minimum length is 0 characters. If you don't want a
resource to have a specific tag value, don't specify a value for this parameter. If you
don't specify a value, Amazon SES sets the value to an empty string.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SES::ContactList

Topic

All content copied from https://docs.aws.amazon.com/.
