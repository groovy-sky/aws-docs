This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::ServiceProfile LoRaWANServiceProfile

LoRaWANServiceProfile object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddGwMetadata" : Boolean,
  "ChannelMask" : String,
  "DevStatusReqFreq" : Integer,
  "DlBucketSize" : Integer,
  "DlRate" : Integer,
  "DlRatePolicy" : String,
  "DrMax" : Integer,
  "DrMin" : Integer,
  "HrAllowed" : Boolean,
  "MinGwDiversity" : Integer,
  "NwkGeoLoc" : Boolean,
  "PrAllowed" : Boolean,
  "RaAllowed" : Boolean,
  "ReportDevStatusBattery" : Boolean,
  "ReportDevStatusMargin" : Boolean,
  "TargetPer" : Integer,
  "UlBucketSize" : Integer,
  "UlRate" : Integer,
  "UlRatePolicy" : String
}

```

### YAML

```yaml

  AddGwMetadata: Boolean
  ChannelMask: String
  DevStatusReqFreq: Integer
  DlBucketSize: Integer
  DlRate: Integer
  DlRatePolicy: String
  DrMax: Integer
  DrMin: Integer
  HrAllowed: Boolean
  MinGwDiversity: Integer
  NwkGeoLoc: Boolean
  PrAllowed: Boolean
  RaAllowed: Boolean
  ReportDevStatusBattery: Boolean
  ReportDevStatusMargin: Boolean
  TargetPer: Integer
  UlBucketSize: Integer
  UlRate: Integer
  UlRatePolicy: String

```

## Properties

`AddGwMetadata`

The AddGWMetaData value.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ChannelMask`

The ChannelMask value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DevStatusReqFreq`

The DevStatusReqFreq value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DlBucketSize`

The DLBucketSize value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DlRate`

The DLRate value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DlRatePolicy`

The DLRatePolicy value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DrMax`

The DRMax value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DrMin`

The DRMin value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HrAllowed`

The HRAllowed value that describes whether handover roaming is allowed.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinGwDiversity`

The MinGwDiversity value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NwkGeoLoc`

The NwkGeoLoc value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrAllowed`

The PRAllowed value that describes whether passive roaming is allowed.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RaAllowed`

The RAAllowed value that describes whether roaming activation is allowed.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReportDevStatusBattery`

The ReportDevStatusBattery value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReportDevStatusMargin`

The ReportDevStatusMargin value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetPer`

The TargetPer value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UlBucketSize`

The UlBucketSize value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UlRate`

The ULRate value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UlRatePolicy`

The ULRatePolicy value.

This property is `ReadOnly` and can't be inputted for create. It's returned
with `Fn::GetAtt`

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTWireless::ServiceProfile

Tag

All content copied from https://docs.aws.amazon.com/.
