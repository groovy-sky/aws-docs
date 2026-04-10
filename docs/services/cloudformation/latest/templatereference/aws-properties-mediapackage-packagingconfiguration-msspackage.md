This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::PackagingConfiguration MssPackage

Parameters for a packaging configuration that uses Microsoft Smooth Streaming (MSS) packaging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Encryption" : MssEncryption,
  "MssManifests" : [ MssManifest, ... ],
  "SegmentDurationSeconds" : Integer
}

```

### YAML

```yaml

  Encryption:
    MssEncryption
  MssManifests:
    - MssManifest
  SegmentDurationSeconds: Integer

```

## Properties

`Encryption`

Parameters for encrypting content.

_Required_: No

_Type_: [MssEncryption](aws-properties-mediapackage-packagingconfiguration-mssencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MssManifests`

A list of Microsoft Smooth manifest configurations that are available from this endpoint.

_Required_: Yes

_Type_: Array of [MssManifest](aws-properties-mediapackage-packagingconfiguration-mssmanifest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentDurationSeconds`

Duration (in seconds) of each fragment. Actual fragments are rounded to the nearest multiple of the source fragment duration.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MssManifest

SpekeKeyProvider

All content copied from https://docs.aws.amazon.com/.
