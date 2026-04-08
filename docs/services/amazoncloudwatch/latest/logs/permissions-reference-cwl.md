# CloudWatch Logs permissions reference

When you are setting up [Access control](auth-and-access-control-cwl.md#access-control-cwl) and writing permissions policies that you
can attach to an IAM identity (identity-based policies), you can use the following
table as a reference. The table lists each CloudWatch Logs API operation and the corresponding
actions for which you can grant permissions to perform the action. You specify the
actions in the policy's `Action` field. For the `Resource`
field, you can specify the ARN of a log group or log stream, or specify
`*` to represent all CloudWatch Logs resources.

You can use AWS-wide condition keys in your CloudWatch Logs policies to express
conditions. For a complete list of AWS-wide keys, see [AWS Global and\
IAM Condition Context Keys](../../../iam/latest/userguide/reference-policies-condition-keys.md) in the
_IAM User Guide_.

###### Note

To specify an action, use the `logs:` prefix followed by the API
operation name. For example: `logs:CreateLogGroup`,
`logs:CreateLogStream`, or `logs:*` (for all CloudWatch Logs
actions).

CloudWatch Logs API operations and required permissions for actionsCloudWatch Logs API operationsRequired permissions (API actions)

[CancelExportTask](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-cancelexporttask.md)

`logs:CancelExportTask`

Required to cancel a pending or running export task.

[CreateExportTask](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-createexporttask.md)

`logs:CreateExportTask`

Required to export data from a log group to an Amazon S3
bucket.

[CreateLogGroup](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-createloggroup.md)

`logs:CreateLogGroup`

Required to create a new log group.

[CreateLogStream](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-createlogstream.md)

`logs:CreateLogStream`

Required to create a new log stream in a log group.

[DeleteDestination](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletedestination.md)

`logs:DeleteDestination`

Required to delete a log destination and disables any
subscription filters to it.

[DeleteLogGroup](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deleteloggroup.md)

`logs:DeleteLogGroup`

Required to delete a log group and any associated archived log
events.

[DeleteLogStream](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletelogstream.md)

`logs:DeleteLogStream`

Required to delete a log stream and any associated archived
log events.

[DeleteMetricFilter](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletemetricfilter.md)

`logs:DeleteMetricFilter`

Required to delete a metric filter associated with a log
group.

[DeleteQueryDefinition](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletequerydefinition.md)

`logs:DeleteQueryDefinition`

Required to delete a saved query definition in CloudWatch Logs
Insights.

[DeleteResourcePolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deleteresourcepolicy.md)

`logs:DeleteResourcePolicy`

Required to delete a CloudWatch Logs resource policy.

[DeleteRetentionPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deleteretentionpolicy.md)

`logs:DeleteRetentionPolicy`

Required to delete a log group's retention policy.

[DeleteSubscriptionFilter](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletesubscriptionfilter.md)

`logs:DeleteSubscriptionFilter`

Required to delete the subscription filter associated with a
log group.

[DescribeDestinations](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describedestinations.md)

`logs:DescribeDestinations`

Required to view all destinations associated with the
account.

[DescribeExportTasks](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describeexporttasks.md)

`logs:DescribeExportTasks`

Required to view all export tasks associated with the
account.

[DescribeLogGroups](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describeloggroups.md)

`logs:DescribeLogGroups`

Required to view all log groups associated with the
account.

[DescribeLogStreams](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describelogstreams.md)

`logs:DescribeLogStreams`

Required to view all log streams associated with a log
group.

[DescribeMetricFilters](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describemetricfilters.md)

`logs:DescribeMetricFilters`

Required to view all metrics associated with a log
group.

[DescribeQueryDefinitions](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describequerydefinitions.md)

`logs:DescribeQueryDefinitions`

Required to see the list of saved query definitions in CloudWatch Logs
Insights.

[DescribeQueries](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describequeries.md)

`logs:DescribeQueries`

Required to see the list of CloudWatch Logs Insights queries that are
scheduled, executing, or have recently excecuted.

[DescribeResourcePolicies](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describeresourcepolicies.md)

`logs:DescribeResourcePolicies`

Required to view a list of CloudWatch Logs resource policies.

[DescribeSubscriptionFilters](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describesubscriptionfilters.md)

`logs:DescribeSubscriptionFilters`

Required to view all subscription filters associated with a
log group.

[FilterLogEvents](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-filterlogevents.md)

`logs:FilterLogEvents`

Required to sort log events by log group filter
pattern.

[GetLogEvents](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getlogevents.md)

`logs:GetLogEvents`

Required to retrieve log events from a log stream.

[GetLogGroupFields](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getloggroupfields.md)

`logs:GetLogGroupFields`

Required to retrieve the list of fields that are included in
the log events in a log group.

[GetLogRecord](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getlogrecord.md)

`logs:GetLogRecord`

Required to retrieve the details from a single log
event.

[GetLogObject](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getlogobject.md)

`logs:GetLogRecord`

Required to fetch the content of large portions of log events that have been ingested through the PutOpenTelemetryLogs API.

[GetQueryResults](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getqueryresults.md)

`logs:GetQueryResults`

Required to retrieve the results of CloudWatch Logs Insights
queries.

ListEntitiesForLogGroup

(CloudWatch console-only permission)

`logs:ListEntitiesForLogGroup`

Required to find the entities associated with a log group.
Required to explore related logs within the CloudWatch console.

ListLogGroupsForEntity

(CloudWatch console-only permission)

`logs:ListLogGroupsForEntity`

Required to find the log groups associated with an entity.
Required to explore related logs within the CloudWatch console.

[ListTagsLogGroup](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-listtagsloggroup.md)

`logs:ListTagsLogGroup`

Required to list the tags associated with a log group.

[ListLogGroups](../../../../reference/amazoncloudwatch/latest/apireference/api-api-listloggroups.md)

`logs:DescribeLogGroups`

Required to view all log groups associated with the account.

[PutDestination](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putdestination.md)

`logs:PutDestination`

Required to create or update a destination log stream (such as
an Kinesis stream).

[PutDestinationPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putdestinationpolicy.md)

`logs:PutDestinationPolicy`

Required to create or update an access policy associated with
an existing log destination.

[PutLogEvents](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putlogevents.md)

`logs:PutLogEvents`

Required to upload a batch of log events to a log
stream.

[PutMetricFilter](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putmetricfilter.md)

`logs:PutMetricFilter`

Required to create or update a metric filter and associate it
with a log group.

[PutQueryDefinition](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putquerydefinition.md)

`logs:PutQueryDefinition`

Required to save a query in CloudWatch Logs Insights, including
saved queries with parameters.

[PutResourcePolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putresourcepolicy.md)

`logs:PutResourcePolicy`

Required to create a CloudWatch Logs resource policy.

[PutRetentionPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putretentionpolicy.md)

`logs:PutRetentionPolicy`

Required to set the number of days to keep log events
(retention) in a log group.

[PutSubscriptionFilter](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putsubscriptionfilter.md)

`logs:PutSubscriptionFilter`

Required to create or update a subscription filter and
associate it with a log group.

[StartQuery](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-startquery.md)

`logs:StartQuery`

Required to start CloudWatch Logs Insights queries. To run a
saved query with parameters, you also need
`logs:DescribeQueryDefinitions`.

[StopQuery](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-stopquery.md)

`logs:StopQuery`

Required to stop a CloudWatch Logs Insights query that is in
progress.

[TagLogGroup](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-tagloggroup.md)

`logs:TagLogGroup`

Required to add or update log group tags.

[TestMetricFilter](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-testmetricfilter.md)

`logs:TestMetricFilter`

Required to test a filter pattern against a sampling of log
event messages.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Using identity-based policies (IAM policies)

Using service-linked roles

All content copied from https://docs.aws.amazon.com/.
