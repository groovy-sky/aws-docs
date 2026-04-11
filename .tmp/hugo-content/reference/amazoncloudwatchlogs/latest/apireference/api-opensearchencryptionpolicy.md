# OpenSearchEncryptionPolicy

This structure contains information about the OpenSearch Service encryption policy used
for this integration. The encryption policy was created automatically when you created the
integration. For more information, see [Encryption policies](../../../../services/opensearch-service/latest/developerguide/serverless-encryption.md#serverless-encryption-policies) in the OpenSearch Service Developer Guide.

## Contents

**policyName**

The name of the encryption policy.

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

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/opensearchencryptionpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/opensearchencryptionpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/opensearchencryptionpolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenSearchDataSource

OpenSearchIntegrationDetails

All content copied from https://docs.aws.amazon.com/.
