This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::DataSource EventBridgeConfig

The data source. This can be an API destination, resource, or AWS
service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventBusArn" : String
}

```

### YAML

```yaml

  EventBusArn: String

```

## Properties

`EventBusArn`

The event bus pipeline's ARN. For more information about event buses, see [EventBridge event buses](../../../eventbridge/latest/userguide/eb-event-bus.md).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynamoDBConfig

HttpConfig

All content copied from https://docs.aws.amazon.com/.
