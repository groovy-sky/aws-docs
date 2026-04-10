This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::ActionConnector ClientCredentialsGrantMetadata

Configuration for OAuth 2.0 client credentials grant authentication, including client ID, client secret, token endpoint, and optional scopes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaseEndpoint" : String,
  "ClientCredentialsDetails" : ClientCredentialsDetails,
  "ClientCredentialsSource" : String
}

```

### YAML

```yaml

  BaseEndpoint: String
  ClientCredentialsDetails:
    ClientCredentialsDetails
  ClientCredentialsSource: String

```

## Properties

`BaseEndpoint`

The base endpoint URL for the external service.

_Required_: Yes

_Type_: String

_Pattern_: `^https://.*`

_Minimum_: `1`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientCredentialsDetails`

The detailed client credentials configuration including client ID, client secret, and token endpoint.

_Required_: No

_Type_: [ClientCredentialsDetails](aws-properties-quicksight-actionconnector-clientcredentialsdetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientCredentialsSource`

The source of the client credentials configuration.

_Required_: No

_Type_: String

_Allowed values_: `PLAIN_CREDENTIALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientCredentialsGrantDetails

IAMConnectionMetadata

All content copied from https://docs.aws.amazon.com/.
