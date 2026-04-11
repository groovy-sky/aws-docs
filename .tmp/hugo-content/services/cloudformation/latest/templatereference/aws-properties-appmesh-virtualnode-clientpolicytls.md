This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode ClientPolicyTls

A reference to an object that represents a Transport Layer Security (TLS) client policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Certificate" : ClientTlsCertificate,
  "Enforce" : Boolean,
  "Ports" : [ Integer, ... ],
  "Validation" : TlsValidationContext
}

```

### YAML

```yaml

  Certificate:
    ClientTlsCertificate
  Enforce: Boolean
  Ports:
    - Integer
  Validation:
    TlsValidationContext

```

## Properties

`Certificate`

A reference to an object that represents a client's TLS certificate.

_Required_: No

_Type_: [ClientTlsCertificate](aws-properties-appmesh-virtualnode-clienttlscertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enforce`

Whether the policy is enforced. The default is `True`, if a value isn't specified.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ports`

One or more ports that the policy is enforced for.

_Required_: No

_Type_: Array of Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Validation`

A reference to an object that represents a TLS validation context.

_Required_: Yes

_Type_: [TlsValidationContext](aws-properties-appmesh-virtualnode-tlsvalidationcontext.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientPolicy

ClientTlsCertificate

All content copied from https://docs.aws.amazon.com/.
