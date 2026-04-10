This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Partnership X12OutboundEdiHeaders

A structure containing the details for an outbound EDI object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ControlNumbers" : X12ControlNumbers,
  "Delimiters" : X12Delimiters,
  "FunctionalGroupHeaders" : X12FunctionalGroupHeaders,
  "Gs05TimeFormat" : String,
  "InterchangeControlHeaders" : X12InterchangeControlHeaders,
  "ValidateEdi" : Boolean
}

```

### YAML

```yaml

  ControlNumbers:
    X12ControlNumbers
  Delimiters:
    X12Delimiters
  FunctionalGroupHeaders:
    X12FunctionalGroupHeaders
  Gs05TimeFormat: String
  InterchangeControlHeaders:
    X12InterchangeControlHeaders
  ValidateEdi: Boolean

```

## Properties

`ControlNumbers`

Specifies control number configuration for outbound X12 EDI headers. These settings
determine the starting values for interchange, functional group, and transaction set
control numbers.

_Required_: No

_Type_: [X12ControlNumbers](aws-properties-b2bi-partnership-x12controlnumbers.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Delimiters`

The delimiters, for example semicolon ( `;`), that separates sections of the
headers for the X12 object.

_Required_: No

_Type_: [X12Delimiters](aws-properties-b2bi-partnership-x12delimiters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunctionalGroupHeaders`

The functional group headers for the X12 object.

_Required_: No

_Type_: [X12FunctionalGroupHeaders](aws-properties-b2bi-partnership-x12functionalgroupheaders.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Gs05TimeFormat`

Specifies the time format in the GS05 element (time) of the functional group header. The
following formats use 24-hour clock time:

- `HHMM` \- Hours and minutes

- `HHMMSS` \- Hours, minutes, and seconds

- `HHMMSSDD` \- Hours, minutes, seconds, and decimal seconds

Where:

- `HH` \- Hours (00-23)

- `MM` \- Minutes (00-59)

- `SS` \- Seconds (00-59)

- `DD` \- Hundredths of seconds (00-99)

_Required_: No

_Type_: String

_Allowed values_: `HHMM | HHMMSS | HHMMSSDD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InterchangeControlHeaders`

In X12 EDI messages, delimiters are used to mark the end of segments or elements, and
are defined in the interchange control header.

_Required_: No

_Type_: [X12InterchangeControlHeaders](aws-properties-b2bi-partnership-x12interchangecontrolheaders.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidateEdi`

Specifies whether or not to validate the EDI for this X12 object: `TRUE` or
`FALSE`. When enabled, this performs both standard EDI validation and applies
any configured custom validation rules including element length constraints, code list
validations, and element requirement checks. Validation results are returned in the
response validation messages.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12InterchangeControlHeaders

AWS::B2BI::Profile

All content copied from https://docs.aws.amazon.com/.
