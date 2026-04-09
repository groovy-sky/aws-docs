# OpenSearchDataAccessPolicy

This structure contains information about the OpenSearch Service data access policy used
for this integration. The access policy defines the access controls for the collection. This
data access policy was automatically created as part of the integration setup. For more
information about OpenSearch Service data access policies, see [Data access\
control for Amazon OpenSearch Serverless](../../../../services/opensearch-service/latest/developerguide/serverless-data-access.md) in the OpenSearch Service Developer
Guide.

## Contents

**policyName**

The name of the data access policy.

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

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/opensearchdataaccesspolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/opensearchdataaccesspolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/opensearchdataaccesspolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenSearchCollection

OpenSearchDataSource

All content copied from https://docs.aws.amazon.com/.
