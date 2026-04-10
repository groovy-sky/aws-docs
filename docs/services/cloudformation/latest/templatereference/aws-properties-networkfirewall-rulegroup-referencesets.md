This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup ReferenceSets

Configures the reference sets for a stateful rule group. For more information, see the [Using IP set references in Suricata compatible rule groups](../../../network-firewall/latest/developerguide/rule-groups-ip-set-references.md) in the _Network Firewall User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IPSetReferences" : {Key: Value, ...}
}

```

### YAML

```yaml

  IPSetReferences:
    Key: Value

```

## Properties

`IPSetReferences`

The IP set references to use in the stateful rule group.

_Required_: No

_Type_: Object of [IPSetReference](aws-properties-networkfirewall-rulegroup-ipsetreference.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PublishMetricAction

RuleDefinition

All content copied from https://docs.aws.amazon.com/.
