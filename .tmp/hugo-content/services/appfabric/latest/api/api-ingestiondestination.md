# IngestionDestination

Contains information about an ingestion destination.

## Contents

**arn**

The Amazon Resource Name (ARN) of the ingestion destination.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+`

Required: Yes

**destinationConfiguration**

Contains information about the destination of ingested data.

Type: [DestinationConfiguration](api-destinationconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: Yes

**ingestionArn**

The Amazon Resource Name (ARN) of the ingestion.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+`

Required: Yes

**processingConfiguration**

Contains information about how ingested data is processed.

Type: [ProcessingConfiguration](api-processingconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: Yes

**createdAt**

The timestamp of when the ingestion destination was created.

Type: Timestamp

Required: No

**status**

The state of the ingestion destination.

The following states are possible:

- `Active`: The ingestion destination is active and is ready to be
used.

- `Failed`: The ingestion destination has failed. If the ingestion
destination is in this state, you should verify the ingestion destination
configuration and try again.

Type: String

Valid Values: `Active | Failed`

Required: No

**statusReason**

The reason for the current status of the ingestion destination.

Only present when the `status` of ingestion destination is
`Failed`.

Type: String

Required: No

**updatedAt**

The timestamp of when the ingestion destination was last updated.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/ingestiondestination.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/ingestiondestination.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/ingestiondestination.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Ingestion

IngestionDestinationSummary

All content copied from https://docs.aws.amazon.com/.
