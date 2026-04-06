This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket

The `AWS::S3::Bucket` resource creates an Amazon S3 bucket in the same AWS Region where you create the AWS CloudFormation stack.

To control how AWS CloudFormation handles the bucket when the stack is
deleted, you can set a deletion policy for your bucket. You can choose to
_retain_ the bucket or to _delete_ the bucket. For
more information, see [DeletionPolicy\
Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).

###### Important

You can only delete empty buckets. Deletion fails for buckets that have contents.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3::Bucket",
  "Properties" : {
      "AbacStatus" : String,
      "AccelerateConfiguration" : AccelerateConfiguration,
      "AccessControl" : String,
      "AnalyticsConfigurations" : [ AnalyticsConfiguration, ... ],
      "BucketEncryption" : BucketEncryption,
      "BucketName" : String,
      "BucketNamePrefix" : String,
      "BucketNamespace" : String,
      "CorsConfiguration" : CorsConfiguration,
      "IntelligentTieringConfigurations" : [ IntelligentTieringConfiguration, ... ],
      "InventoryConfigurations" : [ InventoryConfiguration, ... ],
      "LifecycleConfiguration" : LifecycleConfiguration,
      "LoggingConfiguration" : LoggingConfiguration,
      "MetadataConfiguration" : MetadataConfiguration,
      "MetadataTableConfiguration" : MetadataTableConfiguration,
      "MetricsConfigurations" : [ MetricsConfiguration, ... ],
      "NotificationConfiguration" : NotificationConfiguration,
      "ObjectLockConfiguration" : ObjectLockConfiguration,
      "ObjectLockEnabled" : Boolean,
      "OwnershipControls" : OwnershipControls,
      "PublicAccessBlockConfiguration" : PublicAccessBlockConfiguration,
      "ReplicationConfiguration" : ReplicationConfiguration,
      "Tags" : [ Tag, ... ],
      "VersioningConfiguration" : VersioningConfiguration,
      "WebsiteConfiguration" : WebsiteConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::S3::Bucket
Properties:
  AbacStatus: String
  AccelerateConfiguration:
    AccelerateConfiguration
  AccessControl: String
  AnalyticsConfigurations:
    - AnalyticsConfiguration
  BucketEncryption:
    BucketEncryption
  BucketName: String
  BucketNamePrefix: String
  BucketNamespace: String
  CorsConfiguration:
    CorsConfiguration
  IntelligentTieringConfigurations:
    - IntelligentTieringConfiguration
  InventoryConfigurations:
    - InventoryConfiguration
  LifecycleConfiguration:
    LifecycleConfiguration
  LoggingConfiguration:
    LoggingConfiguration
  MetadataConfiguration:
    MetadataConfiguration
  MetadataTableConfiguration:
    MetadataTableConfiguration
  MetricsConfigurations:
    - MetricsConfiguration
  NotificationConfiguration:
    NotificationConfiguration
  ObjectLockConfiguration:
    ObjectLockConfiguration
  ObjectLockEnabled: Boolean
  OwnershipControls:
    OwnershipControls
  PublicAccessBlockConfiguration:
    PublicAccessBlockConfiguration
  ReplicationConfiguration:
    ReplicationConfiguration
  Tags:
    - Tag
  VersioningConfiguration:
    VersioningConfiguration
  WebsiteConfiguration:
    WebsiteConfiguration

```

## Properties

`AbacStatus`

The ABAC status of the general purpose bucket. When ABAC is enabled for the general purpose bucket, you can use tags to manage access to the general purpose buckets as well as for cost tracking purposes. When ABAC is disabled for the general purpose buckets, you can only use tags for cost tracking purposes. For more information, see [Using tags with S3 general purpose buckets](../../../s3/latest/userguide/buckets-tagging.md).

_Required_: No

_Type_: String

_Allowed values_: `Enabled | Disabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccelerateConfiguration`

Configures the transfer acceleration state for an Amazon S3 bucket. For more information, see [Amazon S3 Transfer\
Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the _Amazon S3 User Guide_.

_Required_: No

_Type_: [AccelerateConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-accelerateconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccessControl`

###### Important

This is a legacy property, and it is not recommended for most use cases. A majority of
modern use cases in Amazon S3 no longer require the use of ACLs, and we recommend that you
keep ACLs disabled. For more information, see [Controlling object\
ownership](../../../s3/latest/userguide/about-object-ownership.md) in the _Amazon S3 User Guide_.

A canned access control list (ACL) that grants predefined permissions to the bucket. For
more information about canned ACLs, see [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) in the
_Amazon S3 User Guide_.

S3 buckets are created with ACLs disabled by default. Therefore, unless you explicitly set the [AWS::S3::OwnershipControls](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ownershipcontrols.html) property to enable ACLs, your resource will fail to deploy
with any value other than Private. Use cases requiring ACLs are uncommon.

The majority of access control configurations can be successfully and more easily
achieved with bucket policies. For more information, see [AWS::S3::BucketPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html). For examples of common policy configurations, including S3
Server Access Logs buckets and more, see [Bucket policy examples](../../../s3/latest/userguide/example-bucket-policies.md) in the
_Amazon S3 User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `AuthenticatedRead | AwsExecRead | BucketOwnerFullControl | BucketOwnerRead | LogDeliveryWrite | Private | PublicRead | PublicReadWrite`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AnalyticsConfigurations`

Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.

_Required_: No

_Type_: Array of [AnalyticsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-analyticsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketEncryption`

Specifies default encryption for a bucket using server-side encryption with Amazon
S3-managed keys (SSE-S3), AWS KMS-managed keys (SSE-KMS), or dual-layer server-side encryption with KMS-managed keys (DSSE-KMS). For
information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3\
Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the _Amazon S3 User Guide_.

_Required_: No

_Type_: [BucketEncryption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-bucketencryption.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketName`

A name for the bucket. If you don't specify a name, AWS CloudFormation
generates a unique ID and uses that ID for the bucket name. The bucket name must contain only
lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket\
restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html). For more information, see [Rules for naming Amazon\
S3 buckets](../../../s3/latest/userguide/bucketnamingrules.md) in the _Amazon S3 User Guide_.

###### Important

If you specify a name, you can't perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you need to
replace the resource, specify a new name.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BucketNamePrefix`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BucketNamespace`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `global | account-regional`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CorsConfiguration`

Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information,
see [Enabling Cross-Origin Resource\
Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the _Amazon S3 User Guide_.

_Required_: No

_Type_: [CorsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-corsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntelligentTieringConfigurations`

Defines how Amazon S3 handles Intelligent-Tiering storage.

_Required_: No

_Type_: Array of [IntelligentTieringConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-intelligenttieringconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InventoryConfigurations`

Specifies the S3 Inventory configuration for an Amazon S3 bucket. For more information, see [GET Bucket\
inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the _Amazon S3 API Reference_.

_Required_: No

_Type_: Array of [InventoryConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-inventoryconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecycleConfiguration`

Specifies the lifecycle configuration for objects in an Amazon S3 bucket. For more information, see
[Object Lifecycle\
Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the _Amazon S3 User Guide_.

_Required_: No

_Type_: [LifecycleConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-lifecycleconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingConfiguration`

Settings that define where logs are stored.

_Required_: No

_Type_: [LoggingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-loggingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetadataConfiguration`

The S3 Metadata configuration for a general purpose bucket.

_Required_: No

_Type_: [MetadataConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-metadataconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetadataTableConfiguration`

The metadata table configuration of an Amazon S3 general purpose bucket.

_Required_: No

_Type_: [MetadataTableConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-metadatatableconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricsConfigurations`

Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics
configuration ID) from an Amazon S3 bucket. If you're updating an existing metrics configuration, note that
this is a full replacement of the existing metrics configuration. If you don't include the elements you
want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html).

_Required_: No

_Type_: Array of [MetricsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-metricsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationConfiguration`

Configuration that defines how Amazon S3 handles bucket notifications.

_Required_: No

_Type_: [NotificationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-notificationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectLockConfiguration`

###### Note

This operation is not supported for directory buckets.

Places an Object Lock configuration on the specified bucket. The rule specified in the Object Lock
configuration will be applied by default to every new object placed in the specified bucket. For more
information, see [Locking\
Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).

###### Note

- The `DefaultRetention` settings require both a mode and a period.

- The `DefaultRetention` period can be either `Days` or `Years`
but you must select one. You cannot specify `Days` and `Years` at the same
time.

- You can enable Object Lock for new or existing buckets. For more information, see [Configuring\
Object Lock](../../../s3/latest/userguide/object-lock-configure.md).

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

_Required_: No

_Type_: [ObjectLockConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-objectlockconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectLockEnabled`

Indicates whether this bucket has an Object Lock configuration enabled. Enable
`ObjectLockEnabled` when you apply `ObjectLockConfiguration` to a
bucket.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OwnershipControls`

Configuration that defines how Amazon S3 handles Object Ownership rules.

_Required_: No

_Type_: [OwnershipControls](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-ownershipcontrols.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublicAccessBlockConfiguration`

Configuration that defines how Amazon S3 handles public access.

_Required_: No

_Type_: [PublicAccessBlockConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-publicaccessblockconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationConfiguration`

Configuration for replicating objects in an S3 bucket. To enable replication, you must
also enable versioning by using the `VersioningConfiguration` property.

Amazon S3 can store replicated objects in a single destination bucket or multiple
destination buckets. The destination bucket or buckets must already exist.

_Required_: No

_Type_: [ReplicationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-replicationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An arbitrary set of tags (key-value pairs) for this S3 bucket.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersioningConfiguration`

Enables multiple versions of all objects in this bucket. You might enable versioning to
prevent objects from being deleted or overwritten by mistake or to archive objects so that you
can retrieve previous versions of them.

###### Note

When you enable versioning on a bucket for the first time, it might take a short
amount of time for the change to be fully propagated. We recommend that you wait for 15
minutes after enabling versioning before issuing write operations
( `PUT`
or
`DELETE`)
on objects in the bucket.

_Required_: No

_Type_: [VersioningConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-versioningconfiguration.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`WebsiteConfiguration`

Information used to configure the bucket as a static website. For more information, see
[Hosting Websites\
on Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html).

_Required_: No

_Type_: [WebsiteConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-websiteconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the bucket name.

Example: `
                            amzn-s3-demo-bucket
                        `

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of the specified bucket.

Example: `arn:aws:s3:::DOC-EXAMPLE-BUCKET`

`DomainName`

Returns the IPv4 DNS name of the specified bucket.

Example: `DOC-EXAMPLE-BUCKET.s3.amazonaws.com`

`DualStackDomainName`

Returns the IPv6 DNS name of the specified bucket.

Example: ` DOC-EXAMPLE-BUCKET.s3.dualstack.us-east-2.amazonaws.com`

For more information about dual-stack endpoints, see [Using Amazon S3 Dual-Stack\
Endpoints](https://docs.aws.amazon.com/AmazonS3/latest/dev/dual-stack-endpoints.html).

`MetadataConfiguration.InventoryTableConfiguration.TableArn`

Property description not available.

`MetadataConfiguration.InventoryTableConfiguration.TableName`

Property description not available.

`MetadataConfiguration.JournalTableConfiguration.TableArn`

Property description not available.

`MetadataConfiguration.JournalTableConfiguration.TableName`

Property description not available.

`MetadataTableConfiguration.S3TablesDestination.TableArn`

The Amazon Resource Name (ARN) for the metadata table in the metadata table configuration. The specified metadata table name must be unique within the `aws_s3_metadata` namespace in the destination table bucket.

Example: `arn:aws:s3tables:region:account-id:bucket/amzn-s3-demo-bucket/table/1234567890abcdef0`

`MetadataTableConfiguration.S3TablesDestination.TableNamespace`

The table bucket namespace for the metadata table in the specified bucket's metadata table configuration. This value is always `aws_s3_metadata`.

`RegionalDomainName`

Returns the regional domain name of the specified bucket.

Example: `DOC-EXAMPLE-BUCKET.s3.us-east-2.amazonaws.com`

`WebsiteURL`

Returns the Amazon S3 website endpoint for the specified bucket.

Example (IPv4): `http://DOC-EXAMPLE-BUCKET.s3-website.us-east-2.amazonaws.com`

Example (IPv6):
`http://DOC-EXAMPLE-BUCKET.s3.dualstack.us-east-2.amazonaws.com`

## Examples

- [Create an S3 bucket](#aws-resource-s3-bucket--examples--Create_an_S3_bucket)

- [Associate a replication configuration IAM role with an S3 bucket](#aws-resource-s3-bucket--examples--Associate_a_replication_configuration_IAM_role_with_an_S3_bucket)

- [Granting public access to S3 buckets](#aws-resource-s3-bucket--examples--Granting_public_access_to_S3_buckets)

- [Enabling ACLs](#aws-resource-s3-bucket--examples--Enabling_ACLs)

- [Configure a static website with a routing rule](#aws-resource-s3-bucket--examples--Configure_a_static_website_with_a_routing_rule)

- [Enable cross-origin resource sharing](#aws-resource-s3-bucket--examples--Enable_cross-origin_resource_sharing)

- [Manage the lifecycle for S3 objects](#aws-resource-s3-bucket--examples--Manage_the_lifecycle_for_S3_objects)

- [Log access requests for a specific S3 bucket](#aws-resource-s3-bucket--examples--Log_access_requests_for_a_specific_S3_bucket)

- [Receive S3 bucket notifications to an SNS topic](#aws-resource-s3-bucket--examples--Receive_S3_bucket_notifications_to_an_SNS_topic)

- [Enable versioning and replicate objects](#aws-resource-s3-bucket--examples--Enable_versioning_and_replicate_objects)

- [Specify analytics and inventory configurations for an S3 bucket](#aws-resource-s3-bucket--examples--Specify_analytics_and_inventory_configurations_for_an_S3_bucket)

### Create an S3 bucket

The following example creates an S3 bucket with a `Retain` deletion
policy.

#### JSON

```json

{
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "DeletionPolicy": "Retain",
            "Properties": {
                "BucketName": "DOC-EXAMPLE-BUCKET"
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  S3Bucket:
    Type: 'AWS::S3::Bucket'
    DeletionPolicy: Retain
    Properties:
      BucketName: DOC-EXAMPLE-BUCKET
```

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

### Granting public access to S3 buckets

When you create a new bucket, all Block Public Access settings are automatically
enabled. We recommend that you keep all Block Public Access settings enabled. If you require some level of public access to your buckets, you can disable Block Public Access settings. The following example shows creating a bucket called `my-bucket` and
then disabling Block Public Access. A public bucket policy is then added to the bucket.

###### Note

The following example assumes the `BlockPublicPolicy` and
`RestrictPublicBuckets` Block Public Access settings have been disabled at
the account level.

#### JSON

```json

        {
          "Resources": {
            "MyBucket": {
              "Type": "AWS::S3::Bucket",
              "Properties": {
                "BucketName": "my-bucket",
                "PublicAccessBlockConfiguration": {
                  "BlockPublicAcls": false,
                  "BlockPublicPolicy": false,
                  "IgnorePublicAcls": false,
                  "RestrictPublicBuckets": false

                }
              }
            },
            "MyBucketPolicy": {
              "Type": "AWS::S3::BucketPolicy",
              "Properties": {
                "Bucket": {
                  "Ref": "MyBucket"
                },
                "PolicyDocument": {
                  "Version": "2012-10-17",
                  "Statement": [
                    {
                       "Effect": "Allow",
                       "Principal": "*",
                       "Action": "s3:GetObject",
                       "Resource": {
                         "Fn::Join": [
                           "",
                           [
                             "arn:aws:s3:::",
                             {
                               "Ref": "MyBucket"
                             },
                             "/*"
                           ]
                         ]
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

        Resources:
          MyBucket:
            Type: 'AWS::S3::Bucket'
            Properties:
              BucketName: my-bucket
              PublicAccessBlockConfiguration:
                BlockPublicAcls: false
                BlockPublicPolicy: false
                IgnorePublicAcls: false
                RestrictPublicBuckets: false
          MyBucketPolicy:
            Type: 'AWS::S3::BucketPolicy'
            Properties:
              Bucket:
                Ref: 'MyBucket'
              PolicyDocument:
                Version: '2012-10-17'
                Statement:
                  - Effect: Allow
                    Principal: '*'
                    Action: 's3:GetObject'
                    Resource:
                      Fn::Join:
                        - ''
                        - - 'arn:aws:s3:::'
                          - Ref: 'MyBucket'
                          - '/*'
```

### Enabling ACLs

By default, S3 Object Ownership is set to `BucketOwnerEnforced` and ACLs are disabled. A majority of modern use cases in S3 no longer require the use of ACLs, and we recommend that you keep ACLs disabled. With ACLs disabled, you can control access to all objects in your bucket, regardless of who uploaded the objects to your bucket.
If your specific use case requires enabling ACLs, you can set S3 Object Ownership to `BucketOwnerPreferred` or `ObjectWriter`. For more information, see [Controlling\
ownership of objects and disabling ACLs](../../../s3/latest/userguide/about-object-ownership.md) in the _Amazon S3 User_
_Guide_.

The following example shows Object Ownership set to `BucketOwnerPreferred`.

#### JSON

```json

        {
          "Resources": {
            "MyBucket": {
              "Type": "AWS::S3::Bucket",
              "Properties": {
                "BucketName": "my-bucket",
                "OwnershipControls": {
                       "Rules": [
                           {
                               "ObjectOwnership": "BucketOwnerPreferred"
                           }
                       ]
                   },
                "AccessControl": "AwsExecRead"
              }
           }
         }
       }
```

#### YAML

```yaml

        Resources:
          MyBucket:
            Type: 'AWS::S3::Bucket'
            Properties:
              BucketName: my-bucket
              OwnershipControls:
                Rules:
                - ObjectOwnership: BucketOwnerPreferred
              AccessControl: AwsExecRead

```

### Configure a static website with a routing rule

In this example, `AWS::S3::Bucket's Fn::GetAtt` values are used to provide
outputs. If an HTTP 404 error occurs, the routing rule redirects requests to an EC2
instance and inserts the object key prefix `report-404/` in the redirect. For
example, if you request a page called `out1/ExamplePage.html` and it results in an
HTTP 404 error, the request is routed to a page called
`report-404/ExamplePage.html` on the specified instance. For all other HTTP
error codes, `error.html` is returned.

This example also specifies a metrics configuration called `EntireBucket`
that enables CloudWatch request metrics at the bucket level.

#### JSON

```json

{
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "PublicRead",
                "BucketName": "public-bucket",
                "MetricsConfigurations": [
                    {
                        "Id": "EntireBucket"
                    }
                ],
                "WebsiteConfiguration": {
                    "IndexDocument": "index.html",
                    "ErrorDocument": "error.html",
                    "RoutingRules": [
                        {
                            "RoutingRuleCondition": {
                                "HttpErrorCodeReturnedEquals": "404",
                                "KeyPrefixEquals": "out1/"
                            },
                            "RedirectRule": {
                                "HostName": "ec2-11-22-333-44.compute-1.amazonaws.com",
                                "ReplaceKeyPrefixWith": "report-404/"
                            }
                        }
                    ]
                }
            },
            "DeletionPolicy": "Retain"
        }
    },
    "Outputs": {
        "WebsiteURL": {
            "Value": {
                "Fn::GetAtt": [
                    "S3Bucket",
                    "WebsiteURL"
                ]
            },
            "Description": "URL for website hosted on S3"
        },
        "S3BucketSecureURL": {
            "Value": {
                "Fn::Join": [
                    "",
                    [
                        "https://",
                        {
                            "Fn::GetAtt": [
                                "S3Bucket",
                                "DomainName"
                            ]
                        }
                    ]
                ]
            },
            "Description": "Name of S3 bucket to hold website content"
        }
    }
}
```

#### YAML

```yaml

Resources:
  S3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      AccessControl: PublicRead
      BucketName: public-bucket
      MetricsConfigurations:
        - Id: EntireBucket
      WebsiteConfiguration:
        IndexDocument: index.html
        ErrorDocument: error.html
        RoutingRules:
          - RoutingRuleCondition:
              HttpErrorCodeReturnedEquals: '404'
              KeyPrefixEquals: out1/
            RedirectRule:
              HostName: ec2-11-22-333-44.compute-1.amazonaws.com
              ReplaceKeyPrefixWith: report-404/
    DeletionPolicy: Retain
Outputs:
  WebsiteURL:
    Value: !GetAtt
      - S3Bucket
      - WebsiteURL
    Description: URL for website hosted on S3
  S3BucketSecureURL:
    Value: !Join
      - ''
      - - 'https://'
        - !GetAtt
          - S3Bucket
          - DomainName
    Description: Name of S3 bucket to hold website content
```

### Enable cross-origin resource sharing

The following example template shows a public S3 bucket with two cross-origin resource
sharing rules.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "PublicRead",
                "CorsConfiguration": {
                    "CorsRules": [
                        {
                            "AllowedHeaders": [
                                "*"
                            ],
                            "AllowedMethods": [
                                "GET"
                            ],
                            "AllowedOrigins": [
                                "*"
                            ],
                            "ExposedHeaders": [
                                "Date"
                            ],
                            "Id": "myCORSRuleId1",
                            "MaxAge": 3600
                        },
                        {
                            "AllowedHeaders": [
                                "x-amz-*"
                            ],
                            "AllowedMethods": [
                                "DELETE"
                            ],
                            "AllowedOrigins": [
                                "http://www.example.com",
                                "http://www.example.net"
                            ],
                            "ExposedHeaders": [
                                "Connection",
                                "Server",
                                "Date"
                            ],
                            "Id": "myCORSRuleId2",
                            "MaxAge": 1800
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "BucketName": {
            "Value": {
                "Ref": "S3Bucket"
            },
            "Description": "Name of the sample Amazon S3 bucket with CORS enabled."
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
      AccessControl: PublicRead
      CorsConfiguration:
        CorsRules:
          - AllowedHeaders:
              - '*'
            AllowedMethods:
              - GET
            AllowedOrigins:
              - '*'
            ExposedHeaders:
              - Date
            Id: myCORSRuleId1
            MaxAge: 3600
          - AllowedHeaders:
              - x-amz-*
            AllowedMethods:
              - DELETE
            AllowedOrigins:
              - 'http://www.example.com'
              - 'http://www.example.net'
            ExposedHeaders:
              - Connection
              - Server
              - Date
            Id: myCORSRuleId2
            MaxAge: 1800
Outputs:
  BucketName:
    Value: !Ref S3Bucket
    Description: Name of the sample Amazon S3 bucket with CORS enabled.
```

### Manage the lifecycle for S3 objects

The following example template shows an S3 bucket with a lifecycle configuration rule.
The rule applies to all objects with the `glacier` key prefix. The objects are
transitioned to Glacier after one day, and deleted after one year.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "Private",
                "LifecycleConfiguration": {
                    "Rules": [
                        {
                            "Id": "GlacierRule",
                            "Prefix": "glacier",
                            "Status": "Enabled",
                            "ExpirationInDays": 365,
                            "Transitions": [
                                {
                                    "TransitionInDays": 1,
                                    "StorageClass": "GLACIER"
                                }
                            ]
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "BucketName": {
            "Value": {
                "Ref": "S3Bucket"
            },
            "Description": "Name of the sample Amazon S3 bucket with a lifecycle configuration."
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
      AccessControl: Private
      LifecycleConfiguration:
        Rules:
          - Id: GlacierRule
            Prefix: glacier
            Status: Enabled
            ExpirationInDays: 365
            Transitions:
              - TransitionInDays: 1
                StorageClass: GLACIER
Outputs:
  BucketName:
    Value: !Ref S3Bucket
    Description: Name of the sample Amazon S3 bucket with a lifecycle configuration.
```

### Log access requests for a specific S3 bucket

The following example template creates two S3 buckets. The `LoggingBucket`
bucket store the logs from the `S3Bucket` bucket. To receive logs from the
`S3Bucket` bucket, the logging bucket requires log delivery write
permissions.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "LoggingConfiguration": {
                    "DestinationBucketName": {
                        "Ref": "LoggingBucket"
                    },
                    "LogFilePrefix": "testing-logs"
                }
            }
        },
        "LoggingBucket": {
            "Type": "AWS::S3::Bucket"
        },
        "S3BucketPolicy": {
            "Type": "AWS::S3::BucketPolicy",
            "Properties": {
                "Bucket": {
                    "Ref": "LoggingBucket"
                },
                "PolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Action": [
                                "s3:PutObject"
                            ],
                            "Effect": "Allow",
                            "Principal": {
                                "Service": "logging.s3.amazonaws.com"
                            },
                            "Resource": {
                                "Fn::Join": [
                                    "",
                                    [
                                        "arn:aws:s3:::",
                                        {
                                            "Ref": "LoggingBucket"
                                        },
                                        "/*"
                                    ]
                                ]
                            },
                            "Condition": {
                                "ArnLike": {
                                    "aws:SourceArn": {
                                        "Fn::GetAtt": [
                                            "S3Bucket",
                                            "Arn"
                                        ]
                                    }
                                },
                                "StringEquals": {
                                    "aws:SourceAccount": {
                                        "Fn::Sub": "${AWS::AccountId}"
                                    }
                                }
                            }
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "BucketName": {
            "Value": {
                "Ref": "S3Bucket"
            },
            "Description": "Name of the sample Amazon S3 bucket with a logging configuration."
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
      LoggingConfiguration:
        DestinationBucketName: !Ref LoggingBucket
        LogFilePrefix: testing-logs
  LoggingBucket:
    Type: 'AWS::S3::Bucket'
  S3BucketPolicy:
    Type: 'AWS::S3::BucketPolicy'
    Properties:
      Bucket: !Ref LoggingBucket
      PolicyDocument:
        Version: 2012-10-17
        Statement:
          - Action:
              - 's3:PutObject'
            Effect: Allow
            Principal:
              Service: logging.s3.amazonaws.com
            Resource: !Join
              - ''
              - - 'arn:aws:s3:::'
                - !Ref LoggingBucket
                - /*
            Condition:
              ArnLike:
                'aws:SourceArn': !GetAtt
                  - S3Bucket
                  - Arn
              StringEquals:
                'aws:SourceAccount': !Sub '${AWS::AccountId}'
Outputs:
  BucketName:
    Value: !Ref S3Bucket
    Description: Name of the sample Amazon S3 bucket with a logging configuration.

```

### Receive S3 bucket notifications to an SNS topic

The following example template shows an Amazon S3 bucket with a notification
configuration that sends an event to the specified SNS topic when S3 has lost all replicas
of an object.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "Private",
                "NotificationConfiguration": {
                    "TopicConfigurations": [
                        {
                            "Topic": "arn:aws:sns:us-east-1:123456789012:TestTopic",
                            "Event": "s3:ReducedRedundancyLostObject"
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "BucketName": {
            "Value": {
                "Ref": "S3Bucket"
            },
            "Description": "Name of the sample Amazon S3 bucket with a notification configuration."
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
      AccessControl: Private
      NotificationConfiguration:
        TopicConfigurations:
          - Topic: 'arn:aws:sns:us-east-1:123456789012:TestTopic'
            Event: 's3:ReducedRedundancyLostObject'
Outputs:
  BucketName:
    Value: !Ref S3Bucket
    Description: Name of the sample Amazon S3 bucket with a notification configuration.
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

### Specify analytics and inventory configurations for an S3 bucket

The following example specifies analytics and inventory results to be generated for an
S3 bucket, including the format of the results and the destination bucket. The inventory
list generates reports weekly and includes the current version of each object.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "S3 Bucket with Inventory and Analytics Configurations",
    "Resources": {
        "Helper": {
            "Type": "AWS::S3::Bucket"
        },
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AnalyticsConfigurations": [
                    {
                        "Id": "AnalyticsConfigurationId",
                        "StorageClassAnalysis": {
                            "DataExport": {
                                "Destination": {
                                    "BucketArn": {
                                        "Fn::GetAtt": [
                                            "Helper",
                                            "Arn"
                                        ]
                                    },
                                    "Format": "CSV",
                                    "Prefix": "AnalyticsDestinationPrefix"
                                },
                                "OutputSchemaVersion": "V_1"
                            }
                        },
                        "Prefix": "AnalyticsConfigurationPrefix",
                        "TagFilters": [
                            {
                                "Key": "AnalyticsTagKey",
                                "Value": "AnalyticsTagValue"
                            }
                        ]
                    }
                ],
                "InventoryConfigurations": [
                    {
                        "Id": "InventoryConfigurationId",
                        "Destination": {
                            "BucketArn": {
                                "Fn::GetAtt": [
                                    "Helper",
                                    "Arn"
                                ]
                            },
                            "Format": "CSV",
                            "Prefix": "InventoryDestinationPrefix"
                        },
                        "Enabled": true,
                        "IncludedObjectVersions": "Current",
                        "Prefix": "InventoryConfigurationPrefix",
                        "ScheduleFrequency": "Weekly"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: S3 Bucket with Inventory and Analytics Configurations
Resources:
  Helper:
    Type: 'AWS::S3::Bucket'
  S3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      AnalyticsConfigurations:
        - Id: AnalyticsConfigurationId
          StorageClassAnalysis:
            DataExport:
              Destination:
                BucketArn: !GetAtt
                  - Helper
                  - Arn
                Format: CSV
                Prefix: AnalyticsDestinationPrefix
              OutputSchemaVersion: V_1
          Prefix: AnalyticsConfigurationPrefix
          TagFilters:
            - Key: AnalyticsTagKey
              Value: AnalyticsTagValue
      InventoryConfigurations:
        - Id: InventoryConfigurationId
          Destination:
            BucketArn: !GetAtt
              - Helper
              - Arn
            Format: CSV
            Prefix: InventoryDestinationPrefix
          Enabled: true
          IncludedObjectVersions: Current
          Prefix: InventoryConfigurationPrefix
          ScheduleFrequency: Weekly
```

## See also

- [Amazon S3 Template Snippets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-s3.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcConfiguration

AbortIncompleteMultipartUpload
