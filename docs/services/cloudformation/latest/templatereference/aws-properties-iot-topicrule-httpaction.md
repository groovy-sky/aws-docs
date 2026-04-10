This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule HttpAction

Send data to an HTTPS endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Auth" : HttpAuthorization,
  "BatchConfig" : BatchConfig,
  "ConfirmationUrl" : String,
  "EnableBatching" : Boolean,
  "Headers" : [ HttpActionHeader, ... ],
  "Url" : String
}

```

### YAML

```yaml

  Auth:
    HttpAuthorization
  BatchConfig:
    BatchConfig
  ConfirmationUrl: String
  EnableBatching: Boolean
  Headers:
    - HttpActionHeader
  Url: String

```

## Properties

`Auth`

The authentication method to use when sending data to an HTTPS endpoint.

_Required_: No

_Type_: [HttpAuthorization](aws-properties-iot-topicrule-httpauthorization.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BatchConfig`

Property description not available.

_Required_: No

_Type_: [BatchConfig](aws-properties-iot-topicrule-batchconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfirmationUrl`

The URL to which AWS IoT sends a confirmation message. The value of the confirmation URL
must be a prefix of the endpoint URL. If you do not specify a confirmation URL AWS IoT uses
the endpoint URL as the confirmation URL. If you use substitution templates in the
confirmationUrl, you must create and enable topic rule destinations that match each
possible value of the substitution template before traffic is allowed to your endpoint
URL.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableBatching`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Headers`

The HTTP headers to send with the message data.

_Required_: No

_Type_: Array of [HttpActionHeader](aws-properties-iot-topicrule-httpactionheader.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The endpoint URL. If substitution templates are used in the URL, you must also specify a
`confirmationUrl`. If this is a new destination, a new
`TopicRuleDestination` is created if possible.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FirehoseAction

HttpActionHeader

All content copied from https://docs.aws.amazon.com/.
