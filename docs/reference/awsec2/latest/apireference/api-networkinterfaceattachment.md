# NetworkInterfaceAttachment

Describes a network interface attachment.

## Contents

**attachmentId**

The ID of the network interface attachment.

Type: String

Required: No

**attachTime**

The timestamp indicating when the attachment initiated.

Type: Timestamp

Required: No

**deleteOnTermination**

Indicates whether the network interface is deleted when the instance is
terminated.

Type: Boolean

Required: No

**deviceIndex**

The device index of the network interface attachment on the instance.

Type: Integer

Required: No

**enaQueueCount**

The number of ENA queues created with the instance.

Type: Integer

Required: No

**enaSrdSpecification**

Configures ENA Express for the network interface that this action attaches to the
instance.

Type: [AttachmentEnaSrdSpecification](api-attachmentenasrdspecification.md) object

Required: No

**instanceId**

The ID of the instance.

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

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/networkinterfaceattachment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/networkinterfaceattachment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/networkinterfaceattachment.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

NetworkInterfaceAssociation

NetworkInterfaceAttachmentChanges
