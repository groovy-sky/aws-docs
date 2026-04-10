This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource SharePointSourceConfiguration

The endpoint information to connect to your SharePoint data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthType" : String,
  "CredentialsSecretArn" : String,
  "Domain" : String,
  "HostType" : String,
  "SiteUrls" : [ String, ... ],
  "TenantId" : String
}

```

### YAML

```yaml

  AuthType: String
  CredentialsSecretArn: String
  Domain: String
  HostType: String
  SiteUrls:
    - String
  TenantId: String

```

## Properties

`AuthType`

The supported authentication type to authenticate and connect
to your SharePoint site/sites.

_Required_: Yes

_Type_: String

_Allowed values_: `OAUTH2_CLIENT_CREDENTIALS | OAUTH2_SHAREPOINT_APP_ONLY_CLIENT_CREDENTIALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CredentialsSecretArn`

The Amazon Resource Name of an AWS Secrets Manager secret that
stores your authentication credentials for your SharePoint site/sites.
For more information on the key-value pairs that must be included in
your secret, depending on your authentication type, see
[SharePoint connection configuration](../../../bedrock/latest/userguide/sharepoint-data-source-connector.md#configuration-sharepoint-connector).

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov):secretsmanager:[a-z0-9-]{1,20}:([0-9]{12}|):secret:[a-zA-Z0-9!/_+=.@-]{1,512}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domain`

The domain of your SharePoint instance or site URL/URLs.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostType`

The supported host type, whether online/cloud or server/on-premises.

_Required_: Yes

_Type_: String

_Allowed values_: `ONLINE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SiteUrls`

A list of one or more SharePoint site URLs.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TenantId`

The identifier of your Microsoft 365 tenant.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SharePointDataSourceConfiguration

Transformation

All content copied from https://docs.aws.amazon.com/.
