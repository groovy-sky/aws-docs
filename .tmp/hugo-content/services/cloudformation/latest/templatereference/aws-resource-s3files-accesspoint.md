This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Files::AccessPoint

The `AWS::S3Files::AccessPoint` resource specifies an access point for
an Amazon S3 Files file system. Access points provide application-specific access with
POSIX user identity and root directory enforcement.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Files::AccessPoint",
  "Properties" : {
      "ClientToken" : String,
      "FileSystemId" : String,
      "PosixUser" : PosixUser,
      "RootDirectory" : RootDirectory,
      "Tags" : [ AccessPointTag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::S3Files::AccessPoint
Properties:
  ClientToken: String
  FileSystemId: String
  PosixUser:
    PosixUser
  RootDirectory:
    RootDirectory
  Tags:
    - AccessPointTag

```

## Properties

`ClientToken`

A string of up to 64 ASCII characters that Amazon S3 Files uses to ensure
idempotent creation.

_Required_: No

_Type_: String

_Pattern_: `^(.+)$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FileSystemId`

The ID of the S3 Files file system that the access point provides access to.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PosixUser`

The POSIX identity configured for this access point.

_Required_: No

_Type_: [PosixUser](aws-properties-s3files-accesspoint-posixuser.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RootDirectory`

The root directory configuration for this access point.

_Required_: No

_Type_: [RootDirectory](aws-properties-s3files-accesspoint-rootdirectory.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [AccessPointTag](aws-properties-s3files-accesspoint-accesspointtag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the access point ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccessPointArn`

The Amazon Resource Name (ARN) of the access point.

`AccessPointId`

The ID of the access point.

`OwnerId`

The AWS account ID of the access point owner.

`Status`

The current status of the access point.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon S3 Files

AccessPointTag

All content copied from https://docs.aws.amazon.com/.
