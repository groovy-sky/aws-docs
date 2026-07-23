---
title: "AWS::AppStream::Stack AgentAccessConfig"
---

This is the new *CloudFormation Template Reference Guide*. Please update your bookmarks and links. For help getting started with CloudFormation, see the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html).

# AWS::AppStream::Stack AgentAccessConfig
<a name="aws-properties-appstream-stack-agentaccessconfig"></a>

<a name="aws-properties-appstream-stack-agentaccessconfig-description"></a>The `AgentAccessConfig` property type specifies Property description not available. for an [AWS::AppStream::Stack](aws-resource-appstream-stack.md).

## Syntax
<a name="aws-properties-appstream-stack-agentaccessconfig-syntax"></a>

To declare this entity in your CloudFormation template, use the following syntax:

### JSON
<a name="aws-properties-appstream-stack-agentaccessconfig-syntax.json"></a>

```
{
  "[S3BucketArn](#cfn-appstream-stack-agentaccessconfig-s3bucketarn)" : {{String}},
  "[ScreenImageFormat](#cfn-appstream-stack-agentaccessconfig-screenimageformat)" : {{String}},
  "[ScreenResolution](#cfn-appstream-stack-agentaccessconfig-screenresolution)" : {{String}},
  "[ScreenshotsUploadEnabled](#cfn-appstream-stack-agentaccessconfig-screenshotsuploadenabled)" : {{Boolean}},
  "[Settings](#cfn-appstream-stack-agentaccessconfig-settings)" : {{[ AgentAccessSetting, ... ]}},
  "[UserControlMode](#cfn-appstream-stack-agentaccessconfig-usercontrolmode)" : {{String}}
}
```

### YAML
<a name="aws-properties-appstream-stack-agentaccessconfig-syntax.yaml"></a>

```
  [S3BucketArn](#cfn-appstream-stack-agentaccessconfig-s3bucketarn): {{String}}
  [ScreenImageFormat](#cfn-appstream-stack-agentaccessconfig-screenimageformat): {{String}}
  [ScreenResolution](#cfn-appstream-stack-agentaccessconfig-screenresolution): {{String}}
  [ScreenshotsUploadEnabled](#cfn-appstream-stack-agentaccessconfig-screenshotsuploadenabled): {{Boolean}}
  [Settings](#cfn-appstream-stack-agentaccessconfig-settings): {{
    - AgentAccessSetting}}
  [UserControlMode](#cfn-appstream-stack-agentaccessconfig-usercontrolmode): {{String}}
```

## Properties
<a name="aws-properties-appstream-stack-agentaccessconfig-properties"></a>

`S3BucketArn`  <a name="cfn-appstream-stack-agentaccessconfig-s3bucketarn"></a>
Property description not available.
*Required*: No
*Type*: String
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`ScreenImageFormat`  <a name="cfn-appstream-stack-agentaccessconfig-screenimageformat"></a>
Property description not available.
*Required*: Yes
*Type*: String
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`ScreenResolution`  <a name="cfn-appstream-stack-agentaccessconfig-screenresolution"></a>
Property description not available.
*Required*: Yes
*Type*: String
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`ScreenshotsUploadEnabled`  <a name="cfn-appstream-stack-agentaccessconfig-screenshotsuploadenabled"></a>
Property description not available.
*Required*: No
*Type*: Boolean
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`Settings`  <a name="cfn-appstream-stack-agentaccessconfig-settings"></a>
Property description not available.
*Required*: Yes
*Type*: Array of [AgentAccessSetting](aws-properties-appstream-stack-agentaccesssetting.md)
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`UserControlMode`  <a name="cfn-appstream-stack-agentaccessconfig-usercontrolmode"></a>
Property description not available.
*Required*: No
*Type*: String
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

All content copied from https://docs.aws.amazon.com/.
