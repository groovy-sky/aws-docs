This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::PolicyGrant GroupPolicyGrantPrincipal

The group principal to whom the policy is granted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupIdentifier" : String
}

```

### YAML

```yaml

  GroupIdentifier: String

```

## Properties

`GroupIdentifier`

The ID Of the group of the group principal.

_Required_: Yes

_Type_: String

_Pattern_: `(^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$|[\p{L}\p{M}\p{S}\p{N}\p{P}\t\n\r  ]+)`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DomainUnitPolicyGrantPrincipal

OverrideDomainUnitOwnersPolicyGrantDetail

All content copied from https://docs.aws.amazon.com/.
