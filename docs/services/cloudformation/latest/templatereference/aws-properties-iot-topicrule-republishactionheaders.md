This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule RepublishActionHeaders

Specifies MQTT Version 5.0 headers information. For more information, see [MQTT](../../../iot/latest/developerguide/mqtt.md) in the IoT
Core Developer Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentType" : String,
  "CorrelationData" : String,
  "MessageExpiry" : String,
  "PayloadFormatIndicator" : String,
  "ResponseTopic" : String,
  "UserProperties" : [ UserProperty, ... ]
}

```

### YAML

```yaml

  ContentType: String
  CorrelationData: String
  MessageExpiry: String
  PayloadFormatIndicator: String
  ResponseTopic: String
  UserProperties:
    - UserProperty

```

## Properties

`ContentType`

A UTF-8 encoded string that describes the content of the publishing message.

For more information, see [Content Type](https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html) in the MQTT Version 5.0 specification.

Supports [substitution\
templates](../../../iot/latest/developerguide/iot-substitution-templates.md).

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CorrelationData`

The base64-encoded binary data used by the sender of the request message to identify
which request the response message is for.

For more information, see [Correlation Data](https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html) in the MQTT Version 5.0 specification.

Supports [substitution\
templates](../../../iot/latest/developerguide/iot-substitution-templates.md).

###### Note

This binary data must be base64-encoded.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageExpiry`

A user-defined integer value that represents the message expiry interval at the broker.
If the messages haven't been sent to the subscribers within that interval, the message
expires and is removed. The value of `messageExpiry` represents the number of
seconds before it expires. For more information about the limits of
`messageExpiry`, see [Message broker and protocol limits and\
quotas](../../../../general/latest/gr/iot-core.md#limits_iot) in the IoT Core Reference Guide.

Supports [substitution\
templates](../../../iot/latest/developerguide/iot-substitution-templates.md).

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PayloadFormatIndicator`

An `Enum` string value that indicates whether the payload is formatted as
UTF-8.

Valid values are `UNSPECIFIED_BYTES` and `UTF8_DATA`.

For more information, see [Payload Format Indicator](https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html) from the MQTT Version 5.0 specification.

Supports [substitution\
templates](../../../iot/latest/developerguide/iot-substitution-templates.md).

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseTopic`

A UTF-8 encoded string that's used as the topic name for a response message. The
response topic is used to describe the topic to which the receiver should publish as part
of the request-response flow. The topic must not contain wildcard characters.

For more information, see [Response Topic](https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html) in the MQTT Version 5.0 specification.

Supports [substitution\
templates](../../../iot/latest/developerguide/iot-substitution-templates.md).

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserProperties`

An array of key-value pairs that you define in the MQTT5 header.

_Required_: No

_Type_: Array of [UserProperty](aws-properties-iot-topicrule-userproperty.md)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RepublishAction

S3Action

All content copied from https://docs.aws.amazon.com/.
