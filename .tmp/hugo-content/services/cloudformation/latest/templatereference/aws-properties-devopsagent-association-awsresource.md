This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Association AWSResource

Defines an AWS resource to be monitored, including its type, ARN, and optional metadata.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceArn" : String,
  "ResourceMetadata" : Json,
  "ResourceType" : String
}

```

### YAML

```yaml

  ResourceArn: String
  ResourceMetadata: Json
  ResourceType: String

```

## Properties

`ResourceArn`

The Amazon Resource Name (ARN) of the resource.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceMetadata`

Additional metadata specific to the resource. This is an optional JSON object that can include resource-specific
information to provide additional context for monitoring and management.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

The type of AWS resource.

_Allowed Values_: `AWS::CloudFormation::Stack` \| `AWS::ECR::Repository` \| `AWS::S3::Bucket` \| `AWS::S3::Object`

_Required_: No

_Type_: String

_Allowed values_: `AWS::CloudFormation::Stack | AWS::ECR::Repository | AWS::S3::Bucket | AWS::S3::Object`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWSConfiguration

DynatraceConfiguration

All content copied from https://docs.aws.amazon.com/.
