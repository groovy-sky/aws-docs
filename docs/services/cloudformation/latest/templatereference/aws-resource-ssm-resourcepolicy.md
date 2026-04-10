This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::ResourcePolicy

Creates or updates a Systems Manager resource policy. A resource policy helps you
to define the IAM entity (for example, an AWS account)
that can manage your Systems Manager resources. Currently, `OpsItemGroup`
is the only resource that supports Systems Manager resource policies. The resource
policy for `OpsItemGroup` enables AWS accounts to view and
interact with OpsCenter operational work items (OpsItems). OpsCenter is a tool in
Systems Manager.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSM::ResourcePolicy",
  "Properties" : {
      "Policy" : Json,
      "ResourceArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::SSM::ResourcePolicy
Properties:
  Policy: Json
  ResourceArn: String

```

## Properties

`Policy`

A policy you want to associate with a resource.

_Required_: Yes

_Type_: Json

_Pattern_: `^(?!\s*$).+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

The Amazon Resource Name (ARN) of the resource to which you want to attach a
policy.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`PolicyHash`

ID of the current policy version. The hash helps to prevent a situation where multiple
users attempt to overwrite a policy. You must provide this hash and the policy ID when
updating or deleting a policy.

`PolicyId`

ID of the current policy version.

## Examples

### Create a resource policy for OpsCenter

The following example specifies the management or delegated administrator
account IDs for working with OpsItems across accounts.

#### YAML

```yaml

---
AWSTemplateFormatVersion: '2010-09-09'
Description: Creates resources needed for a member account to work with OpsCenter OpsItems across multiple accounts.

Parameters:
  AdminAccountIds:
    Description: Allows one or more accounts to access OpsItems. Specify AWS Organizations management account IDs
                 and delegated administrator account IDs in a comma-separated list.
    Type: CommaDelimitedList
  ParentDeploymentRegion:
    Description: Primary AWS Region used for creating global resources such as IAM roles.
    Type: String

Conditions:
  IsParentDeploymentRegion:
    Fn::Equals:
    - !Ref 'AWS::Region'
    - !Ref ParentDeploymentRegion

Resources:
  OpsItemCrossAccountResourcePolicy:
    Type: AWS::SSM::ResourcePolicy
    Properties:
      Policy: !Sub
        - '{"Version":"2012-10-17","Statement":[{"Sid":"AllowAdminAccountsToAccessOpsItems2","Effect":"Allow","Principal":{"AWS":["${AdminAccountIdsString}"]},"Action":["ssm:CreateOpsItem","ssm:AddTagsToResource","ssm:GetOpsItem","ssm:UpdateOpsItem"],"Resource":["arn:${AWS::Partition}:ssm:${AWS::Region}:${AWS::AccountId}:opsitem/*","arn:${AWS::Partition}:ssm:${AWS::Region}:${AWS::AccountId}:opsitemgroup/default"]}]}'
        - AdminAccountIdsString:
            Fn::Join:
            - '\",\"'
            - !Ref AdminAccountIds
      ResourceArn:
        Fn::Sub: arn:${AWS::Partition}:ssm:${AWS::Region}:${AWS::AccountId}:opsitemgroup/default

  OpsItemCrossAccountExecutionRole:
    Type: AWS::IAM::Role
    Condition: IsParentDeploymentRegion
    Properties:
      RoleName: OpsItem-CrossAccountExecutionRole
      Description: 'Role used by the management account or delegated administrator to remediate OpsItems'
      AssumeRolePolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Principal:
              AWS: !Ref AdminAccountIds
            Condition:
              StringLike:
                "aws:PrincipalArn": !Split
                  - ','
                  - !Sub
                    - 'arn:*:iam::${inner}:role/OpsItem-*Role*'
                    - inner: !Join
                        - ':role/OpsItem-*Role*,arn:*:iam::'
                        - Ref: AdminAccountIds
            Action:
              - sts:AssumeRole
      Path: '/'
      ManagedPolicyArns:
        - !Sub 'arn:${AWS::Partition}:iam::aws:policy/ReadOnlyAccess'
```

## See also

- [Setting up OpsCenter to work with OpsItems across accounts](../../../systems-manager/latest/userguide/opscenter-getting-started-multiple-accounts.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SyncSource

Next

All content copied from https://docs.aws.amazon.com/.
