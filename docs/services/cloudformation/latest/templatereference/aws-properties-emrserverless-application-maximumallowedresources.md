This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application MaximumAllowedResources

The maximum allowed cumulative resources for an application. No new resources will be
created once the limit is hit.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cpu" : String,
  "Disk" : String,
  "Memory" : String
}

```

### YAML

```yaml

  Cpu: String
  Disk: String
  Memory: String

```

## Properties

`Cpu`

The maximum allowed CPU for an application.

_Required_: Yes

_Type_: String

_Pattern_: `^[1-9][0-9]*(\s)?(vCPU|vcpu|VCPU)?$`

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Disk`

The maximum allowed disk for an application.

_Required_: No

_Type_: String

_Pattern_: `^[1-9][0-9]*(\s)?(GB|gb|gB|Gb)$`

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Memory`

The maximum allowed resources for an application.

_Required_: Yes

_Type_: String

_Pattern_: `^[1-9][0-9]*(\s)?(GB|gb|gB|Gb)?$`

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedPersistenceMonitoringConfiguration

MonitoringConfiguration

All content copied from https://docs.aws.amazon.com/.
