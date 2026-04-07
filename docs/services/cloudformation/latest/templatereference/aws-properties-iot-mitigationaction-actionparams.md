This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::MitigationAction ActionParams

Defines the type of action and the parameters for that action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddThingsToThingGroupParams" : AddThingsToThingGroupParams,
  "EnableIoTLoggingParams" : EnableIoTLoggingParams,
  "PublishFindingToSnsParams" : PublishFindingToSnsParams,
  "ReplaceDefaultPolicyVersionParams" : ReplaceDefaultPolicyVersionParams,
  "UpdateCACertificateParams" : UpdateCACertificateParams,
  "UpdateDeviceCertificateParams" : UpdateDeviceCertificateParams
}

```

### YAML

```yaml

  AddThingsToThingGroupParams:
    AddThingsToThingGroupParams
  EnableIoTLoggingParams:
    EnableIoTLoggingParams
  PublishFindingToSnsParams:
    PublishFindingToSnsParams
  ReplaceDefaultPolicyVersionParams:
    ReplaceDefaultPolicyVersionParams
  UpdateCACertificateParams:
    UpdateCACertificateParams
  UpdateDeviceCertificateParams:
    UpdateDeviceCertificateParams

```

## Properties

`AddThingsToThingGroupParams`

Specifies the group to which you want to add the devices.

_Required_: No

_Type_: [AddThingsToThingGroupParams](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-mitigationaction-addthingstothinggroupparams.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableIoTLoggingParams`

Specifies the logging level and the role with permissions for logging. You cannot
specify a logging level of `DISABLED`.

_Required_: No

_Type_: [EnableIoTLoggingParams](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-mitigationaction-enableiotloggingparams.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublishFindingToSnsParams`

Specifies the topic to which the finding should be published.

_Required_: No

_Type_: [PublishFindingToSnsParams](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-mitigationaction-publishfindingtosnsparams.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplaceDefaultPolicyVersionParams`

Replaces the policy version with a default or blank policy. You specify the template
name. Only a value of `BLANK_POLICY` is currently supported.

_Required_: No

_Type_: [ReplaceDefaultPolicyVersionParams](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-mitigationaction-replacedefaultpolicyversionparams.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdateCACertificateParams`

Specifies the new state for the CA certificate. Only a value of `DEACTIVATE`
is currently supported.

_Required_: No

_Type_: [UpdateCACertificateParams](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-mitigationaction-updatecacertificateparams.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdateDeviceCertificateParams`

Specifies the new state for a device certificate. Only a value of
`DEACTIVATE` is currently supported.

_Required_: No

_Type_: [UpdateDeviceCertificateParams](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-mitigationaction-updatedevicecertificateparams.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IoT::MitigationAction

AddThingsToThingGroupParams
