This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase RedshiftQueryEngineAwsDataCatalogStorageConfiguration

Contains configurations for storage in AWS Glue Data Catalog.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TableNames" : [ String, ... ]
}

```

### YAML

```yaml

  TableNames:
    - String

```

## Properties

`TableNames`

A list of names of the tables to use.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftProvisionedConfiguration

RedshiftQueryEngineConfiguration

All content copied from https://docs.aws.amazon.com/.
