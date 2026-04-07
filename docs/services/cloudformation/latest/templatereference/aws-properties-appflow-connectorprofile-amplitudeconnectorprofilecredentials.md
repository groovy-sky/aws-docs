This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile AmplitudeConnectorProfileCredentials

The connector-specific credentials required when using Amplitude.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiKey" : String,
  "SecretKey" : String
}

```

### YAML

```yaml

  ApiKey: String
  SecretKey: String

```

## Properties

`ApiKey`

A unique alphanumeric identifier used to authenticate a user, developer, or calling
program to your API.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretKey`

The Secret Access Key portion of the credentials.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [AmplitudeConnectorProfileCredentials](../../../../reference/appflow/1-0/apireference/api-amplitudeconnectorprofilecredentials.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppFlow::ConnectorProfile

ApiKeyCredentials
