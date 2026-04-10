This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AuditManager::Assessment AWSService

The `AWSService` property type specifies an AWS service
such as Amazon S3, AWS CloudTrail, and so on.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ServiceName" : String
}

```

### YAML

```yaml

  ServiceName: String

```

## Properties

`ServiceName`

The name of the AWS service.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-\s().]+$`

_Minimum_: `1`

_Maximum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [AWSService](../../../../reference/audit-manager/latest/apireference/api-awsservice.md)
in the _AWS Audit Manager API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWSAccount

Delegation

All content copied from https://docs.aws.amazon.com/.
