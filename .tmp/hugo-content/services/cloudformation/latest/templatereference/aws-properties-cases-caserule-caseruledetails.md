This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::CaseRule CaseRuleDetails

Represents what rule type should take place, under what conditions.
In the Amazon Connect admin website, case rules are known as _case field conditions_.
For more
information about case field conditions, see [Add case field conditions to a\
case template](../../../connect/latest/adminguide/case-field-conditions.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Hidden" : HiddenCaseRule,
  "Required" : RequiredCaseRule
}

```

### YAML

```yaml

  Hidden:
    HiddenCaseRule
  Required:
    RequiredCaseRule

```

## Properties

`Hidden`

Whether a field is visible, based on values in other fields.

_Required_: No

_Type_: [HiddenCaseRule](aws-properties-cases-caserule-hiddencaserule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Required`

Required rule type, used to indicate whether a field is required.

_Required_: No

_Type_: [RequiredCaseRule](aws-properties-cases-caserule-requiredcaserule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BooleanOperands

HiddenCaseRule

All content copied from https://docs.aws.amazon.com/.
