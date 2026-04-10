This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard SpatialStaticFile

A static file that contains the geospatial data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Source" : StaticFileSource,
  "StaticFileId" : String
}

```

### YAML

```yaml

  Source:
    StaticFileSource
  StaticFileId: String

```

## Properties

`Source`

The source of the spatial static file.

_Required_: No

_Type_: [StaticFileSource](aws-properties-quicksight-dashboard-staticfilesource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticFileId`

The ID of the spatial static file.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Spacing

StaticFile

All content copied from https://docs.aws.amazon.com/.
