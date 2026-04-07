This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DeviceFarm::InstanceProfile

Creates a profile that can be applied to one or more private fleet device
instances.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DeviceFarm::InstanceProfile",
  "Properties" : {
      "Description" : String,
      "ExcludeAppPackagesFromCleanup" : [ String, ... ],
      "Name" : String,
      "PackageCleanup" : Boolean,
      "RebootAfterUse" : Boolean,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DeviceFarm::InstanceProfile
Properties:
  Description: String
  ExcludeAppPackagesFromCleanup:
    - String
  Name: String
  PackageCleanup: Boolean
  RebootAfterUse: Boolean
  Tags:
    - Tag

```

## Properties

`Description`

The description of the instance profile.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeAppPackagesFromCleanup`

An array of strings containing the list of app packages that should not be cleaned up from the device
after a test run completes.

The list of packages is considered only if you set `packageCleanup` to
`true`.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the instance profile.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PackageCleanup`

When set to `true`, Device Farm removes app packages after a test run. The default value is
`false` for private devices.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RebootAfterUse`

When set to `true`, Device Farm reboots the instance after a test run. The default value is
`true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the
_AWS CloudFormation guide_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devicefarm-instanceprofile-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

Not supported for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the instance profile. See [Amazon resource names](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the
_General Reference guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
