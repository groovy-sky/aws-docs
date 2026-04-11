This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::Detector CFNFeatureConfiguration

Information about the configuration of a feature in your account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalConfiguration" : [ CFNFeatureAdditionalConfiguration, ... ],
  "Name" : String,
  "Status" : String
}

```

### YAML

```yaml

  AdditionalConfiguration:
    - CFNFeatureAdditionalConfiguration
  Name: String
  Status: String

```

## Properties

`AdditionalConfiguration`

Information about the additional configuration of a feature in your account.

_Required_: No

_Type_: Array of [CFNFeatureAdditionalConfiguration](aws-properties-guardduty-detector-cfnfeatureadditionalconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of the feature. For a list of allowed values, see [DetectorFeatureConfiguration](../../../../reference/guardduty/latest/apireference/api-detectorfeatureconfiguration.md#guardduty-Type-DetectorFeatureConfiguration-name) in the _GuardDuty API Reference_.

_Required_: Yes

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Status of the feature configuration.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CFNFeatureAdditionalConfiguration

CFNKubernetesAuditLogsConfiguration

All content copied from https://docs.aws.amazon.com/.
