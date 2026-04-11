This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile SAPODataConnectorProfileCredentials

The connector-specific profile credentials required when using SAPOData.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BasicAuthCredentials" : BasicAuthCredentials,
  "OAuthCredentials" : OAuthCredentials
}

```

### YAML

```yaml

  BasicAuthCredentials:
    BasicAuthCredentials
  OAuthCredentials:
    OAuthCredentials

```

## Properties

`BasicAuthCredentials`

The SAPOData basic authentication credentials.

_Required_: No

_Type_: [BasicAuthCredentials](aws-properties-appflow-connectorprofile-basicauthcredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuthCredentials`

The SAPOData OAuth type authentication credentials.

_Required_: No

_Type_: [OAuthCredentials](aws-properties-appflow-connectorprofile-oauthcredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SalesforceConnectorProfileProperties

SAPODataConnectorProfileProperties

All content copied from https://docs.aws.amazon.com/.
