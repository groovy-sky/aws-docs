This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain RStudioServerProDomainSettings

A collection of settings that configure the `RStudioServerPro` Domain-level
app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultResourceSpec" : ResourceSpec,
  "DomainExecutionRoleArn" : String,
  "RStudioConnectUrl" : String,
  "RStudioPackageManagerUrl" : String
}

```

### YAML

```yaml

  DefaultResourceSpec:
    ResourceSpec
  DomainExecutionRoleArn: String
  RStudioConnectUrl: String
  RStudioPackageManagerUrl: String

```

## Properties

`DefaultResourceSpec`

A collection that defines the default `InstanceType`, `SageMakerImageArn`, and
`SageMakerImageVersionArn` for the Domain.

_Required_: No

_Type_: [ResourceSpec](aws-properties-sagemaker-domain-resourcespec.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainExecutionRoleArn`

The ARN of the execution role for the `RStudioServerPro` Domain-level
app.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RStudioConnectUrl`

A URL pointing to an RStudio Connect server.

_Required_: No

_Type_: String

_Pattern_: `^(https:|http:|www\.)\S*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RStudioPackageManagerUrl`

A URL pointing to an RStudio Package Manager server.

_Required_: No

_Type_: String

_Pattern_: `^(https:|http:|www\.)\S*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RStudioServerProAppSettings

S3FileSystemConfig

All content copied from https://docs.aws.amazon.com/.
