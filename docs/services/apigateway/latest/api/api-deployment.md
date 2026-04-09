# Deployment

An immutable representation of a RestApi resource that can be called by users using Stages. A deployment must be associated with a Stage for it to be callable over the Internet.

## Contents

**apiSummary**

A summary of the RestApi at the date and time that the deployment resource was created.

Type: String to string to [MethodSnapshot](api-methodsnapshot.md) object map map

Required: No

**createdDate**

The date and time that the deployment resource was created.

Type: Timestamp

Required: No

**description**

The description for the deployment resource.

Type: String

Required: No

**id**

The identifier for the deployment resource.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/deployment.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/deployment.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/deployment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientCertificate

DeploymentCanarySettings

All content copied from https://docs.aws.amazon.com/.
