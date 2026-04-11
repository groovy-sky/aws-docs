This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::MitigationAction AddThingsToThingGroupParams

Parameters used when defining a mitigation action that move a set of things to a thing group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OverrideDynamicGroups" : Boolean,
  "ThingGroupNames" : [ String, ... ]
}

```

### YAML

```yaml

  OverrideDynamicGroups: Boolean
  ThingGroupNames:
    - String

```

## Properties

`OverrideDynamicGroups`

Specifies if this mitigation action can move the things that triggered the mitigation action even if they are part of one or more dynamic thing groups.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThingGroupNames`

The list of groups to which you want to add the things that triggered the mitigation action. You can add a thing to a maximum of 10 groups, but you can't add a thing to more than one group in the same hierarchy.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `128 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionParams

EnableIoTLoggingParams

All content copied from https://docs.aws.amazon.com/.
