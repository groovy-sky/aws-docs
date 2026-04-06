This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AuditManager::Assessment AssessmentReportsDestination

The `AssessmentReportsDestination` property type specifies the location in
which AWS Audit Manager saves assessment reports for the given assessment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : String,
  "DestinationType" : String
}

```

### YAML

```yaml

  Destination: String
  DestinationType: String

```

## Properties

`Destination`

The destination bucket where Audit Manager stores assessment reports.

_Required_: No

_Type_: String

_Pattern_: `^(S|s)3:\/\/[a-zA-Z0-9\-\.\(\)\'\*\_\!\=\+\@\:\s\,\?\/]+$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationType`

The destination type, such as Amazon S3.

_Required_: No

_Type_: String

_Allowed values_: `S3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [AssessmentReportsDestination](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_AssessmentReportsDestination.html) in the _AWS Audit Manager_
_API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AuditManager::Assessment

AWSAccount
