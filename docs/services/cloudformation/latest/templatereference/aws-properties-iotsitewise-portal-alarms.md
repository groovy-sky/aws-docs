This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Portal Alarms

Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal.
You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range.
For more information, see [Monitoring with alarms](../../../iot-sitewise/latest/appguide/monitor-alarms.md) in the _AWS IoT SiteWise Application Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmRoleArn" : String,
  "NotificationLambdaArn" : String
}

```

### YAML

```yaml

  AlarmRoleArn: String
  NotificationLambdaArn: String

```

## Properties

`AlarmRoleArn`

The [ARN](../../../../general/latest/gr/aws-arns-and-namespaces.md) of the IAM role that allows the alarm to perform actions and access AWS
resources and services, such as AWS IoT Events.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationLambdaArn`

The [ARN](../../../../general/latest/gr/aws-arns-and-namespaces.md) of the Lambda function that manages alarm notifications. For more
information, see [Managing alarm\
notifications](../../../iotevents/latest/developerguide/lambda-support.md) in the _AWS IoT Events Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTSiteWise::Portal

PortalTypeEntry

All content copied from https://docs.aws.amazon.com/.
