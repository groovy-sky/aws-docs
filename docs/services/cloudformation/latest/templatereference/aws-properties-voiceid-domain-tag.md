This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VoiceID::Domain Tag

###### Important

End of support notice: On May 20, 2026, AWS will end support for
Amazon Connect Voice ID. After May 20, 2026, you will no longer be able to access Voice ID on the Amazon Connect
console, access Voice ID features on the Amazon Connect admin website or Contact Control Panel, or access
Voice ID resources. For more information, visit
[Amazon Connect Voice ID end of support](../../../connect/latest/adminguide/amazonconnect-voiceid-end-of-support.md).

The tags used to organize, track, or control access for this resource. For example, {
"tags": {"key1":"value1", "key2":"value2"} }.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The first part of a key:value pair that forms a tag associated with a given resource.
For example, in the tag 'Department':'Sales', the key is 'Department'.

_Required_: Yes

_Type_: String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The second part of a key:value pair that forms a tag associated with a given resource.
For example, in the tag 'Department':'Sales', the value is 'Sales'.

_Required_: Yes

_Type_: String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServerSideEncryptionConfiguration

Next

All content copied from https://docs.aws.amazon.com/.
