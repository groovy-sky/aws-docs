---
title: "AWS::PCS::ComputeNodeGroup ScriptSource"
---

This is the new *CloudFormation Template Reference Guide*. Please update your bookmarks and links. For help getting started with CloudFormation, see the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html).

# AWS::PCS::ComputeNodeGroup ScriptSource
<a name="aws-properties-pcs-computenodegroup-scriptsource"></a>

<a name="aws-properties-pcs-computenodegroup-scriptsource-description"></a>The `ScriptSource` property type specifies Property description not available. for an [AWS::PCS::ComputeNodeGroup](aws-resource-pcs-computenodegroup.md).

## Syntax
<a name="aws-properties-pcs-computenodegroup-scriptsource-syntax"></a>

To declare this entity in your CloudFormation template, use the following syntax:

### JSON
<a name="aws-properties-pcs-computenodegroup-scriptsource-syntax.json"></a>

```
{
  "[Checksum](#cfn-pcs-computenodegroup-scriptsource-checksum)" : {{String}},
  "[S3VersionId](#cfn-pcs-computenodegroup-scriptsource-s3versionid)" : {{String}},
  "[ScriptLocation](#cfn-pcs-computenodegroup-scriptsource-scriptlocation)" : {{String}}
}
```

### YAML
<a name="aws-properties-pcs-computenodegroup-scriptsource-syntax.yaml"></a>

```
  [Checksum](#cfn-pcs-computenodegroup-scriptsource-checksum): {{String}}
  [S3VersionId](#cfn-pcs-computenodegroup-scriptsource-s3versionid): {{String}}
  [ScriptLocation](#cfn-pcs-computenodegroup-scriptsource-scriptlocation): {{String}}
```

## Properties
<a name="aws-properties-pcs-computenodegroup-scriptsource-properties"></a>

`Checksum`  <a name="cfn-pcs-computenodegroup-scriptsource-checksum"></a>
Property description not available.
*Required*: No
*Type*: String
*Pattern*: `^[a-fA-F0-9]{64}$`
*Minimum*: `64`
*Maximum*: `64`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`S3VersionId`  <a name="cfn-pcs-computenodegroup-scriptsource-s3versionid"></a>
Property description not available.
*Required*: No
*Type*: String
*Maximum*: `1024`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`ScriptLocation`  <a name="cfn-pcs-computenodegroup-scriptsource-scriptlocation"></a>
Property description not available.
*Required*: Yes
*Type*: String
*Pattern*: `^(s3://[a-z0-9][a-z0-9.-]{1,61}[a-z0-9]/.+|https://.+)$`
*Minimum*: `1`
*Maximum*: `1024`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

All content copied from https://docs.aws.amazon.com/.
