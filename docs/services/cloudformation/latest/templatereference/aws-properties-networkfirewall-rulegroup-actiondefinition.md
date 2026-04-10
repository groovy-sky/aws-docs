This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup ActionDefinition

A custom action to use in stateless rule actions settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PublishMetricAction" : PublishMetricAction
}

```

### YAML

```yaml

  PublishMetricAction:
    PublishMetricAction

```

## Properties

`PublishMetricAction`

Stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the
matching packet. This setting defines a CloudWatch dimension value to be published.

You can pair this custom action with any of the standard stateless rule actions. For
example, you could pair this in a rule action with the standard action that forwards the
packet for stateful inspection. Then, when a packet matches the rule, Network Firewall
publishes metrics for the packet and forwards it.

_Required_: No

_Type_: [PublishMetricAction](aws-properties-networkfirewall-rulegroup-publishmetricaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::NetworkFirewall::RuleGroup

Address

All content copied from https://docs.aws.amazon.com/.
