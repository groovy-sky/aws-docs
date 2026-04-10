This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Connection OAuth2ClientApplication

The OAuth2 client app used for the connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AWSManagedClientApplicationReference" : String,
  "UserManagedClientApplicationClientId" : String
}

```

### YAML

```yaml

  AWSManagedClientApplicationReference: String
  UserManagedClientApplicationClientId: String

```

## Properties

`AWSManagedClientApplicationReference`

The reference to the SaaS-side client app that is AWS managed.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserManagedClientApplicationClientId`

The client application clientID if the ClientAppType is `USER_MANAGED`.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectionInput

OAuth2Credentials

All content copied from https://docs.aws.amazon.com/.
