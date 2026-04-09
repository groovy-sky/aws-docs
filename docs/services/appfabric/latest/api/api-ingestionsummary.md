# IngestionSummary

Contains a summary of an ingestion.

## Contents

**app**

The name of the application.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**arn**

The Amazon Resource Name (ARN) of the ingestion.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+`

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/ingestionsummary.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/ingestionsummary.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/ingestionsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IngestionDestinationSummary

Oauth2Credential

All content copied from https://docs.aws.amazon.com/.
