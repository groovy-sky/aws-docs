This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::LifecyclePolicy ExclusionRules

Specifies resources that lifecycle policy actions should not apply to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Amis" : AmiExclusionRules,
  "TagMap" : {Key: Value, ...}
}

```

### YAML

```yaml

  Amis:
    AmiExclusionRules
  TagMap:
    Key: Value

```

## Properties

`Amis`

Lists configuration values that apply to AMIs that Image Builder should exclude
from the lifecycle action.

_Required_: No

_Type_: [AmiExclusionRules](aws-properties-imagebuilder-lifecyclepolicy-amiexclusionrules.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagMap`

Contains a list of tags that Image Builder uses to skip lifecycle actions for Image Builder image
resources that have them.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AmiExclusionRules

Filter

All content copied from https://docs.aws.amazon.com/.
