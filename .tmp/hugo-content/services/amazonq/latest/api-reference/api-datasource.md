# DataSource

A data source in an Amazon Q Business application.

## Contents

**createdAt**

The Unix timestamp when the Amazon Q Business data source was created.

Type: Timestamp

Required: No

**dataSourceId**

The identifier of the Amazon Q Business data source.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**displayName**

The name of the Amazon Q Business data source.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`

Required: No

**status**

The status of the Amazon Q Business data source.

Type: String

Valid Values: `PENDING_CREATION | CREATING | ACTIVE | DELETING | FAILED | UPDATING`

Required: No

**type**

The type of the Amazon Q Business data source.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**updatedAt**

The Unix timestamp when the Amazon Q Business data source was last updated.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/datasource.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/datasource.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/datasource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataAccessorIdcTrustedTokenIssuerConfiguration

DataSourceSyncJob

All content copied from https://docs.aws.amazon.com/.
