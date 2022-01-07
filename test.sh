#! /bin/bash
for i in {1..10}; do
	version=$(curl -sSL http://0703b49c-6f02-4010-b27a-98380c244878.k8s.civo.com/version)
	echo $version
done

