This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Gateway CustomJWTAuthorizerConfiguration

Configuration for inbound JWT-based authorization, specifying how incoming requests should be authenticated.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedAudience" : [ String, ... ],
  "AllowedClients" : [ String, ... ],
  "AllowedScopes" : [ String, ... ],
  "CustomClaims" : [ CustomClaimValidationType, ... ],
  "DiscoveryUrl" : String
}

```

### YAML

```yaml

  AllowedAudience:
    - String
  AllowedClients:
    - String
  AllowedScopes:
    - String
  CustomClaims:
    - CustomClaimValidationType
  DiscoveryUrl: String

```

## Properties

`AllowedAudience`

Represents individual audience values that are validated in the incoming JWT token validation process.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedClients`

Represents individual client IDs that are validated in the incoming JWT token validation process.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedScopes`

An array of scopes that are allowed to access the token.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomClaims`

An array of objects that define a custom claim validation name, value, and operation

_Required_: No

_Type_: Array of [CustomClaimValidationType](aws-properties-bedrockagentcore-gateway-customclaimvalidationtype.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DiscoveryUrl`

This URL is used to fetch OpenID Connect configuration or authorization server metadata for validating incoming tokens.

_Required_: Yes

_Type_: String

_Pattern_: `^.+/\.well-known/openid-configuration$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomClaimValidationType

GatewayInterceptorConfiguration

All content copied from https://docs.aws.amazon.com/.
