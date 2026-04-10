This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ResourceDefinition SecretsManagerSecretResourceData

Settings for a secret resource, which references a secret from AWS Secrets Manager.
AWS IoT Greengrass stores a local, encrypted copy of the secret on the Greengrass core,
where it can be securely accessed by connectors and Lambda functions. For
more information, see [Deploy Secrets to the AWS IoT Greengrass Core](../../../greengrass/v1/developerguide/secrets.md) in the _AWS IoT Greengrass Version 1 Developer Guide_.

In an CloudFormation template, `SecretsManagerSecretResourceData` can be
used in the [`ResourceDataContainer`](../userguide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalStagingLabelsToDownload" : [ String, ... ],
  "ARN" : String
}

```

### YAML

```yaml

  AdditionalStagingLabelsToDownload:
    - String
  ARN: String

```

## Properties

`AdditionalStagingLabelsToDownload`

The staging labels whose values you want to make available on the core, in addition to
`AWSCURRENT`.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ARN`

The Amazon Resource Name (ARN) of the Secrets Manager secret to make available
on the core. The value of the secret's latest version (represented by the
`AWSCURRENT` staging label) is included by default.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [SecretsManagerSecretResourceData](../../../../reference/greengrass/v1/apireference/definitions-secretsmanagersecretresourcedata.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SageMakerMachineLearningModelResourceData

AWS::Greengrass::ResourceDefinitionVersion

All content copied from https://docs.aws.amazon.com/.
