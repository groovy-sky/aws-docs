This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::OriginEndpoint DashEncryption

Holds encryption information so that access to the content can be controlled by a DRM solution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KeyRotationIntervalSeconds" : Integer,
  "SpekeKeyProvider" : SpekeKeyProvider
}

```

### YAML

```yaml

  KeyRotationIntervalSeconds: Integer
  SpekeKeyProvider:
    SpekeKeyProvider

```

## Properties

`KeyRotationIntervalSeconds`

Number of seconds before AWS Elemental MediaPackage rotates to a new key. By default, rotation is set to 60 seconds. Set to `0` to disable key rotation.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpekeKeyProvider`

Parameters for the SPEKE key provider.

_Required_: Yes

_Type_: [SpekeKeyProvider](aws-properties-mediapackage-originendpoint-spekekeyprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CmafPackage

DashPackage

All content copied from https://docs.aws.amazon.com/.
