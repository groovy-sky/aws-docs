This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Partnership X12ControlNumbers

Contains configuration for X12 control numbers used in X12 EDI generation. Control
numbers are used to uniquely identify interchanges, functional groups, and transaction
sets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StartingFunctionalGroupControlNumber" : Number,
  "StartingInterchangeControlNumber" : Number,
  "StartingTransactionSetControlNumber" : Number
}

```

### YAML

```yaml

  StartingFunctionalGroupControlNumber:
    Number
  StartingInterchangeControlNumber:
    Number
  StartingTransactionSetControlNumber:
    Number

```

## Properties

`StartingFunctionalGroupControlNumber`

Specifies the starting functional group control number (GS06) to use for X12 EDI
generation. This number is incremented for each new functional group. For the GS
(functional group) envelope, AWS B2B Data Interchange generates a functional
group control number that is unique to the sender ID, receiver ID, and functional
identifier code combination.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `999999999`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartingInterchangeControlNumber`

Specifies the starting interchange control number (ISA13) to use for X12 EDI generation.
This number is incremented for each new interchange. For the ISA (interchange) envelope,
AWS B2B Data Interchange generates an interchange control number that is
unique for the ISA05 and ISA06 (sender) & ISA07 and ISA08 (receiver) combination.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `999999999`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartingTransactionSetControlNumber`

Specifies the starting transaction set control number (ST02) to use for X12 EDI
generation. This number is incremented for each new transaction set.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `999999999`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12AcknowledgmentOptions

X12Delimiters

All content copied from https://docs.aws.amazon.com/.
