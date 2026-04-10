This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint DashProgramInformation

Details about the content that you want MediaPackage to pass through in the manifest to the playback device.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Copyright" : String,
  "LanguageCode" : String,
  "MoreInformationUrl" : String,
  "Source" : String,
  "Title" : String
}

```

### YAML

```yaml

  Copyright: String
  LanguageCode: String
  MoreInformationUrl: String
  Source: String
  Title: String

```

## Properties

`Copyright`

A copyright statement about the content.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LanguageCode`

The language code for this manifest.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_-]*[a-zA-Z0-9]$`

_Minimum_: `2`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MoreInformationUrl`

An absolute URL that contains more information about this content.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

Information about the content provider.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title for the manifest.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashManifestConfiguration

DashSubtitleConfiguration

All content copied from https://docs.aws.amazon.com/.
