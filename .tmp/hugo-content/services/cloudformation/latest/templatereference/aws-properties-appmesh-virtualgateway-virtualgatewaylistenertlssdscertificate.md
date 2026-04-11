This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewayListenerTlsSdsCertificate

An object that represents the virtual gateway's listener's Secret Discovery Service
certificate.The proxy must be configured with a local SDS provider via a Unix Domain
Socket. See App Mesh [TLS\
documentation](../../../app-mesh/latest/userguide/tls.md) for more info.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecretName" : String
}

```

### YAML

```yaml

  SecretName: String

```

## Properties

`SecretName`

A reference to an object that represents the name of the secret secret requested from
the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or
certificate chain.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayListenerTlsFileCertificate

VirtualGatewayListenerTlsValidationContext

All content copied from https://docs.aws.amazon.com/.
