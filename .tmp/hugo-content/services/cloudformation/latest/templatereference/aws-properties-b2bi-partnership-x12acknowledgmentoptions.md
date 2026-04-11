This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Partnership X12AcknowledgmentOptions

Contains options for configuring X12 acknowledgments. These options control how
functional and technical acknowledgments are handled.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FunctionalAcknowledgment" : String,
  "TechnicalAcknowledgment" : String
}

```

### YAML

```yaml

  FunctionalAcknowledgment: String
  TechnicalAcknowledgment: String

```

## Properties

`FunctionalAcknowledgment`

Specifies whether functional acknowledgments (997/999) should be generated for incoming
X12 transactions. Valid values are `DO_NOT_GENERATE`,
`GENERATE_ALL_SEGMENTS` and
`GENERATE_WITHOUT_TRANSACTION_SET_RESPONSE_LOOP`.

If you choose `GENERATE_WITHOUT_TRANSACTION_SET_RESPONSE_LOOP`, AWS B2B Data Interchange
skips the AK2\_Loop when generating an acknowledgment document.

_Required_: Yes

_Type_: String

_Allowed values_: `DO_NOT_GENERATE | GENERATE_ALL_SEGMENTS | GENERATE_WITHOUT_TRANSACTION_SET_RESPONSE_LOOP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TechnicalAcknowledgment`

Specifies whether technical acknowledgments (TA1) should be generated for incoming X12
interchanges. Valid values are `DO_NOT_GENERATE` and
`GENERATE_ALL_SEGMENTS` and.

_Required_: Yes

_Type_: String

_Allowed values_: `DO_NOT_GENERATE | GENERATE_ALL_SEGMENTS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WrapOptions

X12ControlNumbers

All content copied from https://docs.aws.amazon.com/.
