This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::PolicyGrant ProjectPolicyGrantPrincipal

The project policy grant principal.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProjectDesignation" : String,
  "ProjectGrantFilter" : ProjectGrantFilter,
  "ProjectIdentifier" : String
}

```

### YAML

```yaml

  ProjectDesignation: String
  ProjectGrantFilter:
    ProjectGrantFilter
  ProjectIdentifier: String

```

## Properties

`ProjectDesignation`

The project designation of the project policy grant principal.

_Required_: No

_Type_: String

_Allowed values_: `OWNER | CONTRIBUTOR | PROJECT_CATALOG_STEWARD`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProjectGrantFilter`

The project grant filter of the project policy grant principal.

_Required_: No

_Type_: [ProjectGrantFilter](aws-properties-datazone-policygrant-projectgrantfilter.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProjectIdentifier`

The project ID of the project policy grant principal.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProjectGrantFilter

UserPolicyGrantPrincipal

All content copied from https://docs.aws.amazon.com/.
