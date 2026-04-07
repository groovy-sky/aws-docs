This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint KafkaSettings

Provides information that describes an Apache Kafka endpoint. This
information includes the output format of records applied to the endpoint and details of
transaction and control table data information. For more information about other available settings, see
[Using object mapping to migrate data to a Kafka topic](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html#CHAP_Target.Kafka.ObjectMapping)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Broker" : String,
  "IncludeControlDetails" : Boolean,
  "IncludeNullAndEmpty" : Boolean,
  "IncludePartitionValue" : Boolean,
  "IncludeTableAlterOperations" : Boolean,
  "IncludeTransactionDetails" : Boolean,
  "MessageFormat" : String,
  "MessageMaxBytes" : Integer,
  "NoHexPrefix" : Boolean,
  "PartitionIncludeSchemaTable" : Boolean,
  "SaslPassword" : String,
  "SaslUserName" : String,
  "SecurityProtocol" : String,
  "SslCaCertificateArn" : String,
  "SslClientCertificateArn" : String,
  "SslClientKeyArn" : String,
  "SslClientKeyPassword" : String,
  "Topic" : String
}

```

### YAML

```yaml

  Broker: String
  IncludeControlDetails: Boolean
  IncludeNullAndEmpty: Boolean
  IncludePartitionValue: Boolean
  IncludeTableAlterOperations: Boolean
  IncludeTransactionDetails: Boolean
  MessageFormat: String
  MessageMaxBytes: Integer
  NoHexPrefix: Boolean
  PartitionIncludeSchemaTable: Boolean
  SaslPassword: String
  SaslUserName: String
  SecurityProtocol: String
  SslCaCertificateArn: String
  SslClientCertificateArn: String
  SslClientKeyArn: String
  SslClientKeyPassword: String
  Topic: String

```

## Properties

`Broker`

A comma-separated list of one or more broker locations in your Kafka cluster that host your Kafka instance. Specify each broker location
in the form `broker-hostname-or-ip:port`.
For example, `"ec2-12-345-678-901.compute-1.amazonaws.com:2345"`.
For more information and examples of specifying a list of broker locations, see
[Using Apache Kafka as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeControlDetails`

Shows detailed control information for table definition, column definition, and table
and column changes in the Kafka message output. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeNullAndEmpty`

Include NULL and empty columns for records migrated to the endpoint. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludePartitionValue`

Shows the partition value within the Kafka message output unless the partition type is
`schema-table-type`. The default is `false`.

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
`JSON` (default) or `JSON_UNFORMATTED` (a single line with no
tab).

_Required_: No

_Type_: String

_Allowed values_: `json | json-unformatted`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageMaxBytes`

The maximum size in bytes for records created on the endpoint The default is
1,000,000.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoHexPrefix`

Set this optional parameter to `true` to avoid adding a '0x' prefix
to raw data in hexadecimal format. For example, by default, AWS DMS adds a '0x'
prefix to the LOB column type in hexadecimal format moving from an Oracle source to a Kafka
target. Use the `NoHexPrefix` endpoint setting to enable migration of RAW data
type columns without adding the '0x' prefix.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PartitionIncludeSchemaTable`

Prefixes schema and table names to partition values, when the partition type is
`primary-key-type`. Doing this increases data distribution among Kafka
partitions. For example, suppose that a SysBench schema has thousands of tables and each
table has only limited range for a primary key. In this case, the same primary key is sent
from thousands of tables to the same partition, which causes throttling. The default is
`false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SaslPassword`

The secure password that you created when you first set up your Amazon MSK cluster to validate a client identity and
make an encrypted connection between server and client using SASL-SSL authentication.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SaslUserName`

The secure user name you created when you first set up your Amazon MSK cluster to validate a
client identity and make an encrypted connection between server and client using SASL-SSL
authentication.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityProtocol`

Set secure connection to a Kafka target endpoint using Transport Layer Security (TLS). Options include
`ssl-encryption`, `ssl-authentication`, and `sasl-ssl`.
`sasl-ssl` requires `SaslUsername` and `SaslPassword`.

_Required_: No

_Type_: String

_Allowed values_: `plaintext | ssl-authentication | ssl-encryption | sasl-ssl`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SslCaCertificateArn`

The Amazon Resource Name (ARN) for the private certificate authority (CA) cert that AWS DMS uses
to securely connect to your Kafka target endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SslClientCertificateArn`

The Amazon Resource Name (ARN) of the client certificate used to securely connect to a Kafka target endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SslClientKeyArn`

The Amazon Resource Name (ARN) for the client private key used to securely connect to a Kafka target endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SslClientKeyPassword`

The password for the client private key used to securely connect to a Kafka target endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Topic`

The topic to which you migrate the data. If you don't specify a topic, AWS DMS
specifies `"kafka-default-topic"` as the migration topic.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IbmDb2Settings

KinesisSettings
