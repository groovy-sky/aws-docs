# Ingestion

Contains information about an ingestion.

## Contents

**app**

The name of the application.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**appBundleArn**

The Amazon Resource Name (ARN) of the app bundle for the ingestion.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+`

Required: Yes

**arn**

The Amazon Resource Name (ARN) of the ingestion.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+`

Required: Yes

**createdAt**

The timestamp of when the ingestion was created.

Type: Timestamp

Required: Yes

**ingestionType**

The type of the ingestion.

Type: String

Valid Values: `auditLog`

Required: Yes

**state**

The status of the ingestion.

Type: String

Valid Values: `enabled | disabled`

Required: Yes

**tenantId**

The ID of the application tenant.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**updatedAt**

The timestamp of when the ingestion was last updated.

Type: Timestamp

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/ingestion.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/ingestion.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/ingestion.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FirehoseStream

IngestionDestination

All content copied from https://docs.aws.amazon.com/.
