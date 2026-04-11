This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kinesis::StreamConsumer Tag

Metadata assigned to the stream or consumer, consisting of a key-value pair.

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

A unique identifier for the tag. The maximum length for a tag key is 128 characters.

A tag key can only contain the following:

- Unicode letters

- Digits

- White space

- One or more of these symbols: `_`, `.`, `/`, `=`, `+`, `-`, `%`, `@`

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

An optional string, typically used to describe or define the tag. The maximum length for a tag value is 256 characters.

A tag value can only contain the following:

- Unicode letters

- Digits

- White space

- One or more of these symbols: `_`, `.`, `/`, `=`, `+`, `-`, `%`, `@`

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Kinesis::StreamConsumer

Next

All content copied from https://docs.aws.amazon.com/.
