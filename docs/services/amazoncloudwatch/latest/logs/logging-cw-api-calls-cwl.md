# Logging CloudWatch Logs API and console operations in AWS CloudTrail

Amazon CloudWatch Logs is integrated with [AWS CloudTrail](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-user-guide.html), a service that provides a record of actions taken by a user, role, or an
AWS service. CloudTrail captures API calls for CloudWatch Logs as events. The calls captured include calls from the CloudWatch Logs console
and code calls to the CloudWatch Logs API operations. Using the information collected by CloudTrail, you can
determine the request that was made to CloudWatch Logs, the IP address from which the request was
made, when it was made, and additional details.

Every event or log entry contains information about who generated the request. The identity
information helps you determine the following:

- Whether the request was made with root user or user credentials.

- Whether the request was made on behalf of an IAM Identity Center user.

- Whether the request was made with temporary security credentials for a role or federated
user.

- Whether the request was made by another AWS service.

CloudTrail is active in your AWS account when you create the account and you automatically have
access to the CloudTrail **Event history**. The CloudTrail **Event**
**history** provides a viewable, searchable, downloadable, and immutable record of the
past 90 days of recorded management events in an AWS Region. For more information, see [Working\
with CloudTrail Event history](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/view-cloudtrail-events.html) in the _AWS CloudTrail User Guide_. There are no CloudTrail
charges for viewing the **Event history**.

For an ongoing record of events in your AWS account past 90 days, create a trail or a
[CloudTrail\
Lake](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-lake.html) event data store.

**CloudTrail trails**

A _trail_ enables CloudTrail to deliver log files to an Amazon S3 bucket. All trails created using the AWS Management Console are multi-Region. You can create a single-Region or a multi-Region trail by using the AWS CLI. Creating a multi-Region trail is recommended because you capture activity in all AWS Regions in your account. If you create a single-Region trail, you can view only the events logged in the trail's AWS Region. For more information about trails, see [Creating a trail for your AWS account](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-create-and-update-a-trail.html) and [Creating a trail for an organization](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/creating-trail-organization.html) in the _AWS CloudTrail User Guide_.

You can deliver one copy of your ongoing management events to your Amazon S3 bucket at no charge from CloudTrail by creating a trail, however, there are Amazon S3 storage charges. For more information about CloudTrail pricing, see [AWS CloudTrail Pricing](https://aws.amazon.com/cloudtrail/pricing). For information about Amazon S3 pricing, see [Amazon S3 Pricing](https://aws.amazon.com/s3/pricing).

**CloudTrail Lake event data stores**

_CloudTrail Lake_ lets you run SQL-based queries on your events. CloudTrail Lake converts existing events in row-based JSON format to [Apache ORC](https://orc.apache.org/) format. ORC is a columnar storage format that is optimized for fast retrieval of data. Events are aggregated into _event data stores_, which are immutable collections of events based on criteria that you select by applying [advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-lake-concepts.html#adv-event-selectors). The selectors that you apply to an event data store control which events persist and are available for you to query. For more information about CloudTrail Lake, see [Working with AWS CloudTrail Lake](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-lake.html) in the _AWS CloudTrail User Guide_.

CloudTrail Lake event data stores and queries incur costs. When you create an event data store, you choose the [pricing option](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-lake-manage-costs.html#cloudtrail-lake-manage-costs-pricing-option) you want to use for the event data store. The pricing option determines the cost for ingesting and storing events, and the default and maximum retention period for the event data store. For more information about CloudTrail pricing, see [AWS CloudTrail Pricing](https://aws.amazon.com/cloudtrail/pricing).

CloudWatch Logs supports logging the following actions as events in CloudTrail log files:

- [AssociateKmsKey](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-associatekmskey.md)

- [CancelExportTask](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-cancelexporttask.md)

- [CreateDelivery](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-createdelivery.md)

- [CreateExportTask](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-createexporttask.md)

- [CreateLogAnomalyDetector](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-createloganomalydetector.md)

- [CreateLogGroup](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-createloggroup.md)

- [CreateLogStream](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-createlogstream.md)

- [DeleteAccountPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deleteaccountpolicy.md)

- [DeleteDataProtectionPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletedataprotectionpolicy.md)

- [DeleteDelivery](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletedelivery.md)

- [DeleteDeliveryDestination](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletedeliverydestination.md)

- [DeleteDeliveryDestinationPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletedeliverydestinationpolicy.md)

- [DeleteDeliverySource](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletedeliverysource.md)

- [DeleteDestination](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletedestination.md)

- [DeleteIndexPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deleteindexpolicy.md)

- [DeleteIntegration](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deleteintegration.md)

- [DeleteLogAnomalyDetector](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deleteloganomalydetector.md)

- [DeleteLogGroup](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deleteloggroup.md)

- [DeleteLogStream](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletelogstream.md)

- [DeleteMetricFilter](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletemetricfilter.md)

- [DeleteQueryDefinition](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletequerydefinition.md)

- [DeleteResourcePolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deleteresourcepolicy.md)

- [DeleteRetentionPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deleteretentionpolicy.md)

- [DeleteSubscriptionFilter](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletesubscriptionfilter.md)

- [DeleteTransformer](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-deletetransformer.md)

- [DescribeAccountPolicies](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describeaccountpolicies.md)

- [DescribeConfigurationTemplates](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describeconfigurationtemplates.md)

- [DescribeDeliveries](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describedeliveries.md)

- [DescribeDeliveryDestinations](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describedeliverydestinations.md)

- [DescribeDeliverySources](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describedeliverysources.md)

- [DescribeDestinations](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describedestinations.md)

- [DescribeExportTasks](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describeexporttasks.md)

- [DescribeFieldIndexes](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describefieldindexes.md)

- [DescribeIndexPolicies](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describeindexpolicies.md)

- [DescribeLogGroups](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describeloggroups.md)

- [DescribeLogStreams](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describelogstreams.md)

- [DescribeMetricFilters](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describemetricfilters.md)

- [DescribeQueries](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describequeries.md)

- [DescribeQueryDefinitions](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describequerydefinitions.md)

- [DescribeResourcePolicies](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describeresourcepolicies.md)

- [DescribeSubscriptionFilters](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-describesubscriptionfilters.md)

- [DisassociateKmsKey](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-disassociatekmskey.md)

- [FilterLogEvents](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-filterlogevents.md)

- [GetDataProtectionPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getdataprotectionpolicy.md)

- [GetDelivery](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getdelivery.md)

- [GetDeliveryDestination](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getdeliverydestination.md)

- [GetDeliveryDestinationPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getdeliverydestinationpolicy.md)

- [GetDeliverySource](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getdeliverysource.md)

- [GetIntegration](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getintegration.md)

- [GetLogAnomalyDetector](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getloganomalydetector.md)

- [GetLogEvents](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getlogevents.md)

- [GetLogGroupFields](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getloggroupfields.md)

- [GetLogRecord](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getlogrecord.md)

- [GetQueryResults](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getqueryresults.md)

- [GetTransformer](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-gettransformer.md)

- [ListAnomalies](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-listanomalies.md)

- [ListIntegrations](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-listintegrations.md)

- [ListLogAnomalyDetectors](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-listloganomalydetectors.md)

- [ListLogGroups](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-listloggroups.md)

- [ListLogGroupsForQuery](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-listloggroupsforquery.md)

- [ListTagsForResource](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-listtagsforresource.md)

- [ListTagsLogGroup](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-listtagsloggroup.md)

- [PutAccountPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putaccountpolicy.md)

- [PutDataProtectionPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putdataprotectionpolicy.md)

- [PutDeliveryDestination](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putdeliverydestination.md)

- [PutDeliveryDestinationPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putdeliverydestinationpolicy.md)

- [PutDeliverySource](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putdeliverysource.md)

- [PutDestination](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putdestination.md)

- [PutDestinationPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putdestinationpolicy.md)

- [PutIndexPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putindexpolicy.md)

- [PutIntegration](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putintegration.md)

- [PutMetricFilter](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putmetricfilter.md)

- [PutQueryDefinition](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putquerydefinition.md)

- [PutResourcePolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putresourcepolicy.md)

- [PutRetentionPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putretentionpolicy.md)

- [PutSubscriptionFilter](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putsubscriptionfilter.md)

- [PutTransformer](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-puttransformer.md)

- [StartLiveTail](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-startlivetail.md)

- [StartQuery](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-startquery.md)

- [StopQuery](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-stopquery.md)

- [TagResource](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-tagresource.md)

- [TestMetricFilter](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-testmetricfilter.md)

- [TestTransformer](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-testtransformer.md)

- [UntagResource](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-untagresource.md)

- [UpdateAnomaly](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-updateanomaly.md)

- [UpdateDeliveryConfiguration](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-updatedeliveryconfiguration.md)

- [UpdateLogAnomalyDetector](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-updateloganomalydetector.md)

Every event or log entry contains information about who generated the request. The
identity information helps you determine the following:

- Whether the request was made with root or IAM user credentials.

- Whether the request was made with temporary security credentials for a role or
federated user.

- Whether the request was made by another AWS service.

For more information, see the [CloudTrail userIdentity\
Element](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-event-reference-user-identity.html).

## Query generation information in CloudTrail

CloudTrail logging for Query generator console events is also supported. Query generator is currently supported for CloudWatch Logs Insights and
CloudWatch Metric Insights. In these CloudTrail events, the `eventSource` is `monitoring.amazonaws.com`.

The following example shows a
CloudTrail log entry that demonstrates the **GenerateQuery** action in CloudWatch Logs Insights.

```json

{
    "eventVersion": "1.09",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EX_PRINCIPAL_ID",
        "arn": "arn:aws:iam::123456789012:assumed-role/role_name",
        "accountId": "123456789012",
        "accessKeyId": "AKIAIOSFODNN7EXAMPLE",
        "sessionContext": {
            "sessionIssuer": {
                "type": "Role",
                "principalId": "EX_PRINCIPAL_ID",
                "arn": "arn:aws:iam::111222333444:role/Administrator",
                "accountId": "123456789012",
                "userName": "SAMPLE_NAME"
            },
            "attributes": {
                "creationDate": "2020-04-08T21:43:24Z",
                "mfaAuthenticated": "false"
            }
        }
    },
    "eventTime": "2020-04-08T23:06:30Z",
    "eventSource": "monitoring.amazonaws.com",
    "eventName": "GenerateQuery",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "127.0.0.1",
    "userAgent": "exampleUserAgent",
    "requestParameters": {
        "query_ask": "***",
        "query_type": "LogsInsights",
        "logs_insights": {
            "fields": "***",
            "log_group_names": ["yourloggroup"]
        },
        "include_description": true
    },
    "responseElements": null,
    "requestID": "2f56318c-cfbd-4b60-9d93-1234567890",
    "eventID": "52723fd9-4a54-478c-ac55-1234567890",
    "readOnly": true,
    "eventType": "AwsApiCall",
    "managementEvent": true,
    "recipientAccountId": "111122223333",
    "eventCategory": "Management"
}
```

## Understanding log file entries

A trail is a configuration that enables delivery of events as log files to an Amazon S3 bucket
that you specify. CloudTrail log files contain one or more log entries. An event represents a single
request from any source and includes information about the requested action, the date and time
of the action, request parameters, and so on. CloudTrail log files aren't an ordered stack trace of
the public API calls, so they don't appear in any specific order.

The following log file entry shows that a user called the CloudWatch Logs
**CreateExportTask** action.

```

{
        "eventVersion": "1.03",
        "userIdentity": {
            "type": "IAMUser",
            "principalId": "EX_PRINCIPAL_ID",
            "arn": "arn:aws:iam::123456789012:user/someuser",
            "accountId": "123456789012",
            "accessKeyId": "AKIAIOSFODNN7EXAMPLE",
            "userName": "someuser"
        },
        "eventTime": "2016-02-08T06:35:14Z",
        "eventSource": "logs.amazonaws.com",
        "eventName": "CreateExportTask",
        "awsRegion": "us-east-1",
        "sourceIPAddress": "127.0.0.1",
        "userAgent": "aws-sdk-ruby2/2.0.0.rc4 ruby/1.9.3 x86_64-linux Seahorse/0.1.0",
        "requestParameters": {
            "destination": "yourdestination",
            "logGroupName": "yourloggroup",
            "to": 123456789012,
            "from": 0,
            "taskName": "yourtask"
        },
        "responseElements": {
            "taskId": "15e5e534-9548-44ab-a221-64d9d2b27b9b"
        },
        "requestID": "1cd74c1c-ce2e-12e6-99a9-8dbb26bd06c9",
        "eventID": "fd072859-bd7c-4865-9e76-8e364e89307c",
        "eventType": "AwsApiCall",
        "apiVersion": "20140328",
        "recipientAccountId": "123456789012"
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Interface VPC endpoints

Monitoring usage with CloudWatch metrics
