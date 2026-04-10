This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable ReplicaGlobalSecondaryIndexSpecification

Represents the properties of a global secondary index that can be set on a per-replica
basis.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContributorInsightsSpecification" : ContributorInsightsSpecification,
  "IndexName" : String,
  "ReadOnDemandThroughputSettings" : ReadOnDemandThroughputSettings,
  "ReadProvisionedThroughputSettings" : ReadProvisionedThroughputSettings
}

```

### YAML

```yaml

  ContributorInsightsSpecification:
    ContributorInsightsSpecification
  IndexName: String
  ReadOnDemandThroughputSettings:
    ReadOnDemandThroughputSettings
  ReadProvisionedThroughputSettings:
    ReadProvisionedThroughputSettings

```

## Properties

`ContributorInsightsSpecification`

Updates the status for contributor insights for a specific table or index. CloudWatch
Contributor Insights for DynamoDB graphs display the partition key and (if applicable)
sort key of frequently accessed items and frequently throttled items in plaintext. If
you require the use of AWS Key Management Service (KMS) to encrypt this
table’s partition key and sort key data with an AWS managed key or
customer managed key, you should not enable CloudWatch Contributor Insights for DynamoDB
for this table.

_Required_: No

_Type_: [ContributorInsightsSpecification](aws-properties-dynamodb-globaltable-contributorinsightsspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexName`

The name of the global secondary index. The name must be unique among all other indexes
on this table.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReadOnDemandThroughputSettings`

Sets the read request settings for a replica global secondary index. You can only
specify this setting if your resource uses the `PAY_PER_REQUEST` `BillingMode`.

_Required_: No

_Type_: [ReadOnDemandThroughputSettings](aws-properties-dynamodb-globaltable-readondemandthroughputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadProvisionedThroughputSettings`

Allows you to specify the read capacity settings for a replica global secondary index
when the `BillingMode` is set to `PROVISIONED`.

_Required_: No

_Type_: [ReadProvisionedThroughputSettings](aws-properties-dynamodb-globaltable-readprovisionedthroughputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReadProvisionedThroughputSettings

ReplicaSpecification

All content copied from https://docs.aws.amazon.com/.
