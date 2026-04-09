# OpenSearchDataSource

This structure contains information about the OpenSearch Service data source used for this
integration. This data source was created as part of the integration setup. An OpenSearch Service data source defines the source and destination for OpenSearch Service queries. It
includes the role required to execute queries and write to collections.

For more information about OpenSearch Service data sources , see [Creating\
OpenSearch Service data source integrations with Amazon S3.](../../../../services/opensearch-service/latest/developerguide/direct-query-s3-creating.md)

## Contents

**dataSourceName**

The name of the OpenSearch Service data source.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**status**

This structure contains information about the status of this OpenSearch Service
resource.

Type: [OpenSearchResourceStatus](api-opensearchresourcestatus.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/opensearchdatasource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/opensearchdatasource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/opensearchdatasource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenSearchDataAccessPolicy

OpenSearchEncryptionPolicy

All content copied from https://docs.aws.amazon.com/.
