This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewayListenerTlsFileCertificate

An object that represents a local file certificate.
The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](../../../app-mesh/latest/userguide/tls.md#virtual-node-tls-prerequisites).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateChain" : String,
  "PrivateKey" : String
}

```

### YAML

```yaml

  CertificateChain: String
  PrivateKey: String

```

## Properties

`CertificateChain`

The certificate chain for the certificate.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateKey`

The private key for a certificate stored on the file system of the mesh endpoint that
the proxy is running on.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayListenerTlsCertificate

VirtualGatewayListenerTlsSdsCertificate

All content copied from https://docs.aws.amazon.com/.
