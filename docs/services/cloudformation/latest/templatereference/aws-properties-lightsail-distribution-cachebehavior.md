This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Distribution CacheBehavior

`CacheBehavior` is a property of the [AWS::Lightsail::Distribution](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html) resource. It describes the default cache behavior
of an Amazon Lightsail content delivery network (CDN) distribution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Behavior" : String
}

```

### YAML

```yaml

  Behavior: String

```

## Properties

`Behavior`

The cache behavior of the distribution.

The following cache behaviors can be specified:

- **`cache`** \- This option is best for static sites. When specified, your distribution
caches and serves your entire website as static content. This behavior is ideal for
websites with static content that doesn't change depending on who views it, or for
websites that don't use cookies, headers, or query strings to personalize
content.

- **`dont-cache`** \- This option is best for sites that serve a mix of static and dynamic
content. When specified, your distribution caches and serves only the content that is
specified in the distribution’s `CacheBehaviorPerPath` parameter. This
behavior is ideal for websites or web applications that use cookies, headers, and
query strings to personalize content for individual users.

_Required_: No

_Type_: String

_Allowed values_: `dont-cache | cache`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Lightsail::Distribution

CacheBehaviorPerPath
