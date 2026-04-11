This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow CustomConnectorSourceProperties

The properties that are applied when the custom connector is being used as a
source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomProperties" : {Key: Value, ...},
  "DataTransferApi" : DataTransferApi,
  "EntityName" : String
}

```

### YAML

```yaml

  CustomProperties:
    Key: Value
  DataTransferApi:
    DataTransferApi
  EntityName: String

```

## Properties

`CustomProperties`

Custom properties that are required to use the custom connector as a source.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\w]{1,2048}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataTransferApi`

The API of the connector application that Amazon AppFlow uses to transfer your
data.

_Required_: No

_Type_: [DataTransferApi](aws-properties-appflow-flow-datatransferapi.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntityName`

The entity specified in the custom connector as a source in the flow.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomConnectorDestinationProperties

DatadogSourceProperties

All content copied from https://docs.aws.amazon.com/.
