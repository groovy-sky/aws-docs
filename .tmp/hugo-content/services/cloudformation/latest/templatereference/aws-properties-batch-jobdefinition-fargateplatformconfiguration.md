This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition FargatePlatformConfiguration

The platform configuration for jobs that are running on Fargate resources. Jobs that run
on Amazon EC2 resources must not specify this parameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PlatformVersion" : String
}

```

### YAML

```yaml

  PlatformVersion: String

```

## Properties

`PlatformVersion`

The AWS Fargate platform version where the jobs are running. A platform version is
specified only for jobs that are running on Fargate resources. If one isn't specified, the
`LATEST` platform version is used by default. This uses a recent, approved version of
the AWS Fargate platform for compute resources. For more information, see [AWS Fargate\
platform versions](../../../amazonecs/latest/developerguide/platform-versions.md) in the _Amazon Elastic Container Service Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluateOnExit

FirelensConfiguration

All content copied from https://docs.aws.amazon.com/.
