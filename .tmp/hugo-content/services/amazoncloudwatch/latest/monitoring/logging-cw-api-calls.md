# Logging Amazon CloudWatch API and console operations with AWS CloudTrail

Amazon CloudWatch, CloudWatch Synthetics, CloudWatch RUM, Amazon Q Developer operational investigations, Network Flow Monitor, and Internet Monitor are integrated with AWS CloudTrail, a service that provides a record

of actions taken by a user, role, or an AWS service. CloudTrail captures API calls made by or on
behalf of your AWS account. The captured calls include calls from the CloudWatch console and code calls
to CloudWatch API operations. Using the information collected by CloudTrail, you can
determine the request that was made to CloudWatch, the IP address from which the request was
made, when it was made, and additional details.

Every event or log entry contains information about who generated the request. The
identity information helps you determine the following:

- Whether the request was made with root user or user credentials.

- Whether the request was made on behalf of an IAM Identity Center user.

- Whether the request was made with temporary security credentials for a role or federated
user.

- Whether the request was made by another AWS service.

CloudTrail is active in your AWS account when you create the account and you automatically have
access to the CloudTrail **Event history**. The CloudTrail **Event**
**history** provides a viewable, searchable, downloadable, and immutable record of the
past 90 days of recorded management events in an AWS Region. For more information, see [Working\
with CloudTrail Event history](../../../awscloudtrail/latest/userguide/view-cloudtrail-events.md) in the _AWS CloudTrail User Guide_. There are no CloudTrail
charges for viewing the **Event history**.

For an ongoing record of events in your AWS account past 90 days, create a trail or a
[CloudTrail\
Lake](../../../awscloudtrail/latest/userguide/cloudtrail-lake.md) event data store.

**CloudTrail trails**

A _trail_ enables CloudTrail to deliver log files to an Amazon S3 bucket. All trails created using the AWS Management Console are multi-Region. You can create a single-Region or a multi-Region trail by using the AWS CLI. Creating a multi-Region trail is recommended because you capture activity in all AWS Regions in your account. If you create a single-Region trail, you can view only the events logged in the trail's AWS Region. For more information about trails, see [Creating a trail for your AWS account](../../../awscloudtrail/latest/userguide/cloudtrail-create-and-update-a-trail.md) and [Creating a trail for an organization](../../../awscloudtrail/latest/userguide/creating-trail-organization.md) in the _AWS CloudTrail User Guide_.

You can deliver one copy of your ongoing management events to your Amazon S3 bucket at no charge from CloudTrail by creating a trail, however, there are Amazon S3 storage charges. For more information about CloudTrail pricing, see [AWS CloudTrail Pricing](https://aws.amazon.com/cloudtrail/pricing). For information about Amazon S3 pricing, see [Amazon S3 Pricing](https://aws.amazon.com/s3/pricing).

**CloudTrail Lake event data stores**

_CloudTrail Lake_ lets you run SQL-based queries on your events. CloudTrail Lake converts existing events in row-based JSON format to [Apache ORC](https://orc.apache.org/) format. ORC is a columnar storage format that is optimized for fast retrieval of data. Events are aggregated into _event data stores_, which are immutable collections of events based on criteria that you select by applying [advanced event selectors](../../../awscloudtrail/latest/userguide/cloudtrail-lake-concepts.md#adv-event-selectors). The selectors that you apply to an event data store control which events persist and are available for you to query. For more information about CloudTrail Lake, see [Working with AWS CloudTrail Lake](../../../awscloudtrail/latest/userguide/cloudtrail-lake.md) in the _AWS CloudTrail User Guide_.

CloudTrail Lake event data stores and queries incur costs. When you create an event data store, you choose the [pricing option](../../../awscloudtrail/latest/userguide/cloudtrail-lake-manage-costs.md#cloudtrail-lake-manage-costs-pricing-option) you want to use for the event data store. The pricing option determines the cost for ingesting and storing events, and the default and maximum retention period for the event data store. For more information about CloudTrail pricing, see [AWS CloudTrail Pricing](https://aws.amazon.com/cloudtrail/pricing).

###### Note

For information about CloudWatch Logs API calls that are logged in CloudTrail, see
[CloudWatch Logs information in CloudTrail](../logs/logging-cw-api-calls-cwl.md#cwl_info_in_ct).

###### Topics

- [CloudWatch information in CloudTrail](#cw_info_in_ct)

- [CloudWatch data events in CloudTrail](#CloudWatch-data-plane-events)

- [Query generation information in CloudTrail](#cwl_query-generation-cloudtrail)

- [Amazon Q Developer operational investigations events in CloudTrail](#Q-Developer-Investigations-Cloudtrail)

- [Network Flow Monitor in CloudTrail](#CloudWatch-NetworkFlowMonitor-info-in-ct)

- [Network Flow Monitor data plane events in CloudTrail](#CloudWatch-NetworkFlowMonitor-data-plane-events)

- [Internet Monitor in CloudTrail](#cw_im_info_in_ct)

- [CloudWatch Synthetics information in CloudTrail](#cw_synthetics_info_in_ct)

- [CloudWatch RUM information in CloudTrail](#RUM-CloudTrail)

- [CloudWatch RUM data plane events in CloudTrail](#RUM-data-plane-events)

- [Network Synthetic Monitor information in CloudTrail](#cw_network_synthetic_monitor_info_in_ct)

- [CloudWatch Observability Access Manager information in CloudTrail](#cw_observability_access_manager_info_in_ct)

- [CloudWatch Observability Admin information in CloudTrail](#cw_observability_admin_info_in_ct)

- [CloudWatch Application Signals information in CloudTrail](#cw_application_signals_info_in_ct)

- [CloudWatch Application Insights information in CloudTrail](#cw_application_insights_info_in_ct)

## CloudWatch information in CloudTrail

CloudWatch supports logging the following actions as events in CloudTrail log files:

- [DeleteAlarmActions](../../../../reference/amazoncloudwatch/latest/apireference/api-deletealarmactions.md)

- [DeleteAnomalyDetector](../../../../reference/amazoncloudwatch/latest/apireference/api-deleteanomalydetector.md)

- [DeleteDashboards](../../../../reference/amazoncloudwatch/latest/apireference/api-deletedashboards.md)

- [DeleteInsightRules](../../../../reference/amazoncloudwatch/latest/apireference/api-deleteinsightrules.md)

- [DeleteMetricStream](../../../../reference/amazoncloudwatch/latest/apireference/api-deletemetricstream.md)

- [DescribeAlarmHistory](../../../../reference/amazoncloudwatch/latest/apireference/api-describealarmhistory.md)

- [DescribeAlarms](../../../../reference/amazoncloudwatch/latest/apireference/api-describealarms.md)

- [DescribeAlarmsForMetric](../../../../reference/amazoncloudwatch/latest/apireference/api-describealarmsformetric.md)

- [DescribeAnomalyDetectors](../../../../reference/amazoncloudwatch/latest/apireference/api-describeanomalydetectors.md)

- [DescribeInsightRules](../../../../reference/amazoncloudwatch/latest/apireference/api-describeinsightrules.md)

- [DisableAlarmActions](../../../../reference/amazoncloudwatch/latest/apireference/api-disablealarmactions.md)

- [DisableInsightRules](../../../../reference/amazoncloudwatch/latest/apireference/api-disableinsightrules.md)

- [EnableAlarmActions](../../../../reference/amazoncloudwatch/latest/apireference/api-enablealarmactions.md)

- [EnableInsightRules](../../../../reference/amazoncloudwatch/latest/apireference/api-enableinsightrules.md)

- [GetDashboard](../../../../reference/amazoncloudwatch/latest/apireference/api-getdashboard.md)

- [GetInsightRuleReport](../../../../reference/amazoncloudwatch/latest/apireference/api-getinsightrulereport.md)

- [GetMetricStream](../../../../reference/amazoncloudwatch/latest/apireference/api-getmetricstream.md)

- [ListDashboards](../../../../reference/amazoncloudwatch/latest/apireference/api-listdashboards.md)

- [ListManagedInsightRules](../../../../reference/amazoncloudwatch/latest/apireference/api-listmanagedinsightrules.md)

- [ListMetricStreams](../../../../reference/amazoncloudwatch/latest/apireference/api-listmetricstreams.md)

- [ListTagsForResource](../../../../reference/amazoncloudwatch/latest/apireference/api-listtagsforresource.md)

- [PutAnomalyDetector](../../../../reference/amazoncloudwatch/latest/apireference/api-putanomalydetector.md)

- [PutCompositeAlarm](../../../../reference/amazoncloudwatch/latest/apireference/api-putcompositealarm.md)

- [PutDashboard](../../../../reference/amazoncloudwatch/latest/apireference/api-putdashboard.md)

- [PutInsightRule](../../../../reference/amazoncloudwatch/latest/apireference/api-putinsightrule.md)

- [PutManagedInsightRules](../../../../reference/amazoncloudwatch/latest/apireference/api-putmanagedinsightrules.md)

- [PutMetricAlarm](../../../../reference/amazoncloudwatch/latest/apireference/api-putmetricalarm.md)

- [PutMetricStream](../../../../reference/amazoncloudwatch/latest/apireference/api-putmetricstream.md)

- [SetAlarmState](../../../../reference/amazoncloudwatch/latest/apireference/api-setalarmstate.md)

- [StartMetricStreams](../../../../reference/amazoncloudwatch/latest/apireference/api-startmetricstreams.md)

- [StopMetricStreams](../../../../reference/amazoncloudwatch/latest/apireference/api-stopmetricstreams.md)

- [TagResource](../../../../reference/amazoncloudwatch/latest/apireference/api-tagresource.md)

- [UntagResource](../../../../reference/amazoncloudwatch/latest/apireference/api-untagresource.md)

### Example: CloudWatch log file entries

The following example shows a CloudTrail log entry that demonstrates the `PutMetricAlarm` action.

```json

{
    "Records": [{
        "eventVersion": "1.01",
        "userIdentity": {
            "type": "Root",
            "principalId": "EX_PRINCIPAL_ID",
            "arn": "arn:aws:iam::123456789012:root",
            "accountId": "123456789012",
            "accessKeyId": "EXAMPLE_KEY_ID"
        },
        "eventTime": "2014-03-23T21:50:34Z",
        "eventSource": "monitoring.amazonaws.com",
        "eventName": "PutMetricAlarm",
        "awsRegion": "us-east-1",
        "sourceIPAddress": "127.0.0.1",
        "userAgent": "aws-sdk-ruby2/2.0.0.rc4 ruby/1.9.3 x86_64-linux Seahorse/0.1.0",
        "requestParameters": {
            "threshold": 50.0,
            "period": 60,
            "metricName": "CloudTrail Test",
            "evaluationPeriods": 3,
            "comparisonOperator": "GreaterThanThreshold",
            "namespace": "AWS/CloudWatch",
            "alarmName": "CloudTrail Test Alarm",
            "statistic": "Sum"
        },
        "responseElements": null,
        "requestID": "29184022-b2d5-11e3-a63d-9b463e6d0ff0",
        "eventID": "b096d5b7-dcf2-4399-998b-5a53eca76a27"
    },
    ..additional entries
  ]
  }
```

The following log file entry shows that a user called the CloudWatch Events
`PutRule` action.

```nohighlight

{
         "eventVersion":"1.03",
         "userIdentity":{
            "type":"Root",
            "principalId":"123456789012",
            "arn":"arn:aws:iam::123456789012:root",
            "accountId":"123456789012",
            "accessKeyId":"AWS_ACCESS_KEY_ID_REDACTED",
            "sessionContext":{
               "attributes":{
                  "mfaAuthenticated":"false",
                  "creationDate":"2015-11-17T23:56:15Z"
               }
            }
         },
         "eventTime":"2015-11-18T00:11:28Z",
         "eventSource":"events.amazonaws.com",
         "eventName":"PutRule",
         "awsRegion":"us-east-1",
         "sourceIPAddress":"AWS Internal",
         "userAgent":"AWS CloudWatch Console",
         "requestParameters":{
            "description":"",
            "name":"cttest2",
            "state":"ENABLED",
            "eventPattern":"{\"source\":[\"aws.ec2\"],\"detail-type\":[\"EC2 Instance State-change Notification\"]}",
            "scheduleExpression":""
         },
         "responseElements":{
            "ruleArn":"arn:aws:events:us-east-1:123456789012:rule/cttest2"
         },
         "requestID":"e9caf887-8d88-11e5-a331-3332aa445952",
         "eventID":"49d14f36-6450-44a5-a501-b0fdcdfaeb98",
         "eventType":"AwsApiCall",
         "apiVersion":"2015-10-07",
         "recipientAccountId":"123456789012"
}
```

The following log file entry shows that a user called the CloudWatch Logs
`CreateExportTask` action.

```

{
        "eventVersion": "1.03",
        "userIdentity": {
            "type": "IAMUser",
            "principalId": "EX_PRINCIPAL_ID",
            "arn": "arn:aws:iam::123456789012:user/someuser",
            "accountId": "123456789012",
            "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
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

## CloudWatch data events in CloudTrail

CloudTrail can capture API activities related to the CloudWatch data plane operations on metrics [GetMetricData](../../../../reference/amazoncloudwatch/latest/apireference/api-getmetricdata.md), [GetMetricWidgetImage](../../../../reference/amazoncloudwatch/latest/apireference/api-getmetricwidgetimage.md), [PutMetricData](../../../../reference/amazoncloudwatch/latest/apireference/api-putmetricdata.md), [GetMetricStatistics](../../../../reference/amazoncloudwatch/latest/apireference/api-getmetricstatistics.md), and [ListMetrics](../../../../reference/amazoncloudwatch/latest/apireference/api-listmetrics.md) APIs.

[Data events](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md#logging-data-events), also known as data plane operations, give you insight into the resource operations performed
on or within a resource. Data events are often high-volume activities.

By default, CloudTrail doesn’t log
data events. The CloudTrail **Event history** doesn't record data events.

Additional charges apply for data events. For more information about CloudTrail pricing, see
[AWS CloudTrail Pricing](https://aws.amazon.com/cloudtrail/pricing).

You can log data events for the CloudWatch resource types by using the CloudTrail console, AWS CLI,
or CloudTrail API operations. For more information about how to log data events, see [Logging data events with the AWS Management Console](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md#logging-data-events-console) and [Logging data events with the AWS Command Line Interface](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md#creating-data-event-selectors-with-the-AWS-CLI) in the
_AWS CloudTrail User Guide_.

Data plane events can be filtered by resource type. Because there are additional costs for using
data events in CloudTrail, filtering by resource allows you to have more control over what you log and pay for.

Using the information that CloudTrail collects, you can identify any of the metric APIs, the IP address of the requester, the requester's identity, and the date and time of the
request. Logging the **GetMetricData**, **GetMetricWidgetImage**, **PutMetricData**, **GetMetricStatistics**, and **ListMetrics** APIs using CloudTrail helps you enable operational and
risk auditing, governance, and compliance of your AWS account.

###### Note

When you view the **GetMetricData** events in CloudTrail, you might see more calls than the calls that you initiated.
This is because CloudWatch logs events to CloudTrail for **GetMetricData** actions that are initiated by internal components. For example, you'll
see **GetMetricData** calls initiated by CloudWatch dashboards to refresh widget data, and **GetMetricData** calls initiated by a
monitoring account to retrieve data from a source account,
in cross-account observability. These internally-initiated calls don't incur CloudWatch charges, but they do count toward the number of events
logged in CloudTrail, and CloudTrail charges according to the number of events logged.

The following is an example of a CloudTrail event for a **GetMetricData** operation.

```json

{
    "eventVersion": "1.09",
    "userIdentity": {
        "type": "IAMUser",
        "principalId": "AWS_ACCESS_KEY_ID_REDACTED",
        "arn": "arn:aws:iam::111122223333:user/admin",
        "accountId": "111122223333",
        "accessKeyId": "EXAMPLE1234567890",
        "userName": "admin"
    },
    "eventTime": "2024-05-08T16:20:34Z",
    "eventSource": "monitoring.amazonaws.com",
    "eventName": "GetMetricData",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "99.45.3.7",
    "userAgent": "aws-cli/2.13.23 Python/3.11.5 Darwin/23.4.0 exe/x86_64 prompt/off command/cloudwatch.get-metric-data",
    "requestParameters": {
        "metricDataQueries": [{
                "id": "e1",
                "expression": "m1 / m2",
                "label": "ErrorRate"
            },
            {
                "id": "m1",
                "metricStat": {
                    "metric": {
                        "namespace": "CWAgent",
                        "metricName": "disk_used_percent",
                        "dimensions": [{
                            "name": "LoadBalancerName",
                            "value": "EXAMPLE4623a5cb6a7384c5229"
                        }]
                    },
                    "period": 300,
                    "stat": "Sum",
                    "unit": "Count"
                },
                "returnData": false
            },
            {
                "id": "m2",
                "metricStat": {
                    "metric": {
                        "namespace": "CWAgent",
                        "metricName": "disk_used_percent",
                        "dimensions": [{
                            "name": "LoadBalancerName",
                            "value": "EXAMPLE4623a5cb6a7384c5229"
                        }]
                    },
                    "period": 300,
                    "stat": "Sum"
                },
                "returnData": true
            }
        ],
        "startTime": "Apr 19, 2024, 4:00:00 AM",
        "endTime": "May 8, 2024, 4:30:00 AM"
    },
    "responseElements": null,
    "requestID": "EXAMPLE-57ac-47d5-938c-f5917c6799d5",
    "eventID": "EXAMPLE-211c-404b-b13d-36d93c8b4fbf",
    "readOnly": true,
    "resources": [{
        "type": "AWS::CloudWatch::Metric"
    }],
    "eventType": "AwsApiCall",
    "managementEvent": false,
    "recipientAccountId": "111122223333",
    "eventCategory": "Data",
    "tlsDetails": {
        "tlsVersion": "TLSv1.3",
        "cipherSuite": "TLS_AES_128_GCM_SHA256",
        "clientProvidedHostHeader": "monitoring.us-east-1.amazonaws.com"
    }
}
```

The following is an example of a CloudTrail event for a **PutMetricData** operation.

```json

{
      "eventVersion": "1.11",
      "userIdentity": {
        "type": "AssumedRole",
        "principalId": "111122223333:example.amazon.com",
        "arn": "arn:aws:sts::111122223333:assumed-role/cloudwatch.full.access/example.amazon.com",
        "accountId": "111122223333",
        "accessKeyId": "EXAMPLE1234567890",
        "sessionContext": {
          "sessionIssuer": {
            "type": "Role",
            "principalId": "AWS_ACCESS_KEY_ID_REDACTED",
            "arn": "arn:aws:iam::111122223333:role/cloudwatch.full.access",
            "accountId": "111122223333",
            "userName": "cloudwatch.full.access"
          },
          "attributes": {
            "creationDate": "2025-06-19T23:19:50Z",
            "mfaAuthenticated": "false"
          }
        }
      },
      "eventTime": "2025-06-19T23:51:04Z",
      "eventSource": "monitoring.amazonaws.com",
      "eventName": "PutMetricData",
      "awsRegion": "us-east-1",
      "sourceIPAddress": "AWS Internal",
      "userAgent": "AWS Internal",
      "requestParameters": {
        "namespace": "CloudTrailTests",
        "metricData": [
          {
            "metricName": "CloudTrailPutMetricDataTest",
            "dimensions": [
              {
                "name": "TestDimName",
                "value": "TestDimValue"
              }
            ]
          }
        ]
      },
      "responseElements": null,
      "requestID": "877db913-2620-4929-97f3-f3c93c6f689b",
      "eventID": "0c0c4516-75f4-4b27-8a83-213821a96a2b",
      "readOnly": false,
      "resources": [
        {
          "type": "AWS::CloudWatch::Metric"
        }
      ],
      "eventType": "AwsApiCall",
      "managementEvent": false,
      "recipientAccountId": "111122223333",
      "eventCategory": "Data",
      "tlsDetails": {
        "tlsVersion": "TLSv1.3",
        "cipherSuite": "TLS_AES_128_GCM_SHA256",
        "clientProvidedHostHeader": "monitoring.us-east-1.amazonaws.com"
      }
    }
```

## Query generation information in CloudTrail

CloudTrail logging for Query generator console events is also supported. Query generator is currently supported for
CloudWatch Metric Insights and CloudWatch Logs Insights. In these CloudTrail events, the `eventSource` is `monitoring.amazonaws.com`.

The following example shows a
CloudTrail log entry that demonstrates the **GenerateQuery** action in CloudWatch Metrics Insights.

```json

{
    "eventVersion": "1.09",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EX_PRINCIPAL_ID",
        "arn": "arn:aws:iam::123456789012:assumed-role/role_name",
        "accountId": "123456789012",
        "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
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
        "query_type": "MetricsInsights",
        "metrics_insights": {
            "aws_namespaces": [
                "AWS/S3",
                "AWS/Lambda",
                "AWS/DynamoDB"
            ]
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

## Amazon Q Developer operational investigations events in CloudTrail

Amazon Q Developer operational investigations supports logging the following actions as events in CloudTrail log files.

- [CreateInvestigationGroup](../../../../reference/cloudwatchinvestigations/latest/apireference/api-createinvestigationgroup.md)

- [GetInvestigationGroup](../../../../reference/cloudwatchinvestigations/latest/apireference/api-getinvestigationgroup.md)

- [DeleteInvestigationGroup](../../../../reference/cloudwatchinvestigations/latest/apireference/api-deleteinvestigationgroup.md)

- [ListInvestigationGroup](../../../../reference/cloudwatchinvestigations/latest/apireference/api-listinvestigationgroups.md)

- [PutInvestigationGroupPolicy](../../../../reference/cloudwatchinvestigations/latest/apireference/api-putinvestigationgrouppolicy.md)

- [DeleteInvestigationGroupPolicy](../../../../reference/cloudwatchinvestigations/latest/apireference/api-deleteinvestigationgrouppolicy.md)

- [ListTagsForResource](../../../../reference/cloudwatchinvestigations/latest/apireference/api-listtagsforresource.md)

- [GetInvestigationGroupPolicy](../../../../reference/cloudwatchinvestigations/latest/apireference/api-getinvestigationgrouppolicy.md)

- [TagResource](../../../../reference/cloudwatchinvestigations/latest/apireference/api-tagresource.md)

- [UntagResource](../../../../reference/cloudwatchinvestigations/latest/apireference/api-untagresource.md)

- [UpdateInvestigationGroup](../../../../reference/cloudwatchinvestigations/latest/apireference/api-updateinvestigationgroup.md)

### Example: Amazon Q Developer operational investigations log file entries

The following example shows a Amazon Q Developer operational investigations log entry that demonstrates the
`CreateInvestigationGroup` action.

```json

{
	"eventVersion": "1.09",
	"userIdentity": {
		"type": "AssumedRole",
		"principalId": "EX_PRINCIPAL_ID",
		"arn": "arn:aws:iam::123456789012:assumed-role/role_name",
		"accountId": "123456789012",
		"accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
		"sessionContext": {
			"sessionIssuer": {
				"type": "Role",
				"principalId": "EX_PRINCIPAL_ID",
				"arn": "arn:aws:iam::123456789012:role/role_name",
				"accountId": "123456789012",
				"userName": "SAMPLE_NAME"
			},
			"attributes": {
				"creationDate": "2024-10-30T18:42:05Z",
				"mfaAuthenticated": "false"
			}
		}
	},
	"eventTime": "2024-10-30T18:48:26Z",
	"eventSource": "aiops.amazonaws.com",
	"eventName": "CreateInvestigationGroup",
	"awsRegion": "us-east-1",
	"sourceIPAddress": "127.0.0.1",
	"userAgent": "exampleUserAgent",
	"requestParameters": {
		"name": "exampleName",
		"roleArn": "arn:aws:iam::123456789012:role/role_name"	},
	"responseElements": {
		"arn": "arn:aws:aiops:us-east-1:123456789012:investigation-group/021345abcdef67890"
	},
	"requestId": "e9caf887-8d88-11e5-a331-3332aa445952",
	"requestId": "49d14f36-6450-44a5-a501-b0fdcdfaeb98",
	"readOnly": false,
	"eventType": "AwsApiCall",
	"managementEvent": true,
	"recipientAccountId": "123456789012",
	"eventCategory": "Management"
}
```

The following example shows a Amazon Q Developer operational investigations log entry that demonstrates the
`CreateInvestigationEvent` action.

```

{
		"eventVersion": "1.09",
		"userIdentity": {
			"type": "AssumedRole",
			"principalId": "EX_PRINCIPAL_ID",
			"arn": "arn:aws:sts::123456789012:assumed-role/role_name",
			"accountId": "123456789012",
			"accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
			"sessionContext": {
				"sessionIssuer": {
					"type": "Role",
					"principalId": "EX_PRINCIPAL_ID",
					"arn": "arn:aws:iam::123456789012:role/role_name",
					"accountId": "123456789012",
					"userName": "SAMPLE_NAME"
				},
				"attributes": {
					"creationDate": "2024-10-30T16:17:49Z",
					"mfaAuthenticated": "false"
				}
			}
		},
		"eventTime": "2024-10-30T16:35:34Z",
		"eventSource": "aiops.amazonaws.com",
		"eventName": "CreateInvestigationEvent",
		"awsRegion": "us-east-1",
		"sourceIPAddress": "127.0.0.1",
		"userAgent": "exampleUserAgent",
		"requestParameters": {
			"identifier": "arn:aws:aiops:us-east-1:123456789012:investigation-group/021345abcdef67890",
			"investigationId": "bcdef01234567890",
			"clientToken": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
			"type": "METRIC_OBSERVATION",
			"body": "***"
		},
		"responseElements": {
			"investigationGroupArn": "arn:aws:aiops:us-east-1:123456789012:investigation-group/021345abcdef67890",
			"investigationId": "bcdef01234567890",
			"investigationEventId": "14567890abcdef0g"
		},
		"requestId": "e9caf887-8d88-11e5-a331-3332aa445952",
		"eventId": "49d14f36-6450-44a5-a501-b0fdcdfaeb98",
		"readOnly": false,
		"resources": [{
			"accountId": "123456789012",
			"type": "AWS::AIOps::InvestigationGroup",
			"ARN": "arn:aws:aiops:us-east-1:123456789012:investigation-group/021345abcdef67890"
		}],
		"eventType": "AwsApiCall",
		"managementEvent": false,
		"recipientAccountId": "123456789012",
		"eventCategory": "Data"
	}
```

The following example shows a Amazon Q Developer operational investigations log entry that demonstrates the
`UpdateInvestigationEvent` action.

```

{
	"eventVersion": "1.09",
	"userIdentity": {
		"type": "AssumedRole",
		"principalId": "EX_PRINCIPAL_ID",
		"arn": "arn:aws:sts::123456789012:assumed-role/role_name",
		"accountId": "123456789012",
		"accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
		"sessionContext": {
			"sessionIssuer": {
				"type": "Role",
				"principalId": "EX_PRINCIPAL_ID",
				"arn": "arn:aws:iam::123456789012:role/role_name",
				"accountId": "123456789012",
				"userName": "SAMPLE_NAME"
			},
			"attributes": {
				"creationDate": "2024-10-30T16:17:49Z",
				"mfaAuthenticated": "false"
			}
		}
	},
	"eventTime": "2024-10-30T16:24:48Z",
	"eventSource": "aiops.amazonaws.com",
	"eventName": "UpdateInvestigationEvent",
	"awsRegion": "us-east-1",
	"sourceIPAddress": "127.0.0.1",
	"userAgent": "exampleUserAgent",
	"requestParameters": {
		"identifier": "arn:aws:aiops:us-east-1:123456789012:investigation-group/021345abcdef67890",
		"investigationId": "bcdef01234567890",
		"investigationEventId": "14567890abcdef0g",
		"comment": "***"
	},
	"responseElements": null,
	"requestId": "e9caf887-8d88-11e5-a331-3332aa445952",
	"eventId": "49d14f36-6450-44a5-a501-b0fdcdfaeb98",
	"readOnly": false,
	"resources": [{
		"accountId": "123456789012",
		"type": "AWS::AIOps::InvestigationGroup",
		"ARN": "arn:aws:aiops:us-east-1:123456789012:investigation-group/021345abcdef67890"
	}],
	"eventType": "AwsApiCall",
	"managementEvent": false,
	"recipientAccountId": "123456789012",
	"eventCategory": "Data"
}
```

## Network Flow Monitor in CloudTrail

Network Flow Monitor supports logging the following actions as events in CloudTrail log files.

- [CreateMonitor](../../../../reference/networkflowmonitor/2-0/apireference/api-createmonitor.md)

- [CreateScope](../../../../reference/networkflowmonitor/2-0/apireference/api-createscope.md)

- [DeleteMonitor](../../../../reference/networkflowmonitor/2-0/apireference/api-deletemonitor.md)

- [DeleteScope](../../../../reference/networkflowmonitor/2-0/apireference/api-deletescope.md)

- [GetMonitor](../../../../reference/networkflowmonitor/2-0/apireference/api-getmonitor.md)

- [GetScope](../../../../reference/networkflowmonitor/2-0/apireference/api-getscope.md)

- [ListMonitors](../../../../reference/networkflowmonitor/2-0/apireference/api-listmonitors.md)

- [ListScopes](../../../../reference/networkflowmonitor/2-0/apireference/api-listscopes.md)

- [ListTagsForResource](../../../../reference/networkflowmonitor/2-0/apireference/api-listtagsforresource.md)

- [TagResource](../../../../reference/networkflowmonitor/2-0/apireference/api-tagresource.md)

- [UntagResource](../../../../reference/networkflowmonitor/2-0/apireference/api-untagresource.md)

- [UpdateMonitor](../../../../reference/networkflowmonitor/2-0/apireference/api-updatemonitor.md)

- [UpdateScope](../../../../reference/networkflowmonitor/2-0/apireference/api-updatescope.md)

### Example: Network Flow Monitor log file entry

The following example shows a Network Flow Monitor CloudTrail log file entry that demonstrates the
`CreateMonitor` action.

```json

{
    "eventVersion": "1.09",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EX_PRINCIPAL_ID",
        "arn": "arn:aws:iam::000000000000:assumed-role/role_name",
        "accountId": "123456789012",
        "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
        "sessionContext": {
            "sessionIssuer": {
                "type": "Role",
                "principalId": "EX_PRINCIPAL_ID",
                "arn": "arn:aws:iam::000000000000:role/Admin",
                "accountId": "123456789012",
                "userName": "SAMPLE_NAME"
            },
            "attributes": {
                "creationDate": "2024-11-03T15:43:27Z",
                "mfaAuthenticated": "false"
            }
        }
    },
    "eventTime": "2024-11-03T15:58:11Z",
    "eventSource": "networkflowmonitor.amazonaws.com",
    "eventName": "CreateMonitor",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "192.0.2.0",
    "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
    "requestParameters": {
        "MonitorName": "TestMonitor",
        "ClientToken": "33551db7-1618-4aab-cdef-EXAMPLE33333",
        "LocalResources": [
            {
                "Type": "AWS::EC2::Subnet",
                "Identifier": "subnet-cdef-EXAMPLEbbbbb"
            },
            {
                "Type": "AWS::EC2::Subnet",
                "Identifier": "subnet-cdef-EXAMPLEccccc"
            },
            {
                "Type": "AWS::EC2::Subnet",
                "Identifier": "subnet-cdef-EXAMPLEddddd"
            },
            {
                "Type": "AWS::EC2::Subnet",
                "Identifier": "subnet-cdef-EXAMPLEeeeee"
            },
            {
                "Type": "AWS::EC2::Subnet",
                "Identifier": "subnet-cdef-EXAMPLEfffff"
            },
            {
                "Type": "AWS::EC2::Subnet",
                "Identifier": "subnet-cdef-EXAMPLEggggg"
            }
        ]
    },
    "responseElements": {
        "Access-Control-Expose-Headers": "*",
        "MonitorArn": "arn:aws:networkflowmonitor:us-east-1:000000000000:monitor/TestMonitor",
        "MonitorName": "TestMonitor",
        "MonitorStatus": "ACTIVE"
    },
    "requestID": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
    "eventID": "a1b2c3d4-5678-90ab-cdef-EXAMPLEbbbbb",
    "readOnly": false,
    "eventType": "AwsApiCall",
    "managementEvent": true,
    "recipientAccountId": "111122223333",
    "eventCategory": "Management"
}
```

```json

{
        "eventVersion": "1.08",
        "userIdentity": {
            "type": "AssumedRole",
            "principalId": "EX_PRINCIPAL_ID",
            "arn": "arn:aws:iam::000000000000:assumed-role/role_name",
            "accountId":"123456789012",
            "accessKeyId":"AWS_ACCESS_KEY_ID_REDACTED",
            "sessionContext": {
                "sessionIssuer": {
                "type": "Role",
                "principalId": "EX_PRINCIPAL_ID",
                "arn": "arn:aws:iam::000000000000:role/Admin",
                "accountId":"123456789012",
                "userName": "SAMPLE_NAME"
                },
                "webIdFederationData": {},
                "attributes": {
                    "creationDate": "2022-10-11T17:25:41Z",
                    "mfaAuthenticated": "false"
                }
            }
        },
        "eventTime": "2022-10-11T17:30:18Z",
        "eventSource": "networkflowmonitor.amazonaws.com",
        "eventName": "ListMonitors",
        "awsRegion": "us-east-2",
        "sourceIPAddress": "192.0.2.0",
        "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
        "requestParameters": null,
        "responseElements": null,
        "requestID": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
        "eventID": "a1b2c3d4-5678-90ab-cdef-EXAMPLEbbbbb",
        "readOnly": true,
        "eventType": "AwsApiCall",
        "managementEvent": true,
        "recipientAccountId": "111122223333",
        "eventCategory": "Management"
    }
```

## Network Flow Monitor data plane events in CloudTrail

CloudTrail can capture API activities related to the CloudWatch-NetworkFlowMonitor data plane operations.

[Data events](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md#logging-data-events), also known as data plane operations, give you insight into the resource operations performed
on or within a resource. Data events are often high-volume activities.

To enable logging of Network Flow Monitor data events in CloudTrail files, you'll need to enable
the logging of data plane API activity in CloudTrail. See [Logging data events for trails](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md) for more information.

Data plane events can be filtered by resource type. Because there are additional costs for using
data events in CloudTrail, filtering by resource allows you to have more control over what you log and pay for.

Using the information that CloudTrail collects, you can identify a specific request to the CloudWatch-NetworkFlowMonitor
data plane APIs, the IP address of the requester, the requester's identity, and the date and time of the
request. Logging the data plane APIs using CloudTrail can help you with operational and
risk auditing, governance, and compliance of your AWS account.

The following are data plane APIs in Network Flow Monitor.

- [GetQueryResultsMonitorTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-getqueryresultsmonitortopcontributors.md)

- [GetQueryResultsMonitorsTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-getqueryresultsmonitorstopcontributors.md)

- [GetQueryResultsWorkloadInsightsTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-getqueryresultsworkloadinsightstopcontributors.md)

- [GetQueryResultsWorkloadInsightsTopContributorsData](../../../../reference/networkflowmonitor/2-0/apireference/api-getqueryresultsworkloadinsightstopcontributorsdata.md)

- [GetQueryStatusMonitorTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-getquerystatusmonitortopcontributors.md)

- [GetQueryStatusMonitorsTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-getquerystatusmonitorstopcontributors.md)

- [GetQueryStatusWorkloadInsightsTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-getquerystatusworkloadinsightstopcontributors.md)

- [GetQueryStatusWorkloadInsightsTopContributorsData](../../../../reference/networkflowmonitor/2-0/apireference/api-getquerystatusworkloadinsightstopcontributorsdata.md)

- [StartQueryMonitorTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-startquerymonitortopcontributors.md)

- [StartQueryMonitorsTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-startquerymonitorstopcontributors.md)

- [StartQueryWorkloadInsightsTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-startqueryworkloadinsightstopcontributors.md)

- [StartQueryWorkloadInsightsTopContributorsData](../../../../reference/networkflowmonitor/2-0/apireference/api-startqueryworkloadinsightstopcontributorsdata.md)

- [StopQueryMonitorTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-stopquerymonitortopcontributors.md)

- [StopQueryMonitorsTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-stopquerymonitorstopcontributors.md)

- [StopQueryWorkloadInsightsTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-stopqueryworkloadinsightstopcontributors.md)

- [StopQueryWorkloadInsightsTopContributorsData](../../../../reference/networkflowmonitor/2-0/apireference/api-stopqueryworkloadinsightstopcontributorsdata.md)

The following example shows a CloudTrail log entry that demonstrates the [GetQueryResultsMonitorsTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-getqueryresultsmonitorstopcontributors.md) action.

```json

{
  "eventVersion": "1.09",
  "userIdentity": {
    "type": "AssumedRole",
    "principalId": "EX_PRINCIPAL_ID",
    "arn": "arn:aws:iam::000000000000:assumed-role/role_name",
    "accountId": "123456789012",
    "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
    "sessionContext": {
      "sessionIssuer": {
        "type": "Role",
        "principalId": "EX_PRINCIPAL_ID",
        "arn": "arn:aws:iam::000000000000:role/Admin",
        "accountId": "123456789012",
         "userName": "SAMPLE_NAME"
      },
      "attributes": {
      "creationDate": "2024-11-03T15:43:27Z",
      "mfaAuthenticated": "false"
      }
    }
},
  "eventTime": "2024-11-15T14:08:04Z",
  "eventSource": "networkflowmonitor.amazonaws.com",
  "eventName": "GetQueryResultsMonitorTopContributors",
  "awsRegion": "us-east-1",
  "sourceIPAddress": "192.0.2.0",
  "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
  "errorCode": "AccessDenied",
  "requestParameters": {
    "QueryId": "a1b2c3d4-5678-90ab-cdef-EXAMPLEQuery,
    "MaxResults": "20",
    "MonitorName": "TestMonitor"
  },
  "responseElements": null,
  "requestID": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
  "eventID": "a1b2c3d4-5678-90ab-cdef-EXAMPLEbbbbb",
  "readOnly": true,
  "resources": [
    {
      "accountId": "123456789012",
      "type": "AWS::NetworkFlowMonitor::Monitor",
      "ARN": "arn:aws:networkflowmonitor:us-east-1:123456789012:monitor/TestMonitor"
    }
  ],
  "eventType": "AwsApiCall",
  "managementEvent": false,
  "recipientAccountId": "000000000000",
  "eventCategory": "Data"
}
```

The following example shows a CloudTrail log entry that demonstrates the [GetQueryResultsWorkloadInsightsTopContributors](../../../../reference/networkflowmonitor/2-0/apireference/api-getqueryresultsworkloadinsightstopcontributors.md) action.

```json

{
  "eventVersion": "1.09",
  "userIdentity": {
    "type": "AssumedRole",
    "principalId": "EX_PRINCIPAL_ID",
    "arn": "arn:aws:iam::000000000000:assumed-role/role_name",
    "accountId": "123456789012",
    "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
    "sessionContext": {
      "sessionIssuer": {
        "type": "Role",
        "principalId": "EX_PRINCIPAL_ID",
        "arn": "arn:aws:iam::000000000000:role/Admin",
        "accountId": "123456789012",
         "userName": "SAMPLE_NAME"
      },
      "attributes": {
      "creationDate": "2024-11-03T15:43:27Z",
      "mfaAuthenticated": "false"
      }
    }
},
  "eventTime": "2024-11-15T14:08:04Z",
  "eventSource": "networkflowmonitor.amazonaws.com",
  "eventName": "GetQueryResultsWorkloadInsightsTopContributorsData",
  "awsRegion": "us-east-1",
  "sourceIPAddress": "192.0.2.0",
  "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
  "errorCode": "AccessDenied",
  "requestParameters": {
    "QueryId": "a1b2c3d4-5678-90ab-cdef-EXAMPLEQuery",
    "ScopeId": "a1b2c3d4-5678-90ab-cdef-EXAMPLEScope"
  },
  "responseElements": null,
  "requestID": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
  "eventID": "a1b2c3d4-5678-90ab-cdef-EXAMPLEbbbbb",
  "readOnly": true,
  "resources": [
    {
      "accountId": "496383180932",
      "type": "AWS::NetworkFlowMonitor::Scope",
      "ARN": "arn:aws:networkflowmonitor:us-east-1:123456789012:scope/a1b2c3d4-5678-90ab-cdef-EXAMPLEScope"
    }
  ],
  "eventType": "AwsApiCall",
  "managementEvent": false,
  "recipientAccountId": "000000000000",
  "eventCategory": "Data"
}
```

## Internet Monitor in CloudTrail

Internet Monitor supports logging the following actions as events in CloudTrail log files.

- [CreateMonitor](../../../internet-monitor/latest/api/api-createmonitor.md)

- [DeleteMonitor](../../../internet-monitor/latest/api/api-deletemonitor.md)

- [GetHealthEvent](../../../internet-monitor/latest/api/api-gethealthevent.md)

- [GetMonitor](../../../internet-monitor/latest/api/api-getmonitor.md)

- [GetQueryResults](../../../internet-monitor/latest/api/api-getqueryresults.md)

- [GetQueryStatus](../../../internet-monitor/latest/api/api-getquerystatus.md)

- [ListHealthEvents](../../../internet-monitor/latest/api/api-listhealthevents.md)

- [ListInternetEvents](../../../internet-monitor/latest/api/api-listinternetevents.md)

- [ListMonitors](../../../internet-monitor/latest/api/api-listmonitors.md)

- [ListTagsForResource](../../../internet-monitor/latest/api/api-listtagsforresource.md)

- [StartQuery](../../../internet-monitor/latest/api/api-startquery.md)

- [StopQuery](../../../internet-monitor/latest/api/api-stopquery.md)

- [TagResource](../../../internet-monitor/latest/api/api-tagresource.md)

- [UntagResource](../../../internet-monitor/latest/api/api-untagresource.md)

- [UpdateMonitor](../../../internet-monitor/latest/api/api-updatemonitor.md)

### Example: Internet Monitor log file entries

The following example shows a CloudTrail Internet Monitor log entry that demonstrates the
`ListMonitors` action.

```json

{
        "eventVersion": "1.08",
        "userIdentity": {
            "type": "AssumedRole",
            "principalId": "EX_PRINCIPAL_ID",
            "arn": "arn:aws:iam::000000000000:assumed-role/role_name",
            "accountId":"123456789012",
            "accessKeyId":"AWS_ACCESS_KEY_ID_REDACTED",
            "sessionContext": {
                "sessionIssuer": {
                "type": "Role",
                "principalId": "EX_PRINCIPAL_ID",
                "arn": "arn:aws:iam::000000000000:role/Admin",
                "accountId":"123456789012",
                "userName": "SAMPLE_NAME"
                },
                "webIdFederationData": {},
                "attributes": {
                    "creationDate": "2022-10-11T17:25:41Z",
                    "mfaAuthenticated": "false"
                }
            }
        },
        "eventTime": "2022-10-11T17:30:18Z",
        "eventSource": "internetmonitor.amazonaws.com",
        "eventName": "ListMonitors",
        "awsRegion": "us-east-2",
        "sourceIPAddress": "192.0.2.0",
        "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
        "requestParameters": null,
        "responseElements": null,
        "requestID": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
        "eventID": "a1b2c3d4-5678-90ab-cdef-EXAMPLEbbbbb",
        "readOnly": true,
        "eventType": "AwsApiCall",
        "managementEvent": true,
        "recipientAccountId": "111122223333",
        "eventCategory": "Management"
    }
```

The following example shows a CloudTrail Internet Monitor log entry that demonstrates the
`CreateMonitor` action.

```nohighlight

{
        "eventVersion": "1.08",
        "userIdentity": {
            "type": "AssumedRole",
            "principalId": "EX_PRINCIPAL_ID",
            "arn": "arn:aws:iam::000000000000:assumed-role/role_name",
            "accountId":"123456789012",
            "accessKeyId":"AWS_ACCESS_KEY_ID_REDACTED",
            "sessionContext": {
                "sessionIssuer": {
                "type": "Role",
                "principalId": "EX_PRINCIPAL_ID",
                "arn": "arn:aws:iam::000000000000:role/Admin",
                "accountId":"123456789012",
                "userName": "SAMPLE_NAME"
                },
                "webIdFederationData": {},
                "attributes": {
                    "creationDate": "2022-10-11T17:25:41Z",
                    "mfaAuthenticated": "false"
                }
            }
        },
        "eventTime": "2022-10-11T17:30:08Z",
        "eventSource": "internetmonitor.amazonaws.com",
        "eventName": "CreateMonitor",
        "awsRegion": "us-east-2",
        "sourceIPAddress": "192.0.2.0",
        "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
        "requestParameters": {
            "MonitorName": "TestMonitor",
            "Resources": ["arn:aws:ec2:us-east-2:444455556666:vpc/vpc-febc0b95"],
            "ClientToken": "a1b2c3d4-5678-90ab-cdef-EXAMPLE33333"
        },
        "responseElements": {
            "Arn": "arn:aws:internetmonitor:us-east-2:444455556666:monitor/ct-onboarding-test",
            "Status": "PENDING"
        },
        "requestID": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
        "eventID": "a1b2c3d4-5678-90ab-cdef-EXAMPLEbbbbb",
        "readOnly": false,
        "eventType": "AwsApiCall",
        "managementEvent": true,
        "recipientAccountId": "111122223333",
        "eventCategory": "Management"
    }
```

## CloudWatch Synthetics information in CloudTrail

CloudWatch Synthetics supports logging the following actions as events in CloudTrail log files:

- [AssociateResource](../../../../reference/amazonsynthetics/latest/apireference/api-associateresource.md)

- [CreateCanary](../../../../reference/amazonsynthetics/latest/apireference/api-createcanary.md)

- [CreateGroup](../../../../reference/amazonsynthetics/latest/apireference/api-creategroup.md)

- [DeleteCanary](../../../../reference/amazonsynthetics/latest/apireference/api-deletecanary.md)

- [DeleteGroup](../../../../reference/amazonsynthetics/latest/apireference/api-deletegroup.md)

- [DescribeCanaries](../../../../reference/amazonsynthetics/latest/apireference/api-describecanaries.md)

- [DescribeCanariesLastRun](../../../../reference/amazonsynthetics/latest/apireference/api-describecanarieslastrun.md)

- [DescribeRuntimeVersions](../../../../reference/amazonsynthetics/latest/apireference/api-describeruntimeversions.md)

- [DisassociateResource](../../../../reference/amazonsynthetics/latest/apireference/api-disassociateresource.md)

- [GetCanary](../../../../reference/amazonsynthetics/latest/apireference/api-getcanary.md)

- [GetCanaryRuns](../../../../reference/amazonsynthetics/latest/apireference/api-getcanaryruns.md)

- [GetGroup](../../../../reference/amazonsynthetics/latest/apireference/api-getgroup.md)

- [ListAssociatedGroups](../../../../reference/amazonsynthetics/latest/apireference/api-listassociatedgroups.md)

- [ListGroupResources](../../../../reference/amazonsynthetics/latest/apireference/api-listgroupresources.md)

- [ListGroups](../../../../reference/amazonsynthetics/latest/apireference/api-listgroups.md)

- [ListTagsForResource](../../../../reference/amazonsynthetics/latest/apireference/api-listtagsforresource.md)

- [StartCanary](../../../../reference/amazonsynthetics/latest/apireference/api-startcanary.md)

- [StartCanaryDryRun](../../../../reference/amazonsynthetics/latest/apireference/api-startcanarydryrun.md)

- [StopCanary](../../../../reference/amazonsynthetics/latest/apireference/api-stopcanary.md)

- [TagResource](../../../../reference/amazonsynthetics/latest/apireference/api-tagresource.md)

- [UntagResource](../../../../reference/amazonsynthetics/latest/apireference/api-untagresource.md)

- [UpdateCanary](../../../../reference/amazonsynthetics/latest/apireference/api-updatecanary.md)

### Example: CloudWatch Synthetics log file entries

The following example shows a CloudTrail Synthetics log entry that demonstrates the
`DescribeCanaries` action.

```json

{
    "eventVersion": "1.05",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EX_PRINCIPAL_ID",
        "arn": "arn:aws:iam::123456789012:assumed-role/role_name",
        "accountId":"123456789012",
        "accessKeyId":"AWS_ACCESS_KEY_ID_REDACTED",
        "sessionContext": {
            "sessionIssuer": {
                "type": "Role",
                "principalId": "EX_PRINCIPAL_ID",
                "arn": "arn:aws:iam::111222333444:role/Administrator",
                "accountId":"123456789012",
                "userName": "SAMPLE_NAME"
            },
            "webIdFederationData": {},
            "attributes": {
                "mfaAuthenticated": "false",
                "creationDate": "2020-04-08T21:43:24Z"
            }
        }
    },
    "eventTime": "2020-04-08T23:06:47Z",
    "eventSource": "synthetics.amazonaws.com",
    "eventName": "DescribeCanaries",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "127.0.0.1",
    "userAgent": "aws-internal/3 aws-sdk-java/1.11.590 Linux/4.9.184-0.1.ac.235.83.329.metal1.x86_64 OpenJDK_64-Bit_Server_VM/25.212-b03 java/1.8.0_212 vendor/Oracle_Corporation",
    "requestParameters": null,
    "responseElements": null,
    "requestID": "201ed5f3-15db-4f87-94a4-123456789",
    "eventID": "73ddbd81-3dd0-4ada-b246-123456789",
    "readOnly": true,
    "eventType": "AwsApiCall",
    "recipientAccountId": "111122223333"
}
```

The following example shows a CloudTrail Synthetics log entry that demonstrates the
`UpdateCanary` action.

```nohighlight

{
    "eventVersion": "1.05",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EX_PRINCIPAL_ID",
        "arn": "arn:aws:iam::123456789012:assumed-role/role_name",
        "accountId":"123456789012",
        "accessKeyId":"AWS_ACCESS_KEY_ID_REDACTED",
        "sessionContext": {
            "sessionIssuer": {
               "type": "Role",
                "principalId": "EX_PRINCIPAL_ID",
                "arn": "arn:aws:iam::111222333444:role/Administrator",
       "accountId":"123456789012",
                "userName": "SAMPLE_NAME"
            },
            "webIdFederationData": {},
            "attributes": {
                "mfaAuthenticated": "false",
                "creationDate": "2020-04-08T21:43:24Z"
            }
        }
    },
    "eventTime": "2020-04-08T23:06:47Z",
    "eventSource": "synthetics.amazonaws.com",
    "eventName": "UpdateCanary",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "192.0.2.0",
    "userAgent": "aws-internal/3 aws-sdk-java/1.11.590 Linux/4.9.184-0.1.ac.235.83.329.metal1.x86_64 OpenJDK_64-Bit_Server_VM/25.212-b03 java/1.8.0_212 vendor/Oracle_Corporation",
    "requestParameters": {
        "Schedule": {
            "Expression": "rate(1 minute)"
        },
        "name": "sample_canary_name",
        "Code": {
            "Handler": "myOwnScript.handler",
            "ZipFile": "SAMPLE_ZIP_FILE"
        }
    },
    "responseElements": null,
    "requestID": "fe4759b0-0849-4e0e-be71-1234567890",
    "eventID": "9dc60c83-c3c8-4fa5-bd02-1234567890",
    "readOnly": false,
    "eventType": "AwsApiCall",
    "recipientAccountId": "111122223333"
}
```

The following example shows a CloudTrail Synthetics log entry that demonstrates the
`GetCanaryRuns` action.

```nohighlight

{
    "eventVersion": "1.05",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EX_PRINCIPAL_ID",
        "arn": "arn:aws:iam::123456789012:assumed-role/role_name",
        "accountId":"123456789012",
        "accessKeyId":"AWS_ACCESS_KEY_ID_REDACTED",
        "sessionContext": {
            "sessionIssuer": {
                "type": "Role",
                "principalId": "EX_PRINCIPAL_ID",
                "arn": "arn:aws:iam::111222333444:role/Administrator",
       "accountId":"123456789012",
                "userName": "SAMPLE_NAME"
            },
            "webIdFederationData": {},
            "attributes": {
                "mfaAuthenticated": "false",
                "creationDate": "2020-04-08T21:43:24Z"
            }
        }
    },
    "eventTime": "2020-04-08T23:06:30Z",
    "eventSource": "synthetics.amazonaws.com",
    "eventName": "GetCanaryRuns",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "127.0.0.1",
    "userAgent": "aws-internal/3 aws-sdk-java/1.11.590 Linux/4.9.184-0.1.ac.235.83.329.metal1.x86_64 OpenJDK_64-Bit_Server_VM/25.212-b03 java/1.8.0_212 vendor/Oracle_Corporation",
    "requestParameters": {
        "Filter": "TIME_RANGE",
        "name": "sample_canary_name",
        "FilterValues": [
            "2020-04-08T23:00:00.000Z",
            "2020-04-08T23:10:00.000Z"
        ]
    },
    "responseElements": null,
    "requestID": "2f56318c-cfbd-4b60-9d93-1234567890",
    "eventID": "52723fd9-4a54-478c-ac55-1234567890",
    "readOnly": true,
    "eventType": "AwsApiCall",
    "recipientAccountId": "111122223333"
}
```

## CloudWatch RUM information in CloudTrail

CloudWatch RUM supports logging the following actions as events in CloudTrail log files:

- [BatchCreateRumMetricDefinitions](../../../../reference/cloudwatchrum/latest/apireference/api-batchcreaterummetricdefinitions.md)

- [BatchDeleteRumMetricDefinitions](../../../../reference/cloudwatchrum/latest/apireference/api-batchdeleterummetricdefinitions.md)

- [BatchGetRumMetricDefinitions](../../../../reference/cloudwatchrum/latest/apireference/api-batchgetrummetricdefinitions.md)

- [CreateAppMonitor](../../../../reference/cloudwatchrum/latest/apireference/api-createappmonitor.md)

- [DeleteAppMonitor](../../../../reference/cloudwatchrum/latest/apireference/api-deleteappmonitor.md)

- [DeleteResourcePolicy](../../../../reference/cloudwatchrum/latest/apireference/api-deleteresourcepolicy.md)

- [DeleteRumMetricsDestination](../../../../reference/cloudwatchrum/latest/apireference/api-deleterummetricsdestination.md)

- [GetAppMonitor](../../../../reference/cloudwatchrum/latest/apireference/api-getappmonitor.md)

- [GetAppMonitorData](../../../../reference/cloudwatchrum/latest/apireference/api-getappmonitordata.md)

- [GetResourcePolicy](../../../../reference/cloudwatchrum/latest/apireference/api-getresourcepolicy.md)

- [ListAppMonitors](../../../../reference/cloudwatchrum/latest/apireference/api-listappmonitors.md)

- [ListRumMetricsDestinations](../../../../reference/cloudwatchrum/latest/apireference/api-listrummetricsdestinations.md)

- [ListTagsForResource](../../../../reference/cloudwatchrum/latest/apireference/api-listtagsforresource.md)

- [PutResourcePolicy](../../../../reference/cloudwatchrum/latest/apireference/api-putresourcepolicy.md)

- [PutRumMetricsDestination](../../../../reference/cloudwatchrum/latest/apireference/api-putrummetricsdestination.md)

- [TagResource](../../../../reference/cloudwatchrum/latest/apireference/api-tagresource.md)

- [UntagResource](../../../../reference/cloudwatchrum/latest/apireference/api-untagresource.md)

- [UpdateAppMonitor](../../../../reference/cloudwatchrum/latest/apireference/api-updateappmonitor.md)

- [UpdateRumMetricDefinition](../../../../reference/cloudwatchrum/latest/apireference/api-updaterummetricdefinition.md)

### Example: CloudWatch RUM log file entries

This section contains example CloudTrail entries for some CloudWatch RUM APIs.

The following example shows a CloudTrail log entry that demonstrates the [CreateAppMonitor](../../../../reference/cloudwatchrum/latest/apireference/api-createappmonitor.md) action.

```json

{
    "eventVersion": "1.09",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EXAMPLE_PRINCIPAL_ID",
        "arn": "arn:aws:sts::777777777777:assumed-role/EXAMPLE",
        "accountId": "777777777777",
        "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
        "sessionContext": {
            "sessionIssuer": {
                "type": "Role",
                "principalId": "EXAMPLE_PRINCIPAL_ID",
                "arn": "arn:aws:iam::777777777777:role/EXAMPLE",
                "accountId": "777777777777",
                "userName": "USERNAME_EXAMPLE"
            },
            "attributes": {
                "creationDate": "2024-07-23T16:48:47Z",
                "mfaAuthenticated": "false"
            }
        }
    },
    "eventTime": "2024-07-23T18:02:57Z",
    "eventSource": "rum.amazonaws.com",
    "eventName": "CreateAppMonitor",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "54.240.198.39",
    "userAgent": "aws-internal/3 aws-sdk-java/1.12.641 Linux/5.10.219-186.866.amzn2int.x86_64 OpenJDK_64-Bit_Server_VM/25.402-b08 java/1.8.0_402 vendor/Oracle_Corporation cfg/retry-mode/standard",
    "requestParameters": {
        "CustomEvents": {
            "Status": "ENABLED"
        },
        "CwLogEnabled": true,
        "Domain": "*.github.io",
        "AppMonitorConfiguration": {
            "SessionSampleRate": 1,
            "IncludedPages": [],
            "ExcludedPages": [],
            "Telemetries": [
                "performance",
                "errors",
                "http"
            ],
            "EnableXRay": false,
            "AllowCookies": true,
            "IdentityPoolId": "us-east-1:c81b9a1c-a5c9-4de5-8585-eb8df04e66f0"
        },
        "Tags": {
            "TestAppMonitor": ""
        },
        "Name": "TestAppMonitor"
    },
    "responseElements": {
        "Id": "65a8cc63-4ae8-4f2c-b5fc-4a54ef43af51"
    },
    "requestID": "cf7c30ad-25d3-4274-bab1-39c95a558007",
    "eventID": "2d43cc69-7f89-4f1a-95ae-0fc7e9b9fb3b",
    "readOnly": false,
    "eventType": "AwsApiCall",
    "managementEvent": true,
    "recipientAccountId": "777777777777",
    "eventCategory": "Management"
}
```

The following example shows a CloudTrail log entry that demonstrates the [PutRumMetricsDestination](../../../../reference/cloudwatchrum/latest/apireference/api-putrummetricsdestination.md) action.

```json

{
    "eventVersion": "1.09",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EXAMPLE_PRINCIPAL_ID",
        "arn": "arn:aws:sts::777777777777:assumed-role/EXAMPLE",
        "accountId": "777777777777",
        "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
        "sessionContext": {
            "sessionIssuer": {
                "type": "Role",
                "principalId": "EXAMPLE_PRINCIPAL_ID",
                "arn": "arn:aws:iam::777777777777:role/EXAMPLE",
                "accountId": "777777777777",
                "userName": "USERNAME_EXAMPLE"
            },
            "attributes": {
                "creationDate": "2024-07-23T16:48:47Z",
                "mfaAuthenticated": "false"
            }
        }
    },
    "eventTime": "2024-07-23T18:22:22Z",
    "eventSource": "rum.amazonaws.com",
    "eventName": "PutRumMetricsDestination",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "52.94.133.142",
    "userAgent": "aws-cli/2.13.25 Python/3.11.5 Linux/5.10.219-186.866.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/rum.put-rum-metrics-destination",
    "requestParameters": {
        "Destination": "CloudWatch",
        "AppMonitorName": "TestAppMonitor"
    },
    "responseElements": null,
    "requestID": "9b03fcce-b3a2-44fc-b771-900e1702998a",
    "eventID": "6250f9b7-0505-4f96-9668-feb64f82de5b",
    "readOnly": false,
    "eventType": "AwsApiCall",
    "managementEvent": true,
    "recipientAccountId": "777777777777",
    "eventCategory": "Management"
}
```

The following example shows a CloudTrail log entry that demonstrates the [BatchCreateRumMetricsDefinitions](../../../../reference/cloudwatchrum/latest/apireference/api-batchcreaterummetricsdefinitions.md) action.

```json

{
    "eventVersion": "1.09",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EXAMPLE_PRINCIPAL_ID",
        "arn": "arn:aws:sts::777777777777:assumed-role/EXAMPLE",
        "accountId": "777777777777",
        "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
        "sessionContext": {
            "sessionIssuer": {
                "type": "Role",
                "principalId": "EXAMPLE_PRINCIPAL_ID",
                "arn": "arn:aws:iam::777777777777:role/EXAMPLE",
                "accountId": "777777777777",
                "userName": "USERNAME_EXAMPLE"
            },
            "attributes": {
                "creationDate": "2024-07-23T16:48:47Z",
                "mfaAuthenticated": "false"
            }
        }
    },
    "eventTime": "2024-07-23T18:23:11Z",
    "eventSource": "rum.amazonaws.com",
    "eventName": "BatchCreateRumMetricDefinitions",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "52.94.133.142",
    "userAgent": "aws-cli/2.13.25 Python/3.11.5 Linux/5.10.219-186.866.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/rum.batch-create-rum-metric-definitions",
    "requestParameters": {
        "Destination": "CloudWatch",
        "MetricDefinitions": [
            {
                "Name": "NavigationToleratedTransaction",
                "Namespace": "AWS/RUM",
                "DimensionKeys": {
                    "metadata.browserName": "BrowserName"
                },
                "EventPattern": "{\"metadata\":{\"browserName\":[\"Chrome\"]},\"event_type\":[\"com.amazon.rum.performance_navigation_event\"],\"event_details\": {\"duration\": [{\"numeric\": [\"<=\",2000,\"<\",8000]}]}}"
            },
            {
                "Name": "HttpErrorCount",
                "DimensionKeys": {
                    "metadata.browserName": "BrowserName",
                    "metadata.countryCode": "CountryCode"
                },
                "EventPattern": "{\"metadata\":{\"browserName\":[\"Chrome\"], \"countryCode\":[\"US\"]},\"event_type\":[\"com.amazon.rum.http_event\"]}"
            }
        ],
        "AppMonitorName": "TestAppMonitor"
    },
    "responseElements": {
        "Errors": [],
        "MetricDefinitions": []
    },
    "requestID": "b14c5eda-f107-48e5-afae-1ac20d0962a8",
    "eventID": "001b55c6-1de1-48c0-a236-31096dffe249",
    "readOnly": false,
    "eventType": "AwsApiCall",
    "managementEvent": true,
    "recipientAccountId": "777777777777",
    "eventCategory": "Management"
}
```

## CloudWatch RUM data plane events in CloudTrail

CloudTrail can capture API activities related to the CloudWatch RUM data plane operation [PutRumEvents](../../../../reference/cloudwatchrum/latest/apireference/api-putrumevents.md).

[Data events](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md#logging-data-events), also known as data plane operations, give you insight into the resource operations performed
on or within a resource. Data events are often high-volume activities.

To enable logging of the **PutRumEvents** data events in CloudTrail files, you'll need to enable
the logging of data plane API activity in CloudTrail. See [Logging data events for trails](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md) for more information.

Data plane events can be filtered by resource type. Because there are additional costs for using
data events in CloudTrail, filtering by resource allows you to have more control over what you log and pay for.

Using the information that CloudTrail collects, you can identify a specific request to the CloudWatch RUM **PutRumEvents** API, the IP address of the requester, the requester's identity, and the date and time of the
request. Logging the **PutRumEvents** API using CloudTrail helps you enable operational and
risk auditing, governance, and compliance of your AWS account.

The following example shows a CloudTrail log entry that demonstrates the [PutRumEvents](../../../../reference/cloudwatchrum/latest/apireference/api-putrumevents.md) action.

```json

{
 "Records": [
     {
         "eventVersion": "1.09",
         "userIdentity": {
             "type": "AssumedRole",
             "principalId": "EXAMPLE_PRINCIPAL_ID",
             "arn": "arn:aws:sts::777777777777:assumed-role/EXAMPLE",
             "accountId": "777777777777",
             "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
             "sessionContext": {
                 "sessionIssuer": {
                     "type": "Role",
                     "principalId": "EXAMPLE_PRINCIPAL_ID",
                     "arn": "arn:aws:iam::777777777777:role/EXAMPLE",
                     "accountId": "777777777777",
                     "userName": "USERNAME_EXAMPLE"
                 },
                 "attributes": {
                     "creationDate": "2024-05-16T20:32:39Z",
                     "mfaAuthenticated": "false"
                 }
             },
             "invokedBy": "AWS Internal"
         },
         "eventTime": "2024-05-16T20:32:42Z",
         "eventSource": "rum.amazonaws.com",
         "eventName": "PutRumEvents",
         "awsRegion": "us-east-1",
         "sourceIPAddress": "AWS Internal",
         "userAgent": "AWS Internal",
         "requestParameters": {
             "id": "73ddbd81-1234-5678-b246-123456789",
             "batchId": "123456-3dd0-4ada-b246-123456789",
             "appMonitorDetails": {
                 "name": "APP-MONITOR-NAME",
                 "id": "123456-3dd0-4ada-b246-123456789",
                 "version": "1.0.0"
             },
             "userDetails": {
                 "userId": "73ddbd81-1111-9999-b246-123456789",
                 "sessionId": "a1b2c3456-15db-4f87-6789-123456789"
             },
             "rumEvents": [
                 {
                     "id": "201f367a-15db-1234-94a4-123456789",
                     "timestamp": "May 16, 2024, 8:32:20 PM",
                     "type": "com.amazon.rum.dom_event",
                     "metadata": "{}",
                     "details": "{}"
                 }
             ]
         },
         "responseElements": null,
         "requestID": "201ed5f3-15db-4f87-94a4-123456789",
         "eventID": "73ddbd81-3dd0-4ada-b246-123456789",
         "readOnly": false,
         "resources": [
             {
                 "accountId": "777777777777",
                 "type": "AWS::RUM::AppMonitor",
                 "ARN": "arn:aws:rum:us-east-1:777777777777:appmonitor/APPMONITOR_NAME_EXAMPLE"
             }
         ],
         "eventType": "AwsApiCall",
         "managementEvent": false,
         "recipientAccountId": "777777777777",
         "eventCategory": "Data"
     }
 ]
}
```

## Network Synthetic Monitor information in CloudTrail

Network Synthetic Monitor supports logging the following actions as events in CloudTrail log files:

- [CreateMonitor](../../../../reference/networkmonitor/latest/apireference/api-createmonitor.md)

- [CreateProbe](../../../../reference/networkmonitor/latest/apireference/api-createprobe.md)

- [DeleteMonitor](../../../../reference/networkmonitor/latest/apireference/api-deletemonitor.md)

- [DeleteProbe](../../../../reference/networkmonitor/latest/apireference/api-deleteprobe.md)

- [GetMonitor](../../../../reference/networkmonitor/latest/apireference/api-getmonitor.md)

- [GetProbe](../../../../reference/networkmonitor/latest/apireference/api-getprobe.md)

- [ListMonitors](../../../../reference/networkmonitor/latest/apireference/api-listmonitors.md)

- [ListTagsForResource](../../../../reference/networkmonitor/latest/apireference/api-listtagsforresource.md)

- [TagResource](../../../../reference/networkmonitor/latest/apireference/api-tagresource.md)

- [UntagResource](../../../../reference/networkmonitor/latest/apireference/api-untagresource.md)

- [UpdateMonitor](../../../../reference/networkmonitor/latest/apireference/api-updatemonitor.md)

- [UpdateProbe](../../../../reference/networkmonitor/latest/apireference/api-updateprobe.md)

### Example: Network Synthetic Monitor log file entries

The following example shows a Network Synthetic Monitor CloudTrail log entry that demonstrates the
`CreateMonitor` action.

```json

{
    "eventVersion": "1.09",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EX_PRINCIPAL_ID",
        "arn": "arn:aws:iam::111122223333:assumed-role/role_name",
        "accountId": "111122223333",
        "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
        "sessionContext": {
            "sessionIssuer": {
                "type": "Role",
                "principalId": "EX_PRINCIPAL_ID",
                "arn": "arn:aws:iam::111122223333:role/Admin",
                "accountId": "111122223333",
                "userName": "SAMPLE_NAME"
            },
            "attributes": {
                "creationDate": "2024-11-03T15:43:27Z",
                "mfaAuthenticated": "false"
            }
        }
    },
    "eventTime": "2024-11-03T15:58:11Z",
    "eventSource": "networksynthetics.amazonaws.com",
    "eventName": "CreateMonitor",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "192.0.2.0",
    "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
    "requestParameters": {
        "MonitorName": "TestNetworkSyntheticMonitor",
        "ClientToken": "33551db7-1618-4aab-cdef-EXAMPLE33333"
    },
    "responseElements": {
        "MonitorArn": "arn:aws:networksynthetics:us-east-1:111122223333:monitor/TestNetworkSyntheticMonitor",
        "MonitorName": "TestNetworkSyntheticMonitor",
        "MonitorStatus": "ACTIVE"
    },
    "requestID": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
    "eventID": "a1b2c3d4-5678-90ab-cdef-EXAMPLEbbbbb",
    "readOnly": false,
    "eventType": "AwsApiCall",
    "managementEvent": true,
    "recipientAccountId": "111122223333",
    "eventCategory": "Management"
}
```

## CloudWatch Observability Access Manager information in CloudTrail

CloudWatch Observability Access Manager supports logging the following actions as events in CloudTrail log files:

- [CreateLink](../../../../reference/oam/latest/apireference/api-createlink.md)

- [CreateSink](../../../../reference/oam/latest/apireference/api-createsink.md)

- [DeleteLink](../../../../reference/oam/latest/apireference/api-deletelink.md)

- [DeleteSink](../../../../reference/oam/latest/apireference/api-deletesink.md)

- [GetLink](../../../../reference/oam/latest/apireference/api-getlink.md)

- [GetSink](../../../../reference/oam/latest/apireference/api-getsink.md)

- [GetSinkPolicy](../../../../reference/oam/latest/apireference/api-getsinkpolicy.md)

- [ListAttachedLinks](../../../../reference/oam/latest/apireference/api-listattachedlinks.md)

- [ListLinks](../../../../reference/oam/latest/apireference/api-listlinks.md)

- [ListSinks](../../../../reference/oam/latest/apireference/api-listsinks.md)

- [ListTagsForResource](../../../../reference/oam/latest/apireference/api-listtagsforresource.md)

- [PutSinkPolicy](../../../../reference/oam/latest/apireference/api-putsinkpolicy.md)

- [TagResource](../../../../reference/oam/latest/apireference/api-tagresource.md)

- [UntagResource](../../../../reference/oam/latest/apireference/api-untagresource.md)

- [UpdateLink](../../../../reference/oam/latest/apireference/api-updatelink.md)

### Example: CloudWatch Observability Access Manager log file entries

The following example shows a CloudWatch Observability Access Manager CloudTrail log entry that demonstrates the
`CreateSink` action.

```json

{
    "eventVersion": "1.09",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EX_PRINCIPAL_ID",
        "arn": "arn:aws:iam::111122223333:assumed-role/role_name",
        "accountId": "111122223333",
        "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
        "sessionContext": {
            "sessionIssuer": {
                "type": "Role",
                "principalId": "EX_PRINCIPAL_ID",
                "arn": "arn:aws:iam::111122223333:role/Admin",
                "accountId": "111122223333",
                "userName": "SAMPLE_NAME"
            },
            "attributes": {
                "creationDate": "2024-11-03T15:43:27Z",
                "mfaAuthenticated": "false"
            }
        }
    },
    "eventTime": "2024-11-03T15:58:11Z",
    "eventSource": "oam.amazonaws.com",
    "eventName": "CreateSink",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "192.0.2.0",
    "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
    "requestParameters": {
        "Name": "TestObservabilitySink"
    },
    "responseElements": {
        "Arn": "arn:aws:oam:us-east-1:111122223333:sink/TestObservabilitySink",
        "Id": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
        "Name": "TestObservabilitySink"
    },
    "requestID": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
    "eventID": "a1b2c3d4-5678-90ab-cdef-EXAMPLEbbbbb",
    "readOnly": false,
    "eventType": "AwsApiCall",
    "managementEvent": true,
    "recipientAccountId": "111122223333",
    "eventCategory": "Management"
}
```

## CloudWatch Observability Admin information in CloudTrail

CloudWatch Observability Admin supports logging the following actions as events in CloudTrail log files:

- [GetTelemetryEvaluationStatus](../../../cloudwatch/latest/observabilityadmin/api-gettelemetryevaluationstatus.md)

- [GetTelemetryEvaluationStatusForOrganization](../../../cloudwatch/latest/observabilityadmin/api-gettelemetryevaluationstatusfororganization.md)

- [ListResourceTelemetry](../../../cloudwatch/latest/observabilityadmin/api-listresourcetelemetry.md)

- [ListResourceTelemetryForOrganization](../../../cloudwatch/latest/observabilityadmin/api-listresourcetelemetryfororganization.md)

- [StartTelemetryEvaluation](../../../cloudwatch/latest/observabilityadmin/api-starttelemetryevaluation.md)

- [StartTelemetryEvaluationForOrganization](../../../cloudwatch/latest/observabilityadmin/api-starttelemetryevaluationfororganization.md)

- [StopTelemetryEvaluation](../../../cloudwatch/latest/observabilityadmin/api-stoptelemetryevaluation.md)

- [StopTelemetryEvaluationForOrganization](../../../cloudwatch/latest/observabilityadmin/api-stoptelemetryevaluationfororganization.md)

### Example: CloudWatch Observability Admin log file entries

The following example shows a CloudWatch Observability Admin CloudTrail log entry that demonstrates the
`StartTelemetryEvaluation` action.

```json

{
    "eventVersion": "1.09",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EX_PRINCIPAL_ID",
        "arn": "arn:aws:iam::111122223333:assumed-role/role_name",
        "accountId": "111122223333",
        "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
        "sessionContext": {
            "sessionIssuer": {
                "type": "Role",
                "principalId": "EX_PRINCIPAL_ID",
                "arn": "arn:aws:iam::111122223333:role/Admin",
                "accountId": "111122223333",
                "userName": "SAMPLE_NAME"
            },
            "attributes": {
                "creationDate": "2024-11-03T15:43:27Z",
                "mfaAuthenticated": "false"
            }
        }
    },
    "eventTime": "2024-11-03T15:58:11Z",
    "eventSource": "observabilityadmin.amazonaws.com",
    "eventName": "StartTelemetryEvaluation",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "192.0.2.0",
    "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
    "requestParameters": {},
    "responseElements": null,
    "requestID": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
    "eventID": "a1b2c3d4-5678-90ab-cdef-EXAMPLEbbbbb",
    "readOnly": false,
    "eventType": "AwsApiCall",
    "managementEvent": true,
    "recipientAccountId": "111122223333",
    "eventCategory": "Management"
}
```

## CloudWatch Application Signals information in CloudTrail

CloudWatch Application Signals supports logging the following actions as events in CloudTrail log files:

- [BatchGetServiceLevelObjectiveBudgetReport](../../../../reference/applicationsignals/latest/apireference/api-batchgetservicelevelobjectivebudgetreport.md)

- [BatchUpdateExclusionWindows](../../../../reference/applicationsignals/latest/apireference/api-batchupdateexclusionwindows.md)

- [CreateServiceLevelObjective](../../../../reference/applicationsignals/latest/apireference/api-createservicelevelobjective.md)

- [DeleteServiceLevelObjective](../../../../reference/applicationsignals/latest/apireference/api-deleteservicelevelobjective.md)

- [GetService](../../../../reference/applicationsignals/latest/apireference/api-getservice.md)

- [GetServiceLevelObjective](../../../../reference/applicationsignals/latest/apireference/api-getservicelevelobjective.md)

- [ListServiceDependencies](../../../../reference/applicationsignals/latest/apireference/api-listservicedependencies.md)

- [ListServiceDependents](../../../../reference/applicationsignals/latest/apireference/api-listservicedependents.md)

- [ListServiceLevelObjectives](../../../../reference/applicationsignals/latest/apireference/api-listservicelevelobjectives.md)

- [ListServiceOperations](../../../../reference/applicationsignals/latest/apireference/api-listserviceoperations.md)

- [ListServices](../../../../reference/applicationsignals/latest/apireference/api-listservices.md)

- [ListTagsForResource](../../../../reference/applicationsignals/latest/apireference/api-listtagsforresource.md)

- [StartDiscovery](../../../../reference/applicationsignals/latest/apireference/api-startdiscovery.md)

- [TagResource](../../../../reference/applicationsignals/latest/apireference/api-tagresource.md)

- [UntagResource](../../../../reference/applicationsignals/latest/apireference/api-untagresource.md)

- [UpdateServiceLevelObjective](../../../../reference/applicationsignals/latest/apireference/api-updateservicelevelobjective.md)

### Example: CloudWatch Application Signals log file entries

The following example shows a CloudWatch Application Signals CloudTrail log entry that demonstrates the
`CreateServiceLevelObjective` action.

```json

{
    "eventVersion": "1.09",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EX_PRINCIPAL_ID",
        "arn": "arn:aws:iam::111122223333:assumed-role/role_name",
        "accountId": "111122223333",
        "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
        "sessionContext": {
            "sessionIssuer": {
                "type": "Role",
                "principalId": "EX_PRINCIPAL_ID",
                "arn": "arn:aws:iam::111122223333:role/Admin",
                "accountId": "111122223333",
                "userName": "SAMPLE_NAME"
            },
            "attributes": {
                "creationDate": "2024-11-03T15:43:27Z",
                "mfaAuthenticated": "false"
            }
        }
    },
    "eventTime": "2024-11-03T15:58:11Z",
    "eventSource": "applicationsignals.amazonaws.com",
    "eventName": "CreateServiceLevelObjective",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "192.0.2.0",
    "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
    "requestParameters": {
        "Name": "TestSLO",
        "Description": "Test Service Level Objective"
    },
    "responseElements": {
        "Arn": "arn:aws:applicationsignals:us-east-1:111122223333:slo/TestSLO"
    },
    "requestID": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
    "eventID": "a1b2c3d4-5678-90ab-cdef-EXAMPLEbbbbb",
    "readOnly": false,
    "eventType": "AwsApiCall",
    "managementEvent": true,
    "recipientAccountId": "111122223333",
    "eventCategory": "Management"
}
```

## CloudWatch Application Insights information in CloudTrail

CloudWatch Application Insights supports logging the following actions as events in CloudTrail log files:

- [AddWorkload](../../../../reference/cloudwatch/latest/apireference/api-addworkload.md)

- [CreateApplication](../../../../reference/cloudwatch/latest/apireference/api-createapplication.md)

- [CreateComponent](../../../../reference/cloudwatch/latest/apireference/api-createcomponent.md)

- [CreateLogPattern](../../../../reference/cloudwatch/latest/apireference/api-createlogpattern.md)

- [DeleteApplication](../../../../reference/cloudwatch/latest/apireference/api-deleteapplication.md)

- [DeleteComponent](../../../../reference/cloudwatch/latest/apireference/api-deletecomponent.md)

- [DeleteLogPattern](../../../../reference/cloudwatch/latest/apireference/api-deletelogpattern.md)

- [DescribeApplication](../../../../reference/cloudwatch/latest/apireference/api-describeapplication.md)

- [DescribeComponent](../../../../reference/cloudwatch/latest/apireference/api-describecomponent.md)

- [DescribeComponentConfiguration](../../../../reference/cloudwatch/latest/apireference/api-describecomponentconfiguration.md)

- [DescribeComponentConfigurationRecommendation](../../../../reference/cloudwatch/latest/apireference/api-describecomponentconfigurationrecommendation.md)

- [DescribeLogPattern](../../../../reference/cloudwatch/latest/apireference/api-describelogpattern.md)

- [DescribeObservation](../../../../reference/cloudwatch/latest/apireference/api-describeobservation.md)

- [DescribeProblem](../../../../reference/cloudwatch/latest/apireference/api-describeproblem.md)

- [DescribeProblemObservations](../../../../reference/cloudwatch/latest/apireference/api-describeproblemobservations.md)

- [DescribeWorkload](../../../../reference/cloudwatch/latest/apireference/api-describeworkload.md)

- [ListApplications](../../../../reference/cloudwatch/latest/apireference/api-listapplications.md)

- [ListComponents](../../../../reference/cloudwatch/latest/apireference/api-listcomponents.md)

- [ListConfigurationHistory](../../../../reference/cloudwatch/latest/apireference/api-listconfigurationhistory.md)

- [ListLogPatterns](../../../../reference/cloudwatch/latest/apireference/api-listlogpatterns.md)

- [ListLogPatternSets](../../../../reference/cloudwatch/latest/apireference/api-listlogpatternsets.md)

- [ListProblems](../../../../reference/cloudwatch/latest/apireference/api-listproblems.md)

- [ListTagsForResource](../../../../reference/cloudwatch/latest/apireference/api-listtagsforresource.md)

- [ListWorkloads](../../../../reference/cloudwatch/latest/apireference/api-listworkloads.md)

- [RemoveWorkload](../../../../reference/cloudwatch/latest/apireference/api-removeworkload.md)

- [TagResource](../../../../reference/cloudwatch/latest/apireference/api-tagresource.md)

- [UntagResource](../../../../reference/cloudwatch/latest/apireference/api-untagresource.md)

- [UpdateApplication](../../../../reference/cloudwatch/latest/apireference/api-updateapplication.md)

- [UpdateComponent](../../../../reference/cloudwatch/latest/apireference/api-updatecomponent.md)

- [UpdateComponentConfiguration](../../../../reference/cloudwatch/latest/apireference/api-updatecomponentconfiguration.md)

- [UpdateLogPattern](../../../../reference/cloudwatch/latest/apireference/api-updatelogpattern.md)

- [UpdateProblem](../../../../reference/cloudwatch/latest/apireference/api-updateproblem.md)

- [UpdateWorkload](../../../../reference/cloudwatch/latest/apireference/api-updateworkload.md)

### Example: CloudWatch Application Insights log file entries

The following example shows a CloudWatch Application Insights CloudTrail log entry that demonstrates the
`CreateApplication` action.

```json

{
    "eventVersion": "1.09",
    "userIdentity": {
        "type": "AssumedRole",
        "principalId": "EX_PRINCIPAL_ID",
        "arn": "arn:aws:iam::111122223333:assumed-role/role_name",
        "accountId": "111122223333",
        "accessKeyId": "AWS_ACCESS_KEY_ID_REDACTED",
        "sessionContext": {
            "sessionIssuer": {
                "type": "Role",
                "principalId": "EX_PRINCIPAL_ID",
                "arn": "arn:aws:iam::111122223333:role/Admin",
                "accountId": "111122223333",
                "userName": "SAMPLE_NAME"
            },
            "attributes": {
                "creationDate": "2024-11-03T15:43:27Z",
                "mfaAuthenticated": "false"
            }
        }
    },
    "eventTime": "2024-11-03T15:58:11Z",
    "eventSource": "applicationinsights.amazonaws.com",
    "eventName": "CreateApplication",
    "awsRegion": "us-east-1",
    "sourceIPAddress": "192.0.2.0",
    "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
    "requestParameters": {
        "ResourceGroupName": "TestApplicationResourceGroup"
    },
    "responseElements": {
        "ApplicationInfo": {
            "ResourceGroupName": "TestApplicationResourceGroup",
            "LifeCycle": "ACTIVE"
        }
    },
    "requestID": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
    "eventID": "a1b2c3d4-5678-90ab-cdef-EXAMPLEbbbbb",
    "readOnly": false,
    "eventType": "AwsApiCall",
    "managementEvent": true,
    "recipientAccountId": "111122223333",
    "eventCategory": "Management"
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Security considerations for Synthetics canaries

Tagging your CloudWatch resources

All content copied from https://docs.aws.amazon.com/.
