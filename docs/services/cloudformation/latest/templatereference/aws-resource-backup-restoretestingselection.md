This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::RestoreTestingSelection

This request can be sent after CreateRestoreTestingPlan request
returns successfully. This is the second part of creating a resource testing
plan, and it must be completed sequentially.

This consists of `RestoreTestingSelectionName`,
`ProtectedResourceType`, and one of the following:

- `ProtectedResourceArns`

- `ProtectedResourceConditions`

Each protected resource type can have one single value.

A restore testing selection can include a wildcard value ("\*") for
`ProtectedResourceArns` along with `ProtectedResourceConditions`.
Alternatively, you can include up to 30 specific protected resource ARNs in
`ProtectedResourceArns`.

Cannot select by both protected resource types AND specific ARNs.
Request will fail if both are included.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Backup::RestoreTestingSelection",
  "Properties" : {
      "IamRoleArn" : String,
      "ProtectedResourceArns" : [ String, ... ],
      "ProtectedResourceConditions" : ProtectedResourceConditions,
      "ProtectedResourceType" : String,
      "RestoreMetadataOverrides" : {Key: Value, ...},
      "RestoreTestingPlanName" : String,
      "RestoreTestingSelectionName" : String,
      "ValidationWindowHours" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::Backup::RestoreTestingSelection
Properties:
  IamRoleArn: String
  ProtectedResourceArns:
    - String
  ProtectedResourceConditions:
    ProtectedResourceConditions
  ProtectedResourceType: String
  RestoreMetadataOverrides:
    Key: Value
  RestoreTestingPlanName: String
  RestoreTestingSelectionName: String
  ValidationWindowHours: Integer

```

## Properties

`IamRoleArn`

The Amazon Resource Name (ARN) of the IAM role that
AWS Backup uses to create the target resource; for
example: `arn:aws:iam::123456789012:role/S3Access`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtectedResourceArns`

You can include specific ARNs, such as
`ProtectedResourceArns: ["arn:aws:...", "arn:aws:..."]`
or you can include a wildcard: `ProtectedResourceArns: ["*"]`,
but not both.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtectedResourceConditions`

In a resource testing selection, this parameter filters by
specific conditions such as `StringEquals` or
`StringNotEquals`.

_Required_: No

_Type_: [ProtectedResourceConditions](aws-properties-backup-restoretestingselection-protectedresourceconditions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtectedResourceType`

The type of AWS resource included in a resource
testing selection;
for example, an Amazon EBS volume
or an Amazon RDS database.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RestoreMetadataOverrides`

You can override certain restore metadata keys by including the parameter
`RestoreMetadataOverrides` in the body of
`RestoreTestingSelection`. Key values are not case sensitive.

See the complete list of [restore testing\
inferred metadata](../../../aws-backup/latest/devguide/restore-testing-inferred-metadata.md).

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestoreTestingPlanName`

Unique string that is the name of the restore testing plan.

The name cannot be changed after creation. The name must
consist of only alphanumeric characters and underscores.
Maximum length is 50.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RestoreTestingSelectionName`

The unique name of the restore testing selection that
belongs to the related restore testing plan.

The name consists of only alphanumeric characters and underscores.
Maximum length is 50.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidationWindowHours`

This is amount of hours (1 to 168) available to run a validation script on the data. The
data will be deleted upon the completion of the validation script or the end of the
specified retention period, whichever comes first.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

KeyValue

All content copied from https://docs.aws.amazon.com/.
