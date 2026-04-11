This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode TlsValidationContextTrust

An object that represents a Transport Layer Security (TLS) validation context trust.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ACM" : TlsValidationContextAcmTrust,
  "File" : TlsValidationContextFileTrust,
  "SDS" : TlsValidationContextSdsTrust
}

```

### YAML

```yaml

  ACM:
    TlsValidationContextAcmTrust
  File:
    TlsValidationContextFileTrust
  SDS:
    TlsValidationContextSdsTrust

```

## Properties

`ACM`

A reference to an object that represents a Transport Layer Security (TLS) validation context trust for an AWS Certificate Manager certificate.

_Required_: No

_Type_: [TlsValidationContextAcmTrust](aws-properties-appmesh-virtualnode-tlsvalidationcontextacmtrust.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`File`

An object that represents a Transport Layer Security (TLS) validation context trust for a local file.

_Required_: No

_Type_: [TlsValidationContextFileTrust](aws-properties-appmesh-virtualnode-tlsvalidationcontextfiletrust.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SDS`

A reference to an object that represents a Transport Layer Security (TLS) Secret Discovery Service validation
context trust.

_Required_: No

_Type_: [TlsValidationContextSdsTrust](aws-properties-appmesh-virtualnode-tlsvalidationcontextsdstrust.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TlsValidationContextSdsTrust

VirtualNodeConnectionPool

All content copied from https://docs.aws.amazon.com/.
