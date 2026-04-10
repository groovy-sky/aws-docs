This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::SecurityConfiguration TLSCertificateConfiguration

Configurations related to the TLS certificate for the security configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateProviderType" : String,
  "PrivateKeySecretArn" : String,
  "PublicKeySecretArn" : String
}

```

### YAML

```yaml

  CertificateProviderType: String
  PrivateKeySecretArn: String
  PublicKeySecretArn: String

```

## Properties

`CertificateProviderType`

The TLS certificate type. Acceptable values: `PEM` or
`Custom`.

_Required_: No

_Type_: String

_Allowed values_: `PEM`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrivateKeySecretArn`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PublicKeySecretArn`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::EMRContainers::VirtualCluster

All content copied from https://docs.aws.amazon.com/.
