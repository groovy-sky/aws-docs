This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DynamicDefaultValue

Defines different defaults to the users or groups based on mapping.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValueColumn" : ColumnIdentifier,
  "GroupNameColumn" : ColumnIdentifier,
  "UserNameColumn" : ColumnIdentifier
}

```

### YAML

```yaml

  DefaultValueColumn:
    ColumnIdentifier
  GroupNameColumn:
    ColumnIdentifier
  UserNameColumn:
    ColumnIdentifier

```

## Properties

`DefaultValueColumn`

The column that contains the default value of each user or group.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-analysis-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupNameColumn`

The column that contains the group name.

_Required_: No

_Type_: [ColumnIdentifier](aws-properties-quicksight-analysis-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserNameColumn`

The column that contains the username.

_Required_: No

_Type_: [ColumnIdentifier](aws-properties-quicksight-analysis-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DropDownControlDisplayOptions

EmptyVisual

All content copied from https://docs.aws.amazon.com/.
