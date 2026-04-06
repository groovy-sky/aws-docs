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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/DataSource)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/DataSource)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/DataSource)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataAccessorIdcTrustedTokenIssuerConfiguration

DataSourceSyncJob
