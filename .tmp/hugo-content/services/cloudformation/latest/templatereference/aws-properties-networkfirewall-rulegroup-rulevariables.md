This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup RuleVariables

Settings that are available for use in the rules in the rule group
where this is defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IPSets" : {Key: Value, ...},
  "PortSets" : {Key: Value, ...}
}

```

### YAML

```yaml

  IPSets:
    Key: Value
  PortSets:
    Key: Value

```

## Properties

`IPSets`

A list of IP addresses and address ranges, in CIDR notation.

_Required_: No

_Type_: Object of [IPSet](aws-properties-networkfirewall-rulegroup-ipset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortSets`

A list of port ranges.

_Required_: No

_Type_: Object of [PortSet](aws-properties-networkfirewall-rulegroup-portset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RulesSourceList

StatefulRule

All content copied from https://docs.aws.amazon.com/.
