This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile InforNexusConnectorProfileCredentials

The connector-specific profile credentials required by Infor Nexus.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessKeyId" : String,
  "Datakey" : String,
  "SecretAccessKey" : String,
  "UserId" : String
}

```

### YAML

```yaml

  AccessKeyId: String
  Datakey: String
  SecretAccessKey: String
  UserId: String

```

## Properties

`AccessKeyId`

The Access Key portion of the credentials.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Datakey`

The encryption keys used to encrypt data.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretAccessKey`

The secret key used to sign requests.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserId`

The identifier for the user.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [InforNexusConnectorProfileCredentials](../../../../reference/appflow/1-0/apireference/api-infornexusconnectorprofilecredentials.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GoogleAnalyticsConnectorProfileCredentials

InforNexusConnectorProfileProperties
