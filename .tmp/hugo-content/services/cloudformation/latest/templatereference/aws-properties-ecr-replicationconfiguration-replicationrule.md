This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::ReplicationConfiguration ReplicationRule

An array of objects representing the replication destinations and repository filters
for a replication configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destinations" : [ ReplicationDestination, ... ],
  "RepositoryFilters" : [ RepositoryFilter, ... ]
}

```

### YAML

```yaml

  Destinations:
    - ReplicationDestination
  RepositoryFilters:
    - RepositoryFilter

```

## Properties

`Destinations`

An array of objects representing the destination for a replication rule.

_Required_: Yes

_Type_: Array of [ReplicationDestination](aws-properties-ecr-replicationconfiguration-replicationdestination.md)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepositoryFilters`

An array of objects representing the filters for a replication rule. Specifying a
repository filter for a replication rule provides a method for controlling which
repositories in a private registry are replicated.

_Required_: No

_Type_: Array of [RepositoryFilter](aws-properties-ecr-replicationconfiguration-repositoryfilter.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationDestination

RepositoryFilter

All content copied from https://docs.aws.amazon.com/.
