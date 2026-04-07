This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationTelemetryRule TelemetryRule

Defines how telemetry should be configured for specific AWS resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationConfiguration" : TelemetryDestinationConfiguration,
  "ResourceType" : String,
  "Scope" : String,
  "SelectionCriteria" : String,
  "TelemetrySourceTypes" : [ String, ... ],
  "TelemetryType" : String
}

```

### YAML

```yaml

  DestinationConfiguration:
    TelemetryDestinationConfiguration
  ResourceType: String
  Scope: String
  SelectionCriteria: String
  TelemetrySourceTypes:
    - String
  TelemetryType: String

```

## Properties

`DestinationConfiguration`

Configuration specifying where and how the telemetry data should be delivered.

_Required_: No

_Type_: [TelemetryDestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-organizationtelemetryrule-telemetrydestinationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

The type of AWS resource to configure telemetry for (e.g., "AWS::EC2::VPC",
"AWS::EKS::Cluster", "AWS::WAFv2::WebACL").

_Required_: Yes

_Type_: String

_Allowed values_: `AWS::EC2::VPC | AWS::WAFv2::WebACL | AWS::CloudTrail | AWS::EKS::Cluster | AWS::ElasticLoadBalancingV2::LoadBalancer`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

The organizational scope to which the rule applies, specified using accounts or
organizational units.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectionCriteria`

Criteria for selecting which resources the rule applies to, such as resource tags.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TelemetrySourceTypes`

The specific telemetry source types to configure for the resource, such as VPC\_FLOW\_LOGS
or EKS\_AUDIT\_LOGS. TelemetrySourceTypes must be correlated with the specific resource type.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TelemetryType`

The type of telemetry to collect (Logs, Metrics, or Traces).

_Required_: Yes

_Type_: String

_Allowed values_: `Logs`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TelemetryDestinationConfiguration

VPCFlowLogParameters
