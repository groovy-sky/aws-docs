This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::PackagingConfiguration MssEncryption

Holds encryption information so that access to the content can be controlled by a DRM solution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SpekeKeyProvider" : SpekeKeyProvider
}

```

### YAML

```yaml

  SpekeKeyProvider:
    SpekeKeyProvider

```

## Properties

`SpekeKeyProvider`

Parameters for the SPEKE key provider.

_Required_: Yes

_Type_: [SpekeKeyProvider](aws-properties-mediapackage-packagingconfiguration-spekekeyprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HlsPackage

MssManifest

All content copied from https://docs.aws.amazon.com/.
