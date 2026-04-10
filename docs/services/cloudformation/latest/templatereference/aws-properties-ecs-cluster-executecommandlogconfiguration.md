This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Cluster ExecuteCommandLogConfiguration

The log configuration for the results of the execute command actions. The logs can be
sent to CloudWatch Logs or an Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchEncryptionEnabled" : Boolean,
  "CloudWatchLogGroupName" : String,
  "S3BucketName" : String,
  "S3EncryptionEnabled" : Boolean,
  "S3KeyPrefix" : String
}

```

### YAML

```yaml

  CloudWatchEncryptionEnabled: Boolean
  CloudWatchLogGroupName: String
  S3BucketName: String
  S3EncryptionEnabled: Boolean
  S3KeyPrefix: String

```

## Properties

`CloudWatchEncryptionEnabled`

Determines whether to use encryption on the CloudWatch logs. If not specified,
encryption will be off.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchLogGroupName`

The name of the CloudWatch log group to send logs to.

###### Note

The CloudWatch log group must already be created.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketName`

The name of the S3 bucket to send logs to.

###### Note

The S3 bucket must already be created.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3EncryptionEnabled`

Determines whether to use encryption on the S3 logs. If not specified, encryption is
not used.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3KeyPrefix`

An optional folder in the S3 bucket to place logs in.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Define a cluster with the AWS Fargate capacity providers and\
a default capacity provider strategy defined](../userguide/aws-resource-ecs-cluster.md#aws-resource-ecs-cluster--examples--Define_a_cluster_with_the__capacity_providers_and_a_default_capacity_provider_strategy_defined)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExecuteCommandConfiguration

ManagedStorageConfiguration

All content copied from https://docs.aws.amazon.com/.
