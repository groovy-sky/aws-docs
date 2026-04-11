This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::Workflow SourceReference

Contains information about the source reference in a code repository, such as a branch, tag, or commit.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "type" : String,
  "value" : String
}

```

### YAML

```yaml

  type: String
  value: String

```

## Properties

`type`

The type of source reference, such as branch, tag, or commit.

_Required_: No

_Type_: String

_Allowed values_: `BRANCH | TAG | COMMIT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`value`

The value of the source reference, such as the branch name, tag name, or commit ID.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RegistryMapping

WorkflowParameter

All content copied from https://docs.aws.amazon.com/.
