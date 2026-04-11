This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline ArtifactStoreMap

A mapping of `artifactStore` objects and their corresponding AWS Regions. There must be an artifact store for the pipeline Region and for
each cross-region action in the pipeline.

###### Note

You must include either `artifactStore` or
`artifactStores` in your pipeline, but you cannot use both. If you
create a cross-region action in your pipeline, you must use
`artifactStores`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArtifactStore" : ArtifactStore,
  "Region" : String
}

```

### YAML

```yaml

  ArtifactStore:
    ArtifactStore
  Region: String

```

## Properties

`ArtifactStore`

Represents information about the S3 bucket where artifacts are stored for the
pipeline.

###### Note

You must include either `artifactStore` or
`artifactStores` in your pipeline, but you cannot use both. If you
create a cross-region action in your pipeline, you must use
`artifactStores`.

_Required_: Conditional

_Type_: [ArtifactStore](aws-properties-codepipeline-pipeline-artifactstore.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The action declaration's AWS Region, such as us-east-1.

_Required_: Yes

_Type_: String

_Minimum_: `4`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArtifactStore

BeforeEntryConditions

All content copied from https://docs.aws.amazon.com/.
