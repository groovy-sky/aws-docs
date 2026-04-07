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

_Type_: [ApiKeyAuthParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-events-connection-apikeyauthparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BasicAuthParameters`

The authorization parameters for Basic authorization.

_Required_: No

_Type_: [BasicAuthParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-events-connection-basicauthparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectivityParameters`

For private OAuth authentication endpoints. The parameters EventBridge uses to authenticate against the endpoint.

For more information, see [Authorization methods for connections](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-target-connection-auth.html) in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: [ConnectivityParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-events-connection-connectivityparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvocationHttpParameters`

Additional parameters for the connection that are passed through with every invocation to
the HTTP endpoint.

_Required_: No

_Type_: [ConnectionHttpParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-events-connection-connectionhttpparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuthParameters`

The OAuth parameters to use for authorization.

_Required_: No

_Type_: [OAuthParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-events-connection-oauthparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ApiKeyAuthParameters

BasicAuthParameters
