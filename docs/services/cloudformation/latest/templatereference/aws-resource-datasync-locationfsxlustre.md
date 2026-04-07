This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationFSxLustre

The `AWS::DataSync::LocationFSxLustre` resource specifies an endpoint for
an Amazon FSx for Lustre file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataSync::LocationFSxLustre",
  "Properties" : {
      "FsxFilesystemArn" : String,
      "SecurityGroupArns" : [ String, ... ],
      "Subdirectory" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataSync::LocationFSxLustre
Properties:
  FsxFilesystemArn: String
  SecurityGroupArns:
    - String
  Subdirectory: String
  Tags:
    - Tag

```

## Properties

`FsxFilesystemArn`

Specifies the Amazon Resource Name (ARN) of the FSx for Lustre file
system.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):fsx:[a-z\-0-9]+:[0-9]{12}:file-system/fs-[0-9a-f]+$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupArns`

The ARNs of the security groups that are used to configure the FSx for Lustre file
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

Specifies a mount path for your FSx for Lustre file system. The path can include
subdirectories.

When the location is used as a source, DataSync reads data from the mount path.
When the location is used as a destination, DataSync writes data to the mount path.
If you don't include this parameter, DataSync uses the file system's root directory
( `/`).

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies labels that help you categorize, filter, and search for your AWS
resources. We recommend creating at least a name tag for your location.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationfsxlustre-tag.html)

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

The ARN of the specified FSx for Lustre file system location.

`LocationUri`

The URI of the specified FSx for Lustre file system location.

## Examples

### Specify an Amazon FSx for Lustre location for DataSync

The following examples specify an FSx for Lustre location for DataSync,
including the subdirectory `/MySubdirectory`.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Specifies an FSx for Lustre location for DataSync",
    "Resources": {
        "LocationFSxLustre": {
            "Type": "AWS::DataSync::LocationFSxLustre",
            "Properties": {
                "FsxFilesystemArn": "arn:aws:fsx:us-east-2:111222333444:file-system/fs-12345fsx",
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
Description: Specifies an FSx for Lustre location for DataSync
Resources:
  LocationFSxLustre:
    Type: AWS::DataSync::LocationFSxLustre
    Properties:
      FsxFilesystemArn: arn:aws:fsx:us-east-2:111222333444:file-system/fs-12345fsx
      SecurityGroupArns:
        - arn:aws:ec2:us-east-2:11122233344:security-group/sg-12345678901212345
      Subdirectory: /MySubdirectory

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
