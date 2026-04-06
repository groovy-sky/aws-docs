# Collect metrics, logs, and traces using the CloudWatch agent

CloudWatch agent is a software component that collects metrics, logs, and traces from your Amazon EC2
instances, on-premises servers, and containerized applications. It enables you to monitor your
infrastructure and applications more comprehensively than the basic monitoring provided by
default.

**Key benefits**

- Collect system-level metrics (CPU, memory, disk, network)

- Gather custom metrics from your applications

- Collect and centralize logs from various sources

- Monitor both AWS and on-premises environments with a single tool

- Set up alarms and notifications based on collected data

The CloudWatch agent enables you to do the following:

- Collect internal system-level metrics from Amazon EC2 instances across operating systems. The
metrics can include in-guest metrics, in addition to the metrics for EC2 instances. The
additional metrics that can be collected are listed in [Metrics collected by the CloudWatch agent](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/metrics-collected-by-CloudWatch-agent.html).

- Collect system-level metrics from on-premises servers. These can include servers in a
hybrid environment as well as servers not managed by AWS.

- Retrieve custom metrics from your applications or services using the `StatsD`
and `collectd` protocols. `StatsD` is supported on both Linux servers
and servers running Windows Server. `collectd` is supported only on Linux
servers.

- Collect logs from Amazon EC2 instances and on-premises servers, running either Linux or
Windows Server.

###### Note

The CloudWatch agent does not support collecting logs from FIFO pipes.

- Send the metrics to either CloudWatch or Amazon Managed Service for Prometheus, or to both. The CloudWatch agent configuration
file contains a `metrics_destinations` parameter in the `metrics`
section. You can specify `cloudwatch`, `amp`, or both in this
parameter.

- Version 1.300031.0 and later can be used to enable CloudWatch Application Signals. For more
information, see [Application Signals](cloudwatch-application-monitoring-sections.md).

- Version 1.300025.0 and later can collect traces from [OpenTelemetry](https://docs.aws.amazon.com/xray/latest/devguide/xray-instrumenting-your-app.html#xray-instrumenting-opentel) or [X-Ray](https://docs.aws.amazon.com/xray/latest/devguide/xray-instrumenting-your-app.html#xray-instrumenting-xray-sdk) client SDKs, and send them to X-Ray.

Using the CloudWatch agent allows you to collect traces without needing to run a separate
trace collection daemon, helping to reduce the number of agents that you run and
manage.

Metrics sent to CloudWatch can be viewed in CloudWatch just as any other CloudWatch metrics. The default CloudWatch
namespace for metrics collected by the CloudWatch agent is `CWAgent`, although you can
specify a different namespace when you configure the agent.

The logs collected by the CloudWatch agent are processed and stored in Amazon CloudWatch Logs, just like logs
collected by the older CloudWatch Logs agent. For information about CloudWatch Logs pricing, see [Amazon CloudWatch Pricing](http://aws.amazon.com/cloudwatch/pricing).

Metrics collected by the CloudWatch agent are billed as custom metrics. For more information about
CloudWatch metrics pricing, see [Amazon CloudWatch\
Pricing](http://aws.amazon.com/cloudwatch/pricing).

The CloudWatch agent is open-source under the MIT license, and is [hosted on GitHub](https://github.com/aws/amazon-cloudwatch-agent). If you would
like to build, customize or contribute to the CloudWatch agent, see the GitHub repository for the
latest instructions. If you think you’ve found a potential security issue, do not post it on
GitHub or any public forum. Instead, please follow the instructions at [Vulnerability Reporting](https://aws.amazon.com/security/vulnerability-reporting) or [email AWS security directly](mailto:aws-security@amazon.com).

You can download and install the CloudWatch agent manually using the command line, or you can
integrate it with AWS Systems Manager. The general flow of installing the CloudWatch agent is as
follows:

1. Create IAM roles or users that enable the agent to collect metrics from the server and
    optionally to integrate with AWS Systems Manager.

2. Download the agent package.

3. Modify the CloudWatch agent configuration file and specify the metrics that you want to
    collect.

4. Install and start the agent on your servers.

###### Contents

- [Supported operating systems](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/supported-operating-systems.html)

- [Prerequisites](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/prerequisites.html)

- [Download the CloudWatch agent package](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/download-CloudWatch-Agent-on-EC2-Instance-commandline-first.html)

- [Verifying the signature of the CloudWatch agent package](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/verify-CloudWatch-Agent-Package-Signature.html)

- [Installing the CloudWatch agent](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/install-CloudWatch-Agent-on-EC2-Instance.html)

- [Set up the CloudWatch agent with security-enhanced Linux (SELinux)](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Agent-SELinux.html)

- [Create the CloudWatch agent configuration file](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/create-cloudwatch-agent-configuration-file.html)

- [Starting the CloudWatch agent](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/start-CloudWatch-Agent-on-premise-SSM-onprem.html)

- [Metrics collected by the CloudWatch agent](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/metrics-collected-by-CloudWatch-agent.html)

- [Using the CloudWatch agent with related telemetry](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Agent-RelatedEntities.html)

- [Common scenarios with the CloudWatch agent](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Agent-common-scenarios.html)

- [CloudWatch agent credentials preference](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Agent-Credentials-Preference.html)

- [Troubleshooting the CloudWatch agent](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/troubleshooting-CloudWatch-Agent.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Delete a connector to a data source

Supported operating systems
