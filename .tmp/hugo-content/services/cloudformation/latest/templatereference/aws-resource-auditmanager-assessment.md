This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AuditManager::Assessment

The `AWS::AuditManager::Assessment` resource is an Audit Manager
resource type that defines the scope of audit evidence collected by Audit Manager. An
Audit Manager assessment is an implementation of an Audit Manager framework.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AuditManager::Assessment",
  "Properties" : {
      "AssessmentReportsDestination" : AssessmentReportsDestination,
      "AwsAccount" : AWSAccount,
      "Delegations" : [ Delegation, ... ],
      "Description" : String,
      "FrameworkId" : String,
      "Name" : String,
      "Roles" : [ Role, ... ],
      "Scope" : Scope,
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AuditManager::Assessment
Properties:
  AssessmentReportsDestination:
    AssessmentReportsDestination
  AwsAccount:
    AWSAccount
  Delegations:
    - Delegation
  Description: String
  FrameworkId: String
  Name: String
  Roles:
    - Role
  Scope:
    Scope
  Status: String
  Tags:
    - Tag

```

## Properties

`AssessmentReportsDestination`

The destination that evidence reports are stored in for the assessment.

_Required_: No

_Type_: [AssessmentReportsDestination](aws-properties-auditmanager-assessment-assessmentreportsdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsAccount`

The AWS account that's associated with the assessment.

_Required_: No

_Type_: [AWSAccount](aws-properties-auditmanager-assessment-awsaccount.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Delegations`

The delegations that are associated with the assessment.

_Required_: No

_Type_: Array of [Delegation](aws-properties-auditmanager-assessment-delegation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the assessment.

_Required_: No

_Type_: String

_Pattern_: `^[\w\W\s\S]*$`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrameworkId`

The unique identifier for the framework.

_Required_: No

_Type_: String

_Pattern_: `^([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}|.*\S.*)$`

_Minimum_: `32`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the assessment.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_\.]+$`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Roles`

The roles that are associated with the assessment.

_Required_: No

_Type_: Array of [Role](aws-properties-auditmanager-assessment-role.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

The wrapper of AWS accounts and services that are in scope for the
assessment.

_Required_: No

_Type_: [Scope](aws-properties-auditmanager-assessment-scope.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The overall status of the assessment.

When you create a new assessment, the initial `Status` value is always
`ACTIVE`. When you create an assessment, even if you specify the value as
`INACTIVE`, the value overrides to `ACTIVE`.

After you create an assessment, you can change the value of the `Status`
property at any time. For example, when you want to stop collecting evidence for your
assessment, you can change the assessment status to `INACTIVE`.

_Required_: No

_Type_: String

_Allowed values_: `ACTIVE | INACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags that are associated with the assessment.

_Required_: No

_Type_: Array of [Tag](aws-properties-auditmanager-assessment-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the assessment ID. For example:

`{ "Ref": "111A1A1A-22B2-33C3-DDD4-55E5E5E555E5" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the assessment.

`AssessmentId`

The unique identifier for the assessment.

`CreationTime`

Specifies when the assessment was created.

## See also

- [CreateAssessment](../../../../reference/audit-manager/latest/apireference/api-createassessment.md) in the _AWS Audit Manager API_
_Reference_.

- [DeleteAssessment](../../../../reference/audit-manager/latest/apireference/api-deleteassessment.md) in the _AWS Audit Manager API_
_Reference_.

- [GetAssessment](../../../../reference/audit-manager/latest/apireference/api-getassessment.md) in the _AWS Audit Manager API_
_Reference_.

- [UpdateAssessment](../../../../reference/audit-manager/latest/apireference/api-updateassessment.md) in the _AWS Audit Manager API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Audit Manager

AssessmentReportsDestination

All content copied from https://docs.aws.amazon.com/.
