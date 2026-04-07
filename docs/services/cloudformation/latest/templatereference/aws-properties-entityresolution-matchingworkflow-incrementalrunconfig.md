This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::MatchingWorkflow IncrementalRunConfig

Optional. An object that defines the incremental run type. This object contains only the
`incrementalRunType` field, which appears as "Automatic" in the console.

###### Important

For workflows where `resolutionType` is `ML_MATCHING` or `PROVIDER`,
incremental processing is not supported.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IncrementalRunType" : String
}

```

### YAML

```yaml

  IncrementalRunType: String

```

## Properties

`IncrementalRunType`

The type of incremental run. The only valid value is `IMMEDIATE`. This
appears as "Automatic" in the console.

###### Important

For workflows where `resolutionType` is `ML_MATCHING` or `PROVIDER`,
incremental processing is not supported.

_Required_: Yes

_Type_: String

_Allowed values_: `IMMEDIATE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomerProfilesIntegrationConfig

InputSource
