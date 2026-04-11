This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel MsSmoothOutputSettings

Configuration of a Microsoft Smooth output.

The parent of this entity is OutputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "H265PackagingType" : String,
  "NameModifier" : String
}

```

### YAML

```yaml

  H265PackagingType: String
  NameModifier: String

```

## Properties

`H265PackagingType`

Only applicable when this output is referencing an H.265 video description.
Specifies whether MP4 segments should be packaged as HEV1 or HVC1.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NameModifier`

A string that is concatenated to the end of the destination file name. This is required for multiple outputs of
the same type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MsSmoothGroupSettings

MulticastInputSettings

All content copied from https://docs.aws.amazon.com/.
