This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate CustomExtension

Specifies the X.509 extension information for a certificate.

Extensions present in `CustomExtensions` follow the
`ApiPassthrough` [template\
rules](../../../privateca/latest/userguide/usingtemplates.md#template-order-of-operations).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Critical" : Boolean,
  "ObjectIdentifier" : String,
  "Value" : String
}

```

### YAML

```yaml

  Critical: Boolean
  ObjectIdentifier: String
  Value: String

```

## Properties

`Critical`

Specifies the critical flag of the X.509 extension.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ObjectIdentifier`

Specifies the object identifier (OID) of the X.509 extension. For more information,
see the [Global OID reference database.](https://oidref.com/2.5.29)

_Required_: Yes

_Type_: String

_Pattern_: `([0-2])\.([0-9]|([0-3][0-9]))((\.([0-9]+)){0,126})`

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

Specifies the base64-encoded value of the X.509 extension.

_Required_: Yes

_Type_: String

_Pattern_: `(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?`

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomAttribute

EdiPartyName

All content copied from https://docs.aws.amazon.com/.
