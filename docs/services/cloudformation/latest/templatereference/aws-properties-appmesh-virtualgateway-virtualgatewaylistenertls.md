This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewayListenerTls

An object that represents the Transport Layer Security (TLS) properties for a listener.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Certificate" : VirtualGatewayListenerTlsCertificate,
  "Mode" : String,
  "Validation" : VirtualGatewayListenerTlsValidationContext
}

```

### YAML

```yaml

  Certificate:
    VirtualGatewayListenerTlsCertificate
  Mode: String
  Validation:
    VirtualGatewayListenerTlsValidationContext

```

## Properties

`Certificate`

An object that represents a Transport Layer Security (TLS) certificate.

_Required_: Yes

_Type_: [VirtualGatewayListenerTlsCertificate](aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlscertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

Specify one of the following modes.

- STRICT – Listener only accepts connections with TLS
enabled.

- PERMISSIVE – Listener accepts connections with or
without TLS enabled.

- DISABLED – Listener only accepts connections without
TLS.

_Required_: Yes

_Type_: String

_Allowed values_: `STRICT | PERMISSIVE | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Validation`

A reference to an object that represents a virtual gateway's listener's Transport Layer Security (TLS) validation
context.

_Required_: No

_Type_: [VirtualGatewayListenerTlsValidationContext](aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsvalidationcontext.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayListener

VirtualGatewayListenerTlsAcmCertificate

All content copied from https://docs.aws.amazon.com/.
