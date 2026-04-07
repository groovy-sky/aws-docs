This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::Template TemplateRule

An association representing a case rule acting upon a field. In the Amazon Connect admin website, case rules are known as _case field conditions_.
For more
information about case field conditions, see [Add case field conditions to a\
case template](https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CaseRuleId" : String,
  "FieldId" : String
}

```

### YAML

```yaml

  CaseRuleId: String
  FieldId: String

```

## Properties

`CaseRuleId`

Unique identifier of a case rule.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldId`

Unique identifier of a field.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
