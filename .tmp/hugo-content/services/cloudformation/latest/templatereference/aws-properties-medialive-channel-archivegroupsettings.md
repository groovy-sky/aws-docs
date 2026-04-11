This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel ArchiveGroupSettings

The settings for an archive output group.

The parent of this entity is OutputGroupSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArchiveCdnSettings" : ArchiveCdnSettings,
  "Destination" : OutputLocationRef,
  "RolloverInterval" : Integer
}

```

### YAML

```yaml

  ArchiveCdnSettings:
    ArchiveCdnSettings
  Destination:
    OutputLocationRef
  RolloverInterval: Integer

```

## Properties

`ArchiveCdnSettings`

Settings to configure the destination of an Archive output.

_Required_: No

_Type_: [ArchiveCdnSettings](aws-properties-medialive-channel-archivecdnsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destination`

A directory and base file name where archive files should be written.

_Required_: No

_Type_: [OutputLocationRef](aws-properties-medialive-channel-outputlocationref.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RolloverInterval`

The number of seconds to write to an archive file before closing and starting a new one.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArchiveContainerSettings

ArchiveOutputSettings

All content copied from https://docs.aws.amazon.com/.
