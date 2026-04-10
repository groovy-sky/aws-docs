This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis TableFieldLinkContentConfiguration

The URL content (text, icon) for the table link configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomIconContent" : TableFieldCustomIconContent,
  "CustomTextContent" : TableFieldCustomTextContent
}

```

### YAML

```yaml

  CustomIconContent:
    TableFieldCustomIconContent
  CustomTextContent:
    TableFieldCustomTextContent

```

## Properties

`CustomIconContent`

The custom icon content for the table link content configuration.

_Required_: No

_Type_: [TableFieldCustomIconContent](aws-properties-quicksight-analysis-tablefieldcustomiconcontent.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomTextContent`

The custom text content (value, font configuration) for the table link content configuration.

_Required_: No

_Type_: [TableFieldCustomTextContent](aws-properties-quicksight-analysis-tablefieldcustomtextcontent.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableFieldLinkConfiguration

TableFieldOption

All content copied from https://docs.aws.amazon.com/.
