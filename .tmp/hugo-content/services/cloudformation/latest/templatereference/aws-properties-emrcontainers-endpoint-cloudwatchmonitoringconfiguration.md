This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::Endpoint CloudWatchMonitoringConfiguration

A configuration for CloudWatch monitoring. You can configure your jobs to send log
information to CloudWatch Logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroupName" : String,
  "LogStreamNamePrefix" : String
}

```

### YAML

```yaml

  LogGroupName: String
  LogStreamNamePrefix: String

```

## Properties

`LogGroupName`

The name of the log group for log publishing.

_Required_: Yes

_Type_: String

_Pattern_: `[\.\-_/#A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogStreamNamePrefix`

The specified name prefix for log streams.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Certificate

ConfigurationOverrides

All content copied from https://docs.aws.amazon.com/.
