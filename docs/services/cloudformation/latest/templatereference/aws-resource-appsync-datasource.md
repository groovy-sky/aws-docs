This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::DataSource

The `AWS::AppSync::DataSource` resource creates data sources for resolvers
in AWS AppSync to connect to, such as Amazon DynamoDB, AWS Lambda, and Amazon OpenSearch Service. Resolvers use these data sources to
fetch data when clients make GraphQL calls.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppSync::DataSource",
  "Properties" : {
      "ApiId" : String,
      "Description" : String,
      "DynamoDBConfig" : DynamoDBConfig,
      "EventBridgeConfig" : EventBridgeConfig,
      "HttpConfig" : HttpConfig,
      "LambdaConfig" : LambdaConfig,
      "MetricsConfig" : String,
      "Name" : String,
      "OpenSearchServiceConfig" : OpenSearchServiceConfig,
      "RelationalDatabaseConfig" : RelationalDatabaseConfig,
      "ServiceRoleArn" : String,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppSync::DataSource
Properties:
  ApiId: String
  Description: String
  DynamoDBConfig:
    DynamoDBConfig
  EventBridgeConfig:
    EventBridgeConfig
  HttpConfig:
    HttpConfig
  LambdaConfig:
    LambdaConfig
  MetricsConfig: String
  Name: String
  OpenSearchServiceConfig:
    OpenSearchServiceConfig
  RelationalDatabaseConfig:
    RelationalDatabaseConfig
  ServiceRoleArn: String
  Type: String

```

## Properties

`ApiId`

Unique AWS AppSync GraphQL API identifier where this data source will be
created.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the data source.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamoDBConfig`

AWS Region and TableName for an Amazon DynamoDB table in
your account.

_Required_: No

_Type_: [DynamoDBConfig](aws-properties-appsync-datasource-dynamodbconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventBridgeConfig`

An EventBridge configuration that contains a valid ARN of an event bus.

_Required_: No

_Type_: [EventBridgeConfig](aws-properties-appsync-datasource-eventbridgeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpConfig`

Endpoints for an HTTP data source.

_Required_: No

_Type_: [HttpConfig](aws-properties-appsync-datasource-httpconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaConfig`

An ARN of a Lambda function in valid ARN format. This can be the ARN of a Lambda
function that exists in the current account or in another account.

_Required_: No

_Type_: [LambdaConfig](aws-properties-appsync-datasource-lambdaconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricsConfig`

Enables or disables enhanced data source metrics for specified data sources. Note that
`MetricsConfig` won't be used unless the
`dataSourceLevelMetricsBehavior` value is set to
`PER_DATA_SOURCE_METRICS`. If the
`dataSourceLevelMetricsBehavior` is set to
`FULL_REQUEST_DATA_SOURCE_METRICS` instead, `MetricsConfig`
will be ignored. However, you can still set its value.

`MetricsConfig` can be `ENABLED` or
`DISABLED`.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Friendly name for you to identify your AppSync data source after creation.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OpenSearchServiceConfig`

AWS Region and Endpoints for an Amazon OpenSearch Service domain in
your account.

_Required_: No

_Type_: [OpenSearchServiceConfig](aws-properties-appsync-datasource-opensearchserviceconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelationalDatabaseConfig`

Relational Database configuration of the relational database data source.

_Required_: No

_Type_: [RelationalDatabaseConfig](aws-properties-appsync-datasource-relationaldatabaseconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceRoleArn`

The AWS Identity and Access Management service role ARN for the data source. The system assumes
this role when accessing the data source.

Required if `Type` is specified as `AWS_LAMBDA`,
`AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`,
`AMAZON_EVENTBRIDGE`, `AMAZON_OPENSEARCH_SERVICE`,
`RELATIONAL_DATABASE`, or `AMAZON_BEDROCK_RUNTIME`.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the data source.

- **AWS\_LAMBDA**: The data source is an AWS Lambda function.

- **AMAZON\_DYNAMODB**: The data source is an
Amazon DynamoDB table.

- **AMAZON\_ELASTICSEARCH**: The data source is an
Amazon OpenSearch Service domain.

- **AMAZON\_EVENTBRIDGE**: The data source is an
Amazon EventBridge event bus.

- **AMAZON\_OPENSEARCH\_SERVICE**: The data source is
an Amazon OpenSearch Service domain.

- **AMAZON\_BEDROCK\_RUNTIME**: The data source is
the Amazon Bedrock runtime.

- **NONE**: There is no data source. This type is
used when you wish to invoke a GraphQL operation without connecting to a data
source, such as performing data transformation with resolvers or triggering a
subscription to be invoked from a mutation.

- **HTTP**: The data source is an HTTP
endpoint.

- **RELATIONAL\_DATABASE**: The data source is a
relational database.

_Required_: Yes

_Type_: String

_Allowed values_: `AWS_LAMBDA | AMAZON_DYNAMODB | AMAZON_ELASTICSEARCH | NONE | HTTP | RELATIONAL_DATABASE | AMAZON_OPENSEARCH_SERVICE | AMAZON_EVENTBRIDGE | AMAZON_BEDROCK_RUNTIME`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of an `AWS::AppSync::DataSource` resource to
the intrinsic `Ref` function, the function returns the ARN of the Data
Source, such as `
                arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/datasources/datasourcename`.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The
following are the available attributes and sample return values.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`DataSourceArn`

The Amazon Resource Name (ARN) of the API key, such as
`arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/datasources/datasourcename`.

`Name`

Friendly name for you to identify your AWS AppSync data source after
creation.

## Examples

### Data Source Creation Example

The following example creates a data source and associates it with an existing
GraphQL API by passing the GraphQL API ID as a parameter.

#### YAML

```yaml

Parameters:
  graphQlApiId:
    Type: String
  dataSourceName:
    Type: String
  dataSourceDescription:
    Type: String
  serviceRoleArn:
    Type: String
  lambdaFunctionArn:
    Type: String
Resources:
  DataSource:
    Type: AWS::AppSync::DataSource
    Properties:
      ApiId:
        Ref: graphQlApiId
      Name:
        Ref: dataSourceName
      Description:
        Ref: dataSourceDescription
      Type: "AWS_LAMBDA"
      ServiceRoleArn:
        Ref: serviceRoleArn
      LambdaConfig:
        LambdaFunctionArn:
          Ref: lambdaFunctionArn

```

#### JSON

```json

{
  "Parameters": {
    "graphQlApiId": {
      "Type": "String"
    },
    "dataSourceName": {
      "Type": "String"
    },
    "dataSourceDescription": {
      "Type": "String"
    },
    "serviceRoleArn": {
      "Type": "String"
    },
    "lambdaFunctionArn": {
      "Type": "String"
    }
  },
  "Resources": {
    "DataSource": {
      "Type": "AWS::AppSync::DataSource",
      "Properties": {
        "ApiId": {
          "Ref": "graphQlApiId"
        },
        "Name": {
          "Ref": "dataSourceName"
        },
        "Description": {
          "Ref": "dataSourceDescription"
        },
        "Type": "AWS_LAMBDA",
        "ServiceRoleArn": {
           "Ref": "serviceRoleArn"
        },
        "LambdaConfig": {
          "LambdaFunctionArn": {
            "Ref": "lambdaFunctionArn"
          }
        }
      }
    }
  }
}

```

## See also

- [CreateDataSource](../../../../reference/appsync/latest/apireference/api-createdatasource.md) operation in the _AWS AppSync API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AuthorizationConfig

All content copied from https://docs.aws.amazon.com/.
