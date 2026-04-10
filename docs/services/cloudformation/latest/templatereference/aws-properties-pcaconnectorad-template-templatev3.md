This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template TemplateV3

v3 template schema that uses Key Storage Providers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateValidity" : CertificateValidity,
  "EnrollmentFlags" : EnrollmentFlagsV3,
  "Extensions" : ExtensionsV3,
  "GeneralFlags" : GeneralFlagsV3,
  "HashAlgorithm" : String,
  "PrivateKeyAttributes" : PrivateKeyAttributesV3,
  "PrivateKeyFlags" : PrivateKeyFlagsV3,
  "SubjectNameFlags" : SubjectNameFlagsV3,
  "SupersededTemplates" : [ String, ... ]
}

```

### YAML

```yaml

  CertificateValidity:
    CertificateValidity
  EnrollmentFlags:
    EnrollmentFlagsV3
  Extensions:
    ExtensionsV3
  GeneralFlags:
    GeneralFlagsV3
  HashAlgorithm: String
  PrivateKeyAttributes:
    PrivateKeyAttributesV3
  PrivateKeyFlags:
    PrivateKeyFlagsV3
  SubjectNameFlags:
    SubjectNameFlagsV3
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

Enrollment flags describe the enrollment settings for certificates such as using the
existing private key and deleting expired or revoked certificates.

_Required_: Yes

_Type_: [EnrollmentFlagsV3](aws-properties-pcaconnectorad-template-enrollmentflagsv3.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Extensions`

Extensions describe the key usage extensions and application policies for a
template.

_Required_: Yes

_Type_: [ExtensionsV3](aws-properties-pcaconnectorad-template-extensionsv3.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GeneralFlags`

General flags describe whether the template is used for computers or users and if the
template can be used with autoenrollment.

_Required_: Yes

_Type_: [GeneralFlagsV3](aws-properties-pcaconnectorad-template-generalflagsv3.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HashAlgorithm`

Specifies the hash algorithm used to hash the private key.

_Required_: Yes

_Type_: String

_Allowed values_: `SHA256 | SHA384 | SHA512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateKeyAttributes`

Private key attributes allow you to specify the algorithm, minimal key length, key spec,
key usage, and cryptographic providers for the private key of a certificate for v3
templates. V3 templates allow you to use Key Storage Providers.

_Required_: Yes

_Type_: [PrivateKeyAttributesV3](aws-properties-pcaconnectorad-template-privatekeyattributesv3.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateKeyFlags`

Private key flags for v3 templates specify the client compatibility, if the private key
can be exported, if user input is required when using a private key, and if an alternate
signature algorithm should be used.

_Required_: Yes

_Type_: [PrivateKeyFlagsV3](aws-properties-pcaconnectorad-template-privatekeyflagsv3.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubjectNameFlags`

Subject name flags describe the subject name and subject alternate name that is included
in a certificate.

_Required_: Yes

_Type_: [SubjectNameFlagsV3](aws-properties-pcaconnectorad-template-subjectnameflagsv3.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupersededTemplates`

List of templates in Active Directory that are superseded by this template.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `64 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TemplateV2

TemplateV4

All content copied from https://docs.aws.amazon.com/.
