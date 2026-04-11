This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode TlsValidationContextFileTrust

An object that represents a Transport Layer Security (TLS) validation context trust for a local file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateChain" : String
}

```

### YAML

```yaml

  CertificateChain: String

```

## Properties

`CertificateChain`

The certificate trust chain for a certificate stored on the file system of the virtual
node that the proxy is running on.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TlsValidationContextAcmTrust

TlsValidationContextSdsTrust

All content copied from https://docs.aws.amazon.com/.
