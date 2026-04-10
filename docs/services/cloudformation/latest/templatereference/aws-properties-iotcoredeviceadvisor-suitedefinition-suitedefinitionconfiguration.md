This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTCoreDeviceAdvisor::SuiteDefinition SuiteDefinitionConfiguration

The configuration of the Suite Definition. Listed below are the required elements of the `SuiteDefinitionConfiguration`.

**devicePermissionRoleArn**

The device permission arn.

This is a required element.

**Type:** String

**devices**

The list of configured devices under test. For more information on devices under test, see [DeviceUnderTest](https://amazonaws.com/iot/latest/apireference/API_iotdeviceadvisor_DeviceUnderTest.html)

Not a required element.

**Type:** List of devices under test

**intendedForQualification**

The tests intended for qualification in a suite.

Not a required element.

**Type:** Boolean

**rootGroup**

The test suite root group. For more information on creating and using root groups see the [Device Advisor workflow](../../../iot/latest/developerguide/device-advisor-workflow.md).

This is a required element.

**Type:** String

**suiteDefinitionName**

The Suite Definition Configuration name.

This is a required element.

**Type:** String

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DevicePermissionRoleArn" : String,
  "Devices" : [ DeviceUnderTest, ... ],
  "IntendedForQualification" : Boolean,
  "RootGroup" : String,
  "SuiteDefinitionName" : String
}

```

### YAML

```yaml

  DevicePermissionRoleArn: String
  Devices:
    - DeviceUnderTest
  IntendedForQualification: Boolean
  RootGroup: String
  SuiteDefinitionName: String

```

## Properties

`DevicePermissionRoleArn`

Gets the device permission ARN. This is a required parameter.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Devices`

Gets the devices configured.

_Required_: No

_Type_: Array of [DeviceUnderTest](aws-properties-iotcoredeviceadvisor-suitedefinition-deviceundertest.md)

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntendedForQualification`

Gets the tests intended for qualification in a suite.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RootGroup`

Gets the test suite root group. This is a required parameter.
For updating or creating the latest qualification suite,
if `intendedForQualification` is set to true,
`rootGroup` can be an empty string. If `intendedForQualification` is false,
`rootGroup` cannot be an empty string.
If `rootGroup` is empty, and
`intendedForQualification` is set to true,
all the qualification tests are included, and the configuration is default.

For a qualification suite, the minimum length is 0, and the maximum is 2048. For a
non-qualification suite, the minimum length is 1, and the maximum is 2048.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuiteDefinitionName`

Gets the suite definition name. This is a required parameter.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeviceUnderTest

Tag

All content copied from https://docs.aws.amazon.com/.
