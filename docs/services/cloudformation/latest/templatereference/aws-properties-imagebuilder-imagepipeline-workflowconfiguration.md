This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ImagePipeline WorkflowConfiguration

Contains control settings and configurable inputs for a workflow
resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OnFailure" : String,
  "ParallelGroup" : String,
  "Parameters" : [ WorkflowParameter, ... ],
  "WorkflowArn" : String
}

```

### YAML

```yaml

  OnFailure: String
  ParallelGroup: String
  Parameters:
    - WorkflowParameter
  WorkflowArn: String

```

## Properties

`OnFailure`

The action to take if the workflow fails.

_Required_: No

_Type_: String

_Allowed values_: `CONTINUE | ABORT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParallelGroup`

Test workflows are defined within named runtime groups called parallel groups.
The parallel group is the named group that contains this test workflow. Test
workflows within a parallel group can run at the same time. Image Builder starts up to five
test workflows in the group at the same time, and starts additional workflows as
others complete, until all workflows in the group have completed. This field only
applies for test workflows.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9-_+#]{0,99}$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

Contains parameter values for each of the parameters that the workflow
document defined for the workflow resource.

_Required_: No

_Type_: Array of [WorkflowParameter](aws-properties-imagebuilder-imagepipeline-workflowparameter.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowArn`

The Amazon Resource Name (ARN) of the workflow resource.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(?:-[a-z]+)*:imagebuilder:[a-z]{2,}(?:-[a-z]+)+-[0-9]+:(?:[0-9]{12}|aws(?:-[a-z-]+)?):workflow/(build|test|distribution)/[a-z0-9-_]+/(?:(?:([0-9]+|x)\.([0-9]+|x)\.([0-9]+|x))|(?:[0-9]+\.[0-9]+\.[0-9]+/[0-9]+))$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Schedule

WorkflowParameter

All content copied from https://docs.aws.amazon.com/.
