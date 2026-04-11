This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::LifecyclePolicy AmiExclusionRules

Defines criteria for AMIs that are excluded from lifecycle actions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsPublic" : Boolean,
  "LastLaunched" : LastLaunched,
  "Regions" : [ String, ... ],
  "SharedAccounts" : [ String, ... ],
  "TagMap" : {Key: Value, ...}
}

```

### YAML

```yaml

  IsPublic: Boolean
  LastLaunched:
    LastLaunched
  Regions:
    - String
  SharedAccounts:
    - String
  TagMap:
    Key: Value

```

## Properties

`IsPublic`

Configures whether public AMIs are excluded from the lifecycle action.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastLaunched`

Specifies configuration details for Image Builder to exclude the most recent resources
from lifecycle actions.

_Required_: No

_Type_: [LastLaunched](aws-properties-imagebuilder-lifecyclepolicy-lastlaunched.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Regions`

Configures AWS Regions that are excluded from the lifecycle action.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SharedAccounts`

Specifies AWS accounts whose resources are excluded from the lifecycle action.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1536`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagMap`

Lists tags that should be excluded from lifecycle actions for the AMIs that have them.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Action

ExclusionRules

All content copied from https://docs.aws.amazon.com/.
