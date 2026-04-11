This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::CustomActionType ArtifactDetails

Returns information about the details of an artifact.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaximumCount" : Integer,
  "MinimumCount" : Integer
}

```

### YAML

```yaml

  MaximumCount: Integer
  MinimumCount: Integer

```

## Properties

`MaximumCount`

The maximum number of artifacts allowed for the action type.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinimumCount`

The minimum number of artifacts allowed for the action type.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CodePipeline::CustomActionType

ConfigurationProperties

All content copied from https://docs.aws.amazon.com/.
