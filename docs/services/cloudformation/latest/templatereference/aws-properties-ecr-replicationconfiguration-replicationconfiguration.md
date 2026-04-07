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

_Type_: Array of [ReplicationRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecr-replicationconfiguration-replicationrule.html)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ECR::ReplicationConfiguration

ReplicationDestination
