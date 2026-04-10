This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FMS::Policy PolicyOption

Contains the settings to configure a network ACL policy, a AWS Network Firewall firewall policy deployment model, or a third-party firewall policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NetworkAclCommonPolicy" : NetworkAclCommonPolicy,
  "NetworkFirewallPolicy" : NetworkFirewallPolicy,
  "ThirdPartyFirewallPolicy" : ThirdPartyFirewallPolicy
}

```

### YAML

```yaml

  NetworkAclCommonPolicy:
    NetworkAclCommonPolicy
  NetworkFirewallPolicy:
    NetworkFirewallPolicy
  ThirdPartyFirewallPolicy:
    ThirdPartyFirewallPolicy

```

## Properties

`NetworkAclCommonPolicy`

Defines a Firewall Manager network ACL policy.

_Required_: No

_Type_: [NetworkAclCommonPolicy](aws-properties-fms-policy-networkaclcommonpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkFirewallPolicy`

Defines the deployment model to use for the firewall policy.

_Required_: No

_Type_: [NetworkFirewallPolicy](aws-properties-fms-policy-networkfirewallpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThirdPartyFirewallPolicy`

Defines the policy options for a third-party firewall policy.

_Required_: No

_Type_: [ThirdPartyFirewallPolicy](aws-properties-fms-policy-thirdpartyfirewallpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkFirewallPolicy

PolicyTag

All content copied from https://docs.aws.amazon.com/.
