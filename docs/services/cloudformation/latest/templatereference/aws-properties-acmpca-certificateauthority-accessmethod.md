This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::CertificateAuthority AccessMethod

Describes the type and format of extension access. Only one of
`CustomObjectIdentifier` or `AccessMethodType` may be
provided. Providing both results in `InvalidArgsException`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessMethodType" : String,
  "CustomObjectIdentifier" : String
}

```

### YAML

```yaml

  AccessMethodType: String
  CustomObjectIdentifier: String

```

## Properties

`AccessMethodType`

Specifies the `AccessMethod`.

_Required_: No

_Type_: String

_Allowed values_: `CA_REPOSITORY | RESOURCE_PKI_MANIFEST | RESOURCE_PKI_NOTIFY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomObjectIdentifier`

An object identifier (OID) specifying the `AccessMethod`. The OID must
satisfy the regular expression shown below. For more information, see NIST's definition
of [Object Identifier\
(OID)](https://csrc.nist.gov/glossary/term/Object_Identifier).

_Required_: No

_Type_: String

_Pattern_: `([0-2])\.([0-9]|([0-3][0-9]))((\.([0-9]+)){0,126})`

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessDescription

CrlConfiguration

All content copied from https://docs.aws.amazon.com/.
