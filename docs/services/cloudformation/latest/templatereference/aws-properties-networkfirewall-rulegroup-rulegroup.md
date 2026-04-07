This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup RuleGroup

The object that defines the rules in a rule group.

AWS Network Firewall uses a rule group to inspect and control network traffic.
You define stateless rule groups to inspect individual packets and you define stateful rule groups to inspect packets in the context of their
traffic flow.

To use a rule group, you include it by reference in an Network Firewall firewall policy, then you use the policy in a firewall. You can reference a rule group from
more than one firewall policy, and you can use a firewall policy in more than one firewall.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReferenceSets" : ReferenceSets,
  "RulesSource" : RulesSource,
  "RuleVariables" : RuleVariables,
  "StatefulRuleOptions" : StatefulRuleOptions
}

```

### YAML

```yaml

  ReferenceSets:
    ReferenceSets
  RulesSource:
    RulesSource
  RuleVariables:
    RuleVariables
  StatefulRuleOptions:
    StatefulRuleOptions

```

## Properties

`ReferenceSets`

The reference sets for the stateful rule group.

_Required_: No

_Type_: [ReferenceSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkfirewall-rulegroup-referencesets.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RulesSource`

The stateful rules or stateless rules for the rule group.

_Required_: Yes

_Type_: [RulesSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkfirewall-rulegroup-rulessource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleVariables`

Settings that are available for use in the rules in the rule group. You can only use
these for stateful rule groups.

_Required_: No

_Type_: [RuleVariables](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkfirewall-rulegroup-rulevariables.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatefulRuleOptions`

Additional options governing how Network Firewall handles stateful rules. The policies where you use your stateful
rule group must have stateful rule options settings that are compatible with these settings. Some limitations apply; for more information, see [Strict evaluation order](https://docs.aws.amazon.com/network-firewall/latest/developerguide/suricata-limitations-caveats.html) in the _AWS Network Firewall Developer Guide_.

_Required_: No

_Type_: [StatefulRuleOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkfirewall-rulegroup-statefulruleoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RuleDefinition

RuleOption
