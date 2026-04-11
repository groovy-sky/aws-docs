This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel ArchiveContainerSettings

The archive container settings.

The parent of this entity is ArchiveOutputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "M2tsSettings" : M2tsSettings,
  "RawSettings" : Json
}

```

### YAML

```yaml

  M2tsSettings:
    M2tsSettings
  RawSettings: Json

```

## Properties

`M2tsSettings`

The settings for the M2TS in the archive output.

_Required_: No

_Type_: [M2tsSettings](aws-properties-medialive-channel-m2tssettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RawSettings`

The settings for Raw archive output type.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArchiveCdnSettings

ArchiveGroupSettings

All content copied from https://docs.aws.amazon.com/.
