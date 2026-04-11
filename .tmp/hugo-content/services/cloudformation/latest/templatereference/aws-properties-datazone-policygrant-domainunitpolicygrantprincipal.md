This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::PolicyGrant DomainUnitPolicyGrantPrincipal

The domain unit principal to whom the policy is granted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainUnitDesignation" : String,
  "DomainUnitGrantFilter" : DomainUnitGrantFilter,
  "DomainUnitIdentifier" : String
}

```

### YAML

```yaml

  DomainUnitDesignation: String
  DomainUnitGrantFilter:
    DomainUnitGrantFilter
  DomainUnitIdentifier: String

```

## Properties

`DomainUnitDesignation`

Specifes the designation of the domain unit users.

_Required_: No

_Type_: String

_Allowed values_: `OWNER`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainUnitGrantFilter`

The grant filter for the domain unit.

_Required_: No

_Type_: [DomainUnitGrantFilter](aws-properties-datazone-policygrant-domainunitgrantfilter.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainUnitIdentifier`

The ID of the domain unit.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9_\-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DomainUnitGrantFilter

GroupPolicyGrantPrincipal

All content copied from https://docs.aws.amazon.com/.
