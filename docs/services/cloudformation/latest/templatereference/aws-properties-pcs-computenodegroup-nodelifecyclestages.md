---
title: "AWS::PCS::ComputeNodeGroup NodeLifecycleStages"
---

This is the new *CloudFormation Template Reference Guide*. Please update your bookmarks and links. For help getting started with CloudFormation, see the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html).

# AWS::PCS::ComputeNodeGroup NodeLifecycleStages
<a name="aws-properties-pcs-computenodegroup-nodelifecyclestages"></a>

<a name="aws-properties-pcs-computenodegroup-nodelifecyclestages-description"></a>The `NodeLifecycleStages` property type specifies Property description not available. for an [AWS::PCS::ComputeNodeGroup](aws-resource-pcs-computenodegroup.md).

## Syntax
<a name="aws-properties-pcs-computenodegroup-nodelifecyclestages-syntax"></a>

To declare this entity in your CloudFormation template, use the following syntax:

### JSON
<a name="aws-properties-pcs-computenodegroup-nodelifecyclestages-syntax.json"></a>

```
{
  "[NodeBootstrapped](#cfn-pcs-computenodegroup-nodelifecyclestages-nodebootstrapped)" : {{[ NodeLifecycleScript, ... ]}},
  "[NodeReady](#cfn-pcs-computenodegroup-nodelifecyclestages-nodeready)" : {{[ NodeLifecycleScript, ... ]}}
}
```

### YAML
<a name="aws-properties-pcs-computenodegroup-nodelifecyclestages-syntax.yaml"></a>

```
  [NodeBootstrapped](#cfn-pcs-computenodegroup-nodelifecyclestages-nodebootstrapped): {{
    - NodeLifecycleScript}}
  [NodeReady](#cfn-pcs-computenodegroup-nodelifecyclestages-nodeready): {{
    - NodeLifecycleScript}}
```

## Properties
<a name="aws-properties-pcs-computenodegroup-nodelifecyclestages-properties"></a>

`NodeBootstrapped`  <a name="cfn-pcs-computenodegroup-nodelifecyclestages-nodebootstrapped"></a>
Property description not available.
*Required*: No
*Type*: Array of [NodeLifecycleScript](aws-properties-pcs-computenodegroup-nodelifecyclescript.md)
*Minimum*: `1`
*Maximum*: `20`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`NodeReady`  <a name="cfn-pcs-computenodegroup-nodelifecyclestages-nodeready"></a>
Property description not available.
*Required*: No
*Type*: Array of [NodeLifecycleScript](aws-properties-pcs-computenodegroup-nodelifecyclescript.md)
*Minimum*: `1`
*Maximum*: `20`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

All content copied from https://docs.aws.amazon.com/.
