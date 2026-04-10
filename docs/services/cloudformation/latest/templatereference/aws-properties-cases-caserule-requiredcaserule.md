This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::CaseRule RequiredCaseRule

Required rule type, used to indicate whether a field is required.
In the Amazon Connect admin website, case rules are known as _case field conditions_.
For more
information about case field conditions, see [Add case field conditions to a\
case template](../../../connect/latest/adminguide/case-field-conditions.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Conditions" : [ BooleanCondition, ... ],
  "DefaultValue" : Boolean
}

```

### YAML

```yaml

  Conditions:
    - BooleanCondition
  DefaultValue: Boolean

```

## Properties

`Conditions`

List of conditions for the required rule; the first condition to evaluate to true dictates
the value of the rule.

_Required_: Yes

_Type_: Array of [BooleanCondition](aws-properties-cases-caserule-booleancondition.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultValue`

The value of the rule (that is, whether the field is required) should none of the
conditions evaluate to true.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OperandTwo

Tag

All content copied from https://docs.aws.amazon.com/.
