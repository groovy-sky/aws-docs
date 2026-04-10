This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::Detector CFNDataSourceConfigurations

Describes whether S3 data event logs, Kubernetes audit logs, or Malware Protection will
be enabled as a data source when the detector is created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Kubernetes" : CFNKubernetesConfiguration,
  "MalwareProtection" : CFNMalwareProtectionConfiguration,
  "S3Logs" : CFNS3LogsConfiguration
}

```

### YAML

```yaml

  Kubernetes:
    CFNKubernetesConfiguration
  MalwareProtection:
    CFNMalwareProtectionConfiguration
  S3Logs:
    CFNS3LogsConfiguration

```

## Properties

`Kubernetes`

Describes which Kubernetes data sources are enabled for a detector.

_Required_: No

_Type_: [CFNKubernetesConfiguration](aws-properties-guardduty-detector-cfnkubernetesconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MalwareProtection`

Describes whether Malware Protection will be enabled as a data source.

_Required_: No

_Type_: [CFNMalwareProtectionConfiguration](aws-properties-guardduty-detector-cfnmalwareprotectionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Logs`

Describes whether S3 data event logs are enabled as a data source.

_Required_: No

_Type_: [CFNS3LogsConfiguration](aws-properties-guardduty-detector-cfns3logsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GuardDuty::Detector

CFNFeatureAdditionalConfiguration

All content copied from https://docs.aws.amazon.com/.
