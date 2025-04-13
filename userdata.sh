curl -LO https://golang.org/dl/go1.23.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin


curl -LO https://github.com/prometheus/prometheus/releases/download/v3.1.0/prometheus-3.1.0.linux-amd64.tar.gz
tar -xzf prometheus-3.1.0.linux-amd64.tar.gz
mv prometheus-3.1.0.linux-amd64 prometheus
cd prometheus/
./prometheus &