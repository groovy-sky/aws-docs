This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project BatchRestrictions

Specifies restrictions for the batch build.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComputeTypesAllowed" : [ String, ... ],
  "MaximumBuildsAllowed" : Integer
}

```

### YAML

```yaml

  ComputeTypesAllowed:
    - String
  MaximumBuildsAllowed: Integer

```

## Properties

`ComputeTypesAllowed`

An array of strings that specify the compute types that are allowed for the batch
build. See [Build environment\
compute types](../../../codebuild/latest/userguide/build-env-ref-compute-types.md) in the _AWS CodeBuild User Guide_ for these values.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumBuildsAllowed`

Specifies the maximum number of builds allowed.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Artifacts

BuildStatusConfig

All content copied from https://docs.aws.amazon.com/.
