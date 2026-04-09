# GraphqlApi

Describes a GraphQL API.

## Contents

**additionalAuthenticationProviders**

A list of additional authentication providers for the `GraphqlApi`
API.

Type: Array of [AdditionalAuthenticationProvider](api-additionalauthenticationprovider.md) objects

Required: No

**apiId**

The API ID.

Type: String

Required: No

**apiType**

The value that indicates whether the GraphQL API is a standard API
( `GRAPHQL`) or merged API ( `MERGED`).

Type: String

Valid Values: `GRAPHQL | MERGED`

Required: No

**arn**

The Amazon Resource Name (ARN).

Type: String

Required: No

**authenticationType**

The authentication type.

Type: String

Valid Values: `API_KEY | AWS_IAM | AMAZON_COGNITO_USER_POOLS | OPENID_CONNECT | AWS_LAMBDA`

Required: No

**dns**

The DNS records for the API.

Type: String to string map

Required: No

**enhancedMetricsConfig**

The `enhancedMetricsConfig` object.

Type: [EnhancedMetricsConfig](api-enhancedmetricsconfig.md) object

Required: No

**introspectionConfig**

Sets the value of the GraphQL API to enable ( `ENABLED`) or disable
( `DISABLED`) introspection. If no value is provided, the introspection
configuration will be set to `ENABLED` by default. This field will produce an
error if the operation attempts to use the introspection feature while this field is
disabled.

For more information about introspection, see [GraphQL introspection](https://graphql.org/learn/introspection).

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**lambdaAuthorizerConfig**

Configuration for Lambda function authorization.

Type: [LambdaAuthorizerConfig](api-lambdaauthorizerconfig.md) object

Required: No

**logConfig**

The Amazon CloudWatch Logs configuration.

Type: [LogConfig](api-logconfig.md) object

Required: No

**mergedApiExecutionRoleArn**

The Identity and Access Management service role ARN for a merged API. The AppSync
service assumes this role on behalf of the Merged API to validate access to source APIs at
runtime and to prompt the `AUTO_MERGE` to update the merged API endpoint with
the source API changes automatically.

Type: String

Required: No

**name**

The API name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 65536.

Pattern: `[_A-Za-z][_0-9A-Za-z]*`

Required: No

**openIDConnectConfig**

The OpenID Connect configuration.

Type: [OpenIDConnectConfig](api-openidconnectconfig.md) object

Required: No

**owner**

The account owner of the GraphQL API.

Type: String

Required: No

**ownerContact**

The owner contact information for an API resource.

This field accepts any string input with a length of 0 - 256 characters.

Type: String

Required: No

**queryDepthLimit**

The maximum depth a query can have in a single request. Depth refers to the amount of
nested levels allowed in the body of query. The default value is `0` (or
unspecified), which indicates there's no depth limit. If you set a limit, it can be between
`1` and `75` nested levels. This field will produce a limit error
if the operation falls out of bounds.

Note that fields can still be set to nullable or non-nullable. If a non-nullable field
produces an error, the error will be thrown upwards to the first nullable field
available.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 75.

Required: No

**resolverCountLimit**

The maximum number of resolvers that can be invoked in a single request. The default
value is `0` (or unspecified), which will set the limit to `10000`.
When specified, the limit value can be between `1` and `10000`. This
field will produce a limit error if the operation falls out of bounds.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 10000.

Required: No

**tags**

The tags.

Type: String to string map

Map Entries: Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^(?!aws:)[ a-zA-Z+-=._:/]+$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `^[\s\w+-=\.:/@]*$`

Required: No

**uris**

The URIs.

Type: String to string map

Required: No

**userPoolConfig**

The Amazon Cognito user pool configuration.

Type: [UserPoolConfig](api-userpoolconfig.md) object

Required: No

**visibility**

Sets the value of the GraphQL API to public ( `GLOBAL`) or private
( `PRIVATE`). If no value is provided, the visibility will be set to
`GLOBAL` by default. This value cannot be changed once the API has been
created.

Type: String

Valid Values: `GLOBAL | PRIVATE`

Required: No

**wafWebAclArn**

The ARN of the AWS WAF access control list (ACL) associated with this
`GraphqlApi`, if one exists.

Type: String

Required: No

**xrayEnabled**

A flag indicating whether to use AWS X-Ray tracing for this
`GraphqlApi`.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/graphqlapi.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/graphqlapi.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/graphqlapi.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FunctionConfiguration

HandlerConfig

All content copied from https://docs.aws.amazon.com/.
