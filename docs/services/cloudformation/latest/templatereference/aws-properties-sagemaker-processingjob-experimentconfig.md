This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob ExperimentConfig

Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the [CreateProcessingJob](../../../../reference/sagemaker/latest/apireference/api-createprocessingjob.md)
API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExperimentName" : String,
  "RunName" : String,
  "TrialComponentDisplayName" : String,
  "TrialName" : String
}

```

### YAML

```yaml

  ExperimentName: String
  RunName: String
  TrialComponentDisplayName: String
  TrialName: String

```

## Properties

`ExperimentName`

The name of an existing experiment to associate with the trial component.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9]){0,119}`

_Maximum_: `120`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RunName`

The name of the experiment run to associate with the trial component.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9]){0,119}`

_Maximum_: `120`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TrialComponentDisplayName`

The display name for the trial component. If this key isn't specified, the display name is
the trial component name.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9]){0,119}`

_Maximum_: `120`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TrialName`

The name of an existing trial to associate the trial component with. If not specified, a
new trial is created.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9]){0,119}`

_Maximum_: `120`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatasetDefinition

FeatureStoreOutput

All content copied from https://docs.aws.amazon.com/.
