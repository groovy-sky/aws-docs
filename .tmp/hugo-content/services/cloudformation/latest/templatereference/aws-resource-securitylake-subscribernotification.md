This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::SubscriberNotification

Notifies the subscriber when new data is written to the data lake for the sources that
the subscriber consumes in Security Lake. You can create only one subscriber notification per
subscriber.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityLake::SubscriberNotification",
  "Properties" : {
      "NotificationConfiguration" : NotificationConfiguration,
      "SubscriberArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::SecurityLake::SubscriberNotification
Properties:
  NotificationConfiguration:
    NotificationConfiguration
  SubscriberArn: String

```

## Properties

`NotificationConfiguration`

Specify the configurations you want to use for subscriber notification. The subscriber is notified when new data is written to the data lake for sources that the subscriber
consumes in Security Lake.

_Required_: Yes

_Type_: [NotificationConfiguration](aws-properties-securitylake-subscribernotification-notificationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubscriberArn`

The Amazon Resource Name (ARN) of the Security Lake subscriber.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `ref` function, `ref` returns the type of `SubscriberArn`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`SubscriberEndpoint`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

HttpsNotificationConfiguration

All content copied from https://docs.aws.amazon.com/.
