This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline ActionTypeId

Represents information about an action type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Category" : String,
  "Owner" : String,
  "Provider" : String,
  "Version" : String
}

```

### YAML

```yaml

  Category: String
  Owner: String
  Provider: String
  Version: String

```

## Properties

`Category`

A category defines what kind of action can be taken in the stage, and constrains the
provider type for the action. Valid categories are limited to one of the values
below.

- `Source`

- `Build`

- `Test`

- `Deploy`

- `Invoke`

- `Approval`

- `Compute`

_Required_: Yes

_Type_: String

_Allowed values_: `Source | Build | Test | Deploy | Invoke | Approval | Compute`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Owner`

The creator of the action being called. There are three valid values for the
`Owner` field in the action category section within your pipeline
structure: `AWS`, `ThirdParty`, and `Custom`. For more
information, see [Valid Action Types and Providers in CodePipeline](../../../codepipeline/latest/userguide/reference-pipeline-structure.md#actions-valid-providers).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Provider`

The provider of the service being called by the action. Valid providers are
determined by the action category. For example, an action in the Deploy category type
might have a provider of CodeDeploy, which would be specified as
`CodeDeploy`. For more information, see [Valid Action Types and Providers in CodePipeline](../../../codepipeline/latest/userguide/reference-pipeline-structure.md#actions-valid-providers).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

A string that describes the action version.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionDeclaration

ArtifactStore

All content copied from https://docs.aws.amazon.com/.
