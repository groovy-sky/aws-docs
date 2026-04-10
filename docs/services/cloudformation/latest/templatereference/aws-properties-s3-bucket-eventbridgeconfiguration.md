This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket EventBridgeConfiguration

Amazon S3 can send events to Amazon EventBridge whenever certain events happen in your
bucket, see [Using\
EventBridge](../../../s3/latest/userguide/eventbridge.md) in the _Amazon S3 User Guide_.

Unlike other destinations, delivery of events to EventBridge can be either enabled or
disabled for a bucket. If enabled, all events will be sent to EventBridge and you can use
EventBridge rules to route events to additional targets. For more information, see [What Is Amazon\
EventBridge](../../../eventbridge/latest/userguide/eb-what-is.md) in the _Amazon EventBridge User Guide_

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventBridgeEnabled" : Boolean
}

```

### YAML

```yaml

  EventBridgeEnabled: Boolean

```

## Properties

`EventBridgeEnabled`

Enables delivery of events to Amazon EventBridge.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Enable EventBridgeConfiguration

The following example template shows an Amazon S3 bucket with a notification
configuration with EventBridge enabled.

#### JSON

```json

{
"Resources": {
  "S3Bucket": {
    "Type": "AWS::S3::Bucket",
    "Properties": {
      "NotificationConfiguration": {
        "EventBridgeConfiguration": {
          "EventBridgeEnabled": true
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  S3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      NotificationConfiguration:
        EventBridgeConfiguration:
          EventBridgeEnabled: true
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfiguration

FilterRule

All content copied from https://docs.aws.amazon.com/.
