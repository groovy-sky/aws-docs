# SigningRepositoryFilter

A repository filter used to determine which repositories have their
images automatically signed on push. Each filter consists of a filter type and filter value.

## Contents

**filter**

The filter value used to match repository names. When using
`WILDCARD_MATCH`, the `*` character matches any sequence of characters.

Examples:

- `myapp/*` \- Matches all repositories starting with
`myapp/`

- `*/production` \- Matches all repositories ending with
`/production`

- `*prod*` \- Matches all repositories containing
`prod`

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `^(?:[a-z0-9*]+(?:[._-][a-z0-9*]+)*/)*[a-z0-9*]+(?:[._-][a-z0-9*]+)*$`

Required: Yes

**filterType**

The type of filter to apply. Currently, only `WILDCARD_MATCH` is supported,
which uses wildcard patterns to match repository names.

Type: String

Valid Values: `WILDCARD_MATCH`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/signingrepositoryfilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/signingrepositoryfilter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/signingrepositoryfilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SigningConfiguration

SigningRule

All content copied from https://docs.aws.amazon.com/.
