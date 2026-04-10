This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan EksResourceScalingConfiguration

The AWS EKS resource scaling configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityMonitoringApproach" : ,
  "EksClusters" : [ EksCluster, ... ],
  "KubernetesResourceType" : KubernetesResourceType,
  "ScalingResources" : [ {Key: Value, ...}, ... ],
  "TargetPercent" : Number,
  "TimeoutMinutes" : Number,
  "Ungraceful" : EksResourceScalingUngraceful
}

```

### YAML

```yaml

  CapacityMonitoringApproach:

  EksClusters:
    - EksCluster
  KubernetesResourceType:
    KubernetesResourceType
  ScalingResources:
    -
    Key: Value
  TargetPercent: Number
  TimeoutMinutes: Number
  Ungraceful:
    EksResourceScalingUngraceful

```

## Properties

`CapacityMonitoringApproach`

The monitoring approach for the configuration, that is, whether it was sampled in the last
24 hours or autoscaled in the last 24 hours.

_Required_: No

_Type_:

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EksClusters`

The clusters for the configuration.

_Required_: No

_Type_: Array of [EksCluster](aws-properties-arcregionswitch-plan-ekscluster.md)

_Minimum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KubernetesResourceType`

The Kubernetes resource type for the configuration.

_Required_: Yes

_Type_: [KubernetesResourceType](aws-properties-arcregionswitch-plan-kubernetesresourcetype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingResources`

The scaling resources for the configuration.

_Required_: No

_Type_: Array of Object

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetPercent`

The target percentage for the configuration. The default is 100.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutMinutes`

The timeout value specified for the configuration.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ungraceful`

The settings for ungraceful execution.

_Required_: No

_Type_: [EksResourceScalingUngraceful](aws-properties-arcregionswitch-plan-eksresourcescalingungraceful.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksCluster

EksResourceScalingUngraceful

All content copied from https://docs.aws.amazon.com/.
