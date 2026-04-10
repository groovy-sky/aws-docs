This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::Index ValueImportanceItem

Specifies a key-value pair of the search boost value for a document when the key is
part of the metadata of a document.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : Integer
}

```

### YAML

```yaml

  Key: String
  Value: Integer

```

## Properties

`Key`

The document metadata value used for the search boost.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The boost value for a document when the key is part of the metadata of a
document.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserTokenConfiguration

Next

All content copied from https://docs.aws.amazon.com/.
