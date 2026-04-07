This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTCoreDeviceAdvisor::SuiteDefinition

Creates a Device Advisor test suite.

Requires permission to access the [CreateSuiteDefinition](https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions) action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTCoreDeviceAdvisor::SuiteDefinition",
  "Properties" : {
      "SuiteDefinitionConfiguration" : SuiteDefinitionConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTCoreDeviceAdvisor::SuiteDefinition
Properties:
  SuiteDefinitionConfiguration:
    SuiteDefinitionConfiguration
  Tags:
    - Tag

```

## Properties

`SuiteDefinitionConfiguration`

Gets the suite definition configuration.

_Required_: Yes

_Type_: [SuiteDefinitionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that can be used to manage the the Suite Definition.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotcoredeviceadvisor-suitedefinition-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the Suite Definition name.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`SuiteDefinitionArn`

The Arn of the Suite Definition.

`SuiteDefinitionId`

The version of the Suite Definition.

`SuiteDefinitionVersion`

The ID of the Suite Definition.

## Examples

The following example is a template of a `SuiteDefintion`.

#### JSON

```json

{
   "Resources": {
      "TestSuiteDefinition": {
         "Type": "AWS::IoTCoreDeviceAdvisor::SuiteDefinition",
         "Properties": {
            "SuiteDefinitionConfiguration": {
               "SuiteDefinitionName": "SuiteDefinitionName",
               "DevicePermissionRoleArn": "arn:aws:iam::123456789012:role/RoleName",
               "Devices": [
                  {
                     "ThingArn": "arn:aws:iot:us-east-1:123456789012:thing/ThingName"
                  }
               ],
               "RootGroup": "{\n\"configuration\": {},\n\"tests\": [{\n\"name\": \"TestGroup\",\n\"configuration\": {\n\"EXECUTION_TIMEOUT\": \"30\"\n},\n\"tests\": [{\n\"name\": \"MQTTPublishTest\",\n\"configuration\": {\n\"TOPIC_FOR_PUBLISH_VALIDATION\": \"target\"\n},\n\"test\": {\n\"id\": \"MQTT_Publish\",\n\"version\": \"0.0.0\"\n}\n}]\n}]\n}",
               "IntendedForQualification": false
            }
         }
      }
   }
}
```

#### YAML

```yaml

Resources:
  TestSuiteDefinition:
    Type: 'AWS::IoTCoreDeviceAdvisor::SuiteDefinition'
    Properties:
      SuiteDefinitionConfiguration:
        SuiteDefinitionName: "SuiteDefinitionName"
        DevicePermissionRoleArn: "arn:aws:iam::123456789012:role/RoleName"
        Devices:
          - ThingArn: "arn:aws:iot:us-east-1:123456789012:thing/ThingName"
        RootGroup: '{
                "configuration": {},
                "tests": [{
                    "name": "TestGroup",
                    "configuration": {
                        "EXECUTION_TIMEOUT": "30"
                    },
                    "tests": [{
                        "name": "MQTTPublishTest",
                        "configuration": {
                            "TOPIC_FOR_PUBLISH_VALIDATION": "target"
                        },
                        "test": {
                            "id": "MQTT_Publish",
                            "version": "0.0.0"
                        }
                    }]
                }]
            }'
        IntendedForQualification: false
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS IoT Core Device Advisor

DeviceUnderTest
