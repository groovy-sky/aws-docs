---
title: "AWS::ConnectCampaigns::Campaign AnswerMachineDetectionConfig"
---

This is the new *CloudFormation Template Reference Guide*. Please update your bookmarks and links. For help getting started with CloudFormation, see the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html).

# AWS::ConnectCampaigns::Campaign AnswerMachineDetectionConfig
<a name="aws-properties-connectcampaigns-campaign-answermachinedetectionconfig"></a>

Contains information about answering machine detection.

## Syntax
<a name="aws-properties-connectcampaigns-campaign-answermachinedetectionconfig-syntax"></a>

To declare this entity in your CloudFormation template, use the following syntax:

### JSON
<a name="aws-properties-connectcampaigns-campaign-answermachinedetectionconfig-syntax.json"></a>

```
{
  "[AwaitAnswerMachinePrompt](#cfn-connectcampaigns-campaign-answermachinedetectionconfig-awaitanswermachineprompt)" : {{Boolean}},
  "[EnableAnswerMachineDetection](#cfn-connectcampaigns-campaign-answermachinedetectionconfig-enableanswermachinedetection)" : {{Boolean}}
}
```

### YAML
<a name="aws-properties-connectcampaigns-campaign-answermachinedetectionconfig-syntax.yaml"></a>

```
  [AwaitAnswerMachinePrompt](#cfn-connectcampaigns-campaign-answermachinedetectionconfig-awaitanswermachineprompt): {{Boolean}}
  [EnableAnswerMachineDetection](#cfn-connectcampaigns-campaign-answermachinedetectionconfig-enableanswermachinedetection): {{Boolean}}
```

## Properties
<a name="aws-properties-connectcampaigns-campaign-answermachinedetectionconfig-properties"></a>

`AwaitAnswerMachinePrompt`  <a name="cfn-connectcampaigns-campaign-answermachinedetectionconfig-awaitanswermachineprompt"></a>
Whether waiting for answer machine prompt is enabled.
*Required*: No
*Type*: Boolean
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`EnableAnswerMachineDetection`  <a name="cfn-connectcampaigns-campaign-answermachinedetectionconfig-enableanswermachinedetection"></a>
Whether answering machine detection is enabled.
*Required*: Yes
*Type*: Boolean
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

All content copied from https://docs.aws.amazon.com/.
