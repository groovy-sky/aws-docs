This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Connection AuthParameters

Tthe authorization parameters to use for the connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiKeyAuthParameters" : ApiKeyAuthParameters,
  "BasicAuthParameters" : BasicAuthParameters,
  "ConnectivityParameters" : ConnectivityParameters,
  "InvocationHttpParameters" : ConnectionHttpParameters,
  "OAuthParameters" : OAuthParameters
}

```

### YAML

```yaml

  ApiKeyAuthParameters:
    ApiKeyAuthParameters
  BasicAuthParameters:
    BasicAuthParameters
  ConnectivityParameters:
    ConnectivityParameters
  InvocationHttpParameters:
    ConnectionHttpParameters
  OAuthParameters:
    OAuthParameters

```

## Properties

`ApiKeyAuthParameters`

The API Key parameters to use for authorization.

_Required_: No

_Type_: [ApiKeyAuthParameters](aws-properties-events-connection-apikeyauthparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BasicAuthParameters`

The authorization parameters for Basic authorization.

_Required_: No

_Type_: [BasicAuthParameters](aws-properties-events-connection-basicauthparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectivityParameters`

For private OAuth authentication endpoints. The parameters EventBridge uses to authenticate against the endpoint.

For more information, see [Authorization methods for connections](../../../eventbridge/latest/userguide/eb-target-connection-auth.md) in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: [ConnectivityParameters](aws-properties-events-connection-connectivityparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvocationHttpParameters`

Additional parameters for the connection that are passed through with every invocation to
the HTTP endpoint.

_Required_: No

_Type_: [ConnectionHttpParameters](aws-properties-events-connection-connectionhttpparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuthParameters`

The OAuth parameters to use for authorization.

_Required_: No

_Type_: [OAuthParameters](aws-properties-events-connection-oauthparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiKeyAuthParameters

BasicAuthParameters

All content copied from https://docs.aws.amazon.com/.
