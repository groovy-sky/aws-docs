This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerAddonInstance

Creates an Add On instance for the subscription indicated in the request. The
resulting Amazon Resource Name (ARN) can be used in a conditional statement for a rule set or traffic policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::MailManagerAddonInstance",
  "Properties" : {
      "AddonSubscriptionId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SES::MailManagerAddonInstance
Properties:
  AddonSubscriptionId: String
  Tags:
    - Tag

```

## Properties

`AddonSubscriptionId`

The subscription ID for the instance.

_Required_: Yes

_Type_: String

_Pattern_: `^as-[a-zA-Z0-9]{1,64}$`

_Minimum_: `4`

_Maximum_: `67`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-mailmanageraddoninstance-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`AddonInstanceArn`

The Amazon Resource Name (ARN) of the Add On instance.

`AddonInstanceId`

The unique ID of the Add On instance.

`AddonName`

The name of the Add On for the instance.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
