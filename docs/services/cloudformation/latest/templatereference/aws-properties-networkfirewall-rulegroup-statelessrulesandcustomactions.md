This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup StatelessRulesAndCustomActions

Stateless inspection criteria. Each stateless rule group uses exactly one of these data
types to define its stateless rules.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomActions" : [ CustomAction, ... ],
  "StatelessRules" : [ StatelessRule, ... ]
}

```

### YAML

```yaml

  CustomActions:
    - CustomAction
  StatelessRules:
    - StatelessRule

```

## Properties

`CustomActions`

Defines an array of individual custom action definitions that are available for use by
the stateless rules in this `StatelessRulesAndCustomActions` specification. You
name each custom action that you define, and then you can use it by name in your stateless rule definition
`Actions` specification.

_Required_: No

_Type_: Array of [CustomAction](aws-properties-networkfirewall-rulegroup-customaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatelessRules`

Defines the set of stateless rules for use in a stateless rule group.

_Required_: Yes

_Type_: Array of [StatelessRule](aws-properties-networkfirewall-rulegroup-statelessrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StatelessRule

SummaryConfiguration

All content copied from https://docs.aws.amazon.com/.
