This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule IotSiteWiseAction

Describes an action to send data from an MQTT message that triggered the rule to AWS IoT
SiteWise asset properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PutAssetPropertyValueEntries" : [ PutAssetPropertyValueEntry, ... ],
  "RoleArn" : String
}

```

### YAML

```yaml

  PutAssetPropertyValueEntries:
    - PutAssetPropertyValueEntry
  RoleArn: String

```

## Properties

`PutAssetPropertyValueEntries`

A list of asset property value entries.

_Required_: Yes

_Type_: Array of [PutAssetPropertyValueEntry](aws-properties-iot-topicrule-putassetpropertyvalueentry.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the role that grants AWS IoT permission to send an asset property value to AWS IoT SiteWise. ( `"Action": "iotsitewise:BatchPutAssetPropertyValue"`). The trust
policy can restrict access to specific asset hierarchy paths.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IotEventsAction

KafkaAction

All content copied from https://docs.aws.amazon.com/.
