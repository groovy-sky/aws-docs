# Document history for Amazon Q Developer User Guide

The following table describes the document history for the _Amazon Q Developer User_
_Guide_. For notifications about updates to this documentation, you can subscribe to the RSS
feed.

ChangeDescriptionDate

Updated Builder ID unsubscribe instructions

Added guidance for unsubscribing a personal account (Builder ID) when the Amazon Q Developer Profile has been deleted, including steps to re-enable via the console and a link to create a Kiro Profile.

March 19, 2026

Updated account limit for subscriptions

You can now subscribe users in a maximum of 20 accounts per AWS Region, per organization, increased from 10.

February 18, 2026

Deprecated code transformation for GitLab

Code transformation capabilities for Java modernization have been deprecated for [GitLab Duo with Amazon Q Developer](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/gitlab-with-amazon-q.html). The `/q transform` quick action and CI/CD pipeline customization documentation have been removed.

January 16, 2026

Removed deprecated GitLab quick actions

Removed the `/q dev (revise)` bullet point and the entire Unit test generation section with `/q test` command from the [GitLab quick actions](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/gitlab-concepts.html#gitlab-concepts-quick-actions) documentation.

January 13, 2026

Updated GitHub feature development workflow

Updated the GitHub feature development workflow to use conversation-based feedback with the `/q` command instead of the traditional GitHub review process. Users can now provide feedback through natural language commands in pull request conversations.

December 29, 2025

Updated GitHub app installation registration steps

Updated the [Increasing usage limits and configuring details in Amazon Q Developer console](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/github-register-app-install.html)
topic to include steps for navigating the Kiro and Amazon Q Developer shared console.

December 18, 2025

Deprecated code transformation for GitHub

Code transformation capabilities for Java modernization have been deprecated for [Amazon Q Developer for GitHub](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/amazon-q-for-github.html).

December 15, 2025

Q CLI becomes Kiro CLI

The Q CLI has become the Kiro CLI.

November 17, 2025

Updated managed policies: AmazonQFullAccess, AmazonQDeveloperAccess, and AWSServiceRoleForUserSubscriptions

Additional permissions have been added to the [AmazonQFullAccess](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-fullaccess) policy, [AmazonQDeveloperAccess](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-developeraccess) policy, and [AWSServiceRoleForUserSubscriptions](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-AWSServiceRoleForUserSubscriptions).

October 29, 2025

Removed IDE agent documentation

[Agentic chat\
capabilities](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-in-IDE-chat.html#example-tasks) have replaced the /dev, /doc/, /test, and /review agents in the IDE. There are updated steps for [code reviews in the IDE](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/code-reviews.html).

October 21, 2025

Placeholder title

Placeholder content.

October 21, 2025

Added IDE memory bank documentation

Added documentation for
[generating a\
memory bank for Amazon Q chat](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/context-memory-bank.html).

October 21, 2025

Added Q CLI feature documentation

Added documentation for [creating agents with AI assistance](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-agent-generate.html), [prompt management](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-prompts.html), [responding to messages](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-reply-command.html), and [project rules](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-project-rules.html).

October 3, 2025

Enhanced context management documentation

Added comprehensive guidance for [choosing the right context approach](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-context.html#command-line-context-choosing-approach) with comparison table, decision flowchart, and best practices for agent resources, session context, and knowledge bases.

September 23, 2025

Added Q CLI v1.16.0 feature documentation

Added documentation for [agent default behavior](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-agents-default-behavior.html) with priority-based selection, [legacy MCP configuration support](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-agents-legacy-mcp.html), new chat commands ( `/tangent tail`, `/changelog`), and [agent-related settings](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-settings-agent-config.html).

September 18, 2025

Context hooks deprecated

Context hooks are deprecated in favor of [agent hooks](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-custom-agents-configuration.html#command-line-agent-hooks). Existing configurations are automatically migrated to agent files.

September 17, 2025

Added remote MCP server support and built-in tools documentation

Added documentation for [remote MCP servers](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-mcp-config-CLI.html#command-line-mcp-remote-servers) with OAuth authentication and comprehensive built-in tools configuration.

September 17, 2025

Removed Amazon S3 buckets from allowlist table

You no longer need to
[allowlist Amazon S3\
buckets](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/firewall.html#data-perimeters) for software development or unit test
generation in the IDE.

September 9, 2025

New command line tool for transformation version

The most recent [command line tool\
for transformation](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform-CLI-versions.html) version includes bug fixes.

September 9, 2025

Updated Amazon Q Developer for GitHub review agent content

[Amazon Q Developer \
for GitHub](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/amazon-q-for-github) code reviews include a code review summary with threaded findings. You can
interact with Amazon Q Developer in pull request comments about the findings.

September 2, 2025

Added prerequisites to Amazon Q Developer for GitHub review section

[Amazon Q Developer \
for GitHub](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/amazon-q-for-github) code reviews can only be initiated with Write, Maintain, or Admin role in GitHub.

September 2, 2025

Disabling MCP

You can [disable MCP for your organization](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/qdev-mcp.html#disabling-mcp-for-orgs).

August 21, 2025

Updated proxy topic

Updated the [Configuring a corporate proxy in\
Amazon Q](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/firewall.html#corp-proxy) topic to indicate that Proxy Auto-Configuration (PAC) files are not supported.

August 19, 2025

Updated the deployment options topic

Updated the [Deployment options](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/deployment-options.html) topic
to indicate that you can subscribe users in a maximum of 10 accounts per AWS Region, per
organization.

August 15, 2025

Added troubleshooting section and troubleshooting logs section

Added comprehensive [troubleshooting](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-troubleshooting.html)
documentation and detailed [log access and\
analysis](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/troubleshooting-q-logs.html) guide to help users diagnose and resolve issues with Amazon Q Developer features
and functionality.

August 15, 2025

Added command line transformation job history section

Added documentation for the [`qct history`\
command](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/cli-job-history.html) that allows users to view their transformation job history from the command
line.

August 7, 2025

Updated a subscription topic

Updated the [Viewing an aggregated list of\
Amazon Q Developer subscriptions](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/subscribe-visibility.html) topic with new instructions and corrections.

August 6, 2025

Added IDE transformation job history topic

Added documentation for the [transformation job\
history](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transformation-job-history.html) feature, which allows users to view, manage, and retrieve artifacts from
their recent Java transformation jobs in Visual Studio Code.

August 6, 2025

GitLab self-managed instance update

[GitLab Duo with\
Amazon Q](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/ide-chat-history-compaction.html) configuration requires self-managed instances with [GitLab\
17.11.0](https://docs.gitlab.com/update/versions/gitlab_17_changes) or later.

August 5, 2025

Added a troubleshooting topic

Added a [Can't see subscribed\
users](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/ide-chat-history-compaction.html) troubleshooting topic.

August 1, 2025

Added steps for code reviews with agentic chat in VSC

Added procedures for [running a code\
review](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/start-review.html#project-review) and [addressing\
code issues](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/address-issues-jetbrains-visualstudiocode.html) using agentic chat in Visual Studio Code.

July 31, 2025

Added model selection for IDE chat

Added a [Selecting models](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-in-IDE-chat-models.html) topic
with available models and instructions for selecting a model for chat in the IDE.

July 31, 2025

Added keyboard shortcuts topic

Added a [Keyboard shortcuts\
for Amazon Q Developer](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-in-ides-chat-keyboard-shortcuts.html) topic that provides a comprehensive list of keyboard shortcuts
available across different IDEs and platforms.

July 31, 2025

Added custom agents feature

Added [custom agents](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-custom-agents.html)
feature that allows you to customize Amazon Q Developer CLI for specific workflows by configuring which
tools, permissions, and context are available for different use cases.

July 31, 2025

Added chat history compaction topic

Added a [Chat history compaction\
in Amazon Q Developer](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/ide-chat-history-compaction.html) topic that explains how to manage context window limits when chatting
with Amazon Q Developer in your IDE.

July 30, 2025

Added a profile deletion topic

Added a [Deleting the Amazon Q Developer\
profile](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/subscribe-delete-profile.html) topic.

July 24, 2025

Added opt out steps for telemetry collection

Added instructions for [how to opt out of\
telemetry collection](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/opt-out-IDE.html#opt-out-IDE-telemetry) while using the Amazon Q command line tool for transformation.

July 21, 2025

Added GitHub workflow for first-party dependencies content

Updated [Customizing a\
workflow for code transformation](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/github-code-transformation-workflow.html) section with new [GitHub workflow for handling first-party dependencies](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/github-code-transformation-workflow-1p-dependencies.html) page.

July 16, 2025

Updated the user activity metrics topic

Updated the [User activity report\
metrics](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/user-activity-metrics.html) topic with additional metrics related to the [inline chat](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-in-IDE-inline-chat.html) feature.

July 15, 2025

Updated the proxy configuration section

Updated the [Configuring a corporate proxy in\
Amazon Q](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/firewall.html#corp-proxy) to include instructions for Eclipse, JetBrains, and Visual Studio.

July 15, 2025

Image context support added to IDE chat

Amazon Q now supports using [images as\
context](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/ide-chat-context.html#image-context) in the IDE chat panel. This feature allows users to include images when
prompting, enabling scenarios such as generating code from UI mockups or sequence
diagrams.

July 9, 2025

Updated the Builder ID upgrade section

Added a note to the [Upgrading a personal account\
(Builder ID)](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/upgrade-to-pro.html) section to indicate that you cannot link multiple Builder IDs to a single AWS
account.

July 7, 2025

Updated the profile sharing section

The [Enabling profile\
sharing](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-profile-sharing.html) section incorrectly indicated that the profile settings in the management
account would override those in member accounts when profile sharing was enabled. This
statement has been removed, and the purpose of profile sharing has been clarified.

July 2, 2025

Added slash commands details to Amazon Q Developer for GitHub content

[Amazon Q Developer for GitHub](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/amazon-q-for-github.html) can be invoked with [slash\
commands](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/amazon-q-for-github.html#github-slash-commands) to perform feature development, code reviews, and code transformation, as
well invoke for help, which provides a link to to the Amazon Q Developer for GitHub content.

June 30, 2025

Updated details for QCT CLI in IDEs

Added details about maintaining security, pausing or canceling code transformation, new
parameter for `qct transform` command, and LLM summary output in the [IDE code\
transformation](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform-in-IDE.html) content. QCT CLI also now supports
`https://q.eu-central-1.amazonaws.com/`.

June 26, 2025

New content for creating project rules in third-party platforms

You can [create project\
rules in GitLab and GitHub](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/third-party-context-project-rules.html) for standards and best practices.

June 26, 2025

Updated KMS policy for third-party integration

Added `"kms:Decrypt"`, `"kms:DescribeKey"`,
`"kms:Encrypt"`, `"kms:GenerateDataKey"` to [KMS policy for\
managed access to Amazon Q Developer for third-party integration](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/third-party-context-project-rules.html).

June 25, 2025

Added proxy configuration instructions

Added a [Configuring a corporate\
proxy in Amazon Q for Visual Studio Code](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/firewall.html#vsc-corp-proxy) section.

June 25, 2025

Updated user activity reports

[User activity reports](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-user-telemetry.html) are now generated at midnight UTC, and broken into several
files with 1,000 users in each.

June 24, 2025

Updated the Amazon S3 bucket section

Updated the [Amazon S3 bucket URLs and ARNs\
to allowlist](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/firewall.html#data-perimeters) section to indicate version 1.72.0 and later of the Visual Studio Code plugin for
Amazon Q does not require allowlisting of buckets.

June 24, 2025

Rebranded Amazon Q Developer operational investigations

Updated various sections to reflect the rebranding of _Amazon Q Developer operational_
_investigations_ to _CloudWatch investigations_.

June 24, 2025

Updated GitHub feature development and code transformation details

[Amazon Q Developer for GitHub](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/amazon-q-for-github.html) can now be invoked to perform [feature\
development](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/github-feature-development.html) and [transform code](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/github-code-transformation.html)
using slash commands in comments within an issue.

June 19, 2025

Rules button added to VS Code IDE chat

Added Rules button to the chat interface for [creating and managing project\
rules](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/context-project-rules.html) through the UI.

June 18, 2025

Pinned context support added to VS Code IDE chat

Added [context pinning](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/context-pinning.html) functionality in VS Code to maintain selected context items across
chat interactions, with support for manual pinning and system-added context management.

June 18, 2025

Updated details for code transformation in IDEs

Added [code transformation](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform-code.html) details about source code version and target code version, and
additional transformation requirements to upgrade project libraries and dependencies. Removed
details about multiple diffs.

June 17, 2025

Chat with Amazon Q about network security

You can [chat about network\
security](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/chat-network-security.html) with Amazon Q in the console and in chat applications.

June 17, 2025

Added new Data storage topic

Added the [Data storage in Amazon Q Developer](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/data-storage.html) topic with information about where content is
stored.

June 13, 2025

MCP in the IDE

MCP servers can now be used [with Amazon Q Developer in\
the IDE](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-ide-configuration.html).

June 12, 2025

Personal account (Builder ID) upgrade

Users with personal accounts (Builder IDs) can now [upgrade to the Pro\
tier](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/upgrade-to-pro.html#upgrade-builder-id).

June 11, 2025

Updates to subscription statuses

Updated the description of the **Unavailable** subscription status in
[Amazon Q Developer\
subscription statuses](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-setup-subscribe-status.html).

June 9, 2025

Support for specifying dependencies for transformations on the CLI

You can [provide a dependency upgrade file](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/run-CLI-transformations.html#step-3-dependency-upgrade-file) and modify the transformation plan for
Java upgrades with the command line tool for transformations.

June 9, 2025

Expanded command line reference

A new [command reference\
section](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-reference.html) provides more information on the arguments you can use to invoke the
Q CLI.

June 9, 2025

Model selection for chat on the command line

You can now [select the model](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-chat-models.html)
you want Amazon Q to use for chat sessions on the command line.

June 5, 2025

Chat context support expanded to JetBrains, Visual Studio, and Eclipse IDEs

Amazon Q now supports [contexts](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/ide-chat-context.html) in the JetBrains,
Visual Studio, and Eclipse IDE chat panel.

June 5, 2025

Agentic chat in all supported IDEs

Agentic chat is available in [all supported IDEs](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-in-IDE-chat.html#ide-agentic).

June 5, 2025

Updated GitHub code review details

[Amazon Q Developer for GitHub](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/github-code-reviews.html) can now perform code reviews within GitHub pull requests. They
can be initiated with the `/q review` slash command in a new comment.

June 2, 2025

Cost optimization capabilities in chat

In addition to cost analysis, you can chat with Amazon Q to get AWS [cost optimization\
insights](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/chat-costs.html).

June 2, 2025

Updated the firewall page

Updated and added URLs on the [Configuring a firewall, proxy server, or\
data perimeter for Amazon Q Developer Pro](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/firewall.html) page.

May 22, 2025

MCP configuration commands

Commands that you can give directly from your non-Q command line [to configure MCP servers for Amazon Q](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-mcp-configuration.html#command-line-mcp-reference-commands).

May 21, 2025

Updated the customization policy topic

The [Allow administrators to create customizations](id-based-policy-examples-admins.md#id-based-policy-examples-allow-customizations) topic now describes permission errors
that may occur when you create the customization policy.

May 16, 2025

Updated troubleshooting tasks

The [Troubleshooting Amazon Q Developer Pro subscriptions](q-admin-setup-subscribe-troubleshooting.md) topic now includes more information.

May 15, 2025

Updated new code review behavior for GitLab Duo and GitHub

Automatic code reviews have been updated so they are triggered in new or reopened [GitLab](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/gitlab-concepts.html#gitlab-concepts-quick-actions) merge requests and [GitHub](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/github-code-reviews.html) pull requests. Code
reviews are not triggered by subsequent commits.

May 15, 2025

Updated VPC endpoints

Updated the VPC endpoints on the [Amazon Q Developer and interface\
endpoints (AWS PrivateLink)](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/vpc-interface-endpoints.html) page.

May 15, 2025

Improved MCP server loading

MCP servers now load in the background, allowing you to start interacting with Amazon Q
immediately without waiting for all servers to initialize.

May 15, 2025

Added image support in chat

Amazon Q can now analyze and discuss images directly in your chat session using the
`fs_read` tool with the Image mode.

May 15, 2025

Added conversation persistence

Amazon Q automatically remembers your conversations based on the directory where they take
place, and provides `/export` and `/import` commands for manually
managing conversation state.

May 15, 2025

Updated managed policies: AmazonQFullAccess and AmazonQDeveloperAccess

Additional permissions have been added to the [AmazonQFullAccess](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-fullaccess) policy and [AmazonQDeveloperAccess](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-developeraccess) policy to manage conversation history.

May 14, 2025

Support for multiple conversations in console chat

You can [save and\
switch between conversations](chat-with-q.md#manage-conversations-console) with Amazon Q in the AWS console.

May 14, 2025

Added subscription details

The [Subscribe users to Amazon Q Developer Pro across accounts](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/tracking-across-org.html) topic now includes information about
where to install IAM Identity Center and the Amazon Q Developer profile.

May 13, 2025

Transformation website removed

The Amazon Q Developer transformation website has been taken down, and with it, the related
documentation.

May 12, 2025

Updated terminology

Updated the term _identity-aware console sessions_ to
_identity-enhanced console sessions_ throughout this guide.

May 10, 2025

Example SCP added

This [service control policy (SCP)](security-iam-manage-access-with-policies.md#example-scp-deny-q-outside-eu) denies access to Amazon Q outside of EU regions.

May 8, 2025

Amazon Q Developer for GitHub

Information about [Amazon Q Developer for GitHub](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/amazon-q-for-github.html),
including concepts and procedures focused on setting up, key features, and
configuration.

May 5, 2025

Added context hooks

Added support for [context hooks](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-context-hooks.html).

May 3, 2025

Managed policy updated

Permission has been added to [AmazonQFullAccess](managed-policy.md#amazonq-policy-fullaccess).

May 2, 2025

Agentic chat in the IDE

Agentic chat functionality is available [in the IDE](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-in-IDE-chat.html#ide-agentic).

May 1, 2025

Updates to customizations

Customizations now support [additional languages](q-language-ide-support.md#customization-language-support).

April 30, 2025

Managed and example policies update

Permissions have been added to [AmazonQFullAccess](managed-policy.md#amazonq-policy-fullaccess), [GitLabDuoWithAmazonQPermissionsPolicy](managed-policy.md#amazonq-policy-GitLabDuoWithAmazonQPermissionsPolicy), and [Allow administrators to configure plugins](id-based-policy-examples-admins.md#id-based-policy-examples-admin-plugins).

April 30, 2025

Added MCP support

Added support for [MCP in the CLI](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-mcp.html).

April 29, 2025

Added upgrade details

Added information about how to [upgrade from the\
Free tier to the Pro tier](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-free-tier.html#upgrading-to-pro-tier).

April 28, 2025

Support for conversation history

Your [conversation history](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/ide-chat-conversation.html)
is now saved when you chat with Amazon Q in the IDE.

April 21, 2025

Support for code as context

You can now specify classes, functions, and global variables as [context](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/ide-chat-conversation.html) when you chat
with Amazon Q in the IDE.

April 21, 2025

Updates to GitLab Duo with Amazon Q onboarding and policy

[GitLab Duo with Amazon Q](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/gitlab-with-amazon-q.html) has been updated with changes to onboarding and permissions policy
( [GitLabDuoWithAmazonQPermissionsPolicy](managed-policy.md#amazonq-policy-GitLabDuoWithAmazonQPermissionsPolicy)).

April 16, 2025

Updated dashboard permissions

Updated the [list of\
permissions](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/dashboard-troubleshooting.html) required to view the Amazon Q Developer dashboard.

April 15, 2025

Improved security documentation for command line

Reorganized and enhanced [security documentation](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-chat.html#command-line-chat-security) with comprehensive guidance on security considerations, best
practices, and safe usage of tool permissions.

April 13, 2025

Enhanced command-line security and settings

Added a new [command-line settings](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-settings.html)
section with configuration options. Enhanced [tool\
permissions documentation](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-chat.html#command-line-chat-tools.html) with security best practices for sensitive
environments.

April 12, 2025

Update to the subscription experience

The workflow to [subscribe users to\
Amazon Q Developer Pro and install the Amazon Q Developer profile](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-pro-tier-setting-up-access.html) has been moved from the Amazon Q console to
the Amazon Q Developer console.

April 10, 2025

Inline chat is available in Eclipse

You can [chat inline](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-in-IDE-inline-chat.html) with Amazon Q
in Eclipse.

April 10, 2025

Amazon Q Developer profiles are available in Europe (Frankfurt)

When you subscribe to Amazon Q Developer, you can create profiles [in the\
Europe (Frankfurt) Region](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-setup-subscribe-regions.html).

April 10, 2025

/tools feature added to CLI

You can use [the\
`/tools` command](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-chat.html#command-line-chat-tools.html) to manage permissions for tools that Amazon Q uses to
perform actions on your system.

April 10, 2025

Support for natural languages other than English

You can chat with Amazon Q [in the IDE](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-in-IDE.html) and [on the command\
line](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line.html).

April 9, 2025

Updates to GitLab Duo with Amazon Q

[GitLab Duo with Amazon Q](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/gitlab-with-amazon-q.html) has been updated regarding changes to inline policy, and you can
optionally create a CMK policy. The `/fix` feature has been removed.

April 8, 2025

Email notifications for transformations

You may receive [email\
notifications](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform.html#transform-app-emails) for updates related to your transformations.

April 8, 2025

New context, prompt, and project rules topics

The [Adding context to the chat](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/ide-chat-context.html), [Saving prompts](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/context-prompt-library.html), and
[Creating project rules](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/context-project-rules.html) topics have been added.

April 4, 2025

Updates to subscriptions topics

The [Understanding\
subscriptions](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-setup-subscribe-understanding.html), [Viewing aggregated\
subscriptions](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/subscribe-visibility.html), and [Enabling profile\
sharing](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-profile-sharing.html) topics have been corrected.

March 25, 2025

Example policy update

The example policies in [Allow administrators to use the Amazon Q console](id-based-policy-examples-admins.md#q-admin-setup-admin-users-sub) and [Allow administrators to use the Amazon Q Developer console](id-based-policy-examples-admins.md#q-admin-setup-admin-users) have been updated with the
`sso:CreateInstance` permission.

March 24, 2025

Support for C++ and C# in customizations

[Customizations](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/customizations.html) now support C++ and C#.

March 20, 2025

Updates to chatting about resources

You can [chat with Amazon Q about multiple AWS resources and services](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/chat-actions.html) to get answers about
your AWS infrastructure and configurations.

March 13, 2025

Additional language support for documentation generation

The agent for documentation generation now supports [C++\
and C#](q-language-ide-support.md#doc-gen-language-support).

March 12, 2025

New subscription-related limit

Updated the [Subscribing users to\
Amazon Q Developer Pro](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-pro-tier-setting-up-access.html) topic to indicate that you can enable Amazon Q Developer in a maximum of 50
AWS accounts within an organization managed by AWS Organizations.

March 6, 2025

Context integration to CLI chat

Amazon Q CLI now has [context integration](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-chat.html#command-line-chat-context-integration.html), giving Amazon Q enhanced understanding of use cases and enabling it
to provide more relevant and context-aware responses.

March 6, 2025

Policy correction

A JSON sytax error has been corrected in the policy described in [Allow administrators to use the Amazon Q console](id-based-policy-examples-admins.md#q-admin-setup-admin-users-sub).

February 28, 2025

New version of the command line tool for transformation

The [latest version](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform-CLI-versions.html) of the
command line tool for transformation includes support for authenticating with IAM through the
AWS CLI.

February 28, 2025

Upgrading to the Pro tier

Added information about how to upgrade to the Pro tier in the [Amazon Q Developer Free tier](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-free-tier.html) topic.

February 25, 2025

Customizations policy update

A permission has been added to the [customizations policy](id-based-policy-examples-admins.md#id-based-policy-examples-allow-customizations).

February 25, 2025

New dashboard topic

The following topic has been added: [Descriptions of\
Amazon Q Developer dashboard usage metrics](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/dashboard-metrics-descriptions.html).

February 21, 2025

New cross-region processing topic

The [cross-region processing\
topic](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/cross-region-processing.html) describes how Amazon Q Developer processes requests and makes calls across
AWS Regions to provide the service.

February 21, 2025

Managed policy update

Permissions have been added to [AWSServiceRoleForUserSubscriptions](managed-policy.md#amazonq-policy-AWSServiceRoleForUserSubscriptions).

February 21, 2025

/doc enhancement

Amazon Q can now [generate infrastructure\
diagrams](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/doc-generation.html#use-cases) in response to a `/doc` command.

February 20, 2025

New subscription topics

Two subscription-related topics were added: [Amazon Q Developer\
subscription statuses](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-setup-subscribe-status.html) and [Viewing aggregated Amazon Q Developer\
subscriptions](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/subscribe-visibility.html).

February 19, 2025

Amazon Q Developer in chat applications chapter

Amazon Q Developer in chat applications is now Amazon Q Developer in chat applications. A [new chapter](q-in-chat-applications.md) describes
the supported features.

February 19, 2025

Support for Java 21 transformations

You can upgrade Java applications [to Java\
21](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/code-transformation.html#supported-languages-IDEs) in the IDE and [on the command line](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform-CLI.html).

February 14, 2025

New firewall topic

A [Configuring a\
firewall or proxy server for Amazon Q Developer](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/firewall.html) topic has been added.

February 14, 2025

New version of the command line tool for transformation

The [latest version](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform-CLI-versions.html) of the
command line tool for transformation includes support for converting embedded SQL in Java
applications.

February 12, 2025

User activity report correction

The path to the [user\
activity report](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-manage-settings.html#q-admin-user-telemetry) CSV file has been corrected.

February 10, 2025

Update to the retention period of transformed code

Amazon Q now retains [transformed code](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform-in-IDE.html) for 30
days, up from 24 hours.

February 7, 2025

New subscription workflow

The steps to [subscribe users to\
Amazon Q Developer](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-pro-tier-setting-up-access.html) have been improved.

February 6, 2025

New version of the command line tool for transformation

The [latest version](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform-CLI-versions.html) of the
command line for transformation includes the ability to receive your upgraded Java code in
multiple commits.

February 3, 2025

/dev enhancement

Amazon Q can now [test\
the code it generates](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/devfile.html) in response to a `/dev` command.

January 31, 2025

Customizations section update

The [Creating your\
customization](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/customizations-admin-customize.html) topic now indicates you can include any number of repositories in your
customization.

January 24, 2025

Prompt logging examples

The [Enabling prompt\
logging](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-prompt-logging.html) section now includes [example logs](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-prompt-log-examples.html).

January 23, 2025

CloudZero plugin

The [CloudZero plugin](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/cloudzero-plugin.html) is available in Amazon Q chat.

January 15, 2025

User activity report update

[New\
metrics](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/user-activity-metrics.html) have been added to [User\
activity reports](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-manage-settings.html#q-admin-user-telemetry).

December 16, 2024

Dashboard update

Information about the old dashboard has been removed from the [Amazon Q Developer Pro dashboard](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/dashboard.html) section.
Information about filters and metrics has been added.

December 16, 2024

Troubleshooting with Amazon Q

An [Asking Amazon Q to\
troubleshoot your resources](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/chat-actions-troubleshooting.html) section has been added.

December 13, 2024

Identity-enhanced sessions update

The instructions for enabling identity-enhanced console sessions have been clarified in the
[Subscribing users\
to the Amazon Q Developer Pro tier with an organization instance](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-pro-tier-setting-up-access-org.html) section.

December 6, 2024

New test generation agent

You can use Amazon Q [test generation](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/test-generation.html) feature to
generate unit tests.

December 3, 2024

Large-scale transformation

Amazon Q can [transform](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform.html) .NET, mainframe, and VMware workloads in bulk.

December 3, 2024

GitLab Duo with Amazon Q

Information about [GitLab Duo with Amazon Q](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/gitlab-with-amazon-q.html),
including concepts, getting started procedures, and troublehsooting.

December 3, 2024

Documentation generation in the IDE

Amazon Q can [generate READMEs for your\
code](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/doc-generation.html) in supported IDEs.

December 3, 2024

Code reviews in the IDE

Amazon Q code reviews, previously security scans, can [detect and address issues in your\
code](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/code-reviews.html) in supported IDEs.

December 3, 2024

.NET transformation in the IDE

Amazon Q can [port your .NET\
applications](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform-dotnet-IDE.html) to Linux-compatible cross-platform applications in Visual Studio, available in
preview.

December 3, 2024

Transformation on the command line

You can transform Java applications [on the command line](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform-CLI.html), available
in preview.

November 27, 2024

Multiple diffs for transformation in the IDE

You can choose to receive transformation changes from Amazon Q [in multiple\
diffs](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/how-CT-works.html#review-plan-accept-changes).

November 27, 2024

Amazon Q in Eclipse

The [Amazon Q plugin](q-in-ide-setup.md#setup-eclipse)
is available in preview in Eclipse.

November 27, 2024

Cost analysis

The [cost\
analysis](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/chat-costs.html) capability, previously available in preview, is now generally
available.

November 26, 2024

Transformation for embedded SQL code

You can convert [embedded SQL code in your Java\
applications](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform-sql.html) with Amazon Q transformation in the IDE.

November 22, 2024

Dashboard update

The [Amazon Q Developer Pro\
dashboard](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/dashboard.html) has been update with new metrics.

November 22, 2024

CodeConnections repositories

When [creating a customization](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/customizations-admin-customize.html#customizations-admin-connect) using a CodeConnections connection, you can now choose the
repositories you want to use.

November 22, 2024

Amazon Q command line now supports Linux

[Amazon Q\
command line](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line.html) supports Linux environments. It supports Ubuntu 22 and 24, and may
otherwise work with GNOME v42+ or environments where the display server is Xorg and the input
method framework is IBus.

November 21, 2024

Subscribing users

The instructions for subscribing users in [Setting up access to\
the Amazon Q Developer Pro tier](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-pro-tier-setting-up-access.html) have been updated to reflect new user interface (UI)
elements.

November 20, 2024

Changes to customizations

The [Customization in chat](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/customizations.html) feature is now generally available. Also, customizations can
now be created with the following file types: `.md`,
`.mdx`, `.rst`, and `.txt`.

November 20, 2024

Supported IAM Identity Center Regions

A section has been added with information about the [Regions\
where you can set up IAM Identity Center instances](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-pro-tier-setting-up-access.html#idc-regions) for Amazon Q Developer Pro subscriptions.

November 18, 2024

Languages added

[Support](q-language-ide-support.md) has been added for Dart, Lua, R, Swift, SystemVerilog, and Powershell, as
well as expanded support for JSON and YAML.

November 18, 2024

Customer managed key support

Information about using customer managed keys and the features that can be encrypted with
them has been added to the [Data encryption](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/data-encryption.html)
topic.

November 18, 2024

Cross-region inference

A topic on [cross-region inference in\
Amazon Q Developer](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/cross-region-inference.html) has been added.

November 18, 2024

Amazon Q Developer Pro quotas

A [Pro tier\
quotas](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/quotas.html) section has been added.

November 18, 2024

Updated managed policy: AmazonQFullAccess

Additional permissions have been added to the [AmazonQFullAccess](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-fullaccess) policy.

November 13, 2024

Updated managed policy: AmazonQDeveloperAccess

Additional permissions have been added to the [AmazonQDeveloperAccess](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-developeraccess) policy.

November 13, 2024

Amazon Q plugins

[Plugins](plugins.md)
enable users to chat with Amazon Q about metrics provided by third party tools.

November 13, 2024

User activity reports

You can now [enable user activity reports](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-manage-settings.html#q-admin-user-telemetry).

November 8, 2024

Customizations section update

The [Preparing your data](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/customizations-prereq.html#customizations-prereq-data) section now describes file and directory naming
limitations.

November 5, 2024

Clarified the Amazon Q Developer Pro section

The instructions for [subscribing users to\
Amazon Q Developer Pro](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-pro-tier-setting-up-access.html) have been clarified.

November 1, 2024

Inline chat

You can transform code using the new [inline chat](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-in-IDE-inline-chat.html)
feature.

October 29, 2024

Updated managed policies: AmazonQFullAccess and AmazonQDeveloperAccess

Additional permissions have been added to the [AmazonQFullAccess](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-fullaccess) policy and [AmazonQDeveloperAccess](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-developeraccess) policy.

October 28, 2024

Customizations section correction

The [Creating your\
customization](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/customizations-admin-customize.html) section now indicates that your codebase must reside in a folder in
Amazon S3, not the bucket's root.

October 28, 2024

Prompt logging section clarification

The [Enabling prompt logging](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-manage-settings.html#q-admin-prompt-logging) section's wording was clarified.

October 24, 2024

Amazon S3 bucket policy fix

The Amazon S3 bucket policy shown in [Enabling prompt logging](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-manage-settings.html#q-admin-prompt-logging) contained a JSON syntax error that was fixed.

October 22, 2024

Expanded features chapter

The chapter [describing various Amazon Q Developer features](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/features.html) has been significantly expanded.

October 3, 2024

Console-to-Code

Console-to-Code, previously available in preview as a feature of Amazon EC2, [is now generally available](console-to-code.md)
as a feature of Amazon Q Developer. It integrates with Amazon EC2, Amazon VPC, and Amazon RDS.

October 3, 2024

New policy: Use Amazon Q CLI with AWS CloudShell

Identity-based policy [allows\
users to use Amazon Q CLI with AWS CloudShell](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/id-based-policy-examples-allow-cli-cloudshell.html).

October 2, 2024

Prompt logging

You can [log your users' IDE\
prompts](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-prompt-logging.html) in an Amazon S3 bucket.

September 16, 2024

Setup content updated

The [Getting started](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/getting-started-q-dev.html)
chapter has been significantly simplified and restructured.

August 15, 2024

CodeWhisperer endpoint needed for IDE VPC access

[Access\
from a Amazon VPC](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/vpc-interface-endpoints.html#vpc-endpoint-create) must include both `q` and `codewhisperer`
endpoints.

July 18, 2024

New endpoint

Endpoints can now [use the\
string `q`](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/vpc-interface-endpoints.html#vpc-endpoint-create) instead of `codewhisperer`.

July 12, 2024

Customizations are GA

The [customizations](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/customizations.html) feature is generally available.

July 10, 2024

Chatting about customizations (Preview)

In Preview, you can use the [customizations](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/customizations.html) feature to ask
questions related to your codebase.

July 10, 2024

Updated managed policy: AmazonQFullAccess

Additional permissions have been added to the [AmazonQFullAccess](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-fullaccess) policy.

July 9, 2024

New managed policy: AmazonQDeveloperAccess

The [AmazonQDeveloperAccess](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-developeraccess) managed policy provides full access to enable interactions
with Amazon Q Developer, without administrator access.

July 9, 2024

Updated Amazon Q Developer admin policy

The [policy for empowering\
Amazon Q Developer administrators has been updated to include `sso:ListProfiles`.](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-setup-admin-users.html)

June 19, 2024

Trusted access section

A [new section](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-profile-sharing.html) more
clearly explains how a Amazon Q Developer administrator can share settings with member accounts.

June 19, 2024

Updated setup procedures

There's an improved [Getting started](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/getting-started-q-dev.html)
chapter that includes support for [account instances](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/admin-setup-q-dev.html#access-options-account-instances).

June 6, 2024

Updated code examples

The [code\
examples](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/inline-suggestions-code-examples.html) now include C and C++, and have improved examples for C#.

June 6, 2024

Updated managed policy: AmazonQFullAccess

Additional permissions have been added to the [AmazonQFullAccess](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-fullaccess) policy.

April 30, 2024

New service-linked role: AWSServiceRoleForUserSubscriptions

The [AWSServiceRoleForUserSubscriptions](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/using-service-linked-roles-user-subs.html#service-linked-role-permissions-user-subs) service-linked role provides access for User
Subscriptions to your IAM Identity Center resources to automatically update your subscriptions.

April 30, 2024

New service-linked role: AWSServiceRoleForAmazonQDeveloper

The [AWSServiceRoleForAmazonQDeveloper](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/using-service-linked-roles-qdev.html) service-linked role grants permission to access
and emit data, and to create reports.

April 30, 2024

New managed policy: AWSServiceRoleForUserSubscriptionPolicy

The [AWSServiceRoleForUserSubscriptionPolicy](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-AWSServiceRoleForUserSubscriptionPolicy) allows principals to track IAM Identity Center directory
and AWS Organizations changes.

April 30, 2024

New managed policy: AWSServiceRoleForAmazonQDeveloperPolicy

The [AWSServiceRoleForAmazonQDeveloperPolicy](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/managed-policy.html#amazonq-policy-AWSServiceRoleForAmazonQDeveloperPolicy) allows Amazon Q Developer to call CloudWatch and
CodeGuru on your behalf.

April 30, 2024

GA release

Amazon Q Developer is available for general audiences.

April 30, 2024

Amazon CodeWhisperer merge

Amazon CodeWhisperer is now a part of Amazon Q Developer.

April 30, 2024

New guide name

This service and accompanying user guide have been renamed Amazon Q Developer.

March 29, 2024

New permission

The [ListConversations action](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/security_iam_permissions.html#qdev-permissions) is required to chat with Amazon Q in the console.

March 5, 2024

New data protection topic

Amazon Q now uses content for [service improvement\
purposes](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/service-improvement.html).

January 25, 2024

New topic

Added instructions for how to [add Amazon Q to Slack and\
Microsoft Teams channels](https://docs.aws.amazon.com/amazonq/latest/aws-builder-use-ug/q-in-chat-applications.html) that are configured with Amazon Q Developer in chat applications.

January 18, 2024

Preview release

This is the initial preview release of the _Amazon Q Developer User_
_Guide_.

November 28, 2023

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Q Developer service rename
