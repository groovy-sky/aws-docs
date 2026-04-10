This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::OriginEndpoint HlsEncryption

Holds encryption information so that access to the content can be controlled by a DRM solution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConstantInitializationVector" : String,
  "EncryptionMethod" : String,
  "KeyRotationIntervalSeconds" : Integer,
  "RepeatExtXKey" : Boolean,
  "SpekeKeyProvider" : SpekeKeyProvider
}

```

### YAML

```yaml

  ConstantInitializationVector: String
  EncryptionMethod: String
  KeyRotationIntervalSeconds: Integer
  RepeatExtXKey: Boolean
  SpekeKeyProvider:
    SpekeKeyProvider

```

## Properties

`ConstantInitializationVector`

A 128-bit, 16-byte hex value represented by a 32-character string, used with the key for encrypting blocks.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionMethod`

HLS encryption type.

_Required_: No

_Type_: String

_Allowed values_: `AES_128 | SAMPLE_AES`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyRotationIntervalSeconds`

Number of seconds before AWS Elemental MediaPackage rotates to a new key. By default, rotation is set to 60 seconds. Set to `0` to disable key rotation.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepeatExtXKey`

Repeat the `EXT-X-KEY ` directive for every media segment. This
might result in an increase in client requests to the DRM server.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpekeKeyProvider`

Parameters for the SPEKE key provider.

_Required_: Yes

_Type_: [SpekeKeyProvider](aws-properties-mediapackage-originendpoint-spekekeyprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionContractConfiguration

HlsManifest

All content copied from https://docs.aws.amazon.com/.
