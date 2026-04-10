This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::WorkGroup WorkGroupConfiguration

The configuration of the workgroup, which includes the location in Amazon S3 where
query results are stored, the encryption option, if any, used for query results, whether
Amazon CloudWatch Metrics are enabled for the workgroup, and the limit for the amount of
bytes scanned (cutoff) per query, if it is specified. The [EnforceWorkGroupConfiguration](../userguide/aws-properties-athena-workgroup-workgroupconfiguration.md#cfn-athena-workgroup-workgroupconfiguration-enforceworkgroupconfiguration) option determines whether workgroup
settings override client-side query settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalConfiguration" : String,
  "BytesScannedCutoffPerQuery" : Integer,
  "CustomerContentEncryptionConfiguration" : CustomerContentEncryptionConfiguration,
  "EnforceWorkGroupConfiguration" : Boolean,
  "EngineConfiguration" : EngineConfiguration,
  "EngineVersion" : EngineVersion,
  "ExecutionRole" : String,
  "ManagedQueryResultsConfiguration" : ManagedQueryResultsConfiguration,
  "MonitoringConfiguration" : MonitoringConfiguration,
  "PublishCloudWatchMetricsEnabled" : Boolean,
  "RequesterPaysEnabled" : Boolean,
  "ResultConfiguration" : ResultConfiguration
}

```

### YAML

```yaml

  AdditionalConfiguration: String
  BytesScannedCutoffPerQuery: Integer
  CustomerContentEncryptionConfiguration:
    CustomerContentEncryptionConfiguration
  EnforceWorkGroupConfiguration: Boolean
  EngineConfiguration:
    EngineConfiguration
  EngineVersion:
    EngineVersion
  ExecutionRole: String
  ManagedQueryResultsConfiguration:
    ManagedQueryResultsConfiguration
  MonitoringConfiguration:
    MonitoringConfiguration
  PublishCloudWatchMetricsEnabled: Boolean
  RequesterPaysEnabled: Boolean
  ResultConfiguration:
    ResultConfiguration

```

## Properties

`AdditionalConfiguration`

Specifies a user defined JSON string that is passed to the session engine.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BytesScannedCutoffPerQuery`

The upper limit (cutoff) for the amount of bytes a single query in a workgroup is
allowed to scan. No default is defined.

###### Note

This property currently supports integer types. Support for long values is
planned.

_Required_: No

_Type_: Integer

_Minimum_: `10000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomerContentEncryptionConfiguration`

Specifies the KMS key that is used to encrypt the user's data stores in Athena. This setting does not apply to Athena SQL workgroups.

_Required_: No

_Type_: [CustomerContentEncryptionConfiguration](aws-properties-athena-workgroup-customercontentencryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnforceWorkGroupConfiguration`

If set to "true", the settings for the workgroup override client-side settings. If set
to "false", client-side settings are used. For more information, see [Override client-side settings](../../../athena/latest/ug/workgroups-settings-override.md).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EngineConfiguration`

Property description not available.

_Required_: No

_Type_: [EngineConfiguration](aws-properties-athena-workgroup-engineconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EngineVersion`

The engine version that all queries running on the workgroup use.

_Required_: No

_Type_: [EngineVersion](aws-properties-athena-workgroup-engineversion.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRole`

Role used to access user resources in an Athena for Apache Spark session. This
property applies only to Spark-enabled workgroups in Athena.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedQueryResultsConfiguration`

The configuration for storing results in Athena owned storage, which includes whether
this feature is enabled; whether encryption configuration, if any, is used for
encrypting query results.

_Required_: No

_Type_: [ManagedQueryResultsConfiguration](aws-properties-athena-workgroup-managedqueryresultsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringConfiguration`

Contains the configuration settings for managed log persistence, delivering logs to Amazon S3 buckets,
Amazon CloudWatch log groups etc.

_Required_: No

_Type_: [MonitoringConfiguration](aws-properties-athena-workgroup-monitoringconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublishCloudWatchMetricsEnabled`

Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequesterPaysEnabled`

If set to `true`, allows members assigned to a workgroup to reference
Amazon S3 Requester Pays buckets in queries. If set to `false`, workgroup
members cannot query data from Requester Pays buckets, and queries that retrieve data
from Requester Pays buckets cause an error. The default is `false`. For more
information about Requester Pays buckets, see [Requester Pays Buckets](../../../s3/latest/dev/requesterpaysbuckets.md)
in the _Amazon Simple Storage Service Developer Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResultConfiguration`

Specifies the location in Amazon S3 where query results are stored and the encryption
option, if any, used for query results. For more information, see [Work with query results\
and recent queries](../../../athena/latest/ug/querying.md).

_Required_: No

_Type_: [ResultConfiguration](aws-properties-athena-workgroup-resultconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
