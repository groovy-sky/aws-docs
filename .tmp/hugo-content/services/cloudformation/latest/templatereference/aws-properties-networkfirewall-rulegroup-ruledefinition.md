This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup RuleDefinition

The inspection criteria and action for a single stateless rule. AWS Network Firewall inspects each packet for the specified matching
criteria. When a packet matches the criteria, Network Firewall performs the rule's actions on
the packet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ String, ... ],
  "MatchAttributes" : MatchAttributes
}

```

### YAML

```yaml

  Actions:
    - String
  MatchAttributes:
    MatchAttributes

```

## Properties

`Actions`

The actions to take on a packet that matches one of the stateless rule definition's
match attributes. You must specify a standard action and you can add custom actions.

###### Note

Network Firewall only forwards a packet for stateful rule inspection if you specify
`aws:forward_to_sfe` for a rule that the packet matches, or if the packet
doesn't match any stateless rule and you specify `aws:forward_to_sfe` for the
`StatelessDefaultActions` setting for the firewall policy.

For every rule, you must specify exactly one of the following standard actions.

- **aws:pass** \- Discontinues all inspection of
the packet and permits it to go to its intended destination.

- **aws:drop** \- Discontinues all inspection of
the packet and blocks it from going to its intended destination.

- **aws:forward\_to\_sfe** \- Discontinues
stateless inspection of the packet and forwards it to the stateful rule engine for
inspection.

Additionally, you can specify a custom action.
To do this, you define a custom action by name and type, then provide the name you've assigned
to the action in this `Actions` setting.

To provide more than one action in this setting, separate the settings with a comma. For
example, if you have a publish metrics custom action that you've named
`MyMetricsAction`, then you could specify the standard action
`aws:pass` combined with the custom action using `[“aws:pass”,
            “MyMetricsAction”]`.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchAttributes`

Criteria for Network Firewall to use to inspect an individual packet in stateless rule inspection. Each match attributes set can include one or more items such as IP address, CIDR range, port number, protocol, and TCP flags.

_Required_: Yes

_Type_: [MatchAttributes](aws-properties-networkfirewall-rulegroup-matchattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReferenceSets

RuleGroup

All content copied from https://docs.aws.amazon.com/.
