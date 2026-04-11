This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint Encryption

The parameters for encrypting content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CmafExcludeSegmentDrmMetadata" : Boolean,
  "ConstantInitializationVector" : String,
  "EncryptionMethod" : EncryptionMethod,
  "KeyRotationIntervalSeconds" : Integer,
  "SpekeKeyProvider" : SpekeKeyProvider
}

```

### YAML

```yaml

  CmafExcludeSegmentDrmMetadata: Boolean
  ConstantInitializationVector: String
  EncryptionMethod:
    EncryptionMethod
  KeyRotationIntervalSeconds: Integer
  SpekeKeyProvider:
    SpekeKeyProvider

```

## Properties

`CmafExcludeSegmentDrmMetadata`

Excludes SEIG and SGPD boxes from segment metadata in CMAF containers.

When set to `true`, MediaPackage omits these DRM metadata boxes from CMAF segments, which can improve compatibility with certain devices and players that don't support these boxes.

Important considerations:

- This setting only affects CMAF container formats

- Key rotation can still be handled through media playlist signaling

- PSSH and TENC boxes remain unaffected

- Default behavior is preserved when this setting is disabled

Valid values: `true` \| `false`

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConstantInitializationVector`

A 128-bit, 16-byte hex value represented by a 32-character string, used in conjunction with the key for encrypting content. If you don't specify a value, then MediaPackage creates the constant initialization vector (IV).

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-fA-F]+$`

_Minimum_: `32`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionMethod`

The encryption method to use.

_Required_: Yes

_Type_: [EncryptionMethod](aws-properties-mediapackagev2-originendpoint-encryptionmethod.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyRotationIntervalSeconds`

The interval, in seconds, to rotate encryption keys for the origin endpoint.

_Required_: No

_Type_: Integer

_Minimum_: `300`

_Maximum_: `31536000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpekeKeyProvider`

The SPEKE key provider to use for encryption.

_Required_: Yes

_Type_: [SpekeKeyProvider](aws-properties-mediapackagev2-originendpoint-spekekeyprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashUtcTiming

EncryptionContractConfiguration

All content copied from https://docs.aws.amazon.com/.
