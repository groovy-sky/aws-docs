This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL RuleActionOverride

Action setting to use in the place of a rule action that is configured inside the rule
group. You specify one override for each rule whose action you want to change.

You can use overrides for testing, for example you can override all of rule actions to
`Count` and then monitor the resulting count metrics to understand how the
rule group would handle your web traffic. You can also permanently override some or all
actions, to modify how the rule group manages your web traffic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionToUse" : RuleAction,
  "Name" : String
}

```

### YAML

```yaml

  ActionToUse:
    RuleAction
  Name: String

```

## Properties

`ActionToUse`

The override action to use, in place of the configured action of the rule in the rule
group.

_Required_: Yes

_Type_: [RuleAction](aws-properties-wafv2-webacl-ruleaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the rule to override.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z_-]{1,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleAction

RuleGroupReferenceStatement

All content copied from https://docs.aws.amazon.com/.
