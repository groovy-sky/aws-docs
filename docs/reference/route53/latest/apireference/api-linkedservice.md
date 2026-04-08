# LinkedService

If a health check or hosted zone was created by another service,
`LinkedService` is a complex type that describes the service that created
the resource. When a resource is created by another service, you can't edit or delete it
using Amazon Route 53.

## Contents

**Description**

If the health check or hosted zone was created by another service, an optional
description that can be provided by the other service. When a resource is created by
another service, you can't edit or delete it using Amazon Route 53.

Type: String

Length Constraints: Maximum length of 256.

Required: No

**ServicePrincipal**

If the health check or hosted zone was created by another service, the service that
created the resource. When a resource is created by another service, you can't edit or
delete it using Amazon Route 53.

Type: String

Length Constraints: Maximum length of 128.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/linkedservice.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/linkedservice.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/linkedservice.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

KeySigningKey

LocationSummary
