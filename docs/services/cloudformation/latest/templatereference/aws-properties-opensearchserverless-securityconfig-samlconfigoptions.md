This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::SecurityConfig SamlConfigOptions

Describes SAML options for an OpenSearch Serverless security configuration in the form of a key-value
map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupAttribute" : String,
  "Metadata" : String,
  "OpenSearchServerlessEntityId" : String,
  "SessionTimeout" : Integer,
  "UserAttribute" : String
}

```

### YAML

```yaml

  GroupAttribute: String
  Metadata: String
  OpenSearchServerlessEntityId: String
  SessionTimeout: Integer
  UserAttribute: String

```

## Properties

`GroupAttribute`

The group attribute for this SAML integration.

_Required_: No

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metadata`

The XML IdP metadata file generated from your identity provider.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0009\u000A\u000D\u0020-\u007E\u00A1-\u00FF]+`

_Minimum_: `1`

_Maximum_: `51200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenSearchServerlessEntityId`

Custom entity ID attribute to override the default entity ID for this SAML
integration.

_Required_: No

_Type_: String

_Pattern_: `^aws:opensearch:[0-9]{12}:*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionTimeout`

The session timeout, in minutes. Default is 60 minutes (12 hours).

_Required_: No

_Type_: Integer

_Minimum_: `5`

_Maximum_: `720`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserAttribute`

A user attribute for this SAML integration.

_Required_: No

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IamIdentityCenterConfigOptions

AWS::OpenSearchServerless::SecurityPolicy
