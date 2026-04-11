This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Model MultiModelConfig

Specifies additional configuration for hosting multi-model endpoints.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ModelCacheSetting" : String
}

```

### YAML

```yaml

  ModelCacheSetting: String

```

## Properties

`ModelCacheSetting`

Whether to cache models for a multi-model endpoint. By default, multi-model endpoints cache models so that a
model does not have to be loaded into memory each time it is invoked. Some use cases do not benefit from model
caching. For example, if an endpoint hosts a large number of models that are each invoked infrequently, the endpoint
might perform better if you disable model caching. To disable model caching, set the value of this parameter to
Disabled.

_Required_: No

_Type_: String

_Allowed values_: `Enabled | Disabled`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelDataSource

RepositoryAuthConfig

All content copied from https://docs.aws.amazon.com/.
