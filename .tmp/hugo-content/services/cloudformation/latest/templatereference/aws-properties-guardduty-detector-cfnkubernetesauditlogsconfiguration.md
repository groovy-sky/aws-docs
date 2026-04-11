This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::Detector CFNKubernetesAuditLogsConfiguration

Describes which optional data sources are enabled for a detector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enable" : Boolean
}

```

### YAML

```yaml

  Enable: Boolean

```

## Properties

`Enable`

Describes whether Kubernetes audit logs are enabled as a data source for the
detector.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CFNFeatureConfiguration

CFNKubernetesConfiguration

All content copied from https://docs.aws.amazon.com/.
