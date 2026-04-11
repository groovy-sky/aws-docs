This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DefaultNewSheetConfiguration

The configuration for default new sheet settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InteractiveLayoutConfiguration" : DefaultInteractiveLayoutConfiguration,
  "PaginatedLayoutConfiguration" : DefaultPaginatedLayoutConfiguration,
  "SheetContentType" : String
}

```

### YAML

```yaml

  InteractiveLayoutConfiguration:
    DefaultInteractiveLayoutConfiguration
  PaginatedLayoutConfiguration:
    DefaultPaginatedLayoutConfiguration
  SheetContentType: String

```

## Properties

`InteractiveLayoutConfiguration`

The options that determine the default settings for interactive layout configuration.

_Required_: No

_Type_: [DefaultInteractiveLayoutConfiguration](aws-properties-quicksight-analysis-defaultinteractivelayoutconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PaginatedLayoutConfiguration`

The options that determine the default settings for a paginated layout configuration.

_Required_: No

_Type_: [DefaultPaginatedLayoutConfiguration](aws-properties-quicksight-analysis-defaultpaginatedlayoutconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SheetContentType`

The option that determines the sheet content type.

_Required_: No

_Type_: String

_Allowed values_: `PAGINATED | INTERACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultInteractiveLayoutConfiguration

DefaultPaginatedLayoutConfiguration

All content copied from https://docs.aws.amazon.com/.
