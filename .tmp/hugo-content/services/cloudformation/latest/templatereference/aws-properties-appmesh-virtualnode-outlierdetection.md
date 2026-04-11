This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode OutlierDetection

An object that represents the outlier detection for a virtual node's listener.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaseEjectionDuration" : Duration,
  "Interval" : Duration,
  "MaxEjectionPercent" : Integer,
  "MaxServerErrors" : Integer
}

```

### YAML

```yaml

  BaseEjectionDuration:
    Duration
  Interval:
    Duration
  MaxEjectionPercent: Integer
  MaxServerErrors: Integer

```

## Properties

`BaseEjectionDuration`

The base amount of time for which a host is ejected.

_Required_: Yes

_Type_: [Duration](aws-properties-appmesh-virtualnode-duration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interval`

The time interval between ejection sweep analysis.

_Required_: Yes

_Type_: [Duration](aws-properties-appmesh-virtualnode-duration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxEjectionPercent`

Maximum percentage of hosts in load balancing pool for upstream service that can be
ejected. Will eject at least one host regardless of the value.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxServerErrors`

Number of consecutive `5xx` errors required for ejection.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggingFormat

PortMapping

All content copied from https://docs.aws.amazon.com/.
