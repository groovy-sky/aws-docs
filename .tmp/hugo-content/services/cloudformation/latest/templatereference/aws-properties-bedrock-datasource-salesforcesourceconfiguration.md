This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource SalesforceSourceConfiguration

The endpoint information to connect to your Salesforce data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthType" : String,
  "CredentialsSecretArn" : String,
  "HostUrl" : String
}

```

### YAML

```yaml

  AuthType: String
  CredentialsSecretArn: String
  HostUrl: String

```

## Properties

`AuthType`

The supported authentication type to authenticate and connect to your
Salesforce instance.

_Required_: Yes

_Type_: String

_Allowed values_: `OAUTH2_CLIENT_CREDENTIALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CredentialsSecretArn`

The Amazon Resource Name of an AWS Secrets Manager secret that
stores your authentication credentials for your Salesforce instance URL.
For more information on the key-value pairs that must be included in
your secret, depending on your authentication type, see
[Salesforce connection configuration](../../../bedrock/latest/userguide/salesforce-data-source-connector.md#configuration-salesforce-connector).

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov):secretsmanager:[a-z0-9-]{1,20}:([0-9]{12}|):secret:[a-zA-Z0-9!/_+=.@-]{1,512}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostUrl`

The Salesforce host URL or instance URL.

_Required_: Yes

_Type_: String

_Pattern_: `^https://[A-Za-z0-9][^\s]*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SalesforceDataSourceConfiguration

SeedUrl

All content copied from https://docs.aws.amazon.com/.
