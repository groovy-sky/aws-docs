# CloudWatch Logs quotas

You can use the table in this section to review the default service quotas, also referred to as limits, for an AWS account in Amazon CloudWatch Logs.
Most of the service quotas, but not all, are listed under the Amazon CloudWatch Logs namespace in the Service Quotas console.

###### Note

If you want to request a quota increase for any of these quotas, see [the procedure](cloudwatch-limits-cwl.md#service-quotas-manage) in this section.

NameDefaultAdjustableDescriptionActive export taskEach supported Region: 1NoThe number of active (running or pending) export tasks per accountAssociateKmsKey throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of associate-kms-key calls per second per account/per regionAssociateSourceToS3TableIntegration throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of AssociateSourceToS3TableIntegration calls per second per account/per regionAssociateSourceToS3TableIntegration throttle limit in transactions per second in a burstEach supported Region: 10NoThe maximum number of AssociateSourceToS3TableIntegration calls per second per account/per region in a burstBatch sizeEach supported Region: 1 MegabytesNoThe maximum batch size in MB of a put-log-events requestCancelExportTask throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of cancel-export-task calls per second per account/per regionCancelImportTask throttle limit in transactions per secondEach supported Region: 1NoThe maximum number of CancelImportTask calls per second per account/per regionCancelImportTask throttle limit in transactions per second in a burstEach supported Region: 1NoThe maximum number of CancelImportTask calls per second per account/per region in a burstCreateExportTask throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of create-export-task calls per second per account/per regionCreateImportTask throttle limit in transactions per secondEach supported Region: 1NoThe maximum number of CreateImportTask calls per second per account/per regionCreateImportTask throttle limit in transactions per second in a burstEach supported Region: 1NoThe maximum number of CreateImportTask calls per second per account/per region in a burstCreateLogGroup throttle limit in transactions per secondEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/logs/quotas/L-D2832119)The maximum number of create-log-group calls per second per account/per regionCreateLogStream throttle limit in transactions per secondEach supported Region: 50 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/logs/quotas/L-76507CEF)The maximum number of create-log-stream calls per second per account/per regionData archivingEach supported Region: 5 GigabytesNoThe maximum size in GB of free data archivingDeleteDataProtectionPolicy throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of delete-data-protection-policy calls per second per account/per regionDeleteDestination throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of delete-destination calls per second per account/per regionDeleteLogGroup throttle limit in transactions per secondEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/logs/quotas/L-07A912D5)The maximum number of delete-log-group calls per second per account/per regionDeleteLogStream throttle limit in transactions per secondEach supported Region: 15 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/logs/quotas/L-C029A21C)The maximum number of delete-log-stream calls per second per account/per regionDeleteMetricFilter throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of delete-metric-filter calls per second per account/per regionDeleteRetentionPolicy throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of delete-retention-policy calls per second per account/per regionDeleteSubscriptionFilter throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of delete-subscription-filter calls per second per account/per regionDeleteTransformer throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of delete-transformer calls per second per account/per regionDescribeDestinations throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of describe-destinations calls per second per account/per regionDescribeExportTasks throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of describe-export-tasks calls per second per account/per regionDescribeImportTaskBatches throttle limit in transactions per secondEach supported Region: 10NoThe maximum number of DescribeImportTaskBatches calls per second per account/per regionDescribeImportTaskBatches throttle limit in transactions per second in a burstEach supported Region: 10NoThe maximum number of DescribeImportTaskBatches calls per second per account/per region in a burstDescribeImportTasks throttle limit in transactions per secondEach supported Region: 10NoThe maximum number of DescribeImportTasks calls per second per account/per regionDescribeImportTasks throttle limit in transactions per second in a burstEach supported Region: 10NoThe maximum number of DescribeImportTasks calls per second per account/per region in a burstDescribeLogGroups throttle limit in transactions per secondEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/logs/quotas/L-4284EEDE)The maximum number of describe-log-groups calls per second per account/per regionDescribeLogStreams throttle limit in transactions per secondEach supported Region: 25 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/logs/quotas/L-3F243AD0)The maximum number of describe-log-streams calls per second per account/per regionDescribeMetricFilters throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of describe-metric-filters calls per second per account/per regionDescribeSubscriptionFilters throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of describe-subscription-filters calls per second per account/per regionDisassociateSourceFromS3TableIntegration throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of DisassociateSourceFromS3TableIntegration calls per second per account/per regionDisassociateSourceFromS3TableIntegration throttle limit in transactions per second in a burstEach supported Region: 10NoThe maximum number of DisassociateSourceFromS3TableIntegration calls per second per account/per region in a burstEvent sizeEach supported Region: 1,024 KilobytesNoThe maximum log event size in KBFilterLogEvents throttle limit in transactions per second

us-east-1: 25 per second

ap-northeast-3: 5 per second

ap-southeast-3: 5 per second

ca-west-1: 5 per second

eu-central-1: 5 per second

il-central-1: 5 per second

Each of the other supported Regions: 10 per second

NoThe maximum number of filter-log-events calls per second per account/per regionGetDataProtectionPolicy throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of get-data-protection-policy calls per second per account/per regionGetLogEvents throttle limit in transactions per second

us-west-2: 10 per second

ap-northeast-3: 10 per second

ap-southeast-3: 10 per second

ca-west-1: 10 per second

eu-central-1: 10 per second

eu-west-1: 10 per second

eu-west-3: 30 per second

il-central-1: 10 per second

Each of the other supported Regions: 25 per second

NoThe maximum number of get-log-events calls per second per account/per regionGetQueryResults throttle limit in transactions per secondEach supported Region: 10NoThe maximum number of get-query-results calls per second per account/per regionGetTransformer throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of get-transformer calls per second per account/per regionListSourcesForS3TableIntegration throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of ListSourcesForS3TableIntegration calls per second per account/per regionListSourcesForS3TableIntegration throttle limit in transactions per second in a burstEach supported Region: 10NoThe maximum number of ListSourcesForS3TableIntegration calls per second per account/per region in a burstListTagsForResource throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of list-tags-for-resource calls per second per account/per regionListTagsLogGroup throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of list-tags-log-group calls per second per account/per regionLive Tail concurrent sessions limitEach supported Region: 15[Yes](https://console.aws.amazon.com/servicequotas/home/services/logs/quotas/L-D6990B80)The maximum number of concurrent active Live Tail sessions per accountLog groupsEach supported Region: 1,000,000[Yes](https://console.aws.amazon.com/servicequotas/home/services/logs/quotas/L-C7B9AAAB)The maximum number of log groups an account can haveLog groups scanned per Live Tail sessionEach supported Region: 10NoThe maximum number of log groups that can be scanned per Live Tail sessionMetrics filters per log groupEach supported Region: 100NoThe number of metric filters per log groupPutBearerTokenAuthentication throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of put-bearer-token-authentication calls per second per account/per regionPutBearerTokenAuthentication throttle limit in transactions per second in a burstEach supported Region: 10NoThe maximum number of put-bearer-token-authentication calls per second per account/per region in a burstPutDataProtectionPolicy throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of put-data-protection-policy calls per second per account/per regionPutDestination throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of put-destination calls per second per account/per regionPutDestinationPolicy throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of put-destination-policy calls per second per account/per regionPutLogEvents throttle limit in transactions per secondEach supported Region: 5,000 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/logs/quotas/L-7E1FAE88)The maximum number of put-log-events calls per second per account/per regionPutMetricFilter throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of put-metric-filter calls per second per account/per regionPutRetentionPolicy throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of put-retention-policy calls per second per account/per regionPutSubscriptionFilter throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of put-subscription-filter calls per second per account/per regionPutTransformer throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of put-transformer calls per second per account/per regionResource policiesEach supported Region: 10NoThe maximum number of resource policies per account/per regionStartLiveTail throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of start-live-tail calls per second per account/per region. This limit applies independently to both console and APIStartQuery throttle limit in transactions per secondEach supported Region: 10NoThe maximum number of start-query calls per second per account/per regionSubscription filters per log groupEach supported Region: 2NoThe number of subscription filters per log groupTagLogGroup throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of tag-log-group calls per second per account/per regionTagResource throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of tag-resource calls per second per account/per regionTestMetricFilter throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of test-metric-filter calls per second per account/per regionTestTransformer throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of test-transformer calls per second per account/per regionUntagLogGroup throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of untag-log-group calls per second per account/per regionUntagResource throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of untag-resource calls per second per account/per regionlogs anomaly detectorsEach supported Region: 500[Yes](https://console.aws.amazon.com/servicequotas/home/services/logs/quotas/L-B53B9649)The maximum number of active log anomaly detectors

## Managing your CloudWatch Logs service quotas

CloudWatch Logs has integrated with Service Quotas, an AWS service that enables you to view and
manage your quotas from a central location. For more information, see [What Is\
Service Quotas?](../../../servicequotas/latest/userguide/intro.md) in the _Service Quotas User Guide_.

Service Quotas makes it easy to look up the value of your CloudWatch Logs service quotas.

AWS Management Console

###### To view CloudWatch Logs service quotas using the console

1. Open the Service Quotas console at [https://console.aws.amazon.com/servicequotas/](https://console.aws.amazon.com/servicequotas).

2. In the navigation pane, choose **AWS**
**services**.

3. From the **AWS services** list, search for
    and select **Amazon CloudWatch Logs**.

In the **Service quotas** list, you can see
    the service quota name, applied value (if it is available),
    AWS default quota, and whether the quota value is
    adjustable.

4. To view additional information about a service quota, such as
    the description, choose the quota name.

5. (Optional) To request a quota increase,
    select the quota that you want to increase, select
    **Request quota increase**, enter or select
    the required information, and select
    **Request**.

To work more with service quotas using the console see the [Service Quotas User Guide](../../../servicequotas/latest/userguide/intro.md). To request a quota increase, see [Requesting a quota increase](../../../servicequotas/latest/userguide/request-quota-increase.md) in the
_Service Quotas User Guide_.

AWS CLI

###### To view CloudWatch Logs service quotas using the AWS CLI

Run the following command to view the default CloudWatch Logs quotas.

```nohighlight

aws service-quotas list-aws-default-service-quotas \
    --query 'Quotas[*].{Adjustable:Adjustable,Name:QuotaName,Value:Value,Code:QuotaCode}' \
    --service-code logs \
    --output table
```

To work more with service quotas using the AWS CLI, see the [Service Quotas AWS CLI Command Reference](../../../cli/latest/reference/service-quotas/index.md#cli-aws-service-quotas). To request a quota increase,
see the [`request-service-quota-increase`](../../../cli/latest/reference/service-quotas/request-service-quota-increase.md) command in
the [AWS CLI Command Reference](../../../cli/latest/reference/service-quotas/index.md#cli-aws-service-quotas).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Monitoring usage with CloudWatch metrics

Document history

All content copied from https://docs.aws.amazon.com/.
