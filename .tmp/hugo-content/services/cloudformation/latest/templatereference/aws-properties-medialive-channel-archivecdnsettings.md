This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel ArchiveCdnSettings

Settings to configure the destination of an Archive output.

The parent of this entity is ArchiveGroupSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArchiveS3Settings" : ArchiveS3Settings
}

```

### YAML

```yaml

  ArchiveS3Settings:
    ArchiveS3Settings

```

## Properties

`ArchiveS3Settings`

Sets up Amazon S3 as the destination for this Archive output.

_Required_: No

_Type_: [ArchiveS3Settings](aws-properties-medialive-channel-archives3settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnywhereSettings

ArchiveContainerSettings

All content copied from https://docs.aws.amazon.com/.
