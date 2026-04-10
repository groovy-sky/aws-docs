This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::IdentitySource CognitoGroupConfiguration

The type of entity that a policy store maps to groups from an Amazon Cognito user
pool identity source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupEntityType" : String
}

```

### YAML

```yaml

  GroupEntityType: String

```

## Properties

`GroupEntityType`

The name of the schema entity type that's mapped to the user pool group. Defaults
to `AWS::CognitoGroup`.

_Required_: Yes

_Type_: String

_Pattern_: `^([_a-zA-Z][_a-zA-Z0-9]*::)*[_a-zA-Z][_a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::VerifiedPermissions::IdentitySource

CognitoUserPoolConfiguration

All content copied from https://docs.aws.amazon.com/.
