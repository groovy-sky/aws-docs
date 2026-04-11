This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project RegistryCredential

`RegistryCredential` is a property of the
[AWS CodeBuild Project Environment](../userguide/aws-properties-codebuild-project-environment.md) property type that specifies information about credentials that provide access to a private Docker registry. When this is set:

- `imagePullCredentialsType` must be set to `SERVICE_ROLE`.

- images cannot be curated or an Amazon ECR image.

For more information, see [Private Registry with \
AWS Secrets Manager Sample for AWS CodeBuild](../../../codebuild/latest/userguide/sample-private-registry.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Credential" : String,
  "CredentialProvider" : String
}

```

### YAML

```yaml

  Credential: String
  CredentialProvider: String

```

## Properties

`Credential`

The Amazon Resource Name (ARN) or name of credentials created using AWS Secrets Manager.

###### Note

The `credential` can use the name of the credentials only if they
exist in your current AWS Region.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CredentialProvider`

The service that created the credentials to access a private Docker registry. The
valid value, SECRETS\_MANAGER, is for AWS Secrets Manager.

_Required_: Yes

_Type_: String

_Allowed values_: `SECRETS_MANAGER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [RegistryCredential](../../../../reference/codebuild/latest/apireference/api-registrycredential.md) in the _AWS CodeBuild API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProjectTriggers

S3LogsConfig

All content copied from https://docs.aws.amazon.com/.
