This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::PolicyGrant CreateProjectFromProjectProfilePolicyGrantDetail

Specifies whether to create a project from project profile policy grant details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IncludeChildDomainUnits" : Boolean,
  "ProjectProfiles" : [ String, ... ]
}

```

### YAML

```yaml

  IncludeChildDomainUnits: Boolean
  ProjectProfiles:
    - String

```

## Properties

`IncludeChildDomainUnits`

Specifies whether to include child domain units when creating a project from project
profile policy grant details

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProjectProfiles`

Specifies project profiles when creating a project from project profile policy grant
details

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateGlossaryPolicyGrantDetail

CreateProjectPolicyGrantDetail

All content copied from https://docs.aws.amazon.com/.
