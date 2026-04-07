This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign CustomDeliveryConfiguration

Specifies the delivery configuration settings for sending a campaign or
campaign treatment through a custom channel. This object is required if you use
the `CampaignCustomMessage` object to define the message to send for
the campaign or campaign treatment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeliveryUri" : String,
  "EndpointTypes" : [ String, ... ]
}

```

### YAML

```yaml

  DeliveryUri: String
  EndpointTypes:
    - String

```

## Properties

`DeliveryUri`

The destination to send the campaign or treatment to. This value can be one of
the following:

- The name or Amazon Resource Name (ARN) of an AWS Lambda
function to invoke to handle delivery of the campaign or
treatment.

- The URL for a web application or service that supports HTTPS and can
receive the message. The URL has to be a full URL, including the HTTPS
protocol.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointTypes`

The types of endpoints to send the campaign or treatment to. Each valid value
maps to a type of channel that you can associate with an endpoint by using the
`ChannelType` property of an endpoint.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CampaignSmsMessage

DefaultButtonConfiguration
