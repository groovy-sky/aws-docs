This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::ResourcePolicy

Attaches a resource-based policy to a AWS End User Messaging SMS resource(phone number, sender Id, phone poll, or opt-out list) that is used for
sharing the resource. A shared resource can be a Pool, Opt-out list, Sender Id, or Phone number. For more information about
resource-based policies, see [Working with shared resources](https://docs.aws.amazon.com/sms-voice/latest/userguide/shared-resources.html) in the _AWS End User Messaging SMS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SMSVOICE::ResourcePolicy",
  "Properties" : {
      "PolicyDocument" : Json,
      "ResourceArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::SMSVOICE::ResourcePolicy
Properties:
  PolicyDocument: Json
  ResourceArn: String

```

## Properties

`PolicyDocument`

The JSON formatted resource-based policy to attach.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

The Amazon Resource Name (ARN) of the AWS End User Messaging SMS resource attached to the resource-based policy.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:\S+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns
`ResourceArn`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::SMSVOICE::SenderId
