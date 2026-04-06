# CloudWatch service quotas

Amazon CloudWatch provides monitoring and observability for your AWS resources and applications.
To ensure optimal performance and prevent abuse, CloudWatch imposes service quotas on various
aspects of its functionality. This chapter outlines the key quotas for CloudWatch services,
including metrics, alarms, API requests, and notifications. Understanding these quotas is
crucial for effectively planning and managing your CloudWatch usage.

###### Note

For some AWS services including CloudWatch, you can use the CloudWatch usage metrics to
visualize your current service usage on CloudWatch graphs and dashboards. You can use a CloudWatch
metric math function to display the service quotas for those resources on your graphs.
You can also configure alarms that alert you when your usage approaches a service quota.
For more information, see [Visualizing your service quotas and setting alarms](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Quotas-Visualize-Alarms.html).

###### Topics

- [CloudWatch](#CloudWatch-quotas)

- [CloudWatch investigations (Amazon AI Operations)](#CloudWatch-AIO-quotas)

- [CloudWatch Application Insights](#CloudWatch-AppIns-quotas)

- [CloudWatch Application Signals](#CloudWatch-AppSig-quotas)

- [Internet Monitor](#CloudWatch-IM-quotas)

- [Network Flow Monitor](#CloudWatch-NetworkFlowMonitor-quotas)

- [Network Synthetic Monitor](#nw-monitor-quotas)

- [CloudWatch Observability Access Manager](#CloudWatch-OAM-quotas)

- [CloudWatch Observability Admin](#CloudWatch-ObsAdmin-quotas)

- [CloudWatch RUM](#CloudWatch-RUM-quotas)

- [Managing your CloudWatch service quotas](#service-quotas-manage)

## CloudWatch

This section details the service quotas specifically for core CloudWatch functionality.
These quotas cover aspects such as metrics, alarms, API requests, and dashboards.
Familiarizing yourself with these limits helps you optimize your CloudWatch configuration and
avoid potential throttling or service disruptions.

NameDefaultAdjustableDescriptionCanary limit

us-east-1: 300

ap-east-2: 200

ap-northeast-1: 300

Each of the other supported Regions: 500

[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-C1FE0F5C)The maximum number of canaries per account per region.Minimum frequencyEach supported Region: 60,000 MillisecondsNoThe minimum time, in milliseconds, between runs of the same canary.Number of Alarm Mute RulesEach supported Region: 2,000NoThe maximum number of Alarm Mute Rules that you can have in this account in the current region.Number of Contributor Insights rulesEach supported Region: 100[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-DBD11BCC)The maximum number of Contributor Insights rules you can have in this account.Number of Metrics Insights alarmsEach supported Region: 200NoThe maximum number of Metrics Insights alarms that you can have in this account in the current region.Rate of ANOMALY\_DETECTION\_BAND usage in GetMetricDataEach supported Region: 1,000NoThe maximum number of times the ANOMALY\_DETECTION\_BAND function can be used in all GetMetricData requests, per second, in this account in the current region.Rate of DB\_PERF\_INSIGHTS usage in GetMetricDataEach supported Region: 4NoThe maximum number of times the DB\_PERF\_INSIGHTS function can be used in all GetMetricData requests, per second, in this account in the current region.Rate of DeleteAlarmMuteRule requestsEach supported Region: 3 per secondNoThe maximum number of DeleteAlarmMuteRule requests that you can make, per second, in this account in the current region.Rate of DeleteAlarms requestsEach supported Region: 3 per secondNoThe maximum number of DeleteAlarms requests that you can make, per second, in this account in the current region.Rate of DeleteAnomalyDetector requestsEach supported Region: 5 per secondNoThe maximum number of DeleteAnomalyDetector requests that you can make, per second, in this account in the current region.Rate of DeleteDashboards requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-E1508405)The maximum number of DeleteDashboards requests that you can make, per second, in this account in the current region.Rate of DeleteInsightRules requestsEach supported Region: 5 per secondNoThe maximum number of DeleteInsightRules requests you can make per second in this account.Rate of DeleteMetricStream requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-0F4E28CA)The maximum number of DeleteMetricStream requests that you can make, per second, in this account in the current region.Rate of DescribeAlarmContributors requestsEach supported Region: 9 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-7D8B1BC8)The maximum number of DescribeAlarmContributors requests that you can make, per second, in this account in the current region.Rate of DescribeAlarmHistory requestsEach supported Region: 20 per secondNoThe maximum number of DescribeAlarmHistory requests that you can make, per second, in this account in the current region.Rate of DescribeAlarms requestsEach supported Region: 9 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-21CB40A4)The maximum number of DescribeAlarms requests that you can make, per second, in this account in the current region.Rate of DescribeAlarmsForMetric requestsEach supported Region: 9 per secondNoThe maximum number of DescribeAlarmsForMetric requests that you can make, per second, in this account in the current region.Rate of DescribeAnomalyDetectors requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-6300B446)The maximum number of DescribeAnomalyDetectors requests that you can make, per second, in this account in the current region.Rate of DescribeInsightRules requestsEach supported Region: 20 per secondNoThe maximum number of DescribeInsightRules requests you can make per second in this account.Rate of DisableAlarmActions requestsEach supported Region: 3 per secondNoThe maximum number of DisableAlarmActions requests that you can make, per second, in this account in the current region.Rate of DisableInsightRules requestsEach supported Region: 1 per secondNoThe maximum number of DisableInsightRules requests you can make per second in this account.Rate of EnableAlarmActions requestsEach supported Region: 3 per secondNoThe maximum number of EnableAlarmActions requests that you can make, per second, in this account in the current region.Rate of EnableInsightRules requestsEach supported Region: 1 per secondNoThe maximum number of EnableInsightRules requests you can make per second in this account.Rate of GetAlarmMuteRule requestsEach supported Region: 20 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-8445A443)The maximum number of GetAlarmMuteRule requests that you can make, per second, in this account in the current region.Rate of GetDashboard requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-E82C279D)The maximum number of GetDashboard requests that you can make, per second, in this account in the current region.Rate of GetInsightRuleReport requestsEach supported Region: 20 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-1F0C4E0C)The maximum number of GetInsightRuleReport requests you can make, per second in this account.Rate of GetMetricData datapoints for metrics older than three hoursEach supported Region: 396,000NoThe maximum number of GetMetricData datapoints that you can fetch, per second, for a request with a StartTime of more than three hours in this account in the current region.Rate of GetMetricData datapoints for the last three hours of metricsEach supported Region: 180,000NoThe maximum number of GetMetricData datapoints that you can fetch, per second, for a request with a StartTime of less than or equal to three hours in this account in the current region.Rate of GetMetricData datapoints using Metrics InsightsEach supported Region: 4,300,000NoThe maximum number of GetMetricData datapoints that you can fetch using Metrics Insights, per second, for a request with a StartTime of less than or equal to three hours in this account in the current region.Rate of GetMetricData requestsEach supported Region: 500 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-5E141212)The maximum number of GetMetricData requests that you can make, per second, in this account in the current region.Rate of GetMetricStatistics requestsEach supported Region: 400 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-EE839489)The maximum number of GetMetricStatistics requests that you can make, per second, in this account in the current region.Rate of GetMetricStream requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-59022E75)The maximum number of GetMetricStream requests that you can make, per second, in this account in the current region.Rate of GetMetricWidgetImage requestsEach supported Region: 20 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-6FCAAA2E)The maximum number of GetMetricWidgetImage requests that you can make, per second, in this account in the current region.Rate of INSIGHT\_RULE\_METRIC usage in GetMetricDataEach supported Region: 20NoThe maximum number of times the INSIGHT\_RULE\_METRIC function can be used in all GetMetricData requests, per second, in this account in the current region.Rate of LAMBDA usage in GetMetricDataEach supported Region: 5NoThe maximum number of times the LAMBDA function can be used in all GetMetricData requests, per second, in this account in the current region.Rate of ListAlarmMuteRules requestsEach supported Region: 50 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-B1A3FF67)The maximum number of ListAlarmMuteRules requests that you can make, per second, in this account in the current region.Rate of ListDashboards requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-69C44FFD)The maximum number of ListDashboards requests that you can make, per second, in this account in the current region.Rate of ListMetricStreams requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-A1710150)The maximum number of ListMetricStreams requests that you can make, per second, in this account in the current region.Rate of ListMetrics requestsEach supported Region: 25 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-05D334F0)The maximum number of ListMetrics requests that you can make, per second, in this account in the current region.Rate of ListTagsForResource requestsEach supported Region: 10 per secondNoThe maximum number of ListTagsForResource requests that you can make, per second, in this account in the current region.Rate of Metrics Insights usage in GetMetricDataEach supported Region: 10NoThe maximum number of times Metrics Insights can be used in all GetMetricData requests, per second, in this account in the current region.Rate of PutAlarmMuteRule requestsEach supported Region: 3 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-73547960)The maximum number of PutAlarmMuteRule requests that you can make, per second, in this account in the current region.Rate of PutAnomalyDetector requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-8A387C66)The maximum number of PutAnomalyDetector requests that you can make, per second, in this account in the current region.Rate of PutCompositeAlarm requestsEach supported Region: 3 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-515B0B71)The maximum number of PutCompositeAlarm requests that you can make, per second, in this account in the current region.Rate of PutDashboard requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-6753900D)The maximum number of PutDashboard requests that you can make, per second, in this account in the current region.Rate of PutInsightRule requestsEach supported Region: 5 per secondNoThe maximum number of PutInsightRule requests you can make, per second in this account.Rate of PutMetricAlarm requestsEach supported Region: 3 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-0720E68F)The maximum number of PutMetricAlarm requests that you can make, per second, in this account in the current region.Rate of PutMetricData requestsEach supported Region: 500 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-8BC498D4)The maximum number of PutMetricData requests that you can make, per second, in this account in the current region.Rate of PutMetricStream requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-A6D89949)The maximum number of PutMetricStream requests that you can make, per second, in this account in the current region.Rate of SEARCH usage in GetMetricDataEach supported Region: 50NoThe maximum number of times the SEARCH function can be used in all GetMetricData requests, per second, in this account in the current region.Rate of SERVICE\_QUOTA usage in GetMetricDataEach supported Region: 1,000NoThe maximum number of times the SERVICE\_QUOTA function can be used in all GetMetricData requests, per second, in this account in the current region.Rate of SetAlarmState requestsEach supported Region: 3 per secondNoThe maximum number of SetAlarmState requests that you can make, per second, in this account in the current region.Rate of StartMetricStreams requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-787E531D)The maximum number of StartMetricStreams requests that you can make, per second, in this account in the current region.Rate of StopMetricStreams requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/monitoring/quotas/L-A64F5500)The maximum number of StopMetricStreams requests that you can make, per second, in this account in the current region.Rate of TagResource requestsEach supported Region: 20 per secondNoThe maximum number of TagResource requests that you can make, per second, in this account in the current region.Rate of UntagResource requestsEach supported Region: 20 per secondNoThe maximum number of UntagResource requests that you can make, per second, in this account in the current region.

## CloudWatch investigations (Amazon AI Operations)

CloudWatch investigations (Amazon AI Operations) enables intelligent problem detection and root cause analysis. This
section outlines the quotas related to concurrent investigations, investigation groups,
and AI-assisted analyses.

NameDefaultAdjustableDescriptionConcurrent active investigationsEach supported Region: 2NoThe maximum number of concurrent investigations with active AI analysis per account per region.Investigation groupsEach supported Region: 1NoThe maximum number of investigation groups that can be created per account in a region.Monthly investigationsEach supported Region: 150NoThe maximum number of AI-assisted investigations per month in this account in the current region.

## CloudWatch Application Insights

CloudWatch Application Insights helps you monitor your applications and troubleshoot issues quickly. This section
presents the quotas for CloudWatch Application Insights, including limits on API requests, applications, log
streams, and metrics. Knowing these quotas allows you to effectively plan your
application monitoring setup.

Resource Default quotaAdjustable

API requests

All API actions are throttled to 5 TPS

No

Resource Group applications

100 per account

No

Account applications

1 per account

No

Log Streams

5 per resource

No

Observations per problem

20 per dashboard

40 per DescribeProblemObservations action

No

Metrics

60 per resource

No

## CloudWatch Application Signals

CloudWatch Application Signals provides deep insights into your application's performance
and dependencies. This section covers the quotas for Application Signals, including API
request limits and service-level objective (SLO) constraints. Understanding these quotas
is crucial for implementing effective application performance monitoring.

NameDefaultAdjustableDescriptionNumber of SLOs per RegionEach supported Region: 250[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-3FECAFD0)The maximum number of SLOs that you can have in this account in the current region.Number of SLOs per ServiceEach supported Region: 100[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-660777C0)The maximum number of SLOs that you can have per service in the current region.Rate of BatchGetServiceLevelObjectiveBudgetReport requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-C48E6360)The maximum number of BatchGetServiceLevelObjectiveBudgetReport requests you can make per second in this Region.Rate of BatchUpdateExclusionWindows requestsEach supported Region: 5 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-A755C830)The maximum number of BatchUpdateExclusionWindows requests you can make per second in this Region.Rate of CreateServiceLevelObjective requestsEach supported Region: 5 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-BC1D91EF)The maximum number of CreateServiceLevelObjective requests you can make per second in this Region.Rate of DeleteGroupingConfiguration requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-B7AF0347)The maximum number of DeleteGroupingConfiguration requests you can make per second in this Region.Rate of DeleteServiceLevelObjective requestsEach supported Region: 5 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-EA133927)The maximum number of DeleteServiceLevelObjective requests you can make per second in this Region.Rate of GetService requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-A7EC90F1)The maximum number of GetService requests you can make per second in this Region.Rate of GetServiceLevelObjective requestsEach supported Region: 5 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-692A17A9)The maximum number of GetServiceLevelObjective requests you can make per second in this Region.Rate of ListAuditFindings requestsEach supported Region: 1 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-D898488F)The maximum number of ListAuditFindings requests you can make per second in this Region.Rate of ListEntityEvents requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-661F9404)The maximum number of ListEntityEvents requests you can make per second in this Region.Rate of ListGroupingAttributeDefinitions requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-7A03F44C)The maximum number of ListGroupingAttributeDefinitions requests you can make per second in this Region.Rate of ListServiceDependencies requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-A3B40D72)The maximum number of ListServiceDependencies requests you can make per second in this Region.Rate of ListServiceDependents requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-7F8DB46D)The maximum number of ListServiceDependents requests you can make per second in this Region.Rate of ListServiceLevelObjectiveExclusionWindows requestsEach supported Region: 5 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-A1810D82)The maximum number of ListServiceLevelObjectiveExclusionWindows requests you can make per second in this Region.Rate of ListServiceLevelObjectives requestsEach supported Region: 5 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-B4EFDB9E)The maximum number of ListServiceLevelObjectives requests you can make per second in this Region.Rate of ListServiceOperations requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-CD5EC149)The maximum number of ListServiceOperations requests you can make per second in this Region.Rate of ListServiceStates requestsEach supported Region: 5 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-715DC88B)The maximum number of ListServiceStates requests you can make per second in this Region.Rate of ListServices requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-6BBF1A92)The maximum number of ListServices requests you can make per second in this Region.Rate of ListTagsForResource requestsEach supported Region: 5 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-E65D4614)The maximum number of ListTagsForResource requests you can make per second in this Region.Rate of PutGroupingConfiguration requestsEach supported Region: 10 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-2C8EDAF0)The maximum number of PutGroupingConfiguration requests you can make per second in this Region.Rate of StartDiscovery requestsEach supported Region: 5 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-EB837E0B)The maximum number of StartDiscovery requests you can make per second in this Region.Rate of TagResource requestsEach supported Region: 5 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-FCC0FB4B)The maximum number of TagResource requests you can make per second in this Region.Rate of UntagResource requestsEach supported Region: 5 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-7A9D456C)The maximum number of UntagResource requests you can make per second in this Region.Rate of UpdateServiceLevelObjective requestsEach supported Region: 5 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/application-signals/quotas/L-7AB759D6)The maximum number of UpdateServiceLevelObjective requests you can make per second in this Region.

## Internet Monitor

Internet Monitor, a feature of CloudWatch Network Monitoring, helps you understand how internet issues
impact your application's performance. This section provides the quotas specific to
Internet Monitor, to help you to plan an effective internet performance monitoring strategy.

NameDefaultAdjustableDescriptionDays that health events are retainedEach supported Region: 400NoThe number of days that AWS keeps information about a resolved Internet Monitor health event.Monitors per account per AWS RegionEach supported Region: 50NoThe maximum number of monitors in an account in one AWS Region.Resources per monitorEach supported Region: 50NoThe maximum number of resources that a monitor can have.

## Network Flow Monitor

Network Flow Monitor, a feature of CloudWatch Network Monitoring, provides near real-time visibility into
network performance, such as packet loss and latency, for traffic between Amazon EC2
instances, as well as traffic toward certain other AWS services. It also helps you to
determine if AWS issues are causing network performance degradation for your
application. This section provides the quotas for Network Flow Monitor, including limits on monitored
flows, data retention, and API requests. Understanding these quotas can help you to plan
an effective network flow monitoring strategy.

ResourceDefault quotaAdjustable?

Scopes per account per AWS Region

1

No

Monitors per account per AWS Region

20

Yes

Local resources per monitor

25

No

Remote resources per monitor

25

No

## Network Synthetic Monitor

Network Synthetic Monitor, a feature of CloudWatch Network Monitoring, enables you to proactively monitor your
network endpoints and API operations using configurable tests. This section details the
quotas for Network Synthetic Monitor, including limits on synthetic monitors, test frequency, and script
execution. These quotas are important for planning your network testing strategy and
maintaining reliable endpoint monitoring.

NameDefaultAdjustableDescriptionNumber of monitors per account per AWS regionEach supported Region: 100[Yes](https://console.aws.amazon.com/servicequotas/home/services/networkmonitor/quotas/L-A4298AB9)The maximum number of monitors in an account in one AWS Region.Number of probes per monitorEach supported Region: 24[Yes](https://console.aws.amazon.com/servicequotas/home/services/networkmonitor/quotas/L-F192A8D6)The maximum number of probes that a monitor can have.Number of probes per subnet for each monitorEach supported Region: 4[Yes](https://console.aws.amazon.com/servicequotas/home/services/networkmonitor/quotas/L-A8FA6DFE)The maximum number of probes that a subnet in a monitor can have.

## CloudWatch Observability Access Manager

CloudWatch Observability Access Manager allows you to securely share CloudWatch resources across
accounts. This section details the quotas for Observability Access Manager, including
limits on API requests, source account links, and sinks. Understanding these quotas is
essential for implementing cross-account observability effectively.

NameDefaultAdjustableDescriptionLinks per sinkEach supported Region: 100,000NoMaximum number of links that can be attached to a sinkNumber of linksEach supported Region: 5NoMaximum number of links in your accountNumber of sinksEach supported Region: 1NoMaximum number of sinks in your accountRate of CreateLink requestsEach supported Region: 10 per secondNoMaximum number of CreateLink requests you can make per second, in this account in the current regionRate of CreateSink requestsEach supported Region: 10 per secondNoMaximum number of CreateSink requests you can make per second, in this account in the current regionRate of DeleteLink requestsEach supported Region: 10 per secondNoMaximum number of DeleteLink requests you can make per second, in this account in the current regionRate of DeleteSink requestsEach supported Region: 10 per secondNoMaximum number of DeleteSink requests you can make per second, in this account in the current regionRate of GetLink requestsEach supported Region: 10 per secondNoMaximum number of GetLink requests you can make per second, in this account in the current regionRate of GetSink requestsEach supported Region: 10 per secondNoMaximum number of GetSink requests you can make per second, in this account in the current regionRate of GetSinkPolicy requestsEach supported Region: 10 per secondNoMaximum number of GetSinkPolicy requests you can make per second, in this account in the current regionRate of ListAttachedLinks requestsEach supported Region: 10 per secondNoMaximum number of ListAttachedLinks requests you can make per second, in this account in the current regionRate of ListLinks requestsEach supported Region: 10 per secondNoMaximum number of ListLinks requests you can make per second, in this account in the current regionRate of ListSinks requestsEach supported Region: 10 per secondNoMaximum number of ListSinks requests you can make per second, in this account in the current regionRate of ListTagsForResource requestsEach supported Region: 10 per secondNoMaximum number of ListTagsForResource requests you can make per second, in this account in the current regionRate of PutSinkPolicy requestsEach supported Region: 1 per secondNoMaximum number of PutSinkPolicy requests you can make per second, in this account in the current regionRate of TagResource requestsEach supported Region: 10 per secondNoMaximum number of TagResource requests you can make per second, in this account in the current regionRate of UntagResource requestsEach supported Region: 10 per secondNoMaximum number of UntagResource requests you can make per second, in this account in the current regionRate of UpdateLink requestsEach supported Region: 10 per secondNoMaximum number of UpdateLink requests you can make per second, in this account in the current region

## CloudWatch Observability Admin

CloudWatch Observability Admin helps you discover, audit, and manage CloudWatch telemetry
configurations across an AWS account or organization. It provides a consolidated view
of monitoring settings for resources, enabling you to ensure consistent and proper data
collection by simplifying the management of telemetry across multiple accounts. Most quotas are adjustable to
accommodate enterprise-scale deployments that need higher limits for managing
observability rules across large organizations.

NameDefaultAdjustableDescriptionCreateCentralizationRuleForOrganization throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of create-centralization-rule-for-organization calls per second per account/per regionCreateTelemetryPipeline throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of CreateTelemetryPipeline calls per second per account and per RegionCreateTelemetryRule throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of create-telemetry-rule calls per second per account/per regionCreateTelemetryRuleForOrganization throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of create-telemetry-rule-for-organization calls per second per account/per regionDeleteCentralizationRuleForOrganization throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of delete-centralization-rule-for-organization calls per second per account/per regionDeleteTelemetryPipeline throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of DeleteTelemetryPipeline calls per second per account and per RegionDeleteTelemetryRule throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of delete-telemetry-rule calls per second per account/per regionDeleteTelemetryRuleForOrganization throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of delete-telemetry-rule-for-organization calls per second per account/per regionGetCentralizationRuleForOrganization throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of get-centralization-rule-for-organization calls per second per account/per regionGetTelemetryEnrichmentStatus throttle limit in transactions per secondEach supported Region: 5 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/observabilityadmin/quotas/L-B19247CC)The maximum number of GetTelemetryEnrichmentStatus calls per second per account and per regionGetTelemetryPipeline throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of GetTelemetryPipeline calls per second per account and per RegionGetTelemetryRule throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of get-telemetry-rule calls per second per account/per regionGetTelemetryRuleForOrganization throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of get-telemetry-rule-for-organization calls per second per account/per regionListCentralizationRulesForOrganization throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of list-centralization-rules-for-organization calls per second per account/per regionListTagsForResource throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of list-tags-for-resource calls per second per account/per regionListTelemetryPipelines throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of ListTelemetryPipelines calls per second per account and per RegionListTelemetryRules throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of list-telemetry-rules calls per second per account/per regionListTelemetryRulesForOrganization throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of list-telemetry-rules-for-organization calls per second per account/per regionLogs Centralization throughput per destination regionEach supported Region: 2.5 Gigabytes per secondNoThe maximum organization-level throughput for centralization of logs per destination regionOrganization Centralization RulesEach supported Region: 50[Yes](https://console.aws.amazon.com/servicequotas/home/services/observabilityadmin/quotas/L-B8EC8109)The maximum number of organization centralization rules that you can create for an organizationOrganization Telemetry RulesEach supported Region: 50[Yes](https://console.aws.amazon.com/servicequotas/home/services/observabilityadmin/quotas/L-74ABF7EB)The maximum number of organization telemetry rules that you can create for an organizationPipelines with a CloudWatch Logs sourceEach supported Region: 300[Yes](https://console.aws.amazon.com/servicequotas/home/services/observabilityadmin/quotas/L-928E7B34)The maximum number of pipelines with a CloudWatch Logs source allowed per account per RegionPipelines with a third-party integration sourceEach supported Region: 30[Yes](https://console.aws.amazon.com/servicequotas/home/services/observabilityadmin/quotas/L-57445DF8)The maximum number of pipelines with a third-party integration source allowed per account per RegionStartTelemetryEnrichment throttle limit in transactions per secondEach supported Region: 1 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/observabilityadmin/quotas/L-E3DAB469)The maximum number of StartTelemetryEnrichment calls per second per account and per regionStopTelemetryEnrichment throttle limit in transactions per secondEach supported Region: 1 per second[Yes](https://console.aws.amazon.com/servicequotas/home/services/observabilityadmin/quotas/L-C6A78D4B)The maximum number of StopTelemetryEnrichment calls per second per account and per regionTagResource throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of tag-resource calls per second per account/per regionTelemetry RulesEach supported Region: 20[Yes](https://console.aws.amazon.com/servicequotas/home/services/observabilityadmin/quotas/L-A344069D)The maximum number of telemetry rules that you can create for an accountTestTelemetryPipeline throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of TestTelemetryPipeline calls per second per account and per RegionUntagResource throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of untag-resource calls per second per account/per regionUpdateCentralizationRuleForOrganization throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of update-centralization-rule-for-organization calls per second per account/per regionUpdateTelemetryPipeline throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of UpdateTelemetryPipeline calls per second per account and per RegionUpdateTelemetryRule throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of update-telemetry-rule calls per second per account/per regionUpdateTelemetryRuleForOrganization throttle limit in transactions per secondEach supported Region: 5NoThe maximum number of update-telemetry-rule-for-organization calls per second per account/per regionValidateTelemetryPipelineConfiguration throttle limit in transactions per secondEach supported Region: 5 per secondNoThe maximum number of ValidateTelemetryPipelineConfiguration calls per second per account and per Region

## CloudWatch RUM

CloudWatch RUM (Real-User Monitoring) helps you collect and analyze performance data from
real user interactions with your web applications. This section presents the quotas for
CloudWatch RUM, allowing you to plan your real-user monitoring strategy within the service
limits.

NameDefaultAdjustableDescriptionRUM AppMonitorsEach supported Region: 20[Yes](https://console.aws.amazon.com/servicequotas/home/services/rum/quotas/L-3FB7EA17)The maximum number of RUM AppMonitors that you can have in this AWS account.RUM Events per second per AWS AccountEach supported Region: 50[Yes](https://console.aws.amazon.com/servicequotas/home/services/rum/quotas/L-35851224)The maximum number of RUM Events per second that RUM will ingest for this AWS account.

## Managing your CloudWatch service quotas

Effective management of CloudWatch service quotas is crucial for maintaining optimal
performance and avoiding service disruptions. This section provides guidance on viewing
your current quota usage, requesting quota increases when necessary, and best practices
for working within CloudWatch service limits.

CloudWatch quotas are integrated with Service Quotas, an AWS service that enables you to view
and manage your quotas from a central location. For more information, see [What Is Service\
Quotas?](https://docs.aws.amazon.com/servicequotas/latest/userguide/intro.html) in the _Service Quotas User Guide_.

AWS Management Console

###### To view CloudWatch service quotas using the console

1. Open the Service Quotas console at [https://console.aws.amazon.com/servicequotas/](https://console.aws.amazon.com/servicequotas).

2. In the navigation pane, choose **AWS**
**services**.

3. From the **AWS services** list, search for and
    select the CloudWatch service whose quotas you want to view.

In the **Service quotas** list, you can see the
    service quota name, applied value (if it is available), AWS
    default quota, and whether the quota value is adjustable.

4. To view additional information about a service quota, such as the
    description, choose the quota name.

5. (Optional) To request a quota increase, select the quota that you
    want to increase, select **Request quota**
**increase**, enter or select the required information, and
    select **Request**.

To work more with service quotas using the console see the [Service Quotas User Guide](https://docs.aws.amazon.com/servicequotas/latest/userguide/intro.html). To learn more about quota increases, see
[Requesting a quota increase](https://docs.aws.amazon.com/servicequotas/latest/userguide/request-quota-increase.html) in the
_Service Quotas User Guide_.

AWS CLI

###### To view CloudWatch service quotas using the AWS CLI

1. Run the following command to view the default CloudWatch quotas.

```nohighlight

aws service-quotas list-aws-default-service-quotas \
       --query 'Quotas[*].{Adjustable:Adjustable,Name:QuotaName,Value:Value,Code:QuotaCode}' \
       --service-code ServiceCode
       --output table
```

###### Note

The following list contains the CloudWatch
**ServiceCode** values:

- To see the CloudWatch quotas, use `--service-code
                                                  monitoring`.

- To see the Amazon AI Operations quotas, use `--service-code
                                                  aiops`.

- To see the CloudWatch Application Signals quotas, use
`--service-code
                                              application-signals`.

- To see the Amazon CloudWatch Internet Monitor quotas, use `--service-code
                                                  internetmonitor`.

- To see the CloudWatch Network Monitor quotas, use
`--service-code networkmonitor`.

- To see the CloudWatch Observability Access Manager quotas,
use `--service-code oam`.

- To see the CloudWatch RUM (Real-User Monitoring) quotas, use
`--service-code rum`.

2. (Optional) Request a quota increase for a CloudWatch service:

1. Identify the quota code for the quota you want to
       increase.

      ```nohighlight

      aws service-quotas list-service-quotas --service-code ServiceCode
      ```

2. Note the `QuotaCode` value for the item. For
       example if you wanted to increase your quota for the
       **Rate of PutCompositeAlarm requests**
       that are supported by CloudWatch for your account, you would
       record the **QuotaCode** `L-8742A250`.

3. Type the following command using the identified
       `ServiceCode` and `QuotaCode`
       values and providing the numeric value for the desired new
       quota:

```nohighlight

aws service-quotas request-service-quota-increase \
    --service-code ServiceCode \
    --quota-code QuotaCode \
    --desired-value new-quota-value

```

For more details about quota increases, see the [`request-service-quota-increase`](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/request-service-quota-increase.html) command
in the [AWS CLI Command Reference](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/index.html#cli-aws-service-quotas).

To work more with service quotas using the AWS CLI, see the [Service Quotas AWS CLI Command Reference](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/index.html#cli-aws-service-quotas).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Grafana integration

Document history
