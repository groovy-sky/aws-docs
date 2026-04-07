This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Panorama::PackageVersion

###### Important

End of support notice: On May 31, 2026, AWS will end support for AWS Panorama. After May 31, 2026,
you will no longer be able to access the AWS Panorama console or AWS Panorama resources. For more information, see
[AWS Panorama end of support](https://docs.aws.amazon.com/panorama/latest/dev/panorama-end-of-support.html).

Registers a package version.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Panorama::PackageVersion",
  "Properties" : {
      "MarkLatest" : Boolean,
      "OwnerAccount" : String,
      "PackageId" : String,
      "PackageVersion" : String,
      "PatchVersion" : String,
      "UpdatedLatestPatchVersion" : String
    }
}

```

### YAML

```yaml

Type: AWS::Panorama::PackageVersion
Properties:
  MarkLatest: Boolean
  OwnerAccount: String
  PackageId: String
  PackageVersion: String
  PatchVersion: String
  UpdatedLatestPatchVersion: String

```

## Properties

`MarkLatest`

Whether to mark the new version as the latest version.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OwnerAccount`

An owner account.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-z\_]+$`

_Minimum_: `1`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PackageId`

A package ID.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-\_\/]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PackageVersion`

A package version.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9]+)\.([0-9]+)$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PatchVersion`

A patch version.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UpdatedLatestPatchVersion`

If the version was marked latest, the new version to maker as latest.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IsLatestPatch`

Whether the package version is the latest version.

`PackageArn`

The package version's ARN.

`PackageName`

The package version's name.

`RegisteredTime`

The package version's registered time.

`Status`

The package version's status.

`StatusDescription`

The package version's status description.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
