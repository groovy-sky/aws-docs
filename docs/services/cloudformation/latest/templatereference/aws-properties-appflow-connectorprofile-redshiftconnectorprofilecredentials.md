This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile RedshiftConnectorProfileCredentials

The connector-specific profile credentials required when using Amazon Redshift.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Password" : String,
  "Username" : String
}

```

### YAML

```yaml

  Password: String
  Username: String

```

## Properties

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

- [RedshiftConnectorProfileCredentials](../../../../reference/appflow/1-0/apireference/api-redshiftconnectorprofilecredentials.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PardotConnectorProfileProperties

RedshiftConnectorProfileProperties
