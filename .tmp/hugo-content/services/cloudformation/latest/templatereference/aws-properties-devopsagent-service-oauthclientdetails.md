This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service OAuthClientDetails

OAuth client credentials shared by Dynatrace and ServiceNow service types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientId" : String,
  "ClientName" : String,
  "ClientSecret" : String,
  "ExchangeParameters" : Json
}

```

### YAML

```yaml

  ClientId: String
  ClientName: String
  ClientSecret: String
  ExchangeParameters: Json

```

## Properties

`ClientId`

The OAuth client ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClientName`

A friendly name for the OAuth client.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClientSecret`

The OAuth client secret.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExchangeParameters`

Additional parameters for the OAuth token exchange request.

_Required_: No

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NewRelicServiceDetails

RegisteredDynatraceDetails

All content copied from https://docs.aws.amazon.com/.
