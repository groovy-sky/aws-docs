This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeArtifact::PackageGroup RestrictionType

The `RestrictionType` property type specifies Property description not available. for an [AWS::CodeArtifact::PackageGroup](aws-resource-codeartifact-packagegroup.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Repositories" : [ String, ... ],
  "RestrictionMode" : String
}

```

### YAML

```yaml

  Repositories:
    - String
  RestrictionMode: String

```

## Properties

`Repositories`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestrictionMode`

Property description not available.

_Required_: Yes

_Type_: String

_Allowed values_: `ALLOW | BLOCK | ALLOW_SPECIFIC_REPOSITORIES | INHERIT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Restrictions

Tag

All content copied from https://docs.aws.amazon.com/.
