This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::OriginEndpoint CmafEncryption

Holds encryption information so that access to the content can be controlled by a DRM solution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConstantInitializationVector" : String,
  "EncryptionMethod" : String,
  "KeyRotationIntervalSeconds" : Integer,
  "SpekeKeyProvider" : SpekeKeyProvider
}

```

### YAML

```yaml

  ConstantInitializationVector: String
  EncryptionMethod: String
  KeyRotationIntervalSeconds: Integer
  SpekeKeyProvider:
    SpekeKeyProvider

```

## Properties

`ConstantInitializationVector`

An optional 128-bit, 16-byte hex value represented by a 32-character string, used in conjunction with the key for encrypting blocks. If you don't specify a value, then AWS Elemental MediaPackage creates the constant initialization vector (IV).

_Required_: No

_Type_: String

_Pattern_: `\A[0-9a-fA-F]+\Z`

_Minimum_: `32`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionMethod`

The encryption method to use.

_Required_: No

_Type_: String

_Allowed values_: `SAMPLE_AES | AES_CTR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

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

Authorization

CmafPackage

All content copied from https://docs.aws.amazon.com/.
