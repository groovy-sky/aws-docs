This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel ArchiveOutputSettings

The archive output settings.

The parent of this entity is OutputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerSettings" : ArchiveContainerSettings,
  "Extension" : String,
  "NameModifier" : String
}

```

### YAML

```yaml

  ContainerSettings:
    ArchiveContainerSettings
  Extension: String
  NameModifier: String

```

## Properties

`ContainerSettings`

The settings that are specific to the container type of the file.

_Required_: No

_Type_: [ArchiveContainerSettings](aws-properties-medialive-channel-archivecontainersettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Extension`

The output file extension. If excluded, this is auto-selected from the container type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NameModifier`

A string that is concatenated to the end of the destination file name. The string is required for multiple
outputs of the same type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArchiveGroupSettings

ArchiveS3Settings

All content copied from https://docs.aws.amazon.com/.
