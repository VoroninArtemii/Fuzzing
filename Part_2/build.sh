export GOTOOLCHAIN=local
export BUILDDIR="$PWD/.gopath"
export IMPORT_PATH="%import_path"
export GOPATH="$BUILDDIR:%go_path"
export KUBE_GIT_COMMIT=%release
export KUBE_GIT_TREE_STATE="clean"
export KUBE_GIT_VERSION="v%version"

# Fixes https://github.com/golang/go/issues/58425
%ifarch %arm
export CGO_ENABLED=0
%endif

%golang_prepare
# .go-version is needed for successfull build
cp -alv -- .go-version "$BUILDDIR/src/$IMPORT_PATH"
pushd .gopath/src/%import_path
export KUBE_EXTRA_GOPATH=$(pwd)/Godeps/_workspace

GOMAXPROCS=10 make WHAT="${BINS[*]}"

# convert md to man
./hack/update-generated-docs.sh || true
pushd docs
pushd admin
cp kube-apiserver.md kube-controller-manager.md kube-proxy.md kube-scheduler.md kubelet.md ..
popd
cp %SOURCE2 genmanpages.sh
bash genmanpages.sh
popd
popd