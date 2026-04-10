This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket ReplicationTime

A container specifying S3 Replication Time Control (S3 RTC) related information, including whether S3 RTC is enabled and
the time when all objects and operations on objects must be replicated. Must be specified together with
a `Metrics` block.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Status" : String,
  "Time" : ReplicationTimeValue
}

```

### YAML

```yaml

  Status: String
  Time:
    ReplicationTimeValue

```

## Properties

`Status`

Specifies whether the replication time is enabled.

_Required_: Yes

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Time`

A container specifying the time by which replication should be complete for all objects and
operations on objects.

_Required_: Yes

_Type_: [ReplicationTimeValue](aws-properties-s3-bucket-replicationtimevalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Enable S3 Replication Time Control

The following example creates a replication configuration with S3 Replication Time Control
(S3 RTC) enabled. To use this example, replace `amzn-s3-demo-source-bucket`
with the name of your source bucket and replace `amzn-s3-demo-destination-bucket`
with the name of your destination bucket. Make sure to update the AWS Identity and Access Management
(IAM) role and the replication rule as needed.

#### JSON

```json

      {
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS CloudFormation Template for S3 Bucket Replication",
    "Resources": {
      "MyS3Bucket": {
        "Type": "AWS::S3::Bucket",
        "Properties": {
          "BucketName": "amzn-s3-demo-source-bucket",
          "VersioningConfiguration": {
            "Status": "Enabled"
          },
          "ReplicationConfiguration": {
            "Role": "arn:aws:iam::account:role/s3-replication-role",
            "Rules": [
              {
                "Id": "ReplicationRule1",
                "Status": "Enabled",
                "Filter": {
                  "Prefix": ""
                },
                "Destination": {
                  "Bucket": "arn:aws:s3:::amzn-s3-demo-destination-bucket",
                  "ReplicationTime": {
                    "Status": "Enabled",
                    "Time": {
                      "Minutes": 15
                    }
                  },
                  "Metrics": {
                    "Status": "Enabled",
                    "EventThreshold": {
                      "Minutes": 15
                    }
                  }
                },
                "Priority": 1,
                "DeleteMarkerReplication": {
                  "Status": "Enabled"
                },
                "SourceSelectionCriteria": {
                  "ReplicaModifications": {
                    "Status": "Disabled"
                  }
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

        AWSTemplateFormatVersion: '2010-09-09'
Description: 'AWS CloudFormation Template for S3 Bucket Replication'

Resources:
  MyS3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      BucketName: 'amzn-s3-demo-source-bucket'
      VersioningConfiguration:
        Status: 'Enabled'
      ReplicationConfiguration:
        Role: 'arn:aws:iam::account:role/s3-replication-role'
        Rules:
          - Id: 'ReplicationRule1'
            Status: 'Enabled'
            Filter:
              Prefix: ""
            Destination:
              Bucket: 'arn:aws:s3:::amzn-s3-demo-destination-bucket'
              ReplicationTime:
                  Status: Enabled
                  Time:
                    Minutes: 15
              Metrics:
                Status: Enabled
                EventThreshold:
                  Minutes: 15
            Priority: 1
            DeleteMarkerReplication:
              Status: Enabled
            SourceSelectionCriteria:
              ReplicaModifications:
                Status: Disabled

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationRuleFilter

ReplicationTimeValue

All content copied from https://docs.aws.amazon.com/.
