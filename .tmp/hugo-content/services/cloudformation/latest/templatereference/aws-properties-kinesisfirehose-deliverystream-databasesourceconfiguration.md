This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream DatabaseSourceConfiguration

The top level object for configuring streams with database as a source.

Amazon Data Firehose is in preview release and is subject to change.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Columns" : DatabaseColumns,
  "Databases" : Databases,
  "DatabaseSourceAuthenticationConfiguration" : DatabaseSourceAuthenticationConfiguration,
  "DatabaseSourceVPCConfiguration" : DatabaseSourceVPCConfiguration,
  "Digest" : String,
  "Endpoint" : String,
  "Port" : Integer,
  "PublicCertificate" : String,
  "SnapshotWatermarkTable" : String,
  "SSLMode" : String,
  "SurrogateKeys" : [ String, ... ],
  "Tables" : DatabaseTables,
  "Type" : String
}

```

### YAML

```yaml

  Columns:
    DatabaseColumns
  Databases:
    Databases
  DatabaseSourceAuthenticationConfiguration:
    DatabaseSourceAuthenticationConfiguration
  DatabaseSourceVPCConfiguration:
    DatabaseSourceVPCConfiguration
  Digest: String
  Endpoint: String
  Port: Integer
  PublicCertificate: String
  SnapshotWatermarkTable: String
  SSLMode: String
  SurrogateKeys:
    - String
  Tables:
    DatabaseTables
  Type: String

```

## Properties

`Columns`

The list of column patterns in source database endpoint for Firehose to read from.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: No

_Type_: [DatabaseColumns](aws-properties-kinesisfirehose-deliverystream-databasecolumns.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Databases`

The list of database patterns in source database endpoint for Firehose to read from.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: Yes

_Type_: [Databases](aws-properties-kinesisfirehose-deliverystream-databases.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseSourceAuthenticationConfiguration`

The structure to configure the authentication methods for Firehose to connect to source database endpoint.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: Yes

_Type_: [DatabaseSourceAuthenticationConfiguration](aws-properties-kinesisfirehose-deliverystream-databasesourceauthenticationconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseSourceVPCConfiguration`

The details of the VPC Endpoint Service which Firehose uses to create a PrivateLink to the database.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: Yes

_Type_: [DatabaseSourceVPCConfiguration](aws-properties-kinesisfirehose-deliverystream-databasesourcevpcconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Digest`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Endpoint`

The endpoint of the database server.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\s*$).+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Port`

The port of the database. This can be one of the following values.

- 3306 for MySQL database type

- 5432 for PostgreSQL database type

Amazon Data Firehose is in preview release and is subject to change.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PublicCertificate`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotWatermarkTable`

The fully qualified name of the table in source database endpoint that Firehose uses to track snapshot progress.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0001-\uFFFF]*`

_Minimum_: `1`

_Maximum_: `129`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SSLMode`

The mode to enable or disable SSL when Firehose connects to the database endpoint.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: No

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SurrogateKeys`

The optional list of table and column names used as unique key columns when taking snapshot if the tables don’t have primary keys configured.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tables`

The list of table patterns in source database endpoint for Firehose to read from.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: Yes

_Type_: [DatabaseTables](aws-properties-kinesisfirehose-deliverystream-databasetables.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of database engine. This can be one of the following values.

- MySQL

- PostgreSQL

Amazon Data Firehose is in preview release and is subject to change.

_Required_: Yes

_Type_: String

_Allowed values_: `MySQL | PostgreSQL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatabaseSourceAuthenticationConfiguration

DatabaseSourceVPCConfiguration

All content copied from https://docs.aws.amazon.com/.
