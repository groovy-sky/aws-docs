This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase RedshiftServerlessConfiguration

Contains configurations for authentication to Amazon Redshift Serverless.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthConfiguration" : RedshiftServerlessAuthConfiguration,
  "WorkgroupArn" : String
}

```

### YAML

```yaml

  AuthConfiguration:
    RedshiftServerlessAuthConfiguration
  WorkgroupArn: String

```

## Properties

`AuthConfiguration`

Specifies configurations for authentication to an Amazon Redshift provisioned data warehouse.

_Required_: Yes

_Type_: [RedshiftServerlessAuthConfiguration](aws-properties-bedrock-knowledgebase-redshiftserverlessauthconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkgroupArn`

The ARN of the Amazon Redshift workgroup.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:(aws(-[a-z]+)*):redshift-serverless:[a-z]{2}(-gov)?-[a-z]+-\d{1}:\d{12}:workgroup/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftServerlessAuthConfiguration

S3Location

All content copied from https://docs.aws.amazon.com/.
