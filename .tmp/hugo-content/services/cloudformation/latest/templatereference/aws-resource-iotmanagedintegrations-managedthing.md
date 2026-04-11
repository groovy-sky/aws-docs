This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTManagedIntegrations::ManagedThing

Creates a managed thing. A managed thing contains the device identifier, protocol
supported, and capabilities of the device in a protocol-specific format.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTManagedIntegrations::ManagedThing",
  "Properties" : {
      "AuthenticationMaterial" : String,
      "AuthenticationMaterialType" : String,
      "Brand" : String,
      "CapabilityReport" : CapabilityReport,
      "Classification" : String,
      "CredentialLockerId" : String,
      "MetaData" : {Key: Value, ...},
      "Model" : String,
      "Name" : String,
      "Owner" : String,
      "Role" : String,
      "SerialNumber" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::IoTManagedIntegrations::ManagedThing
Properties:
  AuthenticationMaterial: String
  AuthenticationMaterialType: String
  Brand: String
  CapabilityReport:
    CapabilityReport
  Classification: String
  CredentialLockerId: String
  MetaData:
    Key: Value
  Model: String
  Name: String
  Owner: String
  Role: String
  SerialNumber: String
  Tags:
    Key: Value

```

## Properties

`AuthenticationMaterial`

The authentication material defining the device connectivity setup requests. The
authentication materials used are the device bar code.

_Required_: No

_Type_: String

_Pattern_: ``^[0-9A-Za-z!#$%&()*\+\-;<=>?@^_`{|}~\/: ]+$``

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuthenticationMaterialType`

The type of authentication material used for device connectivity setup
requests.

_Required_: No

_Type_: String

_Allowed values_: `WIFI_SETUP_QR_BAR_CODE | ZWAVE_QR_BAR_CODE | ZIGBEE_QR_BAR_CODE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Brand`

The brand of the device.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9-_ ]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CapabilityReport`

A report of the capabilities for the managed thing.

_Required_: No

_Type_: [CapabilityReport](aws-properties-iotmanagedintegrations-managedthing-capabilityreport.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Classification`

The classification of the managed thing such as light bulb or
thermostat.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CredentialLockerId`

The identifier of the credential locker for the managed thing.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetaData`

The metadata for the managed thing.

###### Note

The `managedThing metadata` parameter is used for associating
attributes with a `managedThing` that can be used for grouping over-the-air
(OTA) tasks. Name value pairs in `metadata ` can be used in the
`OtaTargetQueryString` parameter for the `CreateOtaTask` API
operation.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9_.,@/:#-]+`

_Maximum_: `800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Model`

The model of the device.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9-_ ]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the managed thing representing the physical
device.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Owner`

Owner of the device, usually an indication of whom the device belongs to. This
value should not contain personal identifiable information.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.,@-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Role`

The type of device used. This will be the hub controller, cloud device, or IoT
device.

_Required_: Yes

_Type_: String

_Allowed values_: `CONTROLLER | DEVICE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SerialNumber`

The serial number of the device.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9-_ ]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A set of key/value pairs that are used to manage the managed thing.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the managed thing name.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the managed thing.

`CreatedAt`

The timestamp value of when the device creation request occurred.

`Id`

The id of the managed thing.

`Identifier`

The id of the managed thing.

`ProvisioningStatus`

The provisioning status of the device in the provisioning workflow for onboarding to
IoT managed integrations.

`UniversalProductCode`

The universal product code (UPC) of the device model. The UPC is typically used in
the United States of America and Canada.

`UpdatedAt`

The timestamp value of when the managed thing was last updated at.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTManagedIntegrations::CredentialLocker

CapabilityReport

All content copied from https://docs.aws.amazon.com/.
