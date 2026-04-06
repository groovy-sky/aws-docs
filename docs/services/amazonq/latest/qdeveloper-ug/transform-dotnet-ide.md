# Transforming .NET applications with Amazon Q Developer

Amazon Q Developer can port your Windows-based .NET applications to
Linux-compatible cross-platform .NET applications through a generative AI-powered
refactoring workflow. Amazon Q also helps you upgrade outdated versions of cross-platform .NET
applications to newer versions.

To transform a .NET solution or project, Amazon Q analyzes your codebase, determines the
necessary updates to port your application, and generates a transformation plan before the
transformation begins. During this analysis, Amazon Q divides your .NET solution or project
into code groups that you can view in the transformation plan. A _code group_ is a project and all its dependencies that together generate a
buildable unit of code such as a dynamic link library (DLL) or an executable.

During the transformation, Amazon Q provides step-by-step updates in a Transformation Hub
where you can monitor progress. After transforming your application, Amazon Q generates a
summary with the proposed changes in a diff view for you to optionally verify the changes
before you accept them. When you accept the changes, Amazon Q makes in-place updates to your
.NET solution or project.

Amazon Q performs four keys tasks to port .NET applications to Linux:

- **Upgrades language version** – Replaces
outdated C# versions of code with Linux-compatible C# versions.

- **Migrates from .NET Framework to cross-platform**
**.NET** – Migrates projects and packages from Windows dependent
.NET Framework to cross-platform .NET compatible with Linux.

- **Rewrites code for Linux compatibility** –
Refactors and rewrites deprecated and inefficient code components.

- **Generates a Linux compatibility readiness report**– For open-ended tasks where user intervention is needed to make
the code build and run on Linux, Amazon Q provides a detailed report of actions needed
to configure your application after transformation.

For more information about how Amazon Q performs .NET transformations, see
[How it works](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/how-dotnet-transformation-works.html).

###### Topics

- [Quotas](#quotas-dotnet-transformation)

- [Porting a .NET application with Amazon Q Developer in Visual Studio](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/port-dotnet-application.html)

- [How Amazon Q Developer transforms .NET applications](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/how-dotnet-transformation-works.html)

- [Troubleshooting issues with .NET transformations in the IDE](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/troubleshooting-dotnet-transformation-IDE.html)

## Quotas

.NET transformations with Amazon Q in the IDE maintain the following quotas:

- **Lines of code per job** – The
maximum number of code lines that Amazon Q can transform in a given transformation
job. This is also the monthly total limit for .NET transformations.

- **Concurrent jobs** – The maximum number of
transformation jobs you can run at the same time. This quota applies to all
transformations in the IDE, including [Java transformations](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/transform-java.html).

ResourceQuotasLines of code per job100,000 lines of codeConcurrent jobs

1 job per user

2 jobs per AWS account

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Troubleshooting

Porting a .NET application
