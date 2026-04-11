This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL RuleGroupReferenceStatement

A rule statement used to run the rules that are defined in a [AWS::WAFv2::RuleGroup](aws-resource-wafv2-rulegroup.md). To use this, create a rule group with your rules, then provide the ARN of the rule group in this statement.

You cannot nest a `RuleGroupReferenceStatement`, for example for use inside a `NotStatement` or `OrStatement`. You cannot use a rule group
reference statement inside another rule group. You can only reference a rule group as a top-level statement within a rule that you define in a web ACL.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "ExcludedRules" : [ ExcludedRule, ... ],
  "RuleActionOverrides" : [ RuleActionOverride, ... ]
}

```

### YAML

```yaml

  Arn: String
  ExcludedRules:
    - ExcludedRule
  RuleActionOverrides:
    - RuleActionOverride

```

## Properties

`Arn`

The Amazon Resource Name (ARN) of the entity.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludedRules`

Rules in the referenced rule group whose actions are set to `Count`.

###### Note

Instead of this option, use `RuleActionOverrides`. It accepts any valid action setting, including `Count`.

_Required_: No

_Type_: Array of [ExcludedRule](aws-properties-wafv2-webacl-excludedrule.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleActionOverrides`

Action settings to use in the place of the rule actions that are configured inside the rule group. You specify one override for each rule whose action you want to change.

###### Note

Verify the rule names in your overrides carefully. With managed rule groups, AWS WAF silently ignores any override that uses an invalid rule name. With customer-owned rule groups, invalid rule names in your overrides will cause web ACL updates to fail. An invalid rule name is any name that doesn't exactly match the case-sensitive name of an existing rule in the rule group.

You can use overrides for testing, for example you can override all of rule actions to `Count` and then monitor the resulting count metrics to understand how the rule group would handle your web traffic. You can also permanently override some or all actions, to modify how the rule group manages your web traffic.

_Required_: No

_Type_: Array of [RuleActionOverride](aws-properties-wafv2-webacl-ruleactionoverride.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleActionOverride

SingleHeader

All content copied from https://docs.aws.amazon.com/.
