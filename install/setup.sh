#!/bin/sh

echo "Cloning atomix helm-charts repository"
git clone https://github.com/atomix/helm-charts.git

echo "Installing atomix-controller (default namespace)"
cd helm-charts/atomix-controller
helm install atomix-controller .

echo "Installing raft-storage-controller (default namespace)"
cd ../raft-storage-controller
helm install raft-storage-controller .

echo "Installing raft-database (default namespace)"
cd ../raft-database
helm install raft-database .

echo "Cloning req-smr repository"
cd ../../
git clone https://github.com/lucasbfernandes/req-smr.git

echo "Installing req-smr (default namespace)"
cd req-smr/install/helm-chart
helm install req-smr .

echo "Done."
