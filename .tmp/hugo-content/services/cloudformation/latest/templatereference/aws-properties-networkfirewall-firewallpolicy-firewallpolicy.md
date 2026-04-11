This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::FirewallPolicy FirewallPolicy

The traffic filtering behavior of a firewall policy, defined in a collection of stateless
and stateful rule groups and other settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableTLSSessionHolding" : Boolean,
  "PolicyVariables" : PolicyVariables,
  "StatefulDefaultActions" : [ String, ... ],
  "StatefulEngineOptions" : StatefulEngineOptions,
  "StatefulRuleGroupReferences" : [ StatefulRuleGroupReference, ... ],
  "StatelessCustomActions" : [ CustomAction, ... ],
  "StatelessDefaultActions" : [ String, ... ],
  "StatelessFragmentDefaultActions" : [ String, ... ],
  "StatelessRuleGroupReferences" : [ StatelessRuleGroupReference, ... ],
  "TLSInspectionConfigurationArn" : String
}

```

### YAML

```yaml

  EnableTLSSessionHolding: Boolean
  PolicyVariables:
    PolicyVariables
  StatefulDefaultActions:
    - String
  StatefulEngineOptions:
    StatefulEngineOptions
  StatefulRuleGroupReferences:
    - StatefulRuleGroupReference
  StatelessCustomActions:
    - CustomAction
  StatelessDefaultActions:
    - String
  StatelessFragmentDefaultActions:
    - String
  StatelessRuleGroupReferences:
    - StatelessRuleGroupReference
  TLSInspectionConfigurationArn: String

```

## Properties

`EnableTLSSessionHolding`

When true, prevents TCP and TLS packets from reaching destination servers until TLS Inspection has evaluated Server Name Indication (SNI) rules. Requires an associated TLS Inspection configuration.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyVariables`

Contains variables that you can use to override default Suricata settings in your firewall policy.

_Required_: No

_Type_: [PolicyVariables](aws-properties-networkfirewall-firewallpolicy-policyvariables.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatefulDefaultActions`

The default actions to take on a packet that doesn't match any stateful rules. The stateful default action is optional,
and is only valid when using the strict rule order.

Valid values of the stateful default action:

- aws:drop\_strict

- aws:drop\_established

- aws:alert\_strict

- aws:alert\_established

For more information, see
[Strict evaluation order](../../../network-firewall/latest/developerguide/suricata-rule-evaluation-order.md#suricata-strict-rule-evaluation-order.html) in the _AWS Network Firewall Developer Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatefulEngineOptions`

Additional options governing how Network Firewall handles stateful rules. The stateful
rule groups that you use in your policy must have stateful rule options settings that are compatible with these settings.

_Required_: No

_Type_: [StatefulEngineOptions](aws-properties-networkfirewall-firewallpolicy-statefulengineoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatefulRuleGroupReferences`

References to the stateful rule groups that are used in the policy. These define the
inspection criteria in stateful rules.

_Required_: No

_Type_: Array of [StatefulRuleGroupReference](aws-properties-networkfirewall-firewallpolicy-statefulrulegroupreference.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatelessCustomActions`

The custom action definitions that are available for use in the firewall policy's
`StatelessDefaultActions` setting. You name each custom action that you
define, and then you can use it by name in your default actions specifications.

_Required_: No

_Type_: Array of [CustomAction](aws-properties-networkfirewall-firewallpolicy-customaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatelessDefaultActions`

The actions to take on a packet if it doesn't match any of the stateless rules in the
policy. If you want non-matching packets to be forwarded for stateful inspection, specify
`aws:forward_to_sfe`.

You must specify one of the standard actions: `aws:pass`,
`aws:drop`, or `aws:forward_to_sfe`. In addition, you can specify
custom actions that are compatible with your standard section choice.

For example, you could specify `["aws:pass"]` or you could specify
`["aws:pass", “customActionName”]`. For information about compatibility, see
the custom action descriptions.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatelessFragmentDefaultActions`

The actions to take on a fragmented packet if it doesn't match any of the stateless
rules in the policy. If you want non-matching fragmented packets to be forwarded for
stateful inspection, specify `aws:forward_to_sfe`.

You must specify one of the standard actions: `aws:pass`,
`aws:drop`, or `aws:forward_to_sfe`. In addition, you can specify
custom actions that are compatible with your standard section choice.

For example, you could specify `["aws:pass"]` or you could specify
`["aws:pass", “customActionName”]`. For information about compatibility, see
the custom action descriptions.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatelessRuleGroupReferences`

References to the stateless rule groups that are used in the policy. These define the
matching criteria in stateless rules.

_Required_: No

_Type_: Array of [StatelessRuleGroupReference](aws-properties-networkfirewall-firewallpolicy-statelessrulegroupreference.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TLSInspectionConfigurationArn`

The Amazon Resource Name (ARN) of the TLS inspection configuration.

_Required_: No

_Type_: String

_Pattern_: `^(arn:aws.*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Dimension

FlowTimeouts

All content copied from https://docs.aws.amazon.com/.
