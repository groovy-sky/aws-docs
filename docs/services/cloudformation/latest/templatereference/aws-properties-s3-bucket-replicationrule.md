This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket ReplicationRule

Specifies which Amazon S3 objects to replicate and where to store the replicas.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeleteMarkerReplication" : DeleteMarkerReplication,
  "Destination" : ReplicationDestination,
  "Filter" : ReplicationRuleFilter,
  "Id" : String,
  "Prefix" : String,
  "Priority" : Integer,
  "SourceSelectionCriteria" : SourceSelectionCriteria,
  "Status" : String
}

```

### YAML

```yaml

  DeleteMarkerReplication:
    DeleteMarkerReplication
  Destination:
    ReplicationDestination
  Filter:
    ReplicationRuleFilter
  Id: String
  Prefix: String
  Priority: Integer
  SourceSelectionCriteria:
    SourceSelectionCriteria
  Status: String

```

## Properties

`DeleteMarkerReplication`

Specifies whether Amazon S3 replicates delete markers. If you specify a `Filter` in your
replication configuration, you must also include a `DeleteMarkerReplication` element. If your
`Filter` includes a `Tag` element, the `DeleteMarkerReplication` `Status` must be set to Disabled, because Amazon S3 does not support replicating delete markers
for tag-based rules. For an example configuration, see [Basic Rule\
Configuration](../../../s3/latest/dev/replication-add-config.md#replication-config-min-rule-config).

For more information about delete marker replication, see [Basic Rule Configuration](../../../s3/latest/dev/delete-marker-replication.md).

###### Note

If you are using an earlier version of the replication configuration, Amazon S3 handles replication of
delete markers differently. For more information, see [Backward Compatibility](../../../s3/latest/dev/replication-add-config.md#replication-backward-compat-considerations).

_Required_: No

_Type_: [DeleteMarkerReplication](aws-properties-s3-bucket-deletemarkerreplication.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destination`

A container for information about the replication destination and its configurations including
enabling the S3 Replication Time Control (S3 RTC).

_Required_: Yes

_Type_: [ReplicationDestination](aws-properties-s3-bucket-replicationdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

A filter that identifies the subset of objects to which the replication rule applies. A
`Filter` must specify exactly one `Prefix`, `TagFilter`, or
an `And` child element. The use of the filter field indicates that this is a V2
replication configuration. This field isn't supported in a V1 replication
configuration.

###### Note

V1 replication configuration only supports filtering by key prefix. To filter using a V1
replication configuration, add the `Prefix` directly as a child element of the
`Rule` element.

_Required_: No

_Type_: [ReplicationRuleFilter](aws-properties-s3-bucket-replicationrulefilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

A unique identifier for the rule. The maximum value is 255 characters. If you don't
specify a value, AWS CloudFormation generates a random ID. When using a V2
replication configuration this property is capitalized as "ID".

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

An object key name prefix that identifies the object or objects to which the rule applies.
The maximum prefix length is 1,024 characters. To include all objects in a bucket, specify an
empty string. To filter using a V1 replication configuration, add the `Prefix`
directly as a child element of the `Rule` element.

###### Important

Replacement must be made for object keys containing special characters (such as carriage
returns) when using XML requests. For more information, see [XML\
related object key constraints](../../../s3/latest/userguide/object-keys.md#object-key-xml-related-constraints).

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The priority indicates which rule has precedence whenever two or more replication rules conflict.
Amazon S3 will attempt to replicate objects according to all replication rules. However, if there are two or
more rules with the same destination bucket, then objects will be replicated according to the rule with
the highest priority. The higher the number, the higher the priority.

For more information, see [Replication](../../../s3/latest/dev/replication.md) in the _Amazon S3 User Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceSelectionCriteria`

A container that describes additional filters for identifying the source objects that you
want to replicate. You can choose to enable or disable the replication of these
objects.

_Required_: No

_Type_: [SourceSelectionCriteria](aws-properties-s3-bucket-sourceselectioncriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Specifies whether the rule is enabled.

_Required_: Yes

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Associate a replication configuration IAM role with an S3 bucket](#aws-properties-s3-bucket-replicationrule--examples--Associate_a_replication_configuration_IAM_role_with_an_S3_bucket)

- [Enable versioning and replicate objects](#aws-properties-s3-bucket-replicationrule--examples--Enable_versioning_and_replicate_objects)

### Associate a replication configuration IAM role with an S3 bucket

The following example creates an S3 bucket and grants it permission to write to a
replication bucket by using an AWS Identity and Access Management (IAM)
role. To avoid a circular dependency, the role's policy is declared as a separate
resource. The bucket depends on the `WorkItemBucketBackupRole` role. If the
policy is included in the role, the role also depends on the bucket.

#### JSON

```json

{
    "Resources": {
        "RecordServiceS3Bucket": {
            "Type": "AWS::S3::Bucket",
            "DeletionPolicy": "Retain",
            "Properties": {
                "ReplicationConfiguration": {
                    "Role": {
                        "Fn::GetAtt": [
                            "WorkItemBucketBackupRole",
                            "Arn"
                        ]
                    },
                    "Rules": [
                        {
                            "Destination": {
                                "Bucket": {
                                    "Fn::Join": [
                                        "",
                                        [
                                            "arn:aws:s3:::",
                                            {
                                                "Fn::Join": [
                                                    "-",
                                                    [
                                                        {
                                                            "Ref": "AWS::Region"
                                                        },
                                                        {
                                                            "Ref": "AWS::StackName"
                                                        },
                                                        "replicationbucket"
                                                    ]
                                                ]
                                            }
                                        ]
                                    ]
                                },
                                "StorageClass": "STANDARD"
                            },
                            "Id": "Backup",
                            "Prefix": "",
                            "Status": "Enabled"
                        }
                    ]
                },
                "VersioningConfiguration": {
                    "Status": "Enabled"
                }
            }
        },
        "WorkItemBucketBackupRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "AssumeRolePolicyDocument": {
                    "Statement": [
                        {
                            "Action": [
                                "sts:AssumeRole"
                            ],
                            "Effect": "Allow",
                            "Principal": {
                                "Service": [
                                    "s3.amazonaws.com"
                                ]
                            }
                        }
                    ]
                }
            }
        },
        "BucketBackupPolicy": {
            "Type": "AWS::IAM::Policy",
            "Properties": {
                "PolicyDocument": {
                    "Statement": [
                        {
                            "Action": [
                                "s3:GetReplicationConfiguration",
                                "s3:ListBucket"
                            ],
                            "Effect": "Allow",
                            "Resource": [
                                {
                                    "Fn::Join": [
                                        "",
                                        [
                                            "arn:aws:s3:::",
                                            {
                                                "Ref": "RecordServiceS3Bucket"
                                            }
                                        ]
                                    ]
                                }
                            ]
                        },
                        {
                            "Action": [
                                "s3:GetObjectVersion",
                                "s3:GetObjectVersionAcl"
                            ],
                            "Effect": "Allow",
                            "Resource": [
                                {
                                    "Fn::Join": [
                                        "",
                                        [
                                            "arn:aws:s3:::",
                                            {
                                                "Ref": "RecordServiceS3Bucket"
                                            },
                                            "/*"
                                        ]
                                    ]
                                }
                            ]
                        },
                        {
                            "Action": [
                                "s3:ReplicateObject",
                                "s3:ReplicateDelete"
                            ],
                            "Effect": "Allow",
                            "Resource": [
                                {
                                    "Fn::Join": [
                                        "",
                                        [
                                            "arn:aws:s3:::",
                                            {
                                                "Fn::Join": [
                                                    "-",
                                                    [
                                                        {
                                                            "Ref": "AWS::Region"
                                                        },
                                                        {
                                                            "Ref": "AWS::StackName"
                                                        },
                                                        "replicationbucket"
                                                    ]
                                                ]
                                            },
                                            "/*"
                                        ]
                                    ]
                                }
                            ]
                        }
                    ]
                },
                "PolicyName": "BucketBackupPolicy",
                "Roles": [
                    {
                        "Ref": "WorkItemBucketBackupRole"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  RecordServiceS3Bucket:
    Type: 'AWS::S3::Bucket'
    DeletionPolicy: Retain
    Properties:
      ReplicationConfiguration:
        Role: !GetAtt
          - WorkItemBucketBackupRole
          - Arn
        Rules:
          - Destination:
              Bucket: !Join
                - ''
                - - 'arn:aws:s3:::'
                  - !Join
                    - '-'
                    - - !Ref 'AWS::Region'
                      - !Ref 'AWS::StackName'
                      - replicationbucket
              StorageClass: STANDARD
            Id: Backup
            Prefix: ''
            Status: Enabled
      VersioningConfiguration:
        Status: Enabled
  WorkItemBucketBackupRole:
    Type: 'AWS::IAM::Role'
    Properties:
      AssumeRolePolicyDocument:
        Statement:
          - Action:
              - 'sts:AssumeRole'
            Effect: Allow
            Principal:
              Service:
                - s3.amazonaws.com
  BucketBackupPolicy:
    Type: 'AWS::IAM::Policy'
    Properties:
      PolicyDocument:
        Statement:
          - Action:
              - 's3:GetReplicationConfiguration'
              - 's3:ListBucket'
            Effect: Allow
            Resource:
              - !Join
                - ''
                - - 'arn:aws:s3:::'
                  - !Ref RecordServiceS3Bucket
          - Action:
              - 's3:GetObjectVersion'
              - 's3:GetObjectVersionAcl'
            Effect: Allow
            Resource:
              - !Join
                - ''
                - - 'arn:aws:s3:::'
                  - !Ref RecordServiceS3Bucket
                  - /*
          - Action:
              - 's3:ReplicateObject'
              - 's3:ReplicateDelete'
            Effect: Allow
            Resource:
              - !Join
                - ''
                - - 'arn:aws:s3:::'
                  - !Join
                    - '-'
                    - - !Ref 'AWS::Region'
                      - !Ref 'AWS::StackName'
                      - replicationbucket
                  - /*
      PolicyName: BucketBackupPolicy
      Roles:
        - !Ref WorkItemBucketBackupRole
```

### Enable versioning and replicate objects

The following example enables versioning and two replication rules. The rules copy
objects prefixed with either `MyPrefix` and `MyOtherPrefix` and
stores the copied objects in a bucket named `my-replication-bucket`.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "VersioningConfiguration": {
                    "Status": "Enabled"
                },
                "ReplicationConfiguration": {
                    "Role": "arn:aws:iam::123456789012:role/replication_role",
                    "Rules": [
                        {
                            "Id": "MyRule1",
                            "Status": "Enabled",
                            "Prefix": "MyPrefix",
                            "Destination": {
                                "Bucket": "arn:aws:s3:::my-replication-bucket",
                                "StorageClass": "STANDARD"
                            }
                        },
                        {
                            "Status": "Enabled",
                            "Prefix": "MyOtherPrefix",
                            "Destination": {
                                "Bucket": "arn:aws:s3:::my-replication-bucket"
                            }
                        }
                    ]
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
  S3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      VersioningConfiguration:
        Status: Enabled
      ReplicationConfiguration:
        Role: 'arn:aws:iam::123456789012:role/replication_role'
        Rules:
          - Id: MyRule1
            Status: Enabled
            Prefix: MyPrefix
            Destination:
              Bucket: 'arn:aws:s3:::my-replication-bucket'
              StorageClass: STANDARD
          - Status: Enabled
            Prefix: MyOtherPrefix
            Destination:
              Bucket: 'arn:aws:s3:::my-replication-bucket'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationDestination

ReplicationRuleAndOperator

All content copied from https://docs.aws.amazon.com/.
