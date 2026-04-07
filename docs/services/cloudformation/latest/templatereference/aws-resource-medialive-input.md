This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Input

The AWS::MediaLive::Input resource is a MediaLive resource type that creates an input.

A MediaLive input holds information that describes how the MediaLive channel is connected to the upstream system
that is providing the source content that is to be transcoded.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaLive::Input",
  "Properties" : {
      "Destinations" : [ InputDestinationRequest, ... ],
      "InputDevices" : [ InputDeviceSettings, ... ],
      "InputNetworkLocation" : String,
      "InputSecurityGroups" : [ String, ... ],
      "MediaConnectFlows" : [ MediaConnectFlowRequest, ... ],
      "MulticastSettings" : MulticastSettingsCreateRequest,
      "Name" : String,
      "RoleArn" : String,
      "RouterSettings" : RouterSettings,
      "SdiSources" : [ String, ... ],
      "Smpte2110ReceiverGroupSettings" : Smpte2110ReceiverGroupSettings,
      "Sources" : [ InputSourceRequest, ... ],
      "SrtSettings" : SrtSettingsRequest,
      "Tags" : [ Tag, ... ],
      "Type" : String,
      "Vpc" : InputVpcRequest
    }
}

```

### YAML

```yaml

Type: AWS::MediaLive::Input
Properties:
  Destinations:
    - InputDestinationRequest
  InputDevices:
    - InputDeviceSettings
  InputNetworkLocation: String
  InputSecurityGroups:
    - String
  MediaConnectFlows:
    - MediaConnectFlowRequest
  MulticastSettings:
    MulticastSettingsCreateRequest
  Name: String
  RoleArn: String
  RouterSettings:
    RouterSettings
  SdiSources:
    - String
  Smpte2110ReceiverGroupSettings:
    Smpte2110ReceiverGroupSettings
  Sources:
    - InputSourceRequest
  SrtSettings:
    SrtSettingsRequest
  Tags:
    - Tag
  Type: String
  Vpc:
    InputVpcRequest

```

## Properties

`Destinations`

Settings that apply only if the input is a push type of input.

_Required_: No

_Type_: Array of [InputDestinationRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-input-inputdestinationrequest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputDevices`

Settings that apply only if the input is an Elemental Link input.

_Required_: No

_Type_: Array of [InputDeviceSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-input-inputdevicesettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputNetworkLocation`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InputSecurityGroups`

The list of input security groups (referenced by IDs) to attach to the input if the input is a push type.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaConnectFlows`

Settings that apply only if the input is a MediaConnect input.

_Required_: No

_Type_: Array of [MediaConnectFlowRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-input-mediaconnectflowrequest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MulticastSettings`

Property description not available.

_Required_: No

_Type_: [MulticastSettingsCreateRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-input-multicastsettingscreaterequest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the input.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM role for MediaLive to assume when creating a MediaConnect input or Amazon VPC input. This doesn't apply
to other types of inputs. The role is identified by its ARN.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RouterSettings`

Property description not available.

_Required_: No

_Type_: [RouterSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-input-routersettings.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SdiSources`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Smpte2110ReceiverGroupSettings`

Property description not available.

_Required_: No

_Type_: [Smpte2110ReceiverGroupSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-input-smpte2110receivergroupsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sources`

Settings that apply only if the input is a pull type of input.

_Required_: No

_Type_: Array of [InputSourceRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-input-inputsourcerequest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SrtSettings`

Property description not available.

_Required_: No

_Type_: [SrtSettingsRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-input-srtsettingsrequest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A collection of tags for this input. Each tag is a key-value pair.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type for this input.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Vpc`

Settings that apply only if the input is an push input where the source is on Amazon VPC.

_Required_: No

_Type_: [InputVpcRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-input-inputvpcrequest.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the MediaLive id of the input.

For example: `{ "Ref": "myInput" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the MediaLive input. For example: arn:aws:medialive:us-west-1:111122223333:medialive:input:1234567.
MediaLive creates this ARN when it creates the input.

`Destinations`

For a push input, the the destination or destinations for the input. The destinations are the URLs of locations
on MediaLive where the upstream system pushes the content to, for this input. MediaLive creates these addresses when
it creates the input.

`Id`

The unique ID for the device.

`Sources`

For a pull input, the source or sources for the input. The sources are the URLs of locations on the upstream
system where MediaLive pulls the content from, for this input. You included these URLs in the create request.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::MediaLive::EventBridgeRuleTemplateGroup

InputDestinationRequest
