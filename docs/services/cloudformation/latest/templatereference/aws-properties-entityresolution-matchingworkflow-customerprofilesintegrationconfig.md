This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::MatchingWorkflow CustomerProfilesIntegrationConfig

The `CustomerProfilesIntegrationConfig` property type specifies Property description not available. for an [AWS::EntityResolution::MatchingWorkflow](aws-resource-entityresolution-matchingworkflow.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainArn" : String,
  "ObjectTypeArn" : String
}

```

### YAML

```yaml

  DomainArn: String
  ObjectTypeArn: String

```

## Properties

`DomainArn`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):profile:[a-z]{2}-[a-z]{1,10}-[0-9]:[0-9]{12}:(domains/[a-zA-Z_0-9-]{1,255})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectTypeArn`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):profile:[a-z]{2}-[a-z]{1,10}-[0-9]:[0-9]{12}:(domains/[a-zA-Z_0-9-]{1,255}/object-types/[a-zA-Z_0-9-]{1,255})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EntityResolution::MatchingWorkflow

IncrementalRunConfig

All content copied from https://docs.aws.amazon.com/.
