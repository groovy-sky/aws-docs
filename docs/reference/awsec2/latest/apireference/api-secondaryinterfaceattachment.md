# SecondaryInterfaceAttachment

Describes the attachment of a secondary interface to an instance.

## Contents

**attachmentId**

The ID of the attachment.

Type: String

Required: No

**attachTime**

The timestamp when the attachment was created.

Type: Timestamp

Required: No

**deleteOnTermination**

Indicates whether the secondary interface is deleted when the instance is terminated.

The only supported value for this field is `true`.

Type: Boolean

Required: No

**deviceIndex**

The device index of the secondary interface.

Type: Integer

Required: No

**instanceId**

The ID of the instance to which the secondary interface is attached.

Type: String

Required: No

**instanceOwnerId**

The AWS account ID of the owner of the instance.

Type: String

Required: No

**networkCardIndex**

The index of the network card.

Type: Integer

Required: No

**status**

The attachment state.

Type: String

Valid Values: `attaching | attached | detaching | detached`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/secondaryinterfaceattachment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/secondaryinterfaceattachment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/secondaryinterfaceattachment.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SecondaryInterface

SecondaryInterfaceIpv4Address
