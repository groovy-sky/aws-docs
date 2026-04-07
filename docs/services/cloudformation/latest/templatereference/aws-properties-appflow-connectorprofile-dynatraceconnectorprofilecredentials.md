This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile DynatraceConnectorProfileCredentials

The connector-specific profile credentials required by Dynatrace.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiToken" : String
}

```

### YAML

```yaml

  ApiToken: String

```

## Properties

`ApiToken`

The API tokens used by Dynatrace API to authenticate various API calls.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [DynatraceConnectorProfileCredentials](../../../../reference/appflow/1-0/apireference/api-dynatraceconnectorprofilecredentials.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DatadogConnectorProfileProperties

DynatraceConnectorProfileProperties
