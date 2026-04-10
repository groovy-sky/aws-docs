This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template TemplateV2

v2 template schema that uses Legacy Cryptographic Providers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateValidity" : CertificateValidity,
  "EnrollmentFlags" : EnrollmentFlagsV2,
  "Extensions" : ExtensionsV2,
  "GeneralFlags" : GeneralFlagsV2,
  "PrivateKeyAttributes" : PrivateKeyAttributesV2,
  "PrivateKeyFlags" : PrivateKeyFlagsV2,
  "SubjectNameFlags" : SubjectNameFlagsV2,
  "SupersededTemplates" : [ String, ... ]
}

```

### YAML

```yaml

  CertificateValidity:
    CertificateValidity
  EnrollmentFlags:
    EnrollmentFlagsV2
  Extensions:
    ExtensionsV2
  GeneralFlags:
    GeneralFlagsV2
  PrivateKeyAttributes:
    PrivateKeyAttributesV2
  PrivateKeyFlags:
    PrivateKeyFlagsV2
  SubjectNameFlags:
    SubjectNameFlagsV2
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

_Type_: [EnrollmentFlagsV2](aws-properties-pcaconnectorad-template-enrollmentflagsv2.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Extensions`

Extensions describe the key usage extensions and application policies for a
template.

_Required_: Yes

_Type_: [ExtensionsV2](aws-properties-pcaconnectorad-template-extensionsv2.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GeneralFlags`

General flags describe whether the template is used for computers or users and if the
template can be used with autoenrollment.

_Required_: Yes

_Type_: [GeneralFlagsV2](aws-properties-pcaconnectorad-template-generalflagsv2.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateKeyAttributes`

Private key attributes allow you to specify the minimal key length, key spec, and
cryptographic providers for the private key of a certificate for v2 templates. V2 templates
allow you to use Legacy Cryptographic Service Providers.

_Required_: Yes

_Type_: [PrivateKeyAttributesV2](aws-properties-pcaconnectorad-template-privatekeyattributesv2.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateKeyFlags`

Private key flags for v2 templates specify the client compatibility, if the private key
can be exported, and if user input is required when using a private key.

_Required_: Yes

_Type_: [PrivateKeyFlagsV2](aws-properties-pcaconnectorad-template-privatekeyflagsv2.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubjectNameFlags`

Subject name flags describe the subject name and subject alternate name that is included
in a certificate.

_Required_: Yes

_Type_: [SubjectNameFlagsV2](aws-properties-pcaconnectorad-template-subjectnameflagsv2.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupersededTemplates`

List of templates in Active Directory that are superseded by this template.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `64 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TemplateDefinition

TemplateV3

All content copied from https://docs.aws.amazon.com/.
