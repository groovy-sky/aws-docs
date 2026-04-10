This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template DataSetReference

Dataset reference.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSetArn" : String,
  "DataSetPlaceholder" : String
}

```

### YAML

```yaml

  DataSetArn: String
  DataSetPlaceholder: String

```

## Properties

`DataSetArn`

Dataset Amazon Resource Name (ARN).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSetPlaceholder`

Dataset placeholder.

_Required_: Yes

_Type_: String

_Pattern_: `\S`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSetConfiguration

DataSetSchema

All content copied from https://docs.aws.amazon.com/.
