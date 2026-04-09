# AdditionalAuthenticationProvider

Describes an additional authentication provider.

## Contents

**authenticationType**

The authentication type: API key, AWS Identity and Access Management (IAM), OpenID
Connect (OIDC), Amazon Cognito user pools, or AWS Lambda.

Type: String

Valid Values: `API_KEY | AWS_IAM | AMAZON_COGNITO_USER_POOLS | OPENID_CONNECT | AWS_LAMBDA`

Required: No

**lambdaAuthorizerConfig**

Configuration for Lambda function authorization.

Type: [LambdaAuthorizerConfig](api-lambdaauthorizerconfig.md) object

Required: No

**openIDConnectConfig**

The OIDC configuration.

Type: [OpenIDConnectConfig](api-openidconnectconfig.md) object

Required: No

**userPoolConfig**

The Amazon Cognito user pool configuration.

Type: [CognitoUserPoolConfig](api-cognitouserpoolconfig.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/additionalauthenticationprovider.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/additionalauthenticationprovider.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/additionalauthenticationprovider.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data Types

Api

All content copied from https://docs.aws.amazon.com/.
