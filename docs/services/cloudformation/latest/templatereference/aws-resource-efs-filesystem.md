This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EFS::FileSystem

The `AWS::EFS::FileSystem` resource creates a new, empty file system in
Amazon Elastic File System (Amazon EFS). You must create a mount target
( [AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html)) to mount your EFS file system on an
Amazon EC2 or other AWS cloud compute resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EFS::FileSystem",
  "Properties" : {
      "AvailabilityZoneName" : String,
      "BackupPolicy" : BackupPolicy,
      "BypassPolicyLockoutSafetyCheck" : Boolean,
      "Encrypted" : Boolean,
      "FileSystemPolicy" : Json,
      "FileSystemProtection" : FileSystemProtection,
      "FileSystemTags" : [ ElasticFileSystemTag, ... ],
      "KmsKeyId" : String,
      "LifecyclePolicies" : [ LifecyclePolicy, ... ],
      "PerformanceMode" : String,
      "ProvisionedThroughputInMibps" : Number,
      "ReplicationConfiguration" : ReplicationConfiguration,
      "ThroughputMode" : String
    }
}

```

### YAML

```yaml

Type: AWS::EFS::FileSystem
Properties:
  AvailabilityZoneName: String
  BackupPolicy:
    BackupPolicy
  BypassPolicyLockoutSafetyCheck: Boolean
  Encrypted: Boolean
  FileSystemPolicy: Json
  FileSystemProtection:
    FileSystemProtection
  FileSystemTags:
    - ElasticFileSystemTag
  KmsKeyId: String
  LifecyclePolicies:
    - LifecyclePolicy
  PerformanceMode: String
  ProvisionedThroughputInMibps: Number
  ReplicationConfiguration:
    ReplicationConfiguration
  ThroughputMode: String

```

## Properties

`AvailabilityZoneName`

For One Zone file systems, specify the AWS
Availability Zone in which to create the file system. Use the format `us-east-1a` to
specify the Availability Zone. For more information about One Zone file systems, see
[EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/availability-durability.html#file-system-type) in the _Amazon EFS User Guide_.

###### Note

One Zone file systems are not available in all Availability Zones in AWS Regions where Amazon EFS is available.

_Required_: No

_Type_: String

_Pattern_: `.+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BackupPolicy`

Use the `BackupPolicy` to turn automatic backups on or off for the file system.

_Required_: No

_Type_: [BackupPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-efs-filesystem-backuppolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BypassPolicyLockoutSafetyCheck`

(Optional) A boolean that specifies whether or not to bypass the `FileSystemPolicy` lockout safety check. The lockout safety check
determines whether the policy in the request will lock out, or prevent, the IAM principal that is making the request from making future `PutFileSystemPolicy` requests on this file system.
Set `BypassPolicyLockoutSafetyCheck` to `True` only when you intend to prevent
the IAM principal that is making the request from making subsequent `PutFileSystemPolicy` requests on this file system.
The default value is `False`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Encrypted`

A Boolean value that, if true, creates an encrypted file system. When creating an
encrypted file system, you have the option of specifying a KmsKeyId for an existing AWS KMS key. If you don't specify a KMS key, then the default KMS key for
Amazon EFS, `/aws/elasticfilesystem`, is used to protect the encrypted file system.

_Required_: Conditional

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FileSystemPolicy`

The `FileSystemPolicy` for the EFS file system. A file system policy is an IAM resource policy used to control NFS access to an EFS file system.
For more information, see [Using IAM to control NFS access to Amazon EFS](https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html)
in the _Amazon EFS User Guide_.

_Required_: No

_Type_: Json

_Pattern_: `[\s\S]+`

_Minimum_: `1`

_Maximum_: `20000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileSystemProtection`

Describes the protection on the file system.

_Required_: No

_Type_: [FileSystemProtection](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-efs-filesystem-filesystemprotection.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileSystemTags`

Use to create one or more tags associated with the file system. Each
tag is a user-defined key-value pair. Name your file system on creation by including a
`"Key":"Name","Value":"{value}"` key-value pair. Each key must be unique. For more
information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
in the _AWS General Reference Guide_.

_Required_: No

_Type_: Array of [ElasticFileSystemTag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-efs-filesystem-elasticfilesystemtag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The ID of the AWS KMS key to be used to protect the encrypted file system. This
parameter is only required if you want to use a nondefault KMS key. If this parameter is not
specified, the default KMS key for Amazon EFS is used. This ID can be in one of the following
formats:

- Key ID - A unique identifier of the key, for example
`1234abcd-12ab-34cd-56ef-1234567890ab`.

- ARN - An Amazon Resource Name (ARN) for the key, for example
`arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`.

- Key alias - A previously created display name for a key, for example
`alias/projectKey1`.

- Key alias ARN - An ARN for a key alias, for example
`arn:aws:kms:us-west-2:444455556666:alias/projectKey1`.

If `KmsKeyId` is specified, the `Encrypted` parameter must be set to true.

_Required_: No

_Type_: String

_Pattern_: `^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|mrk-[0-9a-f]{32}|alias/[a-zA-Z0-9/_-]+|(arn:aws[-a-z]*:kms:[a-z0-9-]+:\d{12}:((key/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})|(key/mrk-[0-9a-f]{32})|(alias/[a-zA-Z0-9/_-]+))))$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LifecyclePolicies`

An array of `LifecyclePolicy` objects that define the file system's
`LifecycleConfiguration` object. A `LifecycleConfiguration` object
informs Lifecycle management of the following:

- When to move files in the file system from primary storage to IA storage.

- When to move files in the file system from primary storage or IA storage
to Archive storage.

- When to move files that are in IA or Archive storage to primary storage.

###### Note

Amazon EFS requires that each `LifecyclePolicy` object have only a single transition.
This means that in a request body, `LifecyclePolicies` needs to be structured as an array of
`LifecyclePolicy` objects, one object for each transition, `TransitionToIA`, `TransitionToArchive` `TransitionToPrimaryStorageClass`.
See the example requests in the following section for more information.

_Required_: No

_Type_: Array of [LifecyclePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-efs-filesystem-lifecyclepolicy.html)

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerformanceMode`

The performance mode of the file system. We recommend `generalPurpose`
performance mode for all file systems. File systems using the `maxIO` performance
mode can scale to higher levels of aggregate throughput and operations per second with a
tradeoff of slightly higher latencies for most file operations. The performance mode
can't be changed after the file system has been created. The `maxIO` mode is
not supported on One Zone file systems.

###### Important

Due to the higher per-operation latencies with Max I/O, we recommend using General Purpose performance mode for all file systems.

Default is `generalPurpose`.

_Required_: No

_Type_: String

_Allowed values_: `generalPurpose | maxIO`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProvisionedThroughputInMibps`

The throughput, measured in mebibytes per second (MiBps), that you want to provision for a
file system that you're creating. Required if `ThroughputMode` is set to
`provisioned`. Valid values are 1-3414 MiBps, with the upper limit depending on
Region. To increase this limit, contact Support. For more information, see [Amazon EFS quotas\
that you can increase](https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the _Amazon EFS User_
_Guide_.

_Required_: Conditional

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationConfiguration`

Describes the replication configuration for a specific file system.

_Required_: No

_Type_: [ReplicationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-efs-filesystem-replicationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThroughputMode`

Specifies the throughput mode for the file system. The mode can be `bursting`,
`provisioned`, or `elastic`. If you set `ThroughputMode` to
`provisioned`, you must also set a value for
`ProvisionedThroughputInMibps`. After you create the file system, you can
decrease your file system's Provisioned throughput or change between the
throughput modes, with certain time restrictions. For more information, see [Specifying\
throughput with provisioned mode](https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput) in the _Amazon EFS User_
_Guide_.

Default is `bursting`.

_Required_: No

_Type_: String

_Allowed values_: `bursting | provisioned | elastic`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the FileSystem ID. For example:

`{"Ref":"logical_file_system_id"}`

returns `fs-0123456789abcdef2`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the EFS file system.

Example: `arn:aws:elasticfilesystem:us-west-2:1111333322228888:file-system/fs-0123456789abcdef8`

`FileSystemId`

The ID of the EFS file system. For example: `fs-abcdef0123456789a`

## Examples

- [Create an encrypted EFS file system using EFS Standard storage classes](#aws-resource-efs-filesystem--examples--Create_an_encrypted_EFS_file_system_using_EFS_Standard_storage_classes)

- [Create a One Zone file system](#aws-resource-efs-filesystem--examples--Create_a_One_Zone_file_system)

### Create an encrypted EFS file system using EFS Standard storage classes

The following example declares an Amazon EFS file system with the following attributes:

- Uses EFS Standard storage classes.

- maxIO performance mode.

- Lifecycle management and Intelligent Tiering enabled.

- Encrypted at rest.

- Automatic daily backups are enabled.

- File system policy granting read-only access to the EfsReadOnly IAM role.

- File system access:

- Mount targets in three Availability Zones.

- An access point providing an application-specific entry point to the file system.

#### JSON

```json

"{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "MountTargetVPC": {
            "Type": "AWS::EC2::VPC",
            "Properties": {
                "CidrBlock": "172.31.0.0/16"
            }
        },
        "MountTargetSubnetOne": {
            "Type": "AWS::EC2::Subnet",
            "Properties": {
                "CidrBlock": "172.31.1.0/24",
                "VpcId": {
                    "Ref": "MountTargetVPC"
                },
                "AvailabilityZone": "us-east-1a"
            }
        },
        "MountTargetSubnetTwo": {
            "Type": "AWS::EC2::Subnet",
            "Properties": {
                "CidrBlock": "172.31.2.0/24",
                "VpcId": {
                    "Ref": "MountTargetVPC"
                },
                "AvailabilityZone": "us-east-1b"
            }
        },
        "MountTargetSubnetThree": {
            "Type": "AWS::EC2::Subnet",
            "Properties": {
                "CidrBlock": "172.31.3.0/24",
                "VpcId": {
                    "Ref": "MountTargetVPC"
                },
                "AvailabilityZone": "us-east-1c"
            }
        },
       "FileSystemResource": {
            "Type": "AWS::EFS::FileSystem",
            "Properties": {
                "PerformanceMode": "maxIO",
                "LifecyclePolicies":[
                    {
                        "TransitionToIA" : "AFTER_30_DAYS"
                    },
                    {
                        "TransitionToPrimaryStorageClass" : "AFTER_1_ACCESS"
                    }
                ],
                "Encrypted": true,
                "FileSystemTags": [
                    {
                        "Key": "Name",
                        "Value": "TestFileSystem"
                    }
                ],
                "FileSystemPolicy": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Action": [
                                "elasticfilesystem:ClientMount"
                            ],
                            "Principal":  {"AWS": "arn:aws:iam::111122223333:role/EfsReadOnly"}
                        }
                    ]
                },
                "BackupPolicy": {
                    "Status": "ENABLED"
                    },
                "KmsKeyId": {
                    "Fn::GetAtt": [
                        "key",
                        "Arn"
                    ]
                }
            }
        },
        "key": {
            "Type": "AWS::KMS::Key",
            "Properties": {
                "KeyPolicy": {
                    "Version": "2012-10-17",
                    "Id": "key-default-1",
                    "Statement": [
                        {
                            "Sid": "Allow administration of the key",
                            "Effect": "Allow",
                            "Principal": {
                                "AWS": {
                                    "Fn::Join": [
                                        "",
                                        [
                                            "arn:aws:iam::",
                                            {
                                                "Ref": "AWS::AccountId"
                                            },
                                            ":root"
                                        ]
                                    ]
                                }
                            },
                            "Action": [
                                "kms:*"
                            ],
                            "Resource": "*"
                        }
                    ]
                }
            }
        },
        "MountTargetResource1": {
            "Type": "AWS::EFS::MountTarget",
            "Properties": {
                "FileSystemId": {
                    "Ref": "FileSystemResource"
                },
                "SubnetId": {
                    "Ref": "MountTargetSubnetOne"
                },
                "SecurityGroups": [
                    {
                        "Fn::GetAtt": [
                            "MountTargetVPC",
                            "DefaultSecurityGroup"
                        ]
                    }
                ]
            }
        },
        "MountTargetResource2": {
            "Type": "AWS::EFS::MountTarget",
            "Properties": {
                "FileSystemId": {
                    "Ref": "FileSystemResource"
                },
                "SubnetId": {
                    "Ref": "MountTargetSubnetTwo"
                },
                "SecurityGroups": [
                    {
                        "Fn::GetAtt": [
                            "MountTargetVPC",
                            "DefaultSecurityGroup"
                        ]
                    }
                ]
            }
        },
        "MountTargetResource3": {
            "Type": "AWS::EFS::MountTarget",
            "Properties": {
                "FileSystemId": {
                    "Ref": "FileSystemResource"
                },
                "SubnetId": {
                    "Ref": "MountTargetSubnetThree"
                },
                "SecurityGroups": [
                    {
                        "Fn::GetAtt": [
                            "MountTargetVPC",
                            "DefaultSecurityGroup"
                        ]
                    }
                ]
            }
        },
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
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MountTargetVPC:
    Type: AWS::EC2::VPC
    Properties:
      CidrBlock: 172.31.0.0/16

  MountTargetSubnetOne:
    Type: AWS::EC2::Subnet
    Properties:
      CidrBlock: 172.31.1.0/24
      VpcId: !Ref MountTargetVPC
     AvailabilityZone: "us-east-1a"

  MountTargetSubnetTwo:
    Type: AWS::EC2::Subnet
    Properties:
      CidrBlock: 172.31.2.0/24
      VpcId: !Ref MountTargetVPC
      AvailabilityZone: "us-east-1b"

  MountTargetSubnetThree:
    Type: AWS::EC2::Subnet
    Properties:
      CidrBlock: 172.31.3.0/24
      VpcId: !Ref MountTargetVPC
      AvailabilityZone: "us-east-1c"

  FileSystemResource:
    Type: 'AWS::EFS::FileSystem'
    Properties:
      BackupPolicy:
        Status: ENABLED
      PerformanceMode: maxIO
      Encrypted: true
      LifecyclePolicies:
        - TransitionToIA: AFTER_30_DAYS
        - TransitionToPrimaryStorageClass: AFTER_1_ACCESS
      FileSystemTags:
        - Key: Name
          Value: TestFileSystem
      FileSystemPolicy:
        Version: "2012-10-17"
        Statement:
          - Effect: "Allow"
            Action:
              - "elasticfilesystem:ClientMount"
            Principal:
                AWS: 'arn:aws:iam::111122223333:role/EfsReadOnly'
      KmsKeyId: !GetAtt
        - key
        - Arn
  key:
    Type: AWS::KMS::Key
    Properties:
      KeyPolicy:
        Version: 2012-10-17
        Id: key-default-1
        Statement:
          - Sid: Allow administration of the key
            Effect: Allow
            Principal:
              AWS: !Join
                - ''
                - - 'arn:aws:iam::'
                  - !Ref 'AWS::AccountId'
                  - ':root'
            Action:
              - 'kms:*'
            Resource:
              - '*'

  MountTargetResource1:
    Type: AWS::EFS::MountTarget
    Properties:
      FileSystemId: !Ref FileSystemResource
      SubnetId: !Ref MountTargetSubnetOne
      SecurityGroups:
      - !GetAtt MountTargetVPC.DefaultSecurityGroup

  MountTargetResource2:
    Type: AWS::EFS::MountTarget
    Properties:
      FileSystemId: !Ref FileSystemResource
      SubnetId: !Ref MountTargetSubnetTwo
      SecurityGroups:
      - !GetAtt MountTargetVPC.DefaultSecurityGroup

  MountTargetResource3:
    Type: AWS::EFS::MountTarget
    Properties:
      FileSystemId: !Ref FileSystemResource
      SubnetId: !Ref MountTargetSubnetThree
      SecurityGroups:
      - !GetAtt MountTargetVPC.DefaultSecurityGroup

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

### Create a One Zone file system

The following example declares an encrypted One Zone file system in the us-east-1a Availability Zone.

#### JSON

```json

"{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "MountTargetVPC": {
            "Type": "AWS::EC2::VPC",
            "Properties": {
                "CidrBlock": "172.31.0.0/16"
            }
        },
        "MountTargetSubnetOne": {
            "Type": "AWS::EC2::Subnet",
            "Properties": {
                "CidrBlock": "172.31.1.0/24",
                "VpcId": {
                    "Ref": "MountTargetVPC"
                },
                "AvailabilityZone": "us-east-1a"
            }
        },
       "FileSystemResource": {
            "Type": "AWS::EFS::FileSystem",
            "Properties": {
                "AvailabilityZoneName": "us-east-1a",
                "LifecyclePolicies":[
                    {
                        "TransitionToIA" : "AFTER_30_DAYS"
                    },
                    {
                        "TransitionToPrimaryStorageClass" : "AFTER_1_ACCESS"
                    }
                ],
                "Encrypted": true,
                "FileSystemTags": [
                    {
                        "Key": "Name",
                        "Value": "TestFileSystem"
                    }
                ],
                "FileSystemPolicy": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Action": [
                                "elasticfilesystem:ClientMount"
                            ],
                            "Principal":  {"AWS": "arn:aws:iam::111122223333:role/EfsReadOnly"}
                        }
                    ]
                },
                "BackupPolicy": {
                    "Status": "ENABLED"
                    },
                "KmsKeyId": {
                    "Fn::GetAtt": [
                        "key",
                        "Arn"
                    ]
                }
            }
        },
        "key": {
            "Type": "AWS::KMS::Key",
            "Properties": {
                "KeyPolicy": {
                    "Version": "2012-10-17",
                    "Id": "key-default-1",
                    "Statement": [
                        {
                            "Sid": "Allow administration of the key",
                            "Effect": "Allow",
                            "Principal": {
                                "AWS": {
                                    "Fn::Join": [
                                        "",
                                        [
                                            "arn:aws:iam::",
                                            {
                                                "Ref": "AWS::AccountId"
                                            },
                                            ":root"
                                        ]
                                    ]
                                }
                            },
                            "Action": [
                                "kms:*"
                            ],
                            "Resource": "*"
                        }
                    ]
                }
            }
        },
        "MountTargetResource1": {
            "Type": "AWS::EFS::MountTarget",
            "Properties": {
                "FileSystemId": {
                    "Ref": "FileSystemResource"
                },
                "SubnetId": {
                    "Ref": "MountTargetSubnetOne"
                },
                "SecurityGroups": [
                    {
                        "Fn::GetAtt": [
                            "MountTargetVPC",
                            "DefaultSecurityGroup"
                        ]
                    }
                ]
            }
        },
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
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MountTargetVPC:
    Type: AWS::EC2::VPC
    Properties:
      CidrBlock: 172.31.0.0/16

  MountTargetSubnetOne:
    Type: AWS::EC2::Subnet
    Properties:
      CidrBlock: 172.31.1.0/24
      VpcId: !Ref MountTargetVPC
      AvailabilityZone: "us-east-1a"

  FileSystemResource:
    Type: 'AWS::EFS::FileSystem'
    Properties:
      AvailabilityZoneName: us-east-1a
      BackupPolicy:
        Status: ENABLED
      Encrypted: true
      LifecyclePolicies:
        - TransitionToIA: AFTER_30_DAYS
        - TransitionToPrimaryStorageClass: AFTER_1_ACCESS
      FileSystemTags:
        - Key: Name
          Value: TestFileSystem
      FileSystemPolicy:
        Version: "2012-10-17"
        Statement:
          - Effect: "Allow"
            Action:
              - "elasticfilesystem:ClientMount"
            Principal:
                AWS: 'arn:aws:iam::111122223333:role/EfsReadOnly'
      KmsKeyId: !GetAtt
        - key
        - Arn
  key:
    Type: AWS::KMS::Key
    Properties:
      KeyPolicy:
        Version: 2012-10-17
        Id: key-default-1
        Statement:
          - Sid: Allow administration of the key
            Effect: Allow
            Principal:
              AWS: !Join
                - ''
                - - 'arn:aws:iam::'
                  - !Ref 'AWS::AccountId'
                  - ':root'
            Action:
              - 'kms:*'
            Resource:
              - '*'

  MountTargetResource1:
    Type: AWS::EFS::MountTarget
    Properties:
      FileSystemId: !Ref FileSystemResource
      SubnetId: !Ref MountTargetSubnetOne
      SecurityGroups:
      - !GetAtt MountTargetVPC.DefaultSecurityGroup

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

- [Amazon EFS: How it works](https://docs.aws.amazon.com/efs/latest/ug/how-it-works.html)

- [Creating an Amazon EFS file system](https://docs.aws.amazon.com/efs/latest/ug/creating-using-fs.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RootDirectory

BackupPolicy
