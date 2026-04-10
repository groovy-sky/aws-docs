This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup StatelessRule

A single stateless rule. This is used in `StatelessRulesAndCustomActions`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Priority" : Integer,
  "RuleDefinition" : RuleDefinition
}

```

### YAML

```yaml

  Priority: Integer
  RuleDefinition:
    RuleDefinition

```

## Properties

`Priority`

Indicates the order in which to run this rule relative to all of the
rules that are defined for a stateless rule group. Network Firewall evaluates the rules in a
rule group starting with the lowest priority setting. You must ensure that the priority
settings are unique for the rule group.

Each stateless rule group uses exactly one `StatelessRulesAndCustomActions`
object, and each `StatelessRulesAndCustomActions` contains exactly one
`StatelessRules` object. To ensure unique priority settings for your rule
groups, set unique priorities for the stateless rules that you define inside any single
`StatelessRules` object.

You can change the priority settings of your rules at any time. To make it easier to
insert rules later, number them so there's a wide range in between, for example use 100,
200, and so on.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleDefinition`

Defines the stateless 5-tuple packet inspection criteria and the action to take on a
packet that matches the criteria.

_Required_: Yes

_Type_: [RuleDefinition](aws-properties-networkfirewall-rulegroup-ruledefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StatefulRuleOptions

StatelessRulesAndCustomActions

All content copied from https://docs.aws.amazon.com/.
