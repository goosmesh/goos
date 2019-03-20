#!/usr/bin/env bash

current=`pwd`

# start build static resource

cd ../plugin-goos-ui/ui-goos-hidden
npm i
npm run build

# copy default Corefile to production root
cd ${current}
\cp ./Corefile ../production/

# build production
git clone https://github.com/coredns/coredns.git
cd coredns
\cp ././coredns.go ./
\cp ././plugin.cfg ./

go generate
make