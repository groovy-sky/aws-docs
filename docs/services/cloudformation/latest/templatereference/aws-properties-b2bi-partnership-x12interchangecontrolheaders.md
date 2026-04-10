This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Partnership X12InterchangeControlHeaders

In X12, the Interchange Control Header is the first segment of an EDI document and is
part of the Interchange Envelope. It contains information about the sender and receiver,
the date and time of transmission, and the X12 version being used. It also includes
delivery information, such as the sender and receiver IDs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AcknowledgmentRequestedCode" : String,
  "ReceiverId" : String,
  "ReceiverIdQualifier" : String,
  "RepetitionSeparator" : String,
  "SenderId" : String,
  "SenderIdQualifier" : String,
  "UsageIndicatorCode" : String
}

```

### YAML

```yaml

  AcknowledgmentRequestedCode: String
  ReceiverId: String
  ReceiverIdQualifier: String
  RepetitionSeparator: String
  SenderId: String
  SenderIdQualifier: String
  UsageIndicatorCode: String

```

## Properties

`AcknowledgmentRequestedCode`

Located at position ISA-14 in the header. The value "1" indicates that the sender is
requesting an interchange acknowledgment at receipt of the interchange. The value "0" is
used otherwise.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReceiverId`

Located at position ISA-08 in the header. This value (along with the
`receiverIdQualifier`) identifies the intended recipient of the interchange.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9 ]*$`

_Minimum_: `15`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReceiverIdQualifier`

Located at position ISA-07 in the header. Qualifier for the receiver ID. Together, the
ID and qualifier uniquely identify the receiving trading partner.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]*$`

_Minimum_: `2`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepetitionSeparator`

Located at position ISA-11 in the header. This string makes it easier when you need to
group similar adjacent element values together without using extra segments.

###### Note

This parameter is only honored for version greater than 401
( `VERSION_4010` and higher).

For versions less than 401, this field is called [StandardsId](https://www.stedi.com/edi/x12-004010/segment/ISA), in
which case our service sets the value to `U`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SenderId`

Located at position ISA-06 in the header. This value (along with the
`senderIdQualifier`) identifies the sender of the interchange.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9 ]*$`

_Minimum_: `15`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SenderIdQualifier`

Located at position ISA-05 in the header. Qualifier for the sender ID. Together, the ID
and qualifier uniquely identify the sending trading partner.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]*$`

_Minimum_: `2`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsageIndicatorCode`

Located at position ISA-15 in the header. Specifies how this interchange is being
used:

- `T` indicates this interchange is for testing.

- `P` indicates this interchange is for production.

- `I` indicates this interchange is informational.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12InboundEdiOptions

X12OutboundEdiHeaders

All content copied from https://docs.aws.amazon.com/.
