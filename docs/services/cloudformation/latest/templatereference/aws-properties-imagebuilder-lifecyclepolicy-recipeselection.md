This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::LifecyclePolicy RecipeSelection

Specifies an Image Builder recipe that the lifecycle policy uses for resource selection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "SemanticVersion" : String
}

```

### YAML

```yaml

  Name: String
  SemanticVersion: String

```

## Properties

`Name`

The name of an Image Builder recipe that the lifecycle policy uses for resource selection.

_Required_: Yes

_Type_: String

_Pattern_: `^[-_A-Za-z-0-9][-_A-Za-z0-9 ]{1,126}[-_A-Za-z-0-9]$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SemanticVersion`

The version of the Image Builder recipe specified by the `name` field.

_Required_: Yes

_Type_: String

_Pattern_: `^(?:[0-9]+|x)\.(?:[0-9]+|x)\.(?:[0-9]+|x)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PolicyDetail

ResourceSelection

All content copied from https://docs.aws.amazon.com/.
