This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Panorama::Package

###### Important

End of support notice: On May 31, 2026, AWS will end support for AWS Panorama. After May 31, 2026,
you will no longer be able to access the AWS Panorama console or AWS Panorama resources. For more information, see
[AWS Panorama end of support](../../../panorama/latest/dev/panorama-end-of-support.md).

Creates a package and storage location in an Amazon S3 access point.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Panorama::Package",
  "Properties" : {
      "PackageName" : String,
      "StorageLocation" : StorageLocation,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Panorama::Package
Properties:
  PackageName: String
  StorageLocation:
    StorageLocation
  Tags:
    - Tag

```

## Properties

`PackageName`

A name for the package.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-\_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageLocation`

A storage location.

_Required_: No

_Type_: [StorageLocation](aws-properties-panorama-package-storagelocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags for the package.

_Required_: No

_Type_: Array of [Tag](aws-properties-panorama-package-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The package's ARN.

`CreatedTime`

When the package was created.

`PackageId`

The package's ID.

`StorageLocation.BinaryPrefixLocation`

Property description not available.

`StorageLocation.Bucket`

Property description not available.

`StorageLocation.GeneratedPrefixLocation`

Property description not available.

`StorageLocation.ManifestPrefixLocation`

Property description not available.

`StorageLocation.RepoPrefixLocation`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

StorageLocation

All content copied from https://docs.aws.amazon.com/.
