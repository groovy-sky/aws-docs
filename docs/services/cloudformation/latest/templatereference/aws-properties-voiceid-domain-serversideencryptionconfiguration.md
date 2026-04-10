This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VoiceID::Domain ServerSideEncryptionConfiguration

###### Important

End of support notice: On May 20, 2026, AWS will end support for
Amazon Connect Voice ID. After May 20, 2026, you will no longer be able to access Voice ID on the Amazon Connect
console, access Voice ID features on the Amazon Connect admin website or Contact Control Panel, or access
Voice ID resources. For more information, visit
[Amazon Connect Voice ID end of support](../../../connect/latest/adminguide/amazonconnect-voiceid-end-of-support.md).

The configuration containing information about the customer managed key used for
encrypting customer data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String
}

```

### YAML

```yaml

  KmsKeyId: String

```

## Properties

`KmsKeyId`

The identifier of the KMS key to use to encrypt data stored by
Voice ID. Voice ID doesn't support asymmetric customer managed keys.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::VoiceID::Domain

Tag

All content copied from https://docs.aws.amazon.com/.
