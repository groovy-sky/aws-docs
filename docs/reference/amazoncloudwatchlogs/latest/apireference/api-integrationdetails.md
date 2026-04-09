# IntegrationDetails

This structure contains information about the integration configuration. For an
integration with OpenSearch Service, this includes information about OpenSearch Service
resources such as the collection, the workspace, and policies.

This structure is returned by a [GetIntegration](api-getintegration.md) operation.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**openSearchIntegrationDetails**

This structure contains complete information about one integration between CloudWatch Logs and OpenSearch Service.

Type: [OpenSearchIntegrationDetails](api-opensearchintegrationdetails.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/integrationdetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/integrationdetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/integrationdetails.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputLogEvent

IntegrationSummary

All content copied from https://docs.aws.amazon.com/.
