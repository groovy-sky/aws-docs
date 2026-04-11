This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template RowAlternateColorOptions

Determines the row alternate color options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RowAlternateColors" : [ String, ... ],
  "Status" : String,
  "UsePrimaryBackgroundColor" : String
}

```

### YAML

```yaml

  RowAlternateColors:
    - String
  Status: String
  UsePrimaryBackgroundColor: String

```

## Properties

`RowAlternateColors`

Determines the list of row alternate colors.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Determines the widget status.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsePrimaryBackgroundColor`

The primary background color options for alternate rows.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RollingDateConfiguration

SameSheetTargetVisualConfiguration

All content copied from https://docs.aws.amazon.com/.
