This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint MongoDbSettings

Provides information that defines a MongoDB endpoint. This
information includes the output format of records applied to the endpoint and details of
transaction and control table data information. For more information about other available settings, see
[Endpoint configuration settings when using MongoDB as a source for AWS DMS](../../../dms/latest/userguide/chap-source-mongodb.md#CHAP_Source.MongoDB.Configuration)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthMechanism" : String,
  "AuthSource" : String,
  "AuthType" : String,
  "DatabaseName" : String,
  "DocsToInvestigate" : String,
  "ExtractDocId" : String,
  "NestingLevel" : String,
  "Password" : String,
  "Port" : Integer,
  "SecretsManagerAccessRoleArn" : String,
  "SecretsManagerSecretId" : String,
  "ServerName" : String,
  "Username" : String
}

```

### YAML

```yaml

  AuthMechanism: String
  AuthSource: String
  AuthType: String
  DatabaseName: String
  DocsToInvestigate: String
  ExtractDocId: String
  NestingLevel: String
  Password: String
  Port: Integer
  SecretsManagerAccessRoleArn: String
  SecretsManagerSecretId: String
  ServerName: String
  Username: String

```

## Properties

`AuthMechanism`

The authentication mechanism you use to access the MongoDB source endpoint.

For the default value, in MongoDB version 2.x, `"default"` is
`"mongodb_cr"`. For MongoDB version 3.x or later, `"default"` is
`"scram_sha_1"`. This setting isn't used when `AuthType` is set to `"no"`.

_Required_: No

_Type_: String

_Allowed values_: `default | mongodb_cr | scram_sha_1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthSource`

The MongoDB database name. This setting isn't used when `AuthType` is set to `"no"`.

The default is `"admin"`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthType`

The authentication type you use to access the MongoDB source endpoint.

When set to `"no"`, user name and password parameters are not used and can be empty.

_Required_: No

_Type_: String

_Allowed values_: `no | password`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The database name on the MongoDB source endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocsToInvestigate`

Indicates the number of documents to preview to determine the document organization.
Use this setting when `NestingLevel` is set to `"one"`.

Must be a positive value greater than `0`. Default value is `1000`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExtractDocId`

Specifies the document ID. Use this setting when `NestingLevel` is set to
`"none"`. It's required when CDC with
[multi-document transactions](https://www.mongodb.com/docs/manual/reference/method/Session.startTransaction)
is used.

Default value is `"false"`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NestingLevel`

Specifies either document or table mode.

Default value is `"none"`. Specify `"none"` to use document mode.
Specify `"one"` to use table mode.

_Required_: No

_Type_: String

_Allowed values_: `none | one`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Password`

The password for the user account you use to access the MongoDB source endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port value for the MongoDB source endpoint.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerAccessRoleArn`

The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the
trusted entity and grants the required permissions to access the value in
`SecretsManagerSecret`. The role must allow the `iam:PassRole` action.
`SecretsManagerSecret` has the value of the AWS Secrets Manager secret
that allows access to the MongoDB endpoint.

###### Note

You can specify one of two sets of values for these permissions. You can specify the
values for this setting and `SecretsManagerSecretId`. Or you can specify
clear-text values for `UserName`, `Password`,
`ServerName`, and `Port`. You can't specify both.

For more information on creating this `SecretsManagerSecret`, the corresponding
`SecretsManagerAccessRoleArn`, and the `SecretsManagerSecretId`
that is required to access it, see
[Using secrets to access AWS Database Migration Service resources](../../../dms/latest/userguide/chap-security.md#security-iam-secretsmanager)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerSecretId`

The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the MongoDB endpoint connection details.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerName`

The name of the server on the MongoDB source endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The user name you use to access the MongoDB source endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MicrosoftSqlServerSettings

MySqlSettings

All content copied from https://docs.aws.amazon.com/.
