This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation LogsConfigurationPolicy

Provides the information necessary for a user to access the logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedAccountIds" : [ String, ... ],
  "FilterPattern" : String,
  "LogRedactionConfiguration" : LogRedactionConfiguration,
  "LogType" : String
}

```

### YAML

```yaml

  AllowedAccountIds:
    - String
  FilterPattern: String
  LogRedactionConfiguration:
    LogRedactionConfiguration
  LogType: String

```

## Properties

`AllowedAccountIds`

A list of account IDs that are allowed to access the logs.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FilterPattern`

A regular expression pattern that is used to parse the logs and return information that
matches the pattern.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogRedactionConfiguration`

Specifies the log redaction configuration for this policy.

_Required_: No

_Type_: [LogRedactionConfiguration](aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logredactionconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogType`

Specifies the type of log this policy applies to. The currently supported policies are ALL or ERROR\_SUMMARY.

_Required_: No

_Type_: String

_Allowed values_: `ALL | ERROR_SUMMARY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogRedactionConfiguration

MetricsConfigurationPolicy

All content copied from https://docs.aws.amazon.com/.
