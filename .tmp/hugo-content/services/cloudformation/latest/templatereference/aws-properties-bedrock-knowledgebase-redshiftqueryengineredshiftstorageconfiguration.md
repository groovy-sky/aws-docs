This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase RedshiftQueryEngineRedshiftStorageConfiguration

Contains configurations for storage in Amazon Redshift.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseName" : String
}

```

### YAML

```yaml

  DatabaseName: String

```

## Properties

`DatabaseName`

The name of the Amazon Redshift database.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftQueryEngineConfiguration

RedshiftQueryEngineStorageConfiguration

All content copied from https://docs.aws.amazon.com/.
