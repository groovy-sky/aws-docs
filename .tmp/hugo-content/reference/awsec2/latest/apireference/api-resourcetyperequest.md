# ResourceTypeRequest

A resource type to check for image references. Associated options can also be specified if the
resource type is an EC2 instance or launch template.

## Contents

**ResourceType**

The resource type.

Type: String

Valid Values: `ec2:Instance | ec2:LaunchTemplate | ssm:Parameter | imagebuilder:ImageRecipe | imagebuilder:ContainerRecipe`

Required: No

**ResourceTypeOption.N**

The options that affect the scope of the response. Valid only when
`ResourceType` is `ec2:Instance` or
`ec2:LaunchTemplate`.

Type: Array of [ResourceTypeOption](api-resourcetypeoption.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/resourcetyperequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/resourcetyperequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/resourcetyperequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ResourceTypeOption

ResponseError
