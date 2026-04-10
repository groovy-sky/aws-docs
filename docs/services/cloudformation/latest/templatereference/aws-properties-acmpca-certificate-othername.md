This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate OtherName

Defines a custom ASN.1 X.400 `GeneralName` using an object identifier (OID)
and value. The OID must satisfy the regular expression shown below. For more
information, see NIST's definition of [Object Identifier\
(OID)](https://csrc.nist.gov/glossary/term/Object_Identifier).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TypeId" : String,
  "Value" : String
}

```

### YAML

```yaml

  TypeId: String
  Value: String

```

## Properties

`TypeId`

Specifies an OID.

_Required_: Yes

_Type_: String

_Pattern_: `([0-2])\.([0-9]|([0-3][0-9]))((\.([0-9]+)){0,126})`

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

Specifies an OID value.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeyUsage

PolicyInformation

All content copied from https://docs.aws.amazon.com/.
