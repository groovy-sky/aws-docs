This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::ReplicationConfiguration ReplicationConfiguration

The replication configuration for a registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rules" : [ ReplicationRule, ... ]
}

```

### YAML

```yaml

  Rules:
    - ReplicationRule

```

## Properties

`Rules`

An array of objects representing the replication destinations and repository filters
for a replication configuration.

_Required_: Yes

_Type_: Array of [ReplicationRule](aws-properties-ecr-replicationconfiguration-replicationrule.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ECR::ReplicationConfiguration

ReplicationDestination

All content copied from https://docs.aws.amazon.com/.
