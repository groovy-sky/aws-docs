This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::CertificateAuthorityActivation

The `AWS::ACMPCA::CertificateAuthorityActivation` resource creates and
installs a CA certificate on a CA. If no status is specified, the
`AWS::ACMPCA::CertificateAuthorityActivation` resource status defaults to
ACTIVE. Once the CA has a CA certificate installed, you can use the resource to toggle
the CA status field between `ACTIVE` and `DISABLED`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ACMPCA::CertificateAuthorityActivation",
  "Properties" : {
      "Certificate" : String,
      "CertificateAuthorityArn" : String,
      "CertificateChain" : String,
      "Status" : String
    }
}

```

### YAML

```yaml

Type: AWS::ACMPCA::CertificateAuthorityActivation
Properties:
  Certificate: String
  CertificateAuthorityArn: String
  CertificateChain: String
  Status: String

```

## Properties

`Certificate`

The Base64 PEM-encoded certificate authority certificate.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CertificateAuthorityArn`

The Amazon Resource Name (ARN) of your private CA.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CertificateChain`

The Base64 PEM-encoded certificate chain that chains up to the root CA certificate
that you used to sign your private CA certificate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Status of your private CA.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

The Amazon Resource Name (ARN) of the certificate authority.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified
attribute of this type. The following are the available attributes and sample return
values.

For more information about using the `Fn::GetAtt` intrinsic function, see
[Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`CompleteCertificateChain`

The complete Base64 PEM-encoded certificate chain, including the certificate authority
certificate.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::ACMPCA::Permission

All content copied from https://docs.aws.amazon.com/.
