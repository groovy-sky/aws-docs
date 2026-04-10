This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup StatefulRuleOptions

Additional options governing how Network Firewall handles the rule group. You can only use these for stateful rule groups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RuleOrder" : String
}

```

### YAML

```yaml

  RuleOrder: String

```

## Properties

`RuleOrder`

Indicates how to manage the order of the rule evaluation for the rule group. `DEFAULT_ACTION_ORDER` is
the default behavior. Stateful rules are provided to the rule engine as Suricata compatible strings, and Suricata evaluates them
based on certain settings. For more information, see
[Evaluation order for stateful rules](../../../network-firewall/latest/developerguide/suricata-rule-evaluation-order.md) in the _AWS Network Firewall Developer Guide_.

_Required_: No

_Type_: String

_Allowed values_: `DEFAULT_ACTION_ORDER | STRICT_ORDER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StatefulRule

StatelessRule

All content copied from https://docs.aws.amazon.com/.
