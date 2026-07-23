---
title: "AWS::CloudWatch::AlarmMuteRule MuteTargets"
---

This is the new *CloudFormation Template Reference Guide*. Please update your bookmarks and links. For help getting started with CloudFormation, see the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html).

# AWS::CloudWatch::AlarmMuteRule MuteTargets
<a name="aws-properties-cloudwatch-alarmmuterule-mutetargets"></a>

Specifies which alarms this rule applies to.

## Syntax
<a name="aws-properties-cloudwatch-alarmmuterule-mutetargets-syntax"></a>

To declare this entity in your CloudFormation template, use the following syntax:

### JSON
<a name="aws-properties-cloudwatch-alarmmuterule-mutetargets-syntax.json"></a>

```
{
  "[AlarmNames](#cfn-cloudwatch-alarmmuterule-mutetargets-alarmnames)" : {{[ String, ... ]}}
}
```

### YAML
<a name="aws-properties-cloudwatch-alarmmuterule-mutetargets-syntax.yaml"></a>

```
  [AlarmNames](#cfn-cloudwatch-alarmmuterule-mutetargets-alarmnames): {{
    - String}}
```

## Properties
<a name="aws-properties-cloudwatch-alarmmuterule-mutetargets-properties"></a>

`AlarmNames`  <a name="cfn-cloudwatch-alarmmuterule-mutetargets-alarmnames"></a>
Specifies which alarms this rule applies to.
*Required*: Yes
*Type*: Array of String
*Maximum*: `500`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

All content copied from https://docs.aws.amazon.com/.
