This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Pipeline Activity

An activity that performs a transformation on a message.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddAttributes" : AddAttributes,
  "Channel" : Channel,
  "Datastore" : Datastore,
  "DeviceRegistryEnrich" : DeviceRegistryEnrich,
  "DeviceShadowEnrich" : DeviceShadowEnrich,
  "Filter" : Filter,
  "Lambda" : Lambda,
  "Math" : Math,
  "RemoveAttributes" : RemoveAttributes,
  "SelectAttributes" : SelectAttributes
}

```

### YAML

```yaml

  AddAttributes:
    AddAttributes
  Channel:
    Channel
  Datastore:
    Datastore
  DeviceRegistryEnrich:
    DeviceRegistryEnrich
  DeviceShadowEnrich:
    DeviceShadowEnrich
  Filter:
    Filter
  Lambda:
    Lambda
  Math:
    Math
  RemoveAttributes:
    RemoveAttributes
  SelectAttributes:
    SelectAttributes

```

## Properties

`AddAttributes`

Adds other attributes based on existing attributes in the message.

_Required_: No

_Type_: [AddAttributes](aws-properties-iotanalytics-pipeline-addattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Channel`

Determines the source of the messages to be processed.

_Required_: No

_Type_: [Channel](aws-properties-iotanalytics-pipeline-channel.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Datastore`

Specifies where to store the processed message data.

_Required_: No

_Type_: [Datastore](aws-properties-iotanalytics-pipeline-datastore.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceRegistryEnrich`

Adds data from the AWS IoT device registry to your message.

_Required_: No

_Type_: [DeviceRegistryEnrich](aws-properties-iotanalytics-pipeline-deviceregistryenrich.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceShadowEnrich`

Adds information from the AWS IoT Device Shadows service to a message.

_Required_: No

_Type_: [DeviceShadowEnrich](aws-properties-iotanalytics-pipeline-deviceshadowenrich.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

Filters a message based on its attributes.

_Required_: No

_Type_: [Filter](aws-properties-iotanalytics-pipeline-filter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lambda`

Runs a Lambda function to modify the message.

_Required_: No

_Type_: [Lambda](aws-properties-iotanalytics-pipeline-lambda.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Math`

Computes an arithmetic expression using the message's attributes and adds
it to the message.

_Required_: No

_Type_: [Math](aws-properties-iotanalytics-pipeline-math.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveAttributes`

Removes attributes from a message.

_Required_: No

_Type_: [RemoveAttributes](aws-properties-iotanalytics-pipeline-removeattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectAttributes`

Creates a new message using only the specified attributes from the original message.

_Required_: No

_Type_: [SelectAttributes](aws-properties-iotanalytics-pipeline-selectattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTAnalytics::Pipeline

AddAttributes

All content copied from https://docs.aws.amazon.com/.
