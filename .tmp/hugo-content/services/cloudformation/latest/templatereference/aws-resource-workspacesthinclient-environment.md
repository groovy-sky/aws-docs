This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesThinClient::Environment

###### Important

End of support notice: On March 31, 2027, AWS will end support for Amazon WorkSpaces Thin Client. After March 31, 2027,
you will no longer be able to access the WorkSpaces Thin Client console or WorkSpaces Thin Client resources. For more information, see
[Amazon WorkSpaces Thin Client end of support](../../../workspaces-thin-client/latest/ug/workspacesthinclient-end-of-support.md).

Describes an environment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpacesThinClient::Environment",
  "Properties" : {
      "DesiredSoftwareSetId" : String,
      "DesktopArn" : String,
      "DesktopEndpoint" : String,
      "DeviceCreationTags" : [ Tag, ... ],
      "KmsKeyArn" : String,
      "MaintenanceWindow" : MaintenanceWindow,
      "Name" : String,
      "SoftwareSetUpdateMode" : String,
      "SoftwareSetUpdateSchedule" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpacesThinClient::Environment
Properties:
  DesiredSoftwareSetId: String
  DesktopArn: String
  DesktopEndpoint: String
  DeviceCreationTags:
    - Tag
  KmsKeyArn: String
  MaintenanceWindow:
    MaintenanceWindow
  Name: String
  SoftwareSetUpdateMode: String
  SoftwareSetUpdateSchedule: String
  Tags:
    - Tag

```

## Properties

`DesiredSoftwareSetId`

The ID of the software set to apply.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{1,9}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DesktopArn`

The Amazon Resource Name (ARN) of the desktop to stream from Amazon WorkSpaces,
WorkSpaces Secure Browser, or WorkSpaces Applications.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:[a-zA-Z0-9\-]+:[a-zA-Z0-9\-]*:[0-9]{0,12}:[a-zA-Z0-9\-\/\._]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DesktopEndpoint`

The URL for the identity provider login (only for environments that use WorkSpaces Applications).

_Required_: No

_Type_: String

_Pattern_: `^(https:\/\/)[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,32}(:[0-9]{1,5})?(\/.*)?$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceCreationTags`

An array of key-value pairs to apply to the newly created devices for this environment.

_Required_: No

_Type_: Array of [Tag](aws-properties-workspacesthinclient-environment-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The Amazon Resource Name (ARN) of the AWS Key Management Service key used to encrypt the
environment.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:kms:[a-zA-Z0-9\-]*:[0-9]{0,12}:key\/[a-zA-Z0-9-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaintenanceWindow`

A specification for a time window to apply software updates.

_Required_: No

_Type_: [MaintenanceWindow](aws-properties-workspacesthinclient-environment-maintenancewindow.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the environment.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SoftwareSetUpdateMode`

An option to define which software updates to apply.

_Required_: No

_Type_: String

_Allowed values_: `USE_LATEST | USE_DESIRED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SoftwareSetUpdateSchedule`

An option to define if software updates should be applied within a maintenance
window.

_Required_: No

_Type_: String

_Allowed values_: `USE_MAINTENANCE_WINDOW | APPLY_IMMEDIATELY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-workspacesthinclient-environment-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ActivationCode`

The activation code to register a device to the environment.

`Arn`

The Amazon Resource Name (ARN) of the environment.

`CreatedAt`

The timestamp of when the environment was created.

`DesktopType`

The type of streaming desktop for the environment.

`PendingSoftwareSetId`

The ID of the software set that is pending to be installed.

`PendingSoftwareSetVersion`

The version of the software set that is pending to be installed.

`RegisteredDevicesCount`

The number of devices registered to the environment.

`SoftwareSetComplianceStatus`

Describes if the software currently installed on all devices in the environment is a
supported version.

`UpdatedAt`

The timestamp of when the device was updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon WorkSpaces Thin Client

MaintenanceWindow

All content copied from https://docs.aws.amazon.com/.
