This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::StateTemplate

Creates a mechanism for vehicle owners to track the state of their vehicles. State templates determine which signal updates the vehicle sends to the cloud.

For more information, see [State templates](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/state-templates.html) in the _AWS IoT FleetWise Developer Guide_.

###### Important

Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the _AWS IoT FleetWise Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTFleetWise::StateTemplate",
  "Properties" : {
      "DataExtraDimensions" : [ String, ... ],
      "Description" : String,
      "MetadataExtraDimensions" : [ String, ... ],
      "Name" : String,
      "SignalCatalogArn" : String,
      "StateTemplateProperties" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTFleetWise::StateTemplate
Properties:
  DataExtraDimensions:
    - String
  Description: String
  MetadataExtraDimensions:
    - String
  Name: String
  SignalCatalogArn: String
  StateTemplateProperties:
    - String
  Tags:
    - Tag

```

## Properties

`DataExtraDimensions`

A list of vehicle attributes associated with the payload published on the state template's MQTT topic.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `150 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A brief description of the state template.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000-\u001F\u007F]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetadataExtraDimensions`

A list of vehicle attributes to associate with the user properties of the messages published on the
state template's MQTT topic. For example, if you add
`Vehicle.Attributes.Make` and `Vehicle.Attributes.Model` attributes, these attributes are included as user properties with the MQTT message.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `150 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The unique alias of the state template.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z\d\-_:]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SignalCatalogArn`

The Amazon Resource Name (ARN) of the signal catalog associated with the state template.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StateTemplateProperties`

A list of signals from which data is collected. The state template properties contain the fully qualified names of the signals.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `150 | 500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that can be used to manage the state template.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotfleetwise-statetemplate-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the state template.

`CreationTime`

The time the state template was created, in seconds since epoch (January 1, 1970 at midnight UTC time).

`Id`

The unique ID of the state template.

`LastModificationTime`

The time the state template was last updated, in seconds since epoch (January 1, 1970 at midnight UTC time).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
