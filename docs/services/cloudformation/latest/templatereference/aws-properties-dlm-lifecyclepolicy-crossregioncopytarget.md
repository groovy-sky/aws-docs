This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy CrossRegionCopyTarget

**\[Default policies only\]** Specifies a destination Region for cross-Region copy actions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetRegion" : String
}

```

### YAML

```yaml

  TargetRegion: String

```

## Properties

`TargetRegion`

The target Region, for example `us-east-1`.

_Required_: No

_Type_: String

_Pattern_: `([a-z]+-){2,3}\d`

_Minimum_: `0`

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CrossRegionCopyRule

DeprecateRule
