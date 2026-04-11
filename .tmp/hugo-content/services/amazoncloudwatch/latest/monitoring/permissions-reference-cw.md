# Amazon CloudWatch permissions reference

The following table lists each CloudWatch API operation and the corresponding actions for
which you can grant permissions to perform the action. You specify the actions in the
policy's `Action` field, and you specify a wildcard character (\*) as the
resource value in the policy's `Resource` field.

You can use AWS-wide condition keys in your CloudWatch policies to express conditions. For
a complete list of AWS-wide keys, see [AWS Global and IAM\
Condition Context Keys](../../../iam/latest/userguide/reference-policies-condition-keys.md) in the _IAM User Guide_.

###### Note

To specify an action, use the `cloudwatch:` prefix followed by the API
operation name. For example: `cloudwatch:GetMetricData`,
`cloudwatch:ListMetrics`, or `cloudwatch:*` (for all CloudWatch
actions).

###### Topics

- [CloudWatch API operations and required permissions for actions](#cw-permissions-table)

- [CloudWatch Application Signals API operations and required permissions for actions](#cw-application-signals-permissions-table)

- [CloudWatch Contributor Insights API operations and required permissions for actions](#cw-contributor-insights-permissions-table)

- [CloudWatch Events API operations and required permissions for actions](#cwe-permissions-table)

- [CloudWatch Logs API operations and required permissions for actions](#cwl-permissions-table)

- [Amazon EC2 API operations and required permissions for actions](#cw-ec2-permissions-table)

- [Amazon EC2 Auto Scaling API operations and required permissions for actions](#cw-as-permissions-table)

## CloudWatch API operations and required permissions for actions

CloudWatch API operationsRequired permissions (API actions)

[DeleteAlarms](../../../../reference/amazoncloudwatch/latest/apireference/api-deletealarms.md)

`cloudwatch:DeleteAlarms`

Required to delete an alarm.

[DeleteDashboards](../../../../reference/amazoncloudwatch/latest/apireference/api-deletedashboards.md)

`cloudwatch:DeleteDashboards`

Required to delete a dashboard.

[DeleteMetricStream](../../../../reference/amazoncloudwatch/latest/apireference/api-deletemetricstream.md)

`cloudwatch:DeleteMetricStream`

Required to delete a metric stream.

[DescribeAlarmHistory](../../../../reference/amazoncloudwatch/latest/apireference/api-describealarmhistory.md)

`cloudwatch:DescribeAlarmHistory`

Required to view alarm history. To retrieve information about
composite alarms, your
`cloudwatch:DescribeAlarmHistory` permission must
have a `*` scope. You can't return information about
composite alarms if your
`cloudwatch:DescribeAlarmHistory` permission has
a narrower scope.

[DescribeAlarms](../../../../reference/amazoncloudwatch/latest/apireference/api-describealarms.md)

`cloudwatch:DescribeAlarms`

Required to retrieve information about alarms.

To retrieve information about composite alarms, your
`cloudwatch:DescribeAlarms` permission must have
a `*` scope. You can't return information about
composite alarms if your `cloudwatch:DescribeAlarms`
permission has a narrower scope.

[DescribeAlarmsForMetric](../../../../reference/amazoncloudwatch/latest/apireference/api-describealarmsformetric.md)

`cloudwatch:DescribeAlarmsForMetric`

Required to view alarms for a metric.

[DisableAlarmActions](../../../../reference/amazoncloudwatch/latest/apireference/api-disablealarmactions.md)

`cloudwatch:DisableAlarmActions`

Required to disable an alarm action.

[EnableAlarmActions](../../../../reference/amazoncloudwatch/latest/apireference/api-enablealarmactions.md)

`cloudwatch:EnableAlarmActions`

Required to enable an alarm action.

[GetDashboard](../../../../reference/amazoncloudwatch/latest/apireference/api-getdashboard.md)

`cloudwatch:GetDashboard`

Required to display data about existing dashboards.

[GetMetricData](../../../../reference/amazoncloudwatch/latest/apireference/api-getmetricdata.md)

`cloudwatch:GetMetricData`

Required to graph metric data in the CloudWatch console, to retrieve
large batches of metric data, and perform metric math on that
data.

[GetMetricStatistics](../../../../reference/amazoncloudwatch/latest/apireference/api-getmetricstatistics.md)

`cloudwatch:GetMetricStatistics`

Required to view graphs in other parts of the CloudWatch console and
in dashboard widgets.

[GetMetricStream](../../../../reference/amazoncloudwatch/latest/apireference/api-getmetricstream.md)

`cloudwatch:GetMetricStream`

Required to view information about a metric stream.

[GetMetricWidgetImage](../../../../reference/amazoncloudwatch/latest/apireference/api-getmetricwidgetimage.md)

`cloudwatch:GetMetricWidgetImage`

Required to retrieve a snapshot graph of one or more CloudWatch
metrics as a bitmap image.

[ListDashboards](../../../../reference/amazoncloudwatch/latest/apireference/api-listdashboards.md)

`cloudwatch:ListDashboards`

Required to view the list of CloudWatch dashboards in your
account.

ListEntitiesForMetric

(CloudWatch console-only permission)

`cloudwatch:ListEntitiesForMetric`

Required to find the entities associated with a metric.
Required to explore related telemetry within the CloudWatch
console.

[ListMetrics](../../../../reference/amazoncloudwatch/latest/apireference/api-listmetrics.md)

`cloudwatch:ListMetrics`

Required to view or search metric names within the CloudWatch
console and in the CLI. Required to select metrics on dashboard
widgets.

[ListMetricStreams](../../../../reference/amazoncloudwatch/latest/apireference/api-listmetricstreams.md)

`cloudwatch:ListMetricStreams`

Required to view or search the list of metric streams in the
account.

[PutCompositeAlarm](../../../../reference/amazoncloudwatch/latest/apireference/api-putcompositealarm.md)

`cloudwatch:PutCompositeAlarm`

Required to create a composite alarm.

To create a composite alarm, your
`cloudwatch:PutCompositeAlarm` permission must
have a `*` scope. You can't return information about
composite alarms if your
`cloudwatch:PutCompositeAlarm` permission has a
narrower scope.

[PutDashboard](../../../../reference/amazoncloudwatch/latest/apireference/api-putdashboard.md)

`cloudwatch:PutDashboard`

Required to create a dashboard or update an existing
dashboard.

[PutMetricAlarm](../../../../reference/amazoncloudwatch/latest/apireference/api-putmetricalarm.md)

`cloudwatch:PutMetricAlarm`

Required to create or update an alarm.

[PutMetricData](../../../../reference/amazoncloudwatch/latest/apireference/api-putmetricdata.md)

`cloudwatch:PutMetricData`

Required to create metrics.

[PutMetricStream](../../../../reference/amazoncloudwatch/latest/apireference/api-putmetricstream.md)

`cloudwatch:PutMetricStream`

Required to create a metric stream.

[SetAlarmState](../../../../reference/amazoncloudwatch/latest/apireference/api-setalarmstate.md)

`cloudwatch:SetAlarmState`

Required to manually set an alarm's state.

[StartMetricStreams](../../../../reference/amazoncloudwatch/latest/apireference/api-startmetricstreams.md)

`cloudwatch:StartMetricStreams`

Required to start the flow of metrics in a metric
stream.

[StopMetricStreams](../../../../reference/amazoncloudwatch/latest/apireference/api-startmetricstreams.md)

`cloudwatch:StopMetricStreams`

Required to temporarily stop the flow of metrics in a metric
stream.

[TagResource](../../../../reference/amazoncloudwatch/latest/apireference/api-tagresource.md)

`cloudwatch:TagResource`

Required to add or update tags on CloudWatch resources such as
alarms and Contributor Insights rules.

[UntagResource](../../../../reference/amazoncloudwatch/latest/apireference/api-untagresource.md)

`cloudwatch:UntagResource`

Required to remove tags from CloudWatch resources .

## CloudWatch Application Signals API operations and required permissions for actions

CloudWatch Application Signals API operationsRequired permissions (API actions)

[BatchGetServiceLevelObjectiveBudgetReport](../../../../reference/applicationsignals/latest/apireference/api-batchgetservicelevelobjectivebudgetreport.md)

`application-signals:BatchGetServiceLevelObjectiveBudgetReport`

Required to retrieve service level objective budget
reports.

[CreateServiceLevelObjective](../../../../reference/applicationsignals/latest/apireference/api-createservicelevelobjective.md)

`application-signals:CreateServiceLevelObjective`

Required to create a service level objective (SLO).

[DeleteServiceLevelObjective](../../../../reference/applicationsignals/latest/apireference/api-deleteservicelevelobjective.md)

`application-signals:DeleteServiceLevelObjective`

Required to delete a service level objective (SLO).

[GetService](../../../../reference/applicationsignals/latest/apireference/api-getservice.md)

`application-signals:GetService`

Required to retrieve information about a service discovered by
Application Signals.

[GetServiceLevelObjective](../../../../reference/applicationsignals/latest/apireference/api-getservicelevelobjective.md)

`application-signals:GetServiceLevelObjective`

Required to retrieve information about a service level
objective (SLO).

ListObservedEntities

`application-signals:ListObservedEntities`

Grants permission to list entities that are associated with
other entities.

[ListServiceDependencies](../../../../reference/applicationsignals/latest/apireference/api-listservicedependencies.md)

`application-signals:ListServiceDependencies`

Required to retrieve a list of service dependencies of a
service that you specify. This service and the dependencies were
discovered by Application Signals.

[ListServiceDependents](../../../../reference/applicationsignals/latest/apireference/api-listservicedependents.md)

`application-signals:ListServiceDependents`

Required to retrieve a list of dependents that invoked a
service that you specify. This service and the dependents were
discovered by Application Signals.

[ListServiceLevelObjectives](../../../../reference/applicationsignals/latest/apireference/api-listservicelevelobjectives.md)

`application-signals:ListServiceLevelObjectives`

Required to retrieve a list of service level objectives (SLOs)
in the account.

[ListServiceOperations](../../../../reference/applicationsignals/latest/apireference/api-listserviceoperations.md)

`application-signals:ListServiceOperations`

Required to retrieve a list of service operations of a service
that you specify. This service and the dependencies were
discovered by Application Signals.

[ListServices](../../../../reference/applicationsignals/latest/apireference/api-listservices.md)

`application-signals:ListServices`

Required to retrieve a list of services discovered by
Application Signals.

[ListTagsForResource](../../../../reference/applicationsignals/latest/apireference/api-listtagsforresource.md)

`application-signals:ListTagsForResource`

Required to retrieve a list of the tags associated with a
resource.

[StartDiscovery](../../../../reference/applicationsignals/latest/apireference/api-startdiscovery.md)

`application-signals:StartDiscovery`

Required to be able to enable Application Signals in the
account and create the required service-linked role.

[TagResource](../../../../reference/applicationsignals/latest/apireference/api-tagresource.md)

`application-signals:TagResource`

Required to be able to add tags to resources.

[UntagResource](../../../../reference/applicationsignals/latest/apireference/api-untagresource.md)

`application-signals:UntagResource`

Required to be able to remove tags from resources.

[UpdateServiceLevelObjective](../../../../reference/applicationsignals/latest/apireference/api-updateservicelevelobjective.md)

`application-signals:UpdateServiceLevelObjective`

Required to update an existing service level objective

## CloudWatch Contributor Insights API operations and required permissions for actions

###### Important

When you grant a user the `cloudwatch:PutInsightRule` permission,
by default that user can create a rule that evaluates any log group in CloudWatch Logs.
You can add IAM policy conditions that limit these permissions for a user to
include and exclude specific log groups. For more information, see [Using condition keys to limit Contributor Insights users' access to log groups](iam-cw-condition-keys-contributor.md).

CloudWatch Contributor Insights API operationsRequired permissions (API actions)

[DeleteInsightRules](../../../../reference/amazoncloudwatch/latest/apireference/api-deleteinsightrules.md)

`cloudwatch:DeleteInsightRules`

Required to delete Contributor Insights rules.

[DescribeInsightRules](../../../../reference/amazoncloudwatch/latest/apireference/api-describeinsightrules.md)

`cloudwatch:DescribeInsightRules`

Required to view the Contributor Insights rules in your
account.

[EnableInsightRules](../../../../reference/amazoncloudwatch/latest/apireference/api-enableinsightrules.md)

`cloudwatch:EnableInsightRules`

Required to enable Contributor Insights rules.

[GetInsightRuleReport](../../../../reference/amazoncloudwatch/latest/apireference/api-getinsightrulereport.md)

`cloudwatch:GetInsightRuleReport`

Required to retrieve time series data and other statistics
collectd by Contributor Insights rules.

[PutInsightRule](../../../../reference/amazoncloudwatch/latest/apireference/api-putinsightrule.md)

`cloudwatch:PutInsightRule`

Required to create Contributor Insights rules. See the
**Important** note at the
beginning of this table.

## CloudWatch Events API operations and required permissions for actions

CloudWatch Events API operationsRequired permissions (API actions)

[DeleteRule](../../../../reference/amazoncloudwatchevents/latest/apireference/api-deleterule.md)

`events:DeleteRule`

Required to delete a rule.

[DescribeRule](../../../../reference/amazoncloudwatchevents/latest/apireference/api-describerule.md)

`events:DescribeRule`

Required to list the details about a rule.

[DisableRule](../../../../reference/amazoncloudwatchevents/latest/apireference/api-disablerule.md)

`events:DisableRule`

Required to disable a rule.

[EnableRule](../../../../reference/amazoncloudwatchevents/latest/apireference/api-enablerule.md)

`events:EnableRule`

Required to enable a rule.

[ListRuleNamesByTarget](../../../../reference/amazoncloudwatchevents/latest/apireference/api-listrulenamesbytarget.md)

`events:ListRuleNamesByTarget`

Required to list rules associated with a target.

[ListRules](../../../../reference/amazoncloudwatchevents/latest/apireference/api-listrules.md)

`events:ListRules`

Required to list all rules in your account.

[ListTargetsByRule](../../../../reference/amazoncloudwatchevents/latest/apireference/api-listtargetsbyrule.md)

`events:ListTargetsByRule`

Required to list all targets associated with a rule.

[PutEvents](../../../../reference/amazoncloudwatchevents/latest/apireference/api-putevents.md)

`events:PutEvents`

Required to add custom events that can be matched to
rules.

[PutRule](../../../../reference/amazoncloudwatchevents/latest/apireference/api-putrule.md)

`events:PutRule`

Required to create or update a rule.

[PutTargets](../../../../reference/amazoncloudwatchevents/latest/apireference/api-puttargets.md)

`events:PutTargets`

Required to add targets to a rule.

[RemoveTargets](../../../../reference/amazoncloudwatchevents/latest/apireference/api-removetargets.md)

`events:RemoveTargets`

Required to remove a target from a rule.

[TestEventPattern](../../../../reference/amazoncloudwatchevents/latest/apireference/api-testeventpattern.md)

`events:TestEventPattern`

Required to test an event pattern against a given
event.

## CloudWatch Logs API operations and required permissions for actions

###### Note

CloudWatch Logs permissions can be found in the [CloudWatch Logs user\
guide](../logs/permissions-reference-cwl.md).

## Amazon EC2 API operations and required permissions for actions

Amazon EC2 API operationsRequired permissions (API actions)

[DescribeInstanceStatus](../../../../reference/awsec2/latest/apireference/api-describeinstancestatus.md)

`ec2:DescribeInstanceStatus`

Required to view EC2 instance status details.

[DescribeInstances](../../../../reference/awsec2/latest/apireference/api-describeinstances.md)

`ec2:DescribeInstances`

Required to view EC2 instance details.

[RebootInstances](../../../../reference/awsec2/latest/apireference/api-rebootinstances.md)

`ec2:RebootInstances`

Required to reboot an EC2 instance.

[StopInstances](../../../../reference/awsec2/latest/apireference/api-stopinstances.md)

`ec2:StopInstances`

Required to stop an EC2 instance.

[TerminateInstances](../../../../reference/awsec2/latest/apireference/api-terminateinstances.md)

`ec2:TerminateInstances`

Required to terminate an EC2 instance.

## Amazon EC2 Auto Scaling API operations and required permissions for actions

Amazon EC2 Auto Scaling API operationsRequired permissions (API actions)

Scaling

`autoscaling:Scaling`

Required to scale an Auto Scaling group.

Trigger

`autoscaling:Trigger`

Required to trigger an Auto Scaling action.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS managed policies for Application Insights

Compliance validation

All content copied from https://docs.aws.amazon.com/.
