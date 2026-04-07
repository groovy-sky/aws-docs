This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::CloudVmCluster DataCollectionOptions

Information about the data collection options enabled for a VM cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsDiagnosticsEventsEnabled" : Boolean,
  "IsHealthMonitoringEnabled" : Boolean,
  "IsIncidentLogsEnabled" : Boolean
}

```

### YAML

```yaml

  IsDiagnosticsEventsEnabled: Boolean
  IsHealthMonitoringEnabled: Boolean
  IsIncidentLogsEnabled: Boolean

```

## Properties

`IsDiagnosticsEventsEnabled`

Specifies whether diagnostic collection is enabled for the VM cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IsHealthMonitoringEnabled`

Specifies whether health monitoring is enabled for the VM cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IsIncidentLogsEnabled`

Specifies whether incident logs are enabled for the VM cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ODB::CloudVmCluster

DbNode
