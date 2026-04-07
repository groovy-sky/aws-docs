This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DeviceFarm::VPCEConfiguration

Creates a configuration record in Device Farm for your Amazon Virtual Private Cloud (VPC) endpoint
service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DeviceFarm::VPCEConfiguration",
  "Properties" : {
      "ServiceDnsName" : String,
      "Tags" : [ Tag, ... ],
      "VpceConfigurationDescription" : String,
      "VpceConfigurationName" : String,
      "VpceServiceName" : String
    }
}

```

### YAML

```yaml

Type: AWS::DeviceFarm::VPCEConfiguration
Properties:
  ServiceDnsName: String
  Tags:
    - Tag
  VpceConfigurationDescription: String
  VpceConfigurationName: String
  VpceServiceName: String

```

## Properties

`ServiceDnsName`

The DNS name that Device Farm will use to map to the private service you want to access.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the
_AWS CloudFormation guide_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devicefarm-vpceconfiguration-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpceConfigurationDescription`

An optional description that provides details about your VPC endpoint configuration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpceConfigurationName`

The friendly name you give to your VPC endpoint configuration to manage your configurations more
easily.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpceServiceName`

The name of the VPC endpoint service that you want to access from Device Farm.

The name follows the format `com.amazonaws.vpce.us-west-2.vpce-svc-id`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

Not supported for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the VPC endpoint. See [Amazon resource names](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the
_General Reference guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcConfig

Tag
