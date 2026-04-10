This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint KinesisSettings

Provides information that describes an Amazon Kinesis Data Stream endpoint. This
information includes the output format of records applied to the endpoint and details of
transaction and control table data information. For more information about other available settings, see
[Using object mapping to migrate data to a Kinesis data stream](../../../dms/latest/userguide/chap-target-kinesis.md#CHAP_Target.Kinesis.ObjectMapping)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IncludeControlDetails" : Boolean,
  "IncludeNullAndEmpty" : Boolean,
  "IncludePartitionValue" : Boolean,
  "IncludeTableAlterOperations" : Boolean,
  "IncludeTransactionDetails" : Boolean,
  "MessageFormat" : String,
  "NoHexPrefix" : Boolean,
  "PartitionIncludeSchemaTable" : Boolean,
  "ServiceAccessRoleArn" : String,
  "StreamArn" : String
}

```

### YAML

```yaml

  IncludeControlDetails: Boolean
  IncludeNullAndEmpty: Boolean
  IncludePartitionValue: Boolean
  IncludeTableAlterOperations: Boolean
  IncludeTransactionDetails: Boolean
  MessageFormat: String
  NoHexPrefix: Boolean
  PartitionIncludeSchemaTable: Boolean
  ServiceAccessRoleArn: String
  StreamArn: String

```

## Properties

`IncludeControlDetails`

Shows detailed control information for table definition, column definition, and table
and column changes in the Kinesis message output. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeNullAndEmpty`

Include NULL and empty columns for records migrated to the endpoint. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludePartitionValue`

Shows the partition value within the Kinesis message output, unless the partition type
is `schema-table-type`. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeTableAlterOperations`

Includes any data definition language (DDL) operations that change the table in the
control data, such as `rename-table`, `drop-table`,
`add-column`, `drop-column`, and `rename-column`. The
default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeTransactionDetails`

Provides detailed transaction information from the source database. This information
includes a commit timestamp, a log position, and values for `transaction_id`,
previous `transaction_id`, and `transaction_record_id` (the record
offset within a transaction). The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageFormat`

The output format for the records created on the endpoint. The message format is
`JSON` (default) or `JSON_UNFORMATTED` (a single line with no tab).

_Required_: No

_Type_: String

_Allowed values_: `json | json-unformatted`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoHexPrefix`

Set this optional parameter to `true` to avoid adding a '0x' prefix
to raw data in hexadecimal format. For example, by default, AWS DMS adds a '0x'
prefix to the LOB column type in hexadecimal format moving from an Oracle source to an
Amazon Kinesis target. Use the `NoHexPrefix` endpoint setting to enable
migration of RAW data type columns without adding the '0x' prefix.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PartitionIncludeSchemaTable`

Prefixes schema and table names to partition values, when the partition type is
`primary-key-type`. Doing this increases data distribution among Kinesis
shards. For example, suppose that a SysBench schema has thousands of tables and each table
has only limited range for a primary key. In this case, the same primary key is sent from
thousands of tables to the same shard, which causes throttling. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceAccessRoleArn`

The Amazon Resource Name (ARN) for the IAM role
that AWS DMS uses to write to the Kinesis data stream.
The role must allow the `iam:PassRole` action.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamArn`

The Amazon Resource Name (ARN) for the Amazon Kinesis Data Streams endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KafkaSettings

MicrosoftSqlServerSettings

All content copied from https://docs.aws.amazon.com/.
