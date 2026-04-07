This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::GraphQLApi

The `AWS::AppSync::GraphQLApi` resource creates a new AWS AppSync
GraphQL API. This is the top-level construct for your application. For more information,
see [Quick\
Start](https://docs.aws.amazon.com/appsync/latest/devguide/quickstart.html) in the _AWS AppSync Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppSync::GraphQLApi",
  "Properties" : {
      "AdditionalAuthenticationProviders" : [ AdditionalAuthenticationProvider, ... ],
      "ApiType" : String,
      "AuthenticationType" : String,
      "EnhancedMetricsConfig" : EnhancedMetricsConfig,
      "EnvironmentVariables" : {Key: Value, ...},
      "IntrospectionConfig" : String,
      "LambdaAuthorizerConfig" : LambdaAuthorizerConfig,
      "LogConfig" : LogConfig,
      "MergedApiExecutionRoleArn" : String,
      "Name" : String,
      "OpenIDConnectConfig" : OpenIDConnectConfig,
      "OwnerContact" : String,
      "QueryDepthLimit" : Integer,
      "ResolverCountLimit" : Integer,
      "Tags" : [ Tag, ... ],
      "UserPoolConfig" : UserPoolConfig,
      "Visibility" : String,
      "XrayEnabled" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::AppSync::GraphQLApi
Properties:
  AdditionalAuthenticationProviders:
    - AdditionalAuthenticationProvider
  ApiType: String
  AuthenticationType: String
  EnhancedMetricsConfig:
    EnhancedMetricsConfig
  EnvironmentVariables:
    Key: Value
  IntrospectionConfig: String
  LambdaAuthorizerConfig:
    LambdaAuthorizerConfig
  LogConfig:
    LogConfig
  MergedApiExecutionRoleArn: String
  Name: String
  OpenIDConnectConfig:
    OpenIDConnectConfig
  OwnerContact: String
  QueryDepthLimit: Integer
  ResolverCountLimit: Integer
  Tags:
    - Tag
  UserPoolConfig:
    UserPoolConfig
  Visibility: String
  XrayEnabled: Boolean

```

## Properties

`AdditionalAuthenticationProviders`

A list of additional authentication providers for the `GraphqlApi` API.

_Required_: No

_Type_: Array of [AdditionalAuthenticationProvider](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-graphqlapi-additionalauthenticationprovider.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApiType`

The value that indicates whether the GraphQL API is a standard API
( `GRAPHQL`) or merged API ( `MERGED`).

**WARNING**: If the `ApiType` has not been
defined, **explicitly** setting it to `GRAPHQL`
in a template/stack update will result in an API replacement and new DNS values.

The following values are valid:

`GRAPHQL | MERGED`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationType`

Security configuration for your GraphQL API. For allowed values (such as
`API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`,
`OPENID_CONNECT`, or `AWS_LAMBDA`), see [Security](https://docs.aws.amazon.com/appsync/latest/devguide/security.html) in
the _AWS AppSync Developer Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `API_KEY | AWS_IAM | AMAZON_COGNITO_USER_POOLS | OPENID_CONNECT | AWS_LAMBDA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnhancedMetricsConfig`

Enables and controls the enhanced metrics feature. Enhanced metrics emit granular data
on API usage and performance such as AppSync request and error counts, latency, and
cache hits/misses. All enhanced metric data is sent to your CloudWatch account, and you
can configure the types of data that will be sent.

Enhanced metrics can be configured at the resolver, data source, and operation levels.
For more information, see [Monitoring and\
logging](https://docs.aws.amazon.com/appsync/latest/devguide/monitoring.html#cw-metrics) in the _AWS AppSync User Guide_.

_Required_: No

_Type_: [EnhancedMetricsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-graphqlapi-enhancedmetricsconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentVariables`

A map containing the list of resources with their properties and environment
variables. For more information, see [Environmental\
variables](https://docs.aws.amazon.com/appsync/latest/devguide/environmental-variables.html).

_Pattern_: `^[A-Za-z]+\\w*$\\`

_Minimum_: 2

_Maximum_: 64

_Required_: No

_Type_: Object of String

_Pattern_: `^[A-Za-z]+\w*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntrospectionConfig`

Sets the value of the GraphQL API to enable ( `ENABLED`) or disable
( `DISABLED`) introspection. If no value is provided, the introspection
configuration will be set to `ENABLED` by default. This field will produce an
error if the operation attempts to use the introspection feature while this field is
disabled.

For more information about introspection, see [GraphQL introspection](https://graphql.org/learn/introspection).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaAuthorizerConfig`

A `LambdaAuthorizerConfig` holds configuration on how to authorize AWS AppSync API access when using the `AWS_LAMBDA` authorizer mode.
Be aware that an AWS AppSync API may have only one Lambda authorizer
configured at a time.

_Required_: No

_Type_: [LambdaAuthorizerConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-graphqlapi-lambdaauthorizerconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogConfig`

The Amazon CloudWatch Logs configuration.

_Required_: No

_Type_: [LogConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-graphqlapi-logconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MergedApiExecutionRoleArn`

The AWS Identity and Access Management service role ARN for a merged
API. The AppSync service assumes this role on behalf of the Merged API to validate
access to source APIs at runtime and to prompt the `AUTO_MERGE` to update the
merged API endpoint with the source API changes automatically.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The API name.

_Required_: Yes

_Type_: String

_Pattern_: `[_A-Za-z][_0-9A-Za-z]*`

_Minimum_: `1`

_Maximum_: `65536`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenIDConnectConfig`

The OpenID Connect configuration.

_Required_: No

_Type_: [OpenIDConnectConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-graphqlapi-openidconnectconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OwnerContact`

The owner contact information for an API resource.

This field accepts any string input with a length of 0 - 256 characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryDepthLimit`

The maximum depth a query can have in a single request. Depth refers to the amount of
nested levels allowed in the body of query. The default value is `0` (or
unspecified), which indicates there's no depth limit. If you set a limit, it can be
between `1` and `75` nested levels. This field will produce a
limit error if the operation falls out of bounds. Note that fields can still be set to
nullable or non-nullable. If a non-nullable field produces an error, the error will be
thrown upwards to the first nullable field available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResolverCountLimit`

The maximum number of resolvers that can be invoked in a single request. The default
value is `0` (or unspecified), which will set the limit to
`10000`. When specified, the limit value can be between `1` and
`10000`. This field will produce a limit error if the operation falls out
of bounds.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An arbitrary set of tags (key-value pairs) for this GraphQL API.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-graphqlapi-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolConfig`

Optional authorization configuration for using Amazon Cognito user pools with your
GraphQL endpoint.

_Required_: No

_Type_: [UserPoolConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-graphqlapi-userpoolconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

Sets the scope of the GraphQL API to public ( `GLOBAL`) or private
( `PRIVATE`). By default, the scope is set to `Global` if no
value is provided.

**WARNING**: If `Visibility` has not been
defined, **explicitly** setting it to `GLOBAL`
in a template/stack update will result in an API replacement and new DNS values.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XrayEnabled`

A flag indicating whether to use AWS X-Ray tracing for this `GraphqlApi`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of an `AWS::AppSync::GraphQLApi` resource to
the intrinsic `Ref` function, the function returns the ARN of the GraphQL
API, such as `arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid`.

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The
following are the available attributes and sample return values.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt).

`ApiId`

Unique AWS AppSync GraphQL API identifier.

`Arn`

The Amazon Resource Name (ARN) of the API key, such as
`arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid`.

`GraphQLDns`

The fully qualified domain name (FQDN) of the endpoint URL of your GraphQL API.

`GraphQLEndpointArn`

The GraphQL endpoint ARN.

`GraphQLUrl`

The Endpoint URL of your GraphQL API.

`RealtimeDns`

The fully qualified domain name (FQDN) of the real-time endpoint URL of your GraphQL
API.

`RealtimeUrl`

The GraphQL API real-time endpoint URL. For more information, see [Discovering the real-time endpoint from the GraphQL endpoint](https://docs.aws.amazon.com/appsync/latest/devguide/real-time-websocket-client.html#handshake-details-to-establish-the-websocket-connection).

## Examples

### GraphQL API Creation Example

The following example creates a GraphQL API.

#### YAML

```yaml

Parameters:
  graphQlApiName:
    Type: String
  userPoolId:
    Type: String
  userPoolAwsRegion:
    Type: String
  defaultAction:
    Type: String
Resources:
  GraphQLApi:
    Type: AWS::AppSync::GraphQLApi
    Properties:
      Name:
	Ref: graphQlApiName
      AuthenticationType: "AMAZON_COGNITO_USER_POOLS"
      UserPoolConfig:
        UserPoolId:
          Ref: userPoolId
        AwsRegion:
          Ref: userPoolAwsRegion
        DefaultAction:
          Ref: defaultAction

```

#### JSON

```json

{
  "Parameters": {
    "graphQlApiName": {
      "Type": "String"
    },
    "userPoolId": {
      "Type": "String"
    },
    "userPoolAwsRegion": {
      "Type": "String"
    },
    "defaultAction": {
      "Type": "String"
    }
  },
  "Resources": {
    "GraphQLApi": {
      "Type": "AWS::AppSync::GraphQLApi",
      "Properties": {
        "Name": {
          "Ref": "graphQlApiName"
        },
        "AuthenticationType": "AMAZON_COGNITO_USER_POOLS",
        "UserPoolConfig": {
          "UserPoolId": {
            "Ref": "userPoolId"
          },
          "AwsRegion": {
            "Ref": "userPoolAwsRegion"
          },
          "DefaultAction": {
            "Ref": "defaultAction"
          }
        }
      }
    }
  }
}

```

## See also

- [CreateGraphqlApi](https://docs.aws.amazon.com/appsync/latest/APIReference/API_CreateGraphqlApi.html) operation in the _AWS AppSync API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SyncConfig

AdditionalAuthenticationProvider
