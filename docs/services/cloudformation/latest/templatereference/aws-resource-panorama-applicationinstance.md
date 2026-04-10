This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Panorama::ApplicationInstance

###### Important

End of support notice: On May 31, 2026, AWS will end support for AWS Panorama. After May 31, 2026,
you will no longer be able to access the AWS Panorama console or AWS Panorama resources. For more information, see
[AWS Panorama end of support](../../../panorama/latest/dev/panorama-end-of-support.md).

Creates an application instance and deploys it to a device.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Panorama::ApplicationInstance",
  "Properties" : {
      "ApplicationInstanceIdToReplace" : String,
      "DefaultRuntimeContextDevice" : String,
      "Description" : String,
      "ManifestOverridesPayload" : ManifestOverridesPayload,
      "ManifestPayload" : ManifestPayload,
      "Name" : String,
      "RuntimeRoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Panorama::ApplicationInstance
Properties:
  ApplicationInstanceIdToReplace: String
  DefaultRuntimeContextDevice: String
  Description: String
  ManifestOverridesPayload:
    ManifestOverridesPayload
  ManifestPayload:
    ManifestPayload
  Name: String
  RuntimeRoleArn: String
  Tags:
    - Tag

```

## Properties

`ApplicationInstanceIdToReplace`

The ID of an application instance to replace with the new instance.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-\_]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefaultRuntimeContextDevice`

The device's ID.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-\_]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description for the application instance.

_Required_: No

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManifestOverridesPayload`

Setting overrides for the application manifest.

_Required_: No

_Type_: [ManifestOverridesPayload](aws-properties-panorama-applicationinstance-manifestoverridespayload.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManifestPayload`

The application's manifest document.

_Required_: Yes

_Type_: [ManifestPayload](aws-properties-panorama-applicationinstance-manifestpayload.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A name for the application instance.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-\_]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RuntimeRoleArn`

The ARN of a runtime role for the application instance.

_Required_: No

_Type_: String

_Pattern_: `^arn:[a-z0-9][-.a-z0-9]{0,62}:iam::[0-9]{12}:role/.+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags for the application instance.

_Required_: No

_Type_: Array of [Tag](aws-properties-panorama-applicationinstance-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ApplicationInstanceId`

The application instance's ID.

`Arn`

The application instance's ARN.

`CreatedTime`

The application instance's created time.

`DefaultRuntimeContextDeviceName`

The application instance's default runtime context device name.

`HealthStatus`

The application instance's health status.

`LastUpdatedTime`

The application instance's last updated time.

`Status`

The application instance's status.

`StatusDescription`

The application instance's status description.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Panorama

ManifestOverridesPayload

All content copied from https://docs.aws.amazon.com/.
