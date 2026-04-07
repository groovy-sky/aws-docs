This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EFS::AccessPoint

The `AWS::EFS::AccessPoint` resource creates an EFS access point.
An access point is an application-specific view into an EFS file system that applies an operating system user and
group, and a file system path, to any file system request made through the access point. The operating system
user and group override any identity information provided by the NFS client. The file system path is exposed as
the access point's root directory. Applications using the access point can only access data in its own directory and below. To learn more, see
[Mounting a file system using EFS access points](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html).

This operation requires permissions for the `elasticfilesystem:CreateAccessPoint` action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EFS::AccessPoint",
  "Properties" : {
      "AccessPointTags" : [ AccessPointTag, ... ],
      "ClientToken" : String,
      "FileSystemId" : String,
      "PosixUser" : PosixUser,
      "RootDirectory" : RootDirectory
    }
}

```

### YAML

```yaml

Type: AWS::EFS::AccessPoint
Properties:
  AccessPointTags:
    - AccessPointTag
  ClientToken: String
  FileSystemId: String
  PosixUser:
    PosixUser
  RootDirectory:
    RootDirectory

```

## Properties

`AccessPointTags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [AccessPointTag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-efs-accesspoint-accesspointtag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientToken`

The opaque string specified in the request to ensure idempotent creation.

_Required_: No

_Type_: String

_Pattern_: `.+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FileSystemId`

The ID of the EFS file system that the access point applies to. Accepts only the ID format for input when specifying a file system, for example `fs-0123456789abcedf2`.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws[-a-z]*:elasticfilesystem:[0-9a-z-:]+:file-system/fs-[0-9a-f]{8,40}|fs-[0-9a-f]{8,40})$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PosixUser`

The full POSIX identity, including the user ID, group ID, and secondary group IDs on the access point that is used for all file operations by
NFS clients using the access point.

_Required_: No

_Type_: [PosixUser](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-efs-accesspoint-posixuser.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RootDirectory`

The directory on the EFS file system that the access point exposes as the root
directory to NFS clients using the access point.

_Required_: No

_Type_: [RootDirectory](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-efs-accesspoint-rootdirectory.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the AccessPoint ID. For example:

`{"Ref":"access_point-logical_id"}` returns

`fsap-0123456789abcdef0`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccessPointId`

The ID of the EFS access point.

`Arn`

The Amazon Resource Name (ARN) of the access point.

## Examples

### Declare an Access Point for an EFS File System

The following example declares an access point that is associated with an EFS file system. For information about mounting file systems on EC2 instances, see
[Mounting File Systems](https://docs.aws.amazon.com/efs/latest/ug/mounting-fs.html) in the
_EFS User Guide_.

#### JSON

```json

"AccessPointResource": {
            "Type": "AWS::EFS::AccessPoint",
            "Properties": {
                "FileSystemId": {
                    "Ref": "FileSystemResource"
                },
                "PosixUser": {
                    "Uid": "13234",
                    "Gid": "1322",
                    "SecondaryGids": [
                        "1344",
                        "1452"
                    ]
                },
                "RootDirectory": {
                    "CreationInfo": {
                        "OwnerGid": "708798",
                        "OwnerUid": "7987987",
                        "Permissions": "0755"
                    },
                    "Path": "/testcfn/abc"
                }
            }
        }
}
```

#### YAML

```yaml

AccessPointResource:
    Type: 'AWS::EFS::AccessPoint'
    Properties:
      FileSystemId: !Ref FileSystemResource
      PosixUser:
        Uid: "13234"
        Gid: "1322"
        SecondaryGids:
          - "1344"
          - "1452"
      RootDirectory:
        CreationInfo:
          OwnerGid: "708798"
          OwnerUid: "7987987"
          Permissions: "0755"
        Path: "/testcfn/abc"
```

## See also

- [Amazon EFS: How it works](https://docs.aws.amazon.com/efs/latest/ug/how-it-works.html).

- [Working with Amazon EFS access points](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html) in the _Amazon EFS User Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Elastic File System

AccessPointTag
