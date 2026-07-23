---
title: "AWS::PCS::ComputeNodeGroup NodeLifecycleScript"
---

This is the new *CloudFormation Template Reference Guide*. Please update your bookmarks and links. For help getting started with CloudFormation, see the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html).

# AWS::PCS::ComputeNodeGroup NodeLifecycleScript
<a name="aws-properties-pcs-computenodegroup-nodelifecyclescript"></a>

<a name="aws-properties-pcs-computenodegroup-nodelifecyclescript-description"></a>The `NodeLifecycleScript` property type specifies Property description not available. for an [AWS::PCS::ComputeNodeGroup](aws-resource-pcs-computenodegroup.md).

## Syntax
<a name="aws-properties-pcs-computenodegroup-nodelifecyclescript-syntax"></a>

To declare this entity in your CloudFormation template, use the following syntax:

### JSON
<a name="aws-properties-pcs-computenodegroup-nodelifecyclescript-syntax.json"></a>

```
{
  "[Arguments](#cfn-pcs-computenodegroup-nodelifecyclescript-arguments)" : {{[ String, ... ]}},
  "[ExecutionPolicy](#cfn-pcs-computenodegroup-nodelifecyclescript-executionpolicy)" : {{String}},
  "[Name](#cfn-pcs-computenodegroup-nodelifecyclescript-name)" : {{String}},
  "[OnError](#cfn-pcs-computenodegroup-nodelifecyclescript-onerror)" : {{String}},
  "[ScriptSource](#cfn-pcs-computenodegroup-nodelifecyclescript-scriptsource)" : {{ScriptSource}}
}
```

### YAML
<a name="aws-properties-pcs-computenodegroup-nodelifecyclescript-syntax.yaml"></a>

```
  [Arguments](#cfn-pcs-computenodegroup-nodelifecyclescript-arguments): {{
    - String}}
  [ExecutionPolicy](#cfn-pcs-computenodegroup-nodelifecyclescript-executionpolicy): {{String}}
  [Name](#cfn-pcs-computenodegroup-nodelifecyclescript-name): {{String}}
  [OnError](#cfn-pcs-computenodegroup-nodelifecyclescript-onerror): {{String}}
  [ScriptSource](#cfn-pcs-computenodegroup-nodelifecyclescript-scriptsource): {{
    ScriptSource}}
```

## Properties
<a name="aws-properties-pcs-computenodegroup-nodelifecyclescript-properties"></a>

`Arguments`  <a name="cfn-pcs-computenodegroup-nodelifecyclescript-arguments"></a>
Property description not available.
*Required*: No
*Type*: Array of String
*Maximum*: `256 | 20`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`ExecutionPolicy`  <a name="cfn-pcs-computenodegroup-nodelifecyclescript-executionpolicy"></a>
Property description not available.
*Required*: No
*Type*: String
*Allowed values*: `FIRST_BOOT_ONLY | EVERY_BOOT`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`Name`  <a name="cfn-pcs-computenodegroup-nodelifecyclescript-name"></a>
Property description not available.
*Required*: Yes
*Type*: String
*Pattern*: `^[A-Za-z0-9][A-Za-z0-9 _-]*$`
*Minimum*: `1`
*Maximum*: `64`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`OnError`  <a name="cfn-pcs-computenodegroup-nodelifecyclescript-onerror"></a>
Property description not available.
*Required*: No
*Type*: String
*Allowed values*: `TERMINATE | STOP_SEQUENCE | CONTINUE`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`ScriptSource`  <a name="cfn-pcs-computenodegroup-nodelifecyclescript-scriptsource"></a>
Property description not available.
*Required*: Yes
*Type*: [ScriptSource](aws-properties-pcs-computenodegroup-scriptsource.md)
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

All content copied from https://docs.aws.amazon.com/.
