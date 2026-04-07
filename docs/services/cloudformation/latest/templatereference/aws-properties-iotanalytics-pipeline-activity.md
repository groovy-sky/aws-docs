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

_Type_: [AddAttributes](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-pipeline-addattributes.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Channel`

Determines the source of the messages to be processed.

_Required_: No

_Type_: [Channel](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-pipeline-channel.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Datastore`

Specifies where to store the processed message data.

_Required_: No

_Type_: [Datastore](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-pipeline-datastore.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceRegistryEnrich`

Adds data from the AWS IoT device registry to your message.

_Required_: No

_Type_: [DeviceRegistryEnrich](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceShadowEnrich`

Adds information from the AWS IoT Device Shadows service to a message.

_Required_: No

_Type_: [DeviceShadowEnrich](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-pipeline-deviceshadowenrich.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

Filters a message based on its attributes.

_Required_: No

_Type_: [Filter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-pipeline-filter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lambda`

Runs a Lambda function to modify the message.

_Required_: No

_Type_: [Lambda](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-pipeline-lambda.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Math`

Computes an arithmetic expression using the message's attributes and adds
it to the message.

_Required_: No

_Type_: [Math](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-pipeline-math.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveAttributes`

Removes attributes from a message.

_Required_: No

_Type_: [RemoveAttributes](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-pipeline-removeattributes.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectAttributes`

Creates a new message using only the specified attributes from the original message.

_Required_: No

_Type_: [SelectAttributes](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-pipeline-selectattributes.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IoTAnalytics::Pipeline

AddAttributes
