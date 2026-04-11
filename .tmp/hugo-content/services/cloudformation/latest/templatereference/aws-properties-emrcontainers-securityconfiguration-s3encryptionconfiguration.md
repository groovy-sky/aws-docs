This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::SecurityConfiguration S3EncryptionConfiguration

The `S3EncryptionConfiguration` property type specifies Property description not available. for an [AWS::EMRContainers::SecurityConfiguration](aws-resource-emrcontainers-securityconfiguration.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionOption" : String,
  "KMSKeyId" : String
}

```

### YAML

```yaml

  EncryptionOption: String
  KMSKeyId: String

```

## Properties

`EncryptionOption`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `SSE-S3 | SSE-KMS | CSE-KMS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KMSKeyId`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LocalDiskEncryptionConfiguration

SecureNamespaceInfo

All content copied from https://docs.aws.amazon.com/.
