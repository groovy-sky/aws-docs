This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile DatadogConnectorProfileCredentials

The connector-specific credentials required by Datadog.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiKey" : String,
  "ApplicationKey" : String
}

```

### YAML

```yaml

  ApiKey: String
  ApplicationKey: String

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

`ApplicationKey`

Application keys, in conjunction with your API key, give you full access to Datadog’s
programmatic API. Application keys are associated with the user account that created them. The
application key is used to log all requests made to the API.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [DatadogConnectorProfileCredentials](../../../../reference/appflow/1-0/apireference/api-datadogconnectorprofilecredentials.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomConnectorProfileProperties

DatadogConnectorProfileProperties
