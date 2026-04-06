# What is a document?

When you connect Amazon Q Business to a data source, what Amazon Q Business considers—and crawls—as a document varies by
connector.

The following table outlines what each connector crawls as a document.

Data source connectorSupports crawlingDocument definitionAdobe Experience Manager (Cloud and Server)

- Assets

- Pages

- Each Asset is considered a single document.

- Each Page is considered a single document.

Alfresco (Cloud and Server)

- Files

- Comments

- Each File is considered a single document.

- Each Comment is considered a single document.

Amazon FSx (Windows)Files

Each File is considered a single document.

Amazon S3
Objects

Each Object is considered a single document. Any
`object-name.metadata.json` file
and access control list (ACL) file is considered metadata for
the object it is associated with and not treated as a separate
document.

Amazon Q Business Web Crawler

- Web pages

- Attachments

- Each Web page is considered a single document.

- Each Attachment is considered a single
document.

WorkDocs

- Files

- Comments

- Each File is considered a single document.

- Each Comment is considered a single document.

Box

- Files

- Tasks

- Comments

- Weblinks

- Each File is considered a single document.

- Each Task is considered a single document.

- Each Comment is considered a single document.

- Each Weblink is considered a single document.

Confluence (Cloud and Server)

- Spaces

- Pages

- Blogs

- Comments

- Attachments

- Each Space is considered a single document.

- Each Page is considered a single document.

- Each Blog is considered a single document.

- Each Comment is considered a single document.

- Each Attachment is considered a single
document.

Database data sources

- Aurora (MySQL)

- Aurora (PostgreSQL

- Amazon RDS (Microsoft SQL Server)

- Amazon RDS (MySQL)

- Amazon RDS (Oracle)

- Amazon RDS (PostgreSQL)

- IBM DB2

- PostgreSQL

- Microsoft SQL Server

- MySQL

- Oracle Database

- Table data in a single database

- View data in a single database

Each row in a table and view is considered a single
document.

Dropbox

- Files

- Papers

- Paper templates

- Shortcuts

- Each File is considered a single document.

- Each Paper is considered a single document.

- Each Paper template is considered a single
document.

- Each Shortcut is considered a single document.

Drupal

- Articles

- Basic pages

- Basic blocks

- Custom content

- Custom blocks

- Comments on articles, basic pages, basic blocks,
custom content, and custom blocks

- Attachments in articles, basic pages, basic blocks,
custom content, and custom blocks

- Each Article is considered a single document.

- Each Basic page is considered a single
document.

- Each Basic block is considered a single
document.

- Each Custom content is considered a single
document.

- Each Custom block is considered a single
document.

- Each Comment on an article, a basic page, a basic
block, any custom content, and a custom block is
considered a document.

- Each Attachment in an article, a basic page, a basic
block, any custom content, and a custom block is
considered a document.

GitHub (Cloud and Server)

- Respositories

- Repository commits

- Issues

- Issue attachments

- Issue comments

- Pull request documents

- Pull request comments

- Pull request attachments

- Each Repository is considered a single
document.

- Each Repository commit is considered a single
document.

- Each Issue is considered a single document.

- Each Issue attachment is considered a single
document.

- Each Issue comment is considered a single
document.

- Each Pull request is considered a single
document.

- Each Pull request comment is considered a single
document.

- Each Pull request attachment is considered a single
document.

Gmail

- Emails

- Email attachments

- Each Email is considered a single document.

- Each Email attachment is considered a single
document.

Google Calendar

- Calendar

Each calendar is considered a single document.

Google Drive

- Files

- Comments

- Each File is considered a single document.

- Each Comment is considered a single document.

Jira

- Projects

- Issues

- Comments

- Attachments

- Worklog

- Each Project is considered a single document.

- Each Issue is considered a single document.

- Each Comment is considered a single document.

- Each Attachment is considered a single
document.

- Each Worklog is considered a single document.

Microsoft Exchange

- Emails

- Attachments

- Calendar

- Contacts

- Notes

- OneNotes

- Each Email is considered a single document.

- Each Attachment is considered a single
document.

- Each Calendar is considered a single document.

- Each Contact is considered a single document.

- Each Note is considered a single document.

- Each page in OneNotes is considered a single
document.

Microsoft OneDrive

- Files

- OneNotes

- Each File is considered a single document.

- Each page in OneNotes is considered a single
document.

Microsoft SharePoint (Online and Server)

- Events

- Pages

- Files

- Links

- File attachments

- Comments

- OneNotes

- Each Event is considered a single document.

- Each Page is considered a single document.

- Each File is considered a single document.

- Each Link is considered a single document.

- Each File attachment is considered a single
document.

- Each Comment is considered a single document.

- Each page in OneNotes is considered a single
document.

Microsoft Teams

- Chat messages

- Chat attachments

- Channel posts

- Channel wikis

- Channel attachments

- Meeting chats

- Meeting files

- Meeting notes

- Calendar meetings

- OneNotes

- Each Chat message is considered a single
document.

- Each Chat attachment is considered a single
document.

- Each Channel post is considered a single
document.

- Each Channel wiki is considered a single
document.

- Each Channel attachment is considered a single
document.

- Each Meeting chat is considered a single
document.

- Each Meeting file is considered a single
document.

- Each Meeting note is considered a single
document.

- Each Calendar meeting is considered a single
document.

- Each page in OneNotes is considered a single
document.

Microsoft Yammer

- Communities

- Attachments

- Messages

- Users

- Each Community is considered a single document.

- Each Attachment is considered a single
document.

- Each Message and community post is considered a single
document.

- Each User is considered a single document.

Quip

- Files

- Messages

- Threads

- Each File is considered a single document.

- Each Comment is considered a single document.

- Each file and message posted in a Thread is considered
a single document.

Salesforce

- Accounts

- Contacts

- Campaigns

- Contracts

- Cases

- Partners

- Opportunities

- Groups

- Leads

- Users

- Tasks

- Ideas

- Profiles

- Solutions

- Chatters

- Documents

- Custom entities

- Knowledge articles

- Each Account is considered a single document.

- Each Contact is considered a single document.

- Each Campaign is considered a single document.

- Each Contract is considered a single document.

- Each Case is considered a single document.

- Each Partner is considered a single document.

- Each Opportunity is considered a single
document.

- Each Group is considered a single document.

- Each Lead is considered a single document.

- Each User is considered a single document.

- Each Task is considered a single document.

- Each Idea is considered a single document.

- Each Profile is considered a single document.

- Each Solution is considered a single document.

- Each Chatter is considered a single document.

- Each Document (file) is considered a single
document.

- Each Custom entity (record) is considered a single
document.

- Each Knowledge article is considered a single
document.

ServiceNow

- Incidents

- Knowledge articles

- Service catalog

- Attachments

- Each Incident is considered a single document.

- Each Knowledge article is considered a single
document.

- Each Service catalog is considered a single
document.

- Each Attachment is considered a single
document.

Slack

- Messages

- Message attachments

- Channel posts

- Each Message is considered a single document.

- Each Message attachment is considered a single
document.

- Each Channel post is considered a single
document.

Zendesk

- Tickets

- Ticket comments

- Ticket comment attachments

- Articles

- Article attachments

- Article comments

- Community topics

- Community posts

- Community post comments

- Each Ticket is considered a single document.

- Each Ticket comment is considered a single
document.

- Each Ticket comment attachment is considered a single
document.

- Each Article is considered a single document.

- Each Article attachment is considered a single
document.

- Each Article comment is considered a single
document.

- Each Community topic is considered a single
document.

- Each Community post is considered a single
document.

- Each Community post comment is considered a single
document.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Concepts

Configuration best practices
