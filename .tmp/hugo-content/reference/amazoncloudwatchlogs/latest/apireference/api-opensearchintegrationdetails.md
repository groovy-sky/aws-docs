# OpenSearchIntegrationDetails

This structure contains complete information about one CloudWatch Logs integration.
This structure is returned by a [GetIntegration](api-getintegration.md) operation.

## Contents

**accessPolicy**

This structure contains information about the OpenSearch Service data access policy used
for this integration. The access policy defines the access controls for the collection. This
data access policy was automatically created as part of the integration setup. For more
information about OpenSearch Service data access policies, see [Data access\
control for Amazon OpenSearch Serverless](../../../../services/opensearch-service/latest/developerguide/serverless-data-access.md) in the OpenSearch Service Developer
Guide.

Type: [OpenSearchDataAccessPolicy](api-opensearchdataaccesspolicy.md) object

Required: No

**application**

This structure contains information about the OpenSearch Service application used for this
integration. An OpenSearch Service application is the web application that was created by the
integration with CloudWatch Logs. It hosts the vended logs dashboards.

Type: [OpenSearchApplication](api-opensearchapplication.md) object

Required: No

**collection**

This structure contains information about the OpenSearch Service collection used for this
integration. This collection was created as part of the integration setup. An OpenSearch Service collection is a logical grouping of one or more indexes that represent an analytics
workload. For more information, see [Creating and\
managing OpenSearch Service Serverless collections](../../../../services/opensearch-service/latest/developerguide/serverless-collections.md).

Type: [OpenSearchCollection](api-opensearchcollection.md) object

Required: No

**dataSource**

This structure contains information about the OpenSearch Service data source used for this
integration. This data source was created as part of the integration setup. An OpenSearch Service data source defines the source and destination for OpenSearch Service queries. It
includes the role required to execute queries and write to collections.

For more information about OpenSearch Service data sources , see [Creating\
OpenSearch Service data source integrations with Amazon S3.](../../../../services/opensearch-service/latest/developerguide/direct-query-s3-creating.md)

Type: [OpenSearchDataSource](api-opensearchdatasource.md) object

Required: No

**encryptionPolicy**

This structure contains information about the OpenSearch Service encryption policy used
for this integration. The encryption policy was created automatically when you created the
integration. For more information, see [Encryption policies](../../../../services/opensearch-service/latest/developerguide/serverless-encryption.md#serverless-encryption-policies) in the OpenSearch Service Developer Guide.

Type: [OpenSearchEncryptionPolicy](api-opensearchencryptionpolicy.md) object

Required: No

**lifecyclePolicy**

This structure contains information about the OpenSearch Service data lifecycle policy
used for this integration. The lifecycle policy determines the lifespan of the data in the
collection. It was automatically created as part of the integration setup.

For more information, see [Using data\
lifecycle policies with OpenSearch Service Serverless](../../../../services/opensearch-service/latest/developerguide/serverless-lifecycle.md) in the OpenSearch Service
Developer Guide.

Type: [OpenSearchLifecyclePolicy](api-opensearchlifecyclepolicy.md) object

Required: No

**networkPolicy**

This structure contains information about the OpenSearch Service network policy used for
this integration. The network policy assigns network access settings to collections. For more
information, see [Network policies](../../../../services/opensearch-service/latest/developerguide/serverless-network.md#serverless-network-policies) in the OpenSearch Service Developer Guide.

Type: [OpenSearchNetworkPolicy](api-opensearchnetworkpolicy.md) object

Required: No

**workspace**

This structure contains information about the OpenSearch Service workspace used for this
integration. An OpenSearch Service workspace is the collection of dashboards along with other
OpenSearch Service tools. This workspace was created automatically as part of the
integration setup. For more information, see [Centralized OpenSearch user\
interface (Dashboards) with OpenSearch Service](../../../../services/opensearch-service/latest/developerguide/application.md).

Type: [OpenSearchWorkspace](api-opensearchworkspace.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/opensearchintegrationdetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/opensearchintegrationdetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/opensearchintegrationdetails.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenSearchEncryptionPolicy

OpenSearchLifecyclePolicy

All content copied from https://docs.aws.amazon.com/.
