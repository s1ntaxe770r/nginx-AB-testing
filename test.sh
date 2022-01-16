#! /bin/bash
for i in {1..10}; do
	version=$(curl -sSL cluster-id-60cb32e9bc42.k8s.civo.com/version)
	echo $version
done

