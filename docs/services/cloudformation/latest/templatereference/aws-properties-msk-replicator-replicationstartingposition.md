This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Replicator ReplicationStartingPosition

Specifies the position in the topics to start replicating from.

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

The type of replication starting position.

_Required_: No

_Type_: String

_Allowed values_: `LATEST | EARLIEST`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationInfo

ReplicationTopicNameConfiguration

All content copied from https://docs.aws.amazon.com/.
