This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewayClientPolicyTls

An object that represents a Transport Layer Security (TLS) client policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Certificate" : VirtualGatewayClientTlsCertificate,
  "Enforce" : Boolean,
  "Ports" : [ Integer, ... ],
  "Validation" : VirtualGatewayTlsValidationContext
}

```

### YAML

```yaml

  Certificate:
    VirtualGatewayClientTlsCertificate
  Enforce: Boolean
  Ports:
    - Integer
  Validation:
    VirtualGatewayTlsValidationContext

```

## Properties

`Certificate`

A reference to an object that represents a virtual gateway's client's Transport Layer Security (TLS)
certificate.

_Required_: No

_Type_: [VirtualGatewayClientTlsCertificate](aws-properties-appmesh-virtualgateway-virtualgatewayclienttlscertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enforce`

Whether the policy is enforced. The default is `True`, if a value isn't
specified.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ports`

One or more ports that the policy is enforced for.

_Required_: No

_Type_: Array of Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Validation`

A reference to an object that represents a Transport Layer Security (TLS) validation context.

_Required_: Yes

_Type_: [VirtualGatewayTlsValidationContext](aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontext.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayClientPolicy

VirtualGatewayClientTlsCertificate

All content copied from https://docs.aws.amazon.com/.
