This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Transformer X12AdvancedOptions

Contains advanced options specific to X12 EDI processing, such as splitting large X12
files into smaller units.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SplitOptions" : X12SplitOptions,
  "ValidationOptions" : X12ValidationOptions
}

```

### YAML

```yaml

  SplitOptions:
    X12SplitOptions
  ValidationOptions:
    X12ValidationOptions

```

## Properties

`SplitOptions`

Specifies options for splitting X12 EDI files. These options control how large X12 files
are divided into smaller, more manageable units.

_Required_: No

_Type_: [X12SplitOptions](aws-properties-b2bi-transformer-x12splitoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidationOptions`

Specifies validation options for X12 EDI processing. These options control how
validation rules are applied during EDI document processing, including custom validation
rules for element length constraints, code list validations, and element requirement
checks.

_Required_: No

_Type_: [X12ValidationOptions](aws-properties-b2bi-transformer-x12validationoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

X12CodeListValidationRule

All content copied from https://docs.aws.amazon.com/.
