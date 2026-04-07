This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RedshiftServerless::Namespace

A collection of database objects and users.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RedshiftServerless::Namespace",
  "Properties" : {
      "AdminPasswordSecretKmsKeyId" : String,
      "AdminUsername" : String,
      "AdminUserPassword" : String,
      "DbName" : String,
      "DefaultIamRoleArn" : String,
      "FinalSnapshotName" : String,
      "FinalSnapshotRetentionPeriod" : Integer,
      "IamRoles" : [ String, ... ],
      "KmsKeyId" : String,
      "LogExports" : [ String, ... ],
      "ManageAdminPassword" : Boolean,
      "NamespaceName" : String,
      "NamespaceResourcePolicy" : Json,
      "RedshiftIdcApplicationArn" : String,
      "SnapshotCopyConfigurations" : [ SnapshotCopyConfiguration, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RedshiftServerless::Namespace
Properties:
  AdminPasswordSecretKmsKeyId: String
  AdminUsername: String
  AdminUserPassword: String
  DbName: String
  DefaultIamRoleArn: String
  FinalSnapshotName: String
  FinalSnapshotRetentionPeriod: Integer
  IamRoles:
    - String
  KmsKeyId: String
  LogExports:
    - String
  ManageAdminPassword: Boolean
  NamespaceName: String
  NamespaceResourcePolicy: Json
  RedshiftIdcApplicationArn: String
  SnapshotCopyConfigurations:
    - SnapshotCopyConfiguration
  Tags:
    - Tag

```

## Properties

`AdminPasswordSecretKmsKeyId`

The ID of the AWS Key Management Service (KMS) key used to encrypt and store the namespace's admin credentials secret. You can only use this parameter if `ManageAdminPassword` is `true`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdminUsername`

The username of the administrator for the primary database created in the namespace.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z][a-zA-Z_0-9+.@-]*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdminUserPassword`

The password of the administrator for the primary database created in the namespace.

_Required_: No

_Type_: String

_Pattern_: `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^\x00-\x20\x22\x27\x2f\x40\x5c\x7f-\uffff]+`

_Minimum_: `8`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DbName`

The name of the primary database created in the namespace.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z][a-zA-Z_0-9+.@-]*`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultIamRoleArn`

The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FinalSnapshotName`

The name of the snapshot to be created before the namespace is deleted.

_Required_: No

_Type_: String

_Pattern_: `[a-z][a-z0-9]*(-[a-z0-9]+)*`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FinalSnapshotRetentionPeriod`

How long to retain the final snapshot.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamRoles`

A list of IAM roles to associate with the namespace.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The ID of the AWS Key Management Service key used to encrypt your data.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogExports`

The types of logs the namespace can export.
Available export types are `userlog`, `connectionlog`, and `useractivitylog`.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManageAdminPassword`

If true, Amazon Redshift uses AWS Secrets Manager to manage the namespace's admin credentials. You can't use `AdminUserPassword` if `ManageAdminPassword` is true. If `ManageAdminPassword` is `false` or not set, Amazon Redshift uses `AdminUserPassword` for the admin user account's password.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NamespaceName`

The name of the namespace.
Must be between 3-64 alphanumeric characters in lowercase,
and it cannot be a reserved word. A list of reserved words can be found
in [Reserved Words](https://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html) in the Amazon Redshift Database Developer Guide.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9-]+$`

_Minimum_: `3`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NamespaceResourcePolicy`

The resource policy that will be attached to the namespace.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedshiftIdcApplicationArn`

The ARN for the Redshift application that integrates with IAM Identity Center.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotCopyConfigurations`

Property description not available.

_Required_: No

_Type_: Array of [SnapshotCopyConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-redshiftserverless-namespace-snapshotcopyconfiguration.html)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The map of the key-value pairs used to tag the namespace.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-redshiftserverless-namespace-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, Ref
returns the NamespaceName, such as `sample-namespace.` For more
information about using the Ref function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

GetAtt returns a value for a specified attribute of this type. For more information, see
[Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html). The following are the available attributes and sample return
values.

`Namespace.AdminUsername`

The username of the administrator for the first database created in the namespace.

`Namespace.CreationDate`

The date of when the namespace was created.

`Namespace.DbName`

The name of the first database created in the namespace.

`Namespace.DefaultIamRoleArn`

The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace.

`Namespace.IamRoles`

A list of IAM roles to associate with the namespace.

`Namespace.KmsKeyId`

The ID of the AWS Key Management Service key used to encrypt your data.

`Namespace.LogExports`

The types of logs the namespace can export. Available export types are
`User log`, `Connection log`, and `User activity log`.

`Namespace.NamespaceArn`

The Amazon Resource Name (ARN) associated with a namespace.

`Namespace.NamespaceId`

The unique identifier of a namespace.

`Namespace.NamespaceName`

The name of the namespace.
Must be between 3-64 alphanumeric characters in lowercase,
and it cannot be a reserved word. A list of reserved words can be found
in [Reserved Words](https://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html) in the Amazon Redshift Database Developer Guide.

`Namespace.Status`

The status of the namespace.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Redshift Serverless

Namespace
