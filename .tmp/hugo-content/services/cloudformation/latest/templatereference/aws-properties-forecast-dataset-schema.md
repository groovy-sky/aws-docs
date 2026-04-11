This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Forecast::Dataset Schema

Defines the fields of a dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : [ AttributesItems, ... ]
}

```

### YAML

```yaml

  Attributes:
    - AttributesItems

```

## Properties

`Attributes`

An array of attributes specifying the name and type of each field in a dataset.

_Required_: No

_Type_: Array of [AttributesItems](aws-properties-forecast-dataset-attributesitems.md)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfig

TagsItems

All content copied from https://docs.aws.amazon.com/.
