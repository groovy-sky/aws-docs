This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile SAPODataConnectorProfileProperties

The connector-specific profile properties required when using SAPOData.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationHostUrl" : String,
  "ApplicationServicePath" : String,
  "ClientNumber" : String,
  "DisableSSO" : Boolean,
  "LogonLanguage" : String,
  "OAuthProperties" : OAuthProperties,
  "PortNumber" : Integer,
  "PrivateLinkServiceName" : String
}

```

### YAML

```yaml

  ApplicationHostUrl: String
  ApplicationServicePath: String
  ClientNumber: String
  DisableSSO: Boolean
  LogonLanguage: String
  OAuthProperties:
    OAuthProperties
  PortNumber: Integer
  PrivateLinkServiceName: String

```

## Properties

`ApplicationHostUrl`

The location of the SAPOData resource.

_Required_: No

_Type_: String

_Pattern_: `^(https?)://[-a-zA-Z0-9+&amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&amp;@#/%=~_|]`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationServicePath`

The application path to catalog service.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientNumber`

The client number for the client creating the connection.

_Required_: No

_Type_: String

_Pattern_: `^\d{3}$`

_Minimum_: `3`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableSSO`

If you set this parameter to `true`, Amazon AppFlow bypasses the single
sign-on (SSO) settings in your SAP account when it accesses your SAP OData instance.

Whether you need this option depends on the types of credentials that you applied to your
SAP OData connection profile. If your profile uses basic authentication credentials, SAP SSO
can prevent Amazon AppFlow from connecting to your account with your username and
password. In this case, bypassing SSO makes it possible for Amazon AppFlow to connect
successfully. However, if your profile uses OAuth credentials, this parameter has no
affect.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogonLanguage`

The logon language of SAPOData instance.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_]*$`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuthProperties`

The SAPOData OAuth properties required for OAuth type authentication.

_Required_: No

_Type_: [OAuthProperties](aws-properties-appflow-connectorprofile-oauthproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortNumber`

The port number of the SAPOData instance.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateLinkServiceName`

The SAPOData Private Link service name to be used for private data transfers.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SAPODataConnectorProfileCredentials

ServiceNowConnectorProfileCredentials

All content copied from https://docs.aws.amazon.com/.
