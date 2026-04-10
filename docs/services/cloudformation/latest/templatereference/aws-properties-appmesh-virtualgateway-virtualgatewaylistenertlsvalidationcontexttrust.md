This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewayListenerTlsValidationContextTrust

An object that represents a virtual gateway's listener's Transport Layer Security (TLS) validation context
trust.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "File" : VirtualGatewayTlsValidationContextFileTrust,
  "SDS" : VirtualGatewayTlsValidationContextSdsTrust
}

```

### YAML

```yaml

  File:
    VirtualGatewayTlsValidationContextFileTrust
  SDS:
    VirtualGatewayTlsValidationContextSdsTrust

```

## Properties

`File`

An object that represents a Transport Layer Security (TLS) validation context trust for a local file.

_Required_: No

_Type_: [VirtualGatewayTlsValidationContextFileTrust](aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextfiletrust.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SDS`

A reference to an object that represents a virtual gateway's listener's Transport Layer Security (TLS) Secret
Discovery Service validation context trust.

_Required_: No

_Type_: [VirtualGatewayTlsValidationContextSdsTrust](aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextsdstrust.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayListenerTlsValidationContext

VirtualGatewayLogging

All content copied from https://docs.aws.amazon.com/.
