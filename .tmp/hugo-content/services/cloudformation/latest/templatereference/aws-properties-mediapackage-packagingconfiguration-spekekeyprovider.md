This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::PackagingConfiguration SpekeKeyProvider

A configuration for accessing an external Secure Packager and Encoder Key Exchange
(SPEKE) service that provides encryption keys.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionContractConfiguration" : EncryptionContractConfiguration,
  "RoleArn" : String,
  "SystemIds" : [ String, ... ],
  "Url" : String
}

```

### YAML

```yaml

  EncryptionContractConfiguration:
    EncryptionContractConfiguration
  RoleArn: String
  SystemIds:
    - String
  Url: String

```

## Properties

`EncryptionContractConfiguration`

Use `encryptionContractConfiguration` to configure one or more content encryption keys for your
endpoints that use SPEKE Version 2.0. The encryption contract defines which content keys are used to encrypt the
audio and video tracks in your stream. To configure the encryption contract, specify which audio and video encryption presets to use.

_Required_: No

_Type_: [EncryptionContractConfiguration](aws-properties-mediapackage-packagingconfiguration-encryptioncontractconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN for the IAM role that's granted by the key provider to provide access
to the key provider API. Valid format: arn:aws:iam::{accountID}:role/{name}

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SystemIds`

List of unique identifiers for the DRM systems to use, as defined in the CPIX specification.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

URL for the key provider's key retrieval API endpoint. Must start with https://.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MssPackage

StreamSelection

All content copied from https://docs.aws.amazon.com/.
