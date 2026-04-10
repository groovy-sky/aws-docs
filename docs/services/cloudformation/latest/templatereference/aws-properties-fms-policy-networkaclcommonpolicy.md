This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FMS::Policy NetworkAclCommonPolicy

Defines a Firewall Manager network ACL policy. This is used in the `PolicyOption` of a `SecurityServicePolicyData` for a `Policy`, when
the `SecurityServicePolicyData` type is set to `NETWORK_ACL_COMMON`.

For information about network ACLs, see
[Control traffic to subnets using network ACLs](../../../vpc/latest/userguide/vpc-network-acls.md)
in the _Amazon Virtual Private Cloud User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NetworkAclEntrySet" : NetworkAclEntrySet
}

```

### YAML

```yaml

  NetworkAclEntrySet:
    NetworkAclEntrySet

```

## Properties

`NetworkAclEntrySet`

The definition of the first and last rules for the network ACL policy.

_Required_: Yes

_Type_: [NetworkAclEntrySet](aws-properties-fms-policy-networkaclentryset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IEMap

NetworkAclEntry

All content copied from https://docs.aws.amazon.com/.
