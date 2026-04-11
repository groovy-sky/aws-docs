This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template TemplateV4

v4 template schema that can use either Legacy Cryptographic Providers or Key Storage
Providers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateValidity" : CertificateValidity,
  "EnrollmentFlags" : EnrollmentFlagsV4,
  "Extensions" : ExtensionsV4,
  "GeneralFlags" : GeneralFlagsV4,
  "HashAlgorithm" : String,
  "PrivateKeyAttributes" : PrivateKeyAttributesV4,
  "PrivateKeyFlags" : PrivateKeyFlagsV4,
  "SubjectNameFlags" : SubjectNameFlagsV4,
  "SupersededTemplates" : [ String, ... ]
}

```

### YAML

```yaml

  CertificateValidity:
    CertificateValidity
  EnrollmentFlags:
    EnrollmentFlagsV4
  Extensions:
    ExtensionsV4
  GeneralFlags:
    GeneralFlagsV4
  HashAlgorithm: String
  PrivateKeyAttributes:
    PrivateKeyAttributesV4
  PrivateKeyFlags:
    PrivateKeyFlagsV4
  SubjectNameFlags:
    SubjectNameFlagsV4
  SupersededTemplates:
    - String

```

## Properties

`CertificateValidity`

Certificate validity describes the validity and renewal periods of a certificate.

_Required_: Yes

_Type_: [CertificateValidity](aws-properties-pcaconnectorad-template-certificatevalidity.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnrollmentFlags`

Enrollment flags describe the enrollment settings for certificates using the existing
private key and deleting expired or revoked certificates.

_Required_: Yes

_Type_: [EnrollmentFlagsV4](aws-properties-pcaconnectorad-template-enrollmentflagsv4.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Extensions`

Extensions describe the key usage extensions and application policies for a
template.

_Required_: Yes

_Type_: [ExtensionsV4](aws-properties-pcaconnectorad-template-extensionsv4.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GeneralFlags`

General flags describe whether the template is used for computers or users and if the
template can be used with autoenrollment.

_Required_: Yes

_Type_: [GeneralFlagsV4](aws-properties-pcaconnectorad-template-generalflagsv4.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HashAlgorithm`

Specifies the hash algorithm used to hash the private key. Hash algorithm can only be
specified when using Key Storage Providers.

_Required_: No

_Type_: String

_Allowed values_: `SHA256 | SHA384 | SHA512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateKeyAttributes`

Private key attributes allow you to specify the minimal key length, key spec, key usage,
and cryptographic providers for the private key of a certificate for v4 templates. V4
templates allow you to use either Key Storage Providers or Legacy Cryptographic Service
Providers. You specify the cryptography provider category in private key flags.

_Required_: Yes

_Type_: [PrivateKeyAttributesV4](aws-properties-pcaconnectorad-template-privatekeyattributesv4.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateKeyFlags`

Private key flags for v4 templates specify the client compatibility, if the private key
can be exported, if user input is required when using a private key, if an alternate
signature algorithm should be used, and if certificates are renewed using the same private
key.

_Required_: Yes

_Type_: [PrivateKeyFlagsV4](aws-properties-pcaconnectorad-template-privatekeyflagsv4.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubjectNameFlags`

Subject name flags describe the subject name and subject alternate name that is included
in a certificate.

_Required_: Yes

_Type_: [SubjectNameFlagsV4](aws-properties-pcaconnectorad-template-subjectnameflagsv4.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupersededTemplates`

List of templates in Active Directory that are superseded by this template.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `64 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TemplateV3

ValidityPeriod

All content copied from https://docs.aws.amazon.com/.
