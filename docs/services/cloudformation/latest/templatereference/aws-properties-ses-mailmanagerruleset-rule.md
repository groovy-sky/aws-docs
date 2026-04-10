This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet Rule

A rule contains conditions, "unless conditions" and actions. For each
envelope recipient of an email, if all conditions match and none of the "unless
conditions" match, then all of the actions are executed sequentially. If no
conditions are provided, the rule always applies and the actions are implicitly
executed. If only "unless conditions" are provided, the rule applies if the
email does not match the evaluation of the "unless conditions".

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ RuleAction, ... ],
  "Conditions" : [ RuleCondition, ... ],
  "Name" : String,
  "Unless" : [ RuleCondition, ... ]
}

```

### YAML

```yaml

  Actions:
    - RuleAction
  Conditions:
    - RuleCondition
  Name: String
  Unless:
    - RuleCondition

```

## Properties

`Actions`

The list of actions to execute when the conditions match the incoming email, and none
of the "unless conditions" match.

_Required_: Yes

_Type_: Array of [RuleAction](aws-properties-ses-mailmanagerruleset-ruleaction.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Conditions`

The conditions of this rule. All conditions must match the email for the actions to be
executed. An empty list of conditions means that all emails match, but are still subject
to any "unless conditions"

_Required_: No

_Type_: Array of [RuleCondition](aws-properties-ses-mailmanagerruleset-rulecondition.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The user-friendly name of the rule.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unless`

The "unless conditions" of this rule. None of the conditions can match the
email for the actions to be executed. If any of these conditions do match the email,
then the actions are not executed.

_Required_: No

_Type_: Array of [RuleCondition](aws-properties-ses-mailmanagerruleset-rulecondition.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplaceRecipientAction

RuleAction

All content copied from https://docs.aws.amazon.com/.
