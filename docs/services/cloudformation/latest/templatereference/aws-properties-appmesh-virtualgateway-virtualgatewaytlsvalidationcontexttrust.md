This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewayTlsValidationContextTrust

An object that represents a Transport Layer Security (TLS) validation context trust.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ACM" : VirtualGatewayTlsValidationContextAcmTrust,
  "File" : VirtualGatewayTlsValidationContextFileTrust,
  "SDS" : VirtualGatewayTlsValidationContextSdsTrust
}

```

### YAML

```yaml

  ACM:
    VirtualGatewayTlsValidationContextAcmTrust
  File:
    VirtualGatewayTlsValidationContextFileTrust
  SDS:
    VirtualGatewayTlsValidationContextSdsTrust

```

## Properties

`ACM`

A reference to an object that represents a Transport Layer Security (TLS) validation context trust for an AWS Certificate Manager certificate.

_Required_: No

_Type_: [VirtualGatewayTlsValidationContextAcmTrust](aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextacmtrust.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`File`

An object that represents a Transport Layer Security (TLS) validation context trust for a local file.

_Required_: No

_Type_: [VirtualGatewayTlsValidationContextFileTrust](aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextfiletrust.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SDS`

A reference to an object that represents a virtual gateway's Transport Layer Security (TLS) Secret Discovery
Service validation context trust.

_Required_: No

_Type_: [VirtualGatewayTlsValidationContextSdsTrust](aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextsdstrust.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayTlsValidationContextSdsTrust

AWS::AppMesh::VirtualNode

All content copied from https://docs.aws.amazon.com/.
