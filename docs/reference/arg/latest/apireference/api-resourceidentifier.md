---
title: "ResourceIdentifier"
---

# ResourceIdentifier

A structure that contains the ARN of a resource and its resource type.

## Contents

###### Note

In the following list, the required parameters are described first.

**ResourceArn**

The Amazon resource name (ARN) of a resource.

Type: String

Pattern: `arn:aws(-[a-z]+)*:[a-z0-9\-]*:([a-z]{2}(-[a-z]+)+-\d{1})?:([0-9]{12})?:.+`

Required: No

**ResourceType**

The resource type of a resource, such as `AWS::EC2::Instance`.

Type: String

Pattern: `AWS::[a-zA-Z0-9]+::\w+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/resource-groups-2017-11-27/ResourceIdentifier)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/resource-groups-2017-11-27/ResourceIdentifier)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/resource-groups-2017-11-27/ResourceIdentifier)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceFilter

ResourceQuery

All content copied from https://docs.aws.amazon.com/.
