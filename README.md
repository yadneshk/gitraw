# gitraw
Get files and directories from a Github repository without cloning.

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

Get files from a repository

    $ $GOPATH/bin/gitraw get --repository yadneshk/k8s-cluster-deploy --branch main  --output-dir ~/Downloads/ deploy-k8s/bash-scripts/setup-master.sh
    $ ls ~/Downloads/setup-master.sh 
    /home/ykulkarn/Downloads/setup-master.sh

Get directories from a repository

    $ $GOPATH/bin/gitraw get -r yadneshk/k8s-cluster-deploy -b main -o ~/deploy/ deploy-k8s/playbook/
    /home/ykulkarn/deploy/.gitignore
    /home/ykulkarn/deploy/README.md
    /home/ykulkarn/deploy/ansible.cfg
    /home/ykulkarn/deploy/destroy.yml
    /home/ykulkarn/deploy/files/network.xml.j2
    /home/ykulkarn/deploy/inventory/kvm_host
    /home/ykulkarn/deploy/inventory/prepare_dynamic_inventory.py
    /home/ykulkarn/deploy/populate_hosts.yml
    /home/ykulkarn/deploy/roles/configure_cni/.travis.yml
    /home/ykulkarn/deploy/roles/configure_cni/README.md
    /home/ykulkarn/deploy/roles/configure_cni/defaults/main.yml
    /home/ykulkarn/deploy/roles/configure_cni/handlers/main.yml
    /home/ykulkarn/deploy/roles/configure_cni/meta/main.yml
    /home/ykulkarn/deploy/roles/configure_cni/tasks/main.yml
    /home/ykulkarn/deploy/roles/configure_cni/tests/inventory
    /home/ykulkarn/deploy/roles/configure_cni/tests/test.yml
    /home/ykulkarn/deploy/roles/configure_cni/vars/main.yml
    /home/ykulkarn/deploy/roles/configure_master/.travis.yml
    /home/ykulkarn/deploy/roles/configure_master/README.md
    /home/ykulkarn/deploy/roles/configure_master/defaults/main.yml
    /home/ykulkarn/deploy/roles/configure_master/handlers/main.yml
    /home/ykulkarn/deploy/roles/configure_master/meta/main.yml
    /home/ykulkarn/deploy/roles/configure_master/tasks/main.yml
    /home/ykulkarn/deploy/roles/configure_master/tests/inventory
    /home/ykulkarn/deploy/roles/configure_master/tests/test.yml
    /home/ykulkarn/deploy/roles/configure_master/vars/main.yml
    /home/ykulkarn/deploy/roles/configure_worker/.travis.yml
    /home/ykulkarn/deploy/roles/configure_worker/README.md
    /home/ykulkarn/deploy/roles/configure_worker/defaults/main.yml
    /home/ykulkarn/deploy/roles/configure_worker/handlers/main.yml
    /home/ykulkarn/deploy/roles/configure_worker/meta/main.yml
    /home/ykulkarn/deploy/roles/configure_worker/tasks/main.yml
    /home/ykulkarn/deploy/roles/configure_worker/tests/inventory
    /home/ykulkarn/deploy/roles/configure_worker/tests/test.yml
    /home/ykulkarn/deploy/roles/configure_worker/vars/main.yml
    /home/ykulkarn/deploy/scripts/populate_hosts.py
    /home/ykulkarn/deploy/setup_cluster_nodes.yml
    /home/ykulkarn/deploy/setup_hypervisor.yml
    /home/ykulkarn/deploy/setup_nodes.yml
    /home/ykulkarn/deploy/site.yml
    /home/ykulkarn/deploy/vars/cluster_vars.yml
    /home/ykulkarn/deploy/vars/hypervisor_vars.yml
    /home/ykulkarn/deploy/vars/node_vars.yml
