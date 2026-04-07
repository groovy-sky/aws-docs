This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::CertificateAuthority CsrExtensions

Describes the certificate extensions to be added to the certificate signing request
(CSR).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KeyUsage" : KeyUsage,
  "SubjectInformationAccess" : [ AccessDescription, ... ]
}

```

### YAML

```yaml

  KeyUsage:
    KeyUsage
  SubjectInformationAccess:
    - AccessDescription

```

## Properties

`KeyUsage`

Indicates the purpose of the certificate and of the key contained in the
certificate.

_Required_: No

_Type_: [KeyUsage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-acmpca-certificateauthority-keyusage.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubjectInformationAccess`

For CA certificates, provides a path to additional information pertaining to the CA,
such as revocation and policy. For more information, see [Subject\
Information Access](https://datatracker.ietf.org/doc/html/rfc5280) in RFC 5280.

_Required_: No

_Type_: Array of [AccessDescription](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-acmpca-certificateauthority-accessdescription.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CrlDistributionPointExtensionConfiguration

CustomAttribute
