This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis SectionStyle

The options that style a section.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Height" : String,
  "Padding" : Spacing
}

```

### YAML

```yaml

  Height: String
  Padding:
    Spacing

```

## Properties

`Height`

The height of a section.

Heights can only be defined for header and footer sections. The default height margin is 0.5 inches.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Padding`

The spacing between section content and its top, bottom, left, and right edges.

There is no padding by default.

_Required_: No

_Type_: [Spacing](aws-properties-quicksight-analysis-spacing.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SectionPageBreakConfiguration

SelectedSheetsFilterScopeConfiguration

All content copied from https://docs.aws.amazon.com/.
