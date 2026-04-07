This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationFSxWindows

The `AWS::DataSync::LocationFSxWindows` resource specifies an endpoint for
an Amazon FSx for Windows Server file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataSync::LocationFSxWindows",
  "Properties" : {
      "CmkSecretConfig" : CmkSecretConfig,
      "CustomSecretConfig" : CustomSecretConfig,
      "Domain" : String,
      "FsxFilesystemArn" : String,
      "Password" : String,
      "SecurityGroupArns" : [ String, ... ],
      "Subdirectory" : String,
      "Tags" : [ Tag, ... ],
      "User" : String
    }
}

```

### YAML

```yaml

Type: AWS::DataSync::LocationFSxWindows
Properties:
  CmkSecretConfig:
    CmkSecretConfig
  CustomSecretConfig:
    CustomSecretConfig
  Domain: String
  FsxFilesystemArn: String
  Password: String
  SecurityGroupArns:
    - String
  Subdirectory: String
  Tags:
    - Tag
  User: String

```

## Properties

`CmkSecretConfig`

Specifies configuration information for a DataSync-managed secret, such as an
authentication token, secret key, password, or Kerberos keytab that DataSync uses
to access a specific storage location, with a customer-managed AWS KMS key.

###### Note

You can use either `CmkSecretConfig` or `CustomSecretConfig` to
provide credentials for a `CreateLocation` request. Do not provide both
parameters for the same request.

_Required_: No

_Type_: [CmkSecretConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationfsxwindows-cmksecretconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomSecretConfig`

Specifies configuration information for a customer-managed Secrets Manager secret where
a storage location credentials is stored in Secrets Manager as plain text (for authentication
token, secret key, or password) or as binary (for Kerberos keytab). This configuration includes
the secret ARN, and the ARN for an IAM role that provides access to the secret.

###### Note

You can use either `CmkSecretConfig` or `CustomSecretConfig` to
provide credentials for a `CreateLocation` request. Do not provide both
parameters for the same request.

_Required_: No

_Type_: [CustomSecretConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationfsxwindows-customsecretconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domain`

Specifies the name of the Windows domain that the FSx for Windows File Server file system
belongs to.

If you have multiple Active Directory domains in your environment, configuring this
parameter makes sure that DataSync connects to the right file system.

_Required_: No

_Type_: String

_Pattern_: `^([A-Za-z0-9]+[A-Za-z0-9-.]*)*[A-Za-z0-9-]*[A-Za-z0-9]$`

_Maximum_: `253`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FsxFilesystemArn`

Specifies the Amazon Resource Name (ARN) for the FSx for Windows File Server file
system.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):fsx:[a-z\-0-9]*:[0-9]{12}:file-system/fs-.*$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Password`

Specifies the password of the user with the permissions to mount and access the files,
folders, and file metadata in your FSx for Windows File Server file system.

_Required_: No

_Type_: String

_Pattern_: `^.{0,104}$`

_Maximum_: `104`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupArns`

The Amazon Resource Names (ARNs) of the security groups that are used to configure the
FSx for Windows File Server file system.

_Pattern_:
`^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`

_Length constraints_: Maximum length of 128.

_Required_: Yes

_Type_: Array of String

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subdirectory`

Specifies a mount path for your file system using forward slashes. This is where DataSync reads or writes data (depending on if this is a source or destination
location).

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies labels that help you categorize, filter, and search for your AWS
resources. We recommend creating at least a name tag for your location.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationfsxwindows-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`User`

The user who has the permissions to access files and folders in the FSx for Windows
File Server file system.

For information about choosing a user name that ensures sufficient permissions to
files, folders, and metadata, see [user](https://docs.aws.amazon.com/datasync/latest/userguide/create-fsx-location.html#FSxWuser).

_Required_: Yes

_Type_: String

_Pattern_: `^[^\x5B\x5D\\/:;|=,+*?]{1,104}$`

_Maximum_: `104`

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

`CmkSecretConfig.SecretArn`

Property description not available.

`LocationArn`

The ARN of the specified FSx for Windows Server file system location.

`LocationUri`

The URI of the specified FSx for Windows Server file system location.

## Examples

### Specify an Amazon FSx for Windows File Server location for DataSync

The following examples specify an FSx for Windows File Server location for
DataSync, including the subdirectory `/MySubdirectory`, user
`Admin`, and password.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Specifies an FSx for Windows File Server location for DataSync",
    "Resources": {
        "LocationFSxWindows": {
            "Type": "AWS::DataSync::LocationFSxWindows",
            "Properties": {
                "FsxFilesystemArn": "arn:aws:fsx:us-east-2:111222333444:file-system/fs-12345fsx",
                "SecurityGroupArns": [
                    "arn:aws:ec2:us-east-2:11122233344:security-group/sg-12345678901212345"
                ],
                "Subdirectory": "/MySubdirectory",
                "User": "Admin",
                "Password": {
                    "Ref": "Password"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies an FSx for Windows File Server location for DataSync
Resources:
  LocationFSxWindows:
    Type: AWS::DataSync::LocationFSxWindows
    Properties:
      FsxFilesystemArn: arn:aws:fsx:us-east-2:111222333444:file-system/fs-12345fsx
      SecurityGroupArns:
        - arn:aws:ec2:us-east-2:11122233344:security-group/sg-12345678901212345
      Subdirectory: /MySubdirectory
      User: Admin
      Password: !Ref Password

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

CmkSecretConfig
