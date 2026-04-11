This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup RulesSource

The stateless or stateful rules definitions for use in a single rule group. Each rule
group requires a single `RulesSource`. You can use an instance of this for
either stateless rules or stateful rules.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RulesSourceList" : RulesSourceList,
  "RulesString" : String,
  "StatefulRules" : [ StatefulRule, ... ],
  "StatelessRulesAndCustomActions" : StatelessRulesAndCustomActions
}

```

### YAML

```yaml

  RulesSourceList:
    RulesSourceList
  RulesString:
    String
  StatefulRules:
    - StatefulRule
  StatelessRulesAndCustomActions:
    StatelessRulesAndCustomActions

```

## Properties

`RulesSourceList`

Stateful inspection criteria for a domain list rule group.

_Required_: No

_Type_: [RulesSourceList](aws-properties-networkfirewall-rulegroup-rulessourcelist.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RulesString`

Stateful inspection criteria, provided in Suricata compatible rules. Suricata is an open-source threat detection framework that includes a standard
rule-based language for network traffic inspection.

These rules contain the inspection criteria and the action to take for traffic that
matches the criteria, so this type of rule group doesn't have a separate action
setting.

###### Note

You can't use the `priority` keyword if the `RuleOrder` option in StatefulRuleOptions is set to `STRICT_ORDER`.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatefulRules`

An array of individual stateful rules inspection criteria to be used together in a stateful rule group.
Use this option to specify simple Suricata rules with protocol, source and destination, ports, direction, and rule options.
For information about the Suricata `Rules` format, see
[Rules Format](https://suricata.readthedocs.io/en/suricata-7.0.3/rules/intro.html).

_Required_: No

_Type_: Array of [StatefulRule](aws-properties-networkfirewall-rulegroup-statefulrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatelessRulesAndCustomActions`

Stateless inspection criteria to be used in a stateless rule group.

_Required_: No

_Type_: [StatelessRulesAndCustomActions](aws-properties-networkfirewall-rulegroup-statelessrulesandcustomactions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleOption

RulesSourceList

All content copied from https://docs.aws.amazon.com/.
