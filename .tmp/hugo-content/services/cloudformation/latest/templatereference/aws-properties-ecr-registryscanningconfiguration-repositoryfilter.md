This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::RegistryScanningConfiguration RepositoryFilter

The filter settings used with image replication. Specifying a repository filter to a
replication rule provides a method for controlling which repositories in a private
registry are replicated. If no filters are added, the contents of all repositories are
replicated.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Filter" : String,
  "FilterType" : String
}

```

### YAML

```yaml

  Filter: String
  FilterType: String

```

## Properties

`Filter`

The filter to use when scanning.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9*](?:[._\-/a-z0-9*]?[a-z0-9*]+)*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterType`

The type associated with the filter.

_Required_: Yes

_Type_: String

_Allowed values_: `WILDCARD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ECR::RegistryScanningConfiguration

ScanningRule

All content copied from https://docs.aws.amazon.com/.
