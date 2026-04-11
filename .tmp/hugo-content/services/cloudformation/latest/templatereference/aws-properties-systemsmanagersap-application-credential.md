This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SystemsManagerSAP::Application Credential

The credentials of your SAP application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CredentialType" : String,
  "DatabaseName" : String,
  "SecretId" : String
}

```

### YAML

```yaml

  CredentialType: String
  DatabaseName: String
  SecretId: String

```

## Properties

`CredentialType`

The type of the application credentials.

_Required_: No

_Type_: String

_Allowed values_: `ADMIN`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseName`

The name of the SAP HANA database.

_Required_: No

_Type_: String

_Pattern_: `^(?=.{1,100}$).*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecretId`

The secret ID created in AWS Secrets Manager to store the credentials
of the SAP application.

_Required_: No

_Type_: String

_Pattern_: `^(?=.{1,100}$).*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentInfo

Tag

All content copied from https://docs.aws.amazon.com/.
