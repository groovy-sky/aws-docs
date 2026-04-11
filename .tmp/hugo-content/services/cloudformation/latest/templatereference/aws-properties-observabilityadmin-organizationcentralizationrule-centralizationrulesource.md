This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationCentralizationRule CentralizationRuleSource

Configuration specifying the source of telemetry data to be centralized.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Regions" : [ String, ... ],
  "Scope" : String,
  "SourceLogsConfiguration" : SourceLogsConfiguration
}

```

### YAML

```yaml

  Regions:
    - String
  Scope: String
  SourceLogsConfiguration:
    SourceLogsConfiguration

```

## Properties

`Regions`

The list of source regions from which telemetry data should be centralized.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

The organizational scope from which telemetry data should be centralized, specified using
organization id, accounts or organizational unit ids.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceLogsConfiguration`

Log specific configuration for centralization source log groups.

_Required_: No

_Type_: [SourceLogsConfiguration](aws-properties-observabilityadmin-organizationcentralizationrule-sourcelogsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CentralizationRuleDestination

DestinationLogsConfiguration

All content copied from https://docs.aws.amazon.com/.
