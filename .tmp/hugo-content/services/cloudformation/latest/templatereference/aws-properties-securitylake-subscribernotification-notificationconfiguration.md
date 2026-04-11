This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::SubscriberNotification NotificationConfiguration

Specify the configurations you want to use for subscriber notification. The subscriber is notified when new data is written to the data lake for sources that the subscriber
consumes in Security Lake.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HttpsNotificationConfiguration" : HttpsNotificationConfiguration,
  "SqsNotificationConfiguration" : Json
}

```

### YAML

```yaml

  HttpsNotificationConfiguration:
    HttpsNotificationConfiguration
  SqsNotificationConfiguration: Json

```

## Properties

`HttpsNotificationConfiguration`

The configurations used for HTTPS subscriber notification.

_Required_: No

_Type_: [HttpsNotificationConfiguration](aws-properties-securitylake-subscribernotification-httpsnotificationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqsNotificationConfiguration`

The configurations for SQS subscriber notification. The members of this structure are context-dependent.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpsNotificationConfiguration

Next

All content copied from https://docs.aws.amazon.com/.
