This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationTelemetryRule TelemetryDestinationConfiguration

Configuration specifying where and how telemetry data should be delivered for AWS
resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudtrailParameters" : CloudtrailParameters,
  "DestinationPattern" : String,
  "DestinationType" : String,
  "ELBLoadBalancerLoggingParameters" : ELBLoadBalancerLoggingParameters,
  "RetentionInDays" : Integer,
  "VPCFlowLogParameters" : VPCFlowLogParameters,
  "WAFLoggingParameters" : WAFLoggingParameters
}

```

### YAML

```yaml

  CloudtrailParameters:
    CloudtrailParameters
  DestinationPattern: String
  DestinationType: String
  ELBLoadBalancerLoggingParameters:
    ELBLoadBalancerLoggingParameters
  RetentionInDays: Integer
  VPCFlowLogParameters:
    VPCFlowLogParameters
  WAFLoggingParameters:
    WAFLoggingParameters

```

## Properties

`CloudtrailParameters`

Configuration parameters specific to AWS CloudTrail when CloudTrail is the source type.

_Required_: No

_Type_: [CloudtrailParameters](aws-properties-observabilityadmin-organizationtelemetryrule-cloudtrailparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationPattern`

The pattern used to generate the destination path or name, supporting macros like
<resourceId> and <accountId>.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationType`

The type of destination for the telemetry data (e.g., "Amazon CloudWatch Logs", "S3").

_Required_: No

_Type_: String

_Allowed values_: `cloud-watch-logs`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ELBLoadBalancerLoggingParameters`

Configuration parameters specific to ELB load balancer logging when ELB is the resource
type.

_Required_: No

_Type_: [ELBLoadBalancerLoggingParameters](aws-properties-observabilityadmin-organizationtelemetryrule-elbloadbalancerloggingparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionInDays`

The number of days to retain the telemetry data in the destination.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VPCFlowLogParameters`

Configuration parameters specific to VPC Flow Logs when VPC is the resource type.

_Required_: No

_Type_: [VPCFlowLogParameters](aws-properties-observabilityadmin-organizationtelemetryrule-vpcflowlogparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WAFLoggingParameters`

Configuration parameters specific to WAF logging when WAF is the resource type.

_Required_: No

_Type_: [WAFLoggingParameters](aws-properties-observabilityadmin-organizationtelemetryrule-wafloggingparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TelemetryRule

All content copied from https://docs.aws.amazon.com/.
