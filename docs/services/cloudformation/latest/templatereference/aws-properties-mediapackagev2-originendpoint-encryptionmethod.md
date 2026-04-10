This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint EncryptionMethod

The encryption method associated with the origin endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CmafEncryptionMethod" : String,
  "IsmEncryptionMethod" : String,
  "TsEncryptionMethod" : String
}

```

### YAML

```yaml

  CmafEncryptionMethod: String
  IsmEncryptionMethod: String
  TsEncryptionMethod: String

```

## Properties

`CmafEncryptionMethod`

The encryption method to use.

_Required_: No

_Type_: String

_Allowed values_: `CENC | CBCS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsmEncryptionMethod`

The encryption method used for Microsoft Smooth Streaming (MSS) content. This specifies how the MSS segments are encrypted to protect the content during delivery to client players.

_Required_: No

_Type_: String

_Allowed values_: `CENC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TsEncryptionMethod`

The encryption method to use.

_Required_: No

_Type_: String

_Allowed values_: `AES_128 | SAMPLE_AES`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionContractConfiguration

FilterConfiguration

All content copied from https://docs.aws.amazon.com/.
