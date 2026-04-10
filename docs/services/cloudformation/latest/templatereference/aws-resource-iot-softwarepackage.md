This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::SoftwarePackage

Use the `AWS::IoT::SoftwarePackage` resource to create a software
package.

For information about working with software packages, see [AWS IoT Device Management Software Package Catalog](../../../iot/latest/developerguide/software-package-catalog.md) and [Creating a software package and package version](../../../iot/latest/developerguide/creating-package-and-version.md) in the _AWS IoT Developer Guide_. See also, [CreatePackage](../../../../reference/iot/latest/apireference/api-createpackage.md) in the
_API Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::SoftwarePackage",
  "Properties" : {
      "Description" : String,
      "PackageName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoT::SoftwarePackage
Properties:
  Description: String
  PackageName: String
  Tags:
    - Tag

```

## Properties

`Description`

A summary of the package being created. This can be used to outline the package's
contents or purpose.

_Required_: No

_Type_: String

_Pattern_: `^[^\p{C}]+$`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PackageName`

The name of the new software package.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_.]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata that can be used to manage the package.

_Required_: No

_Type_: Array of [Tag](aws-properties-iot-softwarepackage-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the software package name. For example:

```

{ "Ref": "MyPackage" }
```

Response:

```

package-name
```

For a stack named MyStack, a value that is similar to the following is returned:

```

MyStack-MySoftwarePackage-AB1CDEFGHIJK
```

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

For example:

```

{ "GetAtt": "MySoftwarePackage.PackageArn" }
```

`PackageArn`

The Amazon Resource Name (ARN) for the package.

## Examples

The following example declares a software package and the values of its
attributes.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MySoftwarePackage": {
      "Type": "AWS::IoT::SoftwarePackage",
      "Properties": {
        "PackageName": "Package",
        "Description": "description",
        "Tags": [
          {
            "Key": "TagKey",
            "Value": "TagValue"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  MySoftwarePackage:
    Type: AWS::IoT::SoftwarePackage
    Properties:
      PackageName: "Package"
      Description: "description"
      Tags:
        - Key: "TagKey"
          Value: "TagValue"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
