This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::Detector CFNKubernetesConfiguration

Describes which Kubernetes protection data sources are enabled for the detector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuditLogs" : CFNKubernetesAuditLogsConfiguration
}

```

### YAML

```yaml

  AuditLogs:
    CFNKubernetesAuditLogsConfiguration

```

## Properties

`AuditLogs`

Describes whether Kubernetes audit logs are enabled as a data source for the
detector.

_Required_: Yes

_Type_: [CFNKubernetesAuditLogsConfiguration](aws-properties-guardduty-detector-cfnkubernetesauditlogsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CFNKubernetesAuditLogsConfiguration

CFNMalwareProtectionConfiguration

All content copied from https://docs.aws.amazon.com/.
