This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate ExperimentTemplateExperimentOptions

Describes the experiment options for an experiment template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountTargeting" : String,
  "EmptyTargetResolutionMode" : String
}

```

### YAML

```yaml

  AccountTargeting: String
  EmptyTargetResolutionMode: String

```

## Properties

`AccountTargeting`

The account targeting setting for an experiment template.

_Required_: No

_Type_: String

_Allowed values_: `multi-account | single-account`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EmptyTargetResolutionMode`

The empty target resolution mode for an experiment template.

_Required_: No

_Type_: String

_Allowed values_: `fail | skip`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ExperimentTemplateAction

ExperimentTemplateExperimentReportConfiguration
