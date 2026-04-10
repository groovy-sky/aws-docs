This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream Tag

Metadata that you can assign to a Firehose stream, consisting of a key-value
pair.

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

A unique identifier for the tag. Maximum length: 128 characters. Valid characters:
Unicode letters, digits, white space, \_ . / = + - % @

_Required_: Yes

_Type_: String

_Pattern_: `^(?!aws:)[\p{L}\p{Z}\p{N}_.:\/=+\-@%]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

An optional string, which you can use to describe or define the tag. Maximum length:
256 characters. Valid characters: Unicode letters, digits, white space, \_ . / = + - %
@

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}\p{Z}\p{N}_.:\/=+\-@%]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableCreationConfiguration

VpcConfiguration

All content copied from https://docs.aws.amazon.com/.
