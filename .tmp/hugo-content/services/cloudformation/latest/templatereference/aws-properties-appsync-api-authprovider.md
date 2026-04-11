This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::Api AuthProvider

Describes an authorization provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthType" : String,
  "CognitoConfig" : CognitoConfig,
  "LambdaAuthorizerConfig" : LambdaAuthorizerConfig,
  "OpenIDConnectConfig" : OpenIDConnectConfig
}

```

### YAML

```yaml

  AuthType: String
  CognitoConfig:
    CognitoConfig
  LambdaAuthorizerConfig:
    LambdaAuthorizerConfig
  OpenIDConnectConfig:
    OpenIDConnectConfig

```

## Properties

`AuthType`

The authorization type.

_Required_: Yes

_Type_: String

_Allowed values_: `AMAZON_COGNITO_USER_POOLS | AWS_IAM | API_KEY | OPENID_CONNECT | AWS_LAMBDA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CognitoConfig`

Describes an Amazon Cognito user pool configuration.

_Required_: No

_Type_: [CognitoConfig](aws-properties-appsync-api-cognitoconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaAuthorizerConfig`

A `LambdaAuthorizerConfig` specifies how to authorize AWS AppSync API access when
using the `AWS_LAMBDA` authorizer mode. Be aware that an AWS AppSync API can have only
one AWS Lambda authorizer configured at a time.

_Required_: No

_Type_: [LambdaAuthorizerConfig](aws-properties-appsync-api-lambdaauthorizerconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenIDConnectConfig`

Describes an OpenID Connect (OIDC) configuration.

_Required_: No

_Type_: [OpenIDConnectConfig](aws-properties-appsync-api-openidconnectconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthMode

CognitoConfig

All content copied from https://docs.aws.amazon.com/.
