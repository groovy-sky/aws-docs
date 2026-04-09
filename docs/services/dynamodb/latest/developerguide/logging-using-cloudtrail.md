# Logging DynamoDB operations by using AWS CloudTrail

DynamoDB is integrated with AWS CloudTrail, a service that provides a record of actions taken by a
user, role, or an AWS service in DynamoDB. CloudTrail captures all API calls for DynamoDB as events.
The calls captured include calls from the DynamoDB console and code calls to the DynamoDB API
operations, using both PartiQL and the classic API. If you create a trail, you can enable
continuous delivery of CloudTrail events to an Amazon S3 bucket, including events for DynamoDB. If you
don't configure a trail, you can still view the most recent events in the CloudTrail console in
**Event history**. Using the information collected by CloudTrail, you can
determine the request that was made to DynamoDB, the IP address from which the request was
made, who made the request, when it was made, and additional details.

For robust monitoring and alerting, you can also integrate CloudTrail events with [Amazon CloudWatch Logs](../../../amazoncloudwatch/latest/logs/whatiscloudwatchlogs.md). To enhance your
analysis of DynamoDB service activity and identify changes in activities for an AWS account,
you can query AWS CloudTrail logs using [Amazon\
Athena](../../../athena/latest/ug/cloudtrail-logs.md). For example, you can use queries to identify trends and further isolate
activity by attributes such as source IP address or user.

To learn more about CloudTrail, including how to configure and enable it, see the [AWS CloudTrail User Guide](../../../awscloudtrail/latest/userguide.md).

###### Topics

- [DynamoDB information in CloudTrail](#service-name-info-in-cloudtrail)

- [Understanding DynamoDB log file entries](understanding-ddb-log-entries.md)

## DynamoDB information in CloudTrail

CloudTrail is enabled on your AWS account when you create the account. When supported
event activity occurs in DynamoDB, that activity is recorded in a CloudTrail event along with
other AWS service events in **Event history**. You can view, search,
and download recent events in your AWS account. For more information, see [Working with CloudTrail Event history](../../../awscloudtrail/latest/userguide/view-cloudtrail-events.md).

For an ongoing record of events in your AWS account, including events for DynamoDB,
create a trail. A _trail_ enables CloudTrail to deliver log files to an
Amazon S3 bucket. By default, when you create a trail in the console, the trail applies to
all AWS Regions. The trail logs events from all Regions in the AWS partition and
delivers the log files to the Amazon S3 bucket that you specify. Additionally, you can
configure other AWS services to further analyze and act upon the event data collected
in CloudTrail logs. For more information, see the following:

- [Overview\
for creating a trail](../../../awscloudtrail/latest/userguide/cloudtrail-create-and-update-a-trail.md)

- [CloudTrail supported services and integrations](../../../awscloudtrail/latest/userguide/cloudtrail-aws-service-specific-topics.md#cloudtrail-aws-service-specific-topics-integrations)

- [Configuring\
Amazon SNS notifications for CloudTrail](../../../awscloudtrail/latest/userguide/getting-notifications-top-level.md)

- [Receiving CloudTrail log files from multiple regions](../../../awscloudtrail/latest/userguide/receive-cloudtrail-log-files-from-multiple-regions.md) and [Receiving CloudTrail log files from multiple accounts](../../../awscloudtrail/latest/userguide/cloudtrail-receive-logs-from-multiple-accounts.md)

### Control plane events in CloudTrail

The following API actions are logged by default as events in CloudTrail files:

**Amazon DynamoDB**

- [CreateBackup](../../../../reference/amazondynamodb/latest/apireference/api-createbackup.md)

- [CreateGlobalTable](../../../../reference/amazondynamodb/latest/apireference/api-createglobaltable.md)

- [CreateTable](../../../../reference/amazondynamodb/latest/apireference/api-createtable.md)

- [DeleteBackup](../../../../reference/amazondynamodb/latest/apireference/api-deletebackup.md)

- [DeleteTable](../../../../reference/amazondynamodb/latest/apireference/api-deletetable.md)

- [DescribeBackup](../../../../reference/amazondynamodb/latest/apireference/api-describebackup.md)

- [DescribeContinuousBackups](../../../../reference/amazondynamodb/latest/apireference/api-describecontinuousbackups.md)

- [DescribeGlobalTable](../../../../reference/amazondynamodb/latest/apireference/api-describeglobaltable.md)

- [DescribeLimits](../../../../reference/amazondynamodb/latest/apireference/api-describelimits.md)

- [DescribeTable](../../../../reference/amazondynamodb/latest/apireference/api-describetable.md)

- [DescribeTimeToLive](../../../../reference/amazondynamodb/latest/apireference/api-describetimetolive.md)

- [ListBackups](../../../../reference/amazondynamodb/latest/apireference/api-listbackups.md)

- [ListTables](../../../../reference/amazondynamodb/latest/apireference/api-listtables.md)

- [ListTagsOfResource](../../../../reference/amazondynamodb/latest/apireference/api-listtagsofresource.md)

- [ListGlobalTables](../../../../reference/amazondynamodb/latest/apireference/api-listglobaltables.md)

- [RestoreTableFromBackup](../../../../reference/amazondynamodb/latest/apireference/api-restoretablefrombackup.md)

- [RestoreTableToPointInTime](../../../../reference/amazondynamodb/latest/apireference/api-restoretabletopointintime.md)

- [TagResource](../../../../reference/amazondynamodb/latest/apireference/api-tagresource.md)

- [UntagResource](../../../../reference/amazondynamodb/latest/apireference/api-untagresource.md)

- [UpdateGlobalTable](../../../../reference/amazondynamodb/latest/apireference/api-updateglobaltable.md)

- [UpdateTable](../../../../reference/amazondynamodb/latest/apireference/api-updatetable.md)

- [UpdateTimeToLive](../../../../reference/amazondynamodb/latest/apireference/api-updatetimetolive.md)

- [DescribeReservedCapacity](iam-policy-prevent-purchase-reserved-capacity.md)

- [DescribeReservedCapacityOfferings](iam-policy-prevent-purchase-reserved-capacity.md)

- [PurchaseReservedCapacityOfferings](iam-policy-prevent-purchase-reserved-capacity.md)

- [DescribeScalableTargets](../../../../reference/autoscaling/application/apireference/api-describescalabletargets.md)

- [RegisterScalableTarget](../../../../reference/autoscaling/application/apireference/api-registerscalabletarget.md)

**DynamoDB Streams**

- [DescribeStream](../../../../reference/amazondynamodb/latest/apireference/api-streams-describestream.md)

- [ListStreams](../../../../reference/amazondynamodb/latest/apireference/api-streams-liststreams.md)

**DynamoDB Accelerator (DAX)**

- [CreateCluster](../../../../reference/amazondynamodb/latest/apireference/api-dax-createcluster.md)

- [CreateParameterGroup](../../../../reference/amazondynamodb/latest/apireference/api-dax-createparametergroup.md)

- [CreateSubnetGroup](../../../../reference/amazondynamodb/latest/apireference/api-dax-createsubnetgroup.md)

- [DecreaseReplicationFactor](../../../../reference/amazondynamodb/latest/apireference/api-dax-decreasereplicationfactor.md)

- [DeleteCluster](../../../../reference/amazondynamodb/latest/apireference/api-dax-deletecluster.md)

- [DeleteParameterGroup](../../../../reference/amazondynamodb/latest/apireference/api-dax-deleteparametergroup.md)

- [DeleteSubnetGroup](../../../../reference/amazondynamodb/latest/apireference/api-dax-deletesubnetgroup.md)

- [DescribeClusters](../../../../reference/amazondynamodb/latest/apireference/api-dax-describeclusters.md)

- [DescribeDefaultParameters](../../../../reference/amazondynamodb/latest/apireference/api-dax-describedefaultparameters.md)

- [DescribeEvents](../../../../reference/amazondynamodb/latest/apireference/api-dax-describeevents.md)

- [DescribeParameterGroups](../../../../reference/amazondynamodb/latest/apireference/api-dax-describeparametergroups.md)

- [DescribeParameters](../../../../reference/amazondynamodb/latest/apireference/api-dax-describeparameters.md)

- [DescribeSubnetGroups](../../../../reference/amazondynamodb/latest/apireference/api-dax-describesubnetgroups.md)

- [IncreaseReplicationFactor](../../../../reference/amazondynamodb/latest/apireference/api-dax-increasereplicationfactor.md)

- [ListTags](../../../../reference/amazondynamodb/latest/apireference/api-dax-listtags.md)

- [RebootNode](../../../../reference/amazondynamodb/latest/apireference/api-dax-rebootnode.md)

- [TagResource](../../../../reference/amazondynamodb/latest/apireference/api-dax-tagresource.md)

- [UntagResource](../../../../reference/amazondynamodb/latest/apireference/api-dax-untagresource.md)

- [UpdateCluster](../../../../reference/amazondynamodb/latest/apireference/api-dax-updatecluster.md)

- [UpdateParameterGroup](../../../../reference/amazondynamodb/latest/apireference/api-dax-updateparametergroup.md)

- [UpdateSubnetGroup](../../../../reference/amazondynamodb/latest/apireference/api-dax-updatesubnetgroup.md)

### DynamoDB data plane events in CloudTrail

To enable logging of the following API actions in CloudTrail files, you'll need to
enable logging of data plane API activity in CloudTrail.
See [Logging data events for trails](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md) for more information.

Data plane events can be filtered by resource type, for granular control over which DynamoDB API calls you want to
selectively log and pay for in CloudTrail. For example, by specifying `AWS::DynamoDB::Stream` as a resource
type, you can log only calls to the DynamoDB streams APIs. For tables with streams enabled, the resource field in the
data plane event contains both `AWS::DynamoDB::Stream` and `AWS::DynamoDB::Table`. If you
specify `AWS::DynamoDB::Table` as a resource type, it will log both DynamoDB table and DynamoDB streams events
by default. You can add an additional [filter](../../../../reference/awscloudtrail/latest/apireference/api-advancedfieldselector.md) to exclude the streams
events, if you don't want the streams events to be logged. For more information, see [DataResource](../../../../reference/awscloudtrail/latest/apireference/api-dataresource.md) in the _AWS CloudTrail API Reference_.

**Amazon DynamoDB**

- [BatchExecuteStatement](../../../../reference/amazondynamodb/latest/apireference/api-batchexecutestatement.md)

- [BatchGetItem](../../../../reference/amazondynamodb/latest/apireference/api-batchgetitem.md)

- [BatchWriteItem](../../../../reference/amazondynamodb/latest/apireference/api-batchwriteitem.md)

- [DeleteItem](../../../../reference/amazondynamodb/latest/apireference/api-deleteitem.md)

- [ExecuteStatement](../../../../reference/amazondynamodb/latest/apireference/api-executestatement.md)

- [ExecuteTransaction](../../../../reference/amazondynamodb/latest/apireference/api-executetransaction.md)

- [GetItem](../../../../reference/amazondynamodb/latest/apireference/api-getitem.md)

- [PutItem](../../../../reference/amazondynamodb/latest/apireference/api-putitem.md)

- [Query](../../../../reference/amazondynamodb/latest/apireference/api-query.md)

- [Scan](../../../../reference/amazondynamodb/latest/apireference/api-scan.md)

- [TransactGetItems](../../../../reference/amazondynamodb/latest/apireference/api-transactgetitems.md)

- [TransactWriteItems](../../../../reference/amazondynamodb/latest/apireference/api-transactwriteitems.md)

- [UpdateItem](../../../../reference/amazondynamodb/latest/apireference/api-updateitem.md)

###### Note

DynamoDB Time to Live data plane actions are not logged by CloudTrail

**DynamoDB Streams**

- [GetRecords](../../../../reference/amazondynamodb/latest/apireference/api-streams-getrecords.md)

- [GetShardIterator](../../../../reference/amazondynamodb/latest/apireference/api-streams-getsharditerator.md)

###### Important

When you log `GetRecords` data events, you might see `GetRecords` calls from DynamoDB internal operations, such as global tables replication. Although you are not charged by DynamoDB for these `GetRecords` calls, you are charged by CloudTrail for the data event logs. This might result in unexpected CloudTrail charges.

To avoid unexpected CloudTrail charges, do one of the following:

- Use the **Exclude AWS service-initiated events** log selector template.

- Add an advanced event selector filter with `userIdentity.arn` set to `NotStartsWith` `AWSServiceRoleFor`.

For more information, see [Logging data events](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md) in the AWS CloudTrail User Guide.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Creating CloudWatch alarms in DynamoDB

Understanding DynamoDB log file entries

All content copied from https://docs.aws.amazon.com/.
