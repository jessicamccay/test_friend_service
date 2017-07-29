//go:generate bash -c "docker run --rm -v `pwd`/..:/tmp -w /tmp docker.uacf.io/docker_images/grpc:master friends.uacf.io/apps/friendsapi/rpc/friendsapi.proto"

package friends
