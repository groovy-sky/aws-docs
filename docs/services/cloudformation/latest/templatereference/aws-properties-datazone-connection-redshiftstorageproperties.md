This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection RedshiftStorageProperties

The Amazon Redshift storage properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterName" : String,
  "WorkgroupName" : String
}

```

### YAML

```yaml

  ClusterName: String
  WorkgroupName: String

```

## Properties

`ClusterName`

The cluster name in the Amazon Redshift storage properties.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9-]+$`

_Minimum_: `0`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkgroupName`

The workgroup name in the Amazon Redshift storage properties.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9-]+$`

_Minimum_: `3`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftPropertiesInput

S3PropertiesInput

All content copied from https://docs.aws.amazon.com/.
