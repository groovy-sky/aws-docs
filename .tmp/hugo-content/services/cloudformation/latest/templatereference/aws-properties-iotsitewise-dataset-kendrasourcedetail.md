This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Dataset KendraSourceDetail

The source details for the Kendra dataset source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KnowledgeBaseArn" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  KnowledgeBaseArn: String
  RoleArn: String

```

## Properties

`KnowledgeBaseArn`

The `knowledgeBaseArn` details for the Kendra dataset source.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The `roleARN` details for the Kendra dataset source.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatasetSource

SourceDetail

All content copied from https://docs.aws.amazon.com/.
