#!/bin/bash

# Change to the directory with our code that we plan to work from
cd "$GOPATH/src/github.com/TerrenceHo/simplephotohost.com"

echo "==== Releasing simplephotohost.com ===="
echo "  Deleting the local binary if it exists (so it isn't uploaded)..."
rm simplephotohost.com
echo "  Done!"

echo "  Deleting existing code..."
ssh root@www.simplephotohost.com "rm -rf /root/go/src/github.com/TerrenceHo/simplephotohost.com"
echo "  Code deleted successfully!"

echo "  Uploading code..."
rsync -avr --exclude '.git/*' --exclude 'tmp/*' --exclude 'images/*' ./ root@www.simplephotohost.com:/root/go/src/github.com/TerrenceHo/simplephotohost.com/
echo "  Code uploaded successfully!"

echo "  Go getting deps..."
ssh root@www.simplephotohost.com "export GOPATH=/root/go; /usr/local/go/bin/go get golang.org/x/crypto/bcrypt"
ssh root@www.simplephotohost.com "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/gorilla/mux"
ssh root@www.simplephotohost.com "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/gorilla/schema"
ssh root@www.simplephotohost.com "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/lib/pq"
ssh root@www.simplephotohost.com "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/jinzhu/gorm"
ssh root@www.simplephotohost.com "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/gorilla/csrf"

echo "  Building the code on remote server..."
ssh root@www.simplephotohost.com 'export GOPATH=/root/go; cd /root/app; /usr/local/go/bin/go build -o ./server $GOPATH/src/github.com/TerrenceHo/simplephotohost.com/*.go'
echo "  Code built successfully!"

echo "  Moving assets..."
ssh root@www.simplephotohost.com "cd /root/app; cp -R /root/go/src/github.com/TerrenceHo/simplephotohost.com/assets ."
echo "  Assets moved successfully!"

echo "  Moving views..."
ssh root@www.simplephotohost.com "cd /root/app; cp -R /root/go/src/github.com/TerrenceHo/simplephotohost.com/views ."
echo "  Views moved successfully!"

echo "  Moving Caddyfile..."
ssh root@www.simplephotohost.com "cd /root/app; cp /root/go/src/github.com/TerrenceHo/simplephotohost.com/Caddyfile ."
echo "  Views moved successfully!"

echo "  Moving .config file..."
ssh root@www.simplephotohost.com "cd /root/app; cp /root/go/src/github.com/TerrenceHo/simplephotohost.com/.config ."

echo "  Restarting the server..."
ssh root@www.simplephotohost.com "sudo service simplephotohost.com restart"
echo "  Server restarted successfully!"

echo "  Restarting Caddy server..."
ssh root@www.simplephotohost.com "sudo service caddy restart"
echo "  Caddy restarted successfully!"

echo "==== Done releasing simplephotohost.com ===="

