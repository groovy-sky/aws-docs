This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Synthetics::Canary VisualReference

Defines the screenshots to use as the baseline for comparisons during visual monitoring comparisons during future runs of this canary. If you omit this
parameter, no changes are made to any baseline screenshots that the canary might be using already.

Visual monitoring is supported only on canaries running the **syn-puppeteer-node-3.2**
runtime or later. For more information, see [Visual monitoring](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Library_SyntheticsLogger_VisualTesting.html) and [Visual monitoring blueprint](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Blueprints_VisualTesting.html)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaseCanaryRunId" : String,
  "BaseScreenshots" : [ BaseScreenshot, ... ],
  "BrowserType" : String
}

```

### YAML

```yaml

  BaseCanaryRunId: String
  BaseScreenshots:
    - BaseScreenshot
  BrowserType: String

```

## Properties

`BaseCanaryRunId`

Specifies which canary run to use the screenshots from as the baseline for future visual monitoring with this canary. Valid values are
`nextrun` to use the screenshots from the next run after this update is made, `lastrun` to use the screenshots from the most recent run
before this update was made, or the value of `Id` in the [CanaryRun](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CanaryRun.html) from any past run of this canary.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BaseScreenshots`

An array of screenshots that are used as the baseline for comparisons during visual monitoring.

_Required_: No

_Type_: Array of [BaseScreenshot](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-synthetics-canary-basescreenshot.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BrowserType`

The browser type associated with this visual reference configuration. Valid values are `CHROME` and `FIREFOX`.

_Required_: No

_Type_: String

_Allowed values_: `CHROME | FIREFOX`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

VPCConfig
