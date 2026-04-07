# ServiceVolumeConfiguration

The configuration for a volume specified in the task definition as a volume that is
configured at launch time. Currently, the only supported volume type is an Amazon EBS
volume.

## Contents

**name**

The name of the volume. This value must match the volume name from the
`Volume` object in the task definition.

Type: String

Required: Yes

**managedEBSVolume**

The configuration for the Amazon EBS volume that Amazon ECS creates and manages on
your behalf. These settings are used to create each Amazon EBS volume, with one volume
created for each task in the service. The Amazon EBS volumes are visible in your account
in the Amazon EC2 console once they are created.

Type: [ServiceManagedEBSVolumeConfiguration](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ServiceManagedEBSVolumeConfiguration.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/ServiceVolumeConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/ServiceVolumeConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/ServiceVolumeConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ServiceRevisionSummary

Session
