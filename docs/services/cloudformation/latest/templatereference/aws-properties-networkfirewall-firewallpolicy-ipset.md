---
title: "AWS::NetworkFirewall::FirewallPolicy IPSet"
---

This is the new *CloudFormation Template Reference Guide*. Please update your bookmarks and links. For help getting started with CloudFormation, see the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html).

# AWS::NetworkFirewall::FirewallPolicy IPSet
<a name="aws-properties-networkfirewall-firewallpolicy-ipset"></a>

A list of IP addresses and address ranges, in CIDR notation. This is part of a rule variable.

## Syntax
<a name="aws-properties-networkfirewall-firewallpolicy-ipset-syntax"></a>

To declare this entity in your CloudFormation template, use the following syntax:

### JSON
<a name="aws-properties-networkfirewall-firewallpolicy-ipset-syntax.json"></a>

```
{
  "[Definition](#cfn-networkfirewall-firewallpolicy-ipset-definition)" : {{[ String, ... ]}}
}
```

### YAML
<a name="aws-properties-networkfirewall-firewallpolicy-ipset-syntax.yaml"></a>

```
  [Definition](#cfn-networkfirewall-firewallpolicy-ipset-definition): {{
    - String}}
```

## Properties
<a name="aws-properties-networkfirewall-firewallpolicy-ipset-properties"></a>

`Definition`  <a name="cfn-networkfirewall-firewallpolicy-ipset-definition"></a>
The list of IP addresses and address ranges, in CIDR notation.
*Required*: No
*Type*: Array of String
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

All content copied from https://docs.aws.amazon.com/.
