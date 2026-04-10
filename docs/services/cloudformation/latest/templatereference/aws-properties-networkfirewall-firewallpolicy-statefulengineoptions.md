This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::FirewallPolicy StatefulEngineOptions

Configuration settings for the handling of the stateful rule groups in a firewall policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FlowTimeouts" : FlowTimeouts,
  "RuleOrder" : String,
  "StreamExceptionPolicy" : String
}

```

### YAML

```yaml

  FlowTimeouts:
    FlowTimeouts
  RuleOrder: String
  StreamExceptionPolicy: String

```

## Properties

`FlowTimeouts`

Configures the amount of time that can pass without any traffic sent through the firewall before the firewall determines that the connection is idle.

_Required_: No

_Type_: [FlowTimeouts](aws-properties-networkfirewall-firewallpolicy-flowtimeouts.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleOrder`

Indicates how to manage the order of stateful rule evaluation for the policy. `DEFAULT_ACTION_ORDER` is
the default behavior. Stateful rules are provided to the rule engine as Suricata compatible strings, and Suricata evaluates them
based on certain settings. For more information, see
[Evaluation order for stateful rules](../../../network-firewall/latest/developerguide/suricata-rule-evaluation-order.md) in the _AWS Network Firewall Developer Guide_.

_Required_: No

_Type_: String

_Allowed values_: `DEFAULT_ACTION_ORDER | STRICT_ORDER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamExceptionPolicy`

Configures how Network Firewall processes traffic when a network connection breaks midstream. Network connections can break due to disruptions in external networks or within the firewall itself.

- `DROP` \- Network Firewall fails closed and drops all subsequent traffic going to the firewall. This is the default behavior.

- `CONTINUE` \- Network Firewall continues to apply rules to the subsequent traffic without context from traffic before the break. This impacts the behavior of rules that depend on this context. For example, if you have a stateful rule to `drop http` traffic, Network Firewall won't match the traffic for this rule because the service won't have the context from session initialization defining the application layer protocol as HTTP. However, this behavior is rule dependent—a TCP-layer rule using a `flow:stateless` rule would still match, as would the `aws:drop_strict` default action.

- `REJECT` \- Network Firewall fails closed and drops all subsequent traffic going to the firewall. Network Firewall also sends a TCP reject packet back to your client so that the client can immediately establish a new session. Network Firewall will have context about the new session and will apply rules to the subsequent traffic.

_Required_: No

_Type_: String

_Allowed values_: `DROP | CONTINUE | REJECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PublishMetricAction

StatefulRuleGroupOverride

All content copied from https://docs.aws.amazon.com/.
