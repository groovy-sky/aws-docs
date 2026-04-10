This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::LoggingConfiguration Condition

A single match condition for a log filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionCondition" : ActionCondition,
  "LabelNameCondition" : LabelNameCondition
}

```

### YAML

```yaml

  ActionCondition:
    ActionCondition
  LabelNameCondition:
    LabelNameCondition

```

## Properties

`ActionCondition`

A single action condition. This is the action setting that a log record must contain in order to meet the condition.

_Required_: No

_Type_: [ActionCondition](aws-properties-wafv2-loggingconfiguration-actioncondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LabelNameCondition`

A single label name condition. This is the fully qualified label name that a log record must contain in order to meet the condition.
Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label.

_Required_: No

_Type_: [LabelNameCondition](aws-properties-wafv2-loggingconfiguration-labelnamecondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionCondition

FieldToMatch

All content copied from https://docs.aws.amazon.com/.
