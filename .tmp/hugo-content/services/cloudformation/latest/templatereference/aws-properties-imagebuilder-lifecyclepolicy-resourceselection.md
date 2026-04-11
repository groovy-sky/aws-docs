This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::LifecyclePolicy ResourceSelection

Resource selection criteria for the lifecycle policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Recipes" : [ RecipeSelection, ... ],
  "TagMap" : {Key: Value, ...}
}

```

### YAML

```yaml

  Recipes:
    - RecipeSelection
  TagMap:
    Key: Value

```

## Properties

`Recipes`

A list of recipes that are used as selection criteria for the output
images that the lifecycle policy applies to.

_Required_: No

_Type_: Array of [RecipeSelection](aws-properties-imagebuilder-lifecyclepolicy-recipeselection.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagMap`

A list of tags that are used as selection criteria for the Image Builder image
resources that the lifecycle policy applies to.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecipeSelection

AWS::ImageBuilder::Workflow

All content copied from https://docs.aws.amazon.com/.
