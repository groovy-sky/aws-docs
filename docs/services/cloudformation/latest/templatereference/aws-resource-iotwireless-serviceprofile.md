This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::ServiceProfile

Creates a new service profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTWireless::ServiceProfile",
  "Properties" : {
      "LoRaWAN" : LoRaWANServiceProfile,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTWireless::ServiceProfile
Properties:
  LoRaWAN:
    LoRaWANServiceProfile
  Name: String
  Tags:
    - Tag

```

## Properties

`LoRaWAN`

LoRaWAN service profile object.

_Required_: No

_Type_: [LoRaWANServiceProfile](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the new resource.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags are an array of key-value pairs to attach to the specified resource. Tags can
have a minimum of 0 and a maximum of 50 items.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotwireless-serviceprofile-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the service profile ID.

### Fn::GetAtt

`Arn`

The ARN of the service profile created.

`Id`

The ID of the service profile created.

`LoRaWAN.ChannelMask`

The ChannelMask value.

`LoRaWAN.DevStatusReqFreq`

The DevStatusReqFreq value.

`LoRaWAN.DlBucketSize`

The DLBucketSize value.

`LoRaWAN.DlRate`

The DLRate value.

`LoRaWAN.DlRatePolicy`

The DLRatePolicy value.

`LoRaWAN.DrMax`

The DRMax value.

`LoRaWAN.DrMin`

The DRMin value.

`LoRaWAN.HrAllowed`

The HRAllowed value that describes whether handover roaming is allowed.

`LoRaWAN.MinGwDiversity`

The MinGwDiversity value.

`LoRaWAN.NwkGeoLoc`

The NwkGeoLoc value.

`LoRaWAN.ReportDevStatusBattery`

The ReportDevStatusBattery value.

`LoRaWAN.ReportDevStatusMargin`

The ReportDevStatusMargin value.

`LoRaWAN.TargetPer`

The TargetPer value.

`LoRaWAN.UlBucketSize`

The UlBucketSize value.

`LoRaWAN.UlRate`

The ULRate value.

`LoRaWAN.UlRatePolicy`

The ULRatePolicy value.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

LoRaWANServiceProfile
