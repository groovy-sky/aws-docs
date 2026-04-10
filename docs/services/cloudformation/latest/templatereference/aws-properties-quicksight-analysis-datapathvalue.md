This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DataPathValue

The data path that needs to be sorted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataPathType" : DataPathType,
  "FieldId" : String,
  "FieldValue" : String
}

```

### YAML

```yaml

  DataPathType:
    DataPathType
  FieldId: String
  FieldValue: String

```

## Properties

`DataPathType`

The type configuration of the field.

_Required_: No

_Type_: [DataPathType](aws-properties-quicksight-analysis-datapathtype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldId`

The field ID of the field that needs to be sorted.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldValue`

The actual value of the field that needs to be sorted.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataPathType

DataSetIdentifierDeclaration

All content copied from https://docs.aws.amazon.com/.
