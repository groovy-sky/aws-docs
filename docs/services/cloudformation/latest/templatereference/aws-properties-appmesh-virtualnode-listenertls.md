This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode ListenerTls

An object that represents the Transport Layer Security (TLS) properties for a listener.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Certificate" : ListenerTlsCertificate,
  "Mode" : String,
  "Validation" : ListenerTlsValidationContext
}

```

### YAML

```yaml

  Certificate:
    ListenerTlsCertificate
  Mode: String
  Validation:
    ListenerTlsValidationContext

```

## Properties

`Certificate`

A reference to an object that represents a listener's Transport Layer Security (TLS) certificate.

_Required_: Yes

_Type_: [ListenerTlsCertificate](aws-properties-appmesh-virtualnode-listenertlscertificate.md)

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

A reference to an object that represents a listener's Transport Layer Security (TLS) validation context.

_Required_: No

_Type_: [ListenerTlsValidationContext](aws-properties-appmesh-virtualnode-listenertlsvalidationcontext.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListenerTimeout

ListenerTlsAcmCertificate

All content copied from https://docs.aws.amazon.com/.
