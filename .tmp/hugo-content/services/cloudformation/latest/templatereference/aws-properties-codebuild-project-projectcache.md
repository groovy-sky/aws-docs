This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project ProjectCache

`ProjectCache` is a property of the
[AWS CodeBuild Project](../userguide/aws-resource-codebuild-project.md)
resource that specifies information about the cache for the build project. If `ProjectCache` is not specified, then both of its properties
default to `NO_CACHE`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CacheNamespace" : String,
  "Location" : String,
  "Modes" : [ String, ... ],
  "Type" : String
}

```

### YAML

```yaml

  CacheNamespace: String
  Location: String
  Modes:
    - String
  Type: String

```

## Properties

`CacheNamespace`

Defines the scope of the cache. You can use this namespace to share a cache across
multiple projects. For more information, see [Cache sharing \
between projects](../../../codebuild/latest/userguide/caching-s3.md#caching-s3-sharing) in the _AWS CodeBuild User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Location`

Information about the cache location:

- `NO_CACHE` or `LOCAL`: This value is ignored.

- `S3`: This is the S3 bucket name/prefix.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Modes`

An array of strings that specify the local cache modes. You can use one or more local
cache modes at the same time. This is only used for `LOCAL` cache
types.

Possible values are:

LOCAL\_SOURCE\_CACHE

Caches Git metadata for primary and secondary sources. After the cache is
created, subsequent builds pull only the change between commits. This mode
is a good choice for projects with a clean working directory and a source
that is a large Git repository. If you choose this option and your project
does not use a Git repository (GitHub, GitHub Enterprise, or Bitbucket), the
option is ignored.

LOCAL\_DOCKER\_LAYER\_CACHE

Caches existing Docker layers. This mode is a good choice for projects
that build or pull large Docker images. It can prevent the performance
issues caused by pulling large Docker images down from the network.

###### Note

- You can use a Docker layer cache in the Linux environment
only.

- The `privileged` flag must be set so that your
project has the required Docker permissions.

- You should consider the security implications before you use a
Docker layer cache.

LOCAL\_CUSTOM\_CACHE

Caches directories you specify in the buildspec file. This mode is a good
choice if your build scenario is not suited to one of the other three local
cache modes. If you use a custom cache:

- Only directories can be specified for caching. You cannot specify
individual files.

- Symlinks are used to reference cached directories.

- Cached directories are linked to your build before it downloads
its project sources. Cached items are overridden if a source item
has the same name. Directories are specified using cache paths in
the buildspec file.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of cache used by the build project. Valid values include:

- `NO_CACHE`: The build project does not use any cache.

- `S3`: The build project reads and writes from and to S3.

- `LOCAL`: The build project stores a cache locally on a build host
that is only available to that build host.

_Required_: Yes

_Type_: String

_Allowed values_: `NO_CACHE | S3 | LOCAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ProjectCache](../../../../reference/codebuild/latest/apireference/api-projectcache.md) in the _AWS CodeBuild API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProjectBuildBatchConfig

ProjectFileSystemLocation

All content copied from https://docs.aws.amazon.com/.
