This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::IdentitySource CognitoUserPoolConfiguration

A structure that contains configuration information used when creating or updating an
identity source that represents a connection to an Amazon Cognito user pool used as
an identity provider for Verified Permissions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientIds" : [ String, ... ],
  "GroupConfiguration" : CognitoGroupConfiguration,
  "UserPoolArn" : String
}

```

### YAML

```yaml

  ClientIds:
    - String
  GroupConfiguration:
    CognitoGroupConfiguration
  UserPoolArn: String

```

## Properties

`ClientIds`

The unique application client IDs that are associated with the specified Amazon Cognito user
pool.

Example: `"ClientIds": ["&ExampleCogClientId;"]`

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `255 | 1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupConfiguration`

The type of entity that a policy store maps to groups from an Amazon Cognito user
pool identity source.

_Required_: No

_Type_: [CognitoGroupConfiguration](aws-properties-verifiedpermissions-identitysource-cognitogroupconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolArn`

The [Amazon Resource Name (ARN)](../../../../general/latest/gr/aws-arns-and-namespaces.md) of the Amazon Cognito user pool that
contains the identities to be authorized.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-zA-Z0-9-]+:cognito-idp:(([a-zA-Z0-9-]+:\d{12}:userpool/[\w-]+_[0-9a-zA-Z]+))$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CognitoGroupConfiguration

IdentitySourceConfiguration

All content copied from https://docs.aws.amazon.com/.
