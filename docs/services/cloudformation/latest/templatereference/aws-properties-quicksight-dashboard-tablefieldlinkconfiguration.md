This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TableFieldLinkConfiguration

The link configuration of a table field URL.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Content" : TableFieldLinkContentConfiguration,
  "Target" : String
}

```

### YAML

```yaml

  Content:
    TableFieldLinkContentConfiguration
  Target: String

```

## Properties

`Content`

The URL content (text, icon) for the table link configuration.

_Required_: Yes

_Type_: [TableFieldLinkContentConfiguration](aws-properties-quicksight-dashboard-tablefieldlinkcontentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

The URL target (new tab, new window, same tab) for the table link configuration.

_Required_: Yes

_Type_: String

_Allowed values_: `NEW_TAB | NEW_WINDOW | SAME_TAB`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableFieldImageConfiguration

TableFieldLinkContentConfiguration

All content copied from https://docs.aws.amazon.com/.
