This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup RulesSourceList

Stateful inspection criteria for a domain list rule group.

For HTTPS traffic, domain filtering is SNI-based. It uses the server name indicator extension of the TLS handshake.

By default, Network Firewall domain list inspection only includes traffic coming from the VPC where you deploy the firewall. To inspect traffic from IP addresses outside of the deployment VPC, you set the `HOME_NET` rule variable to include the CIDR range of the deployment VPC plus the other CIDR ranges. For more information, see `RuleVariables` in this guide and [Stateful domain list rule groups in AWS Network Firewall](../../../network-firewall/latest/developerguide/stateful-rule-groups-domain-names.md) in the _Network Firewall Developer Guide_

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GeneratedRulesType" : String,
  "Targets" : [ String, ... ],
  "TargetTypes" : [ String, ... ]
}

```

### YAML

```yaml

  GeneratedRulesType: String
  Targets:
    - String
  TargetTypes:
    - String

```

## Properties

`GeneratedRulesType`

Whether you want to apply allow, reject, alert, or drop behavior to the domains in your target list.

###### Note

When logging is enabled and you choose Alert, traffic that matches the domain specifications
generates an alert in the firewall's logs. Then, traffic either passes, is rejected, or drops based on other rules in the firewall policy.

_Required_: Yes

_Type_: String

_Allowed values_: `ALLOWLIST | DENYLIST | ALERTLIST | REJECTLIST`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Targets`

The domains that you want to inspect for in your traffic flows. Valid domain specifications are the following:

- Explicit names. For example, `abc.example.com` matches only the domain `abc.example.com`.

- Names that use a domain wildcard, which you indicate with an initial ' `.`'. For example, `.example.com` matches `example.com` and matches all subdomains of `example.com`, such as `abc.example.com` and `www.example.com`.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetTypes`

The types of targets to inspect for. Valid values are `TLS_SNI` and `HTTP_HOST`.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RulesSource

RuleVariables

All content copied from https://docs.aws.amazon.com/.
