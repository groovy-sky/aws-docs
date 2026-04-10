This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate EdiPartyName

Describes an Electronic Data Interchange (EDI) entity as described in as defined in
[Subject Alternative\
Name](https://datatracker.ietf.org/doc/html/rfc5280) in RFC 5280.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NameAssigner" : String,
  "PartyName" : String
}

```

### YAML

```yaml

  NameAssigner: String
  PartyName: String

```

## Properties

`NameAssigner`

Specifies the name assigner.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PartyName`

Specifies the party name.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomExtension

ExtendedKeyUsage

All content copied from https://docs.aws.amazon.com/.
