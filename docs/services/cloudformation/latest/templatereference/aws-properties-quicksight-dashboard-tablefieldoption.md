This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TableFieldOption

The options for a table field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomLabel" : String,
  "FieldId" : String,
  "URLStyling" : TableFieldURLConfiguration,
  "Visibility" : String,
  "Width" : String
}

```

### YAML

```yaml

  CustomLabel: String
  FieldId: String
  URLStyling:
    TableFieldURLConfiguration
  Visibility: String
  Width: String

```

## Properties

`CustomLabel`

The custom label for a table field.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldId`

The field ID for a table field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`URLStyling`

The URL configuration for a table field.

_Required_: No

_Type_: [TableFieldURLConfiguration](aws-properties-quicksight-dashboard-tablefieldurlconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility of a table field.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Width`

The width for a table field.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableFieldLinkContentConfiguration

TableFieldOptions

All content copied from https://docs.aws.amazon.com/.
