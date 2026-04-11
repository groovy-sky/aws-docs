This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate

The `AWS::ACMPCA::Certificate` resource is used to issue a certificate
using your private certificate authority. For more information, see the [IssueCertificate](../../../../reference/privateca/latest/apireference/api-issuecertificate.md) action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ACMPCA::Certificate",
  "Properties" : {
      "ApiPassthrough" : ApiPassthrough,
      "CertificateAuthorityArn" : String,
      "CertificateSigningRequest" : String,
      "SigningAlgorithm" : String,
      "TemplateArn" : String,
      "Validity" : Validity,
      "ValidityNotBefore" : Validity
    }
}

```

### YAML

```yaml

Type: AWS::ACMPCA::Certificate
Properties:
  ApiPassthrough:
    ApiPassthrough
  CertificateAuthorityArn: String
  CertificateSigningRequest: String
  SigningAlgorithm: String
  TemplateArn: String
  Validity:
    Validity
  ValidityNotBefore:
    Validity

```

## Properties

`ApiPassthrough`

Specifies X.509 certificate information to be included in the issued certificate. An
`APIPassthrough` or `APICSRPassthrough` template variant must
be selected, or else this parameter is ignored.

_Required_: No

_Type_: [ApiPassthrough](aws-properties-acmpca-certificate-apipassthrough.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CertificateAuthorityArn`

The Amazon Resource Name (ARN) for the private CA issues the certificate.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CertificateSigningRequest`

The certificate signing request (CSR) for the certificate.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SigningAlgorithm`

The name of the algorithm that will be used to sign the certificate to be issued.

This parameter should not be confused with the `SigningAlgorithm` parameter
used to sign a CSR in the `CreateCertificateAuthority` action.

###### Note

The specified signing algorithm family (RSA or ECDSA) must match the algorithm
family of the CA's secret key.

_Required_: Yes

_Type_: String

_Allowed values_: `SHA256WITHECDSA | SHA384WITHECDSA | SHA512WITHECDSA | SHA256WITHRSA | SHA384WITHRSA | SHA512WITHRSA | SM3WITHSM2 | ML_DSA_44 | ML_DSA_65 | ML_DSA_87`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TemplateArn`

Specifies a custom configuration template to use when issuing a certificate. If this
parameter is not provided, AWS Private CA defaults to the
`EndEntityCertificate/V1` template. For more information about AWS Private CA templates, see [Using Templates](../../../privateca/latest/userguide/usingtemplates.md).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Validity`

The period of time during which the certificate will be valid.

_Required_: Yes

_Type_: [Validity](aws-properties-acmpca-certificate-validity.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidityNotBefore`

Information describing the start of the validity period of the certificate. This
parameter sets the “Not Before" date for the certificate.

By default, when issuing a certificate, AWS Private CA sets the "Not
Before" date to the issuance time minus 60 minutes. This compensates for clock
inconsistencies across computer systems. The `ValidityNotBefore` parameter
can be used to customize the “Not Before” value.

Unlike the `Validity` parameter, the `ValidityNotBefore`
parameter is optional.

The `ValidityNotBefore` value is expressed as an explicit date and time,
using the `Validity` type value `ABSOLUTE`.

_Required_: No

_Type_: [Validity](aws-properties-acmpca-certificate-validity.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

This reference should not be used in CloudFormation templates. Instead, use
`AWS::ACMPCA::Certificate.Arn` to identify a certificate, and use
`AWS::ACMPCA::Certificate.CertificateAuthorityArn` to identify a
certificate authority.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified
attribute of this type. The following are the available attributes and sample return
values.

For more information about using the `Fn::GetAtt` intrinsic function, see
[Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the issued certificate.

`Certificate`

The issued Base64 PEM-encoded certificate.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Private Certificate Authority

ApiPassthrough

All content copied from https://docs.aws.amazon.com/.
