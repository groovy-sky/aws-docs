This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::CustomActionType ConfigurationProperties

The configuration properties for the custom action.

###### Note

You can refer to a name in the configuration properties of the custom action
within the URL templates by following the format of {Config:name}, as long as the
configuration property is both required and not secret. For more information, see
[Create a\
Custom Action for a Pipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Key" : Boolean,
  "Name" : String,
  "Queryable" : Boolean,
  "Required" : Boolean,
  "Secret" : Boolean,
  "Type" : String
}

```

### YAML

```yaml

  Description: String
  Key: Boolean
  Name: String
  Queryable: Boolean
  Required: Boolean
  Secret: Boolean
  Type: String

```

## Properties

`Description`

The description of the action configuration property that is displayed to
users.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `160`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Key`

Whether the configuration property is a key.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the action configuration property.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Queryable`

Indicates that the property is used with `PollForJobs`. When creating a
custom action, an action can have up to one queryable property. If it has one, that
property must be both required and not secret.

If you create a pipeline with a custom action type, and that custom action contains
a queryable property, the value for that configuration property is subject to other
restrictions. The value must be less than or equal to twenty (20) characters. The value
can contain only alphanumeric characters, underscores, and hyphens.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Required`

Whether the configuration property is a required value.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Secret`

Whether the configuration property is secret. Secrets are hidden from all calls
except for `GetJobDetails`, `GetThirdPartyJobDetails`,
`PollForJobs`, and `PollForThirdPartyJobs`.

When updating a pipeline, passing \* \* \* \* \* without changing any other values of
the action preserves the previous value of the secret.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of the configuration property.

_Required_: No

_Type_: String

_Allowed values_: `String | Number | Boolean`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ArtifactDetails

Settings
