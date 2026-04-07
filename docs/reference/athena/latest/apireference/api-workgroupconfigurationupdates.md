# WorkGroupConfigurationUpdates

The configuration information that will be updated for this workgroup, which includes
the location in Amazon S3 where query and calculation results are stored, the
encryption option, if any, used for query results, whether the Amazon CloudWatch
Metrics are enabled for the workgroup, whether the workgroup settings override the
client-side settings, and the data usage limit for the amount of bytes scanned per
query, if it is specified.

## Contents

**AdditionalConfiguration**

Contains a user defined string in JSON format for a Spark-enabled workgroup.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**BytesScannedCutoffPerQuery**

The upper limit (cutoff) for the amount of bytes a single query in a workgroup is
allowed to scan.

Type: Long

Valid Range: Minimum value of 10000000.

Required: No

**CustomerContentEncryptionConfiguration**

Specifies the customer managed KMS key that is used to encrypt the user's data stores
in Athena. When an AWS managed key is used, this value is
null. This setting does not apply to Athena SQL workgroups.

Type: [CustomerContentEncryptionConfiguration](api-customercontentencryptionconfiguration.md) object

Required: No

**EnableMinimumEncryptionConfiguration**

Enforces a minimal level of encryption for the workgroup for query and calculation
results that are written to Amazon S3. When enabled, workgroup users can set
encryption only to the minimum level set by the administrator or higher when they submit
queries. This setting does not apply to Spark-enabled workgroups.

The `EnforceWorkGroupConfiguration` setting takes precedence over the
`EnableMinimumEncryptionConfiguration` flag. This means that if
`EnforceWorkGroupConfiguration` is true, the
`EnableMinimumEncryptionConfiguration` flag is ignored, and the workgroup
configuration for encryption is used.

Type: Boolean

Required: No

**EnforceWorkGroupConfiguration**

If set to "true", the settings for the workgroup override client-side settings. If set
to "false" client-side settings are used. For more information, see [Workgroup Settings Override Client-Side Settings](../../../../services/athena/latest/ug/workgroups-settings-override.md).

Type: Boolean

Required: No

**EngineConfiguration**

The engine configuration for the workgroup, which includes the minimum/maximum number of Data Processing Units (DPU) that queries should use when
running in provisioned capacity. If not specified, Athena uses default values (Default value for min is 4 and for max is Minimum of 124 and allocated DPUs).

To specify DPU values for PC queries the WG containing EngineConfiguration should have the following values:
The name of the Classifications should be `athena-query-engine-properties`, with the only allowed properties as `max-dpu-count` and `min-dpu-count`.

Type: [EngineConfiguration](api-engineconfiguration.md) object

Required: No

**EngineVersion**

The engine version requested when a workgroup is updated. After the update, all
queries on the workgroup run on the requested engine version. If no value was previously
set, the default is Auto. Queries on the `AmazonAthenaPreviewFunctionality`
workgroup run on the preview engine regardless of this setting.

Type: [EngineVersion](api-engineversion.md) object

Required: No

**ExecutionRole**

The ARN of the execution role used to access user resources for Spark sessions and
Identity Center enabled workgroups. This property applies only to Spark enabled
workgroups and Identity Center enabled workgroups.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

Required: No

**ManagedQueryResultsConfigurationUpdates**

Updates configuration information for managed query results in the workgroup.

Type: [ManagedQueryResultsConfigurationUpdates](https://docs.aws.amazon.com/athena/latest/APIReference/API_ManagedQueryResultsConfigurationUpdates.html) object

Required: No

**MonitoringConfiguration**

Contains the configuration settings for managed log persistence, delivering logs to Amazon S3 buckets,
Amazon CloudWatch log groups etc.

Type: [MonitoringConfiguration](api-monitoringconfiguration.md) object

Required: No

**PublishCloudWatchMetricsEnabled**

Indicates whether this workgroup enables publishing metrics to Amazon CloudWatch.

Type: Boolean

Required: No

**QueryResultsS3AccessGrantsConfiguration**

Specifies whether Amazon S3 access grants are enabled for query
results.

Type: [QueryResultsS3AccessGrantsConfiguration](api-queryresultss3accessgrantsconfiguration.md) object

Required: No

**RemoveBytesScannedCutoffPerQuery**

Indicates that the data usage control limit per query is removed. [WorkGroupConfiguration:BytesScannedCutoffPerQuery](api-workgroupconfiguration.md#athena-Type-WorkGroupConfiguration-BytesScannedCutoffPerQuery)

Type: Boolean

Required: No

**RemoveCustomerContentEncryptionConfiguration**

Removes content encryption configuration from an Apache Spark-enabled Athena workgroup.

Type: Boolean

Required: No

**RequesterPaysEnabled**

If set to `true`, allows members assigned to a workgroup to specify Amazon S3 Requester Pays buckets in queries. If set to `false`, workgroup
members cannot query data from Requester Pays buckets, and queries that retrieve data
from Requester Pays buckets cause an error. The default is `false`. For more
information about Requester Pays buckets, see [Requester Pays Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
in the _Amazon Simple Storage Service Developer Guide_.

Type: Boolean

Required: No

**ResultConfigurationUpdates**

The result configuration information about the queries in this workgroup that will be
updated. Includes the updated results location and an updated option for encrypting
query results.

Type: [ResultConfigurationUpdates](https://docs.aws.amazon.com/athena/latest/APIReference/API_ResultConfigurationUpdates.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/WorkGroupConfigurationUpdates)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/WorkGroupConfigurationUpdates)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/WorkGroupConfigurationUpdates)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WorkGroupConfiguration

WorkGroupSummary
