This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::GraphQLApi AdditionalAuthenticationProvider

Describes an additional authentication provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationType" : String,
  "LambdaAuthorizerConfig" : LambdaAuthorizerConfig,
  "OpenIDConnectConfig" : OpenIDConnectConfig,
  "UserPoolConfig" : CognitoUserPoolConfig
}

```

### YAML

```yaml

  AuthenticationType: String
  LambdaAuthorizerConfig:
    LambdaAuthorizerConfig
  OpenIDConnectConfig:
    OpenIDConnectConfig
  UserPoolConfig:
    CognitoUserPoolConfig

```

## Properties

`AuthenticationType`

The authentication type for API key, AWS Identity and Access Management, OIDC, Amazon Cognito user pools, or AWS Lambda.

Valid Values: `API_KEY` \| `AWS_IAM` \|
`OPENID_CONNECT` \| `AMAZON_COGNITO_USER_POOLS` \|
`AWS_LAMBDA`

_Required_: Yes

_Type_: String

_Allowed values_: `API_KEY | AWS_IAM | AMAZON_COGNITO_USER_POOLS | OPENID_CONNECT | AWS_LAMBDA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaAuthorizerConfig`

Configuration for AWS Lambda function authorization.

_Required_: No

_Type_: [LambdaAuthorizerConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-graphqlapi-lambdaauthorizerconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenIDConnectConfig`

The OIDC configuration.

_Required_: No

_Type_: [OpenIDConnectConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-graphqlapi-openidconnectconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolConfig`

The Amazon Cognito user pool configuration.

_Required_: No

_Type_: [CognitoUserPoolConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-graphqlapi-cognitouserpoolconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppSync::GraphQLApi

CognitoUserPoolConfig
