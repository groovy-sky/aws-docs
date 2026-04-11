This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::FirewallPolicy PolicyVariables

Contains variables that you can use to override default Suricata settings in your firewall policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RuleVariables" : {Key: Value, ...}
}

```

### YAML

```yaml

  RuleVariables:
    Key: Value

```

## Properties

`RuleVariables`

The IPv4 or IPv6 addresses in CIDR notation to use for the Suricata `HOME_NET` variable. If your firewall uses an inspection VPC, you might want to override the `HOME_NET` variable with the CIDRs of your home networks. If you don't override `HOME_NET` with your own CIDRs, Network Firewall by default uses the CIDR of your inspection VPC.

_Required_: No

_Type_: Object of [IPSet](aws-properties-networkfirewall-firewallpolicy-ipset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IPSet

PublishMetricAction

All content copied from https://docs.aws.amazon.com/.
