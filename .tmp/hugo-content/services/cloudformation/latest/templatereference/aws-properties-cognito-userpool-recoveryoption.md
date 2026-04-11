This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool RecoveryOption

A recovery option for a user. The `AccountRecoverySettingType` data type is
an array of this object. Each `RecoveryOptionType` has a priority property
that determines whether it is a primary or secondary option.

For example, if `verified_email` has a priority of `1` and
`verified_phone_number` has a priority of `2`, your user pool
sends account-recovery messages to a verified email address but falls back to an SMS
message if the user has a verified phone number. The `admin_only` option
prevents self-service account recovery.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Priority" : Integer
}

```

### YAML

```yaml

  Name: String
  Priority: Integer

```

## Properties

`Name`

The recovery method that this object sets a recovery option for.

_Required_: No

_Type_: String

_Allowed values_: `verified_email | verified_phone_number | admin_only`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

Your priority preference for using the specified attribute in account recovery. The
highest priority is `1`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PreTokenGenerationConfig

SchemaAttribute

All content copied from https://docs.aws.amazon.com/.
