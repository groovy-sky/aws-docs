This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Endpoint EndpointEventBus

The event buses the endpoint is associated with.

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

The ARN of the event bus the endpoint is associated with.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z-]*:events:[a-z]+-[a-z-]+-\d+:\d{12}:event-bus/[\w.-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Events::Endpoint

FailoverConfig

All content copied from https://docs.aws.amazon.com/.
