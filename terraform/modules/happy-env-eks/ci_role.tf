module "happy_github_ci_role" {
  for_each = var.authorized_github_repos
  source   = "../happy-github-ci-role"

  ecr_repo_arns           = flatten([for ecr in module.ecrs : ecr.repository_arn])
  authorized_github_repos = [each.value]
  happy_app_name          = each.value

  tags = var.tags
}
