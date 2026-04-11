This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Partnership X12FunctionalGroupHeaders

Part of the X12 message structure. These are the functional group headers for the X12
EDI object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationReceiverCode" : String,
  "ApplicationSenderCode" : String,
  "ResponsibleAgencyCode" : String
}

```

### YAML

```yaml

  ApplicationReceiverCode: String
  ApplicationSenderCode: String
  ResponsibleAgencyCode: String

```

## Properties

`ApplicationReceiverCode`

A value representing the code used to identify the party receiving a message, at
position GS-03.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9 ]*$`

_Minimum_: `2`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationSenderCode`

A value representing the code used to identify the party transmitting a message, at
position GS-02.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9 ]*$`

_Minimum_: `2`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponsibleAgencyCode`

A code that identifies the issuer of the standard, at position GS-07.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12Envelope

X12InboundEdiOptions

All content copied from https://docs.aws.amazon.com/.
