This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::DataRepositoryAssociation

Creates an Amazon FSx for Lustre data repository association (DRA). A data
repository association is a link between a directory on the file system and
an Amazon S3 bucket or prefix. You can have a maximum of 8 data repository
associations on a file system. Data repository associations are supported on
all FSx for Lustre 2.12 and newer file systems, excluding `scratch_1` deployment type.

Each data repository association must have a unique Amazon FSx file
system directory and a unique S3 bucket or prefix associated with it. You
can configure a data repository association for automatic import only,
for automatic export only, or for both. To learn more about linking a
data repository to your file system, see
[Linking your file system to an S3 bucket](https://docs.aws.amazon.com/fsx/latest/LustreGuide/create-dra-linked-data-repo.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FSx::DataRepositoryAssociation",
  "Properties" : {
      "BatchImportMetaDataOnCreate" : Boolean,
      "DataRepositoryPath" : String,
      "FileSystemId" : String,
      "FileSystemPath" : String,
      "ImportedFileChunkSize" : Integer,
      "S3" : S3,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::FSx::DataRepositoryAssociation
Properties:
  BatchImportMetaDataOnCreate: Boolean
  DataRepositoryPath: String
  FileSystemId: String
  FileSystemPath: String
  ImportedFileChunkSize: Integer
  S3:
    S3
  Tags:
    - Tag

```

## Properties

`BatchImportMetaDataOnCreate`

A boolean flag indicating whether an import data repository task to import
metadata should run after the data repository association is created. The
task runs if this flag is set to `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataRepositoryPath`

The path to the Amazon S3 data repository that will be linked to the file
system. The path can be an S3 bucket or prefix in the format
`s3://myBucket/myPrefix/`. This path specifies where in the S3
data repository files will be imported from or exported to.

_Required_: Yes

_Type_: String

_Pattern_: `^[^\u0000\u0085\u2028\u2029\r\n]{3,4357}$`

_Minimum_: `3`

_Maximum_: `4357`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FileSystemId`

The ID of the file system on which the data repository association is configured.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FileSystemPath`

A path on the Amazon FSx for Lustre file system that points to a high-level directory (such
as `/ns1/`) or subdirectory (such as `/ns1/subdir/`)
that will be mapped 1-1 with `DataRepositoryPath`.
The leading forward slash in the name is required. Two data repository
associations cannot have overlapping file system paths. For example, if
a data repository is associated with file system path `/ns1/`,
then you cannot link another data repository with file system
path `/ns1/ns2`.

This path specifies where in your file system files will be exported
from or imported to. This file system directory can be linked to only one
Amazon S3 bucket, and no other S3 bucket can be linked to the directory.

###### Note

If you specify only a forward slash ( `/`) as the file system
path, you can link only one data repository to the file system. You can only specify
"/" as the file system path for the first data repository associated with a file system.

_Required_: Yes

_Type_: String

_Pattern_: `^[^\u0000\u0085\u2028\u2029\r\n]{1,4096}$`

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImportedFileChunkSize`

For files imported from a data repository, this value determines the stripe count and
maximum amount of data per file (in MiB) stored on a single physical disk. The maximum
number of disks that a single file can be striped across is limited by the total number
of disks that make up the file system or cache.

The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500
GiB). Amazon S3 objects have a maximum size of 5 TB.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `512000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

The configuration for an Amazon S3 data repository linked to an
Amazon FSx Lustre file system with a data repository association.
The configuration defines which file events (new, changed, or
deleted files or directories) are automatically imported from
the linked data repository to the file system or automatically
exported from the file system to the data repository.

_Required_: No

_Type_: [S3](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-datarepositoryassociation-s3.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of `Tag` values, with a maximum of 50 elements.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-datarepositoryassociation-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the data repository association resource ID. For example:

`{"Ref":"data_repository_association_logical_id"}`

Returns `dra-0123456789abcdef6`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssociationId`

Returns the data repository association's system generated Association ID.

Example: `dra-abcdef0123456789d`

`ResourceARN`

Returns the data repository association's Amazon Resource Name (ARN).

Example: `arn:aws:fsx:us-east-1:111122223333:association/fs-abc012345def6789a/dra-abcdef0123456789b`

## Examples

### Create an Amazon FSx for Lustre data repository association

The following examples create a data repository association (DRA) that
links an Amazon FSx for Lustre file system (with the `Persistent_2`
deployment type) to an Amazon S3 bucket.

#### JSON

```json

{
    "Description": "FSx DRA with all properties.",
    "Parameters": {
        "FSId": {
            "Type": "String"
        },
        "DRAIdExportName": {
            "Type": "String"
        },
        "FileSystemPath": {
            "Type": "String"
        },
        "ImportedFileChunkSize": {
            "Type": "String"
        }
    },
    "Resources": {
        "TestDRA": {
            "Type": "AWS::FSx::DataRepositoryAssociation",
            "Properties": {
                "FileSystemId": {
                    "Ref": "FSId"
                },
                "FileSystemPath": {
                    "Ref": "FileSystemPath"
                },
                "DataRepositoryPath": "s3://example-bucket",
                "BatchImportMetaDataOnCreate": true,
                "ImportedFileChunkSize": {
                    "Ref": "ImportedFileChunkSize"
                },
                "S3": {
                    "AutoImportPolicy": {
                        "Events": [
                            "NEW",
                            "CHANGED",
                            "DELETED"
                        ]
                    },
                    "AutoExportPolicy": {
                        "Events": [
                            "NEW",
                            "CHANGED",
                            "DELETED"
                        ]
                    }
                },
                "Tags": [
                    {
                        "Key": "Location",
                        "Value": "Boston"
                    }
                ]
            }
        }
    },
    "Outputs": {
        "DRAId": {
            "Value": {
                "Ref": "TestDRA"
            },
            "Export": {
                "Name": {
                    "Ref": "DRAIdExportName"
                }
            }
        }
    }
```

#### YAML

```yaml

Description: "FSx DRA with all properties."
Parameters:
  FSId:
    Type: String
  DRAIdExportName:
    Type: String
  FileSystemPath:
    Type: String
  ImportedFileChunkSize:
    Type: String
Resources:
  TestDRA:
    Type: 'AWS::FSx::DataRepositoryAssociation'
    Properties:
      FileSystemId: !Ref FSId
      FileSystemPath: !Ref FileSystemPath
      DataRepositoryPath: "s3://example-bucket"
      BatchImportMetaDataOnCreate: true
      ImportedFileChunkSize: !Ref ImportedFileChunkSize
      S3:
        AutoImportPolicy:
          Events:
            - NEW
            - CHANGED
            - DELETED
        AutoExportPolicy:
          Events:
            - NEW
            - CHANGED
            - DELETED
      Tags:
        - Key: Location
          Value: Boston
Outputs:
  DRAId:
    Value: !Ref TestDRA
    Export:
      Name: !Ref DRAIdExportName
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon FSx

AutoExportPolicy
