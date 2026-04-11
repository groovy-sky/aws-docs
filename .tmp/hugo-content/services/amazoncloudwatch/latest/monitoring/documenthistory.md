# Document history

The following table describes important changes for each release of the _Amazon CloudWatch User Guide_, beginning in June 2018. To receive notifications
about updates to this documentation, you can subscribe to an RSS feed.

ChangeDescriptionDate

Updated CloudWatchSyntheticsFullAccess managed policy

The `cloudwatch:ListMetrics` permission was added to the
**CloudWatchSyntheticsFullAccess** policy. Additionally, the
`apigateway:GET` permission has been changed from allowing all resources to specific API Gateway resources: REST APIs, REST API stages, REST API stage Swagger exports, and HTTP
APIs. For more information, see [CloudWatchSyntheticsFullAccess](managed-policies-cloudwatch.md#managed-policies-cloudwatch-CloudWatchSyntheticsFullAccess).

March 31, 2026

Updated AWSObservabilityAdminTelemetryEnablementServiceRolePolicy policy

CloudWatch updated the
**AWSObservabilityAdminTelemetryEnablementServiceRolePolicy** that
grants CloudWatch the permissions necessary to enable and manage telemetry configurations for
the newly supported AWS resources based on telemetry rules. For more information, see
[AWSObservabilityAdminTelemetryEnablementServiceRolePolicy](using-service-linked-roles.md#service-linked-role-telemetry-enablement).

March 31, 2026

Updated AWSObservabilityAdminTelemetryEnablementServiceRolePolicy policy

CloudWatch updated the
**AWSObservabilityAdminTelemetryEnablementServiceRolePolicy** that
grants CloudWatch the permissions necessary to enable and manage telemetry configurations for
the newly supported AWS resources based on telemetry rules. For more information, see
[AWSObservabilityAdminTelemetryEnablementServiceRolePolicy](using-service-linked-roles.md#service-linked-role-telemetry-enablement).

March 10, 2026

Update to AIOPsAssistantPolicy managed policy

CloudWatch updated the **AIOpsAssistantPolicy** to allow CloudWatch investigations access to
additional resources to assist in queries, troubleshooting, and topology mapping. For more
information, see [AIOpsAssistantPolicy](managed-policies-cloudwatch.md#managed-policies-QInvestigations-AIOpsAssistant).

February 6, 2026

Update to AIOPsAssistantPolicy managed policy

CloudWatch updated the **AIOpsAssistantPolicy** to allow CloudWatch investigations access to
additional resources to assist in queries, troubleshooting, and topology mapping. For more
information, see [AIOpsAssistantPolicy](managed-policies-cloudwatch.md#managed-policies-QInvestigations-AIOpsAssistant).

December 17, 2025

Updated CloudWatchNetworkMonitorServiceRolePolicy managed policy with TGW permissions

CloudWatch added the `ec2:DescribeRouteTables`,
`ec2:DescribeTransitGatewayAttachments`,
`ec2:DescribeTransitGatewayRouteTables`,
`ec2:SearchTransitGatewayRoutes` permissions to
the [CloudWatchNetworkMonitorServiceRolePolicy](security-iam-awsmanpol-nw.md#security-iam-CloudWatchNetworkMonitorServiceRolePolicy)
managed policy to grant permissions for Network Synthetic Monitor to generate the
Network Health Indicator values for customers utilizing Transit Gateways.

December 12, 2025

Updated CloudWatchNetworkFlowMonitorTopologyServiceRolePolicy managed policy with VPC Endpoints and vpce service config permissions

CloudWatch added the `ec2:DescribeVpcEndpoints` and `ec2:DescribeVpcEndpointServiceConfigurations`
permissions to the CloudWatchNetworkFlowMonitorTopologyServiceRolePolicy managed policy to grant permissions
for Network Flow Monitor to generate topology snapshots of VPC endpoints and vpce service config.

December 10, 2025

Updated CloudWatchFullAccessV2 policy

Permissions for observability administration actions were added to [CloudWatchFullAccessV2](managed-policies-cloudwatch.md#managed-policies-cloudwatch-CloudWatchFullAccessV2) to allow full access to telemetry pipelines and S3 table
integrations.

December 2, 2025

Observability for Amazon Bedrock AgentCore policies

You can now monitor invocation metrics for Amazon Bedrock AgentCore policies using Amazon CloudWatch. For
more information, see [Gateways](gateways.md).

December 2, 2025

Observability for Amazon Bedrock AgentCore Evaluations

You can now monitor and assess the performance, quality, and reliability of your AI
agents using the Evaluations dashboard. For more information, see [Agent details -\
Evaluations](session-traces-evaluations.md).

December 2, 2025

New AWSObservabilityAdminTelemetryEnablementServiceRolePolicy policy

CloudWatch updated the
**AWSObservabilityAdminTelemetryEnablementServiceRolePolicy** that
grants CloudWatch the permissions necessary to enable and manage telemetry configurations for
the newly supported AWS resources based on telemetry rules. For more information, see
[AWSObservabilityAdminTelemetryEnablementServiceRolePolicy](using-service-linked-roles.md#service-linked-role-telemetry-enablement).

December 2, 2025

Introduced CloudWatch pipelines

The new pipeline feature of CloudWatch connects data from a catalogue of data sources,
so that you can add and configure pipeline processors from a library to parse, enrich, and
standardize data. For more information, see [CloudWatch\
pipelines](cloudwatch-pipelines.md).

December 2, 2025

Introduced CloudWatch pipelines

The new pipeline feature of CloudWatch connects data from a catalogue of data sources,
so that you can add and configure pipeline processors from a library to parse, enrich, and
standardize data. For more information, see [CloudWatch\
pipelines](cloudwatch-pipelines.md).

December 2, 2025

New CloudWatch investigations feature

Amazon CloudWatch now offers interactive incident report generation, enabling
customers to create comprehensive post-incident analysis reports in minutes. The new
capability, available within CloudWatch investigations, automatically gathers and
correlates your telemetry data, as well as your input and any actions taken during an
investigation, and produces a streamlined incident report. Amazon CloudWatch enhanced
incident report generation capabilities with an AI-powered root-cause workflow that guides
customers through the "Five Why’s" analysis technique. For more information, see [Generate incident reports](investigations-incident-reports.md)

November 30, 2025

Updated CloudWatchReadOnlyAccess policy

CloudWatch updated [CloudWatchReadOnlyAccess](managed-policies-cloudwatch.md#managed-policies-cloudwatch-CloudWatchReadOnlyAccess) to include Resource Explorer permissions to support
uninstrumented service functionalities in Application Signals.

November 20, 2025

Updated CloudWatchFullAccessV2 policy

CloudWatch updated [CloudWatchFullAccessV2](managed-policies-cloudwatch.md#managed-policies-cloudwatch-CloudWatchFullAccessV2) to include Resource Explorer permissions to support
uninstrumented service functionalities in Application Signals.

November 20, 2025

Updated CloudWatchApplicationSignalsServiceRolePolicy policy

CloudWatch updated the policy to include `resource-explorer-2:Search` and
`cloudtrail:CreateServiceLinkedChannel` to enable new Application Signals
features. For more information, see [CloudWatchApplicationSignalsServiceRolePolicy](using-service-linked-roles.md#service-linked-role-signals).

November 20, 2025

Updated CloudWatchApplicationSignalsReadOnlyAccess policy

CloudWatch updated [CloudWatchApplicationSignalsReadOnlyAccess](managed-policies-cloudwatch.md#managed-policies-cloudwatch-CloudWatchApplicationSignalsReadOnlyAccess) to include new Application Signals
operation and OAM and Resource Explorer permissions to support change events, cross
account and uninstrumented service functionalities in Application Signals.

November 20, 2025

Updated CloudWatchApplicationSignalsFullAccess policy

CloudWatch updated [CloudWatchApplicationSignalsFullAccess](managed-policies-cloudwatch.md#managed-policies-cloudwatch-CloudWatchApplicationSignalsFullAccess) to include OAM and Resource Explorer
permissions to support cross account and uninstrumented service functionalities in
Application Signals.

November 20, 2025

CloudWatch updated AIOpsConsoleAdminPolicy managed policy with Amazon Q permissions

CloudWatch updated the **AIOpsConsoleAdminPolicy** managed policy to
include Amazon Q integration permissions. These permissions enable users to interact with
CloudWatch investigations incident reports through Amazon Q's conversational interface. For more information
see, [AIOpsConsoleAdminPolicy](auth-and-access-control-cw.md#managed-policies-QInvestigations-AIOpsConsoleAdminPolicy).

November 17, 2025

CloudWatch updated AIOpsOperatorAccess managed policy with Amazon Q permissions

CloudWatch updated the **AIOpsOperatorAccess** managed policy to include
Amazon Q integration permissions. These permissions enable users to interact with CloudWatch investigations
incident reports through Amazon Q's conversational interface. For more information see,
[AIOpsOperatorAccess](auth-and-access-control-cw.md#managed-policies-QInvestigations-AIOpsOperatorAccess).

November 7, 2025

Updated CloudWatchNetworkFlowMonitorTopologyServiceRolePolicy managed policy with managed prefix lists permissions

CloudWatch added the `ec2:DescribeManagedPrefixLists` and
`ec2:GetManagedPrefixListEntries` permissions to the
CloudWatchNetworkFlowMonitorTopologyServiceRolePolicy managed policy to grant permissions
for Network Flow Monitor to generate topology snapshots of managed prefix lists.

November 6, 2025

Updated CloudWatchReadOnlyAccess Policy with observability administration permissions

CloudWatch updated the **CloudWatchReadOnlyAccess** managed policy to
include observability administration permissions. These permissions provide read-only
access to telemetry rules, centralization configurations, and resource telemetry data
across AWS Organizations. For more information, see [CloudWatchReadOnlyAccess](auth-and-access-control-cw.md#managed-policies-cloudwatch-CloudWatchReadOnlyAccess).

October 10, 2025

CloudWatch updated AIOpsOperatorAccess managed policy with incident report permissions

CloudWatch updated the **AIOpsOperatorAccess** managed policy to include
incident report generation permissions. These permissions allow users to create and manage
incident reports and work with AI-derived facts in CloudWatch investigations. For more information, see [AIOpsOperatorAccess](auth-and-access-control-cw.md#managed-policies-QInvestigations-AIOpsOperatorAccess).

October 10, 2025

CloudWatch investigations added incident reports

CloudWatch updated CloudWatch investigations to provide incident reporting. These AI-generated reports capture
investigation findings, root causes, timeline events, and recommended corrective actions
in a structured format. For more information see, [Generate\
incident reports](investigations-incident-reports.md).

October 10, 2025

CloudWatch added AIOpsAssistantIncidentReportPolicy managed policy to support CloudWatch investigations AI-generated incident reports

CloudWatch added the **AIOpsAssistantIncidentReportPolicy** managed policy
to enable CloudWatch investigations to generate incident reports from investigation data. This policy provides
permissions to access investigations, create and manage reports, and work with AI-derived
facts that form the foundation of incident documentation. For more information, see [AIOpsAssistantIncidentReportPolicy](auth-and-access-control-cw.md#managed-policies-QInvestigations-AIOpsAssistantIncidentReportPolicy).

October 10, 2025

Updated CloudWatchReadOnlyAccess policy with CloudTrail and Service Quotas permissions

CloudWatch updated [CloudWatchReadOnlyAccess](auth-and-access-control-cw.md#managed-policies-cloudwatch-CloudWatchReadOnlyAccess) to include CloudTrail and Service Quotas permissions to support top
observations and change indicators functionality within Application Signals.

October 8, 2025

Updated CloudWatchFullAccessV2 policy

CloudWatch updated [CloudWatchFullAccessV2](auth-and-access-control-cw.md#managed-policies-cloudwatch-CloudWatchFullAccessV2) to include CloudTrail and Service Quotas permissions to support top
observations and change indicators functionality within Application Signals.

October 8, 2025

Updated CloudWatchApplicationSignalsReadOnlyAccess policy

CloudWatch updated the `CloudWatchApplicationSignalsReadOnlyAccess` policy to
view how resources and services change in your account and view top observations for
service anomalies in your account. For more information, see [CloudWatchApplicationSignalsReadOnlyAccess](../../../aws-managed-policy/latest/reference/cloudwatchapplicationsignalsreadonlyaccess.md).

September 29, 2025

Updated CloudWatchApplicationSignalsFullAccess policy

CloudWatch updated the `CloudWatchApplicationSignalsFullAccess` policy to view
how resources and services change in your account and view top observations for service
anomalies in your account. For more information, see [CloudWatchApplicationSignalsFullAccess](../../../aws-managed-policy/latest/reference/cloudwatchapplicationsignalsfullaccess.md).

September 29, 2025

CloudWatch Application Signals added support for dynamic service grouping and filtering, last deployment tracking, and automated audit findings

You can group and filter services with Application Signals' dynamic grouping
capabilities track deployment activities through Application Signals' automatic processing
of CloudTrail events, and discover critical insights through Application Signals' automated
audit findings. For more information, see [Application Signals](cloudwatch-application-monitoring-sections.md).

September 29, 2025

Update to AIOPsAssistantPolicy managed policy

CloudWatch updated the **AIOpsAssistantPolicy** to allow CloudWatch investigations access to
additional resources to assist in queries, troubleshooting, and topology mapping. For more
information, see [AIOpsAssistantPolicy](managed-policies-cloudwatch.md#managed-policies-QInvestigations-AIOpsAssistant).

September 24, 2025

New AWSObservabilityAdminLogsCentralizationServiceRolePolicy policy

CloudWatch added the
**AWSObservabilityAdminLogsCentralizationServiceRolePolicy** that
grants CloudWatch permission to use telemetry data from other AWS accounts that you specify to
create CloudWatch log groups, log streams, and log events in your organization's monitoring
account. The SLR only provides the assume role permission to allow the CloudWatch service to
assume the role in the monitoring account. For more information, see [AWSObservabilityAdminLogsCentralizationServiceRolePolicy](auth-and-access-control-cw.md#service-linked-role-logscentralization).

September 17, 2025

CloudWatch added support for Cross-account cross-Region log centralization

CloudWatch added support for the **Cross-account cross-Region log**
**centralization** that enables CloudWatch Logs to gather log data from AWS Organizations member
accounts into a central account for analysis. For more information, see [Monitor\
across accounts and Regions](cloudwatch-cross-account-methods.md).

September 17, 2025

[Collect EC2 instance store volume NVMe driver metrics](container-insights-metrics-instance-store-collect.md)

You can now use the CloudWatch agent to collect AWS NVMe driver metrics for instance store
volumes attached to Nitro-based Amazon EC2 instances.

September 16, 2025

Removed AWS OpsWorks metrics and dimensions

AWS OpsWorks reached end-of-life and the service is deprecated.

August 20, 2025

New AWSObservabilityAdminTelemetryEnablementServiceRolePolicy policy

CloudWatch added the
**AWSObservabilityAdminTelemetryEnablementServiceRolePolicy** that
grants CloudWatch the permissions necessary to enable and manage telemetry configurations for
AWS resources based on telemetry rules. For more information, see [AWSObservabilityAdminTelemetryEnablementServiceRolePolicy](using-service-linked-roles.md#service-linked-role-telemetry-enablement).

July 17, 2025

Gen AI observability

With CloudWatch, you can observe generative AI workloads, including Amazon Bedrock AgentCore
agents, and gain insights into AI performance, health, and accuracy. For more information,
see [Generative AI\
observability](genai-observability.md).

July 16, 2025

AIOpsReadOnlyAccess policy updated

CloudWatch updated the **AIOpsReadOnlyAccess** policy to permit validation
of investigation groups across accounts. For more information, see [AIOpsReadOnlyAccess policy](managed-policies-cloudwatch-qinvestigations.md#managed-policies-QInvestigations-AIOpsReadOnlyAccess).

June 23, 2025

Updated AIOpsConsoleAdminPolicy policy

CloudWatch updated the AIOpsConsoleAdminPolicy policy to permit validation of investigation
groups across accounts. For more information, see [AIOpsConsoleAdminPolicy policy](managed-policies-cloudwatch-qinvestigations.md#managed-policies-QInvestigations-AIOpsConsoleAdminPolicy).

June 13, 2025

Updated AIOpsAssistantPolicy policy

CloudWatch updated the AIOpsAssistantPolicy policy to provide additional IAM permissions.
For more information, see [AIOpsConsoleAdminPolicy policy](managed-policies-cloudwatch-qinvestigations.md#managed-policies-QInvestigations-AIOpsAssistant).

June 13, 2025

Update managed policy

CloudWatch updated the **AIOpsOperatorAccess** policy to permit validation
of investigation groups across accounts. For more information, see [AIOpsOperatorAccess](managed-policies-cloudwatch-qinvestigations.md#managed-policies-QInvestigations-AIOpsOperatorAccess.html).

June 13, 2025

[AmazonCloudWatchRUMReadOnlyAccess updated](auth-and-access-control-cw.md#managed-policies-cloudwatch-RUM)

The `synthetics: describeCanaries`,
`synthetics:describeCanariesLastRun`, `cloudwatch:GetMetricData`,
`cloudwatch:DescribeAlarms`, `logs:DescribeLogGroups`,
`xray:GetTraceSummaries`, and `rum:ListTagsForResources`
permissions was added to the CloudWatch RUM `AmazonCloudWatchRUMReadOnlyAccess`
policy.

June 12, 2025

Added functionality

You can now monitor Amazon Aurora PostgreSQL Limitless Databases at the instance level
from the Database Instance Dashboard. For more information, see [Monitoring\
Aurora Limitless databases with Database Insights](database-insights-limitless.md).

May 22, 2025

Added functionality

You can now enable Application Signals for your applications on an existing Amazon EKS
cluster using CloudWatch Observability add-on Auto monitor advanced configuration. For more
information, see [Enable Application Signals on an Amazon EKS cluster using the CloudWatch Observability add-on\
advanced configuration](cloudwatch-application-signals-enable-eks.md).

May 19, 2025

CloudWatch Network Monitoring releases multiple account support for Network Flow Monitor

Network Flow Monitor, a feature of CloudWatch Network Monitoring, now supports monitoring
network performance metrics for network flows for resources in multiple accounts. For more
information, see [Initialize Network Flow Monitor for multi-account monitoring](cloudwatch-networkflowmonitor-multi-account.md).

May 6, 2025

[AmazonCloudWatchRUMReadOnlyAccess updated](auth-and-access-control-cw.md#managed-policies-cloudwatch-RUM)

The `rum:GetResourcePolicy` permission was added so that CloudWatch RUM can view
the resource policy attached to the RUM application monitor.

April 28, 2025

[Updated AIOpsAssistantPolicy policy](auth-and-access-control-cw.md#managed-policies-QInvestigations-AIOpsAssistant)

CloudWatch added permissions to the **AIOpsAssistantPolicy** policy, which
Amazon Q Developer operational investigations uses to find information during
investigations.

April 11, 2025

[CloudWatch Synthetics deprecates five runtime metrics on October 1, 2025](cloudwatch-synthetics-library-nodejs-puppeteer.md#CloudWatch_Synthetics_runtimeversion-nodejs-puppeteer-6.0)

CloudWatch Synthetics metrics `syn-nodejs-puppeteer-7.0`,
`syn-nodejs-puppeteer-6.2`, `syn-nodejs-puppeteer-5.2`,
`syn-python-selenium-3.0`, and `syn-python-selenium-2.1` will be
deprecated on October 1, 2025.

April 4, 2025

Updated CloudWatchApplicationSignalsServiceRolePolicy policy

CloudWatch updated the policy to exclude time windows from impacting the SLO attainment
rate, error budget, and burn rate metrics. CloudWatch can retrieve exclusion windows on behalf
of you. For more information, see [CloudWatchApplicationSignalsServiceRolePolicy](using-service-linked-roles.md#service-linked-role-signals).

March 13, 2025

Time window exclusions

SLOs adds support for time window exclusions which is a block of time with a defined
start and end date. This time period is excluded from the SLO's performance metrics and
you can schedule one-time or recurring time exclusions windows. For more information, see
[Service\
level objectives (SLOs)](cloudwatch-servicelevelobjectives.md).

March 13, 2025

[New AIOpsReadOnlyAccess policy](auth-and-access-control-cw.md#managed-policies-QInvestigations-AIOpsReadOnlyAccess)

CloudWatch added the **AIOpsReadOnlyAccess** policy, which grants read-only
permissions for Amazon AI Operations and other related services.

December 3, 2024

[New AIOpsOperatorAccess policy](auth-and-access-control-cw.md#managed-policies-QInvestigations-AIOpsOperatorAccess)

CloudWatch added the **AIOpsOperatorAccess** policy, which grants users
access to Amazon AI Operations actions, and to additional AWS actions that are necessary
for accessing investigation events.

December 3, 2024

[New AIOpsConsoleAdminPolicy policy](auth-and-access-control-cw.md#managed-policies-QInvestigations-AIOpsConsoleAdminPolicy)

CloudWatch added the **AIOpsConsoleAdminPolicy** policy, which grants users
full administrative access for managing Amazon Q investigations, including the management of
trusted identity propagation, and the management of IAM Identity Center and organizational
access.

December 3, 2024

[New AIOpsAssistantPolicy policy](auth-and-access-control-cw.md#managed-policies-QInvestigations-AIOpsAssistant)

CloudWatch added the **AIOpsAssistantBasicAccess** policy, which you can
assign to the Amazon AI Operations assistant to enable the assistant to have basic
functionality.

December 3, 2024

[Amazon Q Developer operational investigations feature in Preview](investigations.md)

CloudWatch released a Preview of Amazon Q Developer operational investigations. This feature is a
generative AI-powered assistant that can help you respond to incidents in your system. It
uses generative AI to scan your system's telemetry and quickly surface suggestions that
might be related to your issue, to help you troubleshoot operational issues faster.

December 3, 2024

[CloudWatchFullAccessV2 and CloudWatchFullAccess policies updated](auth-and-access-control-cw.md#managed-policies-cloudwatch-CloudWatchFullAccessV2)

CloudWatch updated both **CloudWatchFullAccessV2** and
**CloudWatchFullAccess**. Permissions for
Amazon OpenSearch Service were added to to enable CloudWatch Logs integration with OpenSearch Service for
some features.

December 1, 2024

[Telemetry config feature generally available](telemetry-config-cloudwatch.md)

CloudWatch added the telemetry config feature to discover and understand the state of
telemetry configuration for your AWS resources.

November 26, 2024

[New AWSServiceRoleForObservabilityAdmin service-linked role](using-service-linked-roles.md#service-linked-role-telemetry-config)

CloudWatch added the new AWSServiceRoleForObservabilityAdmin service-linked role and
AWSObservabilityAdminServiceRolePolicy policy.

November 26, 2024

[CloudWatch supports exploring related telemetry](explorerelated.md)

CloudWatch supports exploring related telemetry data across interconnected resources.

November 22, 2024

[Updated CloudWatchSyntheticsFullAccess IAM policy](auth-and-access-control-cw.md#managed-policies-cloudwatch-canaries)

CloudWatch updated the **CloudWatchSyntheticsFullAccess** policy to grant
CloudWatch Synthetics access to get AWS Lambda functions, and perform actions on logs available
in CloudWatch Logs. Additionally, Lambda layer version actions apply to all CloudWatch Synthetics layer
ARNs.

November 21, 2024

[New CloudWatch Synthetics support for Node.js with the Playwright runtime](cloudwatch-synthetics-canaries-nodejs-playwright.md)

CloudWatch Synthetics adds support for Node.js canary scripts that use the Playwright
runtime.

November 21, 2024

CloudWatchReadOnlyAccess policy updated

The scope of the policy was changed to include the `xray:List`\*,
`xray:StartTraceRetrieval`, and `xray:CancelTraceRetrieval`
permissions. This change allows CloudWatch to access new X-Ray read APIs. For more infomration,
see [CloudWatchReadOnlyAccess](auth-and-access-control-cw.md#managed-policies-cloudwatch-CloudWatchReadOnlyAccess).

November 21, 2024

CloudWatch Transaction Search is generally available

CloudWatch releases Transaction Search, a search and analytics experience. Transaction
Search provides visibility into application transaction issues. Explore interactions
between application components to establish a root cause. Understand your application
transactions and their impact on customers. Search transactions using terms like customer
name or order. For more information, see [CloudWatch\
Transaction Search](cloudwatch-transaction-search.md).

November 21, 2024

[Application Signals supports Lambda functions](cloudwatch-application-signals-enable-lambda.md)

You can now enable Application Signals for Lambda functions.

November 21, 2024

[OpenTelemetry](cloudwatch-opentelemetry-sections.md)

You can use OpenTelemetry to directly send traces to an OpenTelemetry Protocol (OTLP)
endpoint, and get out-of-the box application performance monitoring experiences in
CloudWatch Application Signals.

November 20, 2024

[Application Signals collects runtime metrics](appsignals-runtimemetrics.md)

Application Signals now collects OpenTelemetry-compatible runtime metrics from your
Java and Python applications.

November 19, 2024

[Application Signals collects runtime metrics](appsignals-runtimemetrics.md)

Application Signals now collects OpenTelemetry-compatible runtime metrics from your
Java and Python applications.

November 19, 2024

[Internet traffic monitoring for VPCs in AWS Local Zones](cloudwatch-internetmonitor-regions.md)

Internet Monitor now supports internet traffic monitoring for VPCs deployed in AWS Local
Zones

November 18, 2024

[Application Signals support for Node.js applications is generally available](cloudwatch-application-signals-enable.md)

CloudWatch supports enabling Node.js applications for Application Signals.

November 15, 2024

[New CloudWatchInternetMonitorReadOnlyAccess IAM policy](cloudwatch-im-permissions.md#security-iam-awsmanpol-CloudWatchInternetMonitorReadOnlyAccess)

CloudWatch added a **CloudWatchInternetMonitorReadOnlyAccess** policy, to
grant read-only access to actions and resources available in the CloudWatch console for
Internet Monitor.

November 14, 2024

[CloudWatch added observability solutions](monitoring-solutions.md)

CloudWatch observability solutions offer a catalog of readily available configurations to
help you quickly implement monitoring for various AWS services and common workloads,
such as Java Virtual Machines (JVM), Apache Kafka, Apache Tomcat, and NGINX.

November 13, 2024

[CloudWatch added burn rate monitoring for Application Signals SLOs](cloudwatch-servicelevelobjectives-burn.md)

CloudWatch added burn rate monitoring to Application Signals SLOs. A burn rate is a metric
that indicates how fast the service is consuming the error budget, relative to the
attainment goal of the SLO.

November 13, 2024

[New Internet Monitor latency reduction suggestions through Route 53 routing updates](cloudwatch-im-insights.md)

Amazon CloudWatch Internet Monitor now provides actionable suggestions to help you optimize your Route 53 IP-based
routing configurations. Easily identify the optimal AWS Regions for routing your end
user traffic, and then configure your Route 53 IP-based routing based on the
recommendations.

November 1, 2024

[New CloudWatchInternetMonitorFullAccess IAM policy](cloudwatch-im-permissions.md#security-iam-awsmanpol-CloudWatchInternetMonitorFullAccess)

CloudWatch added a **CloudWatchInternetMonitorFullAccess** policy, to grant
full access to actions and resources available in the CloudWatch console for Internet Monitor.

October 23, 2024

End of support notice

End of support notice: On October 16, 2025, AWS will discontinue support for CloudWatch
Evidently. After October 16, 2025, you will no longer be able to access the Evidently
console or Evidently resources.

October 17, 2024

End of support notice

End of support notice: On October 16, 2025, AWS will discontinue support for CloudWatch
Evidently. After October 16, 2025, you will no longer be able to access the Evidently
console or Evidently resources.

October 17, 2024

[New CloudWatchLambdaApplicationSignalsExecutionRolePolicy IAM policy](auth-and-access-control-cw.md#managed-policies-CloudWatchLambdaApplicationSignalsExecutionRolePolicy)

CloudWatch added a new policy named
**CloudWatchLambdaApplicationSignalsExecutionRolePolicy**. This policy
is used when CloudWatch Application Signals is enabled for Lambda workloads. It enables write
access to X-Ray and the log group used by CloudWatch Application Signals.

October 16, 2024

[New CloudWatchLambdaApplicationSignalsExecutionRolePolicy IAM policy](auth-and-access-control-cw.md#managed-policies-CloudWatchLambdaApplicationSignalsExecutionRolePolicy)

CloudWatch added a new policy named
**CloudWatchLambdaApplicationSignalsExecutionRolePolicy**. This policy
is used when CloudWatch Application Signals is enabled for Lambda workloads. It enables write
access to X-Ray and the log group used by CloudWatch Application Signals.

October 16, 2024

[CloudWatchSyntheticsFullAccess policy updated](auth-and-access-control-cw.md#managed-policies-cloudwatch-CloudWatchSyntheticsFullAccess)

Permissions were added to the **CloudWatchSyntheticsFullAccess**
IAM policy. The `lambda:ListTags`, `lambda:TagResource`, and
`lambda:UntagResource` permissions were added so that when you apply or
change tags on a canary, you can choose to have Synthetics also apply those same tags or
changes to the Lambda function that the canary uses.

October 11, 2024

Application Signals supports request-based SLOs

Application Signals adds support for request-based service level objects (SLOs). For
more information, see [Service\
level objectives (SLOs)](cloudwatch-servicelevelobjectives.md).

September 6, 2024

[Internet Monitor refreshed dashboard and latency improvement suggestions](cloudwatch-im-monitor-and-optimize.md)

Amazon CloudWatch Internet Monitor has launched an updated console experience, including new features for
visualizing configuration changes that can help you reduce latency for your application.

August 12, 2024

Application Signals previewing support for .NET applications

CloudWatch Application Signals has added [support for .NET applications](cloudwatch-application-signals-supportmatrix.md), with ADOT Instrumentation for .NET on supported
platforms. This feature is in preview release.

August 5, 2024

[New CloudWatch Application Insights service linked role permission added](security-iam-awsmanpol-appinsights.md#security-iam-awsmanpol-appinsights-updates)

Allows CloudWatch Application Insights to enable and disable termination protection on CloudFormation stacks.

July 25, 2024

CloudWatch Metrics Insights support for natural language query generation is generally available.

CloudWatch Metrics Insights supports natural language query to generate and update queries.
For more information, see [Use natural language to generate and update CloudWatch Metric Insights queries](cloudwatch-metrics-insights-query-assist.md).

June 10, 2024

CloudWatch Application Signals is generally available

CloudWatch Application Signals is now generally available. Use Application Signals to
instrument your applications on AWS so that you can monitor current application health,
create service level objectives (SLOs), and track long-term application performance
against your business objectives. For more information, see [Application Signals](cloudwatch-application-monitoring-sections.md).

June 10, 2024

[New CloudWatchApplicationSignalsReadOnlyAccess IAM policy](auth-and-access-control-cw.md#managed-policies-cloudwatch-ApplicationSignals)

CloudWatch added a **CloudWatchApplicationSignalsReadOnlyAccess** policy to
add read only actions and resources available in the CloudWatch console for Application
Signals.

June 7, 2024

[New CloudWatchApplicationSignalsFullAccess policy](auth-and-access-control-cw.md#managed-policies-cloudwatch-ApplicationSignals)

CloudWatch added a **CloudWatchApplicationSignalsFullAccess** policy to add
actions and resources available in the CloudWatch console for Application Signals.

June 7, 2024

[CloudTrail now captures API activities related to CloudWatch data plane operations.](logging-cw-api-calls.md#CloudWatch-data-plane-events)

CloudTrail now logs events in CloudTrail for **GetMetricData** and
**GetMetricWidgetImage** API activities.

June 6, 2024

[CloudWatch Application Signals application map supports canary, RUM clients and AWS service dependency groupings.](servicemap.md)

The Application Signals preview release has added default groupings in the application
map for canaries, RUM clients, and AWS service dependencies of the same type. This
change reduces the number of icons in the application map default view to make it easier
to view and navigate.

May 21, 2024

[CloudWatchReadOnlyAccess policy updated](auth-and-access-control-cw.md#managed-policies-cloudwatch-CloudWatchReadOnlyAccess)

CloudWatch changed the scope of a permission in
**CloudWatchReadOnlyAccess**. The scope of the policy added the
`application-signals:BatchGet*`, `application-signals:Get*`, and
`application-signals:List*` actions so that users can use CloudWatch Application
Signals to view, investigate, and diagnose issues with the health of their services. CloudWatch
also added an `iam:GetRole` action so that users can check if Application
Signals is set up.

May 17, 2024

[CloudWatchFullAccessV2 policy updated](auth-and-access-control-cw.md#managed-policies-cloudwatch-CloudWatchFullAccessV2)

CloudWatch changed the scope of a permission in **CloudWatchFullAccessV2**.
The scope of the policy added the `application-signals:*` so that users can use
CloudWatch Application Signals to view, investigate, and diagnose issues with the health of
their services.

May 17, 2024

[Lambda Insights supports AWS GovCloud (US-East) and AWS GovCloud (US-West)](lambda-insights.md)

CloudWatch Lambda Insights has added support for the AWS GovCloud (US-East) and
AWS GovCloud (US-West) Regions.

April 29, 2024

[CloudWatch cross-account observability supports resource filters](cloudwatch-unified-cross-account.md)

You can now create filters to specify which metric namespaces and log groups are
shared from the source account to the monitoring account, when you create the link between
the accounts.

April 26, 2024

[CloudWatch Application Signals updates](cloudwatch-application-monitoring-sections.md)

The Application Signals preview release has added three features. Application Signals
now supports Python applications. It offers a simpler enablement process for applications
on Amazon EKS architectures. And it includes new configurations that you can use to manage the
cardinality of metrics that are collected.

April 26, 2024

[CloudWatch Container Insights with enhanced observability for Amazon EKS can collect AWS Elastic Fabric Adapter (EFA) metrics](container-insights-metrics-eks.md#Container-Insights-metrics-EFA)

You can now use CloudWatch Container Insights with enhanced observability for Amazon EKS collect
AWS Elastic Fabric Adapter (EFA) metrics from Amazon EKS clusters.

April 23, 2024

[**Updated IAM policy**](using-service-linked-roles.md#service-linked-role-updates)

CloudWatch updated the **CloudWatchApplicationSignalsServiceRolePolicy**
policy. The scoping of the `logs:StartQuery` and
`logs:GetQueryResults` permissions in this policy was changed to add
`arn:aws:logs:*:*:log-group:/aws/appsignals/*:*` and
`"arn:aws:logs:*:*:log-group:/aws/application-signals/data:*"` to enable
Application Signals on more architecures. This policy is attached to the
**AWSServiceRoleForCloudWatchApplicationSignals** service-linked
role.

April 18, 2024

[Internet Monitor provides a global internet weather map to authenticated AWS customers](cloudwatch-internetmonitor-outage-map.md)

Amazon CloudWatch Internet Monitor now displays a global internet weather map that is available in the console
to all authenticated AWS customers. To view the map, in the Amazon CloudWatch console, navigate
to Internet Monitor.

April 16, 2024

[CloudWatch Container Insights with enhanced observability for Amazon EKS can collect AWS Neuron metrics](container-insights-metrics-eks.md#Container-Insights-metrics-EKS-Neuron)

You can now use CloudWatch Container Insights with enhanced observability for Amazon EKS collect
AWS Neuron metrics from Amazon EKS clusters.

April 16, 2024

[CloudWatch Application Signals adds a **Service overview** tab and more metrics to aid in diagnostics](servicedetail.md)

A new **Service overview** tab displays an overview of your service,
including number of operations, dependencies, synthetics, and client pages. The tab shows
key metrics for your entire service, and top operations and dependencies. You can also now
view X-Ray traces that are correlated with issues including faults, errors, and latency
problems.

April 16, 2024

[CloudWatch Container Insights with enhanced observability for Amazon EKS adds support for Windows](containerinsights.md#container-insights-detailed-metrics)

You can now use CloudWatch Container Insights with enhanced observability for Amazon EKS collect
metrics from Windows worker nodes on Amazon EKS clusters.

April 10, 2024

[**CloudWatchApplicationSignalsServiceRolePolicy** IAM policy updated](using-service-linked-roles.md#service-linked-role-signals)

CloudWatch changed the scope of a permission in
**CloudWatchApplicationSignalsServiceRolePolicy**. The scope of the
`cloudwatch:GetMetricData` permission was changed to `*` so that
Application Signals can retrieve metrics from sources in linked accounts.

April 8, 2024

[Amazon CloudWatch Internet Monitor now supports cross-account observability](cwim-cross-account.md)

You can now use Internet Monitor cross-account observability to monitor your applications that
span multiple AWS accounts within a single AWS Region.

March 29, 2024

[**CloudWatchAgentServerPolicy** and **CloudWatchAgentAdminPolicy** policies updated](auth-and-access-control-cw.md#managed-policies-cloudwatch-CloudWatchAgentServerPolicy)

CloudWatch added permissions to both the `CloudWatchAgentServerPolicy` and
`CloudWatchAgentAdminPolicy` policies to allow the CloudWatch agent to publish
X-Ray traces and modify log group retention periods. In both policies, the
`xray:PutTraceSegments`, `xray:PutTelemetryRecords`,
`xray:GetSamplingRules`, `xray:GetSamplingTargets`,
`xray:GetSamplingStatisticSummaries` and `logs:PutRetentionPolicy`
permissions were added

February 12, 2024

[New service linked role and IAM policy for CloudWatch Network Monitor](monitoring-using-service-linked-roles-nw.md)

CloudWatch added a new service-linked role, called **AWSServiceRoleForNetworkMonitor**. CloudWatch added this new service-linked
role to allow you to create monitors to fetch network metrics between source subnets and
destination IP addresses. The new **CloudWatchNetworkMonitorServiceRolePolicy** IAM policy is attached to this
role, and the policy grants permission to CloudWatch to fetch network metrics on your
behalf.

December 22, 2023

[CloudWatch releases Amazon CloudWatch Network Monitor](what-is-network-monitor.md)

CloudWatch released a new feature, Amazon CloudWatch Network Monitor. This is a new active network
monitoring service that identifies if a network issues exists within the AWS network or
your own company network.

December 22, 2023

[**CloudWatchReadOnlyAccess** policy updated](auth-and-access-control-cw.md#managed-policies-cloudwatch-CloudWatchReadOnlyAccess)

CloudWatch added existing read-only permissions for CloudWatch Synthetics, X-Ray, and CloudWatch RUM
and new read-only permissions for CloudWatch Application Signals
to **CloudWatchReadOnlyAccess** so that users with this policy can
triage and diagnose service health issues as reported by CloudWatch Application Signals. The
`cloudwatch:GenerateQuery` permission was added so that users with this
policy can generate a CloudWatch Metrics Insights query string from a natural language prompt.

December 5, 2023

[**CloudWatchFullAccessV2** policy updated](auth-and-access-control-cw.md#managed-policies-cloudwatch-CloudWatchFullAccessV2)

CloudWatch added existing permissions to **CloudWatchFullAccessV2** for
CloudWatch Synthetics, X-Ray, and CloudWatch RUM, and added new permissions for CloudWatch Application
Signals so that users with this policy can fully manage Application Signals to triage and
diagnose issues with service health.

December 5, 2023

[New service-linked role and new IAM policy](using-service-linked-roles.md#service-linked-role-updates)

CloudWatch added a new service-linked role, called
**AWSServiceRoleForCloudWatchApplicationSignals**. CloudWatch added this new
service-linked role to allow CloudWatch Application Signals to collect CloudWatch Logs data, X-Ray trace
data, CloudWatch metrics data, and tagging data from applications that you have enabled for CloudWatch
Application Signals. The new
**CloudWatchApplicationSignalsServiceRolePolicy** IAM policy is
attached to this role, and the policy grants permission to CloudWatch Application Signals to
collect monitoring and tagging data from other relevant AWS services.

November 30, 2023

CloudWatch launches Preview release of Application Signals

CloudWatch Application Signals is in Preview. Use Application Signals to instrument your
applications on AWS so that you can monitor current application health, create service
level objectives (SLOs), and track long-term application performance against your business
objectives. For more information, see [Application Signals](cloudwatch-application-monitoring-sections.md).

November 30, 2023

CloudWatch adds support for querying other data sources

You can use CloudWatch to query, visualize, and create alarms for metrics from other data
sources For more information, see [Querying metrics\
from other data sources](multidatasourcequerying.md).

November 26, 2023

CloudWatch Metrics Insights supports natural language query generation

CloudWatch Metrics Insights supports natural language query to generate and update queries.
For more information, see [Use natural language to generate and update CloudWatch Metric Insights queries](cloudwatch-metrics-insights-query-assist.md).

November 26, 2023

[CloudWatch releases Container Insights with enhanced observability for Amazon EKS](containerinsights.md#container-insights-detailed-metrics)

CloudWatch released a new version of Container Insights. This version supports enhanced
observability for Amazon EKS clusters and can collect more detailed metrics from clusters
running Amazon EKS. After installation, it automatically collects detailed infrastructure
telemetry and container logs for your Amazon EKS clusters. You can then use curated,
immediately usable dashboards to drill down into application and infrastructure telemetry.

November 6, 2023

[CloudWatch metric streams adds quick partner setup](cloudwatch-metric-streams.md)

CloudWatch metric streams now provides a quick partner setup option, which you can use to
quickly set up a metric stream to some third-party providers.

October 17, 2023

[CloudWatch releases alarm recommendations](best-practice-alarms.md)

CloudWatch Synthetics now provides alarm recommendations for metrics from other AWS
services. These recommendations can help you identify the metrics that you should set
alarms for to follow best practices for monitoring these services.

October 16, 2023

[CloudWatch Synthetics releases runtime syn-nodejs-puppeteer-6.0](cloudwatch-synthetics-library-nodejs-puppeteer.md#CloudWatch_Synthetics_runtimeversion-nodejs-puppeteer-6.0)

CloudWatch Synthetics released runtime `syn-nodejs-puppeteer-6.0`.

September 26, 2023

[Adds Amazon CloudWatch Application Insights support for cross-account applications](appinsights-cross-account.md)

You can now share CloudWatch Application Insights applications across account boundaries.

September 26, 2023

[**New service-linked role and new IAM policy**](auth-and-access-control-cw.md#security-iam-awsmanpol-updates)

CloudWatch added a new service-linked role, called
**AWSServiceRoleForCloudWatchMetrics\_DbPerfInsights**. CloudWatch added this
new service-linked role to allow CloudWatch to fetch Performance Insights metrics for alarming,
anomaly detection, and snapshotting. The new
**AWSServiceRoleForCloudWatchMetrics\_DbPerfInsightsServiceRolePolicy**
IAM policy is attached to this role, and the policy grants permission to CloudWatch to fetch
Performance Insights metrics on your behalf.

September 20, 2023

[**Adds new metric math function**](using-metric-math.md#metric-math-syntax)

CloudWatch added a new metric math function, `DB_PERF_INSIGHTS`, that you can use
to fetch Performance Insights metrics from AWS database services for alarming, anomaly
detection, and snapshotting.

September 20, 2023

[**CloudWatchReadOnlyAccess** policy updated](auth-and-access-control-cw.md#security-iam-awsmanpol-updates)

CloudWatch added the `application-autoscaling:DescribeScalingPolicies` permission
to **CloudWatchReadOnlyAccess** so that users with this policy can access
information about Application Auto Scaling policies.

September 14, 2023

[CloudWatch agent added support for AL2023](install-cloudwatch-agent.md)

The CloudWatch agent supports AL2023.

August 8, 2023

[**New managed IAM policy, CloudWatchFullAccessV2**](auth-and-access-control-cw.md#security-iam-awsmanpol-updates)

CloudWatch added a new policy **CloudWatchFullAccessV2**. This policy
grants full access to CloudWatch actions and resources while better scoping the permissions
granted to other services such as Amazon SNS and Amazon EC2 Auto Scaling.

August 1, 2023

[Updated service linked role for Amazon CloudWatch Internet Monitor – update to an existing policy](auth-and-access-control-cw.md#security-iam-awsmanpol-updates)

Adds new permissions, `elasticloadbalancing:DescribeLoadBalancers` and
`ec2:DescribeNetworkInterfaces`, to the service linked role for Internet Monitor, to
support monitoring traffic for specific Network Load Balancer resources.

July 25, 2023

[Added support for Network Load Balancer resources in Amazon CloudWatch Internet Monitor](immonitorresources.md)

Adds support for creating a monitor in Internet Monitor with specific Network Load Balancer resources, to provide
more granular levels of observability for your application.

July 25, 2023

Dashboard variables feature

CloudWatch released _dashboard variables_, which you can use to create
flexible dashboards that can quickly display different contents depending on how you set
one input field within the dashboard. For example, you can create a dashboard that can
quickly switch between different Lambda functions or Amazon EC2 instance IDs, or one that can
switch to different AWS Regions. For more information, see [Create\
flexible dashboards with dashboard variables](cloudwatch-dashboard-variables.md).

June 28, 2023

Internet Monitor now supports customizing the threshold for health events

Internet Monitor added the ability to customize the threshold for when a global performance score
or availability score triggers a heath event. For more information, see [Tracking real-time\
performance and availability in Amazon CloudWatch Internet Monitor](cloudwatch-im-overview.md).

June 26, 2023

[**Internet Monitor now supports all commercial Regions**](cloudwatch-internetmonitor-regions.md)

Internet Monitor added seven new AWS Regions and now supports all commercial Regions.

June 19, 2023

New Lambda Insights extension versions

CloudWatch added the 1.0.229.0 version of the Lambda Insights extension for both x86-64
platforms and ARM64 platforms. For more information, see [Available versions of the Lambda Insights extension](lambda-insights-extension-versions.md).

June 12, 2023

CloudWatchReadOnlyAccess policy updated

CloudWatch added permissions to **CloudWatchReadOnlyAccess**. The
`logs:StartLiveTail` and `logs:StopLiveTail` permissions were
added so that users with this policy can use the console to start and stop CloudWatch Logs live tail
sessions. For more information, see [Use live tail to view\
logs in near real time](../logs/cloudwatchlogs-livetail.md).

June 6, 2023

CloudWatch RUM adds support for custom metrics

You can use CloudWatch RUM app monitors to create custom metrics and send them to CloudWatch and
CloudWatch Evidently. This feature includes an update to the
**AmazonCloudWatchRUMServiceRolePolicy** managed IAM policy. In that
policy, a condition key was changed so that CloudWatch RUM can send custom metrics to custom
metric namespaces.

February 9, 2023

New and updated managed polices for CloudWatch

To support CloudWatch cross-account observability, the `CloudWatchFullAccess` and
`CloudWatchReadOnlyAccess` policies have been updated, and the following new
managed policies have been added: `CloudWatchCrossAccountSharingConfiguration`,
`OAMFullAccess`, and `OAMReadOnlyAccess`. For more information,
see [CloudWatch updates to AWS managed policies](iam-identity-based-access-control-cw.md#iam-awsmanpol-updates).

February 7, 2023

[CloudWatch Application Insights service-linked role policy updates — update to an existing policy.](chap-using-service-linked-roles-appinsights.md)

CloudWatch Application Insights updated an existing AWS service-linked role policy.

December 19, 2022

[Amazon CloudWatch Application Insights support for containerized applications and microservices from the Container Insights console.](cloudwatch-application-insights.md)

You can display CloudWatch Application Insights detected problems for Amazon ECS and Amazon EKS on your Container
Insights dashboard.

November 17, 2021

[Amazon CloudWatch Application Insights monitoring for SAP HANA databases.](cloudwatch-application-insights.md)

You can monitor SAP HANA databases with Application Insights.

November 15, 2021

[Amazon CloudWatch Application Insights support for monitoring all resources in an account.](cloudwatch-application-insights.md)

You can onboard and monitor all resources in an account.

September 15, 2021

[Amazon CloudWatch Application Insights support for Amazon FSx.](cloudwatch-application-insights.md)

You can monitor metrics retrieved from Amazon FSx.

August 31, 2021

SDK Metrics is no longer supported.

CloudWatch SDK Metrics is no longer supported.

August 25, 2021

[Amazon CloudWatch Application Insights support for setting up container monitoring.](cloudwatch-application-insights.md)

You can monitor containers using best practices with Amazon CloudWatch Application Insights.

May 18, 2021

Metric streams is generally available

You can use metric streams to continually stream CloudWatch metrics to a destination of your
choice. For more information, see [Metric streams](cloudwatch-metric-streams.md) in the _Amazon CloudWatch User Guide_.

March 31, 2021

[Amazon CloudWatch Application Insights monitoring for Oracle databases on Amazon RDS and Amazon EC2.](cloudwatch-application-insights.md)

You can monitor metrics and logs retrieved from Oracle with Amazon CloudWatch Application Insights.

January 16, 2021

Lambda Insights is generally available

CloudWatch Lambda Insights is a monitoring and troubleshooting solution for serverless
applications running on AWS Lambda. For more information, see [Using Lambda Insights](lambda-insights.md) in the
_Amazon CloudWatch User Guide_.

December 3, 2020

[Amazon CloudWatch Application Insights monitoring for Prometheus JMX exporter metrics.](cloudwatch-application-insights.md)

You can monitor metrics retrieved from Prometheus JMX exporter with
Amazon CloudWatch Application Insights.

November 20, 2020

CloudWatch Synthetics releases new runtime version

CloudWatch Synthetics has released a new runtime version. For more information, see [Canary Runtime\
Versions](cloudwatch-synthetics-canaries-library.md) in the _Amazon CloudWatch User Guide_.

September 11, 2020

[Amazon CloudWatch Application Insights monitoring for Postgre SQL on Amazon RDS and Amazon EC2.](cloudwatch-application-insights.md)

You can monitor applications built with PostgreSQL running on Amazon RDS or
Amazon EC2.

September 11, 2020

CloudWatch supports dashboard sharing

You can now share CloudWatch dashboards with people outside of your organization and AWS
account. For more information, see [Sharing CloudWatch Dashboards](cloudwatch-dashboard-sharing.md) in
the _Amazon CloudWatch User Guide_.

September 10, 2020

[Set up monitors for .NET applications using SQL Server on the backend with CloudWatch Application Insights](cloudwatch-application-insights.md)

You can use the documentation tutorial to help you to set up monitors for .NET
applications using SQL Server on the backend with CloudWatch Application Insights.

August 19, 2020

[CloudFormation support for Amazon CloudWatch Application Insights applications.](cloudwatch-application-insights.md)

You can add CloudWatch Application Insights monitoring, including key metrics and telemetry, to your
application, database, and web server, directly from CloudFormation templates.

July 30, 2020

[Amazon CloudWatch Application Insights monitoring for Aurora for MySQL database clusters.](cloudwatch-application-insights.md)

You can monitor Aurora for MySQL database clusters (RDS Aurora) with
Amazon CloudWatch Application Insights.

July 2, 2020

CloudWatch Contributor Insights general availability

CloudWatch Contributor Insights is now generally available. It enables you to analyze log
data and create time series that display contributor data. You can see metrics about the
top-N contributors, the total number of unique contributors, and their usage. For more
information, see [Using Contributor\
Insights to Analyze High-Cardinality Data](contributorinsights.md) in the
_Amazon CloudWatch User Guide_.

April 2, 2020

CloudWatch Synthetics public preview

CloudWatch Synthetics is now in public preview. It enables you to create canaries to monitor
your endpoints and APIs. For more information, see [Using Canaries](cloudwatch-synthetics-canaries.md) in the
_Amazon CloudWatch User Guide_.

November 25, 2019

CloudWatch Contributor Insights public preview

CloudWatch Contributor Insights is now in public preview. It enables you to analyze log data
and create time series that display contributor data. You can see metrics about the top-N
contributors, the total number of unique contributors, and their usage. For more
information, see [Using Contributor\
Insights to Analyze High-Cardinality Data](contributorinsights.md) in the
_Amazon CloudWatch User Guide_.

November 25, 2019

CloudWatch launches ServiceLens feature

ServiceLens ehances the observability of your services and applications by enabling
you to integrate traces, metrics, logs, and alarms into one place. ServiceLens integrates
CloudWatch with AWS X-Ray to provide an end-to-end view of your application.

November 21, 2019

Use CloudWatch to proactively manage your AWS service quotas

You can use CloudWatch to proactively manage your AWS service quotas. CloudWatch usage metrics
provide visibility into your account's usage of resources and API operations. For more
information, see [Service Quotas Integration and Usage Metrics](cloudwatch-service-quota-integration.md) in the
_Amazon CloudWatch User Guide_.

November 19, 2019

CloudWatch sends events when alarms change state

CloudWatch now sends an event to Amazon EventBridge when any CloudWatch alarm changes state. For more
information, see [Alarm Events\
and EventBridge](cloudwatch-and-eventbridge.md) in the _Amazon CloudWatch User Guide_.

October 8, 2019

Container Insights

CloudWatch Container Insights is now generally available. It enables you to collect,
aggregate, and summarize metrics and logs from your containerized applications and
microservices. For more information, see [Using Container Insights](containerinsights.md) in the _Amazon CloudWatch User Guide_.

August 30, 2019

Updates for Container Insights preview metrics on Amazon EKS and Kubernetes

The Container Insights on Amazon EKS and Kubernetes public preview has been updated.
InstanceId is now included as a dimension to the cluster EC2 instances. This allows alarms
that have been created on these metrics to trigger the following EC2 actions: Stop,
Terminate, Reboot, or Recover. Additionally, pod and service metrics are now reported by
Kubernetes namespace to simplify the monitoring and alarming on metrics by
namespace.

August 19, 2019

Updates for AWS Systems Manager OpsCenter integration

Updates on how CloudWatch Application Insights integrates with Systems Manager OpsCenter.

August 7, 2019

CloudWatch usage metrics

CloudWatch usage metrics help you track the usage of your CloudWatch resources and stay within
your service limits. For more information, see [https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Usage-Metrics.html](cloudwatch-usage-metrics.md).

August 6, 2019

CloudWatch Container Insights public preview

CloudWatch Container Insights is now in public preview. It enables you to collect,
aggregate, and summarize metrics and logs from your containerized applications and
microservices. For more information, see [Using Container Insights](containerinsights.md) in the _Amazon CloudWatch User Guide_.

July 9, 2019

CloudWatch Anomaly Detection public preview

CloudWatch anomaly detection is now in public preview. CloudWatch applies machine-learning
algorithms to a metric's past data to create a model of the metric's expected values. You
can use this model for visualization and for setting alarms. For more information, see
[Using CloudWatch Anomaly\
Detection](cloudwatch-anomaly-detection.md) in the _Amazon CloudWatch User Guide_.

July 9, 2019

CloudWatch Application Insights for .NET and SQL Server

CloudWatch Application Insights for .NET and SQL Server facilitates observability for .NET
and SQL Server applications. It can help you set up the best monitors for your application
resources to continuously analyze data for signs of problems with your
applications.

June 21, 2019

CloudWatch agent section re-organized

The CloudWatch agent documentation has been rewritten to improve clarity, especially for
customers using the command line to install and configure the agent. For more information,
see [Collecting Metrics and Logs\
from Amazon EC2 Instances and On-Premises Servers with the CloudWatch Agent](install-cloudwatch-agent.md) in the
_Amazon CloudWatch User Guide_.

March 28, 2019

SEARCH function added to metric math expressions

You can now use a SEARCH function in metric math expressions. This enables you to
create dashboards that update automatically as new resources are created that match the
search query. For more information, see [Using Search Expressions in\
Graphs](using-search-expressions.md) in the _Amazon CloudWatch User Guide_.

March 21, 2019

AWS SDK Metrics for Enterprise Support

SDK Metrics helps you assess the health of your AWS services and diagnose latency caused
by reaching your account usage limits or by a service outage. For more information, see
[Monitor Applications Using\
AWS SDK Metrics](../../../../reference/amazoncloudwatch/latest/monitoring/cloudwatch-agent-sdk-metrics.md) in the _Amazon CloudWatch User Guide_.

December 11, 2018

Alarms on math expressions

CloudWatch supports creating alarms based on metric math expressions. For more information,
see [Alarms on Math Expressions](cloudwatch-alarms.md#alarms-on-metric-math-expressions) in the _Amazon CloudWatch User Guide_.

November 20, 2018

New CloudWatch console home page

Amazon has created a new home page in the CloudWatch console, which automatically displays
key metrics and alarms for all the AWS services you are using. For more information, see
[Getting Started with Amazon CloudWatch](gettingstarted.md) in
the _Amazon CloudWatch User Guide_.

November 19, 2018

CloudFormation templates for the CloudWatch Agent

Amazon has uploaded CloudFormation templates that you can use to install and update the CloudWatch
agent. For more information, see [Install\
the CloudWatch Agent on New Instances Using CloudFormation](install-cloudwatch-agent-new-instances-cloudformation.md) in the
_Amazon CloudWatch User Guide_.

November 9, 2018

Enhancements to the CloudWatch Agent

The CloudWatch agent has been updated to work with both the StatsD and collectd protocols.
It also has improved cross-account support. For more information, see [Retrieve Custom Metrics\
with StatsD](cloudwatch-agent-custom-metrics-statsd.md), [Retrieve Custom Metrics\
with collectd](cloudwatch-agent-custom-metrics-collectd.md), and [Sending Metrics and Logs to a Different AWS Account](cloudwatch-agent-common-scenarios.md#CloudWatch-Agent-send-to-different-AWS-account) in the
_Amazon CloudWatch User Guide_.

September 28, 2018

Support for Amazon VPC endpoints

You can now establish a private connection between your VPC and CloudWatch. For more
information, see [Using CloudWatch\
with Interface VPC Endpoints](cloudwatch-and-interface-vpc.md) in the _Amazon CloudWatch User Guide_.

June 28, 2018

The following table describes important changes to the _Amazon CloudWatch User_
_Guide_ before June 2018.

ChangeDescriptionRelease date

Metric math

You can now perform math expressions on CloudWatch metrics, producing new time series
that you can add to graphs on your dashboard. For more information, see [Using math expressions with CloudWatch metrics](using-metric-math.md).

April 4, 2018

"M out of N" alarms

You can now configure an alarm to trigger based on "M out of N" datapoints in any
alarm evaluation interval. For more information, see [Alarm evaluation](alarm-evaluation.md).

December 8, 2017

CloudWatch agent

A new unified CloudWatch agent was released. You can use the unified multi-platform
agent to collect custom both system metrics and log files from Amazon EC2 instances and
on-premises servers. The new agent supports both Windows and Linux and enables
customization of metrics collected, including sub-resource metrics such as per-CPU
core. For more information, see [Collect metrics, logs, and traces using the CloudWatch agent](install-cloudwatch-agent.md).

September 7, 2017

NAT gateway metrics

Added metrics for Amazon VPC NAT gateway.

September 7, 2017

High-resolution metrics

You can now optionally set up custom metrics as high-resolution metrics, with a
granularity of as low as one second. For more information, see [High-resolution metrics](publishingmetrics.md#high-resolution-metrics).

July 26, 2017

Dashboard APIs

You can now create, modify, and delete dashboards using APIs and the AWS CLI. For
more information, see [Creating a customized CloudWatch dashboard](create-dashboard.md).

July 6, 2017

Direct Connect metrics

Added metrics for Direct Connect.

June 29, 2017

Amazon VPC VPN metrics

Added metrics for Amazon VPC VPN.

May 15, 2017

WorkSpaces Applications metrics

Added metrics for WorkSpaces Applications.

March 8, 2017

CloudWatch console color picker

You can now choose the color for each metric on your dashboard widgets. For more
information, see [Editing a graph on a CloudWatch dashboard](edit-graph-dashboard.md).

February 27, 2017

Alarms on dashboards

Alarms can now be added to dashboards. For more information, see [Adding an alarm to a CloudWatch dashboard](add-alarm-dashboard.md).

February 15, 2017

Added metrics for Amazon Polly

Added metrics for Amazon Polly.

December 1, 2016

Added metrics for Amazon Managed Service for Apache Flink

Added metrics for Amazon Managed Service for Apache Flink.

December 1, 2016

Added support for percentile statistics

You can specify any percentile, using up to two decimal places (for example,
p95.45). For more information, see [Percentiles](cloudwatch-concepts.md#Percentiles).

November 17, 2016

Added metrics for Amazon Simple Email Service

Added metrics for Amazon Simple Email Service.

November 2, 2016

Updated metrics retention

Amazon CloudWatch now retains metrics data for 15 months instead of 14 days.

November 1, 2016

Updated metrics console interface

The CloudWatch console is updated with improvements to existing functionality and new
functionality.

November 1, 2016

Added metrics for Amazon Elastic Transcoder

Added metrics for Amazon Elastic Transcoder.

September 20, 2016

Added metrics for Amazon API Gateway

Added metrics for Amazon API Gateway.

September 9, 2016

Added metrics for AWS Key Management Service

Added metrics for AWS Key Management Service.

September 9, 2016

Added metrics for the new Application Load Balancers supported by Elastic Load Balancing

Added metrics for Application Load Balancers.

August 11, 2016

Added new NetworkPacketsIn and NetworkPacketsOut metrics for Amazon EC2

Added new NetworkPacketsIn and NetworkPacketsOut metrics for Amazon EC2.

March 23, 2016

Added new metrics for Amazon EC2 Spot fleet

Added new metrics for Amazon EC2 Spot fleet.

March 21, 2016

Added new CloudWatch Logs metrics

Added new CloudWatch Logs metrics.

March 10, 2016

Added Amazon OpenSearch Service and AWS WAF metrics and dimensions

Added Amazon OpenSearch Service and AWS WAF metrics and dimensions.

October 14, 2015

Added support for CloudWatch dashboards

Dashboards are customizable home pages in the CloudWatch console that you can use to
monitor your resources in a single view, even those that are spread out across
different Regions. For more information, see [Using Amazon CloudWatch dashboards](cloudwatch-dashboards.md).

October 8, 2015

Added AWS Lambda metrics and dimensions

Added AWS Lambda metrics and dimensions.

September 4, 2015

Added Amazon Elastic Container Service metrics and dimensions

Added Amazon Elastic Container Service metrics and dimensions.

August 17, 2015

Added Amazon Simple Storage Service metrics and dimensions

Added Amazon Simple Storage Service metrics and dimensions.

July 26, 2015

New feature: Reboot alarm action

Added the reboot alarm action and new IAM role for use with alarm actions. For
more information, see [Stop, terminate, reboot, or recover an EC2 instance](usingalarmactions.md).

July 23, 2015

Added Amazon WorkSpaces metrics and dimensions

Added Amazon WorkSpaces metrics and dimensions.

April 30, 2015

Added Amazon Machine Learning metrics and dimensions

Added Amazon Machine Learning metrics and dimensions.

April 9, 2015

New feature: Amazon EC2 instance recovery alarm actions

Updated alarm actions to include new EC2 instance recovery action. For more
information, see [Stop, terminate, reboot, or recover an EC2 instance](usingalarmactions.md).

March 12, 2015

Added Amazon CloudFront and Amazon CloudSearch metrics and dimensions

Added Amazon CloudFront and Amazon CloudSearch metrics and dimensions.

March 6, 2015

Added Amazon Simple Workflow Service metrics and dimensions

Added Amazon Simple Workflow Service metrics and dimensions.

May 9, 2014

Updated guide to add support for AWS CloudTrail

Added a new topic to explain how you can use AWS CloudTrail to log activity in Amazon CloudWatch.
For more information, see [Logging Amazon CloudWatch API and console operations with AWS CloudTrail](logging-cw-api-calls.md).

April 30, 2014

Updated guide to use the new AWS Command Line Interface (AWS CLI)

The AWS CLI is a cross-service CLI with a simplified installation, unified
configuration, and consistent command line syntax. The AWS CLI is supported on
Linux/Unix, Windows, and Mac. The CLI examples in this guide have been updated to use
the new AWS CLI.

For information about how to install and configure the new AWS CLI, see [Getting Set Up with the AWS CLI\
Interface](../../../cli/latest/userguide/cli-chap-getting-set-up.md) in the _AWS Command Line Interface User Guide_.

February 21, 2014

Added Amazon Redshift and AWS OpsWorks metrics and dimensions

Added Amazon Redshift and AWS OpsWorks metrics and dimensions.

July 16, 2013

Added Amazon Route 53 metrics and dimensions

Added Amazon Route 53 metrics and dimensions.

June 26, 2013

New feature: Amazon CloudWatch Alarm Actions

Added a new section to document Amazon CloudWatch alarm actions, which you can use to stop
or terminate an Amazon Elastic Compute Cloud instance. For more information, see [Stop, terminate, reboot, or recover an EC2 instance](usingalarmactions.md).

January 8, 2013

Updated EBS metrics

Updated the EBS metrics to include two new metrics for Provisioned IOPS volumes.

November 20, 2012

New billing alerts

You can now monitor your AWS charges using Amazon CloudWatch metrics and create alarms to
notify you when you have exceeded the specified threshold. For more information, see
[Create a billing alarm to monitor your estimated AWS charges](monitor-estimated-charges-with-cloudwatch.md).

May 10, 2012

New metrics

You can now access six new Elastic Load Balancing metrics that provide counts of various HTTP
response codes.

October 19, 2011

New feature

You can now access metrics from Amazon EMR.

June 30, 2011

New feature

You can now access metrics from Amazon Simple Notification Service and Amazon Simple Queue Service.

July 14, 2011

New Feature

Added information about using the `PutMetricData` API to publish custom
metrics. For more information, see [Publish custom metrics](publishingmetrics.md).

May 10, 2011

Updated metrics retention

Amazon CloudWatch now retains the history of an alarm for two weeks rather than six weeks.
With this change, the retention period for alarms matches the retention period for
metrics data.

April 7, 2011

New feature

Added ability to send Amazon Simple Notification Service or Auto Scaling notifications when a metric has crossed a
threshold. For more information, see [Alarms](cloudwatch-concepts.md#CloudWatchAlarms).

December 2, 2010

New feature

A number of CloudWatch actions now include the MaxRecords and NextToken parameters,
which enable you to control pages of results to display.

December 2, 2010

New feature

This service now integrates with AWS Identity and Access Management (IAM).

December 2, 2010

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Service quotas

All content copied from https://docs.aws.amazon.com/.
