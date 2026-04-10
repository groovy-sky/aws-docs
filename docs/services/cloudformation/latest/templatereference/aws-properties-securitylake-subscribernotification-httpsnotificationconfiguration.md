This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::SubscriberNotification HttpsNotificationConfiguration

Specify the configurations you want to use for HTTPS subscriber notification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationApiKeyName" : String,
  "AuthorizationApiKeyValue" : String,
  "Endpoint" : String,
  "HttpMethod" : String,
  "TargetRoleArn" : String
}

```

### YAML

```yaml

  AuthorizationApiKeyName: String
  AuthorizationApiKeyValue: String
  Endpoint: String
  HttpMethod: String
  TargetRoleArn: String

```

## Properties

`AuthorizationApiKeyName`

The key name for the notification subscription.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizationApiKeyValue`

The key value for the notification subscription.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Endpoint`

The subscription endpoint in Security Lake. If you prefer notification with an HTTPS
endpoint, populate this field.

_Required_: Yes

_Type_: String

_Pattern_: `^https?://.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpMethod`

The HTTPS method used for the notification subscription.

_Required_: No

_Type_: String

_Allowed values_: `POST | PUT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetRoleArn`

The Amazon Resource Name (ARN) of the EventBridge API destinations IAM role that you
created. For more information about ARNs and how to use them in policies, see [Managing data access](../../../security-lake/latest/userguide/subscriber-data-access.md) and [AWS\
Managed Policies](../../../security-lake/latest/userguide/security-iam-awsmanpol.md) in the _Amazon Security Lake User Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SecurityLake::SubscriberNotification

NotificationConfiguration

All content copied from https://docs.aws.amazon.com/.
