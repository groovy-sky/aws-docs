This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::DeviceDefinitionVersion

The `AWS::Greengrass::DeviceDefinitionVersion` resource represents a device
definition version for AWS IoT Greengrass. A device definition version contains a list of
devices.

###### Note

To create a device definition version, you must specify the ID of the device
definition that you want to associate with the version. For information about creating a
device definition, see [`AWS::Greengrass::DeviceDefinition`](../userguide/aws-resource-greengrass-devicedefinition.md).

After you create a device definition version that contains the devices you want to
deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::DeviceDefinitionVersion",
  "Properties" : {
      "DeviceDefinitionId" : String,
      "Devices" : [ Device, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::DeviceDefinitionVersion
Properties:
  DeviceDefinitionId: String
  Devices:
    - Device

```

## Properties

`DeviceDefinitionId`

The ID of the device definition associated with this version. This value is a
GUID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Devices`

The devices in this version.

_Required_: Yes

_Type_: Array of [Device](aws-properties-greengrass-devicedefinitionversion-device.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the device definition
version, such as `arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/devices/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Device Definition Version Snippet

The following snippet defines device definition and device definition version
resources. The device definition version references the device definition and
contains a device. This example points to a manually generated device
certificate.

For an example of a complete template, see the [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md) resource.

#### JSON

```json

"TestDeviceDefinition": {
    "Type": "AWS::Greengrass::DeviceDefinition",
    "Properties": {
        "Name": "DemoTestDeviceDefinition"
    }
},
"TestDeviceDefinitionVersion": {
    "Type": "AWS::Greengrass::DeviceDefinitionVersion",
    "Properties": {
        "DeviceDefinitionId": {
            "Fn::GetAtt": [
                "TestDeviceDefinition",
                "Id"
            ]
        },
        "Devices": [
            {
                "Id": "TestDevice1",
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
                "SyncShadow": "true",
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
                }
            }
        ]
    }
}
```

#### YAML

```yaml

TestDeviceDefinition:
  Type: 'AWS::Greengrass::DeviceDefinition'
  Properties:
    Name: DemoTestDeviceDefinition
TestDeviceDefinitionVersion:
  Type: 'AWS::Greengrass::DeviceDefinitionVersion'
  Properties:
    DeviceDefinitionId: !GetAtt
      - TestDeviceDefinition
      - Id
    Devices:
      - Id: TestDevice1
        CertificateArn: !Join
          - ':'
          - - 'arn:aws:iot'
            - !Ref 'AWS::Region'
            - !Ref 'AWS::AccountId'
            - >-
              cert/4db8b7f58d95b7fdb45c38c28a0b1adf6c5f8c03e902de65734935fea83e751f
        SyncShadow: 'true'
        ThingArn: !Join
          - ':'
          - - 'arn:aws:iot'
            - !Ref 'AWS::Region'
            - !Ref 'AWS::AccountId'
            - thing/TestDevice1
```

## See also

- [CreateDeviceDefinitionVersion](../../../../reference/greengrass/v1/apireference/createdevicedefinitionversion-post.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeviceDefinitionVersion

Device

All content copied from https://docs.aws.amazon.com/.
