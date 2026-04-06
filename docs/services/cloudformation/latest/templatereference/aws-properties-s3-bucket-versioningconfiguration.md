This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket VersioningConfiguration

Describes the versioning state of an Amazon S3 bucket. For more information, see [PUT\
Bucket versioning](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTVersioningStatus.html) in the _Amazon S3 API Reference_.

Keep the following timing in mind when enabling, suspending, or transitioning between versioning states:

- **Enabling versioning** \- Changes may take up to 15 minutes to propagate across all AWS regions for full consistency.

- **Suspending versioning** \- Takes effect immediately with no propagation delay.

- **Transitioning between states** \- Any change from Suspended to Enabled has a 15-minute delay.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Status" : String
}

```

### YAML

```yaml

  Status: String

```

## Properties

`Status`

The versioning state of the bucket.

_Required_: Yes

_Type_: String

_Allowed values_: `Enabled | Suspended`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Examples

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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Transition

WebsiteConfiguration
