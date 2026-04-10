This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ColumnHierarchy

The option that determines the hierarchy of the fields for a visual element.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DateTimeHierarchy" : DateTimeHierarchy,
  "ExplicitHierarchy" : ExplicitHierarchy,
  "PredefinedHierarchy" : PredefinedHierarchy
}

```

### YAML

```yaml

  DateTimeHierarchy:
    DateTimeHierarchy
  ExplicitHierarchy:
    ExplicitHierarchy
  PredefinedHierarchy:
    PredefinedHierarchy

```

## Properties

`DateTimeHierarchy`

The option that determines the hierarchy of any `DateTime` fields.

_Required_: No

_Type_: [DateTimeHierarchy](aws-properties-quicksight-analysis-datetimehierarchy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExplicitHierarchy`

The option that determines the hierarchy of the fields that are built within a visual's field wells. These fields can't be duplicated to other visuals.

_Required_: No

_Type_: [ExplicitHierarchy](aws-properties-quicksight-analysis-explicithierarchy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredefinedHierarchy`

The option that determines the hierarchy of the fields that are defined during data preparation. These fields are available to use in any analysis that uses the data source.

_Required_: No

_Type_: [PredefinedHierarchy](aws-properties-quicksight-analysis-predefinedhierarchy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColumnConfiguration

ColumnIdentifier

All content copied from https://docs.aws.amazon.com/.
