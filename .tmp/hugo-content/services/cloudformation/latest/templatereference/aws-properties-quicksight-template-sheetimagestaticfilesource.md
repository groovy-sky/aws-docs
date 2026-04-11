This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template SheetImageStaticFileSource

The source of the static file that contains the image.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StaticFileId" : String
}

```

### YAML

```yaml

  StaticFileId: String

```

## Properties

`StaticFileId`

The ID of the static file that contains the image.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SheetImageSource

SheetImageTooltipConfiguration

All content copied from https://docs.aws.amazon.com/.
