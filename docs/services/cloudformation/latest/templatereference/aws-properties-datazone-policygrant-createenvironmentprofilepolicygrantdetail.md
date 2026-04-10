This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::PolicyGrant CreateEnvironmentProfilePolicyGrantDetail

The details of the policy grant.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainUnitId" : String
}

```

### YAML

```yaml

  DomainUnitId: String

```

## Properties

`DomainUnitId`

The ID of the domain unit.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9_\-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateDomainUnitPolicyGrantDetail

CreateFormTypePolicyGrantDetail

All content copied from https://docs.aws.amazon.com/.
