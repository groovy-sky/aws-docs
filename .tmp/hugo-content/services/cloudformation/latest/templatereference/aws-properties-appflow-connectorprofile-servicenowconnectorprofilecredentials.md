This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile ServiceNowConnectorProfileCredentials

The connector-specific profile credentials required when using ServiceNow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OAuth2Credentials" : OAuth2Credentials,
  "Password" : String,
  "Username" : String
}

```

### YAML

```yaml

  OAuth2Credentials:
    OAuth2Credentials
  Password: String
  Username: String

```

## Properties

`OAuth2Credentials`

The OAuth 2.0 credentials required to authenticate the user.

_Required_: No

_Type_: [OAuth2Credentials](aws-properties-appflow-connectorprofile-oauth2credentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Password`

The password that corresponds to the user name.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The name of the user.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ServiceNowConnectorProfileCredentials](../../../../reference/appflow/1-0/apireference/api-servicenowconnectorprofilecredentials.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SAPODataConnectorProfileProperties

ServiceNowConnectorProfileProperties

All content copied from https://docs.aws.amazon.com/.
