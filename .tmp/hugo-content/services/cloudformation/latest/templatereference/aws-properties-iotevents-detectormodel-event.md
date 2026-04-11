This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel Event

Specifies the `actions` to be performed when the `condition`
evaluates to TRUE.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ Action, ... ],
  "Condition" : String,
  "EventName" : String
}

```

### YAML

```yaml

  Actions:
    - Action
  Condition: String
  EventName: String

```

## Properties

`Actions`

The actions to be performed.

_Required_: No

_Type_: Array of [Action](aws-properties-iotevents-detectormodel-action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Condition`

Optional. The Boolean expression that, when TRUE, causes the `actions` to be
performed. If not present, the actions are performed (=TRUE). If the expression result is not
a Boolean value, the actions are not performed (=FALSE).

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventName`

The name of the event.

_Required_: Yes

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynamoDBv2

Firehose

All content copied from https://docs.aws.amazon.com/.
