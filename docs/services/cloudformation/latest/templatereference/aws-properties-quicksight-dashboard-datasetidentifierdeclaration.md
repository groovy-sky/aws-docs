This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard DataSetIdentifierDeclaration

A data set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSetArn" : String,
  "Identifier" : String
}

```

### YAML

```yaml

  DataSetArn: String
  Identifier: String

```

## Properties

`DataSetArn`

The Amazon Resource Name (ARN) of the data set.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Identifier`

The identifier of the data set, typically the data set's name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataQAEnabledOption

DataSetReference

All content copied from https://docs.aws.amazon.com/.
