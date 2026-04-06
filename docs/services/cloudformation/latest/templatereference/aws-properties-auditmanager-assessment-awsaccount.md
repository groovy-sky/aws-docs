This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AuditManager::Assessment AWSAccount

The `AWSAccount` property type specifies the wrapper of the AWS account details, such as account ID, email address, and so on.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EmailAddress" : String,
  "Id" : String,
  "Name" : String
}

```

### YAML

```yaml

  EmailAddress: String
  Id: String
  Name: String

```

## Properties

`EmailAddress`

The email address that's associated with the AWS account.

_Required_: No

_Type_: String

_Pattern_: `^.*@.*$`

_Minimum_: `1`

_Maximum_: `320`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The identifier for the AWS account.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the AWS account.

_Required_: No

_Type_: String

_Pattern_: `^[\u0020-\u007E]+$`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [AWSAccount](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_AWSAccount.html)
in the _AWS Audit Manager API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AssessmentReportsDestination

AWSService
