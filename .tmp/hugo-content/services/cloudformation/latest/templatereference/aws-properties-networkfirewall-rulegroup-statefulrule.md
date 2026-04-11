This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup StatefulRule

A single Suricata rules specification, for use in a stateful rule group.
Use this option to specify a simple Suricata rule with protocol, source and destination, ports, direction, and rule options.
For information about the Suricata `Rules` format, see
[Rules Format](https://suricata.readthedocs.io/en/suricata-7.0.3/rules/intro.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "Header" : Header,
  "RuleOptions" : [ RuleOption, ... ]
}

```

### YAML

```yaml

  Action: String
  Header:
    Header
  RuleOptions:
    - RuleOption

```

## Properties

`Action`

Defines what Network Firewall should do with the packets in a traffic flow when the flow
matches the stateful rule criteria. For all actions, Network Firewall performs the specified
action and discontinues stateful inspection of the traffic flow.

The actions for a stateful rule are defined as follows:

- **PASS** \- Permits the packets to go to the
intended destination.

- **DROP** \- Blocks the packets from going to
the intended destination and sends an alert log message, if alert logging is configured in the firewall
logging configuration.

- **REJECT** \- Drops traffic that matches the conditions of the stateful rule and sends a TCP reset packet back to sender of the packet. A TCP reset packet is a packet with no payload and a `RST` bit contained in the TCP header flags. `REJECT` is available only for TCP traffic.

- **ALERT** \- Permits the packets to go to the
intended destination and sends an alert log message, if alert logging is configured in the firewall
logging configuration.

You can use this action to test a rule that you intend to use to drop traffic. You
can enable the rule with `ALERT` action, verify in the logs that the rule
is filtering as you want, then change the action to `DROP`.

- **REJECT** \- Drops TCP traffic that matches the conditions of the stateful rule, and sends a TCP reset packet back to sender of the packet. A TCP reset packet is a packet with no payload and a `RST` bit contained in the TCP header flags. Also sends an alert log mesage if alert logging is configured in the firewall
logging configuration.

`REJECT` isn't currently available for use with IMAP and FTP protocols.

_Required_: Yes

_Type_: String

_Allowed values_: `PASS | DROP | ALERT | REJECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Header`

The stateful inspection criteria for this rule, used to inspect traffic flows.

_Required_: Yes

_Type_: [Header](aws-properties-networkfirewall-rulegroup-header.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleOptions`

Additional settings for a stateful rule, provided as keywords and settings.

_Required_: Yes

_Type_: Array of [RuleOption](aws-properties-networkfirewall-rulegroup-ruleoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleVariables

StatefulRuleOptions

All content copied from https://docs.aws.amazon.com/.
