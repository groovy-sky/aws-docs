---
title: "AWS::SSMContacts::Contact Tag"
---

This is the new *CloudFormation Template Reference Guide*. Please update your bookmarks and links. For help getting started with CloudFormation, see the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html).

# AWS::SSMContacts::Contact Tag
<a name="aws-properties-ssmcontacts-contact-tag"></a>

A container of a key-value name pair.

## Syntax
<a name="aws-properties-ssmcontacts-contact-tag-syntax"></a>

To declare this entity in your CloudFormation template, use the following syntax:

### JSON
<a name="aws-properties-ssmcontacts-contact-tag-syntax.json"></a>

```
{
  "[Key](#cfn-ssmcontacts-contact-tag-key)" : {{String}},
  "[Value](#cfn-ssmcontacts-contact-tag-value)" : {{String}}
}
```

### YAML
<a name="aws-properties-ssmcontacts-contact-tag-syntax.yaml"></a>

```
  [Key](#cfn-ssmcontacts-contact-tag-key): {{String}}
  [Value](#cfn-ssmcontacts-contact-tag-value): {{String}}
```

## Properties
<a name="aws-properties-ssmcontacts-contact-tag-properties"></a>

`Key`  <a name="cfn-ssmcontacts-contact-tag-key"></a>
Name of the object key.
*Required*: Yes
*Type*: String
*Minimum*: `1`
*Maximum*: `128`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`Value`  <a name="cfn-ssmcontacts-contact-tag-value"></a>
Value of the tag.
*Required*: Yes
*Type*: String
*Minimum*: `0`
*Maximum*: `256`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

All content copied from https://docs.aws.amazon.com/.
