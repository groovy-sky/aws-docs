This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow ServiceNowSourceProperties

The properties that are applied when ServiceNow is being used as a source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Object" : String
}

```

### YAML

```yaml

  Object: String

```

## Properties

`Object`

The object specified in the ServiceNow flow source.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ServiceNowSourceProperties](../../../../reference/appflow/1-0/apireference/api-servicenowsourceproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ScheduledTriggerProperties

SingularSourceProperties
