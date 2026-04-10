# Resource

Specifies the type and name of a resource referenced by an event.

## Contents

**ResourceName**

The name of the resource referenced by the event returned. These are user-created names
whose values will depend on the environment. For example, the resource name might be
"auto-scaling-test-group" for an Auto Scaling Group or "i-1234567" for an EC2
Instance.

Type: String

Required: No

**ResourceType**

The type of a resource referenced by the event returned. When the resource type cannot
be determined, null is returned. Some examples of resource types are: **Instance** for EC2, **Trail** for CloudTrail, **DBInstance** for Amazon RDS, and **AccessKey** for IAM. To learn more about how to look up and filter
events by the resource types supported for a service, see [Filtering CloudTrail Events](../../../../services/awscloudtrail/latest/userguide/view-cloudtrail-events-console.md#filtering-cloudtrail-events).

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-2013-11-01/resource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-2013-11-01/resource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-2013-11-01/resource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RequestWidget

ResourceTag

All content copied from https://docs.aws.amazon.com/.
