This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile ZendeskConnectorProfileProperties

The connector-specific profile properties required when using Zendesk.

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

The location of the Zendesk resource.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ZendeskConnectorProfileProperties](../../../../reference/appflow/1-0/apireference/api-zendeskconnectorprofileproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ZendeskConnectorProfileCredentials

AWS::AppFlow::Flow

All content copied from https://docs.aws.amazon.com/.
