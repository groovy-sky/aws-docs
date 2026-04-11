This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::DecoderManifest

Creates the decoder manifest associated with a model manifest. To create a decoder
manifest, the following must be true:

- Every signal decoder has a unique name.

- Each signal decoder is associated with a network interface.

- Each network interface has a unique ID.

- The signal decoders are specified in the model manifest.

For more information, see [Decoder manifests](../../../iot-fleetwise/latest/developerguide/decoder-manifests.md) in the _AWS IoT FleetWise Developer Guide_.

###### Important

Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](../../../iot-fleetwise/latest/developerguide/fleetwise-regions.md) in the _AWS IoT FleetWise Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTFleetWise::DecoderManifest",
  "Properties" : {
      "DefaultForUnmappedSignals" : String,
      "Description" : String,
      "ModelManifestArn" : String,
      "Name" : String,
      "NetworkInterfaces" : [ NetworkInterfacesItems, ... ],
      "SignalDecoders" : [ SignalDecodersItems, ... ],
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTFleetWise::DecoderManifest
Properties:
  DefaultForUnmappedSignals: String
  Description: String
  ModelManifestArn: String
  Name: String
  NetworkInterfaces:
    - NetworkInterfacesItems
  SignalDecoders:
    - SignalDecodersItems
  Status: String
  Tags:
    - Tag

```

## Properties

`DefaultForUnmappedSignals`

Use default decoders for all unmapped signals in the model. You don't need to provide any detailed decoding information.

_Required_: No

_Type_: String

_Allowed values_: `CUSTOM_DECODING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A brief description of the decoder manifest.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000-\u001F\u007F]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelManifestArn`

The Amazon Resource Name (ARN) of a vehicle model (model manifest) associated with the decoder
manifest.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the decoder manifest.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z\d\-_:]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkInterfaces`

A list of information about available network interfaces.

_Required_: No

_Type_: Array of [NetworkInterfacesItems](aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.md)

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SignalDecoders`

A list of information about signal decoders.

_Required_: No

_Type_: Array of [SignalDecodersItems](aws-properties-iotfleetwise-decodermanifest-signaldecodersitems.md)

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The state of the decoder manifest. If the status is `ACTIVE`, the decoder
manifest can't be edited. If the status is marked `DRAFT`, you can edit the
decoder manifest.

_Required_: No

_Type_: String

_Allowed values_: `ACTIVE | DRAFT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that can be used to manage the decoder manifest.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotfleetwise-decodermanifest-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the decoder manifest.

`CreationTime`

The time the decoder manifest was created in seconds since epoch (January 1, 1970 at midnight UTC time).

`LastModificationTime`

The time the decoder manifest was last updated in seconds since epoch (January 1, 1970 at midnight UTC time).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimestreamConfig

CanInterface

All content copied from https://docs.aws.amazon.com/.
