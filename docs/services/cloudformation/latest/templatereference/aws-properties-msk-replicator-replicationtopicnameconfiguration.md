This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Replicator ReplicationTopicNameConfiguration

Configuration for specifying replicated topic names will be the same as their corresponding upstream topics or prefixed with source cluster alias.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String
}

```

### YAML

```yaml

  Type: String

```

## Properties

`Type`

The type of replication topic name configuration, identical to upstream topic name or prefixed with source cluster alias.

_Required_: No

_Type_: String

_Allowed values_: `PREFIXED_WITH_SOURCE_CLUSTER_ALIAS | IDENTICAL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationStartingPosition

Tag

All content copied from https://docs.aws.amazon.com/.
