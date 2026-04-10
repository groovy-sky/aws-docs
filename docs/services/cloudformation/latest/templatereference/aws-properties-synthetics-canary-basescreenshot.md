This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Synthetics::Canary BaseScreenshot

A structure representing a screenshot that is used as a baseline during visual monitoring comparisons made by the canary.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IgnoreCoordinates" : [ String, ... ],
  "ScreenshotName" : String
}

```

### YAML

```yaml

  IgnoreCoordinates:
    - String
  ScreenshotName: String

```

## Properties

`IgnoreCoordinates`

Coordinates that define the part of a screen to ignore during screenshot comparisons. To obtain the coordinates to use here, use the
CloudWatch console to draw the boundaries on the screen. For more information, see
[Edit or delete a canary](../../../amazoncloudwatch/latest/monitoring/synthetics-canaries-deletion.md).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScreenshotName`

The name of the screenshot. This is generated the first time the canary is run after the `UpdateCanary` operation that
specified for this canary to perform visual monitoring.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArtifactConfig

BrowserConfig

All content copied from https://docs.aws.amazon.com/.
