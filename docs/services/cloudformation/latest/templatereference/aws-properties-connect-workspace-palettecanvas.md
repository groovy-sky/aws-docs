This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Workspace PaletteCanvas

Contains color configuration for canvas elements in a workspace theme.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActiveBackground" : String,
  "ContainerBackground" : String,
  "PageBackground" : String
}

```

### YAML

```yaml

  ActiveBackground: String
  ContainerBackground: String
  PageBackground: String

```

## Properties

`ActiveBackground`

The background color for active elements.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerBackground`

The background color for container elements.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PageBackground`

The background color for page elements.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MediaItem

PaletteHeader

All content copied from https://docs.aws.amazon.com/.
