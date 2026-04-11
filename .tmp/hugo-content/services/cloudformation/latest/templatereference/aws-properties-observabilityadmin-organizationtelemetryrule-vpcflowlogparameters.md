This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationTelemetryRule VPCFlowLogParameters

Configuration parameters specific to VPC Flow Logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogFormat" : String,
  "MaxAggregationInterval" : Integer,
  "TrafficType" : String
}

```

### YAML

```yaml

  LogFormat: String
  MaxAggregationInterval: Integer
  TrafficType: String

```

## Properties

`LogFormat`

The format in which VPC Flow Log entries should be logged.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxAggregationInterval`

The maximum interval in seconds between the capture of flow log records.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrafficType`

The type of traffic to log (ACCEPT, REJECT, or ALL).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TelemetryRule

WAFLoggingParameters

All content copied from https://docs.aws.amazon.com/.
