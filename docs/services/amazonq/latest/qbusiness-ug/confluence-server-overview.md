# Confluence (Server/Data Center) connector overview

The following table gives an overview of the Amazon Q Business
Confluence (Server/Data Center) connector and its supported features.

CategoryFeatureSupport**Security****Authentication type**Basic, OAuth 2.0 with Refresh Token Flow, Personal Access
Token**Authentication credentials**

For Basic authentication:

- Confluence Server/Data Center username

- Confluence Server/Data Center password

For OAuth 2.0 authentication with Refresh Token
Flow:

- App key

- App secret

- Access token

- [Refresh token](https://developer.atlassian.com/cloud/jira/platform/oauth-2-3lo-apps)

###### Note

Access and refresh tokens expire in 1 hour. For
information on regenerating tokens, see [Atlassian Developer\
Documentation](https://developer.atlassian.com/cloud/jira/platform/oauth-2-3lo-apps).

Personal Access Token

- Personal Access Token

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes. For more information, see [ACL crawling](confluence-server-user-management.md).**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**Yes**Crawl features****Custom metadata**Yes**Entities**Yes. The following entities are supported:

- Space

- Page

- Blog post

- Comment

- Attachment

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. For more information, see [Field mappings](confluence-server-field-mappings.md).**Filters**Yes. The following filters are supported:

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Known limitations

Prerequisites for connecting Amazon Q to Confluence (Server/Data Center)

All content copied from https://docs.aws.amazon.com/.
