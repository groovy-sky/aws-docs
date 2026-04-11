This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Workspace WorkspacePage

Contains information about a page configuration in a workspace, including the view assigned to the page.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputData" : String,
  "Page" : String,
  "ResourceArn" : String,
  "Slug" : String
}

```

### YAML

```yaml

  InputData: String
  Page: String
  ResourceArn: String
  Slug: String

```

## Properties

`InputData`

A JSON string containing input parameters passed to the view when the page is rendered.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Page`

The page identifier. System pages include `HOME` and `AGENT_EXPERIENCE`.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\.$)(?!\.\.$)[\p{L}\p{Z}\p{N}\-_.:=@'|]+$`

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

The Amazon Resource Name (ARN) of the view associated with this page.

_Required_: Yes

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slug`

The URL-friendly identifier for the page.

_Required_: No

_Type_: String

_Pattern_: `^$|^[\p{L}\p{Z}\p{N}\-_.:=@'|]{3,}$`

_Minimum_: `0`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

WorkspaceTheme

All content copied from https://docs.aws.amazon.com/.
