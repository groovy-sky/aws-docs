This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::OriginEndpoint SpekeKeyProvider

Key provider settings for DRM.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateArn" : String,
  "EncryptionContractConfiguration" : EncryptionContractConfiguration,
  "ResourceId" : String,
  "RoleArn" : String,
  "SystemIds" : [ String, ... ],
  "Url" : String
}

```

### YAML

```yaml

  CertificateArn: String
  EncryptionContractConfiguration:
    EncryptionContractConfiguration
  ResourceId: String
  RoleArn: String
  SystemIds:
    - String
  Url: String

```

## Properties

`CertificateArn`

The Amazon Resource Name (ARN) for the certificate that you imported to AWS Certificate Manager to add content key encryption to this endpoint. For this feature to work, your DRM key provider must support content key encryption.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionContractConfiguration`

Use `encryptionContractConfiguration` to configure one or more content encryption keys for your
endpoints that use SPEKE Version 2.0. The encryption contract defines which content keys are used to encrypt the
audio and video tracks in your stream. To configure the encryption contract, specify which audio and video encryption presets to use.

_Required_: No

_Type_: [EncryptionContractConfiguration](aws-properties-mediapackage-originendpoint-encryptioncontractconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceId`

Unique identifier for this endpoint, as it is configured in the key provider service.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN for the IAM role that's granted by the key provider to provide
access to the key provider API. This role must have a trust policy that allows AWS Elemental MediaPackage to assume the role, and it must have a sufficient permissions policy
to allow access to the specific key retrieval URL. Valid format: arn:aws:iam::{accountID}:role/{name}

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SystemIds`

List of unique identifiers for the DRM systems to use, as defined in the CPIX specification.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

URL for the key provider’s key retrieval API endpoint. Must start with https://.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MssPackage

StreamSelection

All content copied from https://docs.aws.amazon.com/.
