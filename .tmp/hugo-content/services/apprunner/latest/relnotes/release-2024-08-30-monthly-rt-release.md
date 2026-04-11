AWS App Runner will no longer be open to new customers starting April 30, 2026. If you would like to use
App Runner, sign up prior to that date. Existing customers can continue to use the service as normal. For more information, see
[AWS App Runner availability\
change](../dg/apprunner-availability-change.md).

# Release: App Runner runtime updates on August 30, 2024

This release provides minor version updates for the following language runtimes:
.NET.

It also provides package updates to the following platforms:
Java, .NET, Python, Ruby.

**Release date:** August 30, 2024

## App Runner managed platforms

App Runner provides convenient platform-specific managed runtimes. When you use a managed runtime, App Runner starts with a managed runtime base image to build a
container image from your source code. For more information, see [App Runner managed platforms](../dg/service-source-code.md#service-source-code.managed-platforms) in the
_AWS App Runner Developer Guide_.

## Changes

The following table lists the changes included in this release.

**Category****Description**

**Base container image updates**

Made the following updates to the base container images:

**Component****Update**

Base container image for AL2023

Updated base container image to version **2023.5.20240819.0**.

###### Note

This image only applies to Node.js 18 and Python 3.11 runtimes.

Base container image for AL2

Updated base container image to version **2.0.20240816.0**.

###### Note

This image applies to all runtimes, except for Node.js 18 and Python 3.11.

To view the Amazon Linux container image on the _Amazon ECR Public Gallery,_ see the _Image tags_ tab on [Amazon ECR Public Gallery -\
amazonlinux](https://gallery.ecr.aws/amazonlinux/amazonlinux).

**Platform-specific updates**

Made these platform-specific updates:

**Platform****Update**

**Corretto**

No updates to language versions.

Tools Updates:

- Updated Maven to 3.9.9.

**.NET Core**

[Supported runtimes](../dg/service-source-code-dotnet-releases.md)

Updated **.NET Core 6.0** to 6.0.33.

Package updates:

- Updated .NET SDK to 6.0.425.

**Python**

[Supported runtimes](../dg/service-source-code-python-releases.md)

No updates to language versions.

Package updates:

- Updated SQLite to 3.46.1.

**Ruby**

[Supported runtimes](../dg/service-source-code-ruby-releases.md)

No updates to language versions.

Package updates:

- Updated SQLite to 3.46.1.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Runtime updates 2024-10-01

Runtime updates 2024-07-31

All content copied from https://docs.aws.amazon.com/.
