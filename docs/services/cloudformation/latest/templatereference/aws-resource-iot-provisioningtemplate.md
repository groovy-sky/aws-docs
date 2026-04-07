This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::ProvisioningTemplate

Creates a fleet provisioning template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::ProvisioningTemplate",
  "Properties" : {
      "Description" : String,
      "Enabled" : Boolean,
      "PreProvisioningHook" : ProvisioningHook,
      "ProvisioningRoleArn" : String,
      "Tags" : [ Tag, ... ],
      "TemplateBody" : String,
      "TemplateName" : String,
      "TemplateType" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoT::ProvisioningTemplate
Properties:
  Description: String
  Enabled: Boolean
  PreProvisioningHook:
    ProvisioningHook
  ProvisioningRoleArn: String
  Tags:
    - Tag
  TemplateBody: String
  TemplateName: String
  TemplateType: String

```

## Properties

`Description`

The description of the fleet provisioning template.

_Required_: No

_Type_: String

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

True to enable the fleet provisioning template, otherwise false.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreProvisioningHook`

Creates a pre-provisioning hook template.

_Required_: No

_Type_: [ProvisioningHook](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-provisioningtemplate-provisioninghook.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisioningRoleArn`

The role ARN for the role associated with the fleet provisioning template. This IoT
role grants permission to provision a device.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that can be used to manage the fleet provisioning template.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-provisioningtemplate-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateBody`

The JSON formatted contents of the fleet provisioning template version.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateName`

The name of the fleet provisioning template.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z_-]+$`

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TemplateType`

The type of the provisioning template.

_Required_: No

_Type_: String

_Allowed values_: `FLEET_PROVISIONING | JITP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the template name. For example:

`{ "Ref": "MyTemplate" }`

For a stack named MyStack, a value similar to the following is returned:

`MyStack-MyTemplate-AB1CDEFGHIJK`

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`TemplateArn`

The ARN that identifies the provisioning template.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IoT::PolicyPrincipalAttachment

ProvisioningHook
