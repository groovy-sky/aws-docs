This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode TlsValidationContext

An object that represents how the proxy will validate its peer during Transport Layer Security (TLS)
negotiation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SubjectAlternativeNames" : SubjectAlternativeNames,
  "Trust" : TlsValidationContextTrust
}

```

### YAML

```yaml

  SubjectAlternativeNames:
    SubjectAlternativeNames
  Trust:
    TlsValidationContextTrust

```

## Properties

`SubjectAlternativeNames`

A reference to an object that represents the SANs for a Transport Layer Security (TLS) validation context. If you
don't specify SANs on the _terminating_ mesh endpoint, the Envoy proxy
for that node doesn't verify the SAN on a peer client certificate. If you don't specify
SANs on the _originating_ mesh endpoint, the SAN on the certificate
provided by the terminating endpoint must match the mesh endpoint service discovery
configuration. Since SPIRE vended certificates have a SPIFFE ID as a name, you must set the
SAN since the name doesn't match the service discovery name.

_Required_: No

_Type_: [SubjectAlternativeNames](aws-properties-appmesh-virtualnode-subjectalternativenames.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Trust`

A reference to where to retrieve the trust chain when validating a peer’s Transport Layer Security (TLS)
certificate.

_Required_: Yes

_Type_: [TlsValidationContextTrust](aws-properties-appmesh-virtualnode-tlsvalidationcontexttrust.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TcpTimeout

TlsValidationContextAcmTrust

All content copied from https://docs.aws.amazon.com/.
