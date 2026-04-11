This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Connection OAuthParameters

Contains the OAuth authorization parameters to use for the connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationEndpoint" : String,
  "ClientParameters" : ClientParameters,
  "HttpMethod" : String,
  "OAuthHttpParameters" : ConnectionHttpParameters
}

```

### YAML

```yaml

  AuthorizationEndpoint: String
  ClientParameters:
    ClientParameters
  HttpMethod: String
  OAuthHttpParameters:
    ConnectionHttpParameters

```

## Properties

`AuthorizationEndpoint`

The URL to the authorization endpoint when OAuth is specified as the authorization
type.

_Required_: Yes

_Type_: String

_Pattern_: `^((%[0-9A-Fa-f]{2}|[-()_.!~*';/?:@\x26=+$,A-Za-z0-9])+)([).!';/?:,])?$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientParameters`

The client parameters for OAuth authorization.

_Required_: Yes

_Type_: [ClientParameters](aws-properties-events-connection-clientparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpMethod`

The method to use for the authorization request.

_Required_: Yes

_Type_: String

_Allowed values_: `GET | POST | PUT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuthHttpParameters`

Details about the additional
parameters to use for the connection.

_Required_: No

_Type_: [ConnectionHttpParameters](aws-properties-events-connection-connectionhttpparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InvocationConnectivityParameters

Parameter

All content copied from https://docs.aws.amazon.com/.
