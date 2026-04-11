This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL AWSManagedRulesAntiDDoSRuleSet

Configures the use of the anti-DDoS managed rule group, `AWSManagedRulesAntiDDoSRuleSet`. This configuration is used in `ManagedRuleGroupConfig`.

The configuration that you provide here determines whether and how the rules in the rule group are used.

For additional information about this and the other intelligent threat mitigation rule groups,
see [Intelligent threat mitigation in AWS WAF](../../../waf/latest/developerguide/waf-managed-protections.md)
and [AWS Managed Rules rule groups list](../../../waf/latest/developerguide/aws-managed-rule-groups-list.md)
in the _AWS WAF Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientSideActionConfig" : ClientSideActionConfig,
  "SensitivityToBlock" : String
}

```

### YAML

```yaml

  ClientSideActionConfig:
    ClientSideActionConfig
  SensitivityToBlock: String

```

## Properties

`ClientSideActionConfig`

Configures the request handling that's applied by the managed rule group rules `ChallengeAllDuringEvent` and `ChallengeDDoSRequests` during a distributed denial of service (DDoS) attack.

_Required_: Yes

_Type_: [ClientSideActionConfig](aws-properties-wafv2-webacl-clientsideactionconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SensitivityToBlock`

The sensitivity that the rule group rule `DDoSRequests` uses when matching against the
DDoS suspicion labeling on a request. The managed rule group adds the labeling during DDoS events, before the `DDoSRequests` rule runs.

The higher the sensitivity, the more levels of labeling that the rule matches:

- Low sensitivity is less sensitive, causing the rule to match only on the most likely participants in an attack, which are the requests with the high suspicion label `awswaf:managed:aws:anti-ddos:high-suspicion-ddos-request`.

- Medium sensitivity causes the rule to match on the medium and high suspicion labels.

- High sensitivity causes the rule to match on all of the suspicion labels: low, medium, and high.

Default: `LOW`

_Required_: No

_Type_: String

_Allowed values_: `LOW | MEDIUM | HIGH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWSManagedRulesACFPRuleSet

AWSManagedRulesATPRuleSet

All content copied from https://docs.aws.amazon.com/.
