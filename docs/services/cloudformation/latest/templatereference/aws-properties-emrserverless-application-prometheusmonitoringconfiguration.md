This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application PrometheusMonitoringConfiguration

The monitoring configuration object you can configure to send metrics to Amazon Managed Service for Prometheus for a job run.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RemoteWriteUrl" : String
}

```

### YAML

```yaml

  RemoteWriteUrl: String

```

## Properties

`RemoteWriteUrl`

The remote write URL in the Amazon Managed Service for Prometheus workspace to send metrics to.

_Required_: No

_Type_: String

_Pattern_: `^https://aps-workspaces.([a-z]{2}-[a-z-]{1,20}-[1-9]).amazonaws(.[0-9A-Za-z]{2,4})+/workspaces/[-_.0-9A-Za-z]{1,100}/api/v1/remote_write$`

_Minimum_: `1`

_Maximum_: `10280`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkConfiguration

S3MonitoringConfiguration

All content copied from https://docs.aws.amazon.com/.
