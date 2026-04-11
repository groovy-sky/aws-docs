This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::DataLake ReplicationConfiguration

Provides replication configuration details for objects stored in the Amazon Security Lake data lake.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Regions" : [ String, ... ],
  "RoleArn" : String
}

```

### YAML

```yaml

  Regions:
    - String
  RoleArn: String

```

## Properties

`Regions`

Specifies one or more centralized rollup Regions. The AWS Region specified in the
region parameter of the `CreateDataLake` or
`UpdateDataLake` operations contributes data to the
rollup Region or Regions specified in this parameter.

Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets. S3 buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can replicate objects to a single destination bucket or to multiple destination buckets. The destination buckets can be in different Regions or within the same Region as the source bucket.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

Replication settings for the Amazon S3 buckets. This parameter uses the AWS Identity and Access Management (IAM) role you created that is managed by Security Lake, to ensure the replication setting is correct.

_Required_: No

_Type_: String

_Pattern_: `^arn:.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LifecycleConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
