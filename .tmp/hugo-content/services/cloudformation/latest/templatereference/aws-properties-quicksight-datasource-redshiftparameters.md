This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource RedshiftParameters

The parameters for Amazon Redshift. The `ClusterId` field can be blank if
`Host` and `Port` are both set. The `Host` and `Port` fields can be blank if the `ClusterId` field is set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterId" : String,
  "Database" : String,
  "Host" : String,
  "IAMParameters" : RedshiftIAMParameters,
  "IdentityCenterConfiguration" : IdentityCenterConfiguration,
  "Port" : Number
}

```

### YAML

```yaml

  ClusterId: String
  Database: String
  Host: String
  IAMParameters:
    RedshiftIAMParameters
  IdentityCenterConfiguration:
    IdentityCenterConfiguration
  Port: Number

```

## Properties

`ClusterId`

Cluster ID. This field can be blank if the `Host` and `Port` are
provided.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Database`

Database.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Host`

Host. This field can be blank if `ClusterId` is provided.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IAMParameters`

An optional parameter that uses IAM authentication to grant Quick Sight access to your cluster. This parameter can be used instead of [DataSourceCredentials](../../../../reference/quicksight/latest/apireference/api-datasourcecredentials.md).

_Required_: No

_Type_: [RedshiftIAMParameters](aws-properties-quicksight-datasource-redshiftiamparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityCenterConfiguration`

An optional parameter that configures IAM Identity Center authentication to grant Quick Sight access to your cluster.

This parameter can only be specified if your Quick Sight account is configured with IAM Identity Center.

_Required_: No

_Type_: [IdentityCenterConfiguration](aws-properties-quicksight-datasource-identitycenterconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

Port. This field can be blank if the `ClusterId` is provided.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftIAMParameters

ResourcePermission

All content copied from https://docs.aws.amazon.com/.
