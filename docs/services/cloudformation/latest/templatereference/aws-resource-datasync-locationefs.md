This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationEFS

The `AWS::DataSync::LocationEFS` resource creates an endpoint for an
Amazon EFS file system. AWS DataSync can access this endpoint as a
source or destination location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataSync::LocationEFS",
  "Properties" : {
      "AccessPointArn" : String,
      "Ec2Config" : Ec2Config,
      "EfsFilesystemArn" : String,
      "FileSystemAccessRoleArn" : String,
      "InTransitEncryption" : String,
      "Subdirectory" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataSync::LocationEFS
Properties:
  AccessPointArn: String
  Ec2Config:
    Ec2Config
  EfsFilesystemArn: String
  FileSystemAccessRoleArn: String
  InTransitEncryption: String
  Subdirectory: String
  Tags:
    - Tag

```

## Properties

`AccessPointArn`

Specifies the Amazon Resource Name (ARN) of the access point that DataSync uses
to mount your Amazon EFS file system.

For more information, see [Accessing\
restricted file systems](https://docs.aws.amazon.com/datasync/latest/userguide/create-efs-location.html#create-efs-location-iam).

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):elasticfilesystem:[a-z\-0-9]+:[0-9]{12}:access-point/fsap-[0-9a-f]{8,40}$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ec2Config`

Specifies the subnet and security groups DataSync uses to connect to one of
your Amazon EFS file system's [mount targets](https://docs.aws.amazon.com/efs/latest/ug/accessing-fs.html).

_Required_: Yes

_Type_: [Ec2Config](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationefs-ec2config.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EfsFilesystemArn`

Specifies the ARN for your Amazon EFS file system.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):elasticfilesystem:[a-z\-0-9]*:[0-9]{12}:file-system/fs-.*$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FileSystemAccessRoleArn`

Specifies an AWS Identity and Access Management (IAM) role that allows DataSync to access your Amazon EFS file system.

For information on creating this role, see [Creating a DataSync\
IAM role for file system access](https://docs.aws.amazon.com/datasync/latest/userguide/create-efs-location.html#create-efs-location-iam-role).

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):iam::[0-9]{12}:role/.*$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InTransitEncryption`

Specifies whether you want DataSync to use Transport Layer Security (TLS) 1.2
encryption when it transfers data to or from your Amazon EFS file system.

If you specify an access point using `AccessPointArn` or an IAM
role using `FileSystemAccessRoleArn`, you must set this parameter to
`TLS1_2`.

_Required_: No

_Type_: String

_Allowed values_: `NONE | TLS1_2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subdirectory`

Specifies a mount path for your Amazon EFS file system. This is where DataSync reads or writes data on your file system (depending on if this is a source or
destination location).

By default, DataSync uses the root directory (or [access point](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html) if you provide one by using
`AccessPointArn`). You can also include subdirectories using forward slashes (for
example, `/path/to/folder`).

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies the key-value pair that represents a tag that you want to add to the
resource. The value can be an empty string. This value helps you manage, filter, and search
for your resources. We recommend that you create a name tag for your location.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationefs-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the location resource ARN. For example:

`arn:aws:datasync:us-east-2:111222333444:location/loc-07db7abfc326c50s3`

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified
attribute of this type. The following are the available attributes and sample return
values.

For more information about using the `Fn::GetAtt` intrinsic function, see
[Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

`LocationArn`

The Amazon Resource Name (ARN) of the Amazon EFS file system.

`LocationUri`

The URI of the Amazon EFS file system.

## Examples

- [Creating an Amazon EFS location](#aws-resource-datasync-locationefs--examples--Creating_an_Amazon_EFS_location)

- [Creating an Amazon EFS location with a higher level of security](#aws-resource-datasync-locationefs--examples--Creating_an_Amazon_EFS_location_with_a_higher_level_of_security)

### Creating an Amazon EFS location

The following example creates a DataSync location for an Amazon EFS file system.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Specifies a DataSync location for an Amazon EFS file system.",
    "Resources": {
        "LocationEFS": {
            "Type": "AWS::DataSync::LocationEFS",
            "Properties": {
                "Ec2Config": {
                    "SecurityGroupArns": [
                        "arn:aws:ec2:us-east-2:11122233344:security-group/sg-1234567890abcdef2"
                    ],
                    "SubnetArn": "arn:aws:ec2:us-east-2:11122233344:subnet/subnet-1234567890abcdef1"
                },
                "EfsFilesystemArn": "arn:aws:elasticfilesystem:us-east-2:111222333444:file-system/fs-021345abcdef6789",
                "Subdirectory": "/mount/path"
            }
        }
    }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies a DataSync location for an Amazon EFS file system.
Resources:
  LocationEFS:
    Type: AWS::DataSync::LocationEFS
      Properties:
        Ec2Config:
          SecurityGroupArns:
            - arn:aws:ec2:us-east-2:11122233344:security-group/sg-1234567890abcdef2
          SubnetArn: arn:aws:ec2:us-east-2:11122233344:subnet/subnet-1234567890abcdef1
        EfsFilesystemArn: arn:aws:elasticfilesystem:us-east-2:111222333444:file-system/fs-021345abcdef6789
        Subdirectory: /mount/path

```

### Creating an Amazon EFS location with a higher level of security

The following example creates a DataSync location for an Amazon EFS file system that's configured for restricted access.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Specifies a DataSync location for an Amazon EFS file system configured for restricted access.",
    "Resources": {
        "LocationEFS": {
            "Type": "AWS::DataSync::LocationEFS",
            "Properties": {
                "AccessPointArn": "arn:aws:elasticfilesystem:us-east-2:111222333444:access-point/fsap-1234567890abcdef0",
                "Ec2Config": {
                    "SecurityGroupArns": [
                        "arn:aws:ec2:us-east-2:11122233344:security-group/sg-1234567890abcdef2"
                    ],
                    "SubnetArn": "arn:aws:ec2:us-east-2:11122233344:subnet/subnet-1234567890abcdef1"
                },
                "EfsFilesystemArn": "arn:aws:elasticfilesystem:us-east-2:111222333444:file-system/fs-021345abcdef6789",
                "FileSystemAccessRoleArn": "arn:aws:iam::111222333444:role/AllowDataSyncAccess",
                "InTransitEncryption": "TLS1_2"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies a DataSync location for an Amazon EFS file system configured for restricted access.
Resources:
  LocationEFS:
    Type: AWS::DataSync::LocationEFS
      Properties:
        AccessPointArn: arn:aws:elasticfilesystem:us-east-2:111222333444:access-point/fsap-1234567890abcdef0
        Ec2Config:
          SecurityGroupArns:
            - arn:aws:ec2:us-east-2:11122233344:security-group/sg-1234567890abcdef2
          SubnetArn: arn:aws:ec2:us-east-2:11122233344:subnet/subnet-1234567890abcdef1
        EfsFilesystemArn: arn:aws:elasticfilesystem:us-east-2:111222333444:file-system/fs-021345abcdef6789
        FileSystemAccessRoleArn: arn:aws:iam::111222333444:role/AllowDataSyncAccess
        InTransitEncryption: TLS1_2
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Ec2Config
