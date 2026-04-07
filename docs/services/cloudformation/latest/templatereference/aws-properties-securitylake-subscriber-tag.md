This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::Subscriber Tag

A _tag_ is a label that you can define and associate with AWS resources, including certain types of Amazon Security Lake resources.
Tags can help you identify, categorize, and manage resources in different ways, such as by owner, environment, or other criteria. You can associate tags with
the following types of Security Lake resources: subscribers, and the data lake configuration for your AWS account in individual AWS Regions.

A resource can have up to 50 tags. Each tag consists of a required _tag key_ and an associated _tag value_. A
_tag key_ is a general label that acts as a category for a more specific tag value. Each tag key must be unique and it can have only one tag
value. A _tag value_ acts as a descriptor for a tag key. Tag keys and values are case sensitive. They can contain letters, numbers, spaces,
or the following symbols: \_ . : / = + @ -

For more information, see [Tagging Amazon Security Lake resources](https://docs.aws.amazon.com/security-lake/latest/userguide/tagging-resources.html) in
the _Amazon Security Lake User Guide_.

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

The name of the tag. This is a general label that acts as a category for a more specific tag value ( `value`).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value that’s associated with the specified tag key ( `key`). This value acts as a descriptor for the tag key. A tag value cannot be
null, but it can be an empty string.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SubscriberIdentity

AWS::SecurityLake::SubscriberNotification
