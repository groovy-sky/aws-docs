This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationFSxONTAP

The `AWS::DataSync::LocationFSxONTAP` resource creates an endpoint for an
Amazon FSx for NetApp ONTAP file system. AWS DataSync can access this
endpoint as a source or destination location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataSync::LocationFSxONTAP",
  "Properties" : {
      "Protocol" : Protocol,
      "SecurityGroupArns" : [ String, ... ],
      "StorageVirtualMachineArn" : String,
      "Subdirectory" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataSync::LocationFSxONTAP
Properties:
  Protocol:
    Protocol
  SecurityGroupArns:
    - String
  StorageVirtualMachineArn: String
  Subdirectory: String
  Tags:
    - Tag

```

## Properties

`Protocol`

Specifies the data transfer protocol that DataSync uses to access your
Amazon FSx file system.

_Required_: No

_Type_: [Protocol](aws-properties-datasync-locationfsxontap-protocol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupArns`

Specifies the Amazon Resource Names (ARNs) of the security groups that DataSync can use to access your FSx for ONTAP file system. You must configure the
security groups to allow outbound traffic on the following ports (depending on the
protocol that you're using):

- **Network File System (NFS)**: TCP ports 111,
635, and 2049

- **Server Message Block (SMB)**: TCP port
445

Your file system's security groups must also allow inbound traffic on the same
port.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `128 | 5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageVirtualMachineArn`

Specifies the ARN of the storage virtual machine (SVM) in your file system where you want
to copy data to or from.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):fsx:[a-z\-0-9]+:[0-9]{12}:storage-virtual-machine/fs-[0-9a-f]+/svm-[0-9a-f]{17,}$`

_Maximum_: `162`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subdirectory`

Specifies a path to the file share in the SVM where you want to transfer data to or
from.

You can specify a junction path (also known as a mount point), qtree path (for NFS file
shares), or share name (for SMB file shares). For example, your mount path might be
`/vol1`, `/vol1/tree1`, or `/share1`.

###### Note

Don't specify a junction path in the SVM's root volume. For more information, see [Managing FSx for ONTAP storage virtual machines](../../../fsx/latest/ontapguide/managing-svms.md) in the _Amazon FSx for NetApp ONTAP User Guide_.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies labels that help you categorize, filter, and search for your AWS
resources. We recommend creating at least a name tag for your location.

_Required_: No

_Type_: Array of [Tag](aws-properties-datasync-locationfsxontap-tag.md)

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

`FsxFilesystemArn`

The ARN of the FSx for ONTAP file system in the specified location.

`LocationArn`

The ARN of the specified location.

`LocationUri`

The URI of the specified location.

`Protocol.SMB.CmkSecretConfig.SecretArn`

Property description not available.

## Examples

- [Creating an FSx for ONTAP location with NFS access](#aws-resource-datasync-locationfsxontap--examples--Creating_an_FSx_for_ONTAP_location_with_NFS_access)

- [Creating an FSx for ONTAP location with SMB access](#aws-resource-datasync-locationfsxontap--examples--Creating_an_FSx_for_ONTAP_location_with_SMB_access)

### Creating an FSx for ONTAP location with NFS access

The following example creates a location for an FSx for ONTAP file
system that DataSync can access using NFS.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Specifies a location for an Amazon FSx for ONTAP file system that DataSync can access using NFS.",
    "Resources": {
        "LocationFSxONTAP": {
            "Type": "AWS::DataSync::LocationFSxONTAP",
            "Properties": {
                "Protocol": {
                    "NFS": {
                        "MountOptions": {
                            "Version": "NFS3"
                        }
                    }
                },
                "SecurityGroupArns": [
                    "arn:aws:ec2:us-east-2:11122233344:security-group/sg-1234567890abcdef2"
                ],
                "StorageVirtualMachineArn": "arn:aws:fsx:us-east-1:11122233344:storage-virtual-machine/fs-abcdef01234567890/svm-021345abcdef6789",
                "Subdirectory": "/vol1"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies a location for an Amazon FSx for ONTAP file system that DataSync can access using NFS.
Resources:
  LocationFSxONTAP:
    Type: AWS::DataSync::LocationFSxONTAP
      Properties:
        Protocol:
          NFS:
            MountOptions:
              Version: NFS3
        SecurityGroupArns:
          - arn:aws:ec2:us-east-2:11122233344:security-group/sg-1234567890abcdef2
        StorageVirtualMachineArn: arn:aws:fsx:us-east-1:11122233344:storage-virtual-machine/fs-abcdef01234567890/svm-021345abcdef6789
        Subdirectory: /vol1
```

### Creating an FSx for ONTAP location with SMB access

The following example creates a location for an FSx for ONTAP file
system that DataSync can access using SMB.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Specifies a location for an Amazon FSx for ONTAP file system that DataSync can access using SMB.",
    "Resources": {
        "LocationFSxONTAP": {
            "Type": "AWS::DataSync::LocationFSxONTAP",
            "Properties": {
                "Protocol": {
                    "SMB": {
                        "Domain": "example.company",
                        "MountOptions": {
                            "Version": "AUTOMATIC"
                        },
                        "Password": "examplepassword",
                        "User": "exampleusername"
                    }
                },
                "SecurityGroupArns": [
                    "arn:aws:ec2:us-east-2:11122233344:security-group/sg-1234567890abcdef2"
                ],
                "StorageVirtualMachineArn": "arn:aws:fsx:us-east-1:11122233344:storage-virtual-machine/fs-abcdef01234567890/svm-021345abcdef6789",
                "Subdirectory": "/vol1"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies a location for an Amazon FSx for ONTAP file system that DataSync can access using SMB.
Resources:
  LocationFSxONTAP:
    Type: AWS::DataSync::LocationFSxONTAP
      Properties:
        Protocol:
          SMB:
            Domain: example.company
            MountOptions:
              Version: AUTOMATIC
            Password: examplepassword
            User: exampleusername
        SecurityGroupArns:
          - arn:aws:ec2:us-east-2:11122233344:security-group/sg-1234567890abcdef2
        StorageVirtualMachineArn: arn:aws:fsx:us-east-1:11122233344:storage-virtual-machine/fs-abcdef01234567890/svm-021345abcdef6789
        Subdirectory: /vol1
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

CmkSecretConfig

All content copied from https://docs.aws.amazon.com/.
