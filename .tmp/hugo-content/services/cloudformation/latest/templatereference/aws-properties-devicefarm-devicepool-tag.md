This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DeviceFarm::DevicePool Tag

The metadata that you apply to a resource to help you categorize and organize it. Each tag consists of a
key and an optional value, both of which you define. Tag keys can have a maximum character length of 128
characters. Tag values can have a maximum length of 256 characters.

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

One part of a key-value pair that makes up a tag. A `key` is a general label that acts like a
category for more specific tag values.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The optional part of a key-value pair that makes up a tag. A `value` acts as a descriptor in a
tag category (key).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Rule

AWS::DeviceFarm::InstanceProfile

All content copied from https://docs.aws.amazon.com/.
