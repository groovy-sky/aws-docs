---
title: "AWS::EC2::SpotFleet TotalLocalStorageGBRequest"
---

This is the new *CloudFormation Template Reference Guide*. Please update your bookmarks and links. For help getting started with CloudFormation, see the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html).

# AWS::EC2::SpotFleet TotalLocalStorageGBRequest
<a name="aws-properties-ec2-spotfleet-totallocalstoragegbrequest"></a>

The minimum and maximum amount of total local storage, in GB.

## Syntax
<a name="aws-properties-ec2-spotfleet-totallocalstoragegbrequest-syntax"></a>

To declare this entity in your CloudFormation template, use the following syntax:

### JSON
<a name="aws-properties-ec2-spotfleet-totallocalstoragegbrequest-syntax.json"></a>

```
{
  "[Max](#cfn-ec2-spotfleet-totallocalstoragegbrequest-max)" : {{Number}},
  "[Min](#cfn-ec2-spotfleet-totallocalstoragegbrequest-min)" : {{Number}}
}
```

### YAML
<a name="aws-properties-ec2-spotfleet-totallocalstoragegbrequest-syntax.yaml"></a>

```
  [Max](#cfn-ec2-spotfleet-totallocalstoragegbrequest-max): {{Number}}
  [Min](#cfn-ec2-spotfleet-totallocalstoragegbrequest-min): {{Number}}
```

## Properties
<a name="aws-properties-ec2-spotfleet-totallocalstoragegbrequest-properties"></a>

`Max`  <a name="cfn-ec2-spotfleet-totallocalstoragegbrequest-max"></a>
The maximum amount of total local storage, in GB. To specify no maximum limit, omit this parameter.
*Required*: No
*Type*: Number
*Update requires*: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

`Min`  <a name="cfn-ec2-spotfleet-totallocalstoragegbrequest-min"></a>
The minimum amount of total local storage, in GB. To specify no minimum limit, omit this parameter.
*Required*: No
*Type*: Number
*Update requires*: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

All content copied from https://docs.aws.amazon.com/.
