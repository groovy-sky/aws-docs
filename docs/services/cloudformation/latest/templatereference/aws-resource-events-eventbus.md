This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::EventBus

Specifies an event bus within your account. This can be a custom event bus which you can
use to receive events from your custom applications and services, or it can be a partner event
bus which can be matched to a partner event source.

###### Note

As an aid to help you jumpstart developing CloudFormation templates, the EventBridge console
enables you to create templates from the existing event buses in your account. For more information, see
[Generating CloudFormation templates from an EventBridge event bus](../../../eventbridge/latest/userguide/eb-generate-event-bus-template.md)
in the _Amazon EventBridge User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Events::EventBus",
  "Properties" : {
      "DeadLetterConfig" : DeadLetterConfig,
      "Description" : String,
      "EventSourceName" : String,
      "KmsKeyIdentifier" : String,
      "LogConfig" : LogConfig,
      "Name" : String,
      "Policy" : Json,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Events::EventBus
Properties:
  DeadLetterConfig:
    DeadLetterConfig
  Description: String
  EventSourceName: String
  KmsKeyIdentifier: String
  LogConfig:
    LogConfig
  Name: String
  Policy: Json
  Tags:
    - Tag

```

## Properties

`DeadLetterConfig`

Configuration details of the Amazon SQS queue for EventBridge to use as a
dead-letter queue (DLQ).

For more information, see [Using dead-letter queues to process undelivered events](../../../eventbridge/latest/userguide/eb-rule-event-delivery.md#eb-rule-dlq) in the _EventBridge User_
_Guide_.

_Required_: No

_Type_: [DeadLetterConfig](aws-properties-events-eventbus-deadletterconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The event bus description.

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventSourceName`

If you are creating a partner event bus, this specifies the partner event source that the
new event bus will be matched with.

_Required_: No

_Type_: String

_Pattern_: `aws\.partner(/[\.\-_A-Za-z0-9]+){2,}`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyIdentifier`

The identifier of the AWS KMS
customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt events on this event bus. The identifier can be the key
Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.

If you do not specify a customer managed key identifier, EventBridge uses an
AWS owned key to encrypt events on the event bus.

For more information, see [Identify and view keys](../../../kms/latest/developerguide/viewing-keys.md) in the _AWS Key Management Service_
_Developer Guide_.

###### Note

Schema discovery is not supported for event buses encrypted using a
customer managed key. EventBridge returns an error if:

- You call `
                                          CreateDiscoverer
                                      ` on an event bus set to use a customer managed key for encryption.

- You call `
                                          UpdatedEventBus
                                      ` to set a customer managed key on an event bus with schema discovery enabled.

To enable schema discovery on an event bus, choose to
use an AWS owned key. For more information, see [Encrypting events](../../../eventbridge/latest/userguide/eb-encryption-event-bus-cmkey.md) in the _Amazon EventBridge User Guide_.

###### Important

If you have specified that EventBridge use a customer managed key for encrypting the source event bus, we strongly recommend you also specify a
customer managed key for any archives for the event bus as well.

For more information, see [Encrypting archives](../../../eventbridge/latest/userguide/encryption-archives.md) in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-/:]*$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogConfig`

The logging configuration settings for the event bus.

For more information, see [Configuring logs for event buses](../../../eventbridge/latest/userguide/eb-event-bus-logs.md) in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: [LogConfig](aws-properties-events-eventbus-logconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the new event bus.

Custom event bus names can't contain the `/` character, but you can use the
`/` character in partner event bus names. In addition, for partner event buses,
the name must exactly match the name of the partner event source that this event bus is
matched to.

You can't use the name `default` for a custom event bus, as this name is
already used for your account's default event bus.

_Required_: Yes

_Type_: String

_Pattern_: `[/\.\-_A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

The permissions policy of the event bus, describing which other AWS
accounts can write events to this event bus.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags to associate with the event bus.

_Required_: No

_Type_: Array of [Tag](aws-properties-events-eventbus-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

The name of the new event bus.

### Fn::GetAtt

The ARN of the task definition to use. If no task revision is supplied, it defaults to the
most recent revision at the time of resource creation.

`Arn`

The ARN of the event bus, such as
`arn:aws:events:us-east-2:123456789012:event-bus/aws.partner/PartnerName/acct1/repo1`.

`Name`

The name of the event bus, such as `PartnerName/acct1/repo1`.

## Examples

- [Create a partner event bus](#aws-resource-events-eventbus--examples--Create_a_partner_event_bus)

- [Create a custom event bus](#aws-resource-events-eventbus--examples--Create_a_custom_event_bus)

### Create a partner event bus

The following example creates a partner event bus named
`aws.partner.repo1`.

#### JSON

```json

"SamplePartnerEventBus": {
    "Type": "AWS::Events::EventBus",
    "Properties": {
        "EventSourceName": "aws.partner/PartnerName/acct1/repo1",
        "Name": "aws.partner.repo1"
    }
}
```

#### YAML

```yaml

SamplePartnerEventBus:
    Type: AWS::Events::EventBus
    Properties:
        EventSourceName: "aws.partner/PartnerName/acct1/repo1"
        Name: "aws.partner.repo1"
```

### Create a custom event bus

The following example creates a custom event bus named
`MyCustomEventBus`.

#### JSON

```json

"SampleCustomEventBus": {
    "Type": "AWS::Events::EventBus",
    "Properties": {
        "Name": "MyCustomEventBus"
    }
}
```

#### YAML

```yaml

SampleCustomEventBus:
    Type: AWS::Events::EventBus
    Properties:
        Name: "MyCustomEventBus"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Secondary

DeadLetterConfig

All content copied from https://docs.aws.amazon.com/.
