This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DeviceFarm::Project

Creates a project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DeviceFarm::Project",
  "Properties" : {
      "DefaultJobTimeoutMinutes" : Integer,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "VpcConfig" : VpcConfig
    }
}

```

### YAML

```yaml

Type: AWS::DeviceFarm::Project
Properties:
  DefaultJobTimeoutMinutes: Integer
  Name: String
  Tags:
    - Tag
  VpcConfig:
    VpcConfig

```

## Properties

`DefaultJobTimeoutMinutes`

Sets the execution timeout value (in minutes) for a project. All test runs in this project use the
specified execution timeout value unless overridden when scheduling a run.

_Required_: No

_Type_: Integer

_Minimum_: `5`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The project's name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to the resource. A tag is an array of key-value pairs. Tag keys can have a maximum
character length of 128 characters. Tag values can have a maximum length of 256 characters.

_Required_: No

_Type_: Array of [Tag](aws-properties-devicefarm-project-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfig`

The VPC security groups and subnets that are attached to a project.

_Required_: No

_Type_: [VpcConfig](aws-properties-devicefarm-project-vpcconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

Not supported for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the project. See [Amazon resource names](../../../../general/latest/gr/aws-arns-and-namespaces.md) in the
_General Reference guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
