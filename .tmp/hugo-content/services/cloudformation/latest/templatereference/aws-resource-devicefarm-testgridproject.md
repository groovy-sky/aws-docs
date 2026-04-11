This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DeviceFarm::TestGridProject

A Selenium testing project. Projects are used to collect and collate sessions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DeviceFarm::TestGridProject",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "VpcConfig" : VpcConfig
    }
}

```

### YAML

```yaml

Type: AWS::DeviceFarm::TestGridProject
Properties:
  Description: String
  Name: String
  Tags:
    - Tag
  VpcConfig:
    VpcConfig

```

## Properties

`Description`

A human-readable description for the project.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A human-readable name for the project.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md) in the
_AWS CloudFormation guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-devicefarm-testgridproject-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfig`

The VPC security groups and subnets that are attached to a project.

_Required_: No

_Type_: [VpcConfig](aws-properties-devicefarm-testgridproject-vpcconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

Not supported for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the `TestGrid` project. See [Amazon resource names](../../../../general/latest/gr/aws-arns-and-namespaces.md) in the
_General Reference guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfig

Tag

All content copied from https://docs.aws.amazon.com/.
