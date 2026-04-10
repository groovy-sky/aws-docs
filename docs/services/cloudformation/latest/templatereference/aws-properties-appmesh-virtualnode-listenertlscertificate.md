This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode ListenerTlsCertificate

An object that represents a listener's Transport Layer Security (TLS) certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ACM" : ListenerTlsAcmCertificate,
  "File" : ListenerTlsFileCertificate,
  "SDS" : ListenerTlsSdsCertificate
}

```

### YAML

```yaml

  ACM:
    ListenerTlsAcmCertificate
  File:
    ListenerTlsFileCertificate
  SDS:
    ListenerTlsSdsCertificate

```

## Properties

`ACM`

A reference to an object that represents an AWS Certificate Manager certificate.

_Required_: No

_Type_: [ListenerTlsAcmCertificate](aws-properties-appmesh-virtualnode-listenertlsacmcertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`File`

A reference to an object that represents a local file certificate.

_Required_: No

_Type_: [ListenerTlsFileCertificate](aws-properties-appmesh-virtualnode-listenertlsfilecertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SDS`

A reference to an object that represents a listener's Secret Discovery Service
certificate.

_Required_: No

_Type_: [ListenerTlsSdsCertificate](aws-properties-appmesh-virtualnode-listenertlssdscertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListenerTlsAcmCertificate

ListenerTlsFileCertificate

All content copied from https://docs.aws.amazon.com/.
