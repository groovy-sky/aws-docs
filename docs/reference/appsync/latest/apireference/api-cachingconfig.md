# CachingConfig

The caching configuration for a resolver that has caching activated.

## Contents

**ttl**

The TTL in seconds for a resolver that has caching activated.

Valid values are 1–3,600 seconds.

Type: Long

Required: Yes

**cachingKeys**

The caching keys for a resolver that has caching activated.

Valid values are entries from the `$context.arguments`,
`$context.source`, and `$context.identity` maps.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/cachingconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/cachingconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/cachingconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BadRequestDetail

ChannelNamespace

All content copied from https://docs.aws.amazon.com/.
