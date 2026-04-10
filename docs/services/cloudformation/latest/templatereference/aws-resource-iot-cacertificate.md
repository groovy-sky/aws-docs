This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::CACertificate

Specifies a CA certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::CACertificate",
  "Properties" : {
      "AutoRegistrationStatus" : String,
      "CACertificatePem" : String,
      "CertificateMode" : String,
      "RegistrationConfig" : RegistrationConfig,
      "RemoveAutoRegistration" : Boolean,
      "Status" : String,
      "Tags" : [ Tag, ... ],
      "VerificationCertificatePem" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoT::CACertificate
Properties:
  AutoRegistrationStatus: String
  CACertificatePem: String
  CertificateMode: String
  RegistrationConfig:
    RegistrationConfig
  RemoveAutoRegistration: Boolean
  Status: String
  Tags:
    - Tag
  VerificationCertificatePem: String

```

## Properties

`AutoRegistrationStatus`

Whether the CA certificate is configured for auto registration of device
certificates. Valid values are "ENABLE" and "DISABLE".

_Required_: No

_Type_: String

_Allowed values_: `ENABLE | DISABLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CACertificatePem`

The certificate data in PEM format.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `1`

_Maximum_: `65536`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CertificateMode`

The mode of the CA.

All the device certificates that are registered using this CA will be registered in the
same mode as the CA. For more information about certificate mode for device certificates,
see [certificate mode](../../../../reference/iot/latest/apireference/api-certificatedescription.md#iot-Type-CertificateDescription-certificateMode).

Valid values are "DEFAULT" and "SNI\_ONLY".

_Required_: No

_Type_: String

_Allowed values_: `DEFAULT | SNI_ONLY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RegistrationConfig`

Information about the registration configuration.

_Required_: No

_Type_: [RegistrationConfig](aws-properties-iot-cacertificate-registrationconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveAutoRegistration`

If true, removes auto registration.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the CA certificate.

Valid values are "ACTIVE" and "INACTIVE".

_Required_: Yes

_Type_: String

_Allowed values_: `ACTIVE | INACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-iot-cacertificate-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerificationCertificatePem`

The private key verification certificate.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `1`

_Maximum_: `65536`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the CA certificate ID.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) for the CA certificate. For example:

`{ "Fn::GetAtt": ["MyCACertificate", "Arn"] }`

A value similar to the following is returned:

`arn:aws:iot:us-east-1:123456789012:cacert/a6be6b84559801927e35a8f901fae08b5971d78d1562e29504ff9663b276a5f5`

`Id`

The CA certificate ID.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

RegistrationConfig

All content copied from https://docs.aws.amazon.com/.
