This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeTargetEventBridgeEventBusParameters

The parameters for using an EventBridge event bus as a target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DetailType" : String,
  "EndpointId" : String,
  "Resources" : [ String, ... ],
  "Source" : String,
  "Time" : String
}

```

### YAML

```yaml

  DetailType: String
  EndpointId: String
  Resources:
    - String
  Source: String
  Time: String

```

## Properties

`DetailType`

A free-form string, with a maximum of 128 characters, used to decide what fields to
expect in the event detail.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointId`

The URL subdomain of the endpoint. For example, if the URL for Endpoint is
https://abcde.veo.endpoints.event.amazonaws.com, then the EndpointId is
`abcde.veo`.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9\-]+[\.][A-Za-z0-9\-]+$`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resources`

AWS resources, identified by Amazon Resource Name (ARN), which the event
primarily concerns. Any number, including zero, may be present.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `1600 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source of the event.

_Required_: No

_Type_: String

_Pattern_: `(?=[/\.\-_A-Za-z0-9]+)((?!aws\.).*)|(\$(\.[\w/_-]+(\[(\d+|\*)\])*)*)`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Time`

The time stamp of the event, per [RFC3339](https://www.rfc-editor.org/rfc/rfc3339.txt). If no time stamp is provided, the time stamp of the [PutEvents](../../../../reference/eventbridge/latest/apireference/api-putevents.md) call is used.

_Required_: No

_Type_: String

_Pattern_: `^\$(\.[\w/_-]+(\[(\d+|\*)\])*)*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeTargetEcsTaskParameters

PipeTargetHttpParameters

All content copied from https://docs.aws.amazon.com/.
