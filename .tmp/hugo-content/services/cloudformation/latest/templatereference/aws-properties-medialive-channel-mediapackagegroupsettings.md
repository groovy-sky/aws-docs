This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel MediaPackageGroupSettings

The settings for the MediaPackage group.

The parent of this entity is OutputGroupSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : OutputLocationRef,
  "MediapackageV2GroupSettings" : MediaPackageV2GroupSettings
}

```

### YAML

```yaml

  Destination:
    OutputLocationRef
  MediapackageV2GroupSettings:
    MediaPackageV2GroupSettings

```

## Properties

`Destination`

The MediaPackage channel destination.

_Required_: No

_Type_: [OutputLocationRef](aws-properties-medialive-channel-outputlocationref.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediapackageV2GroupSettings`

Property description not available.

_Required_: No

_Type_: [MediaPackageV2GroupSettings](aws-properties-medialive-channel-mediapackagev2groupsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MediaPackageAdditionalDestinations

MediaPackageOutputDestinationSettings

All content copied from https://docs.aws.amazon.com/.
