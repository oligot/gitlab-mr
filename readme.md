# gitlab-mr

> Open the GitLab merge request page for the current branch


## Install

```
> git clone https://github.com/oligot/gitlab-mr.git
> cd gitlab-mr
> go build
```

## Configure

* [create a personal access token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html#creating-a-personal-access-token)
* create a file `$HOME/.config/gitlab-mr/config.json`

```json
{
  "token": "..." 
}
```

## Usage

```
> gitlab-mr
```

This should open a tab in your browser with the merge request page, if one exists for the current branch.
