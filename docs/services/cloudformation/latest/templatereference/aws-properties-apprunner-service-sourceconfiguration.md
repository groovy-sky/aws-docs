This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service SourceConfiguration

Describes the source deployed to an AWS App Runner service. It can be a code or an image repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationConfiguration" : AuthenticationConfiguration,
  "AutoDeploymentsEnabled" : Boolean,
  "CodeRepository" : CodeRepository,
  "ImageRepository" : ImageRepository
}

```

### YAML

```yaml

  AuthenticationConfiguration:
    AuthenticationConfiguration
  AutoDeploymentsEnabled: Boolean
  CodeRepository:
    CodeRepository
  ImageRepository:
    ImageRepository

```

## Properties

`AuthenticationConfiguration`

Describes the resources that are needed to authenticate access to some source repositories.

_Required_: No

_Type_: [AuthenticationConfiguration](aws-properties-apprunner-service-authenticationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoDeploymentsEnabled`

If `true`, continuous integration from the source repository is enabled for the App Runner service. Each repository change (including any source
code commit or new image version) starts a deployment.

Default: App Runner sets to `false` for a source image that uses an ECR Public repository or an ECR repository that's in an AWS account other than the one that the service is in. App Runner sets to `true` in all other cases (which currently include a source code
repository or a source image using a same-account ECR repository).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeRepository`

The description of a source code
repository.

You must provide either this member or `ImageRepository` (but not both).

_Required_: No

_Type_: [CodeRepository](aws-properties-apprunner-service-coderepository.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageRepository`

The description of a source image
repository.

You must provide either this member or `CodeRepository` (but not both).

_Required_: No

_Type_: [ImageRepository](aws-properties-apprunner-service-imagerepository.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceCodeVersion

Tag

All content copied from https://docs.aws.amazon.com/.
