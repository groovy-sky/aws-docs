This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template PrivateKeyAttributesV4

Defines the attributes of the private key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Algorithm" : String,
  "CryptoProviders" : [ String, ... ],
  "KeySpec" : String,
  "KeyUsageProperty" : KeyUsageProperty,
  "MinimalKeyLength" : Number
}

```

### YAML

```yaml

  Algorithm: String
  CryptoProviders:
    - String
  KeySpec: String
  KeyUsageProperty:
    KeyUsageProperty
  MinimalKeyLength: Number

```

## Properties

`Algorithm`

Defines the algorithm used to generate the private key.

_Required_: No

_Type_: String

_Allowed values_: `RSA | ECDH_P256 | ECDH_P384 | ECDH_P521`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CryptoProviders`

Defines the cryptographic providers used to generate the private key.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `100 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeySpec`

Defines the purpose of the private key. Set it to "KEY\_EXCHANGE" or "SIGNATURE"
value.

_Required_: Yes

_Type_: String

_Allowed values_: `KEY_EXCHANGE | SIGNATURE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyUsageProperty`

The key usage property defines the purpose of the private key contained in the
certificate. You can specify specific purposes using property flags or all by using
property type ALL.

_Required_: No

_Type_: [KeyUsageProperty](aws-properties-pcaconnectorad-template-keyusageproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimalKeyLength`

Set the minimum key length of the private key.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PrivateKeyAttributesV3

PrivateKeyFlagsV2

All content copied from https://docs.aws.amazon.com/.
