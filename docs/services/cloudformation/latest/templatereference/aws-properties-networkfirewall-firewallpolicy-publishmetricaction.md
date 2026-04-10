This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::FirewallPolicy PublishMetricAction

Stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the
matching packet. This setting defines a CloudWatch dimension value to be published.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimensions" : [ Dimension, ... ]
}

```

### YAML

```yaml

  Dimensions:
    - Dimension

```

## Properties

`Dimensions`

_Required_: Yes

_Type_: Array of [Dimension](aws-properties-networkfirewall-firewallpolicy-dimension.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PolicyVariables

StatefulEngineOptions

All content copied from https://docs.aws.amazon.com/.
