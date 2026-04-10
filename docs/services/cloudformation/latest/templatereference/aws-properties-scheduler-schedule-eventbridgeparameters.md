This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Scheduler::Schedule EventBridgeParameters

The templated target type for the EventBridge [`PutEvents`](../../../../reference/eventbridge/latest/apireference/api-putevents.md) API operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DetailType" : String,
  "Source" : String
}

```

### YAML

```yaml

  DetailType: String
  Source: String

```

## Properties

`DetailType`

A free-form string, with a maximum of 128 characters, used to decide what fields to expect in the event detail.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source of the event.

_Required_: Yes

_Type_: String

_Pattern_: `^(?=[/\.\-_A-Za-z0-9]+)((?!aws\.).*)|(\$(\.[\w_-]+(\[(\d+|\*)\])*)*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EcsParameters

FlexibleTimeWindow

All content copied from https://docs.aws.amazon.com/.
