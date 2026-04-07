This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application Tag

An object that defines the tags associated with an application. A
_tag_ is a label that you optionally define and associate with an
application. Tags can help you categorize and manage resources in different ways, such as
by purpose, owner, environment, or other criteria.

Each tag consists of a required _tag key_ and an associated
_tag value_, both of which you define. A tag key is a general label
that acts as a category for a more specific tag value. A tag value acts as a descriptor
within a tag key. A tag key can contain as many as 128 characters. A tag value can contain
as many as 256 characters. The characters can be Unicode letters, digits, white space, or
one of the following symbols: \_ . : / = + -. The following additional restrictions apply to
tags:

- Tag keys and values are case sensitive.

- For each associated resource, each tag key must be unique and it can have only one
value.

- The `aws:` prefix is reserved for use by AWS; you can’t use it in any
tag keys or values that you define. In addition, you can't edit or remove tag keys or
values that use this prefix.

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

One part of a key-value pair that defines a tag. The maximum length of a tag key is 128
characters. The minimum length is 1 character.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The optional part of a key-value pair that defines a tag. The maximum length of a tag
value is 256 characters. The minimum length is 0 characters. If you don't want an
application to have a specific tag value, don't specify a value for this parameter.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SubComponentTypeConfiguration

WindowsEvent
