This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource AthenaParameters

Parameters for Amazon Athena.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdentityCenterConfiguration" : IdentityCenterConfiguration,
  "RoleArn" : String,
  "WorkGroup" : String
}

```

### YAML

```yaml

  IdentityCenterConfiguration:
    IdentityCenterConfiguration
  RoleArn: String
  WorkGroup: String

```

## Properties

`IdentityCenterConfiguration`

An optional parameter that configures IAM Identity Center authentication to grant Quick Sight access to your workgroup.

This parameter can only be specified if your Quick Sight account is configured with IAM Identity Center.

_Required_: No

_Type_: [IdentityCenterConfiguration](aws-properties-quicksight-datasource-identitycenterconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

Use the `RoleArn` structure to override an account-wide role for a specific Athena data source. For example, say an account administrator has turned off all Athena access with an account-wide role. The administrator can then use `RoleArn` to bypass the account-wide role and allow Athena access for the single Athena data source that is specified in the structure, even if the account-wide role forbidding Athena access is still active.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkGroup`

The workgroup that Amazon Athena uses.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AmazonOpenSearchParameters

AuroraParameters

All content copied from https://docs.aws.amazon.com/.
