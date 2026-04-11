This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::CertificateAuthority AccessDescription

Provides access information used by the `authorityInfoAccess` and
`subjectInfoAccess` extensions described in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessLocation" : GeneralName,
  "AccessMethod" : AccessMethod
}

```

### YAML

```yaml

  AccessLocation:
    GeneralName
  AccessMethod:
    AccessMethod

```

## Properties

`AccessLocation`

The location of `AccessDescription` information.

_Required_: Yes

_Type_: [GeneralName](aws-properties-acmpca-certificateauthority-generalname.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AccessMethod`

The type and format of `AccessDescription` information.

_Required_: Yes

_Type_: [AccessMethod](aws-properties-acmpca-certificateauthority-accessmethod.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ACMPCA::CertificateAuthority

AccessMethod

All content copied from https://docs.aws.amazon.com/.
