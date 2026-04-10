This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::DataSource RedshiftServerlessStorage

The details of the Amazon Redshift Serverless workgroup storage.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "WorkgroupName" : String
}

```

### YAML

```yaml

  WorkgroupName: String

```

## Properties

`WorkgroupName`

The name of the Amazon Redshift Serverless workgroup.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9-]+$`

_Minimum_: `3`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftRunConfigurationInput

RedshiftStorage

All content copied from https://docs.aws.amazon.com/.
