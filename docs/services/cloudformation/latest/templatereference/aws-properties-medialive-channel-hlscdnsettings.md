This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel HlsCdnSettings

The settings for the CDN of an HLS output.

The parent of this entity is HlsGroupSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HlsAkamaiSettings" : HlsAkamaiSettings,
  "HlsBasicPutSettings" : HlsBasicPutSettings,
  "HlsMediaStoreSettings" : HlsMediaStoreSettings,
  "HlsS3Settings" : HlsS3Settings,
  "HlsWebdavSettings" : HlsWebdavSettings
}

```

### YAML

```yaml

  HlsAkamaiSettings:
    HlsAkamaiSettings
  HlsBasicPutSettings:
    HlsBasicPutSettings
  HlsMediaStoreSettings:
    HlsMediaStoreSettings
  HlsS3Settings:
    HlsS3Settings
  HlsWebdavSettings:
    HlsWebdavSettings

```

## Properties

`HlsAkamaiSettings`

Sets up Akamai as the downstream system for the HLS output group.

_Required_: No

_Type_: [HlsAkamaiSettings](aws-properties-medialive-channel-hlsakamaisettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsBasicPutSettings`

The settings for Basic Put for the HLS output.

_Required_: No

_Type_: [HlsBasicPutSettings](aws-properties-medialive-channel-hlsbasicputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsMediaStoreSettings`

Sets up MediaStore as the destination for the HLS output.

_Required_: No

_Type_: [HlsMediaStoreSettings](aws-properties-medialive-channel-hlsmediastoresettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsS3Settings`

Sets up Amazon S3 as the destination for this HLS output.

_Required_: No

_Type_: [HlsS3Settings](aws-properties-medialive-channel-hlss3settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsWebdavSettings`

The settings for Web VTT captions in the HLS output group.

The parent of this entity is HlsGroupSettings.

_Required_: No

_Type_: [HlsWebdavSettings](aws-properties-medialive-channel-hlswebdavsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HlsBasicPutSettings

HlsGroupSettings

All content copied from https://docs.aws.amazon.com/.
