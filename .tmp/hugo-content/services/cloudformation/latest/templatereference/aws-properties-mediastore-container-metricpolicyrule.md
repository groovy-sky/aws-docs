This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaStore::Container MetricPolicyRule

A setting that enables metrics at the object level. Each rule contains an object group and an object group name. If the policy includes the MetricPolicyRules parameter, you must include at least one rule. Each metric policy can include up to five rules by default. You can also [request a quota increase](https://console.aws.amazon.com/servicequotas/home?region=us-east-1) to allow up to 300 rules per policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ObjectGroup" : String,
  "ObjectGroupName" : String
}

```

### YAML

```yaml

  ObjectGroup: String
  ObjectGroupName: String

```

## Properties

`ObjectGroup`

A path or file name that defines which objects to include in the group. Wildcards (\*) are acceptable.

_Required_: Yes

_Type_: String

_Pattern_: `/?(?:[A-Za-z0-9_=:\.\-\~\*]+/){0,10}(?:[A-Za-z0-9_=:\.\-\~\*]+)?/?`

_Minimum_: `1`

_Maximum_: `900`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectGroupName`

A name that allows you to refer to the object group.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_]+`

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricPolicy

Tag

All content copied from https://docs.aws.amazon.com/.
