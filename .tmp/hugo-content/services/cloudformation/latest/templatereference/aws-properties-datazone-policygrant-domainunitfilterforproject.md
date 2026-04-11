This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::PolicyGrant DomainUnitFilterForProject

The domain unit filter of the project grant filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainUnit" : String,
  "IncludeChildDomainUnits" : Boolean
}

```

### YAML

```yaml

  DomainUnit: String
  IncludeChildDomainUnits: Boolean

```

## Properties

`DomainUnit`

The domain unit ID to use in the filter.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9_\-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IncludeChildDomainUnits`

Specifies whether to include child domain units.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateProjectPolicyGrantDetail

DomainUnitGrantFilter

All content copied from https://docs.aws.amazon.com/.
