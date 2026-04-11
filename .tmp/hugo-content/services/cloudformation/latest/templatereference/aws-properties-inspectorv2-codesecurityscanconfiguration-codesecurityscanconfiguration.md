This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CodeSecurityScanConfiguration CodeSecurityScanConfiguration

Contains the configuration settings for code security scans.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "continuousIntegrationScanConfiguration" : ContinuousIntegrationScanConfiguration,
  "periodicScanConfiguration" : PeriodicScanConfiguration,
  "ruleSetCategories" : [ String, ... ]
}

```

### YAML

```yaml

  continuousIntegrationScanConfiguration:
    ContinuousIntegrationScanConfiguration
  periodicScanConfiguration:
    PeriodicScanConfiguration
  ruleSetCategories:
    - String

```

## Properties

`continuousIntegrationScanConfiguration`

Configuration settings for continuous integration scans that run automatically when
code changes are made.

_Required_: No

_Type_: [ContinuousIntegrationScanConfiguration](aws-properties-inspectorv2-codesecurityscanconfiguration-continuousintegrationscanconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`periodicScanConfiguration`

Configuration settings for periodic scans that run on a scheduled basis.

_Required_: No

_Type_: [PeriodicScanConfiguration](aws-properties-inspectorv2-codesecurityscanconfiguration-periodicscanconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ruleSetCategories`

The categories of security rules to be applied during the scan.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::InspectorV2::CodeSecurityScanConfiguration

ContinuousIntegrationScanConfiguration

All content copied from https://docs.aws.amazon.com/.
