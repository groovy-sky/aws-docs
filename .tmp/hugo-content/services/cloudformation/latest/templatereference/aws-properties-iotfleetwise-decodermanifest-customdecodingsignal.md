This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::DecoderManifest CustomDecodingSignal

Information about signals using a custom decoding protocol as defined by the customer.

###### Important

AWS IoT FleetWise will no longer be open to new customers starting April 30, 2026.
If you would like to use AWS IoT FleetWise, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see [AWS IoT FleetWise availability change](../../../iot-fleetwise/latest/developerguide/iotfleetwise-availability-change.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String
}

```

### YAML

```yaml

  Id: String

```

## Properties

`Id`

The ID of the signal.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!.*\.\.)[a-zA-Z0-9_\-#:.]+$`

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomDecodingNetworkInterface

CustomDecodingSignalDecoder

All content copied from https://docs.aws.amazon.com/.
