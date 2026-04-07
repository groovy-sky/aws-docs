This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::ProvisioningTemplate ProvisioningHook

Structure that contains payloadVersion and targetArn. Provisioning hooks can be used
when fleet provisioning to validate device parameters before allowing the device to be
provisioned.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PayloadVersion" : String,
  "TargetArn" : String
}

```

### YAML

```yaml

  PayloadVersion: String
  TargetArn: String

```

## Properties

`PayloadVersion`

The payload that was sent to the target function. The valid payload is
`"2020-04-01"`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArn`

The ARN of the target function.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IoT::ProvisioningTemplate

Tag
