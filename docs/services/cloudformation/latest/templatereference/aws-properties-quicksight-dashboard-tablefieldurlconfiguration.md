This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TableFieldURLConfiguration

The URL configuration for a table field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImageConfiguration" : TableFieldImageConfiguration,
  "LinkConfiguration" : TableFieldLinkConfiguration
}

```

### YAML

```yaml

  ImageConfiguration:
    TableFieldImageConfiguration
  LinkConfiguration:
    TableFieldLinkConfiguration

```

## Properties

`ImageConfiguration`

The image configuration of a table field URL.

_Required_: No

_Type_: [TableFieldImageConfiguration](aws-properties-quicksight-dashboard-tablefieldimageconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LinkConfiguration`

The link configuration of a table field URL.

_Required_: No

_Type_: [TableFieldLinkConfiguration](aws-properties-quicksight-dashboard-tablefieldlinkconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableFieldOptions

TableFieldWells

All content copied from https://docs.aws.amazon.com/.
