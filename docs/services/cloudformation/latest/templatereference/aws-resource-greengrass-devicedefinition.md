This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::DeviceDefinition

The `AWS::Greengrass::DeviceDefinition` resource represents a device
definition for AWS IoT Greengrass. Device definitions are used to organize your device
definition versions.

Device definitions can reference multiple device definition versions. All device
definition versions must be associated with a device definition. Each device definition
version can contain one or more devices.

###### Note

When you create a device definition, you can optionally include an initial device
definition version. To associate a device definition version later, create an [`AWS::Greengrass::DeviceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html) resource and
specify the ID of this device definition.

After you create the device definition version that contains the devices you want to
deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::DeviceDefinition",
  "Properties" : {
      "InitialVersion" : DeviceDefinitionVersion,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::DeviceDefinition
Properties:
  InitialVersion:
    DeviceDefinitionVersion
  Name: String
  Tags:
    - Tag

```

## Properties

`InitialVersion`

The device definition version to include when the device definition is created. A device
definition version contains a list of [`device`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-devicedefinition-device.html) property types.

###### Note

To associate a device definition version after the device definition is created,
create an [`AWS::Greengrass::DeviceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html) resource and
specify the ID of this device definition.

_Required_: No

_Type_: [DeviceDefinitionVersion](aws-properties-greengrass-devicedefinition-devicedefinitionversion.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the device definition.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Application-specific metadata to attach to the device definition. You can use tags in
IAM policies to control access to AWS IoT Greengrass resources. You
can also use tags to categorize your resources. For more information, see [Tagging Your\
AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/v1/developerguide/tagging.html) in the _AWS IoT Greengrass Version 1 Developer Guide_.

This `Json` property type is processed as a map of key-value pairs. It uses
the following format, which is different from most `Tags` implementations in
CloudFormation templates.

```json

"Tags": {
    "KeyName0": "value",
    "KeyName1": "value",
    "KeyName2": "value"
}
```

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the device definition, such as
`1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the `DeviceDefinition`, such as
`arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/devices/1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

`Id`

The ID of the `DeviceDefinition`, such as
`1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

`LatestVersionArn`

The ARN of the last `DeviceDefinitionVersion` that was added to the
`DeviceDefinition`, such as `arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/devices/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

`Name`

The name of the device definition.

## Examples

### Device Definition Snippet

The following snippet defines a device definition resource with an initial version
that contains a device. This example points to a manually generated device
certificate.

For an example of a complete template, see the [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) resource.

#### JSON

```json

"TestDeviceDefinition": {
    "Type": "AWS::Greengrass::DeviceDefinition",
    "Properties": {
        "Name": "DemoTestDeviceDefinition",
        "InitialVersion": {
            "Devices": [
                {
                    "Id": "TestDevice1",
                    "ThingArn": {
                        "Fn::Join": [
                            ":",
                            [
                                "arn:aws:iot",
                                {
                                    "Ref": "AWS::Region"
                                },
                                {
                                    "Ref": "AWS::AccountId"
                                },
                                "thing/TestDevice1"
                            ]
                        ]
                    },
                    "CertificateArn": {
                        "Fn::Join": [
                            ":",
                            [
                                "arn:aws:iot",
                                {
                                    "Ref": "AWS::Region"
                                },
                                {
                                    "Ref": "AWS::AccountId"
                                },
                                "cert/4db8b7f58d95b7fdb45c38c28a0b1adf6c5f8c03e902de65734935fea83e751f"
                            ]
                        ]
                    },
                    "SyncShadow": "true"
                }
            ]
        }
    }
}
```

#### YAML

```yaml

TestDeviceDefinition:
  Type: 'AWS::Greengrass::DeviceDefinition'
  Properties:
    Name: DemoTestDeviceDefinition
    InitialVersion:
      Devices:
        - Id: TestDevice1
          ThingArn: !Join
            - ':'
            - - 'arn:aws:iot'
              - !Ref 'AWS::Region'
              - !Ref 'AWS::AccountId'
              - thing/TestDevice1
          CertificateArn: !Join
            - ':'
            - - 'arn:aws:iot'
              - !Ref 'AWS::Region'
              - !Ref 'AWS::AccountId'
              - >-
                cert/4db8b7f58d95b7fdb45c38c28a0b1adf6c5f8c03e902de65734935fea83e751f
          SyncShadow: 'true'
```

## See also

- [CreateDeviceDefinition](https://docs.aws.amazon.com/greengrass/v1/apireference/createdevicedefinition-post.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Core

Device
