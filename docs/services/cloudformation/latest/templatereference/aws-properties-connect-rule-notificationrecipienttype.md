This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Rule NotificationRecipientType

The type of notification recipient.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "UserArns" : [ String, ... ],
  "UserTags" : String
}

```

### YAML

```yaml

  UserArns:
    - String
  UserTags: String

```

## Properties

`UserArns`

The Amazon Resource Name (ARN) of the user account.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserTags`

The tags used to organize, track, or control access for this resource. For example, {
"tags": {"key1":"value1", "key2":"value2"} }. Amazon Connect users with the
specified tags will be notified.

_Required_: No

_Type_: String

_Pattern_: `^(?=.{1,128}$).+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldValue

Reference

All content copied from https://docs.aws.amazon.com/.
