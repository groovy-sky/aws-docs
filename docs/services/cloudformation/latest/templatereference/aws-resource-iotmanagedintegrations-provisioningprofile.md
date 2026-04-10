This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTManagedIntegrations::ProvisioningProfile

Create a provisioning profile for a device to execute the provisioning flows using
a provisioning template. The provisioning template is a document that defines the set of
resources and policies applied to a device during the provisioning process.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTManagedIntegrations::ProvisioningProfile",
  "Properties" : {
      "CaCertificate" : String,
      "Name" : String,
      "ProvisioningType" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::IoTManagedIntegrations::ProvisioningProfile
Properties:
  CaCertificate: String
  Name: String
  ProvisioningType: String
  Tags:
    Key: Value

```

## Properties

`CaCertificate`

The id of the certificate authority (CA) certificate.

_Required_: No

_Type_: String

_Pattern_: `^-----BEGIN CERTIFICATE-----.*(.|\
)*-----END CERTIFICATE-----\n?$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the provisioning template.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z_-]+$`

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProvisioningType`

The type of provisioning workflow the device uses for onboarding to IoT managed
integrations.

_Required_: Yes

_Type_: String

_Allowed values_: `FLEET_PROVISIONING | JITR`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A set of key/value pairs that are used to manage the provisioning
profile.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the provisioning profile template name

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the provisioning template used in the
provisioning profile.

`ClaimCertificate`

The id of the claim certificate.

`Id`

The provisioning profile id.

`Identifier`

The provisioning template the device uses for the provisioning process.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapabilityReportEndpoint

Next

All content copied from https://docs.aws.amazon.com/.
