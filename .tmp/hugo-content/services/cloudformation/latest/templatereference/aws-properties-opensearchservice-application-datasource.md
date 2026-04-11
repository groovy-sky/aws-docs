This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Application DataSource

Data sources that are associated with an OpenSearch application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSourceArn" : String,
  "DataSourceDescription" : String
}

```

### YAML

```yaml

  DataSourceArn: String
  DataSourceDescription: String

```

## Properties

`DataSourceArn`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSourceDescription`

Detailed description of a data source.

_Required_: No

_Type_: String

_Pattern_: `^([a-zA-Z0-9_])*[\\a-zA-Z0-9_@#%*+=:?./!\s-]*$`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AppConfig

IamIdentityCenterOptions

All content copied from https://docs.aws.amazon.com/.
