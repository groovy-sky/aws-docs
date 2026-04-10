This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile DatadogConnectorProfileProperties

The connector-specific profile properties required by Datadog.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InstanceUrl" : String
}

```

### YAML

```yaml

  InstanceUrl: String

```

## Properties

`InstanceUrl`

The location of the Datadog resource.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [DatadogConnectorProfileProperties](../../../../reference/appflow/1-0/apireference/api-datadogconnectorprofileproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatadogConnectorProfileCredentials

DynatraceConnectorProfileCredentials

All content copied from https://docs.aws.amazon.com/.
