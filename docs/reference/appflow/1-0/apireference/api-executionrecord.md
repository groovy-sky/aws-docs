# ExecutionRecord

Specifies information about the past flow run instances for a given flow.

## Contents

**dataPullEndTime**

The timestamp that indicates the last new or updated record to be transferred in the flow
run.

Type: Timestamp

Required: No

**dataPullStartTime**

The timestamp that determines the first new or updated record to be transferred in the
flow run.

Type: Timestamp

Required: No

**executionId**

Specifies the identifier of the given flow run.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**executionResult**

Describes the result of the given flow run.

Type: [ExecutionResult](api-executionresult.md) object

Required: No

**executionStatus**

Specifies the flow run status and whether it is in progress, has completed successfully,
or has failed.

Type: String

Valid Values: `InProgress | Successful | Error | CancelStarted | Canceled`

Required: No

**lastUpdatedAt**

Specifies the time of the most recent update.

Type: Timestamp

Required: No

**metadataCatalogDetails**

Describes the metadata catalog, metadata table, and data partitions that Amazon AppFlow used for the associated flow run.

Type: Array of [MetadataCatalogDetail](api-metadatacatalogdetail.md) objects

Required: No

**startedAt**

Specifies the start time of the flow run.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/executionrecord.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/executionrecord.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/executionrecord.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExecutionDetails

ExecutionResult

All content copied from https://docs.aws.amazon.com/.
