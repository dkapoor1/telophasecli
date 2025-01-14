<p align="center">
  <a href="https://telophase.dev"><img src="https://github.com/Santiago-Labs/telophasecli/assets/3019043/ff5ed6db-9e91-44e7-9feb-bcf4f608bce8" alt="Logo" height=170></a>
</p>
<h1 align="center">Telophase</h1>
<br/>

## Why Telophase?
Automation and Compliance are key concerns when adopting multi-account AWS. Telophase orchestrates the management of AWS Organizations alongside your infrastructure-as-code (IaC) provider, like Terraform or CDK. Using a single tool for these allows:
1. **Workflow Automation**: Automates account creation and decommissioning, integrating with existing automation workflows, like CI or ServiceNow.
2. **IaC <> Account Binding**: Enables binding accounts to specific IaC stacks for automatic provisioning of baseline resources.
3. **Easier Compliance Deployment**: Enables binding Service Control Policies (SCPs) to accounts as part of your Account provisioning workflow to make sure every Account is compliant. We make it easy to test SCPs before they are deployed.

Currently, Telophase is a CLI tool only. In the future, we plan to offer a web UI.

## Install
Go is the only supported installation method. If you'd like another method, please let us know by opening an issue!
```
go install github.com/santiago-labs/telophasecli@latest
```

## Quick links

- Intro
  - [Quickstart](https://github.com/Santiago-Labs/telophasecli/blob/main/docs/quickstart.md)
- Features
  - [Manage AWS Organization](https://github.com/Santiago-Labs/telophasecli/blob/main/docs/features.md#aws-organization)
  - [Assign IaC to Accounts](https://github.com/Santiago-Labs/telophasecli/blob/main/docs/features.md#assign-iac-stacks-to-accounts)
  - [Pass Outputs Across Stacks](https://github.com/Santiago-Labs/telophasecli/blob/main/docs/features.md#pass-outputs-across-accounts-and-regions-cdk-only)
  - [Terminal UI](https://github.com/Santiago-Labs/telophasecli/blob/main/docs/features.md#terminal-ui)
- CLI
  - [`telophase diff`](https://github.com/Santiago-Labs/telophasecli/blob/main/docs/commands.md#telophasecli-diff)
  - [`telophase deploy`](https://github.com/Santiago-Labs/telophasecli/blob/main/docs/commands.md#telophasecli-deploy)
  - [`telophase account import`](https://github.com/Santiago-Labs/telophasecli/blob/main/docs/commands.md#telophasecli-account-import)
  - [`telophase account diff`](https://github.com/Santiago-Labs/telophasecli/blob/main/docs/commands.md#telophasecli-account-diff)
  - [`telophase account deploy`](https://github.com/Santiago-Labs/telophasecli/blob/main/docs/commands.md#telophasecli-account-deploy)
- Organization.yml Reference
  - [Reference](https://github.com/Santiago-Labs/telophasecli/blob/main/docs/organizationyml.md)


### Future Development
- [ ] Support for multi-cloud organizations with a unified account factory.
  - [ ] Azure
  - [ ] GCP
- [ ] Drift detection/prevention
- [ ] Guardrails around account resources 
- [ ] Guardrails around new Accounts, similar to Control Tower rules.

### Comparisons
#### Telophase vs Control Tower
Manage Accounts via code not a UI. Telophase leaves the controls up to you and your IaC.

#### Telophase vs CDK with multiple environments
Telophase wraps your usage of CDK so that you can apply the cdk to multiple
accounts in parallel. Telophase lets you focus on your actual infrastructure and
not worrying about setting up the right IAM roles for multi account management.
