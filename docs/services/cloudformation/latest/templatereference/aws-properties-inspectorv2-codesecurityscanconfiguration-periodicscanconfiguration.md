This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CodeSecurityScanConfiguration PeriodicScanConfiguration

Configuration settings for periodic scans that run on a scheduled basis.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "frequency" : String,
  "frequencyExpression" : String
}

```

### YAML

```yaml

  frequency: String
  frequencyExpression: String

```

## Properties

`frequency`

The frequency at which periodic scans are performed (such as weekly or monthly).

If you don't provide the `frequencyExpression` Amazon Inspector chooses day for the scan
to run. If you provide the `frequencyExpression`, the schedule must match the
specified `frequency`.

_Required_: No

_Type_: String

_Allowed values_: `WEEKLY | MONTHLY | NEVER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`frequencyExpression`

The schedule expression for periodic scans, in cron format.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContinuousIntegrationScanConfiguration

ScopeSettings

All content copied from https://docs.aws.amazon.com/.
