This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate ApiPassthrough

Contains X.509 certificate information to be placed in an issued certificate. An
`APIPassthrough` or `APICSRPassthrough` template variant must
be selected, or else this parameter is ignored.

If conflicting or duplicate certificate information is supplied from other sources,
AWS Private CA applies [order of\
operation rules](https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html#template-order-of-operations) to determine what information is used.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Extensions" : Extensions,
  "Subject" : Subject
}

```

### YAML

```yaml

  Extensions:
    Extensions
  Subject:
    Subject

```

## Properties

`Extensions`

Specifies X.509 extension information for a certificate.

_Required_: No

_Type_: [Extensions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-acmpca-certificate-extensions.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subject`

Contains information about the certificate subject. The Subject field in the
certificate identifies the entity that owns or controls the public key in the
certificate. The entity can be a user, computer, device, or service. The Subject must
contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished
names (RDNs). The RDNs are separated by commas in the certificate.

_Required_: No

_Type_: [Subject](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-acmpca-certificate-subject.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ACMPCA::Certificate

CustomAttribute
