This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FMS::Policy NetworkAclEntrySet

The configuration of the first and last rules for the network ACL policy, and the remediation settings for each.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FirstEntries" : [ NetworkAclEntry, ... ],
  "ForceRemediateForFirstEntries" : Boolean,
  "ForceRemediateForLastEntries" : Boolean,
  "LastEntries" : [ NetworkAclEntry, ... ]
}

```

### YAML

```yaml

  FirstEntries:
    - NetworkAclEntry
  ForceRemediateForFirstEntries: Boolean
  ForceRemediateForLastEntries: Boolean
  LastEntries:
    - NetworkAclEntry

```

## Properties

`FirstEntries`

The rules that you want to run first in the Firewall Manager managed network ACLs.

###### Note

Provide these in the order in which you want them to run. Firewall Manager will assign
the specific rule numbers for you, in the network ACLs that it creates.

You must specify at least one first entry or one last entry in any network ACL policy.

_Required_: No

_Type_: Array of [NetworkAclEntry](aws-properties-fms-policy-networkaclentry.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForceRemediateForFirstEntries`

Applies only when remediation is enabled for the policy as a whole. Firewall Manager uses this setting when it finds policy
violations that involve conflicts between the custom entries and the policy entries.

If forced remediation is disabled, Firewall Manager marks the network ACL as noncompliant and does not try to
remediate. For more information about the remediation behavior, see
[Remediation for managed network ACLs](../../../waf/latest/developerguide/network-acl-policies.md#network-acls-remediation)
in the _AWS Firewall Manager Developer Guide_.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForceRemediateForLastEntries`

Applies only when remediation is enabled for the policy as a whole. Firewall Manager uses this setting when it finds policy
violations that involve conflicts between the custom entries and the policy entries.

If forced remediation is disabled, Firewall Manager marks the network ACL as noncompliant and does not try to
remediate. For more information about the remediation behavior, see
[Remediation for managed network ACLs](../../../waf/latest/developerguide/network-acl-policies.md#network-acls-remediation)
in the _AWS Firewall Manager Developer Guide_.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastEntries`

The rules that you want to run last in the Firewall Manager managed network ACLs.

###### Note

Provide these in the order in which you want them to run. Firewall Manager will assign
the specific rule numbers for you, in the network ACLs that it creates.

You must specify at least one first entry or one last entry in any network ACL policy.

_Required_: No

_Type_: Array of [NetworkAclEntry](aws-properties-fms-policy-networkaclentry.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkAclEntry

NetworkFirewallPolicy

All content copied from https://docs.aws.amazon.com/.
