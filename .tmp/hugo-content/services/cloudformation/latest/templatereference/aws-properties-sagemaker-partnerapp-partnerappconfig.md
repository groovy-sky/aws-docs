This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::PartnerApp PartnerAppConfig

A collection of configuration settings for the PartnerApp.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdminUsers" : [ String, ... ],
  "Arguments" : {Key: Value, ...}
}

```

### YAML

```yaml

  AdminUsers:
    - String
  Arguments:
    Key: Value

```

## Properties

`AdminUsers`

A list of users that will have administrative access to the Partner AI App.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Arguments`

Additional arguments passed to the Partner AI App during initialization or runtime.

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!\s*$).{1,256}$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SageMaker::PartnerApp

PartnerAppMaintenanceConfig

All content copied from https://docs.aws.amazon.com/.
