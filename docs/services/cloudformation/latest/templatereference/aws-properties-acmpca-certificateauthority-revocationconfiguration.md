This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::CertificateAuthority RevocationConfiguration

Certificate revocation information used by the [CreateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html) and [UpdateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_UpdateCertificateAuthority.html) actions. Your private certificate authority (CA)
can configure Online Certificate Status Protocol (OCSP) support and/or maintain a
certificate revocation list (CRL). OCSP returns validation information about
certificates as requested by clients, and a CRL contains an updated list of certificates
revoked by your CA. For more information, see [RevokeCertificate](https://docs.aws.amazon.com/privateca/latest/APIReference/API_RevokeCertificate.html) in the _AWS Private CA API_
_Reference_ and [Setting up a certificate\
revocation method](https://docs.aws.amazon.com/privateca/latest/userguide/revocation-setup.html) in the _AWS Private CA User_
_Guide_.

The following requirements and constraints apply to revocation configurations.

- A configuration disabling CRLs or OCSP must contain only the
`Enabled=False` parameter, and will fail if other parameters
such as `CustomCname` or `ExpirationInDays` are
included.

- In a CRL configuration, the `S3BucketName` parameter must
conform to the [Amazon S3 bucket\
naming rules](../../../s3/latest/userguide/bucketnamingrules.md).

- A configuration containing a custom Canonical Name (CNAME) parameter for
CRLs or OCSP must conform to [RFC2396](https://www.ietf.org/rfc/rfc2396.txt) restrictions
on the use of special characters in a CNAME.

- In a CRL or OCSP configuration, the value of a CNAME parameter must not
include a protocol prefix such as "http://" or "https://".

- To revoke a certificate, delete the resource from your template, and call the AWS Private CA [RevokeCertificate](https://docs.aws.amazon.com/privateca/latest/APIReference/API_RevokeCertificate.html) API and specify the resource's certificate authority ARN.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrlConfiguration" : CrlConfiguration,
  "OcspConfiguration" : OcspConfiguration
}

```

### YAML

```yaml

  CrlConfiguration:
    CrlConfiguration
  OcspConfiguration:
    OcspConfiguration

```

## Properties

`CrlConfiguration`

Configuration of the certificate revocation list (CRL), if any, maintained by your
private CA.

_Required_: No

_Type_: [CrlConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-acmpca-certificateauthority-crlconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OcspConfiguration`

Configuration of Online Certificate Status Protocol (OCSP) support, if any, maintained
by your private CA.

_Required_: No

_Type_: [OcspConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-acmpca-certificateauthority-ocspconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OtherName

Subject
