This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Workspace MediaItem

Contains information about a media asset used in a workspace.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Source" : String,
  "Type" : String
}

```

### YAML

```yaml

  Source: String
  Type: String

```

## Properties

`Source`

The source URL or data for the media asset.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `533333`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of media. Valid values are: `IMAGE_LOGO_FAVICON` and
`IMAGE_LOGO_HORIZONTAL`.

_Required_: Yes

_Type_: String

_Allowed values_: `IMAGE_LOGO_LIGHT_FAVICON | IMAGE_LOGO_DARK_FAVICON | IMAGE_LOGO_LIGHT_HORIZONTAL | IMAGE_LOGO_DARK_HORIZONTAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FontFamily

PaletteCanvas

All content copied from https://docs.aws.amazon.com/.
