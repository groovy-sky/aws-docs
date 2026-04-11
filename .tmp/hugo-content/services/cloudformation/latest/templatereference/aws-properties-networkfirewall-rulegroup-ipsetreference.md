This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup IPSetReference

Configures one or more IP set references for a Suricata-compatible rule group. An IP set reference is a rule variable that references a resource that you create and manage in another AWS service, such as an Amazon VPC prefix list. Network Firewall IP set references enable you to dynamically update the contents of your rules. When you create, update, or delete the IP set you are referencing in your rule, Network Firewall automatically updates the rule's content with the changes. For more information about IP set references in Network Firewall, see [Using IP set references](../../../network-firewall/latest/developerguide/rule-groups-ip-set-references.md) in the _Network Firewall Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReferenceArn" : String
}

```

### YAML

```yaml

  ReferenceArn: String

```

## Properties

`ReferenceArn`

The Amazon Resource Name (ARN) of the resource to include in the IP set reference.

_Required_: No

_Type_: String

_Pattern_: `^(arn:aws.*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IPSet

MatchAttributes

All content copied from https://docs.aws.amazon.com/.
