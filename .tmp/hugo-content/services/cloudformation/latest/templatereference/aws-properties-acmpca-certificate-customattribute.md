This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate CustomAttribute

Defines the X.500 relative distinguished name (RDN).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ObjectIdentifier" : String,
  "Value" : String
}

```

### YAML

```yaml

  ObjectIdentifier: String
  Value: String

```

## Properties

`ObjectIdentifier`

Specifies the object identifier (OID) of the attribute type of the relative
distinguished name (RDN).

_Required_: Yes

_Type_: String

_Pattern_: `([0-2])\.([0-9]|([0-3][0-9]))((\.([0-9]+)){0,126})`

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

Specifies the attribute value of relative distinguished name (RDN).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiPassthrough

CustomExtension

All content copied from https://docs.aws.amazon.com/.
