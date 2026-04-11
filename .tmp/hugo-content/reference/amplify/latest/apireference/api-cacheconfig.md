# CacheConfig

Describes the cache configuration for an Amplify app.

For more information about how Amplify applies an optimal cache
configuration for your app based on the type of content that is being served, see [Managing cache configuration](../../../../services/amplify/latest/userguide/managing-cache-configuration.md) in the _Amplify User_
_guide_.

## Contents

**type**

The type of cache configuration to use for an Amplify app.

The `AMPLIFY_MANAGED` cache configuration automatically applies an
optimized cache configuration for your app based on its platform, routing rules, and
rewrite rules.

The `AMPLIFY_MANAGED_NO_COOKIES` cache configuration type is the same as
`AMPLIFY_MANAGED`, except that it excludes all cookies from the cache
key. This is the default setting.

Type: String

Valid Values: `AMPLIFY_MANAGED | AMPLIFY_MANAGED_NO_COOKIES`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/amplify-2017-07-25/cacheconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/amplify-2017-07-25/cacheconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/amplify-2017-07-25/cacheconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Branch

Certificate

All content copied from https://docs.aws.amazon.com/.
