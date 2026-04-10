This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode ListenerTlsAcmCertificate

An object that represents an AWS Certificate Manager certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateArn" : String
}

```

### YAML

```yaml

  CertificateArn: String

```

## Properties

`CertificateArn`

The Amazon Resource Name (ARN) for the certificate. The certificate must meet specific requirements and you must have proxy authorization enabled. For more information, see [Transport Layer Security (TLS)](../../../app-mesh/latest/userguide/tls.md#virtual-node-tls-prerequisites).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListenerTls

ListenerTlsCertificate

All content copied from https://docs.aws.amazon.com/.
