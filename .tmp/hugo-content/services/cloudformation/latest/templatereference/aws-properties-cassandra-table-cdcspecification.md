This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Table CdcSpecification

The settings for the CDC stream of a table. For more information about CDC streams, see [Working with change data capture (CDC) streams in Amazon Keyspaces](../../../keyspaces/latest/devguide/cdc.md) in the _Amazon Keyspaces Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Status" : String,
  "Tags" : [ Tag, ... ],
  "ViewType" : String
}

```

### YAML

```yaml

  Status: String
  Tags:
    - Tag
  ViewType: String

```

## Properties

`Status`

The status of the CDC stream. You can enable or disable a stream for a table.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags (key-value pairs) that you want to apply to the stream.

_Required_: No

_Type_: Array of [Tag](aws-properties-cassandra-table-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ViewType`

The view type specifies the changes Amazon Keyspaces records for each changed row in the stream.
After you create the stream, you can't make changes to this selection.

The options are:

- `NEW_AND_OLD_IMAGES` \- both versions of the row, before and after the change. This is the default.

- `NEW_IMAGE` \- the version of the row after the change.

- `OLD_IMAGE` \- the version of the row before the change.

- `KEYS_ONLY` \- the partition and clustering keys of the row that was changed.

_Required_: No

_Type_: String

_Allowed values_: `NEW_IMAGE | OLD_IMAGE | KEYS_ONLY | NEW_AND_OLD_IMAGES`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BillingMode

ClusteringKeyColumn

All content copied from https://docs.aws.amazon.com/.
