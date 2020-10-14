
if [ "x$1" = "x" ]; then
  echo "You have to specify which action to be excuted. [ rm / start / down / deps ]" 1>&2
  exit 1
fi

# rm
if [ "x$1" = "xrm" ]; then
    rm_container_id=$(docker ps -a -q --filter name=trip-planner)
    docker rm -f ${rm_container_id}
fi

# start
if [ "x$1" = "xstart" ]; then
    cd tools
    docker-compose up
fi

# down
if [ "x$1" = "xdown" ]; then
    cd tools
    docker-compose down
fi

# deps
if [ "x$1" = "xdeps" ]; then
    cd src
    go mod vendor
fi
