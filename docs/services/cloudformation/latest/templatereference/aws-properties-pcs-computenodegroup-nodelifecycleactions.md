---
title: "AWS::PCS::ComputeNodeGroup NodeLifecycleActions"
---

This is the new *CloudFormation Template Reference Guide*. Please update your bookmarks and links. For help getting started with CloudFormation, see the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html).

# AWS::PCS::ComputeNodeGroup NodeLifecycleActions
<a name="aws-properties-pcs-computenodegroup-nodelifecycleactions"></a>

<a name="aws-properties-pcs-computenodegroup-nodelifecycleactions-description"></a>The `NodeLifecycleActions` property type specifies Property description not available. for an [AWS::PCS::ComputeNodeGroup](aws-resource-pcs-computenodegroup.md).

## Syntax
<a name="aws-properties-pcs-computenodegroup-nodelifecycleactions-syntax"></a>

To declare this entity in your CloudFormation template, use the following syntax:

### JSON
<a name="aws-properties-pcs-computenodegroup-nodelifecycleactions-syntax.json"></a>

```
{
  "[ScriptCachingPolicy](#cfn-pcs-computenodegroup-nodelifecycleactions-scriptcachingpolicy)" : {{String}},
  "[Stages](#cfn-pcs-computenodegroup-nodelifecycleactions-stages)" : {{NodeLifecycleStages}}
}
```

### YAML
<a name="aws-properties-pcs-computenodegroup-nodelifecycleactions-syntax.yaml"></a>

```
  [ScriptCachingPolicy](#cfn-pcs-computenodegroup-nodelifecycleactions-scriptcachingpolicy): {{String}}
  [Stages](#cfn-pcs-computenodegroup-nodelifecycleactions-stages): {{
    NodeLifecycleStages}}
```

## Properties
<a name="aws-properties-pcs-computenodegroup-nodelifecycleactions-properties"></a>

`ScriptCachingPolicy`  <a name="cfn-pcs-computenodegroup-nodelifecycleactions-scriptcachingpolicy"></a>
Property description not available.
*Required*: No
*Type*: String
*Allowed values*: `CACHE_ONCE | REFRESH_ON_REBOOT`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`Stages`  <a name="cfn-pcs-computenodegroup-nodelifecycleactions-stages"></a>
Property description not available.
*Required*: Yes
*Type*: [NodeLifecycleStages](aws-properties-pcs-computenodegroup-nodelifecyclestages.md)
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

All content copied from https://docs.aws.amazon.com/.
