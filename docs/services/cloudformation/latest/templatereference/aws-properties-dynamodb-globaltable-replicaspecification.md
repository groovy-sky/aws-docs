This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable ReplicaSpecification

Defines settings specific to a single replica of a global table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContributorInsightsSpecification" : ContributorInsightsSpecification,
  "DeletionProtectionEnabled" : Boolean,
  "GlobalSecondaryIndexes" : [ ReplicaGlobalSecondaryIndexSpecification, ... ],
  "GlobalTableSettingsReplicationMode" : String,
  "KinesisStreamSpecification" : KinesisStreamSpecification,
  "PointInTimeRecoverySpecification" : PointInTimeRecoverySpecification,
  "ReadOnDemandThroughputSettings" : ReadOnDemandThroughputSettings,
  "ReadProvisionedThroughputSettings" : ReadProvisionedThroughputSettings,
  "Region" : String,
  "ReplicaStreamSpecification" : ReplicaStreamSpecification,
  "ResourcePolicy" : ResourcePolicy,
  "SSESpecification" : ReplicaSSESpecification,
  "TableClass" : String,
  "Tags" : [ Tag, ... ]
}

```

### YAML

```yaml

  ContributorInsightsSpecification:
    ContributorInsightsSpecification
  DeletionProtectionEnabled: Boolean
  GlobalSecondaryIndexes:
    - ReplicaGlobalSecondaryIndexSpecification
  GlobalTableSettingsReplicationMode: String
  KinesisStreamSpecification:
    KinesisStreamSpecification
  PointInTimeRecoverySpecification:
    PointInTimeRecoverySpecification
  ReadOnDemandThroughputSettings:
    ReadOnDemandThroughputSettings
  ReadProvisionedThroughputSettings:
    ReadProvisionedThroughputSettings
  Region: String
  ReplicaStreamSpecification:
    ReplicaStreamSpecification
  ResourcePolicy:
    ResourcePolicy
  SSESpecification:
    ReplicaSSESpecification
  TableClass: String
  Tags:
    - Tag

```

## Properties

`ContributorInsightsSpecification`

The settings used to enable or disable CloudWatch Contributor Insights for the specified
replica. When not specified, defaults to contributor insights disabled for the
replica.

_Required_: No

_Type_: [ContributorInsightsSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-globaltable-contributorinsightsspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeletionProtectionEnabled`

Determines if a replica is protected from deletion. When enabled, the table cannot be
deleted by any user or process. This setting is disabled by default. For more information,
see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the _Amazon DynamoDBDeveloper Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalSecondaryIndexes`

Defines additional settings for the global secondary indexes of this replica.

_Required_: No

_Type_: Array of [ReplicaGlobalSecondaryIndexSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-globaltable-replicaglobalsecondaryindexspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalTableSettingsReplicationMode`

The replication mode for global table settings across multiple accounts. In a multi-account global table,
you cannot make changes to a synchronized setting using CloudFormation. You need to use console or CLI to make
any settings changes.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisStreamSpecification`

Defines the Kinesis Data Streams configuration for the specified replica.

_Required_: No

_Type_: [KinesisStreamSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-globaltable-kinesisstreamspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PointInTimeRecoverySpecification`

The settings used to enable point in time recovery. When not specified, defaults to
point in time recovery disabled for the replica.

_Required_: No

_Type_: [PointInTimeRecoverySpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-globaltable-pointintimerecoveryspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadOnDemandThroughputSettings`

Sets read request settings for the replica table.

_Required_: No

_Type_: [ReadOnDemandThroughputSettings](aws-properties-dynamodb-globaltable-readondemandthroughputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadProvisionedThroughputSettings`

Defines read capacity settings for the replica table.

_Required_: No

_Type_: [ReadProvisionedThroughputSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-globaltable-readprovisionedthroughputsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The region in which this replica exists.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicaStreamSpecification`

Represents the DynamoDB Streams configuration for a global table replica.

_Required_: No

_Type_: [ReplicaStreamSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-globaltable-replicastreamspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourcePolicy`

A resource-based policy document that contains permissions to add to the specified
replica of a DynamoDB global table. Resource-based policies let you define access
permissions by specifying who has access to each resource, and the actions they are allowed
to perform on each resource.

In a CloudFormation template, you can provide the policy in JSON or YAML format
because CloudFormation converts YAML to JSON before submitting it to DynamoDB. For more information about resource-based policies, see [Using\
resource-based policies for DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/access-control-resource-based.html) and [Resource-based policy\
examples](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-examples.html).

_Required_: No

_Type_: [ResourcePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-globaltable-resourcepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SSESpecification`

Allows you to specify a customer-managed key for the replica. When using
customer-managed keys for server-side encryption, this property must have a value in all
replicas.

_Required_: No

_Type_: [ReplicaSSESpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-globaltable-replicassespecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableClass`

The table class of the specified table. Valid values are `STANDARD` and
`STANDARD_INFREQUENT_ACCESS`.

_Required_: No

_Type_: String

_Allowed values_: `STANDARD | STANDARD_INFREQUENT_ACCESS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this replica.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-globaltable-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ReplicaGlobalSecondaryIndexSpecification

ReplicaSSESpecification
