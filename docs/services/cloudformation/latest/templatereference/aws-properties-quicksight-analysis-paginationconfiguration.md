This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis PaginationConfiguration

The pagination configuration for a table visual or boxplot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PageNumber" : Number,
  "PageSize" : Number
}

```

### YAML

```yaml

  PageNumber:
    Number
  PageSize: Number

```

## Properties

`PageNumber`

Indicates the page number.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PageSize`

Indicates how many items render in one page.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumericSeparatorConfiguration

PanelConfiguration

All content copied from https://docs.aws.amazon.com/.
