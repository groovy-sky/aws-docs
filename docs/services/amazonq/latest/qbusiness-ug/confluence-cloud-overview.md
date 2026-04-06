# Confluence (Cloud) connector overview

The following table contains an overview of the Amazon Q Business
Confluence (Cloud) connector and its supported features.

CategoryFeatureSupport**Security****Authentication type**Basic, OAuth 2.0 with Refresh Token Flow**Authentication credentials**

For Basic authentication

- Confluence Cloud URL

- Confluence username

- Password (Confluence (Cloud) site
token)

For OAuth 2.0 authentication with Refresh
Token Flow

- App key

- App secret

- Access token

- [Refresh token](https://developer.atlassian.com/cloud/jira/platform/oauth-2-3lo-apps)

###### Note

Access and refresh tokens expire in 1 hour. For
information on regenerating tokens, see [Atlassian Developer\
Documentation](https://developer.atlassian.com/cloud/jira/platform/oauth-2-3lo-apps).

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes. For more information, see [ACL crawling](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/confluence-cloud-user-management.html).**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**Yes**Crawl features****Custom metadata**Yes**Entities**Yes. The following entities are supported:

- Space

- Page

- Blog post

- Comment

- Attachment

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. For more information, see [Field mappings](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/confluence-cloud-field-mappings.html).**Filters**Yes. The following filters are supported:

- Inclusion exclusion filters for **Space**
**key** and **Space URL**

- Inclusion exclusion filters on **File**
**Type** for **Attachment**
**entity**

- Supports regex filters for entities

- Supports inclusion and exclusion filters for
**File size**

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and incremental (new, modified, and deleted)
sync.**[File types](doc-types.md)**Supports all files supported by Amazon Q.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Confluence (Cloud)

Prerequisites
