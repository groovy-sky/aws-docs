This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TableFieldCustomTextContent

The custom text content (value, font configuration) for the table link content configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FontConfiguration" : FontConfiguration,
  "Value" : String
}

```

### YAML

```yaml

  FontConfiguration:
    FontConfiguration
  Value: String

```

## Properties

`FontConfiguration`

The font configuration of the custom text content for the table URL link content.

_Required_: Yes

_Type_: [FontConfiguration](aws-properties-quicksight-dashboard-fontconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The string value of the custom text content for the table URL link content.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableFieldCustomIconContent

TableFieldImageConfiguration

All content copied from https://docs.aws.amazon.com/.
