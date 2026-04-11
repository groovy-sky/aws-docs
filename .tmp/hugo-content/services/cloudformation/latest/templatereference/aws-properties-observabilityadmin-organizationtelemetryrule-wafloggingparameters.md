This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationTelemetryRule WAFLoggingParameters

Configuration parameters for WAF logging, including redacted fields and logging filters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LoggingFilter" : LoggingFilter,
  "LogType" : String,
  "RedactedFields" : [ FieldToMatch, ... ]
}

```

### YAML

```yaml

  LoggingFilter:
    LoggingFilter
  LogType: String
  RedactedFields:
    - FieldToMatch

```

## Properties

`LoggingFilter`

A filter configuration that determines which WAF log records to include or exclude.

_Required_: No

_Type_: [LoggingFilter](aws-properties-observabilityadmin-organizationtelemetryrule-loggingfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogType`

The type of WAF logs to collect (currently supports WAF\_LOGS).

_Required_: No

_Type_: String

_Allowed values_: `WAF_LOGS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedactedFields`

The fields to redact from WAF logs to protect sensitive information.

_Required_: No

_Type_: Array of [FieldToMatch](aws-properties-observabilityadmin-organizationtelemetryrule-fieldtomatch.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VPCFlowLogParameters

AWS::ObservabilityAdmin::S3TableIntegration

All content copied from https://docs.aws.amazon.com/.
