This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Synthetics::Canary VisualReference

Defines the screenshots to use as the baseline for comparisons during visual monitoring comparisons during future runs of this canary. If you omit this
parameter, no changes are made to any baseline screenshots that the canary might be using already.

Visual monitoring is supported only on canaries running the **syn-puppeteer-node-3.2**
runtime or later. For more information, see [Visual monitoring](../../../amazoncloudwatch/latest/monitoring/cloudwatch-synthetics-library-syntheticslogger-visualtesting.md) and [Visual monitoring blueprint](../../../amazoncloudwatch/latest/monitoring/cloudwatch-synthetics-canaries-blueprints-visualtesting.md)

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
before this update was made, or the value of `Id` in the [CanaryRun](../../../../reference/amazonsynthetics/latest/apireference/api-canaryrun.md) from any past run of this canary.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BaseScreenshots`

An array of screenshots that are used as the baseline for comparisons during visual monitoring.

_Required_: No

_Type_: Array of [BaseScreenshot](aws-properties-synthetics-canary-basescreenshot.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BrowserType`

The browser type associated with this visual reference configuration. Valid values are `CHROME` and `FIREFOX`.

_Required_: No

_Type_: String

_Allowed values_: `CHROME | FIREFOX`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

VPCConfig

All content copied from https://docs.aws.amazon.com/.
