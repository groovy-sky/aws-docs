This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationCentralizationRule CentralizationRuleDestination

Configuration specifying the primary destination for centralized telemetry data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Account" : String,
  "DestinationLogsConfiguration" : DestinationLogsConfiguration,
  "Region" : String
}

```

### YAML

```yaml

  Account: String
  DestinationLogsConfiguration:
    DestinationLogsConfiguration
  Region: String

```

## Properties

`Account`

The destination account (within the organization) to which the telemetry data should be
centralized.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationLogsConfiguration`

Log specific configuration for centralization destination log groups.

_Required_: No

_Type_: [DestinationLogsConfiguration](aws-properties-observabilityadmin-organizationcentralizationrule-destinationlogsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The primary destination region to which telemetry data should be centralized.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CentralizationRule

CentralizationRuleSource

All content copied from https://docs.aws.amazon.com/.
