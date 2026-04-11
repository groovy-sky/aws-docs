This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewayTlsValidationContextAcmTrust

An object that represents a Transport Layer Security (TLS) validation context trust for an AWS Certificate Manager
certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateAuthorityArns" : [ String, ... ]
}

```

### YAML

```yaml

  CertificateAuthorityArns:
    - String

```

## Properties

`CertificateAuthorityArns`

One or more ACM Amazon Resource Name (ARN)s.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayTlsValidationContext

VirtualGatewayTlsValidationContextFileTrust

All content copied from https://docs.aws.amazon.com/.
