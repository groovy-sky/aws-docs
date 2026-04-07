This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::SecurityConfig IamFederationConfigOptions

Describes IAM federation options for an OpenSearch Serverless security configuration
in the form of a key-value map. These options define how OpenSearch Serverless
integrates with external identity providers using federation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupAttribute" : String,
  "UserAttribute" : String
}

```

### YAML

```yaml

  GroupAttribute: String
  UserAttribute: String

```

## Properties

`GroupAttribute`

The group attribute for this IAM federation integration. This attribute is used to map
identity provider groups to OpenSearch Serverless permissions.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z][A-Za-z0-9_.:/=+\-@]*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserAttribute`

The user attribute for this IAM federation integration. This attribute is used to
identify users in the federated authentication process.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z][A-Za-z0-9_.:/=+\-@]*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::OpenSearchServerless::SecurityConfig

IamIdentityCenterConfigOptions
