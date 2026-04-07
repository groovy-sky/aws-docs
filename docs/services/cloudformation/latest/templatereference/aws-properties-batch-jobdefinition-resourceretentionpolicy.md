This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition ResourceRetentionPolicy

Specifies the resource retention policy settings for a job definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SkipDeregisterOnUpdate" : Boolean
}

```

### YAML

```yaml

  SkipDeregisterOnUpdate: Boolean

```

## Properties

`SkipDeregisterOnUpdate`

Specifies whether the previous revision of the job definition is retained in an active status after
UPDATE events for the resource. The default value is `false`. When the property is set to
`false`, the previous revision of the job definition is de-registered after a new revision
is created. When the property is set to `true`, the previous revision of the job definition
is not de-registered.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResourceRequirement

RetryStrategy
