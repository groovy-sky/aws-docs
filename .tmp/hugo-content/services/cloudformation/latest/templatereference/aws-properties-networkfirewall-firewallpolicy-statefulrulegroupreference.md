This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::FirewallPolicy StatefulRuleGroupReference

Identifier for a single stateful rule group, used in a firewall policy to refer to a
rule group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeepThreatInspection" : Boolean,
  "Override" : StatefulRuleGroupOverride,
  "Priority" : Integer,
  "ResourceArn" : String
}

```

### YAML

```yaml

  DeepThreatInspection: Boolean
  Override:
    StatefulRuleGroupOverride
  Priority: Integer
  ResourceArn: String

```

## Properties

`DeepThreatInspection`

AWS Network Firewall plans to augment the active threat defense managed rule group with an additional deep threat inspection capability. When this capability is released, AWS will analyze service logs of network traffic processed by these rule groups to identify threat indicators across customers.
AWS will use these threat indicators to improve the active threat defense managed rule groups and protect the security of AWS customers and services.

###### Note

Customers can opt-out of deep threat inspection at any time through the AWS Network Firewall console or API. When customers opt out, AWS Network Firewall will not use the network traffic processed by those customers' active threat defense rule groups for rule group improvement.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Override`

The action that allows the policy owner to override the behavior of the rule group within a policy.

_Required_: No

_Type_: [StatefulRuleGroupOverride](aws-properties-networkfirewall-firewallpolicy-statefulrulegroupoverride.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

An integer setting that indicates the order in which to run the stateful rule groups in
a single firewall policy. This setting only applies to firewall policies
that specify the `STRICT_ORDER` rule order in the stateful engine options settings.

Network Firewall evalutes each stateful rule group
against a packet starting with the group that has the lowest priority setting. You must ensure
that the priority settings are unique within each policy.

You can change the priority settings of your rule groups at any time. To make it easier to
insert rule groups later, number them so there's a wide range in between, for example use 100,
200, and so on.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

The Amazon Resource Name (ARN) of the stateful rule group.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws.*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StatefulRuleGroupOverride

StatelessRuleGroupReference

All content copied from https://docs.aws.amazon.com/.
