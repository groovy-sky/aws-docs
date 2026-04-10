This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewayClientTlsCertificate

An object that represents the virtual gateway's client's Transport Layer Security (TLS) certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "File" : VirtualGatewayListenerTlsFileCertificate,
  "SDS" : VirtualGatewayListenerTlsSdsCertificate
}

```

### YAML

```yaml

  File:
    VirtualGatewayListenerTlsFileCertificate
  SDS:
    VirtualGatewayListenerTlsSdsCertificate

```

## Properties

`File`

An object that represents a local file certificate. The certificate must meet specific
requirements and you must have proxy authorization enabled. For more information, see
[Transport Layer Security (TLS)](../../../app-mesh/latest/userguide/tls.md).

_Required_: No

_Type_: [VirtualGatewayListenerTlsFileCertificate](aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SDS`

A reference to an object that represents a virtual gateway's client's Secret Discovery
Service certificate.

_Required_: No

_Type_: [VirtualGatewayListenerTlsSdsCertificate](aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlssdscertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayClientPolicyTls

VirtualGatewayConnectionPool

All content copied from https://docs.aws.amazon.com/.
