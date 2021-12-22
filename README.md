# gitraw
Get raw files from a Github repository without cloning.

How to build `gitraw`
---------------------

    $ git clone https://github.com/yadneshk/gitraw.git
    $ cd gitraw; make build
    $ ls $GOPATH/bin/gitraw

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
