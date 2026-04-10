This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTCoreDeviceAdvisor::SuiteDefinition DeviceUnderTest

Information of a test device. A thing ARN, certificate ARN
or device role ARN is required.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateArn" : String,
  "ThingArn" : String
}

```

### YAML

```yaml

  CertificateArn: String
  ThingArn: String

```

## Properties

`CertificateArn`

Lists device's certificate ARN.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThingArn`

Lists device's thing ARN.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTCoreDeviceAdvisor::SuiteDefinition

SuiteDefinitionConfiguration

All content copied from https://docs.aws.amazon.com/.
