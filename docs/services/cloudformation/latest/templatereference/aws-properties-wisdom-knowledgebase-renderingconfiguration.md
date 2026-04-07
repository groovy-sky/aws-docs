This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase RenderingConfiguration

Information about how to render the content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TemplateUri" : String
}

```

### YAML

```yaml

  TemplateUri: String

```

## Properties

`TemplateUri`

A URI template containing exactly one variable in `${variableName} ` format. This can only be set for
`EXTERNAL` knowledge bases. For Salesforce, ServiceNow, and Zendesk, the variable must be one of the
following:

- Salesforce: `Id`, `ArticleNumber`, `VersionNumber`, `Title`,
`PublishStatus`, or `IsDeleted`

- ServiceNow: `number`, `short_description`, `sys_mod_count`,
`workflow_state`, or `active`

- Zendesk: `id`, `title`, `updated_at`, or `draft`

The variable is replaced with the actual value for a piece of content when calling [GetContent](https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_GetContent.html).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ParsingPrompt

SeedUrl
