This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::Authorizer JWTConfiguration

The `JWTConfiguration` property specifies the configuration of a JWT
authorizer. Required for the `JWT` authorizer type. Supported only for
HTTP APIs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Audience" : [ String, ... ],
  "Issuer" : String
}

```

### YAML

```yaml

  Audience:
    - String
  Issuer: String

```

## Properties

`Audience`

A list of the intended recipients of the JWT. A valid JWT must provide an
`aud` that matches at least one entry in this list. See [RFC 7519](https://tools.ietf.org/html/rfc7519).
Required for the `JWT` authorizer type. Supported only for HTTP APIs.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Issuer`

The base domain of the identity provider that issues JSON Web Tokens. For example,
an Amazon Cognito user pool has the following format:
`https://cognito-idp.{region}.amazonaws.com/{userPoolId}`.
Required for the `JWT` authorizer type. Supported only for HTTP APIs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGatewayV2::Authorizer

AWS::ApiGatewayV2::Deployment
