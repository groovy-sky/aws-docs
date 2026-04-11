This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::ReplicationConfiguration RepositoryFilter

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

The repository filter details. When the `PREFIX_MATCH` filter type is
specified, this value is required and should be the repository name prefix to configure
replication for.

_Required_: Yes

_Type_: String

_Pattern_: `^(?:[a-z0-9]+(?:[._-][a-z0-9]*)*/)*[a-z0-9]*(?:[._-][a-z0-9]*)*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterType`

The repository filter type. The only supported value is `PREFIX_MATCH`,
which is a repository name prefix specified with the `filter`
parameter.

_Required_: Yes

_Type_: String

_Allowed values_: `PREFIX_MATCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationRule

AWS::ECR::Repository

All content copied from https://docs.aws.amazon.com/.
