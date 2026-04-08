# CreateDataSource

Creates a `DataSource` object.

## Request Syntax

```nohighlight

POST /v1/apis/apiId/datasources HTTP/1.1
Content-type: application/json

{
   "description": "string",
   "dynamodbConfig": {
      "awsRegion": "string",
      "deltaSyncConfig": {
         "baseTableTTL": number,
         "deltaSyncTableName": "string",
         "deltaSyncTableTTL": number
      },
      "tableName": "string",
      "useCallerCredentials": boolean,
      "versioned": boolean
   },
   "elasticsearchConfig": {
      "awsRegion": "string",
      "endpoint": "string"
   },
   "eventBridgeConfig": {
      "eventBusArn": "string"
   },
   "httpConfig": {
      "authorizationConfig": {
         "authorizationType": "string",
         "awsIamConfig": {
            "signingRegion": "string",
            "signingServiceName": "string"
         }
      },
      "endpoint": "string"
   },
   "lambdaConfig": {
      "lambdaFunctionArn": "string"
   },
   "metricsConfig": "string",
   "name": "string",
   "openSearchServiceConfig": {
      "awsRegion": "string",
      "endpoint": "string"
   },
   "relationalDatabaseConfig": {
      "rdsHttpEndpointConfig": {
         "awsRegion": "string",
         "awsSecretStoreArn": "string",
         "databaseName": "string",
         "dbClusterIdentifier": "string",
         "schema": "string"
      },
      "relationalDatabaseSourceType": "string"
   },
   "serviceRoleArn": "string",
   "type": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[apiId](#API_CreateDataSource_RequestSyntax)**

The API ID for the GraphQL API for the `DataSource`.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[description](#API_CreateDataSource_RequestSyntax)**

A description of the `DataSource`.

Type: String

Required: No

**[dynamodbConfig](#API_CreateDataSource_RequestSyntax)**

Amazon DynamoDB settings.

Type: [DynamodbDataSourceConfig](api-dynamodbdatasourceconfig.md) object

Required: No

**[elasticsearchConfig](#API_CreateDataSource_RequestSyntax)**

Amazon OpenSearch Service settings.

As of September 2021, Amazon Elasticsearch service is Amazon OpenSearch Service. This
configuration is deprecated. For new data sources, use [CreateDataSource:openSearchServiceConfig](#appsync-CreateDataSource-request-openSearchServiceConfig) to create an OpenSearch data source.

Type: [ElasticsearchDataSourceConfig](api-elasticsearchdatasourceconfig.md) object

Required: No

**[eventBridgeConfig](#API_CreateDataSource_RequestSyntax)**

Amazon EventBridge settings.

Type: [EventBridgeDataSourceConfig](api-eventbridgedatasourceconfig.md) object

Required: No

**[httpConfig](#API_CreateDataSource_RequestSyntax)**

HTTP endpoint settings.

Type: [HttpDataSourceConfig](api-httpdatasourceconfig.md) object

Required: No

**[lambdaConfig](#API_CreateDataSource_RequestSyntax)**

AWS Lambda settings.

Type: [LambdaDataSourceConfig](api-lambdadatasourceconfig.md) object

Required: No

**[metricsConfig](#API_CreateDataSource_RequestSyntax)**

Enables or disables enhanced data source metrics for specified data sources. Note that
`metricsConfig` won't be used unless the
`dataSourceLevelMetricsBehavior` value is set to
`PER_DATA_SOURCE_METRICS`. If the `dataSourceLevelMetricsBehavior`
is set to `FULL_REQUEST_DATA_SOURCE_METRICS` instead, `metricsConfig`
will be ignored. However, you can still set its value.

`metricsConfig` can be `ENABLED` or `DISABLED`.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**[name](#API_CreateDataSource_RequestSyntax)**

A user-supplied name for the `DataSource`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 65536.

Pattern: `[_A-Za-z][_0-9A-Za-z]*`

Required: Yes

**[openSearchServiceConfig](#API_CreateDataSource_RequestSyntax)**

Amazon OpenSearch Service settings.

Type: [OpenSearchServiceDataSourceConfig](api-opensearchservicedatasourceconfig.md) object

Required: No

**[relationalDatabaseConfig](#API_CreateDataSource_RequestSyntax)**

Relational database settings.

Type: [RelationalDatabaseDataSourceConfig](api-relationaldatabasedatasourceconfig.md) object

Required: No

**[serviceRoleArn](#API_CreateDataSource_RequestSyntax)**

The AWS Identity and Access Management (IAM) service role Amazon Resource Name (ARN)
for the data source. The system assumes this role when accessing the data source.

Type: String

Required: No

**[type](#API_CreateDataSource_RequestSyntax)**

The type of the `DataSource`.

Type: String

Valid Values: `AWS_LAMBDA | AMAZON_DYNAMODB | AMAZON_ELASTICSEARCH | NONE | HTTP | RELATIONAL_DATABASE | AMAZON_OPENSEARCH_SERVICE | AMAZON_EVENTBRIDGE | AMAZON_BEDROCK_RUNTIME`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "dataSource": {
      "dataSourceArn": "string",
      "description": "string",
      "dynamodbConfig": {
         "awsRegion": "string",
         "deltaSyncConfig": {
            "baseTableTTL": number,
            "deltaSyncTableName": "string",
            "deltaSyncTableTTL": number
         },
         "tableName": "string",
         "useCallerCredentials": boolean,
         "versioned": boolean
      },
      "elasticsearchConfig": {
         "awsRegion": "string",
         "endpoint": "string"
      },
      "eventBridgeConfig": {
         "eventBusArn": "string"
      },
      "httpConfig": {
         "authorizationConfig": {
            "authorizationType": "string",
            "awsIamConfig": {
               "signingRegion": "string",
               "signingServiceName": "string"
            }
         },
         "endpoint": "string"
      },
      "lambdaConfig": {
         "lambdaFunctionArn": "string"
      },
      "metricsConfig": "string",
      "name": "string",
      "openSearchServiceConfig": {
         "awsRegion": "string",
         "endpoint": "string"
      },
      "relationalDatabaseConfig": {
         "rdsHttpEndpointConfig": {
            "awsRegion": "string",
            "awsSecretStoreArn": "string",
            "databaseName": "string",
            "dbClusterIdentifier": "string",
            "schema": "string"
         },
         "relationalDatabaseSourceType": "string"
      },
      "serviceRoleArn": "string",
      "type": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[dataSource](#API_CreateDataSource_ResponseSyntax)**

The `DataSource` object.

Type: [DataSource](api-datasource.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The request is not well formed. For example, a value is invalid or a required field is
missing. Check the field values, and then try again.

**detail**

Provides further details for the reason behind the bad request. For reason type
`CODE_ERROR`, the detail will contain a list of code errors.

**reason**

Provides context for the cause of the bad request. The only supported value is
`CODE_ERROR`.

HTTP Status Code: 400

**ConcurrentModificationException**

Another modification is in progress at this time and it must complete before you can
make your change.

HTTP Status Code: 409

**InternalFailureException**

An internal AWS AppSync error occurred. Try your request again.

HTTP Status Code: 500

**NotFoundException**

The resource specified in the request was not found. Check the resource, and then try
again.

HTTP Status Code: 404

**UnauthorizedException**

You aren't authorized to perform this operation.

HTTP Status Code: 401

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appsync-2017-07-25/createdatasource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appsync-2017-07-25/createdatasource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/createdatasource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appsync-2017-07-25/createdatasource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/createdatasource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appsync-2017-07-25/createdatasource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appsync-2017-07-25/createdatasource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appsync-2017-07-25/createdatasource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appsync-2017-07-25/createdatasource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/createdatasource.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateChannelNamespace

CreateDomainName
