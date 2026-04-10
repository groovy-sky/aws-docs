This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Transformer X12ValidationOptions

Contains configuration options for X12 EDI validation. This structure allows you to
specify custom validation rules that will be applied during EDI document processing,
including element length constraints, code list modifications, and element requirement
changes. These validation options provide flexibility to accommodate trading
partner-specific requirements while maintaining EDI compliance. The validation rules are
applied in addition to standard X12 validation to ensure documents meet both standard and
custom requirements.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ValidationRules" : [ X12ValidationRule, ... ]
}

```

### YAML

```yaml

  ValidationRules:
    - X12ValidationRule

```

## Properties

`ValidationRules`

Specifies a list of validation rules to apply during EDI document processing. These
rules can include code list modifications, element length constraints, and element
requirement changes.

_Required_: No

_Type_: Array of [X12ValidationRule](aws-properties-b2bi-transformer-x12validationrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12SplitOptions

X12ValidationRule

All content copied from https://docs.aws.amazon.com/.
