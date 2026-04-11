This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::DecoderManifest SignalDecodersItems

Information about a signal decoder.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CanSignalDecoder" : CanSignalDecoder,
  "CustomDecodingSignalDecoder" : CustomDecodingSignalDecoder,
  "ObdSignalDecoder" : ObdSignalDecoder
}

```

### YAML

```yaml

  CanSignalDecoder:
    CanSignalDecoder
  CustomDecodingSignalDecoder:
    CustomDecodingSignalDecoder
  ObdSignalDecoder:
    ObdSignalDecoder

```

## Properties

`CanSignalDecoder`

Information about signal decoder using the Controller Area Network (CAN)
protocol.

_Required_: No

_Type_: [CanSignalDecoder](aws-properties-iotfleetwise-decodermanifest-cansignaldecoder.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomDecodingSignalDecoder`

Information about a [custom signal \
decoder](../../../../reference/iot-fleetwise/latest/apireference/api-customdecodingsignal.md).

_Required_: No

_Type_: [CustomDecodingSignalDecoder](aws-properties-iotfleetwise-decodermanifest-customdecodingsignaldecoder.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObdSignalDecoder`

Information about signal decoder using the on-board diagnostic (OBD) II
protocol.

_Required_: No

_Type_: [ObdSignalDecoder](aws-properties-iotfleetwise-decodermanifest-obdsignaldecoder.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObdSignalDecoder

Tag

All content copied from https://docs.aws.amazon.com/.
