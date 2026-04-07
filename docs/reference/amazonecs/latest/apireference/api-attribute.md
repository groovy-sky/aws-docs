# Attribute

An attribute is a name-value pair that's associated with an Amazon ECS object. Use
attributes to extend the Amazon ECS data model by adding custom metadata to your
resources. For more information, see [Attributes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html#attributes) in the _Amazon Elastic Container Service Developer_
_Guide_.

## Contents

**name**

The name of the attribute. The `name` must contain between 1 and 128
characters. The name may contain letters (uppercase and lowercase), numbers, hyphens
(-), underscores (\_), forward slashes (/), back slashes (\\), or periods (.).

Type: String

Required: Yes

**targetId**

The ID of the target. You can specify the short form ID for a resource or the full
Amazon Resource Name (ARN).

Type: String

Required: No

**targetType**

The type of the target to attach the attribute with. This parameter is required if you
use the short form ID for a resource instead of the full ARN.

Type: String

Valid Values: `container-instance`

Required: No

**value**

The value of the attribute. The `value` must contain between 1 and 128
characters. It can contain letters (uppercase and lowercase), numbers, hyphens (-),
underscores (\_), periods (.), at signs (@), forward slashes (/), back slashes (\\),
colons (:), or spaces. The value can't start or end with a space.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/Attribute)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/Attribute)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/Attribute)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AttachmentStateChange

AutoScalingGroupProvider
