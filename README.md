# gitraw
Get raw files from a Github repository without cloning.

How to build `gitraw`
---------------------

    $ git clone https://github.com/yadneshk/gitraw.git
    $ cd gitraw; make build
    $ ls $GOPATH/bin/gitraw

Set `AUTH_GITHUB` environment variable to personal access token
--------------------------------------------------------------

    $ export AUTH_GITHUB=<token>

You can create one from [here](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) if you do not have it.

Usage
-----

List files and directory of a repository

    $ $GOPATH/bin/gitraw list --repository yadneshk/k8s-cluster-deploy --branch main
    .gitignore
    LICENSE
    README.md
    deploy-k8s/
    issues/
    learn-k8s/
    prometheus_k8s/
