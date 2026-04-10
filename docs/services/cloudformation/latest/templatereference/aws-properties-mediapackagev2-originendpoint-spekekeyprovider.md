This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint SpekeKeyProvider

The parameters for the SPEKE key provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateArn" : String,
  "DrmSystems" : [ String, ... ],
  "EncryptionContractConfiguration" : EncryptionContractConfiguration,
  "ResourceId" : String,
  "RoleArn" : String,
  "Url" : String
}

```

### YAML

```yaml

  CertificateArn: String
  DrmSystems:
    - String
  EncryptionContractConfiguration:
    EncryptionContractConfiguration
  ResourceId: String
  RoleArn: String
  Url: String

```

## Properties

`CertificateArn`

The ARN for the certificate that you imported to AWS Certificate Manager to add content key encryption to this endpoint. For this feature to work, your DRM key provider must support content key encryption.

_Required_: No

_Type_: String

_Pattern_: `^arn:([^:\n]+):acm:([^:\n]+):([0-9]+):certificate/[a-zA-Z0-9-_]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DrmSystems`

The DRM solution provider you're using to protect your content during distribution.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionContractConfiguration`

The encryption contract configuration associated with the SPEKE key provider.

_Required_: Yes

_Type_: [EncryptionContractConfiguration](aws-properties-mediapackagev2-originendpoint-encryptioncontractconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceId`

The unique identifier for the content. The service sends this identifier to the key server to identify the current endpoint. How unique you make this identifier depends on how fine-grained you want access controls to be. The service does not permit you to use the same ID for two simultaneous encryption processes. The resource ID is also known as the content ID.

The following example shows a resource ID: `MovieNight20171126093045`

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN for the IAM role granted by the key provider that provides access to the key provider API. This role must have a trust policy that allows MediaPackage to assume the role, and it must have a sufficient permissions policy to allow access to the specific key retrieval URL. Get this from your DRM solution provider.

Valid format: `arn:aws:iam::{accountID}:role/{name}`. The following example shows a role ARN: `arn:aws:iam::444455556666:role/SpekeAccess`

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL of the SPEKE key provider.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Segment

StartTag

All content copied from https://docs.aws.amazon.com/.
