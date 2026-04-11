This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::PackagingConfiguration MssManifest

Parameters for a Microsoft Smooth manifest.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ManifestName" : String,
  "StreamSelection" : StreamSelection
}

```

### YAML

```yaml

  ManifestName: String
  StreamSelection:
    StreamSelection

```

## Properties

`ManifestName`

A short string that's appended to the end of the endpoint URL to create a unique path to this packaging configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamSelection`

Video bitrate limitations for outputs from this packaging configuration.

_Required_: No

_Type_: [StreamSelection](aws-properties-mediapackage-packagingconfiguration-streamselection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MssEncryption

MssPackage

All content copied from https://docs.aws.amazon.com/.
