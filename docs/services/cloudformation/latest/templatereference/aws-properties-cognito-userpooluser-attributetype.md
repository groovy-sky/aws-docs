This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolUser AttributeType

The name and value of a user attribute.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

The name of the attribute, for example `email` or
`custom:department`.

In some older user pools, the regex pattern for acceptable values of this parameter is
`[\p{L}\p{M}\p{S}\p{N}\p{P}]+`. Older pools will eventually be updated to
use the new pattern. Affected user pools are those created before May 2024 in
US East (N. Virginia), US East (Ohio), US West (N. California), US West (Oregon),
Asia Pacific (Mumbai), Asia Pacific (Tokyo), Asia Pacific (Seoul),
Asia Pacific (Singapore), Asia Pacific (Sydney), Canada (Central),
Europe (Frankfurt), Europe (Ireland), Europe (London), Europe (Paris),
Europe (Stockholm), Middle East (Bahrain), and South America (São Paulo).

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}\t\n\r ]+`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The value of the attribute.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Cognito::UserPoolUser

AWS::Cognito::UserPoolUserToGroupAttachment
