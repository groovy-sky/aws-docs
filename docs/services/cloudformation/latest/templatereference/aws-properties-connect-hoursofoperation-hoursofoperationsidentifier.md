This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::HoursOfOperation HoursOfOperationsIdentifier

Identifier for a hours of operations resource: ARN, ID, Name

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "Name" : String
}

```

### YAML

```yaml

  Id: String
  Name: String

```

## Properties

`Id`

Unique identifier of the hours of operation.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of the hours of operation

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HoursOfOperationOverrideConfig

HoursOfOperationTimeSlice

All content copied from https://docs.aws.amazon.com/.
