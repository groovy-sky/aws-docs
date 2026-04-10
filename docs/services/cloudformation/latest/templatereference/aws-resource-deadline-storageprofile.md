This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::StorageProfile

Creates a storage profile that specifies the operating system, file type, and file
location of resources used on a farm.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Deadline::StorageProfile",
  "Properties" : {
      "DisplayName" : String,
      "FarmId" : String,
      "FileSystemLocations" : [ FileSystemLocation, ... ],
      "OsFamily" : String
    }
}

```

### YAML

```yaml

Type: AWS::Deadline::StorageProfile
Properties:
  DisplayName: String
  FarmId: String
  FileSystemLocations:
    - FileSystemLocation
  OsFamily: String

```

## Properties

`DisplayName`

The display name of the storage profile summary to update.

###### Important

This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FarmId`

The unique identifier of the farm that contains the storage profile.

_Required_: Yes

_Type_: String

_Pattern_: `^farm-[0-9a-f]{32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FileSystemLocations`

Operating system specific file system path to the storage location.

_Required_: No

_Type_: Array of [FileSystemLocation](aws-properties-deadline-storageprofile-filesystemlocation.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OsFamily`

The operating system (OS) family.

_Required_: Yes

_Type_: String

_Allowed values_: `WINDOWS | LINUX | MACOS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the storage profile.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`StorageProfileId`

The storage profile ID.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Deadline::QueueLimitAssociation

FileSystemLocation

All content copied from https://docs.aws.amazon.com/.
