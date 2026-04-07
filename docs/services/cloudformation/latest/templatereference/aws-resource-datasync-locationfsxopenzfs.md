This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationFSxOpenZFS

The `AWS::DataSync::LocationFSxOpenZFS` resource specifies an endpoint for
an Amazon FSx for OpenZFS file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataSync::LocationFSxOpenZFS",
  "Properties" : {
      "FsxFilesystemArn" : String,
      "Protocol" : Protocol,
      "SecurityGroupArns" : [ String, ... ],
      "Subdirectory" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataSync::LocationFSxOpenZFS
Properties:
  FsxFilesystemArn: String
  Protocol:
    Protocol
  SecurityGroupArns:
    - String
  Subdirectory: String
  Tags:
    - Tag

```

## Properties

`FsxFilesystemArn`

The Amazon Resource Name (ARN) of the FSx for OpenZFS file system.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):fsx:[a-z\-0-9]+:[0-9]{12}:file-system/fs-[0-9a-f]+$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Protocol`

The type of protocol that AWS DataSync uses to access your file system.

_Required_: Yes

_Type_: [Protocol](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationfsxopenzfs-protocol.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupArns`

The ARNs of the security groups that are used to configure the FSx for OpenZFS file
system.

_Pattern_:
`^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`

_Length constraints_: Maximum length of 128.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `128 | 5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subdirectory`

A subdirectory in the location's path that must begin with `/fsx`. DataSync uses this subdirectory to read or write data (depending on whether the file
system is a source or destination location).

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The key-value pair that represents a tag that you want to add to the resource. The value
can be an empty string. This value helps you manage, filter, and search for your resources. We
recommend that you create a name tag for your location.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationfsxopenzfs-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the location resource ARN. For example:

`arn:aws:datasync:us-east-2:111222333444:location/loc-07db7abfc326c50s3`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`LocationArn`

The ARN of the specified FSx for OpenZFS file system location.

`LocationUri`

The URI of the specified FSx for OpenZFS file system location.

## Examples

### Specify an Amazon FSx for OpenZFS location for DataSync

The following examples specify an FSx for OpenZFS location for
DataSync.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Specifies an FSx for OpenZFS location for DataSync",
    "Resources": {
        "LocationFSxOpenZFS": {
            "Type": "AWS::DataSync::LocationFSxOpenZFS",
            "Properties": {
                "FsxFilesystemArn": "arn:aws:fsx:us-east-2:111222333444:file-system/fs-12345fsx",
                "Protocol": {
                    "NFS": {
                        "MountOptions": {
                            "Version": "NFS4_1"
                        }
                    }
                },
                "SecurityGroupArns": [
                    "arn:aws:ec2:us-east-2:11122233344:security-group/sg-12345678901212345"
                ],
                "Subdirectory": "/MySubdirectory"
            }
        }
    }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies an FSx for OpenZFS location for DataSync
Resources:
  LocationFSxOpenZFS:
    Type: AWS::DataSync::LocationFSxOpenZFS
    Properties:
      FsxFilesystemArn: arn:aws:fsx:us-east-2:111222333444:file-system/fs-12345fsx
      Protocol:
        NFS:
          MountOptions:
            Version: NFS4_1
      SecurityGroupArns:
        - arn:aws:ec2:us-east-2:11122233344:security-group/sg-12345678901212345
      Subdirectory: /MySubdirectory

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

MountOptions
