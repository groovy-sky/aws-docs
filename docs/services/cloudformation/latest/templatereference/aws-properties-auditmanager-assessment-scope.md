This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AuditManager::Assessment Scope

The `Scope` property type specifies the wrapper that contains the AWS accounts and services that are in scope for the assessment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsAccounts" : [ AWSAccount, ... ],
  "AwsServices" : [ AWSService, ... ]
}

```

### YAML

```yaml

  AwsAccounts:
    - AWSAccount
  AwsServices:
    - AWSService

```

## Properties

`AwsAccounts`

The AWS accounts that are included in the scope of the assessment.

_Required_: No

_Type_: Array of [AWSAccount](aws-properties-auditmanager-assessment-awsaccount.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsServices`

The AWS services that are included in the scope of the assessment.

###### Important

This API parameter is no longer supported. If you use this parameter to specify one
or more AWS services, Audit Manager ignores this input. Instead, the
value for `awsServices` will show as empty.

_Required_: No

_Type_: Array of [AWSService](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-auditmanager-assessment-awsservice.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Scope](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_Scope.html) in the
_AWS Audit Manager API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Role

Tag
