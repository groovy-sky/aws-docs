This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::FirewallPolicy StatelessRuleGroupReference

Identifier for a single stateless rule group, used in a firewall policy to refer to the
rule group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Priority" : Integer,
  "ResourceArn" : String
}

```

### YAML

```yaml

  Priority: Integer
  ResourceArn: String

```

## Properties

`Priority`

An integer setting that indicates the order in which to run the stateless rule groups in
a single firewall policy. Network Firewall applies each stateless rule group
to a packet starting with the group that has the lowest priority setting. You must ensure
that the priority settings are unique within each policy.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

The Amazon Resource Name (ARN) of the stateless rule group.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws.*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StatefulRuleGroupReference

Tag

All content copied from https://docs.aws.amazon.com/.
