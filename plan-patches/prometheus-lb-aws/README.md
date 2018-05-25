## prometheus-lb-aws

This patch will deploy an aws network load balancer that forwards traffic to the nginx that fronts a prometheus cluster deployed by https://github.com/bosh-prometheus/prometheus-boshrelease.

You'll be able to reach grafana, alertmanager, and prometheus on the load balancer ports 3000, 9093, and 9090 respectively.

1. Do the plan-patch dance:
   ```
   export BBL_SOURCE=${GOPATH}/src/github.com/cloudfoundry/bosh-bootloader/
	 mkdir prometheus-env && cd prometheus-env/
   bbl plan --name prometheus-env
   cp -r ${BBL_SOURCE}/plan-patches/prometheus-lb-aws/. .
   bbl up
   ```
1. Once you've bbl'd up, deploy your prometheus cluster:
   ```
   git clone https://github.com/bosh-prometheus/prometheus-boshrelease.git
   bosh deploy -d prometheus prometheus-boshrelease/manifests/prometheus.yml \
     -o add-nginx-network-properties.yml
   ```

1. Wait a hot minute for your load balancers to find their targets, then log in to grafana via a web browser:
   ```
   open "http://$(bbl outputs | bosh int --path=/prometheus_lb_url -):3000"
   ```

1. check out this PR to upstream the add-nginx-network-properties opsfile: https://github.com/bosh-prometheus/prometheus-boshrelease/pull/203