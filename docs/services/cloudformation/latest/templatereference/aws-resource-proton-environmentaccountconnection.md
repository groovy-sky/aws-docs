This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Proton::EnvironmentAccountConnection

Detailed data of an AWS Proton environment account connection resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Proton::EnvironmentAccountConnection",
  "Properties" : {
      "CodebuildRoleArn" : String,
      "ComponentRoleArn" : String,
      "EnvironmentAccountId" : String,
      "EnvironmentName" : String,
      "ManagementAccountId" : String,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Proton::EnvironmentAccountConnection
Properties:
  CodebuildRoleArn: String
  ComponentRoleArn: String
  EnvironmentAccountId: String
  EnvironmentName: String
  ManagementAccountId: String
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`CodebuildRoleArn`

The Amazon Resource Name (ARN) of an IAM service role in the environment account. AWS Proton uses this role to provision infrastructure resources
using CodeBuild-based provisioning in the associated environment account.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov):iam::\d{12}:role/([\w+=,.@-]{1,512}[/:])*([\w+=,.@-]{1,64})$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComponentRoleArn`

The Amazon Resource Name (ARN) of the IAM service role that AWS Proton uses when provisioning directly defined components in the associated
environment account. It determines the scope of infrastructure that a component can provision in the account.

The environment account connection must have a `componentRoleArn` to allow directly defined components to be associated with any
environments running in the account.

For more information about components, see
[AWS Proton components](../../../proton/latest/userguide/ag-components.md) in the
_AWS Proton User Guide_.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov):iam::\d{12}:role/([\w+=,.@-]{1,512}[/:])*([\w+=,.@-]{1,64})$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentAccountId`

The environment account that's connected to the environment account connection.

_Required_: No

_Type_: String

_Pattern_: `^\d{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentName`

The name of the environment that's associated with the environment account connection.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z]+[0-9A-Za-z_\-]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagementAccountId`

The ID of the management account that's connected to the environment account connection.

_Required_: No

_Type_: String

_Pattern_: `^\d{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM service role that's associated with the environment account connection.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov):iam::\d{12}:role/([\w+=,.@-]{1,512}[/:])*([\w+=,.@-]{1,64})$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An optional list of metadata items that you can associate with the AWS Proton environment account connection. A tag is a key-value pair.

For more information, see [AWS Proton resources and tagging](../../../proton/latest/userguide/resources.md) in the
_AWS Proton User Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-proton-environmentaccountconnection-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the environment account connection.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the environment account connection ARN.

`Id`

Returns the environment account connection ID.

`Status`

Returns the environment account connection status.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Proton

Tag

All content copied from https://docs.aws.amazon.com/.
