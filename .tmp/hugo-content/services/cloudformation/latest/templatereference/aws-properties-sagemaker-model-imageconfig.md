This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Model ImageConfig

Specifies whether the model container is in Amazon ECR or a private Docker registry
accessible from your Amazon Virtual Private Cloud (VPC).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RepositoryAccessMode" : String,
  "RepositoryAuthConfig" : RepositoryAuthConfig
}

```

### YAML

```yaml

  RepositoryAccessMode: String
  RepositoryAuthConfig:
    RepositoryAuthConfig

```

## Properties

`RepositoryAccessMode`

Set this to one of the following values:

- `Platform` \- The model image is hosted in Amazon ECR.

- `Vpc` \- The model image is hosted in a private Docker registry in
your VPC.

_Required_: Yes

_Type_: String

_Allowed values_: `Platform | Vpc`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RepositoryAuthConfig`

(Optional) Specifies an authentication configuration for the private docker registry
where your model image is hosted. Specify a value for this property only if you
specified `Vpc` as the value for the `RepositoryAccessMode` field,
and the private Docker registry where the model image is hosted requires
authentication.

_Required_: No

_Type_: [RepositoryAuthConfig](aws-properties-sagemaker-model-repositoryauthconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HubAccessConfig

InferenceExecutionConfig

All content copied from https://docs.aws.amazon.com/.
