This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::DataSource RedshiftClusterStorage

The details of the Amazon Redshift cluster storage.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterName" : String
}

```

### YAML

```yaml

  ClusterName: String

```

## Properties

`ClusterName`

The name of an Amazon Redshift cluster.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-z].[a-z0-9\-]*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecommendationConfiguration

RedshiftCredentialConfiguration

All content copied from https://docs.aws.amazon.com/.
