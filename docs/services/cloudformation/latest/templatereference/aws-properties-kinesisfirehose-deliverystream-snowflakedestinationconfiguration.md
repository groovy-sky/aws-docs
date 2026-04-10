This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream SnowflakeDestinationConfiguration

Configure Snowflake destination

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountUrl" : String,
  "BufferingHints" : SnowflakeBufferingHints,
  "CloudWatchLoggingOptions" : CloudWatchLoggingOptions,
  "ContentColumnName" : String,
  "Database" : String,
  "DataLoadingOption" : String,
  "KeyPassphrase" : String,
  "MetaDataColumnName" : String,
  "PrivateKey" : String,
  "ProcessingConfiguration" : ProcessingConfiguration,
  "RetryOptions" : SnowflakeRetryOptions,
  "RoleARN" : String,
  "S3BackupMode" : String,
  "S3Configuration" : S3DestinationConfiguration,
  "Schema" : String,
  "SecretsManagerConfiguration" : SecretsManagerConfiguration,
  "SnowflakeRoleConfiguration" : SnowflakeRoleConfiguration,
  "SnowflakeVpcConfiguration" : SnowflakeVpcConfiguration,
  "Table" : String,
  "User" : String
}

```

### YAML

```yaml

  AccountUrl: String
  BufferingHints:
    SnowflakeBufferingHints
  CloudWatchLoggingOptions:
    CloudWatchLoggingOptions
  ContentColumnName: String
  Database: String
  DataLoadingOption: String
  KeyPassphrase: String
  MetaDataColumnName: String
  PrivateKey: String
  ProcessingConfiguration:
    ProcessingConfiguration
  RetryOptions:
    SnowflakeRetryOptions
  RoleARN: String
  S3BackupMode: String
  S3Configuration:
    S3DestinationConfiguration
  Schema: String
  SecretsManagerConfiguration:
    SecretsManagerConfiguration
  SnowflakeRoleConfiguration:
    SnowflakeRoleConfiguration
  SnowflakeVpcConfiguration:
    SnowflakeVpcConfiguration
  Table: String
  User: String

```

## Properties

`AccountUrl`

URL for accessing your Snowflake account. This URL must include your [account identifier](https://docs.snowflake.com/en/user-guide/admin-account-identifier).
Note that the protocol (https://) and port number are optional.

_Required_: Yes

_Type_: String

_Pattern_: `.+?\.snowflakecomputing\.com`

_Minimum_: `24`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BufferingHints`

Describes the buffering to perform before delivering data to the Snowflake destination. If you do not specify any value, Firehose uses the default values.

_Required_: No

_Type_: [SnowflakeBufferingHints](aws-properties-kinesisfirehose-deliverystream-snowflakebufferinghints.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchLoggingOptions`

Property description not available.

_Required_: No

_Type_: [CloudWatchLoggingOptions](aws-properties-kinesisfirehose-deliverystream-cloudwatchloggingoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentColumnName`

The name of the record content column.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Database`

All data in Snowflake is maintained in databases.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLoadingOption`

Choose to load JSON keys mapped to table column names or choose to split the JSON payload where content is mapped to a record content column and source metadata is mapped to a record metadata column.

_Required_: No

_Type_: String

_Allowed values_: `JSON_MAPPING | VARIANT_CONTENT_MAPPING | VARIANT_CONTENT_AND_METADATA_MAPPING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyPassphrase`

Passphrase to decrypt the private key when the key is encrypted. For information, see [Using Key Pair Authentication & Key Rotation](https://docs.snowflake.com/en/user-guide/data-load-snowpipe-streaming-configuration).

_Required_: No

_Type_: String

_Minimum_: `7`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetaDataColumnName`

Specify a column name in the table, where the metadata information has to be loaded.
When you enable this field, you will see the following column in the snowflake table, which
differs based on the source type.

For Direct PUT as source

`{ "firehoseDeliveryStreamName" : "streamname", "IngestionTime" : "timestamp"
            }`

For Kinesis Data Stream as source

` "kinesisStreamName" : "streamname", "kinesisShardId" : "Id",
            "kinesisPartitionKey" : "key", "kinesisSequenceNumber" : "1234", "subsequenceNumber" :
            "2334", "IngestionTime" : "timestamp" }`

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateKey`

The private key used to encrypt your Snowflake client. For information, see [Using Key Pair Authentication & Key Rotation](https://docs.snowflake.com/en/user-guide/data-load-snowpipe-streaming-configuration).

_Required_: No

_Type_: String

_Pattern_: `^(?:[A-Za-z0-9+\/]{4})*(?:[A-Za-z0-9+\/]{2}==|[A-Za-z0-9+\/]{3}=)?$`

_Minimum_: `256`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessingConfiguration`

Property description not available.

_Required_: No

_Type_: [ProcessingConfiguration](aws-properties-kinesisfirehose-deliverystream-processingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryOptions`

The time period where Firehose will retry sending data to the chosen HTTP endpoint.

_Required_: No

_Type_: [SnowflakeRetryOptions](aws-properties-kinesisfirehose-deliverystream-snowflakeretryoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

The Amazon Resource Name (ARN) of the Snowflake role

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BackupMode`

Choose an S3 backup mode

_Required_: No

_Type_: String

_Allowed values_: `FailedDataOnly | AllData`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Configuration`

Property description not available.

_Required_: Yes

_Type_: [S3DestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-s3destinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schema`

Each database consists of one or more schemas, which are logical groupings of database objects, such as tables and views

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerConfiguration`

The configuration that defines how you access secrets for Snowflake.

_Required_: No

_Type_: [SecretsManagerConfiguration](aws-properties-kinesisfirehose-deliverystream-secretsmanagerconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnowflakeRoleConfiguration`

Optionally configure a Snowflake role. Otherwise the default user role will be used.

_Required_: No

_Type_: [SnowflakeRoleConfiguration](aws-properties-kinesisfirehose-deliverystream-snowflakeroleconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnowflakeVpcConfiguration`

The VPCE ID for Firehose to privately connect with Snowflake. The ID format is
com.amazonaws.vpce.\[region\].vpce-svc-<\[id\]>. For more information, see [Amazon PrivateLink & Snowflake](https://docs.snowflake.com/en/user-guide/admin-security-privatelink)

_Required_: No

_Type_: [SnowflakeVpcConfiguration](aws-properties-kinesisfirehose-deliverystream-snowflakevpcconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Table`

All data in Snowflake is stored in database tables, logically structured as collections of columns and rows.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`User`

User login name for the Snowflake account.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnowflakeBufferingHints

SnowflakeRetryOptions

All content copied from https://docs.aws.amazon.com/.
