This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::DeviceProfile LoRaWANDeviceProfile

LoRaWAN device profile object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClassBTimeout" : Integer,
  "ClassCTimeout" : Integer,
  "FactoryPresetFreqsList" : [ Integer, ... ],
  "MacVersion" : String,
  "MaxDutyCycle" : Integer,
  "MaxEirp" : Integer,
  "PingSlotDr" : Integer,
  "PingSlotFreq" : Integer,
  "PingSlotPeriod" : Integer,
  "RegParamsRevision" : String,
  "RfRegion" : String,
  "RxDataRate2" : Integer,
  "RxDelay1" : Integer,
  "RxDrOffset1" : Integer,
  "RxFreq2" : Integer,
  "Supports32BitFCnt" : Boolean,
  "SupportsClassB" : Boolean,
  "SupportsClassC" : Boolean,
  "SupportsJoin" : Boolean
}

```

### YAML

```yaml

  ClassBTimeout: Integer
  ClassCTimeout: Integer
  FactoryPresetFreqsList:
    - Integer
  MacVersion: String
  MaxDutyCycle: Integer
  MaxEirp: Integer
  PingSlotDr: Integer
  PingSlotFreq: Integer
  PingSlotPeriod: Integer
  RegParamsRevision: String
  RfRegion: String
  RxDataRate2: Integer
  RxDelay1: Integer
  RxDrOffset1: Integer
  RxFreq2: Integer
  Supports32BitFCnt: Boolean
  SupportsClassB: Boolean
  SupportsClassC: Boolean
  SupportsJoin: Boolean

```

## Properties

`ClassBTimeout`

The ClassBTimeout value.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClassCTimeout`

The ClassCTimeout value.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FactoryPresetFreqsList`

The list of values that make up the FactoryPresetFreqs value. Valid range of values
include a minimum value of 1000000 and a maximum value of 16700000.

_Required_: No

_Type_: Array of Integer

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MacVersion`

The MAC version (such as OTAA 1.1 or OTAA 1.0.3) to use with this device profile.

_Required_: No

_Type_: String

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxDutyCycle`

The MaxDutyCycle value.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxEirp`

The MaxEIRP value.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PingSlotDr`

The PingSlotDR value.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PingSlotFreq`

The PingSlotFreq value.

_Required_: No

_Type_: Integer

_Minimum_: `1000000`

_Maximum_: `16700000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PingSlotPeriod`

The PingSlotPeriod value.

_Required_: No

_Type_: Integer

_Minimum_: `128`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RegParamsRevision`

The version of regional parameters.

_Required_: No

_Type_: String

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RfRegion`

The frequency band (RFRegion) value.

_Required_: No

_Type_: String

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RxDataRate2`

The RXDataRate2 value.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RxDelay1`

The RXDelay1 value.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RxDrOffset1`

The RXDROffset1 value.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `7`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RxFreq2`

The RXFreq2 value.

_Required_: No

_Type_: Integer

_Minimum_: `1000000`

_Maximum_: `16700000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Supports32BitFCnt`

The Supports32BitFCnt value.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SupportsClassB`

The SupportsClassB value.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SupportsClassC`

The SupportsClassC value.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SupportsJoin`

The SupportsJoin value.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTWireless::DeviceProfile

Tag

All content copied from https://docs.aws.amazon.com/.
