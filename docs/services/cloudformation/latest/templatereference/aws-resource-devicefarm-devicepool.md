This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DeviceFarm::DevicePool

Represents a request to the create device pool operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DeviceFarm::DevicePool",
  "Properties" : {
      "Description" : String,
      "MaxDevices" : Integer,
      "Name" : String,
      "ProjectArn" : String,
      "Rules" : [ Rule, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DeviceFarm::DevicePool
Properties:
  Description: String
  MaxDevices: Integer
  Name: String
  ProjectArn: String
  Rules:
    - Rule
  Tags:
    - Tag

```

## Properties

`Description`

The device pool's description.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxDevices`

The number of devices that Device Farm can add to your device pool. Device Farm adds devices that are
available and meet the criteria that you assign for the `rules` parameter. Depending on how many
devices meet these constraints, your device pool might contain fewer devices than the value for this
parameter.

By specifying the maximum number of devices, you can control the costs that you incur
by running tests.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The device pool's name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectArn`

The ARN of the project for the device pool.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.+`

_Minimum_: `32`

_Maximum_: `1011`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Rules`

The device pool's rules.

_Required_: Yes

_Type_: Array of [Rule](aws-properties-devicefarm-devicepool-rule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md) in the
_AWS CloudFormation guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-devicefarm-devicepool-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

Not supported for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the device pool. See [Amazon resource names](../../../../general/latest/gr/aws-arns-and-namespaces.md) in the
_General Reference guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Device Farm

Rule

All content copied from https://docs.aws.amazon.com/.
