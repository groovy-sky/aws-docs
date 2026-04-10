This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TablePinnedFieldOptions

The settings for the pinned columns of a table visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PinnedLeftFields" : [ String, ... ]
}

```

### YAML

```yaml

  PinnedLeftFields:
    - String

```

## Properties

`PinnedLeftFields`

A list of columns to be pinned to the left of a table visual.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `512 | 201`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TablePaginatedReportOptions

TableRowConditionalFormatting

All content copied from https://docs.aws.amazon.com/.
