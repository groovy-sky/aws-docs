This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode ClientTlsCertificate

An object that represents the client's certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "File" : ListenerTlsFileCertificate,
  "SDS" : ListenerTlsSdsCertificate
}

```

### YAML

```yaml

  File:
    ListenerTlsFileCertificate
  SDS:
    ListenerTlsSdsCertificate

```

## Properties

`File`

An object that represents a local file certificate. The certificate must meet specific
requirements and you must have proxy authorization enabled. For more information, see
[Transport Layer Security (TLS)](../../../app-mesh/latest/userguide/tls.md).

_Required_: No

_Type_: [ListenerTlsFileCertificate](aws-properties-appmesh-virtualnode-listenertlsfilecertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SDS`

A reference to an object that represents a client's TLS Secret Discovery Service
certificate.

_Required_: No

_Type_: [ListenerTlsSdsCertificate](aws-properties-appmesh-virtualnode-listenertlssdscertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientPolicyTls

DnsServiceDiscovery

All content copied from https://docs.aws.amazon.com/.
