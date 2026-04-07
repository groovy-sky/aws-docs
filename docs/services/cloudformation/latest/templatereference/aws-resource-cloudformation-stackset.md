This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::StackSet

The `AWS::CloudFormation::StackSet` resource contains information about a
StackSet. With StackSets, you can provision stacks across AWS accounts
and Regions from a single CloudFormation template. Each stack is based on the
same CloudFormation template, but you can customize individual stacks using
parameters.

###### Important

Run deployments to nested StackSets from the parent stack, not directly through
the StackSet API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::StackSet",
  "Properties" : {
      "AdministrationRoleARN" : String,
      "AutoDeployment" : AutoDeployment,
      "CallAs" : String,
      "Capabilities" : [ String, ... ],
      "Description" : String,
      "ExecutionRoleName" : String,
      "ManagedExecution" : ManagedExecution,
      "OperationPreferences" : OperationPreferences,
      "Parameters" : [ Parameter, ... ],
      "PermissionModel" : String,
      "StackInstancesGroup" : [ StackInstances, ... ],
      "StackSetName" : String,
      "Tags" : [ Tag, ... ],
      "TemplateBody" : String,
      "TemplateURL" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::StackSet
Properties:
  AdministrationRoleARN: String
  AutoDeployment:
    AutoDeployment
  CallAs: String
  Capabilities:
    - String
  Description: String
  ExecutionRoleName: String
  ManagedExecution:
    ManagedExecution
  OperationPreferences:
    OperationPreferences
  Parameters:
    - Parameter
  PermissionModel: String
  StackInstancesGroup:
    - StackInstances
  StackSetName: String
  Tags:
    - Tag
  TemplateBody: String
  TemplateURL: String

```

## Properties

`AdministrationRoleARN`

The Amazon Resource Number (ARN) of the IAM role to use to create this
StackSet. Specify an IAM role only if you are using customized
administrator roles to control which users or groups can manage specific StackSets
within the same administrator account.

Use customized administrator roles to control which users or groups can manage
specific StackSets within the same administrator account. For more information, see
[Grant\
self-managed permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html) in the _CloudFormation User_
_Guide_.

Valid only if the permissions model is `SELF_MANAGED`.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoDeployment`

Describes whether StackSets automatically deploys to AWS Organizations accounts
that are added to a target organization or organizational unit (OU). For more
information, see [Enable or disable automatic deployments for StackSets in AWS Organizations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-manage-auto-deployment.html) in the _CloudFormation User Guide_.

Required if the permissions model is `SERVICE_MANAGED`. (Not used with
self-managed permissions.)

_Required_: Conditional

_Type_: [AutoDeployment](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-stackset-autodeployment.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CallAs`

Specifies whether you are acting as an account administrator in the organization's
management account or as a delegated administrator in a member account.

By default, `SELF` is specified. Use `SELF` for StackSets with
self-managed permissions.

- To create a StackSet with service-managed permissions while signed in to the
management account, specify `SELF`.

- To create a StackSet with service-managed permissions while signed in to a
delegated administrator account, specify `DELEGATED_ADMIN`.

Your AWS account must be registered as a delegated admin in the
management account. For more information, see [Register a delegated administrator](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html) in the _CloudFormation User Guide_.

StackSets with service-managed permissions are created in the management account,
including StackSets that are created by delegated administrators.

Valid only if the permissions model is `SERVICE_MANAGED`.

_Required_: No

_Type_: String

_Allowed values_: `SELF | DELEGATED_ADMIN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Capabilities`

The capabilities that are allowed in the StackSet. Some StackSet templates might
include resources that can affect permissions in your AWS account—for
example, by creating new IAM users. For more information, see [Acknowledging IAM resources in CloudFormation\
templates](../userguide/control-access-with-iam.md#using-iam-capabilities) in the _CloudFormation User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the StackSet.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRoleName`

The name of the IAM execution role to use to create the StackSet. If
you don't specify an execution role, CloudFormation uses the
`AWSCloudFormationStackSetExecutionRole` role for the StackSet
operation.

Valid only if the permissions model is `SELF_MANAGED`.

_Pattern_: `[a-zA-Z_0-9+=,.@-]+`

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedExecution`

Describes whether StackSets performs non-conflicting operations concurrently and
queues conflicting operations.

When active, StackSets performs non-conflicting operations concurrently and queues
conflicting operations. After conflicting operations finish, StackSets starts queued
operations in request order.

###### Note

If there are already running or queued operations, StackSets queues all incoming
operations even if they are non-conflicting.

You can't modify your StackSet's execution configuration while there are running
or queued operations for that StackSet.

When inactive (default), StackSets performs one operation at a time in request
order.

_Required_: No

_Type_: [ManagedExecution](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-stackset-managedexecution.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperationPreferences`

The user-specified preferences for how CloudFormation performs a StackSet
operation.

_Required_: No

_Type_: [OperationPreferences](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-stackset-operationpreferences.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

The input parameters for the StackSet template.

_Required_: No

_Type_: Array of [Parameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-stackset-parameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PermissionModel`

Describes how the IAM roles required for StackSet operations are
created.

- With `SELF_MANAGED` permissions, you must create the administrator
and execution roles required to deploy to target accounts. For more information,
see [Grant self-managed permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html) in the _CloudFormation_
_User Guide_.

- With `SERVICE_MANAGED` permissions, StackSets automatically creates
the IAM roles required to deploy to accounts managed by AWS Organizations. For more information, see [Activate trusted access for StackSets with AWS Organizations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-activate-trusted-access.html) in
the _CloudFormation User Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `SERVICE_MANAGED | SELF_MANAGED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StackInstancesGroup`

A group of stack instances with parameters in some specific accounts and
Regions.

_Required_: No

_Type_: Array of [StackInstances](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-stackset-stackinstances.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackSetName`

The name to associate with the StackSet. The name must be unique in the Region where
you create your StackSet.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9\-]{0,127}$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Key-value pairs to associate with this stack. CloudFormation also propagates these tags to
supported resources in the stack. You can specify a maximum number of 50 tags.

If you don't specify this parameter, CloudFormation doesn't modify the stack's tags. If you
specify an empty value, CloudFormation removes all associated tags.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-stackset-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateBody`

The structure that contains the template body, with a minimum length of 1 byte and a
maximum length of 51,200 bytes.

You must include either `TemplateURL` or `TemplateBody` in a
StackSet, but you can't use both. Dynamic references in the `TemplateBody`
may not work correctly in all cases. It's recommended to pass templates that contain
dynamic references through `TemplateUrl` instead.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `51200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateURL`

The URL of a file that contains the template body. The URL must point to a template
(max size: 1 MB) that's located in an Amazon S3 bucket or a Systems Manager
document. The location for an Amazon S3 bucket must start with
`https://`.

Conditional: You must specify only one of the following parameters:
`TemplateBody`, `TemplateURL`.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `5120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the `StackSetId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`StackSetId`

Returns the unique identifier of the resource.

## Examples

- [Activate managed execution for your StackSet](#aws-resource-cloudformation-stackset--examples--Activate_managed_execution_for_your_StackSet)

- [Specifying Secrets Manager secrets in CloudFormation](#aws-resource-cloudformation-stackset--examples--Specifying_secrets_in)

### Activate managed execution for your StackSet

The following example creates a StackSet and specifies
`ManagedExecution`. With managed execution activated, StackSets
performs non-conflicting operations concurrently and queues conflicting
operations.

#### JSON

```json

{
    "TestStackSet1": {
        "Type": "AWS::CloudFormation::StackSet",
        "DeletionPolicy": "Retain",
        "Properties": {
            "StackSetName": "TestStackSet12345",
            "Description": "Updatedescription1",
            "PermissionModel": "SELF_MANAGED",
            "ManagedExecution": {
                "Active": true
            },
            "Tags": [
                {
                    "Key": "tag1",
                    "Value": "value1"
                }
            ],
            "TemplateBody": "{\n  \"AWSTemplateFormatVersion\": \"2010-09-09\",\n  \"Resources\": {\n    \"testWaitHandle\": {\n      \"Type\": \"AWS::CloudFormation::WaitConditionHandle\"\n    }\n  }\n}\n"
        }
    }
}
```

#### YAML

```yaml

TestStackSet1:
  Type: AWS::CloudFormation::StackSet
  DeletionPolicy: Retain
  Properties:
    StackSetName: TestStackSet12345
    Description: Updatedescription1
    PermissionModel: SELF_MANAGED
    ManagedExecution:
      Active: true
    Tags:
      - Key: tag1
        Value: value1
    TemplateBody: |
      {
        "AWSTemplateFormatVersion": "2010-09-09",
        "Resources": {
          "testWaitHandle": {
            "Type": "AWS::CloudFormation::WaitConditionHandle"
          }
        }
      }
```

### Specifying Secrets Manager secrets in CloudFormation

When using the `TemplateBody` property, if the template intends to
resolve secrets from Secrets Manager secret's through an `ARN` and
`!Join` is used to construct Secrets Manager's dynamic
reference, secrets resolution needs to be avoided at stack level so that it will
only be performed upon stack instance creation.

In the following example, a secrets resolution is avoided at the stack level
by providing `{{` and `resolve:secretsmanager:` as
separate strings to !Join instead of `{{resolve:secretsmanager:`
being provided as a single string:

#### JSON

```json

{
    "Fn::Join": [
        "",
        [
            "{{",
            "resolve:secretsmanager:",
            {
                "Fn::Sub": "arn:aws:secretsmanager:${AWS::Region}:${AWS::AccountId}:secret:my-secret"
            },
            "::my-secret-key::}}"
        ]
    ]
}
```

#### YAML

```yaml

!Join
- ''
- - '{{'
  - 'resolve:secretsmanager:'
  - !Sub 'arn:aws:secretsmanager:${AWS::Region}:${AWS::AccountId}:secret:my-secret'
  - '::my-secret-key::}}'
```

## See also

- [CloudFormation StackSets sample templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-sampletemplates.html) in the
_CloudFormation User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AutoDeployment
