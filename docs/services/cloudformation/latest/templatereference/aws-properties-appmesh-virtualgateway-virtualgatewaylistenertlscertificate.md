This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewayListenerTlsCertificate

An object that represents a listener's Transport Layer Security (TLS) certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ACM" : VirtualGatewayListenerTlsAcmCertificate,
  "File" : VirtualGatewayListenerTlsFileCertificate,
  "SDS" : VirtualGatewayListenerTlsSdsCertificate
}

```

### YAML

```yaml

  ACM:
    VirtualGatewayListenerTlsAcmCertificate
  File:
    VirtualGatewayListenerTlsFileCertificate
  SDS:
    VirtualGatewayListenerTlsSdsCertificate

```

## Properties

`ACM`

A reference to an object that represents an AWS Certificate Manager certificate.

_Required_: No

_Type_: [VirtualGatewayListenerTlsAcmCertificate](aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsacmcertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`File`

A reference to an object that represents a local file certificate.

_Required_: No

_Type_: [VirtualGatewayListenerTlsFileCertificate](aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SDS`

A reference to an object that represents a virtual gateway's listener's Secret Discovery
Service certificate.

_Required_: No

_Type_: [VirtualGatewayListenerTlsSdsCertificate](aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlssdscertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayListenerTlsAcmCertificate

VirtualGatewayListenerTlsFileCertificate

All content copied from https://docs.aws.amazon.com/.
