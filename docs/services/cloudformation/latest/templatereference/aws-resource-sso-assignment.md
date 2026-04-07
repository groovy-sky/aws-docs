This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSO::Assignment

Assigns access to a Principal for a specified AWS account using a
specified permission set.

###### Note

The term _principal_ here refers to a user or group that is
defined in IAM Identity Center.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSO::Assignment",
  "Properties" : {
      "InstanceArn" : String,
      "PermissionSetArn" : String,
      "PrincipalId" : String,
      "PrincipalType" : String,
      "TargetId" : String,
      "TargetType" : String
    }
}

```

### YAML

```yaml

Type: AWS::SSO::Assignment
Properties:
  InstanceArn: String
  PermissionSetArn: String
  PrincipalId: String
  PrincipalType: String
  TargetId: String
  TargetType: String

```

## Properties

`InstanceArn`

The ARN of the IAM Identity Center instance under which the operation will be executed.
For more information about ARNs, see [Amazon Resource Names (ARNs) and\
AWS Service Namespaces](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the _AWS General Reference_.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PermissionSetArn`

The ARN of the permission set.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::permissionSet/(sso)?ins-[a-zA-Z0-9-.]{16}/ps-[a-zA-Z0-9-./]{16}`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrincipalId`

An identifier for an object in IAM Identity Center, such as a user or group. PrincipalIds are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6). For more information about PrincipalIds in IAM Identity Center, see the [IAM Identity Center Identity Store API Reference](https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/welcome.html).

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$`

_Minimum_: `1`

_Maximum_: `47`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrincipalType`

The entity type for which the assignment will be created.

_Required_: Yes

_Type_: String

_Allowed values_: `USER | GROUP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetId`

TargetID is an AWS account identifier, (For example, 123456789012).

_Required_: Yes

_Type_: String

_Pattern_: `\d{12}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetType`

The entity type for which the assignment will be created.

_Required_: Yes

_Type_: String

_Allowed values_: `AWS_ACCOUNT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a generated ID, combined by all fields with the delimiter
`|`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Creating a new assignment for IAM Identity Center

The following example creates a custom assignment, assigning the user
`"user_id"` access to account
`"arn:aws:organizations::org_master_id:account/org_id/accountId"` with
the permissions `"PermissionSet"`.

#### JSON

```json

{
   "Assignment": {
      "Type": "AWS::SSO::Assignment",
      "Properties": {
         "InstanceArn": "arn:aws:sso:::instance/ssoins-instanceId",
         "PermissionSetArn": {
            "Fn::GetAtt": [
               "PermissionSet",
               "PermissionSetArn"
            ]
         },
         "TargetId": "accountId",
         "TargetType": "AWS_ACCOUNT",
         "PrincipalType": "USER",
         "PrincipalId": "user_id"
      }
   }
}
```

#### YAML

```yaml

Assignment:
    Type: AWS::SSO::Assignment
    Properties:
      InstanceArn: 'arn:aws:sso:::instance/ssoins-instanceId'
      PermissionSetArn: !GetAtt PermissionSet.PermissionSetArn
      TargetId: 'accountId'
      TargetType: 'AWS_ACCOUNT'
      PrincipalType: 'USER'
      PrincipalId: 'user_id'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SSO::ApplicationAssignment

AWS::SSO::Instance
