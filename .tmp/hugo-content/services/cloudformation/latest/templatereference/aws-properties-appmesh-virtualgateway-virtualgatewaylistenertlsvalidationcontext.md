This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewayListenerTlsValidationContext

An object that represents a virtual gateway's listener's Transport Layer Security (TLS) validation
context.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SubjectAlternativeNames" : SubjectAlternativeNames,
  "Trust" : VirtualGatewayListenerTlsValidationContextTrust
}

```

### YAML

```yaml

  SubjectAlternativeNames:
    SubjectAlternativeNames
  Trust:
    VirtualGatewayListenerTlsValidationContextTrust

```

## Properties

`SubjectAlternativeNames`

A reference to an object that represents the SANs for a virtual gateway listener's Transport Layer Security (TLS)
validation context.

_Required_: No

_Type_: [SubjectAlternativeNames](aws-properties-appmesh-virtualgateway-subjectalternativenames.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Trust`

A reference to where to retrieve the trust chain when validating a peer’s Transport Layer Security (TLS)
certificate.

_Required_: Yes

_Type_: [VirtualGatewayListenerTlsValidationContextTrust](aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsvalidationcontexttrust.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayListenerTlsSdsCertificate

VirtualGatewayListenerTlsValidationContextTrust

All content copied from https://docs.aws.amazon.com/.
