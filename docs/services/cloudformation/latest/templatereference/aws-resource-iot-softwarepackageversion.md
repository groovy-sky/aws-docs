This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::SoftwarePackageVersion

Use the `AWS::IoT::SoftwarePackageVersion` resource to create a software
package version.

For information about working with software package versions, see [AWS IoT Device ManagementSoftware Package Catalog](../../../iot/latest/developerguide/software-package-catalog.md) and [Creating a software package and package version](../../../iot/latest/developerguide/creating-package-and-version.md) in the _AWS IoT Developer Guide_. See also, [CreatePackageVersion](../../../../reference/iot/latest/apireference/api-createpackageversion.md)
in the _API Guide_.

###### Note

The associated software package must exist before the package version is created. If
you create a software package and package version in the same CloudFormation template,
set the software package as a [dependency](../userguide/aws-attribute-dependson.md) of the package version. If they are created out of sequence, you
will receive an error.

Package versions and created in a `draft` state, for more information, see
[Package version lifecycle](../../../iot/latest/developerguide/preparing-to-use-software-package-catalog.md#package-version-lifecycle). To change the package version state after it’s
created, use the [UpdatePackageVersionAPI](../../../../reference/iot/latest/apireference/api-updatepackageversion.md) command.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::SoftwarePackageVersion",
  "Properties" : {
      "Artifact" : PackageVersionArtifact,
      "Attributes" : {Key: Value, ...},
      "Description" : String,
      "PackageName" : String,
      "Recipe" : String,
      "Sbom" : Sbom,
      "Tags" : [ Tag, ... ],
      "VersionName" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoT::SoftwarePackageVersion
Properties:
  Artifact:
    PackageVersionArtifact
  Attributes:
    Key: Value
  Description: String
  PackageName: String
  Recipe: String
  Sbom:
    Sbom
  Tags:
    - Tag
  VersionName: String

```

## Properties

`Artifact`

Property description not available.

_Required_: No

_Type_: [PackageVersionArtifact](aws-properties-iot-softwarepackageversion-packageversionartifact.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Attributes`

Metadata that can be used to define a package version’s configuration. For example, the
S3 file location, configuration options that are being sent to the device or fleet.

The combined size of all the attributes on a package version is limited to 3KB.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9:_-]+$`

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A summary of the package version being created. This can be used to outline the
package's contents or purpose.

_Required_: No

_Type_: String

_Pattern_: `^[^\p{C}]+$`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PackageName`

The name of the associated software package.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_.]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Recipe`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sbom`

Property description not available.

_Required_: No

_Type_: [Sbom](aws-properties-iot-softwarepackageversion-sbom.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that can be used to manage the package version.

_Required_: No

_Type_: Array of [Tag](aws-properties-iot-softwarepackageversion-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionName`

The name of the new package version.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_.]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the software package name and package version. For
example:

```

{ "Ref": "MyPackageVersion" }
```

Response:

```

package-name|version-name
```

For a stack named MyStack, a value that is similar to the following is returned:

```

MyStack-MyPackageVersion-AB1CDEFGHIJK
```

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

For example:

```

{ "GetAtt": "MySoftwarePackageVersion.PackageVersionArn" }
```

`ErrorReason`

Error reason for a package version failure during creation or update.

`PackageVersionArn`

The Amazon Resource Name (ARN) for the package.

`SbomValidationStatus`

Property description not available.

`Status`

The status of the package version. For more information, see [Package version lifecycle](../../../iot/latest/developerguide/preparing-to-use-software-package-catalog.md#package-version-lifecycle).

## Examples

The following example declares a package version and the values of its
attributes.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MyPackageVersion": {
      "Type": "AWS::IoT::SoftwarePackageVersion",
      "Properties": {
        "PackageName": "PackageName",
        "VersionName": "1.0.0",
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
  MyPackageVersion:
    Type: AWS::IoT::SoftwarePackageVersion
    Properties:
      PackageName: "PackageName"
      VersionName: "1.0.0"
      Description: "description"
      Tags:
        - Key: "TagKey"
          Value: "TagValue"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

PackageVersionArtifact

All content copied from https://docs.aws.amazon.com/.
